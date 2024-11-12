// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/apiquery"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// OrgService contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options      []option.RequestOption
	V3           *OrgV3Service
	ProtoOrg     *OrgProtoOrgService
	V2           *OrgV2Service
	Users        *OrgUserService
	Teams        *OrgTeamService
	Credits      *OrgCreditService
	StarfleetIDs *OrgStarfleetIDService
	Metering     *OrgMeteringService
	AuditLogs    *OrgAuditLogService
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r *OrgService) {
	r = &OrgService{}
	r.Options = opts
	r.V3 = NewOrgV3Service(opts...)
	r.ProtoOrg = NewOrgProtoOrgService(opts...)
	r.V2 = NewOrgV2Service(opts...)
	r.Users = NewOrgUserService(opts...)
	r.Teams = NewOrgTeamService(opts...)
	r.Credits = NewOrgCreditService(opts...)
	r.StarfleetIDs = NewOrgStarfleetIDService(opts...)
	r.Metering = NewOrgMeteringService(opts...)
	r.AuditLogs = NewOrgAuditLogService(opts...)
	return
}

// Create a new organization based on the org info provided in the request.
func (r *OrgService) New(ctx context.Context, params OrgNewParams, opts ...option.RequestOption) (res *OrgResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get organization information
func (r *OrgService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *OrgResponse, err error) {
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
func (r *OrgService) Update(ctx context.Context, orgName string, body OrgUpdateParams, opts ...option.RequestOption) (res *OrgResponse, err error) {
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

// info about an organizations
type OrgResponse struct {
	// Information about the Organization
	Organizations OrgResponseOrganizations `json:"organizations"`
	RequestStatus OrgResponseRequestStatus `json:"requestStatus"`
	JSON          orgResponseJSON          `json:"-"`
}

// orgResponseJSON contains the JSON metadata for the struct [OrgResponse]
type orgResponseJSON struct {
	Organizations apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type OrgResponseOrganizations struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact OrgResponseOrganizationsAlternateContact `json:"alternateContact"`
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
	InfinityManagerSettings OrgResponseOrganizationsInfinityManagerSettings `json:"infinityManagerSettings"`
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
	OrgOwner OrgResponseOrganizationsOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []OrgResponseOrganizationsOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                        `json:"pecSfdcId"`
	ProductEnablements   []OrgResponseOrganizationsProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []OrgResponseOrganizationsProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings OrgResponseOrganizationsRepoScanSettings `json:"repoScanSettings"`
	Type             OrgResponseOrganizationsType             `json:"type"`
	// Users information.
	UsersInfo OrgResponseOrganizationsUsersInfo `json:"usersInfo"`
	JSON      orgResponseOrganizationsJSON      `json:"-"`
}

// orgResponseOrganizationsJSON contains the JSON metadata for the struct
// [OrgResponseOrganizations]
type orgResponseOrganizationsJSON struct {
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

func (r *OrgResponseOrganizations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type OrgResponseOrganizationsAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                       `json:"fullName"`
	JSON     orgResponseOrganizationsAlternateContactJSON `json:"-"`
}

// orgResponseOrganizationsAlternateContactJSON contains the JSON metadata for the
// struct [OrgResponseOrganizationsAlternateContact]
type orgResponseOrganizationsAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgResponseOrganizationsAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgResponseOrganizationsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgResponseOrganizationsInfinityManagerSettingsJSON `json:"-"`
}

// orgResponseOrganizationsInfinityManagerSettingsJSON contains the JSON metadata
// for the struct [OrgResponseOrganizationsInfinityManagerSettings]
type orgResponseOrganizationsInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgResponseOrganizationsInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type OrgResponseOrganizationsOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                               `json:"lastLoginDate"`
	JSON          orgResponseOrganizationsOrgOwnerJSON `json:"-"`
}

// orgResponseOrganizationsOrgOwnerJSON contains the JSON metadata for the struct
// [OrgResponseOrganizationsOrgOwner]
type orgResponseOrganizationsOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgResponseOrganizationsOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type OrgResponseOrganizationsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type OrgResponseOrganizationsProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                               `json:"expirationDate"`
	PoDetails      []OrgResponseOrganizationsProductEnablementsPoDetail `json:"poDetails"`
	JSON           orgResponseOrganizationsProductEnablementJSON        `json:"-"`
}

// orgResponseOrganizationsProductEnablementJSON contains the JSON metadata for the
// struct [OrgResponseOrganizationsProductEnablement]
type orgResponseOrganizationsProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrgResponseOrganizationsProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type OrgResponseOrganizationsProductEnablementsType string

const (
	OrgResponseOrganizationsProductEnablementsTypeNgcAdminEval       OrgResponseOrganizationsProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgResponseOrganizationsProductEnablementsTypeNgcAdminNfr        OrgResponseOrganizationsProductEnablementsType = "NGC_ADMIN_NFR"
	OrgResponseOrganizationsProductEnablementsTypeNgcAdminCommercial OrgResponseOrganizationsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgResponseOrganizationsProductEnablementsTypeEmsEval            OrgResponseOrganizationsProductEnablementsType = "EMS_EVAL"
	OrgResponseOrganizationsProductEnablementsTypeEmsNfr             OrgResponseOrganizationsProductEnablementsType = "EMS_NFR"
	OrgResponseOrganizationsProductEnablementsTypeEmsCommercial      OrgResponseOrganizationsProductEnablementsType = "EMS_COMMERCIAL"
	OrgResponseOrganizationsProductEnablementsTypeNgcAdminDeveloper  OrgResponseOrganizationsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgResponseOrganizationsProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgResponseOrganizationsProductEnablementsTypeNgcAdminEval, OrgResponseOrganizationsProductEnablementsTypeNgcAdminNfr, OrgResponseOrganizationsProductEnablementsTypeNgcAdminCommercial, OrgResponseOrganizationsProductEnablementsTypeEmsEval, OrgResponseOrganizationsProductEnablementsTypeEmsNfr, OrgResponseOrganizationsProductEnablementsTypeEmsCommercial, OrgResponseOrganizationsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgResponseOrganizationsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                                 `json:"pkId"`
	JSON orgResponseOrganizationsProductEnablementsPoDetailJSON `json:"-"`
}

// orgResponseOrganizationsProductEnablementsPoDetailJSON contains the JSON
// metadata for the struct [OrgResponseOrganizationsProductEnablementsPoDetail]
type orgResponseOrganizationsProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgResponseOrganizationsProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type OrgResponseOrganizationsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type OrgResponseOrganizationsProductSubscriptionsType `json:"type"`
	JSON orgResponseOrganizationsProductSubscriptionJSON  `json:"-"`
}

