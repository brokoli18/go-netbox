// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// NewExtrasTagsUpdateParams creates a new ExtrasTagsUpdateParams object
// with the default values initialized.
func NewExtrasTagsUpdateParams() *ExtrasTagsUpdateParams {
	var ()
	return &ExtrasTagsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsUpdateParamsWithTimeout creates a new ExtrasTagsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTagsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsUpdateParams {
	var ()
	return &ExtrasTagsUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasTagsUpdateParamsWithContext creates a new ExtrasTagsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTagsUpdateParamsWithContext(ctx context.Context) *ExtrasTagsUpdateParams {
	var ()
	return &ExtrasTagsUpdateParams{

		Context: ctx,
	}
}

// NewExtrasTagsUpdateParamsWithHTTPClient creates a new ExtrasTagsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTagsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsUpdateParams {
	var ()
	return &ExtrasTagsUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasTagsUpdateParams contains all the parameters to send to the API endpoint
for the extras tags update operation typically these are written to a http.Request
*/
type ExtrasTagsUpdateParams struct {

	/*Data*/
	Data *models.Tag
	/*ID
	  A unique integer value identifying this tag.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithContext(ctx context.Context) *ExtrasTagsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithData(data *models.Tag) *ExtrasTagsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WithID adds the id to the extras tags update params
func (o *ExtrasTagsUpdateParams) WithID(id int64) *ExtrasTagsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags update params
func (o *ExtrasTagsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
