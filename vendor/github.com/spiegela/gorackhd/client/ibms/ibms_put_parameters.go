package ibms

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

	"github.com/spiegela/gorackhd/models"
)

// NewIbmsPutParams creates a new IbmsPutParams object
// with the default values initialized.
func NewIbmsPutParams() *IbmsPutParams {
	var ()
	return &IbmsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsPutParamsWithTimeout creates a new IbmsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsPutParamsWithTimeout(timeout time.Duration) *IbmsPutParams {
	var ()
	return &IbmsPutParams{

		timeout: timeout,
	}
}

// NewIbmsPutParamsWithContext creates a new IbmsPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewIbmsPutParamsWithContext(ctx context.Context) *IbmsPutParams {
	var ()
	return &IbmsPutParams{

		Context: ctx,
	}
}

// NewIbmsPutParamsWithHTTPClient creates a new IbmsPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIbmsPutParamsWithHTTPClient(client *http.Client) *IbmsPutParams {
	var ()
	return &IbmsPutParams{
		HTTPClient: client,
	}
}

/*IbmsPutParams contains all the parameters to send to the API endpoint
for the ibms put operation typically these are written to a http.Request
*/
type IbmsPutParams struct {

	/*Body
	  The IBM settings information to create

	*/
	Body *models.SSHIbmServiceIbm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ibms put params
func (o *IbmsPutParams) WithTimeout(timeout time.Duration) *IbmsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ibms put params
func (o *IbmsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ibms put params
func (o *IbmsPutParams) WithContext(ctx context.Context) *IbmsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ibms put params
func (o *IbmsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ibms put params
func (o *IbmsPutParams) WithHTTPClient(client *http.Client) *IbmsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ibms put params
func (o *IbmsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ibms put params
func (o *IbmsPutParams) WithBody(body *models.SSHIbmServiceIbm) *IbmsPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ibms put params
func (o *IbmsPutParams) SetBody(body *models.SSHIbmServiceIbm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.SSHIbmServiceIbm)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