// orgResponseOrganizationsProductSubscriptionJSON contains the JSON metadata for
// the struct [OrgResponseOrganizationsProductSubscription]
type orgResponseOrganizationsProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrgResponseOrganizationsProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType string

const (
	OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval       OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgResponseOrganizationsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsEval, OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgResponseOrganizationsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgResponseOrganizationsProductSubscriptionsType string

const (
	OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminEval       OrgResponseOrganizationsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminNfr        OrgResponseOrganizationsProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminCommercial OrgResponseOrganizationsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgResponseOrganizationsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminEval, OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminNfr, OrgResponseOrganizationsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type OrgResponseOrganizationsRepoScanSettings struct {
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
	RepoScanShowResults bool                                         `json:"repoScanShowResults"`
	JSON                orgResponseOrganizationsRepoScanSettingsJSON `json:"-"`
}

// orgResponseOrganizationsRepoScanSettingsJSON contains the JSON metadata for the
// struct [OrgResponseOrganizationsRepoScanSettings]
type orgResponseOrganizationsRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgResponseOrganizationsRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrgResponseOrganizationsType string

const (
	OrgResponseOrganizationsTypeUnknown    OrgResponseOrganizationsType = "UNKNOWN"
	OrgResponseOrganizationsTypeCloud      OrgResponseOrganizationsType = "CLOUD"
	OrgResponseOrganizationsTypeEnterprise OrgResponseOrganizationsType = "ENTERPRISE"
	OrgResponseOrganizationsTypeIndividual OrgResponseOrganizationsType = "INDIVIDUAL"
)

func (r OrgResponseOrganizationsType) IsKnown() bool {
	switch r {
	case OrgResponseOrganizationsTypeUnknown, OrgResponseOrganizationsTypeCloud, OrgResponseOrganizationsTypeEnterprise, OrgResponseOrganizationsTypeIndividual:
		return true
	}
	return false
}

// Users information.
type OrgResponseOrganizationsUsersInfo struct {
	// Total number of users.
	TotalUsers int64                                 `json:"totalUsers"`
	JSON       orgResponseOrganizationsUsersInfoJSON `json:"-"`
}

// orgResponseOrganizationsUsersInfoJSON contains the JSON metadata for the struct
// [OrgResponseOrganizationsUsersInfo]
type orgResponseOrganizationsUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgResponseOrganizationsUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseOrganizationsUsersInfoJSON) RawJSON() string {
	return r.raw
}

type OrgResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                             `json:"statusDescription"`
	JSON              orgResponseRequestStatusJSON       `json:"-"`
}

