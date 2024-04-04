// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v2/models"
)

// FindConfigCompactionEnforceMinThresholdReader is a Reader for the FindConfigCompactionEnforceMinThreshold structure.
type FindConfigCompactionEnforceMinThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionEnforceMinThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCompactionEnforceMinThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCompactionEnforceMinThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionEnforceMinThresholdOK creates a FindConfigCompactionEnforceMinThresholdOK with default headers values
func NewFindConfigCompactionEnforceMinThresholdOK() *FindConfigCompactionEnforceMinThresholdOK {
	return &FindConfigCompactionEnforceMinThresholdOK{}
}

/*
FindConfigCompactionEnforceMinThresholdOK handles this case with default header values.

Config value
*/
type FindConfigCompactionEnforceMinThresholdOK struct {
	Payload bool
}

func (o *FindConfigCompactionEnforceMinThresholdOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigCompactionEnforceMinThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionEnforceMinThresholdDefault creates a FindConfigCompactionEnforceMinThresholdDefault with default headers values
func NewFindConfigCompactionEnforceMinThresholdDefault(code int) *FindConfigCompactionEnforceMinThresholdDefault {
	return &FindConfigCompactionEnforceMinThresholdDefault{
		_statusCode: code,
	}
}

/*
FindConfigCompactionEnforceMinThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionEnforceMinThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction enforce min threshold default response
func (o *FindConfigCompactionEnforceMinThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionEnforceMinThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCompactionEnforceMinThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCompactionEnforceMinThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
