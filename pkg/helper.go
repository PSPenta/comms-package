package pkg

import (
	"comms-package/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CurlRequest(senderConfig *config.SenderInfo) (*http.Request, error) {

	requestPayload, err := json.Marshal(senderConfig.Payload)

	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(requestPayload))

	req, err := http.NewRequest(senderConfig.Method, senderConfig.URL, payload)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func CurlResponse(requst *http.Request) []byte {
	response, _ := http.DefaultClient.Do(requst)

	body, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	return body
}
