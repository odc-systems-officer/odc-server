package https

import (
	// "fmt"
	// "log"
	"net/http"
	// "encoding/json"
	"odcserver/domain"
	"odcserver/domain/commands"
)

type PublicController struct {
	CommandHandler domain.CommandHandler
}

func (publicController *PublicController)  HandleRequests() {
	http.HandleFunc("/api/slack/equipmentRequest", publicController.slackEquipmentRequestHandle)
}

func (publicController *PublicController) slackEquipmentRequestHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// apiKey := r.Header.Get("Api-Key")

			// body, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	http.Error(w, "Error reading request body", http.StatusInternalServerError)
			// 	return
			// }
			// defer r.Body.Close()

			// // Process the body (e.g., unmarshal JSON)
			// var data map[string]interface{}
			// err = json.Unmarshal(body, &data)
			// if err != nil {
			// 	http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			// 	return
			// }
			// command := toEquipmentSlackMessageCommand(apiKey, data)
			command := commands.EquipmentRequestCommand{
				ApiKey: "1234",
				Email: "aden@odc.com",
				Equipment: []string{"equipment1", "equipment2"},
				StartDate: "2023-01-01",
				EndDate: "2023-01-31",
				Message: "This is a test message",
				Channel: "#odc-equipment",
			}
			publicController.CommandHandler.SendEquipmentSlackMessage(command)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func toEquipmentSlackMessageCommand(apiKey string, data map[string]interface{}) commands.EquipmentRequestCommand {
	return commands.EquipmentRequestCommand{
		ApiKey: apiKey,
		Email: data["email"].(string),
		Equipment: data["equipment"].([]string),
		StartDate: data["startDate"].(string),
		EndDate: data["endDate"].(string),
		Message: data["message"].(string),
		Channel: data["channel"].(string),
	}
}