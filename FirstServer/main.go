package main

import (
	"fmt"
	"net/http"
	"time"
)

//We can create a handler struct
type HttpHandler struct {}

type PingPongHandler struct {}

//We need to implement the serveHTTP method on HttpHandler
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Println("/ request made")
	//we can create a byte object
	data := []byte("Hello World")
	res.Write(data)
}

func (h PingPongHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	//we can create a byte object
	fmt.Println("/ping request made")
	currTime := time.Now()

	// Calling String method
	time := currTime.String()
	data := []byte("Pong at "+time)
	res.Write(data)
}

func main(){
	handler := HttpHandler{}
	pingPong := PingPongHandler{}


	http.Handle("/ping",pingPong)
	http.Handle("/",handler)
	if err := http.ListenAndServe("0.0.0.0:5000",nil); err != nil {
		fmt.Println("Server started at PORT 8000")
	}
}
