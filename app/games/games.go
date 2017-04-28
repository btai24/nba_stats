package games

import "time"

type Header struct {
	GameDate       time.Time // TODO: turn into a date type
	GameSequence   int
	GameID         string
	GameStatusID   int
	GameStatusText string
	GameCode       string
	HomeTeamID     int
	VisitorTeamID  int
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
	GameDate     time.Time
	GameSequence int
	GameID       string
	TeamID       int
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

type LastMeeting struct {
	GameID       string
	LastGameID   string
	LastGameDate time.Time
	HomeID       int
	HomeCity     string
	HomeName     string
	HomeAbv      string
	HomePts      int
	VisitorID    int
	VisitorCity  string
	VisitorName  string
	VisitorAbv   string
	VisitorPts   int
}

type TeamLeaders struct {
	GameID        string
	TeamID        int
	TeamCity      string
	TeamName      string
	TeamAbv       string
	PtsPlayerID   int
	PtsPlayerName string
	Pts           int
	RebPlayerID   int
	RebPlayerName string
	Reb           int
	AstPlayerID   int
	AstPlayerName string
	Ast           int
}
