package models

import "encoding/json"

// GameData 구조체
type GameData struct {
	InGameBoard InGameBoard `json:"in_game_board"`
	LatestPolicy LatestPolicy `json:"latest_policy"`
	VersionInfos VersionInfos `json:"version_infos"`
	Notice      Notice      `json:"notice"`
	Maintenance Maintenance `json:"maintenance"`
	StoreLink   StoreLink   `json:"store_link"`
}