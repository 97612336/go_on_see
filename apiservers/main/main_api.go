package main

import (
	"go_on_see/apiservers/heartbeat"
	"go_on_see/apiservers/locate"
	"go_on_see/apiservers/objects"
	"log"
	"net/http"
	"os"
)

func main(){
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/",objects.Handler)
	http.HandleFunc("/objects/",locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
