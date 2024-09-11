package main

import (
	"log"
	"net/http"
	"odcserver/adapter/https"
	"odcserver/domain"
	// "odcserver/domain/persistence"
)

func main() {
	// apiRepository := https.ApiRepository{}
	slackRepository := https.SlackController{}
	commandHandler := domain.CommandHandler{
		// apiRepository: apiRepository,
		SlackRepository: slackRepository,
	}

	publicController := https.PublicController{CommandHandler: commandHandler}
	publicController.HandleRequests()
	// https.SendSlackMessage("Hello from Aden :)")
	// https.HandleRequests()
	log.Fatal(http.ListenAndServe(":8080", nil))
}