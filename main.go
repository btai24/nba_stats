package main

import (
	"github.com/btai24/nba_stats/app/games"
	"github.com/btai24/nba_stats/app/standings"
	"github.com/btai24/nba_stats/middleware"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	api := gin.Default()
	api.Use(middleware.SetNbaEndpoint)
	api.Use(middleware.SetRedisClient)

	api.GET("/scores", games.GetScores)
	api.GET("/conference_standings", standings.GetConferenceStandings)

	api.Run()
}
