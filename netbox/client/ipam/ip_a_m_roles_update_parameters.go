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

package ipam

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

// NewIPAMRolesUpdateParams creates a new IPAMRolesUpdateParams object
// with the default values initialized.
func NewIPAMRolesUpdateParams() *IPAMRolesUpdateParams {
	var ()
	return &IPAMRolesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMRolesUpdateParamsWithTimeout creates a new IPAMRolesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMRolesUpdateParamsWithTimeout(timeout time.Duration) *IPAMRolesUpdateParams {
	var ()
	return &IPAMRolesUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMRolesUpdateParamsWithContext creates a new IPAMRolesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMRolesUpdateParamsWithContext(ctx context.Context) *IPAMRolesUpdateParams {
	var ()
	return &IPAMRolesUpdateParams{

		Context: ctx,
	}
}

// NewIPAMRolesUpdateParamsWithHTTPClient creates a new IPAMRolesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMRolesUpdateParamsWithHTTPClient(client *http.Client) *IPAMRolesUpdateParams {
	var ()
	return &IPAMRolesUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMRolesUpdateParams contains all the parameters to send to the API endpoint
for the ipam roles update operation typically these are written to a http.Request
*/
type IPAMRolesUpdateParams struct {

	/*Data*/
	Data *models.Role
	/*ID
	  A unique integer value identifying this role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam roles update params
func (o *IPAMRolesUpdateParams) WithTimeout(timeout time.Duration) *IPAMRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles update params
func (o *IPAMRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles update params
func (o *IPAMRolesUpdateParams) WithContext(ctx context.Context) *IPAMRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles update params
func (o *IPAMRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles update params
func (o *IPAMRolesUpdateParams) WithHTTPClient(client *http.Client) *IPAMRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles update params
func (o *IPAMRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam roles update params
func (o *IPAMRolesUpdateParams) WithData(data *models.Role) *IPAMRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam roles update params
func (o *IPAMRolesUpdateParams) SetData(data *models.Role) {
	o.Data = data
}

// WithID adds the id to the ipam roles update params
func (o *IPAMRolesUpdateParams) WithID(id int64) *IPAMRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles update params
func (o *IPAMRolesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
