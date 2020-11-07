package server

import "net/http"
import "fmt"

import "github.com/mateuszz0000/gojwt/src/auth"


func Hello(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w, "elo\n")
}

func Login(w http.ResponseWriter, request *http.Request) {

}

func Logout(w http.ResponseWriter, request *http.Request) {

}

func SignUp(w http.ResponseWriter, request *http.Request) {
	headers := GetHeadersFromRequest(request)
	x := auth.SignUp(headers["Username"], headers["Password"])
	fmt.Fprintf(w, "signup\n", x)
}
