package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	v := url.Values{}
	v.Set("client_id", "tf99t1lw9dcsxpprca8h4uwew53yos")
	v.Set("redirect_uri", "http://localhost:8080")
	v.Set("scope", "user_read channel_editor channel_read")
	v.Set("response_type", "token")
	fmt.Println("https://id.twitch.tv/oauth2/authorize?" + v.Encode())
	// resp, err := http.Get("https://id.twitch.tv/oauth2/authorize?" + v.Encode())
	// if err != nil {
	// 	panic(err)
	// }
	// err = checkResponse(resp)
	// if err != nil {
	// 	panic(err)
	// }
	// // var r map[string]interface{}
	// // err = json.NewDecoder(resp.Body).Decode(&r)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// fmt.Println(resp)
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
