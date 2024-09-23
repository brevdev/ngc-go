// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/option"
)

// OrgV3TeamService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV3TeamService] method instead.
type OrgV3TeamService struct {
	Options []option.RequestOption
	Users   *OrgV3TeamUserService
}

// NewOrgV3TeamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgV3TeamService(opts ...option.RequestOption) (r *OrgV3TeamService) {
	r = &OrgV3TeamService{}
	r.Options = opts
	r.Users = NewOrgV3TeamUserService(opts...)
	return
}
