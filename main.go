package main

import (
	"log"
	"net/http"

	"github.com/AustinTi/MetaGrabAPI/misc"
	"github.com/AustinTi/MetaGrabAPI/search"
	"github.com/AustinTi/MetaGrabAPI/track"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/track", track.Handler)
	mux.HandleFunc("/search", search.Handler)

	log.Fatal(http.ListenAndServe(misc.Config.BindAddress, rateLimit(mux)))
}
