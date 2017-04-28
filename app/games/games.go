package games

import "time"

type Header struct {
	GameDateEst    time.Time // TODO: turn into a date type
	GameSequence   int
	GameID         string
	GameStatusID   int
	GameStatusText string
	GameCode       string
	HomeTeamID     string
	VisitorTeamID  string
	Season         string
	LivePeriod     int
	LivePcTime     string
	// TODO: Nullable fields
	// NatlTVBroadcater string
	// HomeTVBroadcater string
	// AwayTVBroadcater string
	// LivePeriodTimeBcast string
	// WhStatus            int
}

type Scoreboard struct {
	GameDateEst  time.Time
	GameSequence int
	GameID       string
	TeamID       string
	TeamAbv      string
	TeamCity     string
	TeamWL       string
	Q1           int
	Q2           int
	Q3           int
	Q4           int
	OT1          int
	OT2          int
	OT3          int
	OT4          int
	OT5          int
	OT6          int
	OT7          int
	OT8          int
	OT9          int
	OT10         int
	PTS          int
	FG           float64
	FT           float64
	FG3          float64
	AST          int
	REB          int
	TOV          int
}
