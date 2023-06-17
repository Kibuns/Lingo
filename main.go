package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kibuns/Lingo/DAL"
	"github.com/Kibuns/Lingo/Logic"
	"github.com/Kibuns/Lingo/Models"
	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
    server := mux.NewRouter().StrictSlash(true)
	server.Use(CORS)

	server.HandleFunc("/", helloWorld)
	server.HandleFunc("/query/{query}", guessQuery)
	server.HandleFunc("/start", startSession)

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

func guessQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	var id Models.ID

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, "Could not decode body into id", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	session, err := DAL.GetSession(id.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("secret word is: " + session.SecretWord)
	json.NewEncoder(w).Encode(Logic.ParseUserInput(query, "test"))
}

func startSession(w http.ResponseWriter, r *http.Request) {
	idString, err := DAL.StoreSession()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "started session: " + idString)
}
