// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// OrgProtoOrgService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgProtoOrgService] method instead.
type OrgProtoOrgService struct {
	Options []option.RequestOption
}

// NewOrgProtoOrgService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgProtoOrgService(opts ...option.RequestOption) (r *OrgProtoOrgService) {
	r = &OrgProtoOrgService{}
	r.Options = opts
	return
}

// Create a new organization based on the org info retrieved from the ProtoOrg.
func (r *OrgProtoOrgService) New(ctx context.Context, params OrgProtoOrgNewParams, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/orgs/proto-org"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Validate org creation from proto org
func (r *OrgProtoOrgService) Validate(ctx context.Context, query OrgProtoOrgValidateParams, opts ...option.RequestOption) (res *OrgInvitation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/orgs/proto-org/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Invitation Validation Response.
type OrgInvitation struct {
	// Org invitation to NGC
	OrgInvitation OrgInvitationOrgInvitation `json:"orgInvitation"`
	RequestStatus OrgInvitationRequestStatus `json:"requestStatus"`
	JSON          orgInvitationJSON          `json:"-"`
}

// orgInvitationJSON contains the JSON metadata for the struct [OrgInvitation]
type orgInvitationJSON struct {
	OrgInvitation apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgInvitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationJSON) RawJSON() string {
	return r.raw
}

// Org invitation to NGC
type OrgInvitationOrgInvitation struct {
	// Email address of the user.
	Email string `json:"email"`
	// Proto Org identifier.
	ProtoOrgID string                         `json:"protoOrgId"`
	JSON       orgInvitationOrgInvitationJSON `json:"-"`
}

// orgInvitationOrgInvitationJSON contains the JSON metadata for the struct
// [OrgInvitationOrgInvitation]
type orgInvitationOrgInvitationJSON struct {
	Email       apijson.Field
	ProtoOrgID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgInvitationOrgInvitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationOrgInvitationJSON) RawJSON() string {
	return r.raw
}

type OrgInvitationRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgInvitationRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                               `json:"statusDescription"`
	JSON              orgInvitationRequestStatusJSON       `json:"-"`
}

// orgInvitationRequestStatusJSON contains the JSON metadata for the struct
// [OrgInvitationRequestStatus]
type orgInvitationRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgInvitationRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgInvitationRequestStatusStatusCode string

const (
	OrgInvitationRequestStatusStatusCodeUnknown                    OrgInvitationRequestStatusStatusCode = "UNKNOWN"
	OrgInvitationRequestStatusStatusCodeSuccess                    OrgInvitationRequestStatusStatusCode = "SUCCESS"
	OrgInvitationRequestStatusStatusCodeUnauthorized               OrgInvitationRequestStatusStatusCode = "UNAUTHORIZED"
	OrgInvitationRequestStatusStatusCodePaymentRequired            OrgInvitationRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgInvitationRequestStatusStatusCodeForbidden                  OrgInvitationRequestStatusStatusCode = "FORBIDDEN"
	OrgInvitationRequestStatusStatusCodeTimeout                    OrgInvitationRequestStatusStatusCode = "TIMEOUT"
	OrgInvitationRequestStatusStatusCodeExists                     OrgInvitationRequestStatusStatusCode = "EXISTS"
	OrgInvitationRequestStatusStatusCodeNotFound                   OrgInvitationRequestStatusStatusCode = "NOT_FOUND"
	OrgInvitationRequestStatusStatusCodeInternalError              OrgInvitationRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgInvitationRequestStatusStatusCodeInvalidRequest             OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST"
	OrgInvitationRequestStatusStatusCodeInvalidRequestVersion      OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgInvitationRequestStatusStatusCodeInvalidRequestData         OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgInvitationRequestStatusStatusCodeMethodNotAllowed           OrgInvitationRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgInvitationRequestStatusStatusCodeConflict                   OrgInvitationRequestStatusStatusCode = "CONFLICT"
	OrgInvitationRequestStatusStatusCodeUnprocessableEntity        OrgInvitationRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgInvitationRequestStatusStatusCodeTooManyRequests            OrgInvitationRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgInvitationRequestStatusStatusCodeInsufficientStorage        OrgInvitationRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgInvitationRequestStatusStatusCodeServiceUnavailable         OrgInvitationRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgInvitationRequestStatusStatusCodePayloadTooLarge            OrgInvitationRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgInvitationRequestStatusStatusCodeNotAcceptable              OrgInvitationRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgInvitationRequestStatusStatusCodeUnavailableForLegalReasons OrgInvitationRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgInvitationRequestStatusStatusCodeBadGateway                 OrgInvitationRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgInvitationRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgInvitationRequestStatusStatusCodeUnknown, OrgInvitationRequestStatusStatusCodeSuccess, OrgInvitationRequestStatusStatusCodeUnauthorized, OrgInvitationRequestStatusStatusCodePaymentRequired, OrgInvitationRequestStatusStatusCodeForbidden, OrgInvitationRequestStatusStatusCodeTimeout, OrgInvitationRequestStatusStatusCodeExists, OrgInvitationRequestStatusStatusCodeNotFound, OrgInvitationRequestStatusStatusCodeInternalError, OrgInvitationRequestStatusStatusCodeInvalidRequest, OrgInvitationRequestStatusStatusCodeInvalidRequestVersion, OrgInvitationRequestStatusStatusCodeInvalidRequestData, OrgInvitationRequestStatusStatusCodeMethodNotAllowed, OrgInvitationRequestStatusStatusCodeConflict, OrgInvitationRequestStatusStatusCodeUnprocessableEntity, OrgInvitationRequestStatusStatusCodeTooManyRequests, OrgInvitationRequestStatusStatusCodeInsufficientStorage, OrgInvitationRequestStatusStatusCodeServiceUnavailable, OrgInvitationRequestStatusStatusCodePayloadTooLarge, OrgInvitationRequestStatusStatusCodeNotAcceptable, OrgInvitationRequestStatusStatusCodeUnavailableForLegalReasons, OrgInvitationRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgProtoOrgNewParams struct {
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
	OrgOwner param.Field[OrgProtoOrgNewParamsOrgOwner] `json:"orgOwner"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                                  `json:"pecSfdcId"`
	ProductEnablements param.Field[[]OrgProtoOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]OrgProtoOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Proto org identifier
	ProtoOrgID param.Field[string] `json:"protoOrgId"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                     `json:"sendEmail"`
	Type      param.Field[OrgProtoOrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]                   `cookie:"ncid"`
	VisitorID param.Field[string]                   `cookie:"VisitorID"`
}

func (r OrgProtoOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type OrgProtoOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName"`
	// Identity Provider ID of the org owner.
	IdpID param.Field[string] `json:"idpId"`
	// Starfleet ID of the org owner.
	StarfleetID param.Field[string] `json:"starfleetId"`
}

func (r OrgProtoOrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type OrgProtoOrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[OrgProtoOrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                           `json:"expirationDate"`
	PoDetails      param.Field[[]OrgProtoOrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r OrgProtoOrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type OrgProtoOrgNewParamsProductEnablementsType string

const (
	OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval       OrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminNfr        OrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminCommercial OrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgProtoOrgNewParamsProductEnablementsTypeEmsEval            OrgProtoOrgNewParamsProductEnablementsType = "EMS_EVAL"
	OrgProtoOrgNewParamsProductEnablementsTypeEmsNfr             OrgProtoOrgNewParamsProductEnablementsType = "EMS_NFR"
	OrgProtoOrgNewParamsProductEnablementsTypeEmsCommercial      OrgProtoOrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminDeveloper  OrgProtoOrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgProtoOrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval, OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminNfr, OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminCommercial, OrgProtoOrgNewParamsProductEnablementsTypeEmsEval, OrgProtoOrgNewParamsProductEnablementsTypeEmsNfr, OrgProtoOrgNewParamsProductEnablementsTypeEmsCommercial, OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgProtoOrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r OrgProtoOrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type OrgProtoOrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[OrgProtoOrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r OrgProtoOrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgProtoOrgNewParamsProductSubscriptionsType string

const (
	OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval       OrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminNfr        OrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial OrgProtoOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgProtoOrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval, OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminNfr, OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type OrgProtoOrgNewParamsType string

const (
	OrgProtoOrgNewParamsTypeUnknown    OrgProtoOrgNewParamsType = "UNKNOWN"
	OrgProtoOrgNewParamsTypeCloud      OrgProtoOrgNewParamsType = "CLOUD"
	OrgProtoOrgNewParamsTypeEnterprise OrgProtoOrgNewParamsType = "ENTERPRISE"
	OrgProtoOrgNewParamsTypeIndividual OrgProtoOrgNewParamsType = "INDIVIDUAL"
)

func (r OrgProtoOrgNewParamsType) IsKnown() bool {
	switch r {
	case OrgProtoOrgNewParamsTypeUnknown, OrgProtoOrgNewParamsTypeCloud, OrgProtoOrgNewParamsTypeEnterprise, OrgProtoOrgNewParamsTypeIndividual:
		return true
	}
	return false
}

type OrgProtoOrgValidateParams struct {
	// JWT that contains org owner email and proto org identifier
	InvitationToken param.Field[string] `query:"invitation_token,required"`
}

// URLQuery serializes [OrgProtoOrgValidateParams]'s query parameters as
// `url.Values`.
func (r OrgProtoOrgValidateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
