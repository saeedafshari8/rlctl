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
	}
	defer response.Body.Close()

	if err != nil {
		ch <- ChannelResponse{Error: err, Success: false}
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- ChannelResponse{Error: err, Success: false}
	}
	ch <- ChannelResponse{Error: nil, Success: true, Data: responseData}
}
