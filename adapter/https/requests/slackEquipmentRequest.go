package requests

type SlackEquipmentRequest struct {
	Email string `json:"email"`
	Equipment []string `json:"equipment"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	Message string `json:"message"`
}