package main

import (
	"fmt"
	"net/http"
	"time"
)

//We need to implement the serveHTTP method on HttpHandler
func handleHome(res http.ResponseWriter, req *http.Request){
	fmt.Println("/ request made")
	//we can create a byte object
	data := []byte("Hello World")
	res.Write(data)
}

func handlePingPong(res http.ResponseWriter, req *http.Request){
	//we can create a byte object
	fmt.Println("/ping request made")
	currTime := time.Now()

	// Calling String method
	time := currTime.String()
	data := []byte("Pong at "+time)
	res.Write(data)
}

func main(){

	http.HandleFunc("/ping",handlePingPong)
	http.HandleFunc("/",handleHome)
	if err := http.ListenAndServe("0.0.0.0:5000",nil); err != nil {
		fmt.Println("Server started at PORT 8000")
	}
}
