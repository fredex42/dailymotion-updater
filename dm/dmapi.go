package dm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var BASE_URL = "https://api.dailymotion.com"
var MAX_ATTEMPTS int16 = 100

/**
read and close the HTTP body
*/
func readBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	rtn, readErr := ioutil.ReadAll(response.Body)

	return rtn, readErr
}

/**
perform a GET request to the DM API
*/
func makeGetRequest(subpath string, attempt int16) ([]byte, error) {
	response, httpErr := http.Get(BASE_URL + subpath)

	log.Printf("URL is %s", BASE_URL+subpath)
	if httpErr != nil {
		return nil, httpErr
	} else {
		switch response.StatusCode {
		case 200:
			return readBody(response)
		case 400:
			body, readErr := readBody(response)
			if readErr != nil {
				return nil, readErr
			} else {
				bodyString := string(body)
				errMsg := fmt.Sprintf("API returned bad data: %s", bodyString)
				return nil, errors.New(errMsg)
			}
		case 403:
			body, readErr := readBody(response)
			if readErr != nil {
				return nil, readErr
			} else {
				bodyString := string(body)
				errMsg := fmt.Sprintf("API returned permission denied: %s", bodyString)
				return nil, errors.New(errMsg)
			}
		case 500:
		case 502:
		case 503:
		case 504:
			body, readErr := readBody(response)
			if readErr != nil {
				return nil, readErr
			} else {
				bodyString := string(body)
				log.Printf("{%d/%d}: API returned server error, retrying in 3s: %s", attempt, MAX_ATTEMPTS, bodyString)
				time.Sleep(3 * time.Second)
				if attempt >= MAX_ATTEMPTS {
					return nil, errors.New("failed after exhausting retries")
				}
				return makeGetRequest(subpath, attempt+1)
			}
		default:
			body, readErr := readBody(response)
			if readErr != nil {
				return nil, readErr
			} else {
				bodyString := string(body)
				errMsg := fmt.Sprintf("API returned unexpected status %d: %s", response.StatusCode, bodyString)
				return nil, errors.New(errMsg)
			}
		}
	}
	return nil, errors.New("internal error, should not reach this point")
}

func GetChannels() (*[]DMChannel, error) {
	var rtn []DMChannel
	page := 1 //minimum page is 1
	for {
		url := fmt.Sprintf("/channels?page=%d", page)
		bytes, getErr := makeGetRequest(url, 0)
		if getErr != nil {
			return nil, getErr
		}

		var channelList DMChannelList
		marshalErr := json.Unmarshal(bytes, &channelList)
		if marshalErr != nil {
			return nil, marshalErr
		}

		rtn = append(rtn, channelList.List...)
		if channelList.HasMore == false {
			break
		}
	}

	return &rtn, nil
}
