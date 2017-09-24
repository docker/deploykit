package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTemplatesGetByNameParams creates a new TemplatesGetByNameParams object
// with the default values initialized.
func NewTemplatesGetByNameParams() *TemplatesGetByNameParams {
	var ()
	return &TemplatesGetByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesGetByNameParamsWithTimeout creates a new TemplatesGetByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTemplatesGetByNameParamsWithTimeout(timeout time.Duration) *TemplatesGetByNameParams {
	var ()
	return &TemplatesGetByNameParams{

		timeout: timeout,
	}
}

// NewTemplatesGetByNameParamsWithContext creates a new TemplatesGetByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewTemplatesGetByNameParamsWithContext(ctx context.Context) *TemplatesGetByNameParams {
	var ()
	return &TemplatesGetByNameParams{

		Context: ctx,
	}
}

// NewTemplatesGetByNameParamsWithHTTPClient creates a new TemplatesGetByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTemplatesGetByNameParamsWithHTTPClient(client *http.Client) *TemplatesGetByNameParams {
	var ()
	return &TemplatesGetByNameParams{
		HTTPClient: client,
	}
}

/*TemplatesGetByNameParams contains all the parameters to send to the API endpoint
for the templates get by name operation typically these are written to a http.Request
*/
type TemplatesGetByNameParams struct {

	/*Macs
	  List of valid MAC addresses to lookup

	*/
	Macs []string
	/*Name
	  The name of the template

	*/
	Name string
	/*NodeID
	  The node identifier

	*/
	NodeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the templates get by name params
func (o *TemplatesGetByNameParams) WithTimeout(timeout time.Duration) *TemplatesGetByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates get by name params
func (o *TemplatesGetByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates get by name params
func (o *TemplatesGetByNameParams) WithContext(ctx context.Context) *TemplatesGetByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates get by name params
func (o *TemplatesGetByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates get by name params
func (o *TemplatesGetByNameParams) WithHTTPClient(client *http.Client) *TemplatesGetByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates get by name params
func (o *TemplatesGetByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacs adds the macs to the templates get by name params
func (o *TemplatesGetByNameParams) WithMacs(macs []string) *TemplatesGetByNameParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the templates get by name params
func (o *TemplatesGetByNameParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WithName adds the name to the templates get by name params
func (o *TemplatesGetByNameParams) WithName(name string) *TemplatesGetByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the templates get by name params
func (o *TemplatesGetByNameParams) SetName(name string) {
	o.Name = name
}

// WithNodeID adds the nodeID to the templates get by name params
func (o *TemplatesGetByNameParams) WithNodeID(nodeID *string) *TemplatesGetByNameParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the templates get by name params
func (o *TemplatesGetByNameParams) SetNodeID(nodeID *string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesGetByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesMacs := o.Macs

	joinedMacs := swag.JoinByFormat(valuesMacs, "multi")
	// query array param macs
	if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.NodeID != nil {

		// query param nodeId
		var qrNodeID string
		if o.NodeID != nil {
			qrNodeID = *o.NodeID
		}
		qNodeID := qrNodeID
		if qNodeID != "" {
			if err := r.SetQueryParam("nodeId", qNodeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}