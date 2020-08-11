package main

import (
	"log"
	"net/http"

	"github.com/AustinTi/SpotGrabAPI/misc"
	"github.com/AustinTi/SpotGrabAPI/search"
	"github.com/AustinTi/SpotGrabAPI/track"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/track", track.Handler)
	mux.HandleFunc("/search", search.Handler)

	log.Fatal(http.ListenAndServe(misc.Config.BindAddress, rateLimit(mux)))
}
