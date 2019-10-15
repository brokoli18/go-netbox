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

package dcim

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

// NewDcimConsoleServerPortsUpdateParams creates a new DcimConsoleServerPortsUpdateParams object
// with the default values initialized.
func NewDcimConsoleServerPortsUpdateParams() *DcimConsoleServerPortsUpdateParams {
	var ()
	return &DcimConsoleServerPortsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsUpdateParamsWithTimeout creates a new DcimConsoleServerPortsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortsUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsUpdateParams {
	var ()
	return &DcimConsoleServerPortsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsUpdateParamsWithContext creates a new DcimConsoleServerPortsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortsUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortsUpdateParams {
	var ()
	return &DcimConsoleServerPortsUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortsUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortsUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsUpdateParams {
	var ()
	return &DcimConsoleServerPortsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortsUpdateParams contains all the parameters to send to the API endpoint
for the dcim console server ports update operation typically these are written to a http.Request
*/
type DcimConsoleServerPortsUpdateParams struct {

	/*Data*/
	Data *models.WritableConsoleServerPort
	/*ID
	  A unique integer value identifying this console server port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) WithData(data *models.WritableConsoleServerPort) *DcimConsoleServerPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) SetData(data *models.WritableConsoleServerPort) {
	o.Data = data
}

// WithID adds the id to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) WithID(id int64) *DcimConsoleServerPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server ports update params
func (o *DcimConsoleServerPortsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
