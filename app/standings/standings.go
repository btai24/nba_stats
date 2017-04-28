package standings

import "time"

type Series struct {
	GameID       string
	HomeID       int
	VisitorID    int
	GameDateEST  time.Time
	HomeWins     int
	HomeLosses   int
	SeriesLeader string
}
