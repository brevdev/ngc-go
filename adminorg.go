// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/apiquery"
	"github.com/stainless-sdks/ngc-go/internal/param"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// AdminOrgService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgService] method instead.
type AdminOrgService struct {
	Options    []option.RequestOption
	NcaIDs     *AdminOrgNcaIDService
	Offboarded *AdminOrgOffboardedService
	Teams      *AdminOrgTeamService
	Users      *AdminOrgUserService
}

// NewAdminOrgService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgService(opts ...option.RequestOption) (r *AdminOrgService) {
	r = &AdminOrgService{}
	r.Options = opts
	r.NcaIDs = NewAdminOrgNcaIDService(opts...)
	r.Offboarded = NewAdminOrgOffboardedService(opts...)
	r.Teams = NewAdminOrgTeamService(opts...)
	r.Users = NewAdminOrgUserService(opts...)
	return
}

// OrgCreateRequest is used to create the organization or when no nca_id is
// provided upfront, the OrgCreateRequest is stored as proto org, and proto org
// flow initiates (SuperAdmin privileges required)
func (r *AdminOrgService) New(ctx context.Context, params AdminOrgNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/admin/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all organizations. (SuperAdmin privileges required)
func (r *AdminOrgService) List(ctx context.Context, query AdminOrgListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/org"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgNewParams struct {
	// user country
	Country param.Field[string] `json:"country"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identify the initiator of the org request
	Initiator param.Field[string] `json:"initiator"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Organization name
	Name param.Field[string] `json:"name"`
	// NVIDIA Cloud Account Identifier
	NcaID param.Field[string] `json:"ncaId"`
	// NVIDIA Cloud Account Number
	NcaNumber param.Field[string] `json:"ncaNumber"`
	// Org owner.
	OrgOwner param.Field[AdminOrgNewParamsOrgOwner] `json:"orgOwner"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                               `json:"pecSfdcId"`
	ProductEnablements param.Field[[]AdminOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]AdminOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Proto org identifier
	ProtoOrgID param.Field[string] `json:"protoOrgId"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                  `json:"sendEmail"`
	Type      param.Field[AdminOrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]                `cookie:"ncid"`
	VisitorID param.Field[string]                `cookie:"VisitorID"`
}

func (r AdminOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type AdminOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName"`
	// Identity Provider ID of the org owner.
	IdpID param.Field[string] `json:"idpId"`
	// Starfleet ID of the org owner.
	StarfleetID param.Field[string] `json:"starfleetId"`
}

func (r AdminOrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                        `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r AdminOrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgNewParamsProductEnablementsType string

const (
	AdminOrgNewParamsProductEnablementsTypeNgcAdminEval       AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminNfr        AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgNewParamsProductEnablementsTypeEmsEval            AdminOrgNewParamsProductEnablementsType = "EMS_EVAL"
	AdminOrgNewParamsProductEnablementsTypeEmsNfr             AdminOrgNewParamsProductEnablementsType = "EMS_NFR"
	AdminOrgNewParamsProductEnablementsTypeEmsCommercial      AdminOrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper  AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductEnablementsTypeNgcAdminEval, AdminOrgNewParamsProductEnablementsTypeNgcAdminNfr, AdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial, AdminOrgNewParamsProductEnablementsTypeEmsEval, AdminOrgNewParamsProductEnablementsTypeEmsNfr, AdminOrgNewParamsProductEnablementsTypeEmsCommercial, AdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type AdminOrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r AdminOrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgNewParamsProductSubscriptionsType string

const (
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval       AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr        AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval, AdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr, AdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type AdminOrgNewParamsType string

const (
	AdminOrgNewParamsTypeUnknown    AdminOrgNewParamsType = "UNKNOWN"
	AdminOrgNewParamsTypeCloud      AdminOrgNewParamsType = "CLOUD"
	AdminOrgNewParamsTypeEnterprise AdminOrgNewParamsType = "ENTERPRISE"
	AdminOrgNewParamsTypeIndividual AdminOrgNewParamsType = "INDIVIDUAL"
)

func (r AdminOrgNewParamsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsTypeUnknown, AdminOrgNewParamsTypeCloud, AdminOrgNewParamsTypeEnterprise, AdminOrgNewParamsTypeIndividual:
		return true
	}
	return false
}

type AdminOrgListParams struct {
	FilterUsingOrgDisplayName param.Field[string]                                     `query:"Filter using org display name"`
	FilterUsingOrgOwnerEmail  param.Field[AdminOrgListParamsFilterUsingOrgOwnerEmail] `query:"Filter using org owner email"`
	FilterUsingOrgOwnerName   param.Field[string]                                     `query:"Filter using org owner name"`
	// Org description to search
	OrgDesc param.Field[string] `query:"org-desc"`
	// Org name to search
	OrgName param.Field[string] `query:"org-name"`
	// Org Type to search
	OrgType param.Field[AdminOrgListParamsOrgType] `query:"org-type"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// PEC ID to search
	PecID param.Field[string] `query:"pec-id"`
}

// URLQuery serializes [AdminOrgListParams]'s query parameters as `url.Values`.
func (r AdminOrgListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgListParamsFilterUsingOrgOwnerEmail struct {
	EmailShouldBeBase64Encoded param.Field[string] `query:" Email should be base-64-encoded"`
}

// URLQuery serializes [AdminOrgListParamsFilterUsingOrgOwnerEmail]'s query
// parameters as `url.Values`.
func (r AdminOrgListParamsFilterUsingOrgOwnerEmail) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Org Type to search
type AdminOrgListParamsOrgType string

const (
	AdminOrgListParamsOrgTypeUnknown    AdminOrgListParamsOrgType = "UNKNOWN"
	AdminOrgListParamsOrgTypeCloud      AdminOrgListParamsOrgType = "CLOUD"
	AdminOrgListParamsOrgTypeEnterprise AdminOrgListParamsOrgType = "ENTERPRISE"
	AdminOrgListParamsOrgTypeIndividual AdminOrgListParamsOrgType = "INDIVIDUAL"
)

func (r AdminOrgListParamsOrgType) IsKnown() bool {
	switch r {
	case AdminOrgListParamsOrgTypeUnknown, AdminOrgListParamsOrgTypeCloud, AdminOrgListParamsOrgTypeEnterprise, AdminOrgListParamsOrgTypeIndividual:
		return true
	}
	return false
}
