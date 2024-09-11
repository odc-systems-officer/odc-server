package domain

type ApiKeyProfile struct {
	OwnerEmail string `json:"ownerEmail"`
	UsageCount int `json:"usageCount"`
	Created string `json:"created"`
	LastUpdated string `json:"lastUpdated"`
	Privileges []string `json:"privileges"`
}
