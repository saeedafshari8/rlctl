package util

import (
	"io/ioutil"
	"net/http"
)

func MakeHttpRequest(req *http.Request, ch chan<- ChannelResponse) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ch <- ChannelResponse{Error: err, Success: false}
		return
	}

	if response == nil || response.Body == nil {
		ch <- ChannelResponse{Error: err, Success: false}
		return
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- ChannelResponse{Error: err, Success: false}
		return
	}
	ch <- ChannelResponse{Error: nil, Success: true, Data: responseData}
}
