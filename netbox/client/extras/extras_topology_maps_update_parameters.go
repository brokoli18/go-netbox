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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// NewExtrasTopologyMapsUpdateParams creates a new ExtrasTopologyMapsUpdateParams object
// with the default values initialized.
func NewExtrasTopologyMapsUpdateParams() *ExtrasTopologyMapsUpdateParams {
	var ()
	return &ExtrasTopologyMapsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsUpdateParamsWithTimeout creates a new ExtrasTopologyMapsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsUpdateParams {
	var ()
	return &ExtrasTopologyMapsUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsUpdateParamsWithContext creates a new ExtrasTopologyMapsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsUpdateParamsWithContext(ctx context.Context) *ExtrasTopologyMapsUpdateParams {
	var ()
	return &ExtrasTopologyMapsUpdateParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsUpdateParamsWithHTTPClient creates a new ExtrasTopologyMapsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsUpdateParams {
	var ()
	return &ExtrasTopologyMapsUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsUpdateParams contains all the parameters to send to the API endpoint
for the extras topology maps update operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsUpdateParams struct {

	/*Data*/
	Data *models.TopologyMap
	/*ID
	  A unique integer value identifying this topology map.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) WithContext(ctx context.Context) *ExtrasTopologyMapsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) WithData(data *models.TopologyMap) *ExtrasTopologyMapsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) SetData(data *models.TopologyMap) {
	o.Data = data
}

// WithID adds the id to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) WithID(id int64) *ExtrasTopologyMapsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras topology maps update params
func (o *ExtrasTopologyMapsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
