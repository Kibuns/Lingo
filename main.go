package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kibuns/Lingo/Logic"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world")
	handleRequests()
}

func handleRequests() {
    server := mux.NewRouter().StrictSlash(true)
	server.Use(CORS)

	server.HandleFunc("/", helloWorld)
	server.HandleFunc("/query/{query}", helloQuery)

    log.Println("API Gateway listening on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", server))
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: fix cors headers once frontend is developed

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		//return
	})

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("'Hello world' endpoint hit")
}

func helloQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	json.NewEncoder(w).Encode(Logic.ParseUserInput(query, "secret"))
}