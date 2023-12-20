package models

type Fixture struct {
	Id                   int           `json:"id"`
	Code                 int           `json:"code"`
	PulseId              int           `json:"pulse_id"`
	TeamH                int           `json:"team_h"`
	TeamHScore           *int          `json:"team_h_score"`
	TeamHDifficult       int           `json:"team_h_difficulty"`
	TeamA                int           `json:"team_a"`
	TeamAScore           *int          `json:"team_a_score"`
	TeamADifficult       int           `json:"team_a_difficulty"`
	Event                int           `json:"event"`
	Finished             bool          `json:"finished"`
	Minutes              int           `json:"minutes"`
	ProvisionalStartTime bool          `json:"provisional_start_time"`
	KickoffTime          string        `json:"kickoff_time"`
	Stats                []FixtureStat `json:"stats"`
}

type FixtureStat struct {
	Identifier string               `json:"identifier"`
	A          []FixtureStatElement `json:"a"`
	H          []FixtureStatElement `json:"h"`
}

type FixtureStatElement struct {
	Value   int `json:"value"`
	Element int `json:"element"`
}

type UpcomingFixture struct {
	Id                   int    `json:"id"`
	Code                 int    `json:"code"`
	TeamH                int    `json:"team_h"`
	TeamHScore           *int   `json:"team_h_score"`
	TeamA                int    `json:"team_a"`
	TeamAScore           *int   `json:"team_a_score"`
	Event                int    `json:"event"`
	Finished             bool   `json:"finished"`
	Minutes              int    `json:"minutes"`
	ProvisionalStartTime bool   `json:"provisional_start_time"`
	KickoffTime          string `json:"kickoff_time"`
	EventName            string `json:"event_name"`
	IsHome               bool   `json:"is_home"`
	Difficulty           int    `json:"difficulty"`
}

type PlayerMatchHisoricalStat struct {
	Element                  int    `json:"element"`
	Fixture                  int    `json:"fixture"`
	Team                     int    `json:"team"`
	TotalPoints              int    `json:"total_points"`
	MinutesPlayed            int    `json:"minutes"`
	GoalsScored              int    `json:"goals_scored"`
	Assists                  int    `json:"assists"`
	CleanSheets              int    `json:"clean_sheets"`
	GoalsConceded            int    `json:"goals_conceded"`
	OwnGoals                 int    `json:"own_goals"`
	PenaltiesSaved           int    `json:"penalties_saved"`
	PenaltiesMissed          int    `json:"penalties_missed"`
	YellowCards              int    `json:"yellow_cards"`
	RedCards                 int    `json:"red_cards"`
	Saves                    int    `json:"saves"`
	Bonus                    int    `json:"bonus"`
	Bps                      int    `json:"bps"`
	Influence                string `json:"influence"`
	Creativity               string `json:"creativity"`
	Threat                   string `json:"threat"`
	IctIndex                 string `json:"ict_index"`
	Starts                   int    `json:"starts"`
	ExpectedGoals            string `json:"expected_goals"`
	ExpectedAssists          string `json:"expected_assists"`
	ExpectedGoalInvolvements string `json:"expected_goal_involvements"`
	ExpectedGoalsConceded    string `json:"expected_goals_conceded"`
	Value                    int    `json:"value"`
	TransfersBalance         int    `json:"transfers_balance"`
	Selected                 int    `json:"selected"`
	TransfersIn              int    `json:"transfers_in"`
	TransfersOut             int    `json:"transfers_out"`
}

type PlayerPastSeasonStat struct {
	SeasonName               string `json:"season_name"`
	ElementCode              int    `json:"element_code"`
	StartCost                int    `json:"start_cost"`
	EndCost                  int    `json:"end_cost"`
	TotalPoints              int    `json:"total_points"`
	Minutes                  int    `json:"minutes"`
	GoalsScored              int    `json:"goals_scored"`
	Assists                  int    `json:"assists"`
	CleanSheets              int    `json:"clean_sheets"`
	GoalsConceded            int    `json:"goals_conceded"`
	OwnGoals                 int    `json:"own_goals"`
	PenaltiesSaved           int    `json:"penalties_saved"`
	PenaltiesMissed          int    `json:"penalties_missed"`
	YellowCards              int    `json:"yellow_cards"`
	RedCards                 int    `json:"red_cards"`
	Saves                    int    `json:"saves"`
	Bonus                    int    `json:"bonus"`
	Bps                      int    `json:"bps"`
	Influence                string `json:"influence"`
	Creativity               string `json:"creativity"`
	Threat                   string `json:"threat"`
	IctIndex                 string `json:"ict_index"`
	Starts                   int    `json:"starts"`
	ExpectedGoals            string `json:"expected_goals"`
	ExpectedAssists          string `json:"expected_assists"`
	ExpectedGoalInvolvements string `json:"expected_goal_involvements"`
	ExpectedGoalsConceded    string `json:"expected_goals_conceded"`
}
