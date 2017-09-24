package nodes

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

// NewNodesPostWorkflowByIDParams creates a new NodesPostWorkflowByIDParams object
// with the default values initialized.
func NewNodesPostWorkflowByIDParams() *NodesPostWorkflowByIDParams {
	var ()
	return &NodesPostWorkflowByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesPostWorkflowByIDParamsWithTimeout creates a new NodesPostWorkflowByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesPostWorkflowByIDParamsWithTimeout(timeout time.Duration) *NodesPostWorkflowByIDParams {
	var ()
	return &NodesPostWorkflowByIDParams{

		timeout: timeout,
	}
}

// NewNodesPostWorkflowByIDParamsWithContext creates a new NodesPostWorkflowByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesPostWorkflowByIDParamsWithContext(ctx context.Context) *NodesPostWorkflowByIDParams {
	var ()
	return &NodesPostWorkflowByIDParams{

		Context: ctx,
	}
}

// NewNodesPostWorkflowByIDParamsWithHTTPClient creates a new NodesPostWorkflowByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesPostWorkflowByIDParamsWithHTTPClient(client *http.Client) *NodesPostWorkflowByIDParams {
	var ()
	return &NodesPostWorkflowByIDParams{
		HTTPClient: client,
	}
}

/*NodesPostWorkflowByIDParams contains all the parameters to send to the API endpoint
for the nodes post workflow by Id operation typically these are written to a http.Request
*/
type NodesPostWorkflowByIDParams struct {

	/*Body
	  The name property set to the injectableName property of the workflow graph

	*/
	Body *models.PostNodeWorkflow
	/*Identifier
	  The node identifier

	*/
	Identifier string
	/*Name
	  The optional name of the workflow graph to run

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithTimeout(timeout time.Duration) *NodesPostWorkflowByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithContext(ctx context.Context) *NodesPostWorkflowByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithHTTPClient(client *http.Client) *NodesPostWorkflowByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithBody(body *models.PostNodeWorkflow) *NodesPostWorkflowByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetBody(body *models.PostNodeWorkflow) {
	o.Body = body
}

// WithIdentifier adds the identifier to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithIdentifier(identifier string) *NodesPostWorkflowByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WithName adds the name to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) WithName(name *string) *NodesPostWorkflowByIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the nodes post workflow by Id params
func (o *NodesPostWorkflowByIDParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *NodesPostWorkflowByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.PostNodeWorkflow)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
