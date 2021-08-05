package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func main() {
// 	http.HandleFunc("/", hello)
// 	fmt.Println("Server Started")
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message":"hello world!"}`))
// }

type Response struct {
	Message string `json:"message"`
}

func handleRequests() {
	mainRouter := mux.NewRouter()

	mainRouter.Handle("/healthcheck", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Im OK."
		responseJSON(message, w, http.StatusOK)
	}))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	handleRequests()
}

func responseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
