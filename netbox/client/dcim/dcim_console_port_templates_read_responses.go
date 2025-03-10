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

// DcimConsolePortTemplatesReadReader is a Reader for the DcimConsolePortTemplatesRead structure.
type DcimConsolePortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesReadOK creates a DcimConsolePortTemplatesReadOK with default headers values
func NewDcimConsolePortTemplatesReadOK() *DcimConsolePortTemplatesReadOK {
	return &DcimConsolePortTemplatesReadOK{}
}

/*DcimConsolePortTemplatesReadOK handles this case with default header values.

DcimConsolePortTemplatesReadOK dcim console port templates read o k
*/
type DcimConsolePortTemplatesReadOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesReadOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
