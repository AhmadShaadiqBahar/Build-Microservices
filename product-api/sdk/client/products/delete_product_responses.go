// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/AhmadShaadiqBahar/build-microservices/product-api/sdk/models"
)

// DeleteProductReader is a Reader for the DeleteProduct structure.
type DeleteProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductCreated creates a DeleteProductCreated with default headers values
func NewDeleteProductCreated() *DeleteProductCreated {
	return &DeleteProductCreated{}
}

/*DeleteProductCreated handles this case with default header values.

No content is returned by this API endpoint
*/
type DeleteProductCreated struct {
}

func (o *DeleteProductCreated) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductCreated ", 201)
}

func (o *DeleteProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductNotFound creates a DeleteProductNotFound with default headers values
func NewDeleteProductNotFound() *DeleteProductNotFound {
	return &DeleteProductNotFound{}
}

/*DeleteProductNotFound handles this case with default header values.

Generic error message returned as a string
*/
type DeleteProductNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteProductNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProductNotImplemented creates a DeleteProductNotImplemented with default headers values
func NewDeleteProductNotImplemented() *DeleteProductNotImplemented {
	return &DeleteProductNotImplemented{}
}

/*DeleteProductNotImplemented handles this case with default header values.

Generic error message returned as a string
*/
type DeleteProductNotImplemented struct {
	Payload *models.GenericError
}

func (o *DeleteProductNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotImplemented  %+v", 501, o.Payload)
}

func (o *DeleteProductNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
