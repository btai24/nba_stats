package middleware

import "github.com/gin-gonic/gin"

func SetNbaEndpoint(c *gin.Context) {
	c.Set("NBA_ENDPOINT", "http://data.nba.net/10s/prod/v1")

	c.Next()
}

func Authenticate(c *gin.Context) {

}
