package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	data := url.Values{
		"key": {os.Getenv("KEY")},
	}

	resp, err := http.PostForm("https://www.googleapis.com/geolocation/v1/geolocate", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["location"])
}
