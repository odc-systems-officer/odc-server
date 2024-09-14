package https

import (
	"fmt"
  "strings"
	"github.com/ashwanthkumar/slack-go-webhook"
  "odcserver/domain/commands"
)

type SlackController struct {}

func (slackController SlackController) SendEquipmentRequestMessage(command commands.EquipmentRequestCommand, slackHookUrl string) error {
	webhookUrl := slackHookUrl
    attachment1 := slack.Attachment {}
    attachment1.AddField(slack.Field { Title: "Email", Value: command.Email })
    attachment1.AddField(slack.Field { Title: "Equipment", Value: strings.Join(command.Equipment, ", ") })
    attachment1.AddField(slack.Field { Title: "StartDate", Value: command.StartDate })
    attachment1.AddField(slack.Field { Title: "EndDate", Value: command.EndDate })
    attachment1.AddField(slack.Field { Title: "Comments", Value: command.Message })
    
    payload := slack.Payload {
      Text: "A new equipment request has been received.",
      Username: "odc-bot",
      Channel: "#odc-equipment",
      IconEmoji: ":monkey_face:",
      Attachments: []slack.Attachment{attachment1},
    }
    err := slack.Send(webhookUrl, "", payload)
    if len(err) > 0 {
      fmt.Printf("error: %s\n", err)
    }
    return nil
}