package planes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// POSTPlanesReader is a Reader for the POSTPlanes structure.
type POSTPlanesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTPlanesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPOSTPlanesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPOSTPlanesDefault creates a POSTPlanesDefault with default headers values
func NewPOSTPlanesDefault(code int) *POSTPlanesDefault {
	return &POSTPlanesDefault{
		_statusCode: code,
	}
}

/*POSTPlanesDefault handles this case with default header values.

POSTPlanesDefault p o s t planes default
*/
type POSTPlanesDefault struct {
	_statusCode int

	Payload POSTPlanesDefaultBody
}

// Code gets the status code for the p o s t planes default response
func (o *POSTPlanesDefault) Code() int {
	return o._statusCode
}

func (o *POSTPlanesDefault) Error() string {
	return fmt.Sprintf("[POST /planes][%d] POST_planes default  %+v", o._statusCode, o.Payload)
}

func (o *POSTPlanesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*POSTPlanesBody p o s t planes body
swagger:model POSTPlanesBody
*/
type POSTPlanesBody struct {

	// ciclos
	Ciclos int64 `json:"ciclos,omitempty"`

	// codigo comercio
	CodigoComercio string `json:"codigo_comercio,omitempty"`

	// gracia
	Gracia int64 `json:"gracia,omitempty"`

	// gracia medida
	GraciaMedida string `json:"gracia_medida,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// intervalo
	Intervalo int64 `json:"intervalo,omitempty"`

	// moneda
	Moneda string `json:"moneda,omitempty"`

	// monto
	Monto string `json:"monto,omitempty"`

	// nombre
	Nombre string `json:"nombre,omitempty"`

	// periodo
	Periodo string `json:"periodo,omitempty"`
}

/*POSTPlanesDefaultBody p o s t planes default body
swagger:model POSTPlanesDefaultBody
*/
type POSTPlanesDefaultBody interface{}
