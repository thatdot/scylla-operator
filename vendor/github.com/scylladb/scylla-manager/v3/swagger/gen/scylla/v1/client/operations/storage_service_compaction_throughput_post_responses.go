// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceCompactionThroughputPostReader is a Reader for the StorageServiceCompactionThroughputPost structure.
type StorageServiceCompactionThroughputPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceCompactionThroughputPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceCompactionThroughputPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceCompactionThroughputPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceCompactionThroughputPostOK creates a StorageServiceCompactionThroughputPostOK with default headers values
func NewStorageServiceCompactionThroughputPostOK() *StorageServiceCompactionThroughputPostOK {
	return &StorageServiceCompactionThroughputPostOK{}
}

/*
StorageServiceCompactionThroughputPostOK handles this case with default header values.

Success
*/
type StorageServiceCompactionThroughputPostOK struct {
}

func (o *StorageServiceCompactionThroughputPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceCompactionThroughputPostDefault creates a StorageServiceCompactionThroughputPostDefault with default headers values
func NewStorageServiceCompactionThroughputPostDefault(code int) *StorageServiceCompactionThroughputPostDefault {
	return &StorageServiceCompactionThroughputPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceCompactionThroughputPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceCompactionThroughputPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service compaction throughput post default response
func (o *StorageServiceCompactionThroughputPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceCompactionThroughputPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceCompactionThroughputPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceCompactionThroughputPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
