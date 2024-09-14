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
		var email string
		var usageCount int
		var created string
		var lastUpdated string
		var privilegeLevel int
		row.Scan(&apiKey, &email, &usageCount, &created, &lastUpdated, &privilegeLevel)
		log.Println("apiKey: ", apiKey, " email: ", email, " usageCount: ", usageCount, " created: ", created, " lastUpdated: ", lastUpdated, " privilegeLevel: ", privilegeLevel)
		profile := models.ApiProfile{
			ApiKey: apiKey,
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

func Initialise() (*sql.DB, error) {
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

		profile := models.ApiProfile{
			ApiKey: "<admin-api-key-here>",
			Email: "aden@odc.com",
			UsageCount: 0,
			Created: "2024-09-12",
			LastUpdated: "",
			PrivilegeLevel: 100,
		}
		saveApiProfile(db, profile)
	}

	return db, nil
}

func createApiKeysTable(db *sql.DB) {
	createApiKeysTableSQL := `CREATE TABLE apikeys (
		"apiKey" TEXT NOT NULL PRIMARY KEY,
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

func saveApiProfile(db *sql.DB, profile models.ApiProfile) error {
	log.Println("Inserting apiProfile record ...")
	insertApiProfileSQL := `INSERT INTO apiKeys (apiKey, email, usageCount, created, lastUpdated, privilegeLevel) VALUES (?, ?, ?, ?, ?, ?);`
	statement, err := db.Prepare(insertApiProfileSQL) // Prepare statement (this is good to avoid SQL injections)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(profile.ApiKey, profile.Email, profile.UsageCount, profile.Created, profile.LastUpdated, profile.PrivilegeLevel)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Inserted.")
	return nil
}