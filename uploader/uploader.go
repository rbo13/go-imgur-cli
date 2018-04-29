package uploader

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

type Upload struct {
	File          *os.File
	Response      *http.Response
	ContentLength int
	Done          bool
}

func (upload *Upload) StartUpload() {
	_, err := io.Copy(upload.File, upload.Response.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
		upload.Done = true
		return
	}

	return
}

func NewUpload(fileName string) *Upload {
	upload := new(Upload)

	input, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("ERROR OPENING FILE DUE TO: %v", err)
		return nil
	}

	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		fmt.Printf("ERROR GETTING INPUT INFO DUE TO: %v", err)
		return nil
	}

	upload.ContentLength = int(math.Round(float64(inputInfo.Size()) / 1000))
	log.Println(upload.ContentLength)
	return upload
}
