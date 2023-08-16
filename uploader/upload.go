package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	url := "http://52.35.27.227:8888/srv-applet-mgr/v0/event/receiver"
	method := "POST"

	jsonFile, err := os.Open("SGW_dbc.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := strings.Replace(string(byteValue), "\n", "", -1)
	data = strings.Replace(data, "\t", "", -1)
	data = strings.Replace(data, "\r", "", -1)
	data = strings.Replace(data, "\"", "\\\"", -1)
	payload := strings.NewReader(`{
		"header": {
			"event_type": "DEFAULT",
			"pub_id": "receiver",
			"pub_time": 1676074425,
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjoiMTEyNjk2NTA1NjEyODYxNDgiLCJpc3MiOiJ3M2JzdHJlYW0iLCJleHAiOjE2NzYwNzU2NzJ9.Kv4J7tYlUz8p6JTlSHU4zpUmIfEsljNJ7n65Ox7fQVY"                                
		},
			"payload":"` + data + `"
		}`)
	println(`{
			"header": {
				"event_type": "DEFAULT",
				"pub_id": "receiver",
				"pub_time": 1676074425,
				"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjoiMTEyNjk2NTA1NjEyODYxNDgiLCJpc3MiOiJ3M2JzdHJlYW0iLCJleHAiOjE2NzYwNzU2NzJ9.Kv4J7tYlUz8p6JTlSHU4zpUmIfEsljNJ7n65Ox7fQVY"                                
			},
				"payload":"` + data + `"
			}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
