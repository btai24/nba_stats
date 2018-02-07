package games

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/btai24/nba_stats/api"
	"github.com/gin-gonic/gin"
)

// type Header struct {
// 	GameDate       time.Time // TODO: turn into a date type
// 	GameSequence   int
// 	GameID         string
// 	GameStatusID   int
// 	GameStatusText string
// 	GameCode       string
// 	HomeTeamID     int
// 	VisitorTeamID  int
// 	Season         string
// 	LivePeriod     int
// 	LivePcTime     string
// 	// TODO: Nullable fields
// 	// NatlTVBroadcater string
// 	// HomeTVBroadcater string
// 	// AwayTVBroadcater string
// 	// LivePeriodTimeBcast string
// 	// WhStatus            int
// }
//
// type Scoreboard struct {
// 	GameDate     time.Time
// 	GameSequence int
// 	GameID       string
// 	TeamID       int
// 	TeamAbv      string
// 	TeamCity     string
// 	TeamWL       string
// 	Q1           int
// 	Q2           int
// 	Q3           int
// 	Q4           int
// 	OT1          int
// 	OT2          int
// 	OT3          int
// 	OT4          int
// 	OT5          int
// 	OT6          int
// 	OT7          int
// 	OT8          int
// 	OT9          int
// 	OT10         int
// 	PTS          int
// 	FG           float64
// 	FT           float64
// 	FG3          float64
// 	AST          int
// 	REB          int
// 	TOV          int
// }
//
// type LastMeeting struct {
// 	GameID       string
// 	LastGameID   string
// 	LastGameDate time.Time
// 	HomeID       int
// 	HomeCity     string
// 	HomeName     string
// 	HomeAbv      string
// 	HomePts      int
// 	VisitorID    int
// 	VisitorCity  string
// 	VisitorName  string
// 	VisitorAbv   string
// 	VisitorPts   int
// }
//
// type TeamLeaders struct {
// 	GameID        string
// 	TeamID        int
// 	TeamCity      string
// 	TeamName      string
// 	TeamAbv       string
// 	PtsPlayerID   int
// 	PtsPlayerName string
// 	Pts           int
// 	RebPlayerID   int
// 	RebPlayerName string
// 	Reb           int
// 	AstPlayerID   int
// 	AstPlayerName string
// 	Ast           int
// }

type BoxScore struct {
	GameId   string `json:"game_id"`
	HomeTeam Team   `json:"home_team"`
	AwayTeam Team   `json:"away_team"`
}

type Team struct {
	Id      float64   `json:"id"`
	City    string    `json:"city"`
	Name    string    `json:"name"`
	Abv     string    `json:"abbreviation"`
	WinLoss string    `json:"win_loss"`
	Score   Score     `json:"score"`
	Stats   TeamStats `json:"stats"`
}

type Score struct {
	Total int   `json:"total"`
	Q1    int   `json:"q1"`
	Q2    int   `json:"q2"`
	Q3    int   `json:"q3"`
	Q4    int   `json:"q4"`
	OT    []int `json:"overtime"`
}

type TeamStats struct {
	FieldGoalPct float64 `json:"fg_pct"`
	FreeThrowPct float64 `json:"ft_pct"`
	ThreePtPct   float64 `json:"three_pt_pct"`
	Assists      int     `json:"assists"`
	Rebounds     int     `json:"rebounds"`
	Turnovers    int     `json:"turnovers"`
}

func GetScores(c *gin.Context) {
	// today := time.Now().Local().Format("01/02/2006")
	endpoint := fmt.Sprintf("%s/scoreboardV2?DayOffset=0&LeagueID=00&gameDate=%s", c.GetString("NBA_ENDPOINT"), "02/06/2018")

	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Referer", c.GetString("NBA_REFERER"))
	req.Header.Set("User-Agent", c.GetString("NBA_USER_AGENT"))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var scoreboard api.Scoreboard
	json.NewDecoder(res.Body).Decode(&scoreboard)

	boxScores := []BoxScore{}
	for _, set := range scoreboard.ResultSets {
		if set.Name == "LineScore" {
			for i, row := range set.RowSet {
				if i%2 != 0 {
					continue
				}

				var boxScore BoxScore
				var awayTeam Team
				var homeTeam Team
				awayTeam.populateInfo(&boxScore, row, set.Headers)
				homeTeam.populateInfo(&boxScore, set.RowSet[i+1], set.Headers)

				boxScore.HomeTeam = homeTeam
				boxScore.AwayTeam = awayTeam

				boxScores = append(boxScores, boxScore)
			}
		}
	}

	json, _ := json.Marshal(boxScores)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(json)
}

func (team *Team) populateInfo(boxScore *BoxScore, row []interface{}, headers []string) {
	var score Score
	var stats TeamStats
	for i, _ := range row {
		if headers[i] == "GAME_ID" {
			boxScore.GameId = row[i].(string)
			continue
		}

		// Team Info
		if headers[i] == "TEAM_ID" {
			team.Id = row[i].(float64)
			continue
		}

		if headers[i] == "TEAM_ABBREVIATION" {
			team.Abv = row[i].(string)
			continue
		}

		if headers[i] == "TEAM_CITY_NAME" {
			team.City = row[i].(string)
			continue
		}

		if headers[i] == "TEAM_NAME" {
			team.Name = row[i].(string)
			continue
		}

		if headers[i] == "TEAM_WINS_LOSSES" {
			team.WinLoss = row[i].(string)
			continue
		}

		// Team Score Info
		if headers[i] == "PTS_QTR1" {
			score.Q1 = int(row[i].(float64))
			continue
		}

		if headers[i] == "PTS_QTR2" {
			score.Q2 = int(row[i].(float64))
			continue
		}

		if headers[i] == "PTS_QTR3" {
			score.Q3 = int(row[i].(float64))
			continue
		}

		if headers[i] == "PTS_QTR4" {
			score.Q4 = int(row[i].(float64))

			// Also do OT
			otArr := []int{}
			for j := 1; j <= 10; j++ {
				pts := int(row[i+j].(float64))
				if pts < 1 {
					break
				}
				otArr = append(otArr, pts)
			}
			score.OT = otArr
			continue
		}

		if headers[i] == "PTS" {
			score.Total = int(row[i].(float64))
			continue
		}

		// Team Stats Info
		if headers[i] == "FG_PCT" {
			stats.FieldGoalPct = row[i].(float64)
			continue
		}

		if headers[i] == "FT_PCT" {
			stats.FreeThrowPct = row[i].(float64)
			continue
		}

		if headers[i] == "FG3_PCT" {
			stats.ThreePtPct = row[i].(float64)
			continue
		}

		if headers[i] == "AST" {
			stats.Assists = int(row[i].(float64))
			continue
		}

		if headers[i] == "REB" {
			stats.Rebounds = int(row[i].(float64))
			continue
		}

		if headers[i] == "TOV" {
			stats.Turnovers = int(row[i].(float64))
			continue
		}
	}

	team.Score = score
	team.Stats = stats
}
