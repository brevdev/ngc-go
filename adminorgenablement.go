// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminOrgEnablementService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgEnablementService] method instead.
type AdminOrgEnablementService struct {
	Options []option.RequestOption
}

// NewAdminOrgEnablementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgEnablementService(opts ...option.RequestOption) (r *AdminOrgEnablementService) {
	r = &AdminOrgEnablementService{}
	r.Options = opts
	return
}

// Create org product enablement
func (r *AdminOrgEnablementService) New(ctx context.Context, orgName string, body AdminOrgEnablementNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/enablement", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AdminOrgEnablementNewParams struct {
	// False only if called by SbMS.
	CreateSubscription param.Field[bool] `json:"createSubscription,required"`
	// Product Enablement
	ProductEnablement param.Field[AdminOrgEnablementNewParamsProductEnablement] `json:"productEnablement,required"`
}

func (r AdminOrgEnablementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgEnablementNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgEnablementNewParamsProductEnablementType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                                 `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgEnablementNewParamsProductEnablementPoDetail] `json:"poDetails"`
}

func (r AdminOrgEnablementNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgEnablementNewParamsProductEnablementType string

const (
	AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminEval       AdminOrgEnablementNewParamsProductEnablementType = "NGC_ADMIN_EVAL"
	AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminNfr        AdminOrgEnablementNewParamsProductEnablementType = "NGC_ADMIN_NFR"
	AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminCommercial AdminOrgEnablementNewParamsProductEnablementType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgEnablementNewParamsProductEnablementTypeEmsEval            AdminOrgEnablementNewParamsProductEnablementType = "EMS_EVAL"
	AdminOrgEnablementNewParamsProductEnablementTypeEmsNfr             AdminOrgEnablementNewParamsProductEnablementType = "EMS_NFR"
	AdminOrgEnablementNewParamsProductEnablementTypeEmsCommercial      AdminOrgEnablementNewParamsProductEnablementType = "EMS_COMMERCIAL"
	AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminDeveloper  AdminOrgEnablementNewParamsProductEnablementType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgEnablementNewParamsProductEnablementType) IsKnown() bool {
	switch r {
	case AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminEval, AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminNfr, AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminCommercial, AdminOrgEnablementNewParamsProductEnablementTypeEmsEval, AdminOrgEnablementNewParamsProductEnablementTypeEmsNfr, AdminOrgEnablementNewParamsProductEnablementTypeEmsCommercial, AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgEnablementNewParamsProductEnablementPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgEnablementNewParamsProductEnablementPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
