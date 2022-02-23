// Code generated by go-swagger; DO NOT EDIT.

package override

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

	"github.com/RafaySystems/rcloud-base/components/common/api/def/clients/config/models"
)

// NewOverrideAddDefaultOverridesParams creates a new OverrideAddDefaultOverridesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOverrideAddDefaultOverridesParams() *OverrideAddDefaultOverridesParams {
	return &OverrideAddDefaultOverridesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOverrideAddDefaultOverridesParamsWithTimeout creates a new OverrideAddDefaultOverridesParams object
// with the ability to set a timeout on a request.
func NewOverrideAddDefaultOverridesParamsWithTimeout(timeout time.Duration) *OverrideAddDefaultOverridesParams {
	return &OverrideAddDefaultOverridesParams{
		timeout: timeout,
	}
}

// NewOverrideAddDefaultOverridesParamsWithContext creates a new OverrideAddDefaultOverridesParams object
// with the ability to set a context for a request.
func NewOverrideAddDefaultOverridesParamsWithContext(ctx context.Context) *OverrideAddDefaultOverridesParams {
	return &OverrideAddDefaultOverridesParams{
		Context: ctx,
	}
}

// NewOverrideAddDefaultOverridesParamsWithHTTPClient creates a new OverrideAddDefaultOverridesParams object
// with the ability to set a custom HTTPClient for a request.
func NewOverrideAddDefaultOverridesParamsWithHTTPClient(client *http.Client) *OverrideAddDefaultOverridesParams {
	return &OverrideAddDefaultOverridesParams{
		HTTPClient: client,
	}
}

/* OverrideAddDefaultOverridesParams contains all the parameters to send to the API endpoint
   for the override add default overrides operation.

   Typically these are written to a http.Request.
*/
type OverrideAddDefaultOverridesParams struct {

	// Body.
	Body *models.RPCAddDefaultOverridesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the override add default overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideAddDefaultOverridesParams) WithDefaults() *OverrideAddDefaultOverridesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the override add default overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideAddDefaultOverridesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) WithTimeout(timeout time.Duration) *OverrideAddDefaultOverridesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) WithContext(ctx context.Context) *OverrideAddDefaultOverridesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) WithHTTPClient(client *http.Client) *OverrideAddDefaultOverridesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) WithBody(body *models.RPCAddDefaultOverridesRequest) *OverrideAddDefaultOverridesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the override add default overrides params
func (o *OverrideAddDefaultOverridesParams) SetBody(body *models.RPCAddDefaultOverridesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OverrideAddDefaultOverridesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}