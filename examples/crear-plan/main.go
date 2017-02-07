package main

import (
	"log"

	"github.com/glendc/culqi-go/api/culqi"
	"github.com/glendc/culqi-go/api/culqi/planes"

	"github.com/go-openapi/strfmt"
)

func main() {
	c := culqi.NewHTTPClient(strfmt.Default)
	auth := culqi.AuthInfoWriter("vk9Xjpe2YZMEOSBzEwiRcPDibnx2NlPBYsusKbDobAk")

	err := c.Planes.POSTPlanes(
		planes.NewPOSTPlanesParams().WithBody(planes.POSTPlanesBody{
			CodigoComercio: "X1QjlYMBBSV8",
			Moneda:         "PEN",
			Monto:          "1000",
			ID:             "plan-12345",
			Periodo:        "dias",
			Nombre:         "Plan de prueba",
			Intervalo:      2,
			Gracia:         5,
			GraciaMedida:   "dias",
			Ciclos:         12,
		}), auth)
	if err != nil {
		log.Fatalf("no se pudo crear un plan: %q", err)
	}
}