// orgResponseRequestStatusJSON contains the JSON metadata for the struct
// [OrgResponseRequestStatus]
type orgResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgResponseRequestStatusStatusCode string

const (
	OrgResponseRequestStatusStatusCodeUnknown                    OrgResponseRequestStatusStatusCode = "UNKNOWN"
	OrgResponseRequestStatusStatusCodeSuccess                    OrgResponseRequestStatusStatusCode = "SUCCESS"
	OrgResponseRequestStatusStatusCodeUnauthorized               OrgResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgResponseRequestStatusStatusCodePaymentRequired            OrgResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgResponseRequestStatusStatusCodeForbidden                  OrgResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgResponseRequestStatusStatusCodeTimeout                    OrgResponseRequestStatusStatusCode = "TIMEOUT"
	OrgResponseRequestStatusStatusCodeExists                     OrgResponseRequestStatusStatusCode = "EXISTS"
	OrgResponseRequestStatusStatusCodeNotFound                   OrgResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgResponseRequestStatusStatusCodeInternalError              OrgResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgResponseRequestStatusStatusCodeInvalidRequest             OrgResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgResponseRequestStatusStatusCodeInvalidRequestVersion      OrgResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgResponseRequestStatusStatusCodeInvalidRequestData         OrgResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgResponseRequestStatusStatusCodeMethodNotAllowed           OrgResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgResponseRequestStatusStatusCodeConflict                   OrgResponseRequestStatusStatusCode = "CONFLICT"
	OrgResponseRequestStatusStatusCodeUnprocessableEntity        OrgResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgResponseRequestStatusStatusCodeTooManyRequests            OrgResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgResponseRequestStatusStatusCodeInsufficientStorage        OrgResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgResponseRequestStatusStatusCodeServiceUnavailable         OrgResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgResponseRequestStatusStatusCodePayloadTooLarge            OrgResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgResponseRequestStatusStatusCodeNotAcceptable              OrgResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgResponseRequestStatusStatusCodeBadGateway                 OrgResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgResponseRequestStatusStatusCodeUnknown, OrgResponseRequestStatusStatusCodeSuccess, OrgResponseRequestStatusStatusCodeUnauthorized, OrgResponseRequestStatusStatusCodePaymentRequired, OrgResponseRequestStatusStatusCodeForbidden, OrgResponseRequestStatusStatusCodeTimeout, OrgResponseRequestStatusStatusCodeExists, OrgResponseRequestStatusStatusCodeNotFound, OrgResponseRequestStatusStatusCodeInternalError, OrgResponseRequestStatusStatusCodeInvalidRequest, OrgResponseRequestStatusStatusCodeInvalidRequestVersion, OrgResponseRequestStatusStatusCodeInvalidRequestData, OrgResponseRequestStatusStatusCodeMethodNotAllowed, OrgResponseRequestStatusStatusCodeConflict, OrgResponseRequestStatusStatusCodeUnprocessableEntity, OrgResponseRequestStatusStatusCodeTooManyRequests, OrgResponseRequestStatusStatusCodeInsufficientStorage, OrgResponseRequestStatusStatusCodeServiceUnavailable, OrgResponseRequestStatusStatusCodePayloadTooLarge, OrgResponseRequestStatusStatusCodeNotAcceptable, OrgResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
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
