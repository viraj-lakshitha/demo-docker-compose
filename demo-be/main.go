package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Users struct {
	Data []struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		Email     string `json:"email"`
		Description string `json:"description"`
	} `json:"data"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving user details")
	
	var users Users
	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to read data (%s)", err.Error()), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to unmarshal data (%s)", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Handling CORS
	router := mux.NewRouter()
	headerHandle := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodHandle := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	originHandle := handlers.AllowedOrigins([]string{"*"})

	// Requests
	router.HandleFunc("/api/v1/users", getUsers).Methods("GET")

	// Start Server
	fmt.Println("BE running on http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", handlers.CORS(headerHandle, methodHandle, originHandle)(router))
	if err != nil {
		fmt.Println("Unable to start the server - ", err.Error())
		os.Exit(1)
	}
}