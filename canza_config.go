package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SMSRequestBody ...
type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func main() {

	body := SMSRequestBody{
		APIKey:    os.Getenv("NEXMO_API_KEY"),
		APISecret: os.Getenv("NEXMP_API_SECRET"),
		To:        "PHONE_NUMBER",
		From:      "...",
		Text:      "Hello! how are you?",
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respBody))
}