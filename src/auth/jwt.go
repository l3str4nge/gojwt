package auth

import "encoding/json"
import "time"
import "github.com/mateuszz0000/gojwt/src/redis_cli"
import "github.com/go-redis/redis/v8"
import "context"

var ctx = context.Background()

func ValidateToken(token string) {

}


var JWTHeader = map[string]string {
	"alg": "HS256",
	"typ": "JWT",
}

func CreateToken(username, password string) string {
	red :=  redis_cli.GetRedis()

	header := map2JsonString(JWTHeader)
	headerSha := createSha256FromString(header)

	payload := map2JsonString(createJWTPayload(username))
	payloadSha := createSha256FromString(payload)


	secret, _ := red.HGet(ctx, username, "secret").Result()

	signature := header + "." + payload + "." + secret
	signatureSha := createSha256FromString(signature)
	
	jwtToken := headerSha + "." + payloadSha + "." + signatureSha
	return jwtToken
}


func createJWTPayload(username string) map[string]string {
	now := time.Now()
	
	sub := string(now.Unix())
	iat := string(now.AddDate(0, 0, 7).Unix()) // define expiration delta somwehere else
	
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

