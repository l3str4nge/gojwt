package auth


import "fmt"
import "github.com/mateuszz0000/gojwt/src/redis_cli"
import "github.com/go-redis/redis/v8"
//import "context"

//var ctx = context.Background()

func SignUp(username, password string) string{
	red :=  redis_cli.GetRedis()
	fmt.Println("before redis", username)
	_, err := red.Get(ctx, username).Result()

	if err == redis.Nil {
		fmt.Println("Username does not exists. Creating.")
		hashedPassword := HashPassword(password)
		token := storeUserInRedis(red, username, hashedPassword)
		return token
	
	}else{
		fmt.Println("errrr", err)
	}

	return ""
}

func storeUserInRedis(red *redis.Client, username, password string) string{	
	
	// TODO: error handling, use HMSet
	red.HSet(ctx, username, "password", password)
	secret := GenerateSecretKey()
	red.HSet(ctx, username, "secret", secret)

	token := CreateToken(username, password)
	return token
}