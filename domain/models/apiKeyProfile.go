package models

type ApiProfile struct {
	ApiKey string `json:"apiKey"`
	SlackHookUrl string `json:"slackHookUrl"`
	Email string `json:"email"`
	UsageCount int `json:"usageCount"`
	Created string `json:"created"`
	LastUpdated string `json:"lastUpdated"`
	PrivilegeLevel int `json:"privilegeLevel"`
}
