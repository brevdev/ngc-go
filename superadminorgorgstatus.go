// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// SuperAdminOrgOrgStatusService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminOrgOrgStatusService] method instead.
type SuperAdminOrgOrgStatusService struct {
	Options []option.RequestOption
}

// NewSuperAdminOrgOrgStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperAdminOrgOrgStatusService(opts ...option.RequestOption) (r *SuperAdminOrgOrgStatusService) {
	r = &SuperAdminOrgOrgStatusService{}
	r.Options = opts
	return
}

// Create org product enablement
func (r *SuperAdminOrgOrgStatusService) New(ctx context.Context, orgName string, body SuperAdminOrgOrgStatusNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
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

type SuperAdminOrgOrgStatusNewParams struct {
	// False only if called by SbMS.
	CreateSubscription param.Field[bool] `json:"createSubscription,required"`
	// Product Enablement
	ProductEnablement param.Field[SuperAdminOrgOrgStatusNewParamsProductEnablement] `json:"productEnablement,required"`
}

func (r SuperAdminOrgOrgStatusNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type SuperAdminOrgOrgStatusNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[SuperAdminOrgOrgStatusNewParamsProductEnablementType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                                     `json:"expirationDate"`
	PoDetails      param.Field[[]SuperAdminOrgOrgStatusNewParamsProductEnablementPoDetail] `json:"poDetails"`
}

func (r SuperAdminOrgOrgStatusNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type SuperAdminOrgOrgStatusNewParamsProductEnablementType string

const (
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminEval       SuperAdminOrgOrgStatusNewParamsProductEnablementType = "NGC_ADMIN_EVAL"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminNfr        SuperAdminOrgOrgStatusNewParamsProductEnablementType = "NGC_ADMIN_NFR"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminCommercial SuperAdminOrgOrgStatusNewParamsProductEnablementType = "NGC_ADMIN_COMMERCIAL"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsEval            SuperAdminOrgOrgStatusNewParamsProductEnablementType = "EMS_EVAL"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsNfr             SuperAdminOrgOrgStatusNewParamsProductEnablementType = "EMS_NFR"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsCommercial      SuperAdminOrgOrgStatusNewParamsProductEnablementType = "EMS_COMMERCIAL"
	SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminDeveloper  SuperAdminOrgOrgStatusNewParamsProductEnablementType = "NGC_ADMIN_DEVELOPER"
)

func (r SuperAdminOrgOrgStatusNewParamsProductEnablementType) IsKnown() bool {
	switch r {
	case SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminEval, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminNfr, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminCommercial, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsEval, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsNfr, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeEmsCommercial, SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type SuperAdminOrgOrgStatusNewParamsProductEnablementPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r SuperAdminOrgOrgStatusNewParamsProductEnablementPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
