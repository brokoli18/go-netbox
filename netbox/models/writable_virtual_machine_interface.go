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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableVirtualMachineInterface writable virtual machine interface
// swagger:model WritableVirtualMachineInterface
type WritableVirtualMachineInterface struct {

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// Mode
	// Enum: [100 200 300]
	Mode *int64 `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [0 200 800 1000 1120 1130 1150 1170 1050 1100 1200 1300 1310 1320 1350 1400 1420 1500 1510 1650 1520 1550 1600 1700 1750 2600 2610 2620 2630 2640 2810 2820 2830 6100 6200 6300 6400 6500 6600 6700 3010 3020 3040 3080 3160 3320 3400 7010 7020 7030 7040 7050 7060 7070 7080 7090 4000 4010 4040 4050 5000 5050 5100 5150 5200 5300 5310 5320 5330 32767]
	Type int64 `json:"type,omitempty"`

	// Untagged VLAN
	UntaggedVlan *int64 `json:"untagged_vlan,omitempty"`

	// Virtual machine
	VirtualMachine *int64 `json:"virtual_machine,omitempty"`
}

// Validate validates this writable virtual machine interface
func (m *WritableVirtualMachineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

var writableVirtualMachineInterfaceTypeModePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[100,200,300]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineInterfaceTypeModePropEnum = append(writableVirtualMachineInterfaceTypeModePropEnum, v)
	}
}

// prop value enum
func (m *WritableVirtualMachineInterface) validateModeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableVirtualMachineInterfaceTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

var writableVirtualMachineInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,200,800,1000,1120,1130,1150,1170,1050,1100,1200,1300,1310,1320,1350,1400,1420,1500,1510,1650,1520,1550,1600,1700,1750,2600,2610,2620,2630,2640,2810,2820,2830,6100,6200,6300,6400,6500,6600,6700,3010,3020,3040,3080,3160,3320,3400,7010,7020,7030,7040,7050,7060,7070,7080,7090,4000,4010,4040,4050,5000,5050,5100,5150,5200,5300,5310,5320,5330,32767]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineInterfaceTypeTypePropEnum = append(writableVirtualMachineInterfaceTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableVirtualMachineInterface) validateTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableVirtualMachineInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVirtualMachineInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVirtualMachineInterface) UnmarshalBinary(b []byte) error {
	var res WritableVirtualMachineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
