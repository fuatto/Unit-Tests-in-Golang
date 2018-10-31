package types

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//People is a struct with Number attribute indicating a number of people
type People struct {
	Number int `json:"number"`
}

//GetWebRequest interface contains FetchBytes function
type GetWebRequest interface {
	FetchBytes(url string) []byte
}

//LiveGetWebRequest struct will be a receiver in FetchBytes function implemented for the GetWebRequest interface
type LiveGetWebRequest struct{}

//FetchBytes function is implemented with LiveGetWebRequest receiver
func (LiveGetWebRequest) FetchBytes(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

//GetAstronauts function parses and returns a number of astronauts in the astros.json file
func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"

	bodyBytes := getWebRequest.FetchBytes(url)
	peopleResult := People{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return peopleResult.Number
}
