package main

import (
	"go_on_see/dataservers/heartbeat"
	"go_on_see/dataservers/locate"
	"go_on_see/dataservers/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
