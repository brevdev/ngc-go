// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// OrgService contains methods and other services that help with interacting with
// the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options   []option.RequestOption
	Users     *OrgUserService
	Teams     *OrgTeamService
	ProtoOrg  *OrgProtoOrgService
	Credits   *OrgCreditService
	AuditLogs *OrgAuditLogService
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r *OrgService) {
	r = &OrgService{}
	r.Options = opts
	r.Users = NewOrgUserService(opts...)
	r.Teams = NewOrgTeamService(opts...)
	r.ProtoOrg = NewOrgProtoOrgService(opts...)
	r.Credits = NewOrgCreditService(opts...)
	r.AuditLogs = NewOrgAuditLogService(opts...)
	return
}

// Create a new organization based on the org info provided in the request.
func (r *OrgService) New(ctx context.Context, params OrgNewParams, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get organization information
func (r *OrgService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/orgs/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update organization information
func (r *OrgService) Update(ctx context.Context, orgName string, body OrgUpdateParams, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/orgs/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all organizations of the user
func (r *OrgService) List(ctx context.Context, query OrgListParams, opts ...option.RequestOption) (res *OrgList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// info about an organizations
type Org struct {
	// Information about the Organization
	Organizations OrgOrganizations `json:"organizations"`
	RequestStatus OrgRequestStatus `json:"requestStatus"`
	JSON          orgJSON          `json:"-"`
}

// orgJSON contains the JSON metadata for the struct [Org]
type orgJSON struct {
	Organizations apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Org) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type OrgOrganizations struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact OrgOrganizationsAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings OrgOrganizationsInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner OrgOrganizationsOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []OrgOrganizationsOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                `json:"pecSfdcId"`
	ProductEnablements   []OrgOrganizationsProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []OrgOrganizationsProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings OrgOrganizationsRepoScanSettings `json:"repoScanSettings"`
	Type             OrgOrganizationsType             `json:"type"`
	// Users information.
	UsersInfo OrgOrganizationsUsersInfo `json:"usersInfo"`
	JSON      orgOrganizationsJSON      `json:"-"`
}

// orgOrganizationsJSON contains the JSON metadata for the struct
// [OrgOrganizations]
type orgOrganizationsJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *OrgOrganizations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type OrgOrganizationsAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                               `json:"fullName"`
	JSON     orgOrganizationsAlternateContactJSON `json:"-"`
}

// orgOrganizationsAlternateContactJSON contains the JSON metadata for the struct
// [OrgOrganizationsAlternateContact]
type orgOrganizationsAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgOrganizationsAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgOrganizationsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                        `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgOrganizationsInfinityManagerSettingsJSON `json:"-"`
}

// orgOrganizationsInfinityManagerSettingsJSON contains the JSON metadata for the
// struct [OrgOrganizationsInfinityManagerSettings]
type orgOrganizationsInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgOrganizationsInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type OrgOrganizationsOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                       `json:"lastLoginDate"`
	JSON          orgOrganizationsOrgOwnerJSON `json:"-"`
}

// orgOrganizationsOrgOwnerJSON contains the JSON metadata for the struct
// [OrgOrganizationsOrgOwner]
type orgOrganizationsOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgOrganizationsOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type OrgOrganizationsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type OrgOrganizationsProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                       `json:"expirationDate"`
	PoDetails      []OrgOrganizationsProductEnablementsPoDetail `json:"poDetails"`
	JSON           orgOrganizationsProductEnablementJSON        `json:"-"`
}

// orgOrganizationsProductEnablementJSON contains the JSON metadata for the struct
// [OrgOrganizationsProductEnablement]
type orgOrganizationsProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrgOrganizationsProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type OrgOrganizationsProductEnablementsType string

const (
	OrgOrganizationsProductEnablementsTypeNgcAdminEval       OrgOrganizationsProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgOrganizationsProductEnablementsTypeNgcAdminNfr        OrgOrganizationsProductEnablementsType = "NGC_ADMIN_NFR"
	OrgOrganizationsProductEnablementsTypeNgcAdminCommercial OrgOrganizationsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgOrganizationsProductEnablementsTypeEmsEval            OrgOrganizationsProductEnablementsType = "EMS_EVAL"
	OrgOrganizationsProductEnablementsTypeEmsNfr             OrgOrganizationsProductEnablementsType = "EMS_NFR"
	OrgOrganizationsProductEnablementsTypeEmsCommercial      OrgOrganizationsProductEnablementsType = "EMS_COMMERCIAL"
	OrgOrganizationsProductEnablementsTypeNgcAdminDeveloper  OrgOrganizationsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgOrganizationsProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgOrganizationsProductEnablementsTypeNgcAdminEval, OrgOrganizationsProductEnablementsTypeNgcAdminNfr, OrgOrganizationsProductEnablementsTypeNgcAdminCommercial, OrgOrganizationsProductEnablementsTypeEmsEval, OrgOrganizationsProductEnablementsTypeEmsNfr, OrgOrganizationsProductEnablementsTypeEmsCommercial, OrgOrganizationsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgOrganizationsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                         `json:"pkId"`
	JSON orgOrganizationsProductEnablementsPoDetailJSON `json:"-"`
}

// orgOrganizationsProductEnablementsPoDetailJSON contains the JSON metadata for
// the struct [OrgOrganizationsProductEnablementsPoDetail]
type orgOrganizationsProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgOrganizationsProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type OrgOrganizationsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType OrgOrganizationsProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type OrgOrganizationsProductSubscriptionsType `json:"type"`
	JSON orgOrganizationsProductSubscriptionJSON  `json:"-"`
}

// orgOrganizationsProductSubscriptionJSON contains the JSON metadata for the
// struct [OrgOrganizationsProductSubscription]
type orgOrganizationsProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrgOrganizationsProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgOrganizationsProductSubscriptionsEmsEntitlementType string

const (
	OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval       OrgOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgOrganizationsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval, OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgOrganizationsProductSubscriptionsType string

const (
	OrgOrganizationsProductSubscriptionsTypeNgcAdminEval       OrgOrganizationsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgOrganizationsProductSubscriptionsTypeNgcAdminNfr        OrgOrganizationsProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgOrganizationsProductSubscriptionsTypeNgcAdminCommercial OrgOrganizationsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgOrganizationsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgOrganizationsProductSubscriptionsTypeNgcAdminEval, OrgOrganizationsProductSubscriptionsTypeNgcAdminNfr, OrgOrganizationsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type OrgOrganizationsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                 `json:"repoScanShowResults"`
	JSON                orgOrganizationsRepoScanSettingsJSON `json:"-"`
}

// orgOrganizationsRepoScanSettingsJSON contains the JSON metadata for the struct
// [OrgOrganizationsRepoScanSettings]
type orgOrganizationsRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgOrganizationsRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrgOrganizationsType string

const (
	OrgOrganizationsTypeUnknown    OrgOrganizationsType = "UNKNOWN"
	OrgOrganizationsTypeCloud      OrgOrganizationsType = "CLOUD"
	OrgOrganizationsTypeEnterprise OrgOrganizationsType = "ENTERPRISE"
	OrgOrganizationsTypeIndividual OrgOrganizationsType = "INDIVIDUAL"
)

func (r OrgOrganizationsType) IsKnown() bool {
	switch r {
	case OrgOrganizationsTypeUnknown, OrgOrganizationsTypeCloud, OrgOrganizationsTypeEnterprise, OrgOrganizationsTypeIndividual:
		return true
	}
	return false
}

// Users information.
type OrgOrganizationsUsersInfo struct {
	// Total number of users.
	TotalUsers int64                         `json:"totalUsers"`
	JSON       orgOrganizationsUsersInfoJSON `json:"-"`
}

// orgOrganizationsUsersInfoJSON contains the JSON metadata for the struct
// [OrgOrganizationsUsersInfo]
type orgOrganizationsUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgOrganizationsUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgOrganizationsUsersInfoJSON) RawJSON() string {
	return r.raw
}

type OrgRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                     `json:"statusDescription"`
	JSON              orgRequestStatusJSON       `json:"-"`
}

// orgRequestStatusJSON contains the JSON metadata for the struct
// [OrgRequestStatus]
type orgRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgRequestStatusStatusCode string

const (
	OrgRequestStatusStatusCodeUnknown                    OrgRequestStatusStatusCode = "UNKNOWN"
	OrgRequestStatusStatusCodeSuccess                    OrgRequestStatusStatusCode = "SUCCESS"
	OrgRequestStatusStatusCodeUnauthorized               OrgRequestStatusStatusCode = "UNAUTHORIZED"
	OrgRequestStatusStatusCodePaymentRequired            OrgRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgRequestStatusStatusCodeForbidden                  OrgRequestStatusStatusCode = "FORBIDDEN"
	OrgRequestStatusStatusCodeTimeout                    OrgRequestStatusStatusCode = "TIMEOUT"
	OrgRequestStatusStatusCodeExists                     OrgRequestStatusStatusCode = "EXISTS"
	OrgRequestStatusStatusCodeNotFound                   OrgRequestStatusStatusCode = "NOT_FOUND"
	OrgRequestStatusStatusCodeInternalError              OrgRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgRequestStatusStatusCodeInvalidRequest             OrgRequestStatusStatusCode = "INVALID_REQUEST"
	OrgRequestStatusStatusCodeInvalidRequestVersion      OrgRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgRequestStatusStatusCodeInvalidRequestData         OrgRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgRequestStatusStatusCodeMethodNotAllowed           OrgRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgRequestStatusStatusCodeConflict                   OrgRequestStatusStatusCode = "CONFLICT"
	OrgRequestStatusStatusCodeUnprocessableEntity        OrgRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgRequestStatusStatusCodeTooManyRequests            OrgRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgRequestStatusStatusCodeInsufficientStorage        OrgRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgRequestStatusStatusCodeServiceUnavailable         OrgRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgRequestStatusStatusCodePayloadTooLarge            OrgRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgRequestStatusStatusCodeNotAcceptable              OrgRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgRequestStatusStatusCodeUnavailableForLegalReasons OrgRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgRequestStatusStatusCodeBadGateway                 OrgRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgRequestStatusStatusCodeUnknown, OrgRequestStatusStatusCodeSuccess, OrgRequestStatusStatusCodeUnauthorized, OrgRequestStatusStatusCodePaymentRequired, OrgRequestStatusStatusCodeForbidden, OrgRequestStatusStatusCodeTimeout, OrgRequestStatusStatusCodeExists, OrgRequestStatusStatusCodeNotFound, OrgRequestStatusStatusCodeInternalError, OrgRequestStatusStatusCodeInvalidRequest, OrgRequestStatusStatusCodeInvalidRequestVersion, OrgRequestStatusStatusCodeInvalidRequestData, OrgRequestStatusStatusCodeMethodNotAllowed, OrgRequestStatusStatusCodeConflict, OrgRequestStatusStatusCodeUnprocessableEntity, OrgRequestStatusStatusCodeTooManyRequests, OrgRequestStatusStatusCodeInsufficientStorage, OrgRequestStatusStatusCodeServiceUnavailable, OrgRequestStatusStatusCodePayloadTooLarge, OrgRequestStatusStatusCodeNotAcceptable, OrgRequestStatusStatusCodeUnavailableForLegalReasons, OrgRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// List of organizations.
type OrgList struct {
	Organizations []OrgListOrganization `json:"organizations"`
	// object that describes the pagination information
	PaginationInfo OrgListPaginationInfo `json:"paginationInfo"`
	RequestStatus  OrgListRequestStatus  `json:"requestStatus"`
	JSON           orgListJSON           `json:"-"`
}

// orgListJSON contains the JSON metadata for the struct [OrgList]
type orgListJSON struct {
	Organizations  apijson.Field
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrgList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type OrgListOrganization struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact OrgListOrganizationsAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings OrgListOrganizationsInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner OrgListOrganizationsOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []OrgListOrganizationsOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                    `json:"pecSfdcId"`
	ProductEnablements   []OrgListOrganizationsProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []OrgListOrganizationsProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings OrgListOrganizationsRepoScanSettings `json:"repoScanSettings"`
	Type             OrgListOrganizationsType             `json:"type"`
	// Users information.
	UsersInfo OrgListOrganizationsUsersInfo `json:"usersInfo"`
	JSON      orgListOrganizationJSON       `json:"-"`
}

// orgListOrganizationJSON contains the JSON metadata for the struct
// [OrgListOrganization]
type orgListOrganizationJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *OrgListOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type OrgListOrganizationsAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                   `json:"fullName"`
	JSON     orgListOrganizationsAlternateContactJSON `json:"-"`
}

// orgListOrganizationsAlternateContactJSON contains the JSON metadata for the
// struct [OrgListOrganizationsAlternateContact]
type orgListOrganizationsAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgListOrganizationsAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgListOrganizationsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                            `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgListOrganizationsInfinityManagerSettingsJSON `json:"-"`
}

// orgListOrganizationsInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [OrgListOrganizationsInfinityManagerSettings]
type orgListOrganizationsInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgListOrganizationsInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type OrgListOrganizationsOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                           `json:"lastLoginDate"`
	JSON          orgListOrganizationsOrgOwnerJSON `json:"-"`
}

