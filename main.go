package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	rsp := response{
		Message: "Hello",
		Status:  "OK",
	}

	response, _ := json.MarshalIndent(rsp, "", "")
	w.Write(response)
}
