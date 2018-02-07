package middleware

import "github.com/gin-gonic/gin"

func SetNbaEndpoint(c *gin.Context) {
	c.Set("NBA_ENDPOINT", "http://stats.nba.com/stats")
	c.Set("NBA_REFERER", "")
	c.Set("NBA_USER_AGENT", "Safari/537.36")

	c.Next()
}

func Authenticate(c *gin.Context) {

}
