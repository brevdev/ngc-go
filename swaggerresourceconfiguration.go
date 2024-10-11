// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/brevdev/ngc-go/option"
)

// SwaggerResourceConfigurationService contains methods and other services that
// help with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwaggerResourceConfigurationService] method instead.
type SwaggerResourceConfigurationService struct {
	Options  []option.RequestOption
	Ui       *SwaggerResourceConfigurationUiService
	Security *SwaggerResourceConfigurationSecurityService
}

// NewSwaggerResourceConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSwaggerResourceConfigurationService(opts ...option.RequestOption) (r *SwaggerResourceConfigurationService) {
	r = &SwaggerResourceConfigurationService{}
	r.Options = opts
	r.Ui = NewSwaggerResourceConfigurationUiService(opts...)
	r.Security = NewSwaggerResourceConfigurationSecurityService(opts...)
	return
}
