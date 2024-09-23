// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/brevdev/ngc-go/option"
)

// UserV2Service contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV2Service] method instead.
type UserV2Service struct {
	Options []option.RequestOption
	APIKey  *UserV2APIKeyService
}

// NewUserV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserV2Service(opts ...option.RequestOption) (r *UserV2Service) {
	r = &UserV2Service{}
	r.Options = opts
	r.APIKey = NewUserV2APIKeyService(opts...)
	return
}
