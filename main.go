package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Scoreboard struct {
	Parameters Parameters  `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

type ResultSet struct {
	Name    string          `json:"resource"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:rowSet`
}

type Parameters struct {
	GameDate  string `json:"GameDate"`
	LeagueID  string `json:"LeagueID"`
	DayOffset int    `json:"DayOffset"`
}

func main() {
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", "http://stats.nba.com/stats/scoreboardV2?DayOffset=0&LeagueID=00&gameDate=04/23/2017", nil)
	req.Header.Set("Referer", "http://stats.nba.com/scores/")
	req.Header.Set("User-Agent", "Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var scoreboard Scoreboard
	json.NewDecoder(res.Body).Decode(&scoreboard)
	fmt.Println(scoreboard)
}
