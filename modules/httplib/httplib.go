package httplib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Request(method, url, params string, client *http.Client) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	bf := bytes.NewBufferString(params)
	req.Body = ioutil.NopCloser(bf)
	req.ContentLength = int64(len(params))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.124 Safari/537.36")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cache-Control", "max-age=0")

	if client == nil {
		client = &http.Client{}
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return resp.Body, nil
	}

	defer resp.Body.Close()
	if resp.StatusCode == 404 { // 403 can be rate limit error.  || resp.StatusCode == 403 {
		err = fmt.Errorf("resource not found: %s", url)
	} else {
		err = fmt.Errorf("get %s -> %d", url, resp.StatusCode)
	}
	return nil, err
}

func ResponseBytes(method, url, params string, client *http.Client) ([]byte, error) {
	rc, err := Request(method, url, params, client)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

func ResponseJSON(method, url, params string, client *http.Client, v interface{}) error {
	data, err := ResponseBytes(method, url, params, client)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}
