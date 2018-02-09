package standings

import (
	"encoding/json"
	"fmt"
	"math"
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
	Id             string      `json:"id"`
	Profile        TeamProfile `json:"profile"`
	Wins           int         `json:"wins"`
	Losses         int         `json:"losses"`
	GamesBehind    float64     `json:"games_behind"`
	WinPct         float64     `json:"win_percentage"`
	ConferenceRank int         `json:"conference_rank"`
	Streak         int         `json:"streak"`
	Record         string      `json:"record"`
	HomeRecord     string      `json:"home_record"`
	AwayRecord     string      `json:"away_record"`
	LastTen        string      `json:"last_ten"`
}

func GetConferenceStandings(c *gin.Context) {
	// TODO: handle error
	standings, _ := api.GetConferenceStandings(c)
	var eastStandings [15]Team
	var westStandings [15]Team

	for i, team := range standings.Conference.East {
		eastStandings[i] = formTeam(c, team)
	}

	for i, team := range standings.Conference.West {
		westStandings[i] = formTeam(c, team)
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

// Redis struct - move?
type TeamProfile struct {
	Abbreviation string `json:"triCode"`
	City         string `json:"city"`
	FullName     string `json:"fullName"`
	Name         string `json:"nickname"`
	UrlName      string `json:"urlName"`
	Conference   string `json:"confName"`
	Division     string `json:"divName"`
}

func formTeam(c *gin.Context, team api.StandingsTeam) Team {
	// TODO: handle error
	redisClient, _ := deps.RedisClient(c)

	wins, _ := strconv.Atoi(team.Wins)
	losses, _ := strconv.Atoi(team.Losses)
	winPct, _ := strconv.ParseFloat(team.WinPct, 32)
	gamesBehind, _ := strconv.ParseFloat(team.GamesBehind, 32)
	confRank, _ := strconv.Atoi(team.ConferenceRank)

	streak, _ := strconv.Atoi(team.Streak)
	if !team.IsWinStreak {
		streak = streak * -1
	}

	formedTeam := Team{
		Id:             team.Id,
		Wins:           wins,
		Losses:         losses,
		WinPct:         winPct,
		GamesBehind:    gamesBehind,
		ConferenceRank: confRank,
		Streak:         streak,
		Record:         fmt.Sprintf("%s-%s", team.Wins, team.Losses),
		HomeRecord:     fmt.Sprintf("%s-%s", team.HomeWins, team.HomeLosses),
		AwayRecord:     fmt.Sprintf("%s-%s", team.AwayWins, team.AwayLosses),
		LastTen:        fmt.Sprintf("%s-%s", team.LastTenWins, team.LastTenLosses),
	}

	teamKey := fmt.Sprintf("team:info#%s", team.Id)
	teamInfo, _ := redisClient.Get(teamKey).Result()

	if teamInfo == "" {
		teams, _ := api.GetTeams(c)
		for _, t := range teams {
			if t.Id == team.Id {
				profile := TeamProfile{
					Abbreviation: t.Abbreviation,
					City:         t.City,
					FullName:     t.FullName,
					Name:         t.Name,
					UrlName:      t.UrlName,
					Conference:   t.Conference,
					Division:     t.Division,
				}
				json, _ := json.Marshal(profile)
				redisClient.Set(teamKey, json, math.MaxInt32*time.Second)
				formedTeam.Profile = profile
			}
		}
	} else {
		var profile TeamProfile
		json.Unmarshal([]byte(teamInfo), &profile)
		formedTeam.Profile = profile
	}

	return formedTeam
}
