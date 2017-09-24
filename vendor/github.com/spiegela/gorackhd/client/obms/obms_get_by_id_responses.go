package obms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ObmsGetByIDReader is a Reader for the ObmsGetByID structure.
type ObmsGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObmsGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewObmsGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewObmsGetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewObmsGetByIDOK creates a ObmsGetByIDOK with default headers values
func NewObmsGetByIDOK() *ObmsGetByIDOK {
	return &ObmsGetByIDOK{}
}

/*ObmsGetByIDOK handles this case with default header values.

Successfully retrieved the specified OBM service
*/
type ObmsGetByIDOK struct {
	Payload ObmsGetByIDOKBody
}

func (o *ObmsGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /obms/{identifier}][%d] obmsGetByIdOK  %+v", 200, o.Payload)
}

func (o *ObmsGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObmsGetByIDDefault creates a ObmsGetByIDDefault with default headers values
func NewObmsGetByIDDefault(code int) *ObmsGetByIDDefault {
	return &ObmsGetByIDDefault{
		_statusCode: code,
	}
}

/*ObmsGetByIDDefault handles this case with default header values.

Unexpected error
*/
type ObmsGetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the obms get by Id default response
func (o *ObmsGetByIDDefault) Code() int {
	return o._statusCode
}

func (o *ObmsGetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /obms/{identifier}][%d] obmsGetById default  %+v", o._statusCode, o.Payload)
}

func (o *ObmsGetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ObmsGetByIDOKBody obms get by ID o k body
swagger:model ObmsGetByIDOKBody
*/
type ObmsGetByIDOKBody interface{}
