// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// UserAPIKeyService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAPIKeyService] method instead.
type UserAPIKeyService struct {
	Options []option.RequestOption
}

// NewUserAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAPIKeyService(opts ...option.RequestOption) (r *UserAPIKeyService) {
	r = &UserAPIKeyService{}
	r.Options = opts
	return
}

// Generate API Key
func (r *UserAPIKeyService) Generate(ctx context.Context, opts ...option.RequestOption) (res *UserKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me/api-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
