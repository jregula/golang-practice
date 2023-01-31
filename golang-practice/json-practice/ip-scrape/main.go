package main
// 

import (
	"fmt"
	"encoding/json"
    "net/http"
)

type Response struct {
	Ip string `json:"ip"`
}

var response Response

func main() {
	resp, err := http.Get("http://ip.jsontest.com/")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
    fmt.Println("Response status:", resp.Status)

	d := json.NewDecoder(resp.Body)

	if err := d.Decode(&response); err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Ip)

}