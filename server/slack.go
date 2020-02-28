package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func startSlackServer() {
	http.HandleFunc("/slack", handleSlack)

	log.Printf("starting slack server...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func handleSlack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(`get method is not allowed`)
		return
	}
	defer closeCloser(r.Body)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(data)

	switch data["type"].(string) {
	case "url_verification":
		if challenge, ok := data["challenge"].(string); !ok {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("challenge required at url verification")); err != nil {
				log.Printf("failed to write challenge response: %#v", err)
			}
		} else {
			if _, err := w.Write([]byte(challenge)); err != nil {
				log.Printf("failed to write challenge response: %#v", err)
				w.WriteHeader(http.StatusServiceUnavailable)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	case "event_callback":
		event, ok := data[`event`].(map[string]interface{})
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(`failed to cast event to byte array`)
			return
		}
		w.WriteHeader(http.StatusOK)
		putComment(event)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println(`bad request`)
	}
}

func putComment(event map[string]interface{}) {
	fmt.Sprintf("%+v", event)

}
