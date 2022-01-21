package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	port := os.Args[1]
	fmt.Printf("Port is %v", port)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/headers", headersHandler)
	http.ListenAndServe(":8090", nil)
}
