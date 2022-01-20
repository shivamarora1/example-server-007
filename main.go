package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{"response": "success"}
	json.NewEncoder(w).Encode(response)
}
func headersHandler(w http.ResponseWriter, req *http.Request) {
	allHeaders := map[string]string{}
	for name, headers := range req.Header {
		for _, h := range headers {
			allHeaders[name] = h
		}
	}
	w.Header().Add("content-type", "json")
	json.NewEncoder(w).Encode(allHeaders)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/headers", headersHandler)
	http.ListenAndServe(":8090", nil)
}
