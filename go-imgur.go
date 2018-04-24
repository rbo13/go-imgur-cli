package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var clientID = GetEnvVar()
var imgur = "https://api.imgur.com/3/image"

type ImgurResponse struct {
	Data    ImageData `json:"data"`
	Status  int       `json:"status"`
	Success bool      `json:"success"`
}

type ImageData struct {
	Account_id int    `json:"account_id"`
	Animated   bool   `json:"animated"`
	Bandwidth  int    `json:"bandwidth"`
	DateTime   int    `json:"datetime"`
	Deletehash string `json:"deletehash"`
	Favorite   bool   `json:"favorite"`
	Height     int    `json:"height"`
	Id         string `json:"id"`
	In_gallery bool   `json:"in_gallery"`
	Is_ad      bool   `json:"is_ad"`
	Link       string `json:"link"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Views      int    `json:"views"`
	Width      int    `json:"width"`
}

func main() {
	fileName := os.Args[1]
	fileEncoded, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parameters := url.Values{"image": {base64.StdEncoding.EncodeToString(fileEncoded)}}
	req, err := http.NewRequest("POST", imgur, strings.NewReader(parameters.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Client-ID "+clientID)
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var imgurResponse ImgurResponse
	json.NewDecoder(r.Body).Decode(&imgurResponse)
	fmt.Println("Image Link: " + imgurResponse.Data.Link)
	fmt.Println("Deletion Link: http://imgur.com/delete/" + imgurResponse.Data.Deletehash)
}

func GetEnvVar() string {
	env := os.Getenv("IMGUR_SECRET")

	if env == "" {
		return ""
	}

	return env
}
