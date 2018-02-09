package deps

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func RedisClient(c *gin.Context) (*redis.Client, bool) {
	client, ok := c.Get("REDIS_CLIENT")
	return client.(*redis.Client), ok
}
