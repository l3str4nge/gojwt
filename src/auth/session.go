package auth

import "time"
import "fmt"
import "github.com/mateuszz0000/gojwt/src/redis_cli"


func CreateSession(username string) {
	fmt.Println("Create session for user ", username)
	red :=  redis_cli.GetRedis()

	sessionKey := username + "_session"
	createdAt := fmt.Sprint(time.Now())
	red.Set(ctx, sessionKey, createdAt, -1)
}