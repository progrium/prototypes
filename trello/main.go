package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/now"
	"github.com/mitchellh/hashstructure"

	"github.com/adlio/trello"
)

var (
	zeroTime  = time.Date(0, 0, 0, 0, 0, 0, 0, tz)
	tz        = time.Now().Location()
	tagLabels = []string{"tigl3d", "workbench"}
	schedule  = map[string][]Slot{
		"workbench": []Slot{
			{
				time.Tuesday,
				time.Date(0, 0, 0, 17, 0, 0, 0, tz),
				time.Hour * 2,
			},
			{
				time.Thursday,
				time.Date(0, 0, 0, 17, 0, 0, 0, tz),
				time.Hour * 2,
			},
		},
		"tigl3d": []Slot{
			{
				time.Monday,
				time.Date(0, 0, 0, 15, 0, 0, 0, tz),
				time.Hour * 4,
			},
			{
				time.Tuesday,
				time.Date(0, 0, 0, 14, 0, 0, 0, tz),
				time.Hour * 2,
			},
			{
				time.Wednesday,
				time.Date(0, 0, 0, 15, 0, 0, 0, tz),
				time.Hour * 4,
			},
			{
				time.Thursday,
				time.Date(0, 0, 0, 14, 0, 0, 0, tz),
				time.Hour * 2,
			},
		},
	}
)

func main() {
	resync := false

	state, err := LoadState()
	fatal(err)

	client := trello.NewClient(os.Getenv("TRELLO_KEY"), os.Getenv("TRELLO_TOKEN"))
	board, err := client.GetBoard("EsjNNP3c", trello.Defaults())
	fatal(err)
	lists, err := board.GetLists(trello.Defaults())
	fatal(err)
	var events []Event
	for _, l := range lists {
		if l.Name == "Schedule" {
			cards, err := l.GetCards(trello.Defaults())
			fatal(err)
			for _, c := range cards {
				events = append(events, NewEvent(c))
			}
			break
		}
	}
	trelloHash := HashEvents(events)
	if state.TrelloHash != trelloHash {
		resync = true
	}

	twitch, err := LoadTwitchMock()
	fatal(err)
	for id, hash := range state.TwitchHashes {
		e, _ := twitch.EventByID(id)
		if hash != HashTwitchEvent(e) {
			resync = true
			break
		}
	}

	now.WeekStartDay = time.Monday
	thisWeek := now.BeginningOfWeek()
	nextWeek := now.New(now.EndOfWeek().Add(time.Hour)).BeginningOfWeek()

	if resync {
		fmt.Println("Resyncing...")
		state.TrelloHash = trelloHash
		var event *Event
		for _, weekBegin := range []time.Time{thisWeek, nextWeek} {
			for tag, slots := range schedule {
				for _, slot := range slots {
					if slot.EndTime(weekBegin).Before(time.Now()) {
						continue
					}
					event, events = ShiftEvent(events, tag)
					if event == nil {
						continue
					}
					twitchEvent, _ := twitch.EventAt(slot.StartTime(weekBegin))
					var err error
					if twitchEvent == nil {
						fmt.Println("NEW:", event.Name)
						twitchEvent, err = twitch.Post(TwitchEvent{
							Title:       event.Name,
							Description: event.Description + "\n#" + tag,
							StartTime:   slot.StartTime(weekBegin),
							EndTime:     slot.EndTime(weekBegin),
						})
						fatal(err)
					} else {
						fmt.Println("MOD:", event.Name)
						twitchEvent, err = twitch.Put(twitchEvent.ID, TwitchEvent{
							Title:       event.Name,
							Description: event.Description + "\n#" + tag,
							StartTime:   slot.StartTime(weekBegin),
							EndTime:     slot.EndTime(weekBegin),
						})
						fatal(err)
					}
					state.TwitchHashes[twitchEvent.ID] = HashTwitchEvent(twitchEvent)
					state.Mapping[event.ID] = twitchEvent.ID
				}
			}
		}

	}

	EnsureSlotEvents(twitch, thisWeek)
	EnsureSlotEvents(twitch, nextWeek)

	fatal(SaveTwitchMock(twitch))
	fatal(SaveState(state))
}

func EnsureSlotEvents(twitch TwitchAPI, weekBegin time.Time) {
	for tag, slots := range schedule {
		for _, slot := range slots {
			if slot.EndTime(weekBegin).Before(time.Now()) {
				continue
			}
			e, _ := twitch.EventAt(slot.StartTime(weekBegin))
			if e == nil {
				twitch.Post(TwitchEvent{
					Title:       "TBD",
					Description: "#" + tag,
					StartTime:   slot.StartTime(weekBegin),
					EndTime:     slot.EndTime(weekBegin),
				})
			}
		}
	}
}

func ShiftEvent(events []Event, tag string) (*Event, []Event) {
	var idx *int
	for i, e := range events {
		if e.Tag == tag {
			idx = &i
			break
		}
	}
	if idx == nil {
		return nil, events
	}
	e := events[*idx]
	return &e, append(events[:*idx], events[(*idx)+1:]...)
}

func HashEvents(events []Event) string {
	hash, err := hashstructure.Hash(events, nil)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hash)
}

func HashTwitchEvent(event *TwitchEvent) string {
	hash, err := hashstructure.Hash(event, nil)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hash)
}

func LabelNames(labels []*trello.Label) []string {
	var names []string
	for _, l := range labels {
		names = append(names, l.Name)
	}
	return names
}
