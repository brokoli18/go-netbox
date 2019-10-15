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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/brokoli18/go-netbox/netbox/models"
)

// DcimRackGroupsPartialUpdateReader is a Reader for the DcimRackGroupsPartialUpdate structure.
type DcimRackGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackGroupsPartialUpdateOK creates a DcimRackGroupsPartialUpdateOK with default headers values
func NewDcimRackGroupsPartialUpdateOK() *DcimRackGroupsPartialUpdateOK {
	return &DcimRackGroupsPartialUpdateOK{}
}

/*DcimRackGroupsPartialUpdateOK handles this case with default header values.

DcimRackGroupsPartialUpdateOK dcim rack groups partial update o k
*/
type DcimRackGroupsPartialUpdateOK struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-groups/{id}/][%d] dcimRackGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackGroupsPartialUpdateOK) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
