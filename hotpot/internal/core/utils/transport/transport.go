// Package transport provides an abstraction for communication between services,
// supporting multiple transport types such as HTTP.
package transport

import "context"

// Transport defines the interface for sending requests and receiving responses
// using a specific transport mechanism.
// Transport defines the interface for sending requests and receiving responses
// using a specific transport mechanism.
type Transport interface {
	// Send sends a request using the given context, method, and route.
	//
	// Arguments:
	//   ctx - Context for managing request lifecycle and deadlines.
	//   method - The HTTP method or equivalent transport method (e.g., "GET", "POST").
	//   path - The endpoint or route to which the request is sent.
	//   request - The payload of the request.
	//   contentType - The content type of the request (e.g., "application/json", "application/xml").
	//
	// Returns:
	//   - The response from the transport as an `any` type.
	//   - An error if the request fails or the response cannot be processed.
	Send(ctx context.Context, method string, path string, request any, contentType string) (any, error)
}

// Type represents the type of transport to be created.
type Type string

const (
	HTTP Type = "http" // HTTP transport type.
)

// Factory is responsible for creating instances of Transport implementations.
type Factory struct{}

// New creates a new instance of Factory.
//
// Returns:
//
//	A pointer to a newly initialized Factory.
func New() *Factory {
	return &Factory{}
}

// CreateTransport creates a transport instance based on the specified transport type.
//
// Arguments:
//
//	transportType - The type of transport to create (e.g., HTTP).
//	address - The address or endpoint for the transport.
//
// Returns:
//
//	A Transport implementation corresponding to the transport type, or nil if the type is unsupported.
func (f *Factory) CreateTransport(transportType Type, address string) Transport {
	switch transportType {
	case HTTP:
		return NewHTTPTransport(address) // Assumes NewHTTPTransport is implemented elsewhere.
	default:
		return nil
	}
}
