package auth


import "crypto/sha256"
import "crypto/hmac"
//import "bytes"
//import "encoding"
import "encoding/base64"
import "encoding/hex"
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
	return base64.RawURLEncoding.EncodeToString(secret.Sum(nil))
}

func createHMACSHA256FromString(s, secret string) string{
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write([]byte(s))
	return hex.EncodeToString(hm.Sum(nil))

}

func string2Base64URL(s string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}

func Base64URL2String(s string) string {
	b, _ := base64.RawURLEncoding.DecodeString(s)
	return string(b)
}