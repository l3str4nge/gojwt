package auth

import "fmt"
import "strings"


func SignIn(headers map[string]string) bool{
	logged := false
	if JWTProvided(headers){
		logged = checkTokenAndSignIn(headers["Authorization"])
	} else {
		// login with plain user and pswd
		// username := headers["Username"]
		// password := headers["Password"]
		// logged := createTokenAndLogin(username, password)
		return false
	}

	return logged
}


func JWTProvided(headers map[string]string) bool {
	if val, ok := headers["Authorization"]; ok {
		fmt.Println("Autorization header exists. Check for schema")
		if !strings.HasPrefix(val, "Bearer"){
			fmt.Println("Invalid schema")
			return false
		} 

		return true
	}

	return false
}

func checkTokenAndSignIn(token string) bool {
	// tokenWitoutSchema := strings.ReplaceAll(token, "Bearer", "")
	return true
}

func createTokenAndLogin(username, password string) bool{

	// _ := CreateToken(username, password)
	return true
}