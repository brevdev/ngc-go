// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminEntitlementService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminEntitlementService] method instead.
type AdminEntitlementService struct {
	Options []option.RequestOption
}

// NewAdminEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminEntitlementService(opts ...option.RequestOption) (r *AdminEntitlementService) {
	r = &AdminEntitlementService{}
	r.Options = opts
	return
}

// List all organizations with entitlements. (SuperAdmin, ECM and Billing
// privileges required)
func (r *AdminEntitlementService) List(ctx context.Context, query AdminEntitlementListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminEntitlementListParams struct {
	// get is paid subscription entitlements
	IsPaidSubscription param.Field[bool] `query:"is-paid-subscription"`
	// filter by product-name
	ProductName param.Field[string] `query:"product-name"`
}

// URLQuery serializes [AdminEntitlementListParams]'s query parameters as
// `url.Values`.
func (r AdminEntitlementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
