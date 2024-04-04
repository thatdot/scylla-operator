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

// NewColumnFamilyMetricsMaxRowSizeGetParams creates a new ColumnFamilyMetricsMaxRowSizeGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsMaxRowSizeGetParams() *ColumnFamilyMetricsMaxRowSizeGetParams {

	return &ColumnFamilyMetricsMaxRowSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsMaxRowSizeGetParamsWithTimeout creates a new ColumnFamilyMetricsMaxRowSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsMaxRowSizeGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsMaxRowSizeGetParams {

	return &ColumnFamilyMetricsMaxRowSizeGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsMaxRowSizeGetParamsWithContext creates a new ColumnFamilyMetricsMaxRowSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsMaxRowSizeGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsMaxRowSizeGetParams {

	return &ColumnFamilyMetricsMaxRowSizeGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsMaxRowSizeGetParamsWithHTTPClient creates a new ColumnFamilyMetricsMaxRowSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsMaxRowSizeGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsMaxRowSizeGetParams {

	return &ColumnFamilyMetricsMaxRowSizeGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsMaxRowSizeGetParams contains all the parameters to send to the API endpoint
for the column family metrics max row size get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsMaxRowSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsMaxRowSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsMaxRowSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsMaxRowSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics max row size get params
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsMaxRowSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
