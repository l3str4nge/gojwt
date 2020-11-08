package server

import "net/http"
import "fmt"

import "github.com/mateuszz0000/gojwt/src/auth"


func Hello(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w, "elo\n")
}

func Login(w http.ResponseWriter, request *http.Request) {
	headers := GetHeadersFromRequest(request)
	status := auth.SignIn(headers)
	
	response := "NOK"
	if status {
		response = "OK"
	}
	fmt.Fprintf(w, "Login status: " + response)
}

func Logout(w http.ResponseWriter, request *http.Request) {

}

func SignUp(w http.ResponseWriter, request *http.Request) {
	headers := GetHeadersFromRequest(request)
	token := auth.SignUp(headers["Username"], headers["Password"])
	fmt.Fprintf(w, "token\n", token)
}
