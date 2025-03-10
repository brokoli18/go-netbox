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
)

// NewVirtualizationVirtualMachinesListParams creates a new VirtualizationVirtualMachinesListParams object
// with the default values initialized.
func NewVirtualizationVirtualMachinesListParams() *VirtualizationVirtualMachinesListParams {
	var ()
	return &VirtualizationVirtualMachinesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithTimeout creates a new VirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationVirtualMachinesListParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesListParams {
	var ()
	return &VirtualizationVirtualMachinesListParams{

		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithContext creates a new VirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationVirtualMachinesListParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesListParams {
	var ()
	return &VirtualizationVirtualMachinesListParams{

		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithHTTPClient creates a new VirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationVirtualMachinesListParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesListParams {
	var ()
	return &VirtualizationVirtualMachinesListParams{
		HTTPClient: client,
	}
}

/*VirtualizationVirtualMachinesListParams contains all the parameters to send to the API endpoint
for the virtualization virtual machines list operation typically these are written to a http.Request
*/
type VirtualizationVirtualMachinesListParams struct {

	/*Cluster*/
	Cluster *string
	/*ClusterGroup*/
	ClusterGroup *string
	/*ClusterGroupID*/
	ClusterGroupID *string
	/*ClusterID*/
	ClusterID *string
	/*ClusterType*/
	ClusterType *string
	/*ClusterTypeID*/
	ClusterTypeID *string
	/*Disk*/
	Disk *string
	/*ID*/
	ID *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MacAddress*/
	MacAddress *string
	/*Memory*/
	Memory *string
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Platform*/
	Platform *string
	/*PlatformID*/
	PlatformID *string
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
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
	/*Vcpus*/
	Vcpus *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithCluster(cluster *string) *VirtualizationVirtualMachinesListParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithClusterGroup adds the clusterGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroup(clusterGroup *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroup(clusterGroup)
	return o
}

// SetClusterGroup adds the clusterGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroup(clusterGroup *string) {
	o.ClusterGroup = clusterGroup
}

// WithClusterGroupID adds the clusterGroupID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroupID(clusterGroupID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroupID(clusterGroupID)
	return o
}

// SetClusterGroupID adds the clusterGroupId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroupID(clusterGroupID *string) {
	o.ClusterGroupID = clusterGroupID
}

// WithClusterID adds the clusterID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterID(clusterID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterType adds the clusterType to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterType(clusterType *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterType(clusterType *string) {
	o.ClusterType = clusterType
}

// WithClusterTypeID adds the clusterTypeID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterTypeID(clusterTypeID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterTypeID(clusterTypeID)
	return o
}

// SetClusterTypeID adds the clusterTypeId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterTypeID(clusterTypeID *string) {
	o.ClusterTypeID = clusterTypeID
}

// WithDisk adds the disk to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDisk(disk *string) *VirtualizationVirtualMachinesListParams {
	o.SetDisk(disk)
	return o
}

// SetDisk adds the disk to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDisk(disk *string) {
	o.Disk = disk
}

// WithID adds the id to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithID(id *string) *VirtualizationVirtualMachinesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIn adds the iDIn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDIn(iDIn *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLimit(limit *int64) *VirtualizationVirtualMachinesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddress(macAddress *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMemory adds the memory to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemory(memory *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemory(memory)
	return o
}

// SetMemory adds the memory to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemory(memory *string) {
	o.Memory = memory
}

// WithName adds the name to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithName(name *string) *VirtualizationVirtualMachinesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithOffset(offset *int64) *VirtualizationVirtualMachinesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatform(platform *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformID adds the platformID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatformID(platformID *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithQ adds the q to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithQ(q *string) *VirtualizationVirtualMachinesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegion(region *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegionID(regionID *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRole adds the role to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRole(role *string) *VirtualizationVirtualMachinesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRoleID(roleID *string) *VirtualizationVirtualMachinesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSite adds the site to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSite(site *string) *VirtualizationVirtualMachinesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSiteID(siteID *string) *VirtualizationVirtualMachinesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithStatus(status *string) *VirtualizationVirtualMachinesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTag(tag *string) *VirtualizationVirtualMachinesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenant(tenant *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroup(tenantGroup *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroupID(tenantGroupID *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantID(tenantID *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithVcpus adds the vcpus to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpus(vcpus *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpus(vcpus)
	return o
}

// SetVcpus adds the vcpus to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpus(vcpus *string) {
	o.Vcpus = vcpus
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string
		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {
			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroup != nil {

		// query param cluster_group
		var qrClusterGroup string
		if o.ClusterGroup != nil {
			qrClusterGroup = *o.ClusterGroup
		}
		qClusterGroup := qrClusterGroup
		if qClusterGroup != "" {
			if err := r.SetQueryParam("cluster_group", qClusterGroup); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroupID != nil {

		// query param cluster_group_id
		var qrClusterGroupID string
		if o.ClusterGroupID != nil {
			qrClusterGroupID = *o.ClusterGroupID
		}
		qClusterGroupID := qrClusterGroupID
		if qClusterGroupID != "" {
			if err := r.SetQueryParam("cluster_group_id", qClusterGroupID); err != nil {
				return err
			}
		}

	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string
		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}

	}

	if o.ClusterType != nil {

		// query param cluster_type
		var qrClusterType string
		if o.ClusterType != nil {
			qrClusterType = *o.ClusterType
		}
		qClusterType := qrClusterType
		if qClusterType != "" {
			if err := r.SetQueryParam("cluster_type", qClusterType); err != nil {
				return err
			}
		}

	}

	if o.ClusterTypeID != nil {

		// query param cluster_type_id
		var qrClusterTypeID string
		if o.ClusterTypeID != nil {
			qrClusterTypeID = *o.ClusterTypeID
		}
		qClusterTypeID := qrClusterTypeID
		if qClusterTypeID != "" {
			if err := r.SetQueryParam("cluster_type_id", qClusterTypeID); err != nil {
				return err
			}
		}

	}

	if o.Disk != nil {

		// query param disk
		var qrDisk string
		if o.Disk != nil {
			qrDisk = *o.Disk
		}
		qDisk := qrDisk
		if qDisk != "" {
			if err := r.SetQueryParam("disk", qDisk); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
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

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}

	}

	if o.Memory != nil {

		// query param memory
		var qrMemory string
		if o.Memory != nil {
			qrMemory = *o.Memory
		}
		qMemory := qrMemory
		if qMemory != "" {
			if err := r.SetQueryParam("memory", qMemory); err != nil {
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

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string
		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {
			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
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

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
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

	if o.Vcpus != nil {

		// query param vcpus
		var qrVcpus string
		if o.Vcpus != nil {
			qrVcpus = *o.Vcpus
		}
		qVcpus := qrVcpus
		if qVcpus != "" {
			if err := r.SetQueryParam("vcpus", qVcpus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
