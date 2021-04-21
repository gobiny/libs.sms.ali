package utils

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
	//response = string(body)
}
