// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/brevdev/ngc-go/option"
)

// AdminOrgRegistryService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgRegistryService] method instead.
type AdminOrgRegistryService struct {
	Options  []option.RequestOption
	Metering *AdminOrgRegistryMeteringService
}

// NewAdminOrgRegistryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgRegistryService(opts ...option.RequestOption) (r *AdminOrgRegistryService) {
	r = &AdminOrgRegistryService{}
	r.Options = opts
	r.Metering = NewAdminOrgRegistryMeteringService(opts...)
	return
}
