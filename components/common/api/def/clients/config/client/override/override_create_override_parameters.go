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

// NewOverrideCreateOverrideParams creates a new OverrideCreateOverrideParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOverrideCreateOverrideParams() *OverrideCreateOverrideParams {
	return &OverrideCreateOverrideParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOverrideCreateOverrideParamsWithTimeout creates a new OverrideCreateOverrideParams object
// with the ability to set a timeout on a request.
func NewOverrideCreateOverrideParamsWithTimeout(timeout time.Duration) *OverrideCreateOverrideParams {
	return &OverrideCreateOverrideParams{
		timeout: timeout,
	}
}

// NewOverrideCreateOverrideParamsWithContext creates a new OverrideCreateOverrideParams object
// with the ability to set a context for a request.
func NewOverrideCreateOverrideParamsWithContext(ctx context.Context) *OverrideCreateOverrideParams {
	return &OverrideCreateOverrideParams{
		Context: ctx,
	}
}

// NewOverrideCreateOverrideParamsWithHTTPClient creates a new OverrideCreateOverrideParams object
// with the ability to set a custom HTTPClient for a request.
func NewOverrideCreateOverrideParamsWithHTTPClient(client *http.Client) *OverrideCreateOverrideParams {
	return &OverrideCreateOverrideParams{
		HTTPClient: client,
	}
}

/* OverrideCreateOverrideParams contains all the parameters to send to the API endpoint
   for the override create override operation.

   Typically these are written to a http.Request.
*/
type OverrideCreateOverrideParams struct {

	// Body.
	Body *models.ConfigOverride

	/* MetadataProject.

	   Project of the resource
	*/
	MetadataProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the override create override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideCreateOverrideParams) WithDefaults() *OverrideCreateOverrideParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the override create override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideCreateOverrideParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the override create override params
func (o *OverrideCreateOverrideParams) WithTimeout(timeout time.Duration) *OverrideCreateOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the override create override params
func (o *OverrideCreateOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the override create override params
func (o *OverrideCreateOverrideParams) WithContext(ctx context.Context) *OverrideCreateOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the override create override params
func (o *OverrideCreateOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the override create override params
func (o *OverrideCreateOverrideParams) WithHTTPClient(client *http.Client) *OverrideCreateOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the override create override params
func (o *OverrideCreateOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the override create override params
func (o *OverrideCreateOverrideParams) WithBody(body *models.ConfigOverride) *OverrideCreateOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the override create override params
func (o *OverrideCreateOverrideParams) SetBody(body *models.ConfigOverride) {
	o.Body = body
}

// WithMetadataProject adds the metadataProject to the override create override params
func (o *OverrideCreateOverrideParams) WithMetadataProject(metadataProject string) *OverrideCreateOverrideParams {
	o.SetMetadataProject(metadataProject)
	return o
}

// SetMetadataProject adds the metadataProject to the override create override params
func (o *OverrideCreateOverrideParams) SetMetadataProject(metadataProject string) {
	o.MetadataProject = metadataProject
}

// WriteToRequest writes these params to a swagger request
func (o *OverrideCreateOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metadata.project
	if err := r.SetPathParam("metadata.project", o.MetadataProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}