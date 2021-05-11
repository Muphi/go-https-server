package api

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/listener", listener)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func listener(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		listenerGet()
	case "POST":
		w.WriteHeader(http.StatusCreated)
		listenerAdd()
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		listenerUpdate()
	case "DELETE":
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func addMessageForListener(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
