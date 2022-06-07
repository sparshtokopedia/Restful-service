package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	ID          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type EventAll []Event

var events = EventAll{
	{
		ID:          "1",
		Title:       "Go",
		Description: "GO",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter title and description ")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter title and description ")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func homeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeMessage)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	log.Println("--Server started on :8080--")
	log.Fatal(http.ListenAndServe(":8080", router))

}
