// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/brevdev/ngc-go/option"
)

// OrgV3TeamUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV3TeamUserService] method instead.
type OrgV3TeamUserService struct {
	Options        []option.RequestOption
	NcaInvitations *OrgV3TeamUserNcaInvitationService
}

// NewOrgV3TeamUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgV3TeamUserService(opts ...option.RequestOption) (r *OrgV3TeamUserService) {
	r = &OrgV3TeamUserService{}
	r.Options = opts
	r.NcaInvitations = NewOrgV3TeamUserNcaInvitationService(opts...)
	return
}
