package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/btai35/nba_stats/app/games"
	"github.com/btai35/nba_stats/app/standings"
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
				game := games.Header{
					GameDate:       gameTime,
					GameSequence:   int(row[1].(float64)),
					GameID:         row[2].(string),
					GameStatusID:   int(row[3].(float64)),
					GameStatusText: row[4].(string),
					GameCode:       row[5].(string),
					HomeTeamID:     int(row[6].(float64)),
					VisitorTeamID:  int(row[7].(float64)),
					Season:         row[8].(string),
					LivePeriod:     int(row[9].(float64)),
					LivePcTime:     row[10].(string),
					// NatlTVBroadcater: row[11].(string),
					// HomeTVBroadcater: row[12].(string),
					// AwayTVBroadcater: row[13].(string),
				}
				fmt.Println(game)
			}
		case "LineScore":
			for _, row := range set.RowSet {
				gameTime, err := time.Parse("2006-01-02T15:04:05", row[0].(string))
				if err != nil {
					fmt.Println(err)
				}
				score := games.Scoreboard{
					GameDate:     gameTime,
					GameSequence: int(row[1].(float64)),
					GameID:       row[2].(string),
					TeamID:       int(row[3].(float64)),
					TeamAbv:      row[4].(string),
					TeamCity:     row[5].(string),
					TeamWL:       row[6].(string),
					Q1:           int(row[7].(float64)),
					Q2:           int(row[8].(float64)),
					Q3:           int(row[9].(float64)),
					Q4:           int(row[10].(float64)),
					OT1:          int(row[11].(float64)),
					OT2:          int(row[12].(float64)),
					OT3:          int(row[13].(float64)),
					OT4:          int(row[14].(float64)),
					OT5:          int(row[15].(float64)),
					OT6:          int(row[16].(float64)),
					OT7:          int(row[17].(float64)),
					OT8:          int(row[18].(float64)),
					OT9:          int(row[19].(float64)),
					OT10:         int(row[20].(float64)),
					PTS:          int(row[21].(float64)),
					FG:           row[22].(float64),
					FT:           row[23].(float64),
					FG3:          row[24].(float64),
					AST:          int(row[25].(float64)),
					REB:          int(row[26].(float64)),
					TOV:          int(row[27].(float64)),
				}
				fmt.Println(score)
			}
		case "SeriesStandings":
			for _, row := range set.RowSet {
				gameTime, err := time.Parse("2006-01-02T15:04:05", row[3].(string))
				if err != nil {
					fmt.Println(err)
				}
				series := standings.Series{
					GameID:       row[0].(string),
					HomeID:       int(row[1].(float64)),
					VisitorID:    int(row[2].(float64)),
					GameDate:     gameTime,
					HomeWins:     int(row[4].(float64)),
					HomeLosses:   int(row[5].(float64)),
					SeriesLeader: row[6].(string),
				}
				fmt.Println(series)
			}
		case "LastMeeting":
			for _, row := range set.RowSet {
				gameTime, err := time.Parse("2006-01-02T15:04:05", row[2].(string))
				if err != nil {
					fmt.Println(err)
				}
				lastMeeting := games.LastMeeting{
					GameID:       row[0].(string),
					LastGameID:   row[1].(string),
					LastGameDate: gameTime,
					HomeID:       int(row[3].(float64)),
					HomeCity:     row[4].(string),
					HomeName:     row[5].(string),
					HomeAbv:      row[6].(string),
					HomePts:      int(row[7].(float64)),
					VisitorID:    int(row[8].(float64)),
					VisitorCity:  row[9].(string),
					VisitorName:  row[10].(string),
					VisitorAbv:   row[11].(string),
					VisitorPts:   int(row[12].(float64)),
				}
				fmt.Println(lastMeeting)
			}
		case "EastConfStandingsByDay", "WestConfStandingsByDay":
			conf := standings.ConferenceByDay{}
			if set.Name == "EastConfStandingsByDay" {
				conf.Conference = "East"
			} else {
				conf.Conference = "West"
			}
			var teams [15]standings.Team
			for i, row := range set.RowSet {
				if i == 0 {
					confDate, err := time.Parse("01/02/2006", row[3].(string))
					if err != nil {
						fmt.Println(err)
					}
					conf.Date = confDate
				}
				teams[i] = standings.Team{
					TeamID:     int(row[0].(float64)),
					SeasonID:   row[2].(string),
					City:       row[5].(string),
					Win:        int(row[7].(float64)),
					Loss:       int(row[8].(float64)),
					WinPct:     row[9].(float64),
					HomeRecord: row[10].(string),
					RoadRecord: row[11].(string),
				}
			}
			conf.Teams = teams
			fmt.Println(conf)
		case "TeamLeaders":
			for _, row := range set.RowSet {
				leaders := games.TeamLeaders{
					GameID:        row[0].(string),
					TeamID:        int(row[1].(float64)),
					TeamCity:      row[2].(string),
					TeamName:      row[3].(string),
					TeamAbv:       row[4].(string),
					PtsPlayerID:   int(row[5].(float64)),
					PtsPlayerName: row[6].(string),
					Pts:           int(row[7].(float64)),
					RebPlayerID:   int(row[8].(float64)),
					RebPlayerName: row[9].(string),
					Reb:           int(row[10].(float64)),
					AstPlayerID:   int(row[11].(float64)),
					AstPlayerName: row[12].(string),
					Ast:           int(row[13].(float64)),
				}
				fmt.Println(leaders)
			}
		}
	}
}
