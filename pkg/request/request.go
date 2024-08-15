package request

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{Timeout: 5 * time.Second, Transport: tr, CheckRedirect: nil, Jar: jar}
	return &Client{client: &httpClient}
}

func (client *Client) MakeRequest(method, urlRequest string, data map[string]interface{}) ([]byte, error) {

	req, err := client.MakeNewRequest(method, urlRequest, data)
	if err != nil {
		return nil, err
	}
	res, er := client.client.Do(req)
	if er != nil {
		return nil, er
	}
	defer res.Body.Close()

	// Continue if Response code is success
	if res.StatusCode >= 400 {
		return nil, er
	}

	// encode response body with zip type
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
	case "deflate":
		reader = flate.NewReader(res.Body)
	case "br":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer res.Body.Close()
	defer reader.Close()
	return io.ReadAll(reader)
}

func (client *Client) MakeNewRequest(method, urlStr string, data map[string]interface{}) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	var reqBody *bytes.Buffer
	if method == GET {
		params := ""
		countLength := len(data)
		idx := 0
		for key, val := range data {
			pieces := fmt.Sprintf("%s=%v", key, val)
			if idx+1 == countLength {
				params += pieces
			} else {
				params += pieces + "&"
			}
		}
		if params != "" {
			urlStr += "?" + params
		}
		return http.NewRequest(method, u.String(), nil)
	}
	params := url.Values{}
	for key, val := range data {
		params.Add(key, val.(string))
	}
	reqBody = bytes.NewBufferString(params.Encode())
	return http.NewRequest(method, u.String(), reqBody)
}
