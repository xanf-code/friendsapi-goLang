package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Friend struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Rating *Rating `json:"ratings"`
}

type Rating struct {
	Ratings string `json : "ratings"`
}

var friends []Friend

func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func getFriend(w http.ResponseWriter, r *http.Request) {

}

func addFriend(w http.ResponseWriter, r *http.Request) {

}

func updateFriend(w http.ResponseWriter, r *http.Request) {

}

func deleteFriend(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// Mock Data

	friends = append(friends, Friend{ID: "1", Name: "Shiva", Type: "Friend", Rating: &Rating{Ratings: "10/10"}})
	friends = append(friends, Friend{ID: "2", Name: "Pranav", Type: "Enemey", Rating: &Rating{Ratings: "0/10"}})
	friends = append(friends, Friend{ID: "3", Name: "Raghav", Type: "Friend", Rating: &Rating{Ratings: "10/10"}})

	r.HandleFunc("/api/friends", getFriends).Methods("GET")
	r.HandleFunc("/api/friend/{id}", getFriend).Methods("GET")
	r.HandleFunc("/api/friends", addFriend).Methods("POST")
	r.HandleFunc("/api/friends/{id}", updateFriend).Methods("PUT")
	r.HandleFunc("/api/friends/{id}", deleteFriend).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
