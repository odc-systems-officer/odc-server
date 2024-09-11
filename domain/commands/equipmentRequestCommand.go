package commands

type EquipmentRequestCommand struct {
	ApiKey string
	Email string
	Equipment []string
	StartDate string
	EndDate string
	Message string
	Channel string
}
