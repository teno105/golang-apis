package models

// UpdateMessage JSON 구조체
type UpdateMessage struct {
	UpdateMessage []UpdateMessageData `json:"update_message"`
}

type UpdateMessageData struct {
	LanguageType string `json:"language_type"`
	Title        string `json:"title"`
	Body         string `json:"body"`
}
