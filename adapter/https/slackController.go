package https

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
  // "odcserver/domain/persistence"
)

type SlackController struct {}

func (slackController SlackController) SendEquipmentRequestMessage(message string, channel string) error {
	webhookUrl := "<add webhook url here>"
    attachment1 := slack.Attachment {}
    attachment1.AddField(slack.Field { Title: "Author", Value: "Aden Huen" }).AddField(slack.Field { Title: "Status", Value: "Completed" })
    attachment1.AddAction(slack.Action { Type: "button", Text: "Test Action", Url: "https://google.com", Style: "primary" })
    attachment1.AddAction(slack.Action { Type: "button", Text: "Cancel", Url: "https://google.com", Style: "danger" })
    payload := slack.Payload {
      Text: message,
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