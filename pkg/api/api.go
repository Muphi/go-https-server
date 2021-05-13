package api

import (
	"log"
	"net/http"
)

var listenerGlobal *[]Listener

func HandleRequests(l *[]Listener) {
	listenerGlobal = l
	http.HandleFunc("/listener", listener)
	http.HandleFunc("/message", message)
	http.HandleFunc("/transmitter", transmitter)
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
		listenerAdd(w, r, listenerGlobal)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		listenerUpdate()
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		listenerDelete()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		messageGet()
	case "POST":
		w.WriteHeader(http.StatusCreated)
		messageAdd()
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		messageUpdate()
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		messageDelete()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func transmitter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		transmitterGet()
	case "POST":
		w.WriteHeader(http.StatusCreated)
		transmitterAdd()
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		transmitterUpdate()
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		transmitterDelete()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
