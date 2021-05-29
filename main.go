package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hypebeast/go-osc/osc"
	log "github.com/schollz/logger"
)

func main() {
	log.SetLevel("debug")
	port := 8098
	log.Infof("listening on :%d", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UTC()
	err := handle(w, r)
	if err != nil {
		log.Error(err)
	}
	log.Infof("%v %v %v %s\n", r.RemoteAddr, r.Method, r.URL.Path, time.Since(t))
}

func handle(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// very special paths
	client := osc.NewClient("localhost", 57120)
	msg := osc.NewMessage("/osc/heartbeat")
	msg.Append(strings.TrimPrefix(r.URL.Path, "/"))
	client.Send(msg)

	return
}
