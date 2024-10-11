// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminOrgTeamService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgTeamService] method instead.
type AdminOrgTeamService struct {
	Options []option.RequestOption
	Users   *AdminOrgTeamUserService
}

// NewAdminOrgTeamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgTeamService(opts ...option.RequestOption) (r *AdminOrgTeamService) {
	r = &AdminOrgTeamService{}
	r.Options = opts
	r.Users = NewAdminOrgTeamUserService(opts...)
	return
}
