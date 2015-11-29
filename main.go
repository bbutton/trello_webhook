package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bbutton/trello_webhook/Godeps/_workspace/src/github.com/VojtechVitek/go-trello"
	"github.com/bbutton/trello_webhook/Godeps/_workspace/src/github.com/cloudfoundry-community/go-cfenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
		handlePing(w)
	case "POST":
		handleTrelloUpdate(w, r)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handlePing(w http.ResponseWriter) {
	log.Printf("Hello. World!!\n")
	w.WriteHeader(http.StatusOK)
}

func handleTrelloUpdate(w http.ResponseWriter, r *http.Request) {
	var trelloAction trello.Action
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&trelloAction)
	if err != nil {
		log.Printf("Failed to decode JSON. Error: %s", err.Error())
	}

	log.Printf("<JSON>%v</JSON>\n", trelloAction)
}

func findPort() string {
	var portString string

	appEnv, err := cfenv.Current()
	if err == nil {
		portString = ":" + strconv.Itoa(appEnv.Port)
	} else {
		portString = ":8081"
	}

	return portString
}

func main() {
	log.Printf("Server starting, listening on port %s", findPort())
	http.HandleFunc("/trello_webhook/", handler)
	http.ListenAndServe(findPort(), nil)
}
