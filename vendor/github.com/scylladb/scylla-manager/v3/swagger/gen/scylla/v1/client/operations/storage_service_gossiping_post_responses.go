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

// StorageServiceGossipingPostReader is a Reader for the StorageServiceGossipingPost structure.
type StorageServiceGossipingPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceGossipingPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceGossipingPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceGossipingPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceGossipingPostOK creates a StorageServiceGossipingPostOK with default headers values
func NewStorageServiceGossipingPostOK() *StorageServiceGossipingPostOK {
	return &StorageServiceGossipingPostOK{}
}

/*
StorageServiceGossipingPostOK handles this case with default header values.

Success
*/
type StorageServiceGossipingPostOK struct {
}

func (o *StorageServiceGossipingPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceGossipingPostDefault creates a StorageServiceGossipingPostDefault with default headers values
func NewStorageServiceGossipingPostDefault(code int) *StorageServiceGossipingPostDefault {
	return &StorageServiceGossipingPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceGossipingPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceGossipingPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service gossiping post default response
func (o *StorageServiceGossipingPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceGossipingPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceGossipingPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceGossipingPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
