package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Scoreboard struct {
	NumGames int    `json:"numGames"`
	Games    []Game `json:"games"`
}

type Game struct {
	Id               string   `json:"gameId"`
	Arena            Arena    `json:"arena"`
	StartTimeEastern string   `json:"startTimeEastern"`
	IsGameOngoing    bool     `json:"isGameActivated"`
	Clock            string   `json:"clock"`
	Attendance       string   `json:"attendance"`
	Duration         Duration `json:gameDuration`
	Period           Period   `json:"Period"`
	VisitingTeam     Team     `json:"vTeam"`
	HomeTeam         Team     `json:"hTeam"`
}

type Arena struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Duration struct {
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
}

type Period struct {
	Current       int  `json:"current"`
	IsHalftime    bool `json:"isHalftime"`
	IsEndOfPeriod bool `json:"isEndOfPeriod"`
}

type Score struct {
	Score string `json:"score"`
}

func GetScoreboard(c *gin.Context, date string) (*Scoreboard, error) {
	endpoint := fmt.Sprintf("%s/%s/scoreboard.json", c.GetString("NBA_ENDPOINT"), date)
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", endpoint, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var scoreboard Scoreboard
	json.NewDecoder(res.Body).Decode(&scoreboard)

	return &scoreboard, nil
}
