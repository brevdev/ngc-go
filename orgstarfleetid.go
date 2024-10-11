// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

// OrgStarfleetIDService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgStarfleetIDService] method instead.
type OrgStarfleetIDService struct {
	Options []option.RequestOption
}

// NewOrgStarfleetIDService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgStarfleetIDService(opts ...option.RequestOption) (r *OrgStarfleetIDService) {
	r = &OrgStarfleetIDService{}
	r.Options = opts
	return
}

// Get User details in org by starfleet Id
func (r *OrgStarfleetIDService) Get(ctx context.Context, orgName string, starfleetID string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if starfleetID == "" {
		err = errors.New("missing required starfleet-id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/starfleetIds/%s", orgName, starfleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
