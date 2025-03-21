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

// NewIpamPrefixesAvailableIpsCreateParams creates a new IpamPrefixesAvailableIpsCreateParams object
// with the default values initialized.
func NewIpamPrefixesAvailableIpsCreateParams() *IpamPrefixesAvailableIpsCreateParams {
	var ()
	return &IpamPrefixesAvailableIpsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithTimeout creates a new IpamPrefixesAvailableIpsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamPrefixesAvailableIpsCreateParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsCreateParams {
	var ()
	return &IpamPrefixesAvailableIpsCreateParams{

		timeout: timeout,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithContext creates a new IpamPrefixesAvailableIpsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamPrefixesAvailableIpsCreateParamsWithContext(ctx context.Context) *IpamPrefixesAvailableIpsCreateParams {
	var ()
	return &IpamPrefixesAvailableIpsCreateParams{

		Context: ctx,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithHTTPClient creates a new IpamPrefixesAvailableIpsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamPrefixesAvailableIpsCreateParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsCreateParams {
	var ()
	return &IpamPrefixesAvailableIpsCreateParams{
		HTTPClient: client,
	}
}

/*IpamPrefixesAvailableIpsCreateParams contains all the parameters to send to the API endpoint
for the ipam prefixes available ips create operation typically these are written to a http.Request
*/
type IpamPrefixesAvailableIpsCreateParams struct {

	/*Data*/
	Data *models.WritablePrefix
	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithContext(ctx context.Context) *IpamPrefixesAvailableIpsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithData(data *models.WritablePrefix) *IpamPrefixesAvailableIpsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithID(id int64) *IpamPrefixesAvailableIpsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailableIpsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
