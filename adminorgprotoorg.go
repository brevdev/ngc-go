// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// AdminOrgProtoOrgService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgProtoOrgService] method instead.
type AdminOrgProtoOrgService struct {
	Options []option.RequestOption
}

// NewAdminOrgProtoOrgService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgProtoOrgService(opts ...option.RequestOption) (r *AdminOrgProtoOrgService) {
	r = &AdminOrgProtoOrgService{}
	r.Options = opts
	return
}

// OrgCreateRequest is used to create the organization or when no nca_id is
// provided upfront, the OrgCreateRequest is stored as proto org, and proto org
// flow initiates (SuperAdmin privileges required)
func (r *AdminOrgProtoOrgService) New(ctx context.Context, params AdminOrgProtoOrgNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/admin/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AdminOrgProtoOrgNewParams struct {
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
	OrgOwner param.Field[AdminOrgProtoOrgNewParamsOrgOwner] `json:"orgOwner"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                                       `json:"pecSfdcId"`
	ProductEnablements param.Field[[]AdminOrgProtoOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]AdminOrgProtoOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Proto org identifier
	ProtoOrgID param.Field[string] `json:"protoOrgId"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                          `json:"sendEmail"`
	Type      param.Field[AdminOrgProtoOrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]                        `cookie:"ncid"`
	VisitorID param.Field[string]                        `cookie:"VisitorID"`
}

func (r AdminOrgProtoOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type AdminOrgProtoOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName"`
	// Identity Provider ID of the org owner.
	IdpID param.Field[string] `json:"idpId"`
	// Starfleet ID of the org owner.
	StarfleetID param.Field[string] `json:"starfleetId"`
}

func (r AdminOrgProtoOrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgProtoOrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgProtoOrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                                `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgProtoOrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r AdminOrgProtoOrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgProtoOrgNewParamsProductEnablementsType string

const (
	AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval       AdminOrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminNfr        AdminOrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminCommercial AdminOrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsEval            AdminOrgProtoOrgNewParamsProductEnablementsType = "EMS_EVAL"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsNfr             AdminOrgProtoOrgNewParamsProductEnablementsType = "EMS_NFR"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsCommercial      AdminOrgProtoOrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminDeveloper  AdminOrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgProtoOrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval, AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminNfr, AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminCommercial, AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsEval, AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsNfr, AdminOrgProtoOrgNewParamsProductEnablementsTypeEmsCommercial, AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgProtoOrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgProtoOrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type AdminOrgProtoOrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgProtoOrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r AdminOrgProtoOrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgProtoOrgNewParamsProductSubscriptionsType string

const (
	AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval       AdminOrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminNfr        AdminOrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial AdminOrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgProtoOrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval, AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminNfr, AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type AdminOrgProtoOrgNewParamsType string

const (
	AdminOrgProtoOrgNewParamsTypeUnknown    AdminOrgProtoOrgNewParamsType = "UNKNOWN"
	AdminOrgProtoOrgNewParamsTypeCloud      AdminOrgProtoOrgNewParamsType = "CLOUD"
	AdminOrgProtoOrgNewParamsTypeEnterprise AdminOrgProtoOrgNewParamsType = "ENTERPRISE"
	AdminOrgProtoOrgNewParamsTypeIndividual AdminOrgProtoOrgNewParamsType = "INDIVIDUAL"
)

func (r AdminOrgProtoOrgNewParamsType) IsKnown() bool {
	switch r {
	case AdminOrgProtoOrgNewParamsTypeUnknown, AdminOrgProtoOrgNewParamsTypeCloud, AdminOrgProtoOrgNewParamsTypeEnterprise, AdminOrgProtoOrgNewParamsTypeIndividual:
		return true
	}
	return false
}