// orgListOrganizationsOrgOwnerJSON contains the JSON metadata for the struct
// [OrgListOrganizationsOrgOwner]
type orgListOrganizationsOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgListOrganizationsOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type OrgListOrganizationsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type OrgListOrganizationsProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                           `json:"expirationDate"`
	PoDetails      []OrgListOrganizationsProductEnablementsPoDetail `json:"poDetails"`
	JSON           orgListOrganizationsProductEnablementJSON        `json:"-"`
}

// orgListOrganizationsProductEnablementJSON contains the JSON metadata for the
// struct [OrgListOrganizationsProductEnablement]
type orgListOrganizationsProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrgListOrganizationsProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type OrgListOrganizationsProductEnablementsType string

const (
	OrgListOrganizationsProductEnablementsTypeNgcAdminEval       OrgListOrganizationsProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgListOrganizationsProductEnablementsTypeNgcAdminNfr        OrgListOrganizationsProductEnablementsType = "NGC_ADMIN_NFR"
	OrgListOrganizationsProductEnablementsTypeNgcAdminCommercial OrgListOrganizationsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgListOrganizationsProductEnablementsTypeEmsEval            OrgListOrganizationsProductEnablementsType = "EMS_EVAL"
	OrgListOrganizationsProductEnablementsTypeEmsNfr             OrgListOrganizationsProductEnablementsType = "EMS_NFR"
	OrgListOrganizationsProductEnablementsTypeEmsCommercial      OrgListOrganizationsProductEnablementsType = "EMS_COMMERCIAL"
	OrgListOrganizationsProductEnablementsTypeNgcAdminDeveloper  OrgListOrganizationsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgListOrganizationsProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgListOrganizationsProductEnablementsTypeNgcAdminEval, OrgListOrganizationsProductEnablementsTypeNgcAdminNfr, OrgListOrganizationsProductEnablementsTypeNgcAdminCommercial, OrgListOrganizationsProductEnablementsTypeEmsEval, OrgListOrganizationsProductEnablementsTypeEmsNfr, OrgListOrganizationsProductEnablementsTypeEmsCommercial, OrgListOrganizationsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgListOrganizationsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                             `json:"pkId"`
	JSON orgListOrganizationsProductEnablementsPoDetailJSON `json:"-"`
}

