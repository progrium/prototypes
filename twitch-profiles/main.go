package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kr/pretty"
)

const (
	ChannelID = "5031651"

	CommunityGolang          = "003860d3-270f-4082-a8f7-b1aa926272f4"
	CommunityOpensource      = "4ca22a66-fbfe-4b82-8433-40b9509bc913"
	CommunityProgramming     = "9d175334-ccdd-4da8-a3aa-d9631f95610e"
	CommunityUnity3D         = "beed41df-c336-40a3-ae50-db9909b360f1"
	CommunityGameDevelopment = "5d29f46d-4ac3-4fac-92e4-dbd15be5ff6f"
)

type ChannelUpdate struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Status string `json:"status"`
}

func UpdateChannel(update ChannelUpdate) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&update)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", "https://api.twitch.tv/kraken/channels/"+ChannelID, buf)
	if err != nil {
		return err
	}
	var resp map[string]interface{}
	_, err = Do(req, &resp)
	return err
}

func main() {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/communities?name=gamedevelopment", nil)
	if err != nil {
		panic(err)
	}
	var resp map[string]interface{}
	_, err = Do(req, &resp)
	fmt.Printf("%# v", pretty.Formatter(resp))
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
