package server



//import "fmt"
import "net/http"


func Run(){
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/signup", SignUp)

	http.ListenAndServe(":8090", nil)
}