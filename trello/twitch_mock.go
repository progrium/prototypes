package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type TwitchAPI interface {
	FutureEvents() ([]TwitchEvent, error)
	EventAt(t time.Time) (*TwitchEvent, error)
	EventByID(id string) (*TwitchEvent, error)
	Post(e TwitchEvent) (*TwitchEvent, error)
	Delete(id string) error
	Put(id string, e TwitchEvent) (*TwitchEvent, error)
}

type TwitchAPIMock struct {
	Events []TwitchEvent
}

func (ft *TwitchAPIMock) FutureEvents() ([]TwitchEvent, error) {
	var events []TwitchEvent
	for _, event := range ft.Events {
		if event.StartTime.After(time.Now()) {
			events = append(events, event)
		}
	}
	sort.Slice(events, func(i, j int) bool { return events[i].StartTime.Before(events[j].StartTime) })
	return events, nil
}

func (ft *TwitchAPIMock) EventAt(t time.Time) (*TwitchEvent, error) {
	for _, event := range ft.Events {
		if event.StartTime == t {
			return &event, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (ft *TwitchAPIMock) EventByID(id string) (*TwitchEvent, error) {
	for _, event := range ft.Events {
		if event.ID == id {
			return &event, nil
		}
	}
	return nil, fmt.Errorf("not found: %s", id)
}

func (ft *TwitchAPIMock) Post(e TwitchEvent) (*TwitchEvent, error) {
	e.ID = RandomString(10)
	ft.Events = append(ft.Events, e)
	return &e, nil
}

func (ft *TwitchAPIMock) Delete(id string) error {
	var idx *int
	for i, e := range ft.Events {
		if e.ID == id {
			idx = &i
			break
		}
	}
	if idx == nil {
		return fmt.Errorf("not found: %s", id)
	}
	ft.Events = append(ft.Events[:*idx], ft.Events[(*idx)+1:]...)
	return nil
}

func (ft *TwitchAPIMock) Put(id string, e TwitchEvent) (*TwitchEvent, error) {
	var idx *int
	for i, ee := range ft.Events {
		if ee.ID == id {
			idx = &i
			break
		}
	}
	if idx == nil {
		return nil, fmt.Errorf("not found: %s", id)
	}
	e.ID = id
	ft.Events[*idx] = e
	return &e, nil
}

func LoadTwitchMock() (*TwitchAPIMock, error) {
	b, err := ioutil.ReadFile("twitch.json")
	if err != nil {
		return &TwitchAPIMock{}, nil
	}
	if len(b) == 0 {
		return &TwitchAPIMock{}, nil
	}
	var api TwitchAPIMock
	err = json.Unmarshal(b, &api)
	return &api, err
}

func SaveTwitchMock(t *TwitchAPIMock) error {
	buf, err := json.MarshalIndent(*t, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("twitch.json", buf, 0644)
}