// orgListOrganizationsProductEnablementsPoDetailJSON contains the JSON metadata
// for the struct [OrgListOrganizationsProductEnablementsPoDetail]
type orgListOrganizationsProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgListOrganizationsProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type OrgListOrganizationsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType OrgListOrganizationsProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type OrgListOrganizationsProductSubscriptionsType `json:"type"`
	JSON orgListOrganizationsProductSubscriptionJSON  `json:"-"`
}

// orgListOrganizationsProductSubscriptionJSON contains the JSON metadata for the
// struct [OrgListOrganizationsProductSubscription]
type orgListOrganizationsProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrgListOrganizationsProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgListOrganizationsProductSubscriptionsEmsEntitlementType string

const (
	OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval       OrgListOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgListOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgListOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgListOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgListOrganizationsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval, OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgListOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgListOrganizationsProductSubscriptionsType string

const (
	OrgListOrganizationsProductSubscriptionsTypeNgcAdminEval       OrgListOrganizationsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgListOrganizationsProductSubscriptionsTypeNgcAdminNfr        OrgListOrganizationsProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgListOrganizationsProductSubscriptionsTypeNgcAdminCommercial OrgListOrganizationsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgListOrganizationsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgListOrganizationsProductSubscriptionsTypeNgcAdminEval, OrgListOrganizationsProductSubscriptionsTypeNgcAdminNfr, OrgListOrganizationsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type OrgListOrganizationsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                     `json:"repoScanShowResults"`
	JSON                orgListOrganizationsRepoScanSettingsJSON `json:"-"`
}

// orgListOrganizationsRepoScanSettingsJSON contains the JSON metadata for the
// struct [OrgListOrganizationsRepoScanSettings]
type orgListOrganizationsRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgListOrganizationsRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrgListOrganizationsType string

const (
	OrgListOrganizationsTypeUnknown    OrgListOrganizationsType = "UNKNOWN"
	OrgListOrganizationsTypeCloud      OrgListOrganizationsType = "CLOUD"
	OrgListOrganizationsTypeEnterprise OrgListOrganizationsType = "ENTERPRISE"
	OrgListOrganizationsTypeIndividual OrgListOrganizationsType = "INDIVIDUAL"
)

func (r OrgListOrganizationsType) IsKnown() bool {
	switch r {
	case OrgListOrganizationsTypeUnknown, OrgListOrganizationsTypeCloud, OrgListOrganizationsTypeEnterprise, OrgListOrganizationsTypeIndividual:
		return true
	}
	return false
}

// Users information.
type OrgListOrganizationsUsersInfo struct {
	// Total number of users.
	TotalUsers int64                             `json:"totalUsers"`
	JSON       orgListOrganizationsUsersInfoJSON `json:"-"`
}

// orgListOrganizationsUsersInfoJSON contains the JSON metadata for the struct
// [OrgListOrganizationsUsersInfo]
type orgListOrganizationsUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgListOrganizationsUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListOrganizationsUsersInfoJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type OrgListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                     `json:"totalResults"`
	JSON         orgListPaginationInfoJSON `json:"-"`
}

