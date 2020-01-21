package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HeartBeatHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func SquareHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	num, err := strconv.Atoi(req.FormValue("input"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not parse input\n")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("%d", num*num))
}

func main() {

	http.HandleFunc("/hb", HeartBeatHandler)
	http.HandleFunc("/square", SquareHandler)

	http.ListenAndServe(":8000", nil)
}
