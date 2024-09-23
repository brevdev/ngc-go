// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/option"
)

// OrgV2Service contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV2Service] method instead.
type OrgV2Service struct {
	Options []option.RequestOption
	Users   *OrgV2UserService
	Teams   *OrgV2TeamService
}

// NewOrgV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgV2Service(opts ...option.RequestOption) (r *OrgV2Service) {
	r = &OrgV2Service{}
	r.Options = opts
	r.Users = NewOrgV2UserService(opts...)
	r.Teams = NewOrgV2TeamService(opts...)
	return
}
