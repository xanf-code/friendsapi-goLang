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

// GET localhost:8000/api/friends

func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

// GET localhost:8000/api/friend/:id

func getFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // GET params
	for _, item := range friends {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(friends)
}

// POST localhost:8000/api/friends

func addFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friend Friend
	_ = json.NewDecoder(r.Body).Decode(&friend)
	friends = append(friends, friend)
	json.NewEncoder(w).Encode(friend)
}

// PUT localhost:8000/api/friend/:id

func updateFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range friends {
		if item.ID == params["id"] {
			friends = append(friends[:index], friends[index+1:]...) //DELETE FUNCTION
			var friend Friend
			_ = json.NewDecoder(r.Body).Decode(&friend)
			friends = append(friends, friend)
			json.NewEncoder(w).Encode(friend)
			return
		}
	}
	json.NewEncoder(w).Encode(friends)
}

// DELETE localhost:8000/api/friend/:id

func deleteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range friends {
		if item.ID == params["id"] {
			friends = append(friends[:index], friends[index+1:]...) //DELETE FUNCTION
			break
		}
	}
	json.NewEncoder(w).Encode(friends)
}

func main() {
	r := mux.NewRouter()

	// Mock Data

	friends = append(friends, Friend{ID: "1", Name: "Shiva", Type: "Robot", Rating: &Rating{Ratings: "-3/10"}})
	friends = append(friends, Friend{ID: "2", Name: "Pranav", Type: "Enemey", Rating: &Rating{Ratings: "-0/10"}})
	friends = append(friends, Friend{ID: "3", Name: "Raghav", Type: "CowBoy", Rating: &Rating{Ratings: "11/10"}})

	r.HandleFunc("/api/friends", getFriends).Methods("GET")
	r.HandleFunc("/api/friend/{id}", getFriend).Methods("GET")
	r.HandleFunc("/api/friends", addFriend).Methods("POST")
	r.HandleFunc("/api/friends/{id}", updateFriend).Methods("PUT")
	r.HandleFunc("/api/friends/{id}", deleteFriend).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