// orgListPaginationInfoJSON contains the JSON metadata for the struct
// [OrgListPaginationInfo]
type orgListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrgListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type OrgListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                         `json:"statusDescription"`
	JSON              orgListRequestStatusJSON       `json:"-"`
}

// orgListRequestStatusJSON contains the JSON metadata for the struct
// [OrgListRequestStatus]
type orgListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgListRequestStatusStatusCode string

const (
	OrgListRequestStatusStatusCodeUnknown                    OrgListRequestStatusStatusCode = "UNKNOWN"
	OrgListRequestStatusStatusCodeSuccess                    OrgListRequestStatusStatusCode = "SUCCESS"
	OrgListRequestStatusStatusCodeUnauthorized               OrgListRequestStatusStatusCode = "UNAUTHORIZED"
	OrgListRequestStatusStatusCodePaymentRequired            OrgListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgListRequestStatusStatusCodeForbidden                  OrgListRequestStatusStatusCode = "FORBIDDEN"
	OrgListRequestStatusStatusCodeTimeout                    OrgListRequestStatusStatusCode = "TIMEOUT"
	OrgListRequestStatusStatusCodeExists                     OrgListRequestStatusStatusCode = "EXISTS"
	OrgListRequestStatusStatusCodeNotFound                   OrgListRequestStatusStatusCode = "NOT_FOUND"
	OrgListRequestStatusStatusCodeInternalError              OrgListRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgListRequestStatusStatusCodeInvalidRequest             OrgListRequestStatusStatusCode = "INVALID_REQUEST"
	OrgListRequestStatusStatusCodeInvalidRequestVersion      OrgListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgListRequestStatusStatusCodeInvalidRequestData         OrgListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgListRequestStatusStatusCodeMethodNotAllowed           OrgListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgListRequestStatusStatusCodeConflict                   OrgListRequestStatusStatusCode = "CONFLICT"
	OrgListRequestStatusStatusCodeUnprocessableEntity        OrgListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgListRequestStatusStatusCodeTooManyRequests            OrgListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgListRequestStatusStatusCodeInsufficientStorage        OrgListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgListRequestStatusStatusCodeServiceUnavailable         OrgListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgListRequestStatusStatusCodePayloadTooLarge            OrgListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgListRequestStatusStatusCodeNotAcceptable              OrgListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgListRequestStatusStatusCodeUnavailableForLegalReasons OrgListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgListRequestStatusStatusCodeBadGateway                 OrgListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgListRequestStatusStatusCodeUnknown, OrgListRequestStatusStatusCodeSuccess, OrgListRequestStatusStatusCodeUnauthorized, OrgListRequestStatusStatusCodePaymentRequired, OrgListRequestStatusStatusCodeForbidden, OrgListRequestStatusStatusCodeTimeout, OrgListRequestStatusStatusCodeExists, OrgListRequestStatusStatusCodeNotFound, OrgListRequestStatusStatusCodeInternalError, OrgListRequestStatusStatusCodeInvalidRequest, OrgListRequestStatusStatusCodeInvalidRequestVersion, OrgListRequestStatusStatusCodeInvalidRequestData, OrgListRequestStatusStatusCodeMethodNotAllowed, OrgListRequestStatusStatusCodeConflict, OrgListRequestStatusStatusCodeUnprocessableEntity, OrgListRequestStatusStatusCodeTooManyRequests, OrgListRequestStatusStatusCodeInsufficientStorage, OrgListRequestStatusStatusCodeServiceUnavailable, OrgListRequestStatusStatusCodePayloadTooLarge, OrgListRequestStatusStatusCodeNotAcceptable, OrgListRequestStatusStatusCodeUnavailableForLegalReasons, OrgListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// details about one team
type Team struct {
	RequestStatus TeamRequestStatus `json:"requestStatus"`
	// Information about the team
	Team TeamTeam `json:"team"`
	JSON teamJSON `json:"-"`
}

// teamJSON contains the JSON metadata for the struct [Team]
type teamJSON struct {
	RequestStatus apijson.Field
	Team          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Team) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamJSON) RawJSON() string {
	return r.raw
}

type TeamRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                      `json:"statusDescription"`
	JSON              teamRequestStatusJSON       `json:"-"`
}

// teamRequestStatusJSON contains the JSON metadata for the struct
// [TeamRequestStatus]
type teamRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamRequestStatusStatusCode string

