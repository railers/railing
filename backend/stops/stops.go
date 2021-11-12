package stops

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Stops struct {
	Data []Stop `json:"data"`
}

type Stop struct {
	Id     string `json:"id"`
	Attrib struct {
		Name      string  `json:"name"`
		AtStreet  string  `json:"at_street"`
		OnStreet  string  `json:"on_street"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"attributes"`
}

func GetStops() *Stops {
	url := "https://api-v3.mbta.com/stops"

	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	stops := Stops{}
	err = json.Unmarshal(body, &stops)
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.MarshalIndent(stops, "", " ")

	fmt.Println(string(json))
	return &stops
}
