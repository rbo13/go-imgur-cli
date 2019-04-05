package imgur

import (
	"encoding/base64"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const imgurUploadURL = "https://api.imgur.com/3/image"

// Imgur sets the field for the required fields when uploading images to imgur.
type Imgur struct {
	apiKey string
}

// New returns pointer to imgur
func New(apiKey string) *Imgur {
	return &Imgur{
		apiKey: apiKey,
	}
}

// Upload uploads an image to imgur.
func (imgr *Imgur) Upload(filename string) {
	fileEncoded, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	parameters := url.Values{"image": {base64.StdEncoding.EncodeToString(fileEncoded)}}
	client := client(imgr.apiKey)

	req, err := http.NewRequest("POST", imgurUploadURL, strings.NewReader(parameters.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Client-ID "+imgr.apiKey)

	client.Do(req)
}

func client(apiKey string) *http.Client {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}
