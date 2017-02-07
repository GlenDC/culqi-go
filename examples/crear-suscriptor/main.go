package main

import (
	"fmt"
	"log"

	"github.com/glendc/culqi-go/api/culqi"
	"github.com/glendc/culqi-go/api/culqi/suscripciones"
	"github.com/glendc/culqi-go/api/culqi/tokens"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func main() {
	c := culqi.NewHTTPClient(strfmt.Default)
	auth := culqi.AuthInfoWriter("vk9Xjpe2YZMEOSBzEwiRcPDibnx2NlPBYsusKbDobAk")

	tokenResp, err := c.Tokens.POSTTokens(tokens.NewPOSTTokensParams().
		WithBody(tokens.POSTTokensBody{
			CorreoElectronico: "soporte@culqi.com",
			Nombre:            "Jon",
			Apellido:          "Doe",
			Numero:            4444333322221111,
			Cvv:               123,
			MExp:              9,
			AExp:              2019,
			Guardar:           false,
		}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un token: %q", err)
	}

	resp, err := c.Suscripciones.POSTSuscripciones(
		suscripciones.NewPOSTSuscripcionesParams().WithBody(
			suscripciones.POSTSuscripcionesBody{
				Token:             tokenResp.Payload.ID,
				CodigoPais:        swag.String("PE"),
				Direccion:         swag.String("Avenida Lima 123213"),
				Ciudad:            swag.String("Lima"),
				Usuario:           swag.String("soporteculqi"),
				Telefono:          swag.Int64(1234567789),
				Nombre:            swag.String("Jon"),
				Apellido:          swag.String("Doe"),
				CorreoElectronico: swag.String("soporte@culqi.com"),
				PlanID:            swag.String("plan-basico"),
			}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un suscripor: %q", err)
	}

	// Imprimir la respuesta resultante
	fmt.Println(resp.Payload)
}
