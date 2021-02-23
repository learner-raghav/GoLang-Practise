package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Counter interface {
	up() string
}
type BusHandler struct {
	count int
}

type HundredHandler struct {
	count int
}

type PingPongHandler struct {
	count int
}

func (hundredHandler *HundredHandler) up() string {
	hundredHandler.count += 100
	var res string
	res = "Counter count: "+ strconv.Itoa(hundredHandler.count)
	return res
}

func (pingPongHandler *PingPongHandler) up() string {
	pingPongHandler.count += 1
	var res string
	res = "Counter count: "+ strconv.Itoa(pingPongHandler.count)
	return res
}


func (busHandler *BusHandler) up() string {
	busHandler.count += 1
	var res string
	if busHandler.count%5 == 0 {
		res = "Counter count: BUS"
	} else{
		res = "Counter count: "+ strconv.Itoa(busHandler.count)
	}
	return res
}





func (h *PingPongHandler) pingPongFunctionHandler(res http.ResponseWriter, req *http.Request){

	response := h.up()
	fmt.Fprint(res,response)
}
func (h *HundredHandler) hundredFunctionHandler(res http.ResponseWriter, req *http.Request){

	response := h.up()
	fmt.Fprint(res,response)
}
func (h *BusHandler) busFunctionHandler(res http.ResponseWriter, req *http.Request){

	response := h.up()
	fmt.Fprint(res,response)
}

func main(){
	pingPongHandler := PingPongHandler{}
	busHandler := BusHandler{}
	hundredHandler := HundredHandler{}

	http.Handle("/ping",http.HandlerFunc(pingPongHandler.pingPongFunctionHandler))
	http.Handle("/bus",http.HandlerFunc(busHandler.busFunctionHandler))
	http.Handle("/hundred",http.HandlerFunc(hundredHandler.hundredFunctionHandler))
	http.ListenAndServe("0.0.0.0:5000",nil)
}