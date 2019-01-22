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
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GenesisPeersPingReader is a Reader for the GenesisPeersPing structure.
type GenesisPeersPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenesisPeersPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenesisPeersPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGenesisPeersPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGenesisPeersPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGenesisPeersPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenesisPeersPingOK creates a GenesisPeersPingOK with default headers values
func NewGenesisPeersPingOK() *GenesisPeersPingOK {
	return &GenesisPeersPingOK{}
}

/*GenesisPeersPingOK handles this case with default header values.

Ping received
*/
type GenesisPeersPingOK struct {
}

func (o *GenesisPeersPingOK) Error() string {
	return fmt.Sprintf("[POST /peers/{peerId}/ping][%d] genesisPeersPingOK ", 200)
}

func (o *GenesisPeersPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenesisPeersPingUnauthorized creates a GenesisPeersPingUnauthorized with default headers values
func NewGenesisPeersPingUnauthorized() *GenesisPeersPingUnauthorized {
	return &GenesisPeersPingUnauthorized{}
}

/*GenesisPeersPingUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type GenesisPeersPingUnauthorized struct {
}

func (o *GenesisPeersPingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /peers/{peerId}/ping][%d] genesisPeersPingUnauthorized ", 401)
}

func (o *GenesisPeersPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenesisPeersPingForbidden creates a GenesisPeersPingForbidden with default headers values
func NewGenesisPeersPingForbidden() *GenesisPeersPingForbidden {
	return &GenesisPeersPingForbidden{}
}

/*GenesisPeersPingForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type GenesisPeersPingForbidden struct {
}

func (o *GenesisPeersPingForbidden) Error() string {
	return fmt.Sprintf("[POST /peers/{peerId}/ping][%d] genesisPeersPingForbidden ", 403)
}

func (o *GenesisPeersPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenesisPeersPingNotFound creates a GenesisPeersPingNotFound with default headers values
func NewGenesisPeersPingNotFound() *GenesisPeersPingNotFound {
	return &GenesisPeersPingNotFound{}
}

/*GenesisPeersPingNotFound handles this case with default header values.

Successful query result but no such peer was found.
*/
type GenesisPeersPingNotFound struct {
}

func (o *GenesisPeersPingNotFound) Error() string {
	return fmt.Sprintf("[POST /peers/{peerId}/ping][%d] genesisPeersPingNotFound ", 404)
}

func (o *GenesisPeersPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
