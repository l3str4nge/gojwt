package auth

import "encoding/json"
import "time"
import "fmt"
import "github.com/mateuszz0000/gojwt/src/redis_cli"
import "context"
import "strings"

var ctx = context.Background()


var JWTHeader = map[string]string {
	"alg": "HS256",
	"typ": "JWT",
}

func CreateToken(username, password string) string {
	red :=  redis_cli.GetRedis()

	headerString := map2JsonString(JWTHeader)
	headerBase64 := string2Base64URL(headerString)

	payloadString := map2JsonString(createJWTPayload(username))
	payloadBase64 := string2Base64URL(payloadString)

	secret, _ := red.HGet(ctx, username, "secret").Result()

	signatureString := headerBase64 + "." + payloadBase64
	signatureVerify := createHMACSHA256FromString(signatureString, secret)


	jwtToken := headerBase64 + "." + payloadBase64 + "." + signatureVerify

	red.Set(ctx, username + "_jwt", jwtToken, 30000000000000) 
	return jwtToken
}


func createJWTPayload(username string) map[string]string {
	now := time.Now()
	
	sub := fmt.Sprint(now.Unix())
	iat := fmt.Sprint(now.AddDate(0, 0, 7).Unix()) // define expiration delta somwehere else
	
	return map[string]string {
		"sub": sub,
		"name": username,
		"iat": iat,
	}
}


func map2JsonString(m map[string]string) string {
	data, _ := json.Marshal(m)
	return string(data)
}


func jsonString2map(s string, m *map[string]string) map[string]string {
	json.Unmarshal([]byte(s), &m)
	return *m
}


func IsTokenValid(token string) (bool, string) {
	// TODO: error handling
	red :=  redis_cli.GetRedis()
	payload := decodeTokenPayload(token)
	username := payload["name"]
	
	storedToken, _ := red.Get(ctx, username + "_jwt").Result()

	if token == storedToken {
		return true, username
	}

	return false, ""
}

func decodeTokenPayload(token string) map[string]string{
	payloadb64 := strings.Split(token , ".")[1]
	decodedPayload := Base64URL2String(payloadb64)
	var payloadSkeleton map[string]string
	payload := jsonString2map(decodedPayload, &payloadSkeleton)
	return payload
}