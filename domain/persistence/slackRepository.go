package persistence

type SlackRepository interface {
	SendEquipmentRequestMessage(message string, channel string) error
}