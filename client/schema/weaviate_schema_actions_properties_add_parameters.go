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
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateSchemaActionsPropertiesAddParams creates a new WeaviateSchemaActionsPropertiesAddParams object
// with the default values initialized.
func NewWeaviateSchemaActionsPropertiesAddParams() *WeaviateSchemaActionsPropertiesAddParams {
	var ()
	return &WeaviateSchemaActionsPropertiesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaActionsPropertiesAddParamsWithTimeout creates a new WeaviateSchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaActionsPropertiesAddParamsWithTimeout(timeout time.Duration) *WeaviateSchemaActionsPropertiesAddParams {
	var ()
	return &WeaviateSchemaActionsPropertiesAddParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaActionsPropertiesAddParamsWithContext creates a new WeaviateSchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaActionsPropertiesAddParamsWithContext(ctx context.Context) *WeaviateSchemaActionsPropertiesAddParams {
	var ()
	return &WeaviateSchemaActionsPropertiesAddParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaActionsPropertiesAddParamsWithHTTPClient creates a new WeaviateSchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaActionsPropertiesAddParamsWithHTTPClient(client *http.Client) *WeaviateSchemaActionsPropertiesAddParams {
	var ()
	return &WeaviateSchemaActionsPropertiesAddParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaActionsPropertiesAddParams contains all the parameters to send to the API endpoint
for the weaviate schema actions properties add operation typically these are written to a http.Request
*/
type WeaviateSchemaActionsPropertiesAddParams struct {

	/*Body*/
	Body *models.SemanticSchemaClassProperty
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) WithTimeout(timeout time.Duration) *WeaviateSchemaActionsPropertiesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) WithContext(ctx context.Context) *WeaviateSchemaActionsPropertiesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) WithHTTPClient(client *http.Client) *WeaviateSchemaActionsPropertiesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) WithBody(body *models.SemanticSchemaClassProperty) *WeaviateSchemaActionsPropertiesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) SetBody(body *models.SemanticSchemaClassProperty) {
	o.Body = body
}

// WithClassName adds the className to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) WithClassName(className string) *WeaviateSchemaActionsPropertiesAddParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema actions properties add params
func (o *WeaviateSchemaActionsPropertiesAddParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaActionsPropertiesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
