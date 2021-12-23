package serviceutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostJSON(svc string, path string, query *url.Values, jwt string, h http.Header, d interface{}, r interface{}) (int, error) {
	url := buildUrl(svc, path, query)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(d)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, err
	}

	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	s := res.StatusCode

	if s != http.StatusOK {
		var e RemoteError
		err = json.Unmarshal(rb, &e)
		if err != nil {
			return 0, err
		}
		return s, err
	}

	err = json.Unmarshal(rb, r)
	return s, err
}
