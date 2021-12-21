package serviceutils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Get(svc string, path string, query *url.Values, jwt string, h http.Header) ([]byte, error) {
	url := buildUrl(svc, path, query)

	log.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	return rb, err
}
