package main

import (
	"log"
	"net/http"
	"odcserver/adapter/https"
	"odcserver/domain"
	"odcserver/adapter/sqlite"
	// "odcserver/domain/persistence"
)

func main() {
	db, err := sqlite.Initialise()
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}