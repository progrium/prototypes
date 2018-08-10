package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

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

func main() {
	flag.Parse()

	service, err := youtube.New(getClient())
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	call := service.LiveBroadcasts.List("snippet")
	call.BroadcastType("persistent")
	call.Mine(true)

	resp, err := call.Do()
	if err != nil {
		// The channels.list method call returned an error.
		log.Fatalf("Error making API call: %v", err.Error())
	}

	for _, b := range resp.Items {
		b.Snippet.IsDefaultBroadcast
		fmt.Println(b)
	}
	// 	playlistId := channel.ContentDetails.RelatedPlaylists.Uploads
	// 	// Print the playlist ID for the list of uploaded videos.
	// 	fmt.Printf("Videos in list %s\r\n", playlistId)

	// 	nextPageToken := ""
	// 	for {
	// 		// Call the playlistItems.list method to retrieve the
	// 		// list of uploaded videos. Each request retrieves 50
	// 		// videos until all videos have been retrieved.
	// 		playlistCall := service.PlaylistItems.List("snippet").
	// 			PlaylistId(playlistId).
	// 			MaxResults(50).
	// 			PageToken(nextPageToken)

	// 		playlistResponse, err := playlistCall.Do()

	// 		if err != nil {
	// 			// The playlistItems.list method call returned an error.
	// 			log.Fatalf("Error fetching playlist items: %v", err.Error())
	// 		}

	// 		for _, playlistItem := range playlistResponse.Items {
	// 			title := playlistItem.Snippet.Title
	// 			videoId := playlistItem.Snippet.ResourceId.VideoId
	// 			fmt.Printf("%v, (%v)\r\n", title, videoId)
	// 		}

	// 		// Set the token to retrieve the next page of results
	// 		// or exit the loop if all results have been retrieved.
	// 		nextPageToken = playlistResponse.NextPageToken
	// 		if nextPageToken == "" {
	// 			break
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	// Start making YouTube API calls.
	// Call the channels.list method. Set the mine parameter to true to
	// retrieve the playlist ID for uploads to the authenticated user's
	// channel.
	// call := service.LiveBroadcasts.Update("snippet", &youtube.LiveBroadcast{
	// 	Id: "b072fLxm7sM",
	// 	Snippet: &youtube.LiveBroadcastSnippet{
	// 		Title:       "Hello world",
	// 		Description: "HEre is a description",
	// 	},
	// })

	// _, err = call.Do()
	// if err != nil {
	// 	// The channels.list method call returned an error.
	// 	log.Fatalf("Error making API call: %v", err.Error())
	// }
}
