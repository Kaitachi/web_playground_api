package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// defining router
	mux := http.NewServeMux()
	mux.HandleFunc("/time", getTime)


	// starting server
	fmt.Println("Server is running at port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := []Time{
		{ CurrentTime: http.TimeFormat },
	}

	json.NewEncoder(w).Encode(currentTime)
}
