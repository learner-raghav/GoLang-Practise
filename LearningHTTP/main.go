package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//The counter1,counter2,counter3 types will implement the counter interface here.
//Implies they should take this counter1,counter2,counter3 as receiver arguments
type counter interface {
	up() string
}

//We need to have 3 variables for those counters
type counter1 int
type counter2 int
type counter3 int

//Counter 1 implementation
func (count* counter1) up() string {
	*count += 100
	var res string
	res = "Counter count: "+ strconv.Itoa(int(*count))
	return res
}

//Counter 2 implementation
func (count *counter2) up() string {
	*count += 1
	var res string
	res = "Counter count: "+ strconv.Itoa(int(*count))
	return res
}

//Counter 3 implementation
func (count *counter3) up() string {
	*count += 1
	var res string
	if *count%5 == 0 {
		res = "Counter count: BUS"
	} else{
		res = "Counter count: "+ strconv.Itoa(int(*count))
	}
	return res
}

type generalHandler struct {
	c counter
}

func (handler generalHandler) Pong(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w,handler.c.up())
}

func main(){
	var hundredCounter counter1
	var pingPongCounter counter2
	var busCounter counter3

	hundredCounterObj := generalHandler{&hundredCounter}
	pingPongCounterObj := generalHandler{&pingPongCounter}
	busCounterObj := generalHandler{&busCounter}

	http.Handle("/ping",http.HandlerFunc(pingPongCounterObj.Pong))
	http.Handle("/bus",http.HandlerFunc(busCounterObj.Pong))
	http.Handle("/hundred",http.HandlerFunc(hundredCounterObj.Pong))
	http.ListenAndServe("0.0.0.0:8000",nil)
}