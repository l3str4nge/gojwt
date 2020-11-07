package redis_cli

//import "context"
import "github.com/go-redis/redis/v8"
import "fmt"

var redisInstance *redis.Client


func GetRedis() *redis.Client {

	if redisInstance != nil {
		return redisInstance
	}

	// below values should be in external config 
	redisInstance =  redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
	})
	fmt.Println(redisInstance)
	return redisInstance
}