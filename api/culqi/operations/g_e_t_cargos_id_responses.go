package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GETCargosIDReader is a Reader for the GETCargosID structure.
type GETCargosIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCargosIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETCargosIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETCargosIDOK creates a GETCargosIDOK with default headers values
func NewGETCargosIDOK() *GETCargosIDOK {
	return &GETCargosIDOK{}
}

/*GETCargosIDOK handles this case with default header values.

GETCargosIDOK g e t cargos Id o k
*/
type GETCargosIDOK struct {
	Payload GETCargosIDOKBody
}

func (o *GETCargosIDOK) Error() string {
	return fmt.Sprintf("[GET /cargos/{id}][%d] gETCargosIdOK  %+v", 200, o.Payload)
}

func (o *GETCargosIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GETCargosIDOKBody g e t cargos ID o k body
swagger:model GETCargosIDOKBody
*/
type GETCargosIDOKBody struct {

	// creado
	// Required: true
	Creado *int64 `json:"creado"`

	// descripcion
	// Required: true
	Descripcion *string `json:"descripcion"`

	// id
	// Required: true
	ID *string `json:"id"`

	// moneda
	// Required: true
	Moneda *string `json:"moneda"`

	// monto
	// Required: true
	Monto *int64 `json:"monto"`

	// objeto
	// Required: true
	Objeto *string `json:"objeto"`

	// pedido
	// Required: true
	Pedido *string `json:"pedido"`

	// token
	// Required: true
	Token *GETCargosIDOKBodyToken `json:"token"`

	// usuario
	// Required: true
	Usuario *string `json:"usuario"`
}

// Validate validates this g e t cargos ID o k body
func (o *GETCargosIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreado(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDescripcion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMoneda(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMonto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateObjeto(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePedido(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUsuario(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GETCargosIDOKBody) validateCreado(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"creado", "body", o.Creado); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateDescripcion(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"descripcion", "body", o.Descripcion); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateMoneda(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"moneda", "body", o.Moneda); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateMonto(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"monto", "body", o.Monto); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateObjeto(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"objeto", "body", o.Objeto); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validatePedido(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"pedido", "body", o.Pedido); err != nil {
		return err
	}

	return nil
}

func (o *GETCargosIDOKBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"token", "body", o.Token); err != nil {
		return err
	}

	if o.Token != nil {

		if err := o.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gETCargosIdOK" + "." + "token")
			}
			return err
		}
	}

	return nil
}

func (o *GETCargosIDOKBody) validateUsuario(formats strfmt.Registry) error {

	if err := validate.Required("gETCargosIdOK"+"."+"usuario", "body", o.Usuario); err != nil {
		return err
	}

	return nil
}

/*GETCargosIDOKBodyToken g e t cargos ID o k body token
swagger:model GETCargosIDOKBodyToken
*/
type GETCargosIDOKBodyToken struct {

	// id
	ID string `json:"id,omitempty"`

	// objeto
	Objeto string `json:"objeto,omitempty"`

	// tarjeta
	Tarjeta *GETCargosIDOKBodyTokenTarjeta `json:"tarjeta,omitempty"`
}

// Validate validates this g e t cargos ID o k body token
func (o *GETCargosIDOKBodyToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTarjeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GETCargosIDOKBodyToken) validateTarjeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Tarjeta) { // not required
		return nil
	}

	if o.Tarjeta != nil {

		if err := o.Tarjeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gETCargosIdOK" + "." + "token" + "." + "tarjeta")
			}
			return err
		}
	}

	return nil
}

/*GETCargosIDOKBodyTokenTarjeta g e t cargos ID o k body token tarjeta
swagger:model GETCargosIDOKBodyTokenTarjeta
*/
type GETCargosIDOKBodyTokenTarjeta struct {

	// apellido
	Apellido string `json:"apellido,omitempty"`

	// bin
	Bin string `json:"bin,omitempty"`

	// emisor
	Emisor string `json:"emisor,omitempty"`

	// marca
	Marca string `json:"marca,omitempty"`

	// nombre
	Nombre string `json:"nombre,omitempty"`

	// numero
	Numero string `json:"numero,omitempty"`

	// pais
	Pais string `json:"pais,omitempty"`
}

// Validate validates this g e t cargos ID o k body token tarjeta
func (o *GETCargosIDOKBodyTokenTarjeta) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
