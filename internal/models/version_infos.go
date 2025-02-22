package models

import "encoding/json"

// LatestPolicy JSON 구조체
type LatestPolicy struct {
	VersionInfo VersionInfo `json:"version_info"`
	UpdateMessage []UpdateMessageData `json:"update_message"`
}

type VersionInfo struct {
	StoreType 	string `json:"store_type"`
	VersionNo 	string `json:"version_no"`
	GameServerName 	string `json:"game_server_name"`
	GameServerUrl 	string `json:"game_server_url"`
	VisiblePopup 	string `json:"visible_popup"`
	LatestVersion 	string `json:"latest_version"`
	DynamicConfig	string `json:"dynamic_config"`
}

type UpdateMessageData struct {
	LanguageType string `json:"language_type"`
	Title        string `json:"title"`
	Body         string `json:"body"`
}