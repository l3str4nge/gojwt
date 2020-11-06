package server

import "net/http"
import "fmt"


func Hello(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w, "elo\n")
}