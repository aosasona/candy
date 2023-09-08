package request

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/aosasona/fate/pkg/candy/internal/client"
	"github.com/goccy/go-json"
)

type Request struct {
	// Path to the endpoint without the URL or version, eg: /config
	Path string

	// Headers to be sent with the request
	Headers map[string]string

	// Additional data to be sent with the request
	Data any
}

type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
)

var CaddyAPIAddr = "http://0.0.0.0:2019"

func sendRequest(method RequestMethod, args Request) (*http.Response, error) {
	args.Path = strings.Trim(args.Path, "/")
	url := fmt.Sprintf("%s/%s", CaddyAPIAddr, args.Path)

	if args.Headers == nil {
		args.Headers = make(map[string]string)
	}

	data := &bytes.Buffer{}

	if args.Data != nil {
		body, err := json.Marshal(args.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}

		data = bytes.NewBuffer(body)
	}

	printDebug(debugOpts{URL: url, Method: method, Headers: args.Headers, Body: string(data.Bytes())})

	req, err := http.NewRequest(string(method), url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if method == POST {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	for k, v := range args.Headers {
		req.Header.Set(k, v)
	}

	return client.Client.HttpClient.Do(req)
}

func Get[T any](args Request, out *T) error {
	resp, err := sendRequest(GET, args)
	if err != nil {
		return err
	}
	return decodeResponse(resp, out)
}

func Post[T any](args Request, out *T) error {
	resp, err := sendRequest(POST, args)
	if err != nil {
		return err
	}
	return decodeResponse(resp, out)
}

func decodeResponse[T any](res *http.Response, out *T) error {
	defer res.Body.Close()

	if res.ContentLength == 0 {
		return nil
	}

	err := json.NewDecoder(res.Body).Decode(out)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if res.StatusCode >= 400 || res.StatusCode < 200 {
		raw, _ := json.Marshal(out)
		e := map[string]interface{}{}
		json.Unmarshal(raw, &e)
		serverErr, ok := e["error"]
		if !ok {
			serverErr = e
		}
		return fmt.Errorf("Status ==> %d\nERROR ==> %+v", res.StatusCode, serverErr)
	}

	return nil
}
