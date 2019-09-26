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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// NewTenancyTenantsCreateParams creates a new TenancyTenantsCreateParams object
// with the default values initialized.
func NewTenancyTenantsCreateParams() *TenancyTenantsCreateParams {
	var ()
	return &TenancyTenantsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsCreateParamsWithTimeout creates a new TenancyTenantsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenancyTenantsCreateParamsWithTimeout(timeout time.Duration) *TenancyTenantsCreateParams {
	var ()
	return &TenancyTenantsCreateParams{

		timeout: timeout,
	}
}

// NewTenancyTenantsCreateParamsWithContext creates a new TenancyTenantsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenancyTenantsCreateParamsWithContext(ctx context.Context) *TenancyTenantsCreateParams {
	var ()
	return &TenancyTenantsCreateParams{

		Context: ctx,
	}
}

// NewTenancyTenantsCreateParamsWithHTTPClient creates a new TenancyTenantsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenancyTenantsCreateParamsWithHTTPClient(client *http.Client) *TenancyTenantsCreateParams {
	var ()
	return &TenancyTenantsCreateParams{
		HTTPClient: client,
	}
}

/*TenancyTenantsCreateParams contains all the parameters to send to the API endpoint
for the tenancy tenants create operation typically these are written to a http.Request
*/
type TenancyTenantsCreateParams struct {

	/*Data*/
	Data *models.TenantCreateUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithTimeout(timeout time.Duration) *TenancyTenantsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithContext(ctx context.Context) *TenancyTenantsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithHTTPClient(client *http.Client) *TenancyTenantsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithData(data *models.TenantCreateUpdate) *TenancyTenantsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetData(data *models.TenantCreateUpdate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