const (
	TeamRequestStatusStatusCodeUnknown                    TeamRequestStatusStatusCode = "UNKNOWN"
	TeamRequestStatusStatusCodeSuccess                    TeamRequestStatusStatusCode = "SUCCESS"
	TeamRequestStatusStatusCodeUnauthorized               TeamRequestStatusStatusCode = "UNAUTHORIZED"
	TeamRequestStatusStatusCodePaymentRequired            TeamRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamRequestStatusStatusCodeForbidden                  TeamRequestStatusStatusCode = "FORBIDDEN"
	TeamRequestStatusStatusCodeTimeout                    TeamRequestStatusStatusCode = "TIMEOUT"
	TeamRequestStatusStatusCodeExists                     TeamRequestStatusStatusCode = "EXISTS"
	TeamRequestStatusStatusCodeNotFound                   TeamRequestStatusStatusCode = "NOT_FOUND"
	TeamRequestStatusStatusCodeInternalError              TeamRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamRequestStatusStatusCodeInvalidRequest             TeamRequestStatusStatusCode = "INVALID_REQUEST"
	TeamRequestStatusStatusCodeInvalidRequestVersion      TeamRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamRequestStatusStatusCodeInvalidRequestData         TeamRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamRequestStatusStatusCodeMethodNotAllowed           TeamRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamRequestStatusStatusCodeConflict                   TeamRequestStatusStatusCode = "CONFLICT"
	TeamRequestStatusStatusCodeUnprocessableEntity        TeamRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamRequestStatusStatusCodeTooManyRequests            TeamRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamRequestStatusStatusCodeInsufficientStorage        TeamRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamRequestStatusStatusCodeServiceUnavailable         TeamRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamRequestStatusStatusCodePayloadTooLarge            TeamRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamRequestStatusStatusCodeNotAcceptable              TeamRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamRequestStatusStatusCodeUnavailableForLegalReasons TeamRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamRequestStatusStatusCodeBadGateway                 TeamRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamRequestStatusStatusCodeUnknown, TeamRequestStatusStatusCodeSuccess, TeamRequestStatusStatusCodeUnauthorized, TeamRequestStatusStatusCodePaymentRequired, TeamRequestStatusStatusCodeForbidden, TeamRequestStatusStatusCodeTimeout, TeamRequestStatusStatusCodeExists, TeamRequestStatusStatusCodeNotFound, TeamRequestStatusStatusCodeInternalError, TeamRequestStatusStatusCodeInvalidRequest, TeamRequestStatusStatusCodeInvalidRequestVersion, TeamRequestStatusStatusCodeInvalidRequestData, TeamRequestStatusStatusCodeMethodNotAllowed, TeamRequestStatusStatusCodeConflict, TeamRequestStatusStatusCodeUnprocessableEntity, TeamRequestStatusStatusCodeTooManyRequests, TeamRequestStatusStatusCodeInsufficientStorage, TeamRequestStatusStatusCodeServiceUnavailable, TeamRequestStatusStatusCodePayloadTooLarge, TeamRequestStatusStatusCodeNotAcceptable, TeamRequestStatusStatusCodeUnavailableForLegalReasons, TeamRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type TeamTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             teamTeamJSON             `json:"-"`
}

// teamTeamJSON contains the JSON metadata for the struct [TeamTeam]
type teamTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TeamTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamTeamInfinityManagerSettingsJSON `json:"-"`
}

// teamTeamInfinityManagerSettingsJSON contains the JSON metadata for the struct
// [TeamTeamInfinityManagerSettings]
type teamTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamTeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                         `json:"repoScanShowResults"`
	JSON                teamTeamRepoScanSettingsJSON `json:"-"`
}

// teamTeamRepoScanSettingsJSON contains the JSON metadata for the struct
// [TeamTeamRepoScanSettings]
type teamTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// listing of all teams
type TeamList struct {
	// object that describes the pagination information
	PaginationInfo TeamListPaginationInfo `json:"paginationInfo"`
	RequestStatus  TeamListRequestStatus  `json:"requestStatus"`
	Teams          []TeamListTeam         `json:"teams"`
	JSON           teamListJSON           `json:"-"`
}

// teamListJSON contains the JSON metadata for the struct [TeamList]
type teamListJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Teams          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TeamList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type TeamListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                      `json:"totalResults"`
	JSON         teamListPaginationInfoJSON `json:"-"`
}

// teamListPaginationInfoJSON contains the JSON metadata for the struct
// [TeamListPaginationInfo]
type teamListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type TeamListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                          `json:"statusDescription"`
	JSON              teamListRequestStatusJSON       `json:"-"`
}

// teamListRequestStatusJSON contains the JSON metadata for the struct
// [TeamListRequestStatus]
type teamListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamListRequestStatusStatusCode string

const (
	TeamListRequestStatusStatusCodeUnknown                    TeamListRequestStatusStatusCode = "UNKNOWN"
	TeamListRequestStatusStatusCodeSuccess                    TeamListRequestStatusStatusCode = "SUCCESS"
	TeamListRequestStatusStatusCodeUnauthorized               TeamListRequestStatusStatusCode = "UNAUTHORIZED"
	TeamListRequestStatusStatusCodePaymentRequired            TeamListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamListRequestStatusStatusCodeForbidden                  TeamListRequestStatusStatusCode = "FORBIDDEN"
	TeamListRequestStatusStatusCodeTimeout                    TeamListRequestStatusStatusCode = "TIMEOUT"
	TeamListRequestStatusStatusCodeExists                     TeamListRequestStatusStatusCode = "EXISTS"
	TeamListRequestStatusStatusCodeNotFound                   TeamListRequestStatusStatusCode = "NOT_FOUND"
	TeamListRequestStatusStatusCodeInternalError              TeamListRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamListRequestStatusStatusCodeInvalidRequest             TeamListRequestStatusStatusCode = "INVALID_REQUEST"
	TeamListRequestStatusStatusCodeInvalidRequestVersion      TeamListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamListRequestStatusStatusCodeInvalidRequestData         TeamListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamListRequestStatusStatusCodeMethodNotAllowed           TeamListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamListRequestStatusStatusCodeConflict                   TeamListRequestStatusStatusCode = "CONFLICT"
	TeamListRequestStatusStatusCodeUnprocessableEntity        TeamListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamListRequestStatusStatusCodeTooManyRequests            TeamListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamListRequestStatusStatusCodeInsufficientStorage        TeamListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamListRequestStatusStatusCodeServiceUnavailable         TeamListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamListRequestStatusStatusCodePayloadTooLarge            TeamListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamListRequestStatusStatusCodeNotAcceptable              TeamListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamListRequestStatusStatusCodeUnavailableForLegalReasons TeamListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamListRequestStatusStatusCodeBadGateway                 TeamListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamListRequestStatusStatusCodeUnknown, TeamListRequestStatusStatusCodeSuccess, TeamListRequestStatusStatusCodeUnauthorized, TeamListRequestStatusStatusCodePaymentRequired, TeamListRequestStatusStatusCodeForbidden, TeamListRequestStatusStatusCodeTimeout, TeamListRequestStatusStatusCodeExists, TeamListRequestStatusStatusCodeNotFound, TeamListRequestStatusStatusCodeInternalError, TeamListRequestStatusStatusCodeInvalidRequest, TeamListRequestStatusStatusCodeInvalidRequestVersion, TeamListRequestStatusStatusCodeInvalidRequestData, TeamListRequestStatusStatusCodeMethodNotAllowed, TeamListRequestStatusStatusCodeConflict, TeamListRequestStatusStatusCodeUnprocessableEntity, TeamListRequestStatusStatusCodeTooManyRequests, TeamListRequestStatusStatusCodeInsufficientStorage, TeamListRequestStatusStatusCodeServiceUnavailable, TeamListRequestStatusStatusCodePayloadTooLarge, TeamListRequestStatusStatusCodeNotAcceptable, TeamListRequestStatusStatusCodeUnavailableForLegalReasons, TeamListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type TeamListTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamListTeamsInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamListTeamsRepoScanSettings `json:"repoScanSettings"`
	JSON             teamListTeamJSON              `json:"-"`
}

