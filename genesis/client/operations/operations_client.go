/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GenesisPeersLeave Leave the weaviate network
*/
func (a *Client) GenesisPeersLeave(params *GenesisPeersLeaveParams) (*GenesisPeersLeaveNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenesisPeersLeaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "genesis.peers.leave",
		Method:             "DELETE",
		PathPattern:        "/peers/{peerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenesisPeersLeaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenesisPeersLeaveNoContent), nil

}

/*
GenesisPeersList List the registered peers
*/
func (a *Client) GenesisPeersList(params *GenesisPeersListParams) (*GenesisPeersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenesisPeersListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "genesis.peers.list",
		Method:             "GET",
		PathPattern:        "/peers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenesisPeersListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenesisPeersListOK), nil

}

/*
GenesisPeersPing Ping the Genesis server, to make mark the peer as alive and udpate schema info
*/
func (a *Client) GenesisPeersPing(params *GenesisPeersPingParams) (*GenesisPeersPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenesisPeersPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "genesis.peers.ping",
		Method:             "POST",
		PathPattern:        "/peers/{peerId}/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenesisPeersPingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenesisPeersPingOK), nil

}

/*
GenesisPeersRegister Register a new Weaviate peer in the network
*/
func (a *Client) GenesisPeersRegister(params *GenesisPeersRegisterParams) (*GenesisPeersRegisterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenesisPeersRegisterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "genesis.peers.register",
		Method:             "POST",
		PathPattern:        "/peers/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenesisPeersRegisterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenesisPeersRegisterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
