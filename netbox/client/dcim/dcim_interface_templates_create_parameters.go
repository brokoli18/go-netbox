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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// NewDcimInterfaceTemplatesCreateParams creates a new DcimInterfaceTemplatesCreateParams object
// with the default values initialized.
func NewDcimInterfaceTemplatesCreateParams() *DcimInterfaceTemplatesCreateParams {
	var ()
	return &DcimInterfaceTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesCreateParamsWithTimeout creates a new DcimInterfaceTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfaceTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesCreateParams {
	var ()
	return &DcimInterfaceTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesCreateParamsWithContext creates a new DcimInterfaceTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfaceTemplatesCreateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesCreateParams {
	var ()
	return &DcimInterfaceTemplatesCreateParams{

		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesCreateParamsWithHTTPClient creates a new DcimInterfaceTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfaceTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesCreateParams {
	var ()
	return &DcimInterfaceTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*DcimInterfaceTemplatesCreateParams contains all the parameters to send to the API endpoint
for the dcim interface templates create operation typically these are written to a http.Request
*/
type DcimInterfaceTemplatesCreateParams struct {

	/*Data*/
	Data *models.WritableInterfaceTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates create params
func (o *DcimInterfaceTemplatesCreateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
