package serviceutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(svc string, path string, query *url.Values, jwt string, h http.Header) (int, []byte, error) {
	url := buildUrl(svc, path, query)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}

	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, nil, err
	}

	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	return res.StatusCode, rb, err
}

func GetJSON(svc string, path string, query *url.Values, jwt string, h http.Header, t interface{}) (int, error) {
	s, b, err := Get(svc, path, query, jwt, h)
	if err != nil {
		return s, err
	}

	if s != http.StatusOK {
		var e RemoteError
		err = json.Unmarshal(b, &e)
		if err != nil {
			return 0, err
		}
		return s, err
	}

	err = json.Unmarshal(b, t)
	return s, err
}
