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

// NewVirtualizationClusterGroupsUpdateParams creates a new VirtualizationClusterGroupsUpdateParams object
// with the default values initialized.
func NewVirtualizationClusterGroupsUpdateParams() *VirtualizationClusterGroupsUpdateParams {
	var ()
	return &VirtualizationClusterGroupsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithTimeout creates a new VirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterGroupsUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsUpdateParams {
	var ()
	return &VirtualizationClusterGroupsUpdateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithContext creates a new VirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterGroupsUpdateParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsUpdateParams {
	var ()
	return &VirtualizationClusterGroupsUpdateParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithHTTPClient creates a new VirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterGroupsUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsUpdateParams {
	var ()
	return &VirtualizationClusterGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterGroupsUpdateParams contains all the parameters to send to the API endpoint
for the virtualization cluster groups update operation typically these are written to a http.Request
*/
type VirtualizationClusterGroupsUpdateParams struct {

	/*Data*/
	Data *models.ClusterGroup
	/*ID
	  A unique integer value identifying this cluster group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithData(data *models.ClusterGroup) *VirtualizationClusterGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WithID adds the id to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithID(id int64) *VirtualizationClusterGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
