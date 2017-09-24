package obms

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

// NewObmsGetByIDParams creates a new ObmsGetByIDParams object
// with the default values initialized.
func NewObmsGetByIDParams() *ObmsGetByIDParams {
	var ()
	return &ObmsGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObmsGetByIDParamsWithTimeout creates a new ObmsGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObmsGetByIDParamsWithTimeout(timeout time.Duration) *ObmsGetByIDParams {
	var ()
	return &ObmsGetByIDParams{

		timeout: timeout,
	}
}

// NewObmsGetByIDParamsWithContext creates a new ObmsGetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewObmsGetByIDParamsWithContext(ctx context.Context) *ObmsGetByIDParams {
	var ()
	return &ObmsGetByIDParams{

		Context: ctx,
	}
}

// NewObmsGetByIDParamsWithHTTPClient creates a new ObmsGetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObmsGetByIDParamsWithHTTPClient(client *http.Client) *ObmsGetByIDParams {
	var ()
	return &ObmsGetByIDParams{
		HTTPClient: client,
	}
}

/*ObmsGetByIDParams contains all the parameters to send to the API endpoint
for the obms get by Id operation typically these are written to a http.Request
*/
type ObmsGetByIDParams struct {

	/*Identifier
	  The OBM service identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the obms get by Id params
func (o *ObmsGetByIDParams) WithTimeout(timeout time.Duration) *ObmsGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the obms get by Id params
func (o *ObmsGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the obms get by Id params
func (o *ObmsGetByIDParams) WithContext(ctx context.Context) *ObmsGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the obms get by Id params
func (o *ObmsGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the obms get by Id params
func (o *ObmsGetByIDParams) WithHTTPClient(client *http.Client) *ObmsGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the obms get by Id params
func (o *ObmsGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the obms get by Id params
func (o *ObmsGetByIDParams) WithIdentifier(identifier string) *ObmsGetByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the obms get by Id params
func (o *ObmsGetByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *ObmsGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
