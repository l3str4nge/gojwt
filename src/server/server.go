package server



//import "fmt"
import "net/http"

func Run(){
	http.HandleFunc("/hello", Hello)

	http.ListenAndServe(":8090", nil)
}