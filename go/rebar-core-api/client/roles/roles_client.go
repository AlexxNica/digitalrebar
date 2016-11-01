package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETERoles deletes roles
*/
func (a *Client) DELETERoles(params *DELETERolesParams) (*DELETERolesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETERolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-roles",
		Method:             "DELETE",
		PathPattern:        "/roles/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETERolesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETERolesNoContent), nil

}

/*
GETRoles gets roles
*/
func (a *Client) GETRoles(params *GETRolesParams) (*GETRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-roles",
		Method:             "GET",
		PathPattern:        "/roles/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETRolesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETRolesOK), nil

}

/*
LISTRoles lists roles
*/
func (a *Client) LISTRoles(params *LISTRolesParams) (*LISTRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-roles",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTRolesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTRolesOK), nil

}

/*
POSTRoles creates roles
*/
func (a *Client) POSTRoles(params *POSTRolesParams) (*POSTRolesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-roles",
		Method:             "POST",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTRolesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTRolesCreated), nil

}

/*
PUTRoles updates roles
*/
func (a *Client) PUTRoles(params *PUTRolesParams) (*PUTRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-roles",
		Method:             "PUT",
		PathPattern:        "/roles/{id}",
		ProducesMediaTypes: []string{"application/javascript", "application/json"},
		ConsumesMediaTypes: []string{"application/javascript", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTRolesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTRolesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}