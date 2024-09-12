package domain

import "odcserver/domain/commands"
import "odcserver/domain/persistence"

type CommandHandler struct {
	ApiRepository persistence.ApiRepository
	SlackRepository persistence.SlackRepository
}

func (commandHandler *CommandHandler) SendEquipmentSlackMessage(command commands.EquipmentRequestCommand) error  {
	commandHandler.ApiRepository.GetApiProfile(command.ApiKey)
	// commandHandler.SlackRepository.SendEquipmentRequestMessage(command.Message, command.Channel)
	return nil
}