// teamListTeamJSON contains the JSON metadata for the struct [TeamListTeam]
type teamListTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TeamListTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamListTeamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                     `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamListTeamsInfinityManagerSettingsJSON `json:"-"`
}

// teamListTeamsInfinityManagerSettingsJSON contains the JSON metadata for the
// struct [TeamListTeamsInfinityManagerSettings]
type teamListTeamsInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamListTeamsInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListTeamsInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamListTeamsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                              `json:"repoScanShowResults"`
	JSON                teamListTeamsRepoScanSettingsJSON `json:"-"`
}

// teamListTeamsRepoScanSettingsJSON contains the JSON metadata for the struct
// [TeamListTeamsRepoScanSettings]
type teamListTeamsRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamListTeamsRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListTeamsRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// Response for List User reponse
type UserList struct {
	// object that describes the pagination information
	PaginationInfo UserListPaginationInfo `json:"paginationInfo"`
	RequestStatus  UserListRequestStatus  `json:"requestStatus"`
	// information about the user
	Users []UserListUser `json:"users"`
	JSON  userListJSON   `json:"-"`
}

// userListJSON contains the JSON metadata for the struct [UserList]
type userListJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Users          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type UserListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                      `json:"totalResults"`
	JSON         userListPaginationInfoJSON `json:"-"`
}

// userListPaginationInfoJSON contains the JSON metadata for the struct
// [UserListPaginationInfo]
type userListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type UserListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                          `json:"statusDescription"`
	JSON              userListRequestStatusJSON       `json:"-"`
}

// userListRequestStatusJSON contains the JSON metadata for the struct
// [UserListRequestStatus]
type userListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserListRequestStatusStatusCode string

const (
	UserListRequestStatusStatusCodeUnknown                    UserListRequestStatusStatusCode = "UNKNOWN"
	UserListRequestStatusStatusCodeSuccess                    UserListRequestStatusStatusCode = "SUCCESS"
	UserListRequestStatusStatusCodeUnauthorized               UserListRequestStatusStatusCode = "UNAUTHORIZED"
	UserListRequestStatusStatusCodePaymentRequired            UserListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserListRequestStatusStatusCodeForbidden                  UserListRequestStatusStatusCode = "FORBIDDEN"
	UserListRequestStatusStatusCodeTimeout                    UserListRequestStatusStatusCode = "TIMEOUT"
	UserListRequestStatusStatusCodeExists                     UserListRequestStatusStatusCode = "EXISTS"
	UserListRequestStatusStatusCodeNotFound                   UserListRequestStatusStatusCode = "NOT_FOUND"
	UserListRequestStatusStatusCodeInternalError              UserListRequestStatusStatusCode = "INTERNAL_ERROR"
	UserListRequestStatusStatusCodeInvalidRequest             UserListRequestStatusStatusCode = "INVALID_REQUEST"
	UserListRequestStatusStatusCodeInvalidRequestVersion      UserListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserListRequestStatusStatusCodeInvalidRequestData         UserListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserListRequestStatusStatusCodeMethodNotAllowed           UserListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserListRequestStatusStatusCodeConflict                   UserListRequestStatusStatusCode = "CONFLICT"
	UserListRequestStatusStatusCodeUnprocessableEntity        UserListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserListRequestStatusStatusCodeTooManyRequests            UserListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserListRequestStatusStatusCodeInsufficientStorage        UserListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserListRequestStatusStatusCodeServiceUnavailable         UserListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserListRequestStatusStatusCodePayloadTooLarge            UserListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserListRequestStatusStatusCodeNotAcceptable              UserListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserListRequestStatusStatusCodeUnavailableForLegalReasons UserListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserListRequestStatusStatusCodeBadGateway                 UserListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserListRequestStatusStatusCodeUnknown, UserListRequestStatusStatusCodeSuccess, UserListRequestStatusStatusCodeUnauthorized, UserListRequestStatusStatusCodePaymentRequired, UserListRequestStatusStatusCodeForbidden, UserListRequestStatusStatusCodeTimeout, UserListRequestStatusStatusCodeExists, UserListRequestStatusStatusCodeNotFound, UserListRequestStatusStatusCodeInternalError, UserListRequestStatusStatusCodeInvalidRequest, UserListRequestStatusStatusCodeInvalidRequestVersion, UserListRequestStatusStatusCodeInvalidRequestData, UserListRequestStatusStatusCodeMethodNotAllowed, UserListRequestStatusStatusCodeConflict, UserListRequestStatusStatusCodeUnprocessableEntity, UserListRequestStatusStatusCodeTooManyRequests, UserListRequestStatusStatusCodeInsufficientStorage, UserListRequestStatusStatusCodeServiceUnavailable, UserListRequestStatusStatusCodePayloadTooLarge, UserListRequestStatusStatusCodeNotAcceptable, UserListRequestStatusStatusCodeUnavailableForLegalReasons, UserListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// information about the user
type UserListUser struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType UserListUsersIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []UserListUsersRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []UserListUsersStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata UserListUsersUserMetadata `json:"userMetadata"`
	JSON         userListUserJSON          `json:"-"`
}

// userListUserJSON contains the JSON metadata for the struct [UserListUser]
type userListUserJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *UserListUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUserJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type UserListUsersIdpType string

const (
	UserListUsersIdpTypeNvidia     UserListUsersIdpType = "NVIDIA"
	UserListUsersIdpTypeEnterprise UserListUsersIdpType = "ENTERPRISE"
)

func (r UserListUsersIdpType) IsKnown() bool {
	switch r {
	case UserListUsersIdpTypeNvidia, UserListUsersIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type UserListUsersRole struct {
	// Information about the Organization
	Org UserListUsersRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier UserListUsersRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team UserListUsersRolesTeam `json:"team"`
	// List of team role types that the user have
	TeamRoles []string              `json:"teamRoles"`
	JSON      userListUsersRoleJSON `json:"-"`
}

// userListUsersRoleJSON contains the JSON metadata for the struct
// [UserListUsersRole]
type userListUsersRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserListUsersRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type UserListUsersRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact UserListUsersRolesOrgAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings UserListUsersRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner UserListUsersRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []UserListUsersRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                     `json:"pecSfdcId"`
	ProductEnablements   []UserListUsersRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []UserListUsersRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings UserListUsersRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             UserListUsersRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo UserListUsersRolesOrgUsersInfo `json:"usersInfo"`
	JSON      userListUsersRolesOrgJSON      `json:"-"`
}

// userListUsersRolesOrgJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrg]
type userListUsersRolesOrgJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *UserListUsersRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type UserListUsersRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                    `json:"fullName"`
	JSON     userListUsersRolesOrgAlternateContactJSON `json:"-"`
}

// userListUsersRolesOrgAlternateContactJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgAlternateContact]
type userListUsersRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserListUsersRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                             `json:"infinityManagerEnableTeamOverride"`
	JSON                              userListUsersRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// userListUsersRolesOrgInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [UserListUsersRolesOrgInfinityManagerSettings]
type userListUsersRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserListUsersRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type UserListUsersRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                            `json:"lastLoginDate"`
	JSON          userListUsersRolesOrgOrgOwnerJSON `json:"-"`
}

// userListUsersRolesOrgOrgOwnerJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrgOrgOwner]
type userListUsersRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListUsersRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type UserListUsersRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type UserListUsersRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                            `json:"expirationDate"`
	PoDetails      []UserListUsersRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           userListUsersRolesOrgProductEnablementJSON        `json:"-"`
}

// userListUsersRolesOrgProductEnablementJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgProductEnablement]
type userListUsersRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type UserListUsersRolesOrgProductEnablementsType string

const (
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval       UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr        UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	UserListUsersRolesOrgProductEnablementsTypeEmsEval            UserListUsersRolesOrgProductEnablementsType = "EMS_EVAL"
	UserListUsersRolesOrgProductEnablementsTypeEmsNfr             UserListUsersRolesOrgProductEnablementsType = "EMS_NFR"
	UserListUsersRolesOrgProductEnablementsTypeEmsCommercial      UserListUsersRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper  UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r UserListUsersRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval, UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr, UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial, UserListUsersRolesOrgProductEnablementsTypeEmsEval, UserListUsersRolesOrgProductEnablementsTypeEmsNfr, UserListUsersRolesOrgProductEnablementsTypeEmsCommercial, UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type UserListUsersRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                              `json:"pkId"`
	JSON userListUsersRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// userListUsersRolesOrgProductEnablementsPoDetailJSON contains the JSON metadata
// for the struct [UserListUsersRolesOrgProductEnablementsPoDetail]
type userListUsersRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type UserListUsersRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type UserListUsersRolesOrgProductSubscriptionsType `json:"type"`
	JSON userListUsersRolesOrgProductSubscriptionJSON  `json:"-"`
}

// userListUsersRolesOrgProductSubscriptionJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgProductSubscription]
type userListUsersRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type UserListUsersRolesOrgProductSubscriptionsType string

const (
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval       UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr        UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r UserListUsersRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval, UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr, UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type UserListUsersRolesOrgRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                      `json:"repoScanShowResults"`
	JSON                userListUsersRolesOrgRepoScanSettingsJSON `json:"-"`
}

// userListUsersRolesOrgRepoScanSettingsJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgRepoScanSettings]
type userListUsersRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserListUsersRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type UserListUsersRolesOrgType string

const (
	UserListUsersRolesOrgTypeUnknown    UserListUsersRolesOrgType = "UNKNOWN"
	UserListUsersRolesOrgTypeCloud      UserListUsersRolesOrgType = "CLOUD"
	UserListUsersRolesOrgTypeEnterprise UserListUsersRolesOrgType = "ENTERPRISE"
	UserListUsersRolesOrgTypeIndividual UserListUsersRolesOrgType = "INDIVIDUAL"
)

func (r UserListUsersRolesOrgType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgTypeUnknown, UserListUsersRolesOrgTypeCloud, UserListUsersRolesOrgTypeEnterprise, UserListUsersRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type UserListUsersRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                              `json:"totalUsers"`
	JSON       userListUsersRolesOrgUsersInfoJSON `json:"-"`
}

// userListUsersRolesOrgUsersInfoJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrgUsersInfo]
type userListUsersRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type UserListUsersRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                            `json:"userId"`
	JSON   userListUsersRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// userListUsersRolesTargetSystemUserIdentifierJSON contains the JSON metadata for
// the struct [UserListUsersRolesTargetSystemUserIdentifier]
type userListUsersRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// Information about the team
type UserListUsersRolesTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings UserListUsersRolesTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings UserListUsersRolesTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             userListUsersRolesTeamJSON             `json:"-"`
}

// userListUsersRolesTeamJSON contains the JSON metadata for the struct
// [UserListUsersRolesTeam]
type userListUsersRolesTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserListUsersRolesTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserListUsersRolesTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                              `json:"infinityManagerEnableTeamOverride"`
	JSON                              userListUsersRolesTeamInfinityManagerSettingsJSON `json:"-"`
}

// userListUsersRolesTeamInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [UserListUsersRolesTeamInfinityManagerSettings]
type userListUsersRolesTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserListUsersRolesTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type UserListUsersRolesTeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                       `json:"repoScanShowResults"`
	JSON                userListUsersRolesTeamRepoScanSettingsJSON `json:"-"`
}

// userListUsersRolesTeamRepoScanSettingsJSON contains the JSON metadata for the
// struct [UserListUsersRolesTeamRepoScanSettings]
type userListUsersRolesTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserListUsersRolesTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type UserListUsersStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                         `json:"workspacesUsage"`
	JSON            userListUsersStorageQuotaJSON `json:"-"`
}

// userListUsersStorageQuotaJSON contains the JSON metadata for the struct
// [UserListUsersStorageQuota]
type userListUsersStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserListUsersStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type UserListUsersUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                        `json:"role"`
	JSON userListUsersUserMetadataJSON `json:"-"`
}

// userListUsersUserMetadataJSON contains the JSON metadata for the struct
// [UserListUsersUserMetadata]
type userListUsersUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersUserMetadataJSON) RawJSON() string {
	return r.raw
}

type OrgNewParams struct {
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
	OrgOwner param.Field[OrgNewParamsOrgOwner] `json:"orgOwner"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                          `json:"pecSfdcId"`
	ProductEnablements param.Field[[]OrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]OrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Proto org identifier
	ProtoOrgID param.Field[string] `json:"protoOrgId"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]             `json:"sendEmail"`
	Type      param.Field[OrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]           `cookie:"ncid"`
	VisitorID param.Field[string]           `cookie:"VisitorID"`
}

func (r OrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type OrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName"`
	// Identity Provider ID of the org owner.
	IdpID param.Field[string] `json:"idpId"`
	// Starfleet ID of the org owner.
	StarfleetID param.Field[string] `json:"starfleetId"`
}

