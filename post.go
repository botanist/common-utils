package serviceutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostJSON(svc string, path string, query *url.Values, jwt string, h http.Header, d interface{}) ([]byte, error) {
	url := buildUrl(svc, path, query)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(d)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	return rb, err
}
