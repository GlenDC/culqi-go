package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETCargosIDParams creates a new GETCargosIDParams object
// with the default values initialized.
func NewGETCargosIDParams() *GETCargosIDParams {
	var (
		authorizationDefault = string("Bearer <<apiKey>>")
	)
	return &GETCargosIDParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGETCargosIDParamsWithTimeout creates a new GETCargosIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETCargosIDParamsWithTimeout(timeout time.Duration) *GETCargosIDParams {
	var (
		authorizationDefault = string("Bearer <<apiKey>>")
	)
	return &GETCargosIDParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewGETCargosIDParamsWithContext creates a new GETCargosIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETCargosIDParamsWithContext(ctx context.Context) *GETCargosIDParams {
	var (
		authorizationDefault = string("Bearer <<apiKey>>")
	)
	return &GETCargosIDParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewGETCargosIDParamsWithHTTPClient creates a new GETCargosIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETCargosIDParamsWithHTTPClient(client *http.Client) *GETCargosIDParams {
	var (
		authorizationDefault = string("Bearer <<apiKey>>")
	)
	return &GETCargosIDParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*GETCargosIDParams contains all the parameters to send to the API endpoint
for the g e t cargos id operation typically these are written to a http.Request
*/
type GETCargosIDParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t cargos id params
func (o *GETCargosIDParams) WithTimeout(timeout time.Duration) *GETCargosIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t cargos id params
func (o *GETCargosIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t cargos id params
func (o *GETCargosIDParams) WithContext(ctx context.Context) *GETCargosIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t cargos id params
func (o *GETCargosIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t cargos id params
func (o *GETCargosIDParams) WithHTTPClient(client *http.Client) *GETCargosIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t cargos id params
func (o *GETCargosIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the g e t cargos id params
func (o *GETCargosIDParams) WithAuthorization(authorization string) *GETCargosIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the g e t cargos id params
func (o *GETCargosIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the g e t cargos id params
func (o *GETCargosIDParams) WithID(id string) *GETCargosIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t cargos id params
func (o *GETCargosIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETCargosIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
