package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	filename := "test.txt"

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		log.Fatalln("error writing to buffer", err)
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatalln("error opening file", err)
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Fatalln("something went wrong", err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post("http://localhost:9090/upload", contentType, bodyBuf)
	if err != nil {
		log.Fatalln("something went wrong", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("something went wrong", err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
}
