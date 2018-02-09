package standings

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/btai24/nba_stats/api"
	"github.com/btai24/nba_stats/deps"
	"github.com/gin-gonic/gin"
)

type Series struct {
	GameID       string
	HomeID       int
	VisitorID    int
	GameDate     time.Time
	HomeWins     int
	HomeLosses   int
	SeriesLeader string
}

type ConferenceStandings struct {
	East [15]Team `json:"east"`
	West [15]Team `json:"west"`
}

type Team struct {
	TeamID         string  `json:"id"`
	Wins           int     `json:"wins"`
	Losses         int     `json:"losses"`
	GamesBehind    float64 `json:"games_behind"`
	WinPct         float64 `json:"win_percentage"`
	ConferenceRank int     `json:"conference_rank"`
	Streak         int     `json:"streak"`
	Record         string  `json:"record"`
	HomeRecord     string  `json:"home_record"`
	AwayRecord     string  `json:"away_record"`
	LastTen        string  `json:"last_ten"`
}

func GetConferenceStandings(c *gin.Context) {
	// TODO: handle error
	standings, _ := api.GetConferenceStandings(c)
	var eastStandings [15]Team
	var westStandings [15]Team

	// TODO: handle error
	redisClient, _ := deps.RedisClient(c)

	for i, team := range standings.Conference.East {
		eastStandings[i] = formTeam(team)
	}

	for i, team := range standings.Conference.West {
		westStandings[i] = formTeam(team)
	}

	confStandings := ConferenceStandings{
		East: eastStandings,
		West: westStandings,
	}

	json, _ := json.Marshal(confStandings)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(json)
}

// type StandingsTeam struct {
// 	Id             string `json:"teamId"`
// 	Wins           string `json:"win"`
// 	Losses         string `json:"loss"`
// 	WinPct         string `json:"winPct"`
// 	GamesBehind    string `json:"gamesBehind"`
// 	ConferenceRank string `json:"confRank"`
// 	Streak         string `json:"streak"`
// 	IsWinStreak    bool   `json:"isWinStreak"`
// 	HomeWins       string `json:"homeWin"`
// 	HomeLosses     string `json:"homeLoss"`
// 	AwayWins       string `json:"awayWin"`
// 	AwayLosses     string `json:"awayLoss"`
// 	LastTenWins    string `json:"lastTenWin"`
// 	LastTenLosses  string `json:"lastTenLoss"`
// }

func formTeam(team api.StandingsTeam) Team {
	wins, _ := strconv.Atoi(team.Wins)
	losses, _ := strconv.Atoi(team.Losses)
	winPct, _ := strconv.ParseFloat(team.WinPct, 32)
	gamesBehind, _ := strconv.ParseFloat(team.GamesBehind, 32)
	confRank, _ := strconv.Atoi(team.ConferenceRank)

	streak, _ := strconv.Atoi(team.Streak)
	if !team.IsWinStreak {
		streak = streak * -1
	}

	record := fmt.Sprintf("%s-%s", team.Wins, team.Losses)
	homeRecord := fmt.Sprintf("%s-%s", team.HomeWins, team.HomeLosses)
	awayRecord := fmt.Sprintf("%s-%s", team.AwayWins, team.AwayLosses)
	lastTen := fmt.Sprintf("%s-%s", team.LastTenWins, team.LastTenLosses)

	return Team{
		TeamID:         team.Id,
		Wins:           wins,
		Losses:         losses,
		WinPct:         winPct,
		GamesBehind:    gamesBehind,
		ConferenceRank: confRank,
		Streak:         streak,
		Record:         record,
		HomeRecord:     homeRecord,
		AwayRecord:     awayRecord,
		LastTen:        lastTen,
	}
}
