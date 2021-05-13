package api

import (
	"encoding/json"
	"net/http"
)

func listenerGet() {
	println("ich bin hier bei get")
}

func listenerAdd(w http.ResponseWriter, r *http.Request, lis *[]Listener) {
	var n Listener
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	print(n.user)
	var bc = append(*lis, n)
	lis = &bc
	println("ich bin hier bin post")
	print(bc[0].channel)
}

func listenerUpdate() {
	println("ich bin hier bin update")
}

func listenerDelete() {
	println("ich bin hier bin delete")
}
