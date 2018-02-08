package games

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/btai24/nba_stats/api"
	"github.com/gin-gonic/gin"
)

type BoxScore struct {
	GameId   string `json:"game_id"`
	HomeTeam Team   `json:"home_team"`
	AwayTeam Team   `json:"away_team"`
}

type Team struct {
	Id string `json:"id"`
	// City       string    `json:"city"`
	// Name       string    `json:"name"`
	Abv        string `json:"abbreviation"`
	Wins       string `json:"wins"`
	Losses     string `json:"losses"`
	Total      int    `json:"total"`
	LineScores []int  `json:"line_scores"`
}

func GetScores(c *gin.Context) {
	date, ok := c.GetQuery("date")
	if !ok {
		date = time.Now().Local().Format("20060102")
	} else {
		tmpDate, _ := time.Parse("01/02/2006", date)
		date = tmpDate.Format("20060102")
	}

	endpoint := fmt.Sprintf("%s/%s/scoreboard.json", c.GetString("NBA_ENDPOINT"), date)
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", endpoint, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var scoreboard api.Scoreboard
	json.NewDecoder(res.Body).Decode(&scoreboard)

	boxScores := []BoxScore{}

	for _, game := range scoreboard.Games {
		hTeamScore := []int{}
		for _, quarter := range game.HomeTeam.LineScore {
			score, _ := strconv.Atoi(quarter.Score)
			hTeamScore = append(hTeamScore, score)
		}
		hTeamTotal, _ := strconv.Atoi(game.HomeTeam.Score)
		homeTeam := Team{
			Id:         game.HomeTeam.Id,
			Abv:        game.HomeTeam.Abbreviation,
			Wins:       game.HomeTeam.Wins,
			Losses:     game.HomeTeam.Losses,
			Total:      hTeamTotal,
			LineScores: hTeamScore,
		}

		vTeamScore := []int{}
		for _, quarter := range game.VisitingTeam.LineScore {
			score, _ := strconv.Atoi(quarter.Score)
			vTeamScore = append(hTeamScore, score)
		}

		vTeamTotal, _ := strconv.Atoi(game.VisitingTeam.Score)
		visitingTeam := Team{
			Id:         game.VisitingTeam.Id,
			Abv:        game.VisitingTeam.Abbreviation,
			Wins:       game.VisitingTeam.Wins,
			Losses:     game.VisitingTeam.Losses,
			Total:      vTeamTotal,
			LineScores: vTeamScore,
		}

		boxScore := BoxScore{
			GameId:   game.Id,
			HomeTeam: homeTeam,
			AwayTeam: visitingTeam,
		}

		boxScores = append(boxScores, boxScore)
	}

	// TODO(Brian): handle error
	json, _ := json.Marshal(boxScores)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(json)
}
