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

package virtualization

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

// NewVirtualizationVirtualMachinesUpdateParams creates a new VirtualizationVirtualMachinesUpdateParams object
// with the default values initialized.
func NewVirtualizationVirtualMachinesUpdateParams() *VirtualizationVirtualMachinesUpdateParams {
	var ()
	return &VirtualizationVirtualMachinesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithTimeout creates a new VirtualizationVirtualMachinesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationVirtualMachinesUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesUpdateParams {
	var ()
	return &VirtualizationVirtualMachinesUpdateParams{

		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithContext creates a new VirtualizationVirtualMachinesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationVirtualMachinesUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesUpdateParams {
	var ()
	return &VirtualizationVirtualMachinesUpdateParams{

		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithHTTPClient creates a new VirtualizationVirtualMachinesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationVirtualMachinesUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesUpdateParams {
	var ()
	return &VirtualizationVirtualMachinesUpdateParams{
		HTTPClient: client,
	}
}

/*VirtualizationVirtualMachinesUpdateParams contains all the parameters to send to the API endpoint
for the virtualization virtual machines update operation typically these are written to a http.Request
*/
type VirtualizationVirtualMachinesUpdateParams struct {

	/*Data*/
	Data *models.WritableVirtualMachineWithConfigContext
	/*ID
	  A unique integer value identifying this virtual machine.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithData(data *models.WritableVirtualMachineWithConfigContext) *VirtualizationVirtualMachinesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetData(data *models.WritableVirtualMachineWithConfigContext) {
	o.Data = data
}

// WithID adds the id to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithID(id int64) *VirtualizationVirtualMachinesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
