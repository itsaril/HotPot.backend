// Package transport provides an abstraction for sending requests using different transport mechanisms.
// This implementation focuses on HTTP transport.
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// HTTPTransport is a transport implementation that uses HTTP for communication.
type HTTPTransport struct {
	address string       // Base address of the HTTP server.
	client  *http.Client // HTTP client used to send requests.
}

// NewHTTPTransport creates a new instance of HTTPTransport with a specified address.
//
// Arguments:
//
//	address - The base address of the HTTP server (e.g., "http://localhost:8080").
//
// Returns:
//
//	A pointer to an initialized HTTPTransport.
func NewHTTPTransport(address string) *HTTPTransport {
	return &HTTPTransport{
		address: address,
		client: &http.Client{
			Timeout: 30 * time.Second, // Timeout for HTTP requests.
		},
	}
}

// Send sends an HTTP request to the specified route with the given method and payload.
//
// Arguments:
//
//	ctx - Context for managing request lifecycle and deadlines.
//	request - The payload to include in the request body.
//	method - The HTTP method (e.g., "GET", "POST").
//	route - The endpoint path to which the request is sent (appended to the base address).
//
// Returns:
//   - The response body deserialized into a Go `any` type.
//   - An error if the request fails or the response cannot be processed.
func (h *HTTPTransport) Send(ctx context.Context, method string, path string, request any, contentType string) (any, error) {
	if contentType == "" {
		contentType = "application/json"
	}

	var req *http.Request
	var err error

	if method == http.MethodGet {
		// Convert request struct to query parameters
		values := url.Values{}
		v := reflect.ValueOf(request)
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i).Interface()
			values.Add(field.Tag.Get("json"), fmt.Sprintf("%v", value))
		}

		fullURL := h.address + path + "?" + values.Encode()
		req, err = http.NewRequestWithContext(ctx, method, fullURL, nil)
	} else {
		var reqBody []byte
		switch contentType {
		case "application/json":
			reqBody, err = json.Marshal(request)
		case "application/xml":
			reqBody, err = xml.Marshal(request)
		default:
			return nil, fmt.Errorf("unsupported content type: %s", contentType)
		}

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, method, h.address+path, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", contentType)
	}

	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response any
	switch contentType {
	case "application/json":
		err = json.NewDecoder(resp.Body).Decode(&response)
	case "application/xml":
		err = xml.NewDecoder(resp.Body).Decode(&response)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}
