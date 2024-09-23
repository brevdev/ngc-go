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

// SuperAdminOrgService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminOrgService] method instead.
type SuperAdminOrgService struct {
	Options   []option.RequestOption
	OrgStatus *SuperAdminOrgOrgStatusService
}

// NewSuperAdminOrgService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperAdminOrgService(opts ...option.RequestOption) (r *SuperAdminOrgService) {
	r = &SuperAdminOrgService{}
	r.Options = opts
	r.OrgStatus = NewSuperAdminOrgOrgStatusService(opts...)
	return
}

// Create a new organization. (SuperAdmin privileges required)
func (r *SuperAdminOrgService) New(ctx context.Context, params SuperAdminOrgNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SuperAdminOrgNewParams struct {
	// Org owner.
	OrgOwner param.Field[SuperAdminOrgNewParamsOrgOwner] `json:"orgOwner,required"`
	// user country
	Country param.Field[string] `json:"country"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identity Provider ID.
	IdpID param.Field[string] `json:"idpId"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Organization name
	Name param.Field[string] `json:"name"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                                    `json:"pecSfdcId"`
	ProductEnablements param.Field[[]SuperAdminOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]SuperAdminOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                       `json:"sendEmail"`
	Type      param.Field[SuperAdminOrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]                     `cookie:"ncid"`
	VisitorID param.Field[string]                     `cookie:"VisitorID"`
}

func (r SuperAdminOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type SuperAdminOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `json:"lastLoginDate"`
}

func (r SuperAdminOrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type SuperAdminOrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[SuperAdminOrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                             `json:"expirationDate"`
	PoDetails      param.Field[[]SuperAdminOrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r SuperAdminOrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type SuperAdminOrgNewParamsProductEnablementsType string

const (
	SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminEval       SuperAdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminNfr        SuperAdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial SuperAdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	SuperAdminOrgNewParamsProductEnablementsTypeEmsEval            SuperAdminOrgNewParamsProductEnablementsType = "EMS_EVAL"
	SuperAdminOrgNewParamsProductEnablementsTypeEmsNfr             SuperAdminOrgNewParamsProductEnablementsType = "EMS_NFR"
	SuperAdminOrgNewParamsProductEnablementsTypeEmsCommercial      SuperAdminOrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper  SuperAdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r SuperAdminOrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminEval, SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminNfr, SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial, SuperAdminOrgNewParamsProductEnablementsTypeEmsEval, SuperAdminOrgNewParamsProductEnablementsTypeEmsNfr, SuperAdminOrgNewParamsProductEnablementsTypeEmsCommercial, SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type SuperAdminOrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r SuperAdminOrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type SuperAdminOrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[SuperAdminOrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r SuperAdminOrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type SuperAdminOrgNewParamsProductSubscriptionsType string

const (
	SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval       SuperAdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr        SuperAdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial SuperAdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r SuperAdminOrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval, SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr, SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type SuperAdminOrgNewParamsType string

const (
	SuperAdminOrgNewParamsTypeUnknown    SuperAdminOrgNewParamsType = "UNKNOWN"
	SuperAdminOrgNewParamsTypeCloud      SuperAdminOrgNewParamsType = "CLOUD"
	SuperAdminOrgNewParamsTypeEnterprise SuperAdminOrgNewParamsType = "ENTERPRISE"
	SuperAdminOrgNewParamsTypeIndividual SuperAdminOrgNewParamsType = "INDIVIDUAL"
)

func (r SuperAdminOrgNewParamsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgNewParamsTypeUnknown, SuperAdminOrgNewParamsTypeCloud, SuperAdminOrgNewParamsTypeEnterprise, SuperAdminOrgNewParamsTypeIndividual:
		return true
	}
	return false
}
