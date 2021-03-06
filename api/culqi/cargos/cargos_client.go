package cargos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cargos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cargos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
POSTCargos crears cargo
*/
func (a *Client) POSTCargos(params *POSTCargosParams, authInfo runtime.ClientAuthInfoWriter) (*POSTCargosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTCargosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_cargos",
		Method:             "POST",
		PathPattern:        "/cargos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTCargosReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTCargosOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
