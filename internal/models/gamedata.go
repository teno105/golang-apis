package models

import "encoding/json"

// GameData 구조체 (Maintenance + Notice 포함)
type GameData struct {
	UpdateMessage UpdateMessage `json:"update_message"`
	Maintenance Maintenance `json:"maintenance"`
	Notice      Notice      `json:"notice"`
}