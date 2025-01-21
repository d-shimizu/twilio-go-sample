package model

type VoiceRequest struct {
	From string `json:"From"`
	To   string `json:"To"`
	//Body string `json:"Body"`
}
