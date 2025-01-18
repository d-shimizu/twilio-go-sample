package model

type PhoneNumber struct {
	AccountSid  string `json:"account_sid"`
	PhoneNumber string `json:"phone_number"`
}

type AvailablePhoneNumber struct {
	PhoneNumber  string `json:"phone_number"`
	FriendlyName string `json:"friendly_name"`
	Locality     string `json:"locality,omitempty"`
	Region       string `json:"region,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	Capabilities struct {
		Voice bool `json:"voice"`
		SMS   bool `json:"sms"`
		MMS   bool `json:"mms"`
	} `json:"capabilities"`
}
