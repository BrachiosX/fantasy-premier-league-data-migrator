package models

type BootStrapData struct {
	GameSettings Setting           `json:"game_settings"`
	Players      []Player          `json:"elements"`
	Teams        []Team            `json:"teams"`
	Positions    []Position        `json:"element_types"`
	Stats        []PlayerStatTypes `json:"element_stats"`
}
