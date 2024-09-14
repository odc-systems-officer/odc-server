package persistence

import "odcserver/domain/commands"

type SlackRepository interface {
	SendEquipmentRequestMessage(command commands.EquipmentRequestCommand, slackHookUrl string) error
}