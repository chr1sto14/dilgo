package net

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchUrl(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to fetch %s", url))
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse %s", url))
	}
	return bytes, nil
}
