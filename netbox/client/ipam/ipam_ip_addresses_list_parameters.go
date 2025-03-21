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

// NewIpamIPAddressesListParams creates a new IpamIPAddressesListParams object
// with the default values initialized.
func NewIpamIPAddressesListParams() *IpamIPAddressesListParams {
	var ()
	return &IpamIPAddressesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesListParamsWithTimeout creates a new IpamIPAddressesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamIPAddressesListParamsWithTimeout(timeout time.Duration) *IpamIPAddressesListParams {
	var ()
	return &IpamIPAddressesListParams{

		timeout: timeout,
	}
}

// NewIpamIPAddressesListParamsWithContext creates a new IpamIPAddressesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamIPAddressesListParamsWithContext(ctx context.Context) *IpamIPAddressesListParams {
	var ()
	return &IpamIPAddressesListParams{

		Context: ctx,
	}
}

// NewIpamIPAddressesListParamsWithHTTPClient creates a new IpamIPAddressesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamIPAddressesListParamsWithHTTPClient(client *http.Client) *IpamIPAddressesListParams {
	var ()
	return &IpamIPAddressesListParams{
		HTTPClient: client,
	}
}

/*IpamIPAddressesListParams contains all the parameters to send to the API endpoint
for the ipam ip addresses list operation typically these are written to a http.Request
*/
type IpamIPAddressesListParams struct {

	/*Address*/
	Address *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *float64
	/*DNSName*/
	DNSName *string
	/*Family*/
	Family *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Interface*/
	Interface *string
	/*InterfaceID*/
	InterfaceID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MaskLength*/
	MaskLength *float64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Parent*/
	Parent *string
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*Status*/
	Status *string
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
	/*VirtualMachine*/
	VirtualMachine *string
	/*VirtualMachineID*/
	VirtualMachineID *string
	/*Vrf*/
	Vrf *string
	/*VrfID*/
	VrfID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTimeout(timeout time.Duration) *IpamIPAddressesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithContext(ctx context.Context) *IpamIPAddressesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithHTTPClient(client *http.Client) *IpamIPAddressesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithAddress(address *string) *IpamIPAddressesListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetAddress(address *string) {
	o.Address = address
}

// WithDevice adds the device to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDevice(device *string) *IpamIPAddressesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDeviceID(deviceID *float64) *IpamIPAddressesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDeviceID(deviceID *float64) {
	o.DeviceID = deviceID
}

// WithDNSName adds the dNSName to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSName(dNSName *string) *IpamIPAddressesListParams {
	o.SetDNSName(dNSName)
	return o
}

// SetDNSName adds the dnsName to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSName(dNSName *string) {
	o.DNSName = dNSName
}

// WithFamily adds the family to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithFamily(family *string) *IpamIPAddressesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetFamily(family *string) {
	o.Family = family
}

// WithIDIn adds the iDIn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDIn(iDIn *string) *IpamIPAddressesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithInterface adds the interfaceVar to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterface(interfaceVar *string) *IpamIPAddressesListParams {
	o.SetInterface(interfaceVar)
	return o
}

// SetInterface adds the interface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterface(interfaceVar *string) {
	o.Interface = interfaceVar
}

// WithInterfaceID adds the interfaceID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterfaceID(interfaceID *string) *IpamIPAddressesListParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterfaceID(interfaceID *string) {
	o.InterfaceID = interfaceID
}

// WithLimit adds the limit to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithLimit(limit *int64) *IpamIPAddressesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithMaskLength(maskLength *float64) *IpamIPAddressesListParams {
	o.SetMaskLength(maskLength)
	return o
}

// SetMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetMaskLength(maskLength *float64) {
	o.MaskLength = maskLength
}

// WithOffset adds the offset to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithOffset(offset *int64) *IpamIPAddressesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParent adds the parent to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithParent(parent *string) *IpamIPAddressesListParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithQ adds the q to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithQ(q *string) *IpamIPAddressesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithRole(role *string) *IpamIPAddressesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetRole(role *string) {
	o.Role = role
}

// WithStatus adds the status to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithStatus(status *string) *IpamIPAddressesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTag(tag *string) *IpamIPAddressesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenant(tenant *string) *IpamIPAddressesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroup(tenantGroup *string) *IpamIPAddressesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroupID(tenantGroupID *string) *IpamIPAddressesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantID(tenantID *string) *IpamIPAddressesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVirtualMachine(virtualMachine *string) *IpamIPAddressesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVirtualMachineID(virtualMachineID *string) *IpamIPAddressesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVrf adds the vrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrf(vrf *string) *IpamIPAddressesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfID adds the vrfID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrfID(vrfID *string) *IpamIPAddressesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrfID(vrfID *string) {
	o.VrfID = vrfID
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID float64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatFloat64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.DNSName != nil {

		// query param dns_name
		var qrDNSName string
		if o.DNSName != nil {
			qrDNSName = *o.DNSName
		}
		qDNSName := qrDNSName
		if qDNSName != "" {
			if err := r.SetQueryParam("dns_name", qDNSName); err != nil {
				return err
			}
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
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

	if o.Interface != nil {

		// query param interface
		var qrInterface string
		if o.Interface != nil {
			qrInterface = *o.Interface
		}
		qInterface := qrInterface
		if qInterface != "" {
			if err := r.SetQueryParam("interface", qInterface); err != nil {
				return err
			}
		}

	}

	if o.InterfaceID != nil {

		// query param interface_id
		var qrInterfaceID string
		if o.InterfaceID != nil {
			qrInterfaceID = *o.InterfaceID
		}
		qInterfaceID := qrInterfaceID
		if qInterfaceID != "" {
			if err := r.SetQueryParam("interface_id", qInterfaceID); err != nil {
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

	if o.MaskLength != nil {

		// query param mask_length
		var qrMaskLength float64
		if o.MaskLength != nil {
			qrMaskLength = *o.MaskLength
		}
		qMaskLength := swag.FormatFloat64(qrMaskLength)
		if qMaskLength != "" {
			if err := r.SetQueryParam("mask_length", qMaskLength); err != nil {
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

	if o.Parent != nil {

		// query param parent
		var qrParent string
		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {
			if err := r.SetQueryParam("parent", qParent); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
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

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string
		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {
			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}

	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string
		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {
			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}

	}

	if o.Vrf != nil {

		// query param vrf
		var qrVrf string
		if o.Vrf != nil {
			qrVrf = *o.Vrf
		}
		qVrf := qrVrf
		if qVrf != "" {
			if err := r.SetQueryParam("vrf", qVrf); err != nil {
				return err
			}
		}

	}

	if o.VrfID != nil {

		// query param vrf_id
		var qrVrfID string
		if o.VrfID != nil {
			qrVrfID = *o.VrfID
		}
		qVrfID := qrVrfID
		if qVrfID != "" {
			if err := r.SetQueryParam("vrf_id", qVrfID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
