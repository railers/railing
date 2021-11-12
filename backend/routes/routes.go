package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Routes struct {
	Data []Route `json:"data"`
}

type Route struct {
	Id     string `json:"id"`
	Attrib struct {
		// populate
	} `json:"attributes"`
}

func GetRoutes() *Routes {
	url := "https://api-v3.mbta.com/routes"

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

	routes := Routes{}
	err = json.Unmarshal(body, &routes)
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.MarshalIndent(routes, "", " ")

	fmt.Println(string(json))
	return &routes
}
