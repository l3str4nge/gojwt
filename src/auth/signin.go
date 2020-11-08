package auth

import "fmt"
import "strings"
import "github.com/mateuszz0000/gojwt/src/redis_cli"


func SignIn(headers map[string]string) bool{
	logged := false
	fmt.Println(headers)
	if JWTProvided(headers){
		logged = checkTokenAndSignIn(headers["Authorization"])
	} else {
		username := headers["Username"]
		password := headers["Password"]
		logged = checkCredentialsAndSignIn(username, password)		
	}

	return logged
}

func JWTProvided(headers map[string]string) bool {
	if val, ok := headers["Authorization"]; ok {
		if !strings.HasPrefix(val, "Bearer"){
			return false
		} 
	}
	return false
}

func checkTokenAndSignIn(token string) bool {
	tokenWitoutSchema := strings.ReplaceAll(token, "Bearer", "")
	trimmedToken := strings.Trim(tokenWitoutSchema, " ")
	canLogin, username := IsTokenValid(trimmedToken)
	

	if !canLogin {
		return false
	}
	

	CreateSession(username)
	return true
}


func checkCredentialsAndSignIn(username, password string) bool {
	red :=  redis_cli.GetRedis()
	pswd, _ := red.HGet(ctx, username, "password").Result()
	givenPswdHash := HashPassword(password)
	
	// insecure because we dont have any salt at all :P
	if pswd  == givenPswdHash {
		CreateSession(username)
		return true
	}

	return false
}