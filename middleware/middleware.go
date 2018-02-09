package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func SetNbaEndpoint(c *gin.Context) {
	c.Set("NBA_ENDPOINT", "http://data.nba.net/10s/prod/v1")

	c.Next()
}

func SetRedisClient(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	c.Set("REDIS_CLIENT", client)

	c.Next()
}

func Authenticate(c *gin.Context) {

}
