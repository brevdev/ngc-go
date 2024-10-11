// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminService contains methods and other services that help with interacting with
// the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	Options      []option.RequestOption
	Orgs         *AdminOrgService
	Users        *AdminUserService
	Org          *AdminOrgService
	Entitlements *AdminEntitlementService
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r *AdminService) {
	r = &AdminService{}
	r.Options = opts
	r.Orgs = NewAdminOrgService(opts...)
	r.Users = NewAdminUserService(opts...)
	r.Org = NewAdminOrgService(opts...)
	r.Entitlements = NewAdminEntitlementService(opts...)
	return
}

// Backfill Orgs to Kratos
func (r *AdminService) BackfillOrgsToKratos(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/backfill-orgs-to-kratos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
