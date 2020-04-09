package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	morpheus "github.com/gomorpheus/morpheus-go-sdk"
)

const (
	morpheusTokenEnv = "MORPH_API"
	morpheusURL      = "MORPH_URL"
)

var errMissingMorphToken = errors.New("Missing " + morpheusTokenEnv + " environment variable")
var errMissingMorphURL = errors.New("Missing " + morpheusURL + " environment variable")

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Whoami()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
}

func newClient() (*morpheus.Client, error) {
	url, ok := os.LookupEnv(morpheusURL)
	if !ok {
		return nil, errMissingMorphURL
	}

	token, ok := os.LookupEnv(morpheusTokenEnv)
	if !ok {
		return nil, errMissingMorphToken
	}

	client := morpheus.NewClient(url)
	client.SetAccessToken(token, "", 0, "write")

	return client, nil
}
