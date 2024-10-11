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

// AdminOrgOrganizationService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgOrganizationService] method instead.
type AdminOrgOrganizationService struct {
	Options []option.RequestOption
}

// NewAdminOrgOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgOrganizationService(opts ...option.RequestOption) (r *AdminOrgOrganizationService) {
	r = &AdminOrgOrganizationService{}
	r.Options = opts
	return
}

// List all organizations that match the validate org params
func (r *AdminOrgOrganizationService) Validate(ctx context.Context, query AdminOrgOrganizationValidateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/org/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgOrganizationValidateParams struct {
	// Validate Organization Parameters
	Q param.Field[AdminOrgOrganizationValidateParamsQ] `query:"q,required"`
}

// URLQuery serializes [AdminOrgOrganizationValidateParams]'s query parameters as
// `url.Values`.
func (r AdminOrgOrganizationValidateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Validate Organization Parameters
type AdminOrgOrganizationValidateParamsQ struct {
	// Org owner.
	OrgOwner param.Field[AdminOrgOrganizationValidateParamsQOrgOwner] `query:"orgOwner,required"`
	// Product end customer salesforce.com id (external customer id) for enterprise
	// product.
	PecSfdcID param.Field[string] `query:"pecSfdcId,required"`
	// Product Subscriptions.
	ProductSubscriptions param.Field[[]AdminOrgOrganizationValidateParamsQProductSubscription] `query:"productSubscriptions,required"`
}

// URLQuery serializes [AdminOrgOrganizationValidateParamsQ]'s query parameters as
// `url.Values`.
func (r AdminOrgOrganizationValidateParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Org owner.
type AdminOrgOrganizationValidateParamsQOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `query:"email,required"`
	// Org owner name.
	FullName param.Field[string] `query:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `query:"lastLoginDate"`
}

// URLQuery serializes [AdminOrgOrganizationValidateParamsQOrgOwner]'s query
// parameters as `url.Values`.
func (r AdminOrgOrganizationValidateParamsQOrgOwner) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Product Subscription
type AdminOrgOrganizationValidateParamsQProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `query:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `query:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType] `query:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `query:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `query:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgOrganizationValidateParamsQProductSubscriptionsType] `query:"type"`
}

// URLQuery serializes [AdminOrgOrganizationValidateParamsQProductSubscription]'s
// query parameters as `url.Values`.
func (r AdminOrgOrganizationValidateParamsQProductSubscription) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgOrganizationValidateParamsQProductSubscriptionsType string

const (
	AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminEval       AdminOrgOrganizationValidateParamsQProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminNfr        AdminOrgOrganizationValidateParamsQProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminCommercial AdminOrgOrganizationValidateParamsQProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgOrganizationValidateParamsQProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminEval, AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminNfr, AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}
