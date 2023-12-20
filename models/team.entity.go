package models

type Team struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	ShortName           string `json:"short_name"`
	Strength            int    `json:"strength"`
	StrengthOverallHome int    `json:"strength_overall_home"`
	StrengthOverallAway int    `json:"strength_overall_away"`
	StrengthAttackHome  int    `json:"strength_attack_home"`
	StrengthAttackAway  int    `json:"strength_attack_away"`
	StrengthDefenceHome int    `json:"strength_defence_home"`
	StrengthDefenceAway int    `json:"strength_defence_away"`
	PulseID             int    `json:"pulse_id"`
	Form                string `json:"form"`
	Played              int    `json:"played"`
	Win                 int    `json:"win"`
	Draw                int    `json:"draw"`
	Loss                int    `json:"loss"`
	Points              int    `json:"points"`
	Code                int    `json:"code"`
	Position            int    `json:"position"`
	Unavailable         bool   `json:"unavailable"`
}
