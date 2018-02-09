package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TeamsLeague struct {
	League Teams `json:"league"`
}

type Teams struct {
	Teams []Team `json:"standard"`
}

type Team struct {
	Id           string  `json:"teamId"`
	Abbreviation string  `json:"triCode"`
	Wins         string  `json:"win"`
	Losses       string  `json:"loss"`
	Score        string  `json:"score"`
	LineScore    []Score `json:"linescore"`

	City       string `json:"city"`
	FullName   string `json:"fullName"`
	Name       string `json:"nickname"`
	UrlName    string `json:"urlName"`
	Conference string `json:"confName"`
	Division   string `json:"divName"`
}

func GetTeams(c *gin.Context) ([]Team, error) {
	// TODO: Interpolate the year (2017)
	endpoint := fmt.Sprintf("%s/2017/teams.json", c.GetString("NBA_ENDPOINT"))
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", endpoint, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var teams TeamsLeague
	json.NewDecoder(res.Body).Decode(&teams)

	return teams.League.Teams, nil
}
