// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStorageServiceRelocalSchemaPostParams creates a new StorageServiceRelocalSchemaPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceRelocalSchemaPostParams() *StorageServiceRelocalSchemaPostParams {
	return &StorageServiceRelocalSchemaPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceRelocalSchemaPostParamsWithTimeout creates a new StorageServiceRelocalSchemaPostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceRelocalSchemaPostParamsWithTimeout(timeout time.Duration) *StorageServiceRelocalSchemaPostParams {
	return &StorageServiceRelocalSchemaPostParams{
		timeout: timeout,
	}
}

// NewStorageServiceRelocalSchemaPostParamsWithContext creates a new StorageServiceRelocalSchemaPostParams object
// with the ability to set a context for a request.
func NewStorageServiceRelocalSchemaPostParamsWithContext(ctx context.Context) *StorageServiceRelocalSchemaPostParams {
	return &StorageServiceRelocalSchemaPostParams{
		Context: ctx,
	}
}

// NewStorageServiceRelocalSchemaPostParamsWithHTTPClient creates a new StorageServiceRelocalSchemaPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceRelocalSchemaPostParamsWithHTTPClient(client *http.Client) *StorageServiceRelocalSchemaPostParams {
	return &StorageServiceRelocalSchemaPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceRelocalSchemaPostParams contains all the parameters to send to the API endpoint

	for the storage service relocal schema post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceRelocalSchemaPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service relocal schema post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceRelocalSchemaPostParams) WithDefaults() *StorageServiceRelocalSchemaPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service relocal schema post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceRelocalSchemaPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) WithTimeout(timeout time.Duration) *StorageServiceRelocalSchemaPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) WithContext(ctx context.Context) *StorageServiceRelocalSchemaPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) WithHTTPClient(client *http.Client) *StorageServiceRelocalSchemaPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service relocal schema post params
func (o *StorageServiceRelocalSchemaPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceRelocalSchemaPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}