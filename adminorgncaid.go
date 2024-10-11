// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// AdminOrgNcaIDService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgNcaIDService] method instead.
type AdminOrgNcaIDService struct {
	Options []option.RequestOption
}

// NewAdminOrgNcaIDService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgNcaIDService(opts ...option.RequestOption) (r *AdminOrgNcaIDService) {
	r = &AdminOrgNcaIDService{}
	r.Options = opts
	return
}

// Get Organization by NCA ID or Org Name. (SuperAdmin privileges required)
func (r *AdminOrgNcaIDService) New(ctx context.Context, body AdminOrgNcaIDNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/admin/orgs/ncaIds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AdminOrgNcaIDNewParams struct {
	// List of nca Ids
	NcaIDs param.Field[[]string] `json:"ncaIds"`
	// List of org names
	OrgNames param.Field[[]string] `json:"orgNames"`
}

func (r AdminOrgNcaIDNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
