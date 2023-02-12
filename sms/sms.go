package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func SendSMS(message string, phone string) {
	endpoint := "https://rest.nexmo.com/sms/json"
	data := url.Values{}

	api_key := os.Getenv("NEXMO_API_KEY")
	if api_key == "" {
		panic("NEXMO_API_KEY não foi configurado!")
	}
	api_secret := os.Getenv("NEXMO_API_SECRET")
	if api_secret == "" {
		panic("NEXMO_API_SECRET não foi configurado!")
	}

	data.Set("api_key", api_key)
	data.Set("api_secret", api_secret)
	data.Set("to", phone)
	data.Set("from", "Sistema de alertas XGH")
	data.Set("text", message)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	r.Header.Add("content-type", "application/x-www-form-urlencoded")
	r.Header.Add("content-length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}
