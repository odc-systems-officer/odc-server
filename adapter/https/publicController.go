package https

import (
	"errors"
	"io"
	"net/http"
	"encoding/json"
	"odcserver/domain"
	"odcserver/domain/commands"
	"odcserver/adapter/https/requests"
	"odcserver/domain/models/exceptions"
)

type PublicController struct {
	CommandHandler domain.CommandHandler
}

func (publicController *PublicController)  HandleRequests() {
	http.HandleFunc("/api/notifyEquipmentRequest", publicController.slackEquipmentRequestHandle)
}

func (publicController *PublicController) slackEquipmentRequestHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			apiKey := r.Header.Get("Api-Key")
			if apiKey == "" {
				http.Error(w, exceptions.ApiKeyRequired, http.StatusUnauthorized)
				return
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			// Process the body (e.g., unmarshal JSON)
			var data requests.SlackEquipmentRequest
			err = json.Unmarshal(body, &data)
			if err != nil {
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}
			command := toEquipmentSlackMessageCommand(apiKey, &data)

			// Send Slack Message
			err = publicController.CommandHandler.SendEquipmentSlackMessage(command)
			if err != nil {
				if errors.Is(err, exceptions.ErrApiKeyNotFound) {
					http.Error(w, "API Key not found", http.StatusUnauthorized)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				return
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func toEquipmentSlackMessageCommand(apiKey string, r *requests.SlackEquipmentRequest) commands.EquipmentRequestCommand {
	return commands.EquipmentRequestCommand{
		ApiKey: apiKey,
		Email: r.Email,
		Equipment: r.Equipment,
		StartDate: r.StartDate,
		EndDate: r.EndDate,
		Message: r.Message,
	}
}