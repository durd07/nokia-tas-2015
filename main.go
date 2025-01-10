package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	// http.HandleFunc("/screen", service.ScreenHandler)
	// http.HandleFunc("/api/count", service.CounterHandler)

	http.HandleFunc("/api/members", service.VoteMembersHandler)
	http.HandleFunc("/api/vote", service.VoteHandler)
	http.HandleFunc("/api/dump", service.VoteDumpHandler)
	http.HandleFunc("/api/config", service.ConfigHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
