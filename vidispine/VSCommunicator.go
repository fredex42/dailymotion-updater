package vidispine

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type VidispineCommunicator struct {
	Protocol string
	Hostname string
	Port     int16
	User     string
	Password string
}

/**
read and close the HTTP body
*/
func readBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	rtn, readErr := ioutil.ReadAll(response.Body)

	return rtn, readErr
}

func handleResponse(response *http.Response) ([]byte, error) {
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
			errMsg := fmt.Sprintf("API returned not available: %s", bodyString)
			return nil, errors.New(errMsg)
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
	return nil, errors.New("code bug, should not reach this point")
}

/**
assembles the matrix parameters to a url portion
*/
func assembleMatrixParams(params map[string]string) string {
	var rtn []string

	for k, v := range params {
		rtn = append(rtn, fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v)))
	}
	if len(rtn) > 0 {
		return ";" + strings.Join(rtn, ";")
	} else {
		return ""
	}
}

/**
assembles the query parameters to a url portion
*/
func assembleQueryParams(params map[string]string) string {
	var rtn []string

	for k, v := range params {
		rtn = append(rtn, fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v)))
	}

	if len(rtn) > 0 {
		return "?" + strings.Join(rtn, "&")
	} else {
		return ""
	}
}

/**
builds a URL string out of all the parts we have
*/
func (comm *VidispineCommunicator) assembleUrl(subpath string, matrixParams map[string]string, queryParams map[string]string) string {
	//remove any leading / from subpath
	var actualSubpath string
	if subpath[0] == '/' {
		actualSubpath = subpath[1:]
	} else {
		actualSubpath = subpath
	}
	return fmt.Sprintf("%s://%s:%d/%s%s%s", comm.Protocol, comm.Hostname, comm.Port, actualSubpath, assembleMatrixParams(matrixParams), assembleQueryParams(queryParams))
}

/**
perform a request to the server
*/
func (comm *VidispineCommunicator) MakeRequest(verb string, subpath string, matrixParams map[string]string, queryParams map[string]string, headers map[string]string, body io.Reader) ([]byte, error) {
	client := &http.Client{}

	url := comm.assembleUrl(subpath, matrixParams, queryParams)

	log.Printf("Connecting to %s", url)
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(comm.User, comm.Password)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	for {
		response, doErr := client.Do(req)

		if doErr != nil {
			return nil, err
		}

		rtn, responseErr := handleResponse(response)

		if response.StatusCode == 502 || response.StatusCode == 503 {
			log.Printf("Got a server unavailable error, retrying in 3s...")
			time.Sleep(3 * time.Second)
		} else {
			return rtn, responseErr
		}
	}
}
