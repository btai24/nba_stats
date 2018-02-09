package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Standings struct {
	League League `json:"league"`
}

type League struct {
	Standard ConferenceStandings `json:"standard"`
}

type ConferenceStandings struct {
	StartingYear int        `json:"seasonYear"`
	Conference   Conference `json:"conference"`
}

type Conference struct {
	East []StandingsTeam `json:"east"`
	West []StandingsTeam `json:"west"`
}

type StandingsTeam struct {
	Id             string `json:"teamId"`
	Wins           string `json:"win"`
	Losses         string `json:"loss"`
	WinPct         string `json:"winPct"`
	GamesBehind    string `json:"gamesBehind"`
	ConferenceRank string `json:"confRank"`
	Streak         string `json:"streak"`
	IsWinStreak    bool   `json:"isWinStreak"`
	HomeWins       string `json:"homeWin"`
	HomeLosses     string `json:"homeLoss"`
	AwayWins       string `json:"awayWin"`
	AwayLosses     string `json:"awayLoss"`
	LastTenWins    string `json:"lastTenWin"`
	LastTenLosses  string `json:"lastTenLoss"`
}

func GetConferenceStandings(c *gin.Context) (*ConferenceStandings, error) {
	endpoint := fmt.Sprintf("%s/current/standings_conference.json", c.GetString("NBA_ENDPOINT"))
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", endpoint, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var standings Standings
	json.NewDecoder(res.Body).Decode(&standings)

	return &standings.League.Standard, nil
}
