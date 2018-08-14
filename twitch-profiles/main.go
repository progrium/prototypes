package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	youtube "google.golang.org/api/youtube/v3"
)

const (
	ProgriumChannelID = "5031651"

	StreamNowDescription = "Chat here is ignored, please chat on Twitch: https://twitch.tv/progrium"

	CommunityGolang          = "003860d3-270f-4082-a8f7-b1aa926272f4"
	CommunityOpensource      = "4ca22a66-fbfe-4b82-8433-40b9509bc913"
	CommunityProgramming     = "9d175334-ccdd-4da8-a3aa-d9631f95610e"
	CommunityUnity3D         = "beed41df-c336-40a3-ae50-db9909b360f1"
	CommunityGameDevelopment = "5d29f46d-4ac3-4fac-92e4-dbd15be5ff6f"
	CommunityMusic           = "ec04cef0-0e81-4fa9-a037-d11ac87051b6"
	CommunityLogicPro        = "970a2ae5-ff30-40f3-9bac-7b4e3f3999f0"
)

var profiles = []Profile{
	{
		Name:        "tigl3d",
		Tag:         "[TIGL3D]",
		Communities: []string{CommunityProgramming, CommunityUnity3D, CommunityGameDevelopment},
	},
	{
		Name:        "gcl",
		Tag:         "[GCL]",
		Communities: []string{CommunityProgramming, CommunityGolang, CommunityOpensource},
	},
	{
		Name:        "music",
		Tag:         "[Music]",
		Communities: []string{CommunityMusic, CommunityLogicPro},
	},
}

type Profile struct {
	Name        string
	Tag         string
	Communities []string
}

type ChannelUpdate struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Status string `json:"status"`
}

type EventsResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
}

func UpdateChannel(channelID string, update ChannelUpdate) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&update)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.twitch.tv/kraken/channels/%s", channelID), buf)
	if err != nil {
		return err
	}
	var resp map[string]interface{}
	_, err = Do(req, &resp)
	return err
}

func UpdateChannelCommunities(channelID string, communityIDs ...string) error {
	buf := new(bytes.Buffer)
	payload := make(map[string]interface{})
	payload["community_ids"] = communityIDs
	err := json.NewEncoder(buf).Encode(&payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.twitch.tv/kraken/channels/%s/communities", channelID), buf)
	if err != nil {
		return err
	}
	var resp map[string]interface{}
	_, err = Do(req, &resp)
	return err
}

func FetchEvents(channelID string) ([]Event, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.twitch.tv/kraken/channels/%s/events", channelID), nil)
	if err != nil {
		return nil, err
	}
	var resp EventsResponse
	_, err = Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Events, err
}

func NextEvent(channelID string, descFilter string) (Event, error) {
	events, err := FetchEvents(channelID)
	if err != nil {
		return Event{}, err
	}
	for _, event := range events {
		if strings.Contains(event.Description, descFilter) {
			return event, nil
		}
	}
	return Event{}, errors.New("event not found")
}

func main() {
	flag.Parse()
	profile := flag.Arg(0)
	status := flag.Arg(1)
	for _, p := range profiles {
		if p.Name == profile {
			if status == "" {
				event, _ := NextEvent(ProgriumChannelID, "#"+p.Name)
				if event.Title != "" {
					status = event.Title
				} else {
					log.Fatal("enter status or make sure event exists")
				}
			}
			fullStatus := fmt.Sprintf("%s %s", status, p.Tag)

			// Update YouTube
			broadcast, err := GetDefaultBroadcast()
			if err != nil {
				log.Fatal(err)
			}
			err = UpdateLiveBroadcastSnippet(broadcast.Id, &youtube.LiveBroadcastSnippet{
				Title:       fullStatus,
				Description: StreamNowDescription,
			})
			if err != nil {
				log.Fatal(err)
			}

			// Update Twitch
			err = UpdateChannel(ProgriumChannelID, ChannelUpdate{
				Channel: Channel{
					Status: fullStatus,
				},
			})
			if err != nil {
				log.Fatal(err)
			}
			err = UpdateChannelCommunities(ProgriumChannelID, p.Communities...)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("OK:", fullStatus)
			return
		}
	}
	log.Fatal("profile not found")
}

func Do(req *http.Request, r interface{}) (*http.Response, error) {
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Set("User-Agent", "progrium")
	req.Header.Set("Client-ID", "tf99t1lw9dcsxpprca8h4uwew53yos")
	req.Header.Set("Authorization", "OAuth "+os.Getenv("TWITCH_OAUTH_TOKEN"))
	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = checkResponse(resp); err != nil {
		return resp, err
	}

	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
		if err == io.EOF {
			err = nil
		}
	}
	return resp, err
}

type ErrorResponse struct {
	// HTTP response that cause this error.
	Response *http.Response

	// Error message.
	Message string `json:"message,omitempty"`
}

func checkResponse(r *http.Response) error {
	if 200 <= r.StatusCode && r.StatusCode <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (e *ErrorResponse) Error() string {
	r := e.Response

	return fmt.Sprintf("%v %v: %d %v",
		r.Request.Method, r.Request.URL, r.StatusCode, e.Message)
}

// GOOGLE / YOUTUBE STUFF

func GetDefaultBroadcast() (*youtube.LiveBroadcast, error) {
	service, err := youtube.New(getClient())
	if err != nil {
		return nil, err
	}
	call := service.LiveBroadcasts.List("snippet")
	call.BroadcastType("persistent")
	call.Mine(true)
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}
	for _, b := range resp.Items {
		if b.Snippet.IsDefaultBroadcast {
			return b, nil
		}
	}
	return nil, nil
}

func UpdateLiveBroadcastSnippet(broadcastID string, snippet *youtube.LiveBroadcastSnippet) error {
	service, err := youtube.New(getClient())
	if err != nil {
		return err
	}
	call := service.LiveBroadcasts.Update("snippet", &youtube.LiveBroadcast{
		Id:      broadcastID,
		Snippet: snippet,
	})
	_, err = call.Do()
	return err
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient() *http.Client {
	b, err := ioutil.ReadFile("/Users/progrium/.config/youtube_client_id.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, youtube.YoutubeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	codeCh := make(chan string)
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		codeCh <- r.URL.Query().Get("code")
	})
	h := &http.Server{Addr: ":8080", Handler: mux}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	// fmt.Printf("Go to the following link in your browser then type the "+
	// 	"authorization code: \n%v\n", authURL)
	go h.ListenAndServe()

	openURL(authURL)

	authCode := <-codeCh
	go h.Shutdown(context.Background())
	// var authCode string
	// if _, err := fmt.Scan(&authCode); err != nil {
	// 	log.Fatalf("Unable to read authorization code %v", err)
	// }

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}
