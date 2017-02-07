package main

import (
	"fmt"
	"log"

	"github.com/glendc/culqi-go/api/culqi"
	"github.com/glendc/culqi-go/api/culqi/cargos"
	"github.com/glendc/culqi-go/api/culqi/tokens"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

func main() {
	c := culqi.NewHTTPClient(strfmt.Default)
	auth := culqi.AuthInfoWriter("vk9Xjpe2YZMEOSBzEwiRcPDibnx2NlPBYsusKbDobAk")

	tokenResp, err := c.Tokens.POSTTokens(tokens.NewPOSTTokensParams().
		WithBody(tokens.POSTTokensBody{
			CorreoElectronico: "wmuro@me.com",
			Nombre:            "Will",
			Apellido:          "Muro",
			Numero:            4444333322221111,
			Cvv:               123,
			MExp:              9,
			AExp:              2019,
			Guardar:           false,
		}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un token: %q", err)
	}

	resp, err := c.Cargos.POSTCargos(
		cargos.NewPOSTCargosParams().
			WithBody(cargos.POSTCargosBody{
				Token:             tokenResp.Payload.ID,
				Moneda:            swag.String("PEN"),
				Monto:             swag.Int64(19900),
				Descripcion:       swag.String("Venta de prueba"),
				Pedido:            swag.String("11213351"),
				CodigoPais:        swag.String("PE"),
				Ciudad:            swag.String("Lima"),
				Usuario:           swag.String("71701956"),
				Direccion:         swag.String("Avenida Lima 1232"),
				Telefono:          swag.Int64(12313123),
				Nombres:           swag.String("Will"),
				Apellidos:         swag.String("Muro"),
				CorreoElectronico: conv.Email(strfmt.Email("wmuro@me.com")),
			}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un cargo: %q", err)
	}

	// Imprimir la respuesta resultante
	fmt.Println(resp)
}
