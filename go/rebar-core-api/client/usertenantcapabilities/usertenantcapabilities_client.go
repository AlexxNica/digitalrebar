package usertenantcapabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new usertenantcapabilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for usertenantcapabilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEUsertenantcapability deletes user tenant capability
*/
func (a *Client) DELETEUsertenantcapability(params *DELETEUsertenantcapabilityParams) (*DELETEUsertenantcapabilityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEUsertenantcapabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-usertenantcapability",
		Method:             "DELETE",
		PathPattern:        "/usertenantcapabilities/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEUsertenantcapabilityReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEUsertenantcapabilityNoContent), nil

}

/*
GETUsertenantcapability gets user tenant capability
*/
func (a *Client) GETUsertenantcapability(params *GETUsertenantcapabilityParams) (*GETUsertenantcapabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETUsertenantcapabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-usertenantcapability",
		Method:             "GET",
		PathPattern:        "/usertenantcapabilities/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETUsertenantcapabilityReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETUsertenantcapabilityOK), nil

}

/*
LISTUsertenantcapabilities lists usertenantcapabilities
*/
func (a *Client) LISTUsertenantcapabilities(params *LISTUsertenantcapabilitiesParams) (*LISTUsertenantcapabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTUsertenantcapabilitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-usertenantcapabilities",
		Method:             "GET",
		PathPattern:        "/usertenantcapabilities",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTUsertenantcapabilitiesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTUsertenantcapabilitiesOK), nil

}

/*
POSTUsertenantcapability creates user tenant capability
*/
func (a *Client) POSTUsertenantcapability(params *POSTUsertenantcapabilityParams) (*POSTUsertenantcapabilityCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTUsertenantcapabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-usertenantcapability",
		Method:             "POST",
		PathPattern:        "/usertenantcapabilities",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTUsertenantcapabilityReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTUsertenantcapabilityCreated), nil

}

/*
PUTUsertenantcapability updates user tenant capability
*/
func (a *Client) PUTUsertenantcapability(params *PUTUsertenantcapabilityParams) (*PUTUsertenantcapabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTUsertenantcapabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-usertenantcapability",
		Method:             "PUT",
		PathPattern:        "/usertenantcapabilities/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTUsertenantcapabilityReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTUsertenantcapabilityOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}