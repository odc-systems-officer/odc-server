package sqlite

import (
	"database/sql"
	"log"
	"os"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"odcserver/domain/models"
	"odcserver/domain/models/exceptions"
)

type SqlController struct {
	Db *sql.DB
}

func (sqlController SqlController) GetApiProfile(apiKey string) (*models.ApiProfile, error) {	
	row, err := sqlController.Db.Query("SELECT * FROM apikeys WHERE apiKey = ?", apiKey)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var apiKey string
		var slackHookUrl string
		var email string
		var usageCount int
		var created string
		var lastUpdated string
		var privilegeLevel int

		row.Scan(&apiKey, &slackHookUrl, &email, &usageCount, &created, &lastUpdated, &privilegeLevel)
		profile := models.ApiProfile{
			ApiKey: apiKey,
			SlackHookUrl: slackHookUrl,
			Email: email,
			UsageCount: usageCount,
			Created: created,
			LastUpdated: lastUpdated,
			PrivilegeLevel: privilegeLevel,
		}
		return &profile, nil
	}
	return nil, exceptions.ErrApiKeyNotFound
}

func (sqlController SqlController) SaveApiProfile(apiKey string, profile *models.ApiProfile) error {	
	err := upsertApiProfile(sqlController.Db, profile)	
	return err
}

func Initialise(adminProfile *models.ApiProfile) (*sql.DB, error) {
	os.Remove("odc.db") // Delete the db everytime for now
	var initMode = false
	// If database file does not exist, create database
	if _, err := os.Stat("odc.db"); errors.Is(err, os.ErrNotExist) {
		initMode = true
		// Create database
		log.Println("Creating odc.db...")
		file, err := os.Create("odc.db") 
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("odc.db created")
	}
	// Connect to database
	db, _ := sql.Open("sqlite3", "./odc.db") 

	// Create tables
	if initMode {
		createApiKeysTable(db)
		upsertApiProfile(db, adminProfile)
	}

	return db, nil
}

func createApiKeysTable(db *sql.DB) {
	createApiKeysTableSQL := `CREATE TABLE apikeys (
		"apiKey" TEXT NOT NULL PRIMARY KEY,
		"slackHookUrl" TEXT,
		"email" TEXT,
		"usageCount" INTEGER,
		"created" TEXT,
		"lastUpdated" TEXT,
		"privilegeLevel" INTEGER
	  );` // SQL Statement for Create Table

	log.Println("Create apiKeys table...")
	statement, err := db.Prepare(createApiKeysTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("apiKeys table created")
}

func upsertApiProfile(db *sql.DB, profile *models.ApiProfile) error {
	log.Println("Upserting apiProfile record ...")
	upsertApiProfileSQL := `INSERT INTO apiKeys (apiKey, slackHookUrl, email, usageCount, created, lastUpdated, privilegeLevel) VALUES (?, ?, ?, ?, ?, ?, ?) ON CONFLICT (apiKey) DO UPDATE SET slackHookUrl = ?, email = ?, usageCount = ?, created = ?, lastUpdated = ?, privilegeLevel = ?;`
	statement, err := db.Prepare(upsertApiProfileSQL) // Prepare statement (this is good to avoid SQL injections)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		&profile.ApiKey, 
		&profile.SlackHookUrl,
		&profile.Email, 
		&profile.UsageCount, 
		&profile.Created, 
		&profile.LastUpdated, 
		&profile.PrivilegeLevel, 
		// On conflict
		&profile.SlackHookUrl,
		&profile.Email, 
		&profile.UsageCount, 
		&profile.Created, 
		&profile.LastUpdated, 
		&profile.PrivilegeLevel)

	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}