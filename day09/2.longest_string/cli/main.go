package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	result, err := MakeRequest(os.Args)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	fmt.Println(result)
}

func MakeRequest(args []string) (string, error) {
	if len(args) != 3 {
		return "", errors.New("Incorrect number of arguments")
	}

	cli := http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s", args[2]), bytes.NewReader([]byte(args[1])))
	if err != nil {
		return "", err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
