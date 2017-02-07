package main

import (
	"fmt"
	"log"

	"github.com/glendc/culqi-go/api/culqi"
	"github.com/glendc/culqi-go/api/culqi/tokens"

	"github.com/go-openapi/strfmt"
)

func main() {
	c := culqi.NewHTTPClient(strfmt.Default)
	auth := culqi.AuthInfoWriter("vk9Xjpe2YZMEOSBzEwiRcPDibnx2NlPBYsusKbDobAk")

	resp, err := c.Tokens.POSTTokens(tokens.NewPOSTTokensParams().
		WithBody(tokens.POSTTokensBody{
			CorreoElectronico: "wmuro@me.com",
			Nombre:            "William",
			Apellido:          "Muro",
			Numero:            4444333322221111,
			Cvv:               123,
			MExp:              9,
			AExp:              2019,
			Guardar:           true,
		}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un token: %q", err)
	}

	// Imprimir la respuesta resultante
	fmt.Println(resp.Payload)
}
