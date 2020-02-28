package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack/slackevents"
)

func startSlackServer() {
	http.HandleFunc("/slack", handleSlack)

	log.Printf("starting slack server...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func handleSlack(w http.ResponseWriter, r *http.Request) {
	log.Println("start handle slack message")
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		log.Printf("failed to read from request body: %#v", err)
	}
	body := buf.String()
	eventsAPIEvent, err := slackevents.ParseEvent(
		json.RawMessage(body),
		slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: os.Getenv(`SLACK_VERIFICATION_TOKEN`)}),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to parse event: %#v", err)
		return
	}

	switch eventsAPIEvent.Type {
	case slackevents.URLVerification:
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		if _, err := w.Write([]byte(r.Challenge)); err != nil {
			log.Printf("failed to write challenge response: %#v", err)
		}
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			putAppMentionEvent(ev)
		case *slackevents.MessageEvent:
			putMessageEvent(ev)
		default:
			log.Printf("unhandled callback event %#v", ev)
		}
	default:
		log.Printf("unhandled event: %#v", eventsAPIEvent)
	}

}

func putAppMentionEvent(ev *slackevents.AppMentionEvent) {
	if ignoreMessage(ev.Text) {
		return
	}
	log.Printf("%+v", ev.Text)
	commentCh <- ev.Text
}

func putMessageEvent(ev *slackevents.MessageEvent) {
	if ignoreMessage(ev.Text) {
		return
	}
	log.Printf("%+v", ev.Text)
	commentCh <- ev.Text
}

func ignoreMessage(msg string) bool {
	return msg == "This content can't be displayed."
}
