package models

type Position struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SingularNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_play"`
	SquadMaxPlay       int    `json:"squad_max_play"`
	UiShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
	ElementCount       int    `json:"element_count"`
}