func (r OrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type OrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[OrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                   `json:"expirationDate"`
	PoDetails      param.Field[[]OrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r OrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type OrgNewParamsProductEnablementsType string

const (
	OrgNewParamsProductEnablementsTypeNgcAdminEval       OrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgNewParamsProductEnablementsTypeNgcAdminNfr        OrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	OrgNewParamsProductEnablementsTypeNgcAdminCommercial OrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgNewParamsProductEnablementsTypeEmsEval            OrgNewParamsProductEnablementsType = "EMS_EVAL"
	OrgNewParamsProductEnablementsTypeEmsNfr             OrgNewParamsProductEnablementsType = "EMS_NFR"
	OrgNewParamsProductEnablementsTypeEmsCommercial      OrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	OrgNewParamsProductEnablementsTypeNgcAdminDeveloper  OrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgNewParamsProductEnablementsTypeNgcAdminEval, OrgNewParamsProductEnablementsTypeNgcAdminNfr, OrgNewParamsProductEnablementsTypeNgcAdminCommercial, OrgNewParamsProductEnablementsTypeEmsEval, OrgNewParamsProductEnablementsTypeEmsNfr, OrgNewParamsProductEnablementsTypeEmsCommercial, OrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r OrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type OrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[OrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[OrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r OrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       OrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgNewParamsProductSubscriptionsType string

const (
	OrgNewParamsProductSubscriptionsTypeNgcAdminEval       OrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgNewParamsProductSubscriptionsTypeNgcAdminNfr        OrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgNewParamsProductSubscriptionsTypeNgcAdminCommercial OrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgNewParamsProductSubscriptionsTypeNgcAdminEval, OrgNewParamsProductSubscriptionsTypeNgcAdminNfr, OrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type OrgNewParamsType string

const (
	OrgNewParamsTypeUnknown    OrgNewParamsType = "UNKNOWN"
	OrgNewParamsTypeCloud      OrgNewParamsType = "CLOUD"
	OrgNewParamsTypeEnterprise OrgNewParamsType = "ENTERPRISE"
	OrgNewParamsTypeIndividual OrgNewParamsType = "INDIVIDUAL"
)

func (r OrgNewParamsType) IsKnown() bool {
	switch r {
	case OrgNewParamsTypeUnknown, OrgNewParamsTypeCloud, OrgNewParamsTypeEnterprise, OrgNewParamsTypeIndividual:
		return true
	}
	return false
}

type OrgUpdateParams struct {
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users
	DisplayName param.Field[string] `json:"displayName"`
}

func (r OrgUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgListParams struct {
	FilterUsingOrgDisplayName param.Field[string] `query:"Filter using org display name"`
	FilterUsingOrgName        param.Field[string] `query:"Filter using org name"`
	FilterUsingOrgOwnerEmail  param.Field[string] `query:"Filter using org owner email"`
	FilterUsingOrgOwnerName   param.Field[string] `query:"Filter using org owner name"`
	FilterUsingPecID          param.Field[string] `query:"Filter using PEC ID"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
}

// URLQuery serializes [OrgListParams]'s query parameters as `url.Values`.
func (r OrgListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
