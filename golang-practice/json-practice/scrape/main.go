package main
// http://date.jsontest.com/
// https://www.youtube.com/watch?v=aYk8XAKxhxU

import (
	"fmt"
	"encoding/json"
    "net/http"
    "time"

)

type DateTime struct {
	Date string `json:"date"`
	Milisecond int `json:"milliseconds_since_epoch"`
	Time string `json:"time"`
}

var client *http.Client

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)


	if err != nil {
		return err
	}
    defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)


}

func getIP() {
	var dateTime DateTime
	err := getJson("http://date.jsontest.com/", &dateTime)
	if err != nil {
		fmt.Println(err.Error)
	} else {
		fmt.Println(dateTime)
	}

}


func main() {
	client = &http.Client{Timeout: 10 * time.Second}


	getIP()

}