package main

import (
	"log"
	"net/http"
	"odcserver/adapter/https"
	"odcserver/domain"
	"odcserver/domain/models"
	"odcserver/adapter/sqlite"
)

func main() {
	// Use this admin profile if bootstrapping server
	adminProfile := models.ApiProfile{
		ApiKey: "<admin-api-key-here>",
		SlackHookUrl: "<slack-hook-url-here>",
		Email: "aden@odc.com",
		UsageCount: 0,
		Created: "2024-09-12",
		LastUpdated: "",
		PrivilegeLevel: 100,
	}

	db, err := sqlite.Initialise(&adminProfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	apiRepository := sqlite.SqlController{db}
	slackRepository := https.SlackController{}
	commandHandler := domain.CommandHandler{
		ApiRepository: apiRepository,
		SlackRepository: slackRepository,
	}

	publicController := https.PublicController{CommandHandler: commandHandler}
	publicController.HandleRequests()
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}