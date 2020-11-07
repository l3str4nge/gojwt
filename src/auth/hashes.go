package auth


import "crypto/sha256"
//import "bytes"
//import "encoding"
import "encoding/base64"
import "time"


func HashPassword(password string) string{
	return createSha256FromString(password)
}

func GenerateSecretKey() string {
	timestamp := time.Now().Unix()
	return createSha256FromString(string(timestamp))
}

func createSha256FromString(s string) string {
	secret := sha256.New()
	secret.Write([]byte(string(s)))
	return base64.URLEncoding.EncodeToString(secret.Sum(nil))
}