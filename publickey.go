// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// PublicKeyService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPublicKeyService] method instead.
type PublicKeyService struct {
	Options []option.RequestOption
}

// NewPublicKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPublicKeyService(opts ...option.RequestOption) (r *PublicKeyService) {
	r = &PublicKeyService{}
	r.Options = opts
	return
}

// Used to check the health status of the web service only
func (r *PublicKeyService) Health(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "public-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
