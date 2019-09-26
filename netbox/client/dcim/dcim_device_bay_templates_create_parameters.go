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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// NewDcimDeviceBayTemplatesCreateParams creates a new DcimDeviceBayTemplatesCreateParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesCreateParams() *DcimDeviceBayTemplatesCreateParams {
	var ()
	return &DcimDeviceBayTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesCreateParamsWithTimeout creates a new DcimDeviceBayTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesCreateParams {
	var ()
	return &DcimDeviceBayTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesCreateParamsWithContext creates a new DcimDeviceBayTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesCreateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesCreateParams {
	var ()
	return &DcimDeviceBayTemplatesCreateParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesCreateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesCreateParams {
	var ()
	return &DcimDeviceBayTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesCreateParams contains all the parameters to send to the API endpoint
for the dcim device bay templates create operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesCreateParams struct {

	/*Data*/
	Data *models.DeviceBayTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) WithData(data *models.DeviceBayTemplate) *DcimDeviceBayTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates create params
func (o *DcimDeviceBayTemplatesCreateParams) SetData(data *models.DeviceBayTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
