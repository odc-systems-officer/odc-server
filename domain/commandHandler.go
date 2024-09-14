package domain

import (
	"odcserver/domain/commands"
 	"odcserver/domain/persistence"
)

type CommandHandler struct {
	ApiRepository persistence.ApiRepository
	SlackRepository persistence.SlackRepository
}

func (commandHandler *CommandHandler) SendEquipmentSlackMessage(command commands.EquipmentRequestCommand) error  {
	profile, err := commandHandler.ApiRepository.GetApiProfile(command.ApiKey)
	if err != nil {
		return err
	}
	if profile.PrivilegeLevel > 1 {
		commandHandler.SlackRepository.SendEquipmentRequestMessage(command)
	}
	return nil
}

