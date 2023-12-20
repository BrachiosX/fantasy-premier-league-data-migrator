package models

type Setting struct {
	LeagueJoinPrivateMax         int      `json:"league_join_private_max"`
	LeagueJoinPublicMax          int      `json:"league_join_public_max"`
	LeagueMaxSizePublicClassic   int      `json:"league_max_size_public_classic"`
	LeagueMaxSizePublicH2H       int      `json:"league_max_size_public_h2h"`
	LeagueMaxSizePrivateH2H      int      `json:"league_max_size_private_h2h"`
	LeagueMaxKoRoundsPrivateH2H  int      `json:"league_max_ko_rounds_private_h2h"`
	LeaguePrefixPublic           string   `json:"league_prefix_public"`
	LeaguePointsH2HWin           int      `json:"league_points_h2h_win"`
	LeaguePointsH2HLose          int      `json:"league_points_h2h_lose"`
	LeaguePointsH2HDraw          int      `json:"league_points_h2h_draw"`
	LeagueKoFirstInsteadOfRandom bool     `json:"league_ko_first_instead_of_random"`
	CupStartEventID              int      `json:"cup_start_event_id"`
	CupStopEventID               int      `json:"cup_stop_event_id"`
	CupQualifyingMethod          string   `json:"cup_qualifying_method"`
	CupType                      string   `json:"cup_type"`
	SquadSquadplay               int      `json:"squad_squadplay"`
	SquadSquadsize               int      `json:"squad_squadsize"`
	SquadTeamLimit               int      `json:"squad_team_limit"`
	SquadTotalSpend              int      `json:"squad_total_spend"`
	UiCurrencyMultiplier         int      `json:"ui_currency_multiplier"`
	UiUseSpecialShirts           bool     `json:"ui_use_special_shirts"`
	UiSpecialShirtExclusions     []string `json:"ui_special_shirt_exclusions"`
	StatsFormDays                int      `json:"stats_form_days"`
	SysViceCaptainEnabled        bool     `json:"sys_vice_captain_enabled"`
	TransfersCap                 int      `json:"transfers_cap"`
	TransfersSellOnFee           float64  `json:"transfers_sell_on_fee"`
	LeagueH2HTiebreakStats       []string `json:"league_h2h_tiebreak_stats"`
	Timezone                     string   `json:"timezone"`
}
