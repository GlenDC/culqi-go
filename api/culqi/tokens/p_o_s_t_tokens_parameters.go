package tokens

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

// NewPOSTTokensParams creates a new POSTTokensParams object
// with the default values initialized.
func NewPOSTTokensParams() *POSTTokensParams {
	var (
		authorizationDefault = string("Bearer <<CodigoComercio>>")
	)
	return &POSTTokensParams{
		Authorization: authorizationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTTokensParamsWithTimeout creates a new POSTTokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTTokensParamsWithTimeout(timeout time.Duration) *POSTTokensParams {
	var (
		authorizationDefault = string("Bearer <<CodigoComercio>>")
	)
	return &POSTTokensParams{
		Authorization: authorizationDefault,

		timeout: timeout,
	}
}

// NewPOSTTokensParamsWithContext creates a new POSTTokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTTokensParamsWithContext(ctx context.Context) *POSTTokensParams {
	var (
		authorizationDefault = string("Bearer <<CodigoComercio>>")
	)
	return &POSTTokensParams{
		Authorization: authorizationDefault,

		Context: ctx,
	}
}

// NewPOSTTokensParamsWithHTTPClient creates a new POSTTokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPOSTTokensParamsWithHTTPClient(client *http.Client) *POSTTokensParams {
	var (
		authorizationDefault = string("Bearer <<CodigoComercio>>")
	)
	return &POSTTokensParams{
		Authorization: authorizationDefault,
		HTTPClient:    client,
	}
}

/*POSTTokensParams contains all the parameters to send to the API endpoint
for the p o s t tokens operation typically these are written to a http.Request
*/
type POSTTokensParams struct {

	/*Authorization*/
	Authorization string
	/*Body*/
	Body POSTTokensBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p o s t tokens params
func (o *POSTTokensParams) WithTimeout(timeout time.Duration) *POSTTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t tokens params
func (o *POSTTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t tokens params
func (o *POSTTokensParams) WithContext(ctx context.Context) *POSTTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t tokens params
func (o *POSTTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p o s t tokens params
func (o *POSTTokensParams) WithHTTPClient(client *http.Client) *POSTTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p o s t tokens params
func (o *POSTTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the p o s t tokens params
func (o *POSTTokensParams) WithAuthorization(authorization string) *POSTTokensParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the p o s t tokens params
func (o *POSTTokensParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the p o s t tokens params
func (o *POSTTokensParams) WithBody(body POSTTokensBody) *POSTTokensParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p o s t tokens params
func (o *POSTTokensParams) SetBody(body POSTTokensBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *POSTTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
