package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	spin "github.com/tj/go-spin"
)

var clientID = getEnvVar()

const (
	imgur = "https://api.imgur.com/3/image"
)

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
	s := spin.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// uploader.NewUpload(fileName)

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

	showSpinner(s, "Box1", spin.Box1)

	var imgurResponse ImgurResponse
	json.NewDecoder(r.Body).Decode(&imgurResponse)
	copyToClipboard(imgurResponse.Data.Link)
	fmt.Printf("\n`%s` has been copied to clipboard! \n", imgurResponse.Data.Link)
	fmt.Println("Deletion Link: http://imgur.com/delete/" + imgurResponse.Data.Deletehash)
}

func getEnvVar() string {
	env := os.Getenv("IMGUR_CLIENT_ID")

	if env == "" {
		return ""
	}

	return env
}

func getCopyCommand() *exec.Cmd {
	return exec.Command("pbcopy")
}

func copyToClipboard(text string) error {
	copyCmd := getCopyCommand()
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}

func showSpinner(s *spin.Spinner, name, frames string) {
	s.Set(frames)
	fmt.Printf("\n\n  %s: %s\n\n", name, frames)
	for i := 0; i < 30; i++ {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}
