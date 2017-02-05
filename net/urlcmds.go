package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github/chr1sto14/dilbert/formathipchat"
	"io/ioutil"
	"net/http"
)

func PostMsg(url string, msg formathipchat.Message) error {

	jsonStr, err := json.Marshal(msg)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to marshal msg"))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create request"))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to post to %s", url))
	}
	defer resp.Body.Close()

	return nil
}

func FetchUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to fetch %s", url))
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to interpret retrieved url"))
	}
	return data, nil
}
