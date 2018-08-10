package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/adlio/trello"
)

type Event struct {
	ID          string
	Name        string
	Description string
	Tag         string
	Labels      []string
}

func NewEvent(card *trello.Card) Event {
	labels := LabelNames(card.Labels)
	var tagIdx *int
	for i, l := range labels {
		for _, tl := range tagLabels {
			if l == tl {
				tagIdx = &i
			}
		}
	}
	var tag string
	if tagIdx != nil {
		idx := *tagIdx
		tag = labels[idx]
		copy(labels[idx:], labels[idx+1:])
		labels[len(labels)-1] = ""
		labels = labels[:len(labels)-1]
	}
	if tag == "" {
		tag = "workbench"
	}
	return Event{
		ID:          card.ID,
		Name:        card.Name,
		Description: card.Desc,
		Labels:      labels,
		Tag:         tag,
	}
}

type TwitchEvent struct {
	ID          string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

func (e *TwitchEvent) Tag() string {
	parts := strings.Split(e.Description, "#")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

type Slot struct {
	Weekday  time.Weekday
	Time     time.Time
	Duration time.Duration
}

func (s *Slot) StartTime(weekBegin time.Time) time.Time {
	weekOffset := (int(s.Weekday) - 1) * 24
	return weekBegin.Add(time.Hour * time.Duration(weekOffset)).Add(s.Time.Sub(zeroTime))
}

func (s *Slot) EndTime(weekBegin time.Time) time.Time {
	return s.StartTime(weekBegin).Add(s.Duration)
}

type SyncState struct {
	TrelloHash   string
	Mapping      map[string]string
	TwitchHashes map[string]string
}

func NewSyncState() *SyncState {
	return &SyncState{
		Mapping:      make(map[string]string),
		TwitchHashes: make(map[string]string),
	}
}

func LoadState() (*SyncState, error) {
	b, err := ioutil.ReadFile("state.json")
	if err != nil {
		return NewSyncState(), nil
	}
	if len(b) == 0 {
		return NewSyncState(), nil
	}
	var state SyncState
	err = json.Unmarshal(b, &state)
	return &state, err
}

func SaveState(s *SyncState) error {
	buf, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("state.json", buf, 0644)
}
