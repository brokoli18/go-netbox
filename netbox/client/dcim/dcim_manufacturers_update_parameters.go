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

// NewDcimManufacturersUpdateParams creates a new DcimManufacturersUpdateParams object
// with the default values initialized.
func NewDcimManufacturersUpdateParams() *DcimManufacturersUpdateParams {
	var ()
	return &DcimManufacturersUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersUpdateParamsWithTimeout creates a new DcimManufacturersUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimManufacturersUpdateParamsWithTimeout(timeout time.Duration) *DcimManufacturersUpdateParams {
	var ()
	return &DcimManufacturersUpdateParams{

		timeout: timeout,
	}
}

// NewDcimManufacturersUpdateParamsWithContext creates a new DcimManufacturersUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimManufacturersUpdateParamsWithContext(ctx context.Context) *DcimManufacturersUpdateParams {
	var ()
	return &DcimManufacturersUpdateParams{

		Context: ctx,
	}
}

// NewDcimManufacturersUpdateParamsWithHTTPClient creates a new DcimManufacturersUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimManufacturersUpdateParamsWithHTTPClient(client *http.Client) *DcimManufacturersUpdateParams {
	var ()
	return &DcimManufacturersUpdateParams{
		HTTPClient: client,
	}
}

/*DcimManufacturersUpdateParams contains all the parameters to send to the API endpoint
for the dcim manufacturers update operation typically these are written to a http.Request
*/
type DcimManufacturersUpdateParams struct {

	/*Data*/
	Data *models.Manufacturer
	/*ID
	  A unique integer value identifying this manufacturer.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithTimeout(timeout time.Duration) *DcimManufacturersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithContext(ctx context.Context) *DcimManufacturersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithHTTPClient(client *http.Client) *DcimManufacturersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithData(data *models.Manufacturer) *DcimManufacturersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WithID adds the id to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) WithID(id int64) *DcimManufacturersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers update params
func (o *DcimManufacturersUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
