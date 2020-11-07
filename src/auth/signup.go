package auth


import "fmt"
import "github.com/mateuszz0000/gojwt/src/redis_cli"
import "github.com/go-redis/redis/v8"
import "context"

var ctx = context.Background()

func SignUp(username, password string) bool{
	red :=  redis_cli.GetRedis()
	fmt.Println("before redis", username)
	_, err := red.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("Username does not exists. Creating.")
		hashedPassword := HashPassword(password)
		stored := storeUserInRedis(red, username, hashedPassword)
		if !stored {
			return false
		}	
	
	}else{
		fmt.Println("errrr", err)
	}

	return true
}

func storeUserInRedis(red *redis.Client, username, password string) bool{	
	
	// TODO: error handling, use HMSet
	red.HSet(ctx, username, "password", password)
	secret := GenerateSecretKey()
	red.HSet(ctx, username, "secret", secret)
	return true
}