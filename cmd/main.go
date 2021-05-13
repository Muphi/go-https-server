package main

import (
	"https-server/pkg/api"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

var ListenerGlobal []api.Listener

func main() {
	api.HandleRequests(&ListenerGlobal)
}
