// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/brevdev/ngc-go/option"
)

// OrgV3Service contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV3Service] method instead.
type OrgV3Service struct {
	Options []option.RequestOption
	Users   *OrgV3UserService
	Teams   *OrgV3TeamService
}

// NewOrgV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgV3Service(opts ...option.RequestOption) (r *OrgV3Service) {
	r = &OrgV3Service{}
	r.Options = opts
	r.Users = NewOrgV3UserService(opts...)
	r.Teams = NewOrgV3TeamService(opts...)
	return
}
