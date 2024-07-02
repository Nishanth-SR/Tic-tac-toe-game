package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Game struct {
	Board  [9]string `json:"board"`
	Player string    `json:"player"`
}

var game Game

func handleNewGame(w http.ResponseWriter, r *http.Request) {
	game = Game{
		Board:  [9]string{"", "", "", "", "", "", "", "", ""},
		Player: "X",
	}
	json.NewEncoder(w).Encode(game)
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	var move struct {
		Index int `json:"index"`
	}
	if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if game.Board[move.Index] == "" {
		game.Board[move.Index] = game.Player
		if game.Player == "X" {
			game.Player = "O"
		} else {
			game.Player = "X"
		}
	}
	json.NewEncoder(w).Encode(game)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/move", handleMove).Methods("POST")
	r.HandleFunc("/api/newgame", handleNewGame).Methods("GET")

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
