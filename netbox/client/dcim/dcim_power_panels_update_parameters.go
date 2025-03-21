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

// NewDcimPowerPanelsUpdateParams creates a new DcimPowerPanelsUpdateParams object
// with the default values initialized.
func NewDcimPowerPanelsUpdateParams() *DcimPowerPanelsUpdateParams {
	var ()
	return &DcimPowerPanelsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsUpdateParamsWithTimeout creates a new DcimPowerPanelsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPanelsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsUpdateParams {
	var ()
	return &DcimPowerPanelsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimPowerPanelsUpdateParamsWithContext creates a new DcimPowerPanelsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPanelsUpdateParamsWithContext(ctx context.Context) *DcimPowerPanelsUpdateParams {
	var ()
	return &DcimPowerPanelsUpdateParams{

		Context: ctx,
	}
}

// NewDcimPowerPanelsUpdateParamsWithHTTPClient creates a new DcimPowerPanelsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPanelsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsUpdateParams {
	var ()
	return &DcimPowerPanelsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimPowerPanelsUpdateParams contains all the parameters to send to the API endpoint
for the dcim power panels update operation typically these are written to a http.Request
*/
type DcimPowerPanelsUpdateParams struct {

	/*Data*/
	Data *models.WritablePowerPanel
	/*ID
	  A unique integer value identifying this power panel.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithContext(ctx context.Context) *DcimPowerPanelsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithData(data *models.WritablePowerPanel) *DcimPowerPanelsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetData(data *models.WritablePowerPanel) {
	o.Data = data
}

// WithID adds the id to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithID(id int64) *DcimPowerPanelsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
