// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/brevdev/ngc-go/internal/apiquery"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// AdminOrgValidateService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgValidateService] method instead.
type AdminOrgValidateService struct {
	Options []option.RequestOption
}

// NewAdminOrgValidateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgValidateService(opts ...option.RequestOption) (r *AdminOrgValidateService) {
	r = &AdminOrgValidateService{}
	r.Options = opts
	return
}

// List all organizations that match the validate org params
func (r *AdminOrgValidateService) List(ctx context.Context, query AdminOrgValidateListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/org/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgValidateListParams struct {
	// Validate Organization Parameters
	Q param.Field[AdminOrgValidateListParamsQ] `query:"q,required"`
}

// URLQuery serializes [AdminOrgValidateListParams]'s query parameters as
// `url.Values`.
func (r AdminOrgValidateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Validate Organization Parameters
type AdminOrgValidateListParamsQ struct {
	// Org owner.
	OrgOwner param.Field[AdminOrgValidateListParamsQOrgOwner] `query:"orgOwner,required"`
	// Product end customer salesforce.com id (external customer id) for enterprise
	// product.
	PecSfdcID param.Field[string] `query:"pecSfdcId,required"`
	// Product Subscriptions.
	ProductSubscriptions param.Field[[]AdminOrgValidateListParamsQProductSubscription] `query:"productSubscriptions,required"`
}

// URLQuery serializes [AdminOrgValidateListParamsQ]'s query parameters as
// `url.Values`.
func (r AdminOrgValidateListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Org owner.
type AdminOrgValidateListParamsQOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `query:"email,required"`
	// Org owner name.
	FullName param.Field[string] `query:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `query:"lastLoginDate"`
}

// URLQuery serializes [AdminOrgValidateListParamsQOrgOwner]'s query parameters as
// `url.Values`.
func (r AdminOrgValidateListParamsQOrgOwner) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Product Subscription
type AdminOrgValidateListParamsQProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `query:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `query:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType] `query:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `query:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `query:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgValidateListParamsQProductSubscriptionsType] `query:"type"`
}

// URLQuery serializes [AdminOrgValidateListParamsQProductSubscription]'s query
// parameters as `url.Values`.
func (r AdminOrgValidateListParamsQProductSubscription) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgValidateListParamsQProductSubscriptionsType string

const (
	AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminEval       AdminOrgValidateListParamsQProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminNfr        AdminOrgValidateListParamsQProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminCommercial AdminOrgValidateListParamsQProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgValidateListParamsQProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminEval, AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminNfr, AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}
