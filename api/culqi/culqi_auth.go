package culqi

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthInfoWriter is the authentication info writer to be used for culqi operations
func AuthInfoWriter(apiKey string) runtime.ClientAuthInfoWriter {
	return &authInfoWrither{
		bearer: "Bearer " + apiKey,
	}
}

type authInfoWrither struct {
	bearer string
}

func (w *authInfoWrither) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return req.SetHeaderParam("Authorization", w.bearer)
}
