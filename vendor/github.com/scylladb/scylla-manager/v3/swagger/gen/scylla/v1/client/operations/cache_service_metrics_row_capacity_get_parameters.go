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

// NewCacheServiceMetricsRowCapacityGetParams creates a new CacheServiceMetricsRowCapacityGetParams object
// with the default values initialized.
func NewCacheServiceMetricsRowCapacityGetParams() *CacheServiceMetricsRowCapacityGetParams {

	return &CacheServiceMetricsRowCapacityGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowCapacityGetParamsWithTimeout creates a new CacheServiceMetricsRowCapacityGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsRowCapacityGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowCapacityGetParams {

	return &CacheServiceMetricsRowCapacityGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowCapacityGetParamsWithContext creates a new CacheServiceMetricsRowCapacityGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsRowCapacityGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowCapacityGetParams {

	return &CacheServiceMetricsRowCapacityGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsRowCapacityGetParamsWithHTTPClient creates a new CacheServiceMetricsRowCapacityGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsRowCapacityGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowCapacityGetParams {

	return &CacheServiceMetricsRowCapacityGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsRowCapacityGetParams contains all the parameters to send to the API endpoint
for the cache service metrics row capacity get operation typically these are written to a http.Request
*/
type CacheServiceMetricsRowCapacityGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowCapacityGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowCapacityGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowCapacityGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row capacity get params
func (o *CacheServiceMetricsRowCapacityGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowCapacityGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
