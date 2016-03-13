package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sger/go-apns2/certificate"
	"github.com/sger/go-apns2/client"
)

func main() {

	var deviceToken = "c7800a79efffe8ffc01b280717a936937cb69f8ca307545eb6983c60f12e167a"
	var filename = "certs/PushChatKey.p12"
	var password = "pushchat"

	// POST URL
	url := fmt.Sprintf("%v/3/device/%v", client.Development, deviceToken)

	// Setup payload must contains an aps root label and alert message
	payload := []byte(`{ "aps" : { "alert" : "Hello world" } }`)

	cert, key, err := p12.ReadFile(filename, password)
	if err != nil {
		log.Fatal(err)
	}

	certificate := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
		Leaf:        cert,
	}

	// Setup a new http client
	client, err := client.New(certificate)

	if err != nil {
		log.Fatal(err)
	}

	// Sending the request with valid PAYLOAD (must starts with aps)
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	// Send JSON Header
	// TODO
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read the response
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(body))
}
