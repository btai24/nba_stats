package games

import "time"

type GameHeader struct {
	GameDateEst    time.Time
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
