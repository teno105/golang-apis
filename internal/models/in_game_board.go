package models

import "encoding/json"

// InGameBoard JSON 구조체
type InGameBoard struct {
	Display	bool `json:"display"`
	Url 	string `json:"url"`
}