package standings

import "time"

type Series struct {
	GameID       string
	HomeID       int
	VisitorID    int
	GameDate     time.Time
	HomeWins     int
	HomeLosses   int
	SeriesLeader string
}

type ConferenceByDay struct {
	Date       time.Time
	Conference string
	Teams      [15]Team
}

type Team struct {
	TeamID     int
	SeasonID   string
	City       string
	Win        int
	Loss       int
	WinPct     float64
	HomeRecord string
	RoadRecord string
}
