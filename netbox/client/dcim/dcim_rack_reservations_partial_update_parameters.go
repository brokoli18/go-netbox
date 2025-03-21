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

// NewDcimRackReservationsPartialUpdateParams creates a new DcimRackReservationsPartialUpdateParams object
// with the default values initialized.
func NewDcimRackReservationsPartialUpdateParams() *DcimRackReservationsPartialUpdateParams {
	var ()
	return &DcimRackReservationsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsPartialUpdateParamsWithTimeout creates a new DcimRackReservationsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackReservationsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackReservationsPartialUpdateParams {
	var ()
	return &DcimRackReservationsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimRackReservationsPartialUpdateParamsWithContext creates a new DcimRackReservationsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackReservationsPartialUpdateParamsWithContext(ctx context.Context) *DcimRackReservationsPartialUpdateParams {
	var ()
	return &DcimRackReservationsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimRackReservationsPartialUpdateParamsWithHTTPClient creates a new DcimRackReservationsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackReservationsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackReservationsPartialUpdateParams {
	var ()
	return &DcimRackReservationsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimRackReservationsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim rack reservations partial update operation typically these are written to a http.Request
*/
type DcimRackReservationsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableRackReservation
	/*ID
	  A unique integer value identifying this rack reservation.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackReservationsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) WithContext(ctx context.Context) *DcimRackReservationsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackReservationsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) WithData(data *models.WritableRackReservation) *DcimRackReservationsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) SetData(data *models.WritableRackReservation) {
	o.Data = data
}

// WithID adds the id to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) WithID(id int64) *DcimRackReservationsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack reservations partial update params
func (o *DcimRackReservationsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
