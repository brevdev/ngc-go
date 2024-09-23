// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/option"
)

// UsersManagementService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsersManagementService] method instead.
type UsersManagementService struct {
	Options []option.RequestOption
	Me      *UsersManagementMeService
}

// NewUsersManagementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsersManagementService(opts ...option.RequestOption) (r *UsersManagementService) {
	r = &UsersManagementService{}
	r.Options = opts
	r.Me = NewUsersManagementMeService(opts...)
	return
}
