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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// IPAMPrefixesReadReader is a Reader for the IPAMPrefixesRead structure.
type IPAMPrefixesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMPrefixesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesReadOK creates a IPAMPrefixesReadOK with default headers values
func NewIPAMPrefixesReadOK() *IPAMPrefixesReadOK {
	return &IPAMPrefixesReadOK{}
}

/*IPAMPrefixesReadOK handles this case with default header values.

IPAMPrefixesReadOK ipam prefixes read o k
*/
type IPAMPrefixesReadOK struct {
	Payload *models.Prefix
}

func (o *IPAMPrefixesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/][%d] ipamPrefixesReadOK  %+v", 200, o.Payload)
}

func (o *IPAMPrefixesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
