package model

import (
	"encoding/json"
)

// EventHandler defines a function which gets passed
// the event as instance pointer
type EventHandler func(*Event)

// Event contains event name and data
type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// NewEventFromRaw create an event object
// from raw binary JSON data
func NewEventFromRaw(rawData []byte) (*Event, error) {

	event := &Event{}

	err := json.Unmarshal(rawData, event)

	return event, err
}

// Raw Create binary JSON data from event
func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
