// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/option"
)

// AdminOrgUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgUserService] method instead.
type AdminOrgUserService struct {
	Options []option.RequestOption
}

// NewAdminOrgUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgUserService(opts ...option.RequestOption) (r *AdminOrgUserService) {
	r = &AdminOrgUserService{}
	r.Options = opts
	return
}
