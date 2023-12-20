package models

type PlayerDetail struct {
	Fixtures []UpcomingFixture
	History  []PlayerMatchHisoricalStat
	Seasons  []PlayerPastSeasonStat
}
