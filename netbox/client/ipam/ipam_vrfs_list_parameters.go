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
)

// NewIpamVrfsListParams creates a new IpamVrfsListParams object
// with the default values initialized.
func NewIpamVrfsListParams() *IpamVrfsListParams {
	var ()
	return &IpamVrfsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsListParamsWithTimeout creates a new IpamVrfsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamVrfsListParamsWithTimeout(timeout time.Duration) *IpamVrfsListParams {
	var ()
	return &IpamVrfsListParams{

		timeout: timeout,
	}
}

// NewIpamVrfsListParamsWithContext creates a new IpamVrfsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamVrfsListParamsWithContext(ctx context.Context) *IpamVrfsListParams {
	var ()
	return &IpamVrfsListParams{

		Context: ctx,
	}
}

// NewIpamVrfsListParamsWithHTTPClient creates a new IpamVrfsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamVrfsListParamsWithHTTPClient(client *http.Client) *IpamVrfsListParams {
	var ()
	return &IpamVrfsListParams{
		HTTPClient: client,
	}
}

/*IpamVrfsListParams contains all the parameters to send to the API endpoint
for the ipam vrfs list operation typically these are written to a http.Request
*/
type IpamVrfsListParams struct {

	/*EnforceUnique*/
	EnforceUnique *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Rd*/
	Rd *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *string
	/*TenantID*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTimeout(timeout time.Duration) *IpamVrfsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs list params
func (o *IpamVrfsListParams) WithContext(ctx context.Context) *IpamVrfsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs list params
func (o *IpamVrfsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IpamVrfsListParams) WithHTTPClient(client *http.Client) *IpamVrfsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IpamVrfsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IpamVrfsListParams) WithEnforceUnique(enforceUnique *string) *IpamVrfsListParams {
	o.SetEnforceUnique(enforceUnique)
	return o
}

// SetEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IpamVrfsListParams) SetEnforceUnique(enforceUnique *string) {
	o.EnforceUnique = enforceUnique
}

// WithIDIn adds the iDIn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDIn(iDIn *string) *IpamVrfsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the ipam vrfs list params
func (o *IpamVrfsListParams) WithLimit(limit *int64) *IpamVrfsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vrfs list params
func (o *IpamVrfsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vrfs list params
func (o *IpamVrfsListParams) WithName(name *string) *IpamVrfsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vrfs list params
func (o *IpamVrfsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam vrfs list params
func (o *IpamVrfsListParams) WithOffset(offset *int64) *IpamVrfsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vrfs list params
func (o *IpamVrfsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vrfs list params
func (o *IpamVrfsListParams) WithQ(q *string) *IpamVrfsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vrfs list params
func (o *IpamVrfsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRd adds the rd to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRd(rd *string) *IpamVrfsListParams {
	o.SetRd(rd)
	return o
}

// SetRd adds the rd to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRd(rd *string) {
	o.Rd = rd
}

// WithTag adds the tag to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTag(tag *string) *IpamVrfsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenant(tenant *string) *IpamVrfsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroup(tenantGroup *string) *IpamVrfsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroupID(tenantGroupID *string) *IpamVrfsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantID(tenantID *string) *IpamVrfsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnforceUnique != nil {

		// query param enforce_unique
		var qrEnforceUnique string
		if o.EnforceUnique != nil {
			qrEnforceUnique = *o.EnforceUnique
		}
		qEnforceUnique := qrEnforceUnique
		if qEnforceUnique != "" {
			if err := r.SetQueryParam("enforce_unique", qEnforceUnique); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Rd != nil {

		// query param rd
		var qrRd string
		if o.Rd != nil {
			qrRd = *o.Rd
		}
		qRd := qrRd
		if qRd != "" {
			if err := r.SetQueryParam("rd", qRd); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string
		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {
			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string
		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {
			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
