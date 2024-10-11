// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminOrgEntitlementService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgEntitlementService] method instead.
type AdminOrgEntitlementService struct {
	Options []option.RequestOption
}

// NewAdminOrgEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgEntitlementService(opts ...option.RequestOption) (r *AdminOrgEntitlementService) {
	r = &AdminOrgEntitlementService{}
	r.Options = opts
	return
}

// List all organizations with entitlements. (SuperAdmin, ECM and Billing
// privileges required)
func (r *AdminOrgEntitlementService) List(ctx context.Context, orgName string, query AdminOrgEntitlementListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/entitlements", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgEntitlementListParams struct {
	// get is paid subscription entitlements
	IsPaidSubscription param.Field[bool] `query:"is-paid-subscription"`
	// filter by product-name
	ProductName param.Field[string] `query:"product-name"`
}

// URLQuery serializes [AdminOrgEntitlementListParams]'s query parameters as
// `url.Values`.
func (r AdminOrgEntitlementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
