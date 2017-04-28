package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/btai35/nba_stats/app/games"
)

type Scoreboard struct {
	Parameters Parameters  `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

type ResultSet struct {
	Name    string          `json:"name"`
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

	for _, set := range scoreboard.ResultSets {
		switch set.Name {
		case "GameHeader":
			for _, row := range set.RowSet {
				gameTime, err := time.Parse("2006-01-02T15:04:05", row[0].(string))
				if err != nil {
					fmt.Println(err)
				}
				game := games.GameHeader{
					GameDateEst:    gameTime,
					GameSequence:   int(row[1].(float64)),
					GameID:         row[2].(string),
					GameStatusID:   int(row[3].(float64)),
					GameStatusText: row[4].(string),
					GameCode:       row[5].(string),
					HomeTeamID:     strconv.Itoa(int(row[6].(float64))),
					VisitorTeamID:  strconv.Itoa(int(row[7].(float64))),
					Season:         row[8].(string),
					LivePeriod:     int(row[9].(float64)),
					LivePcTime:     row[10].(string),
					// NatlTVBroadcater: row[11].(string),
					// HomeTVBroadcater: row[12].(string),
					// AwayTVBroadcater: row[13].(string),
				}

				fmt.Println(game)
			}
		}
	}
}
