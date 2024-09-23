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

// SuperAdminOrgControllerOrgService contains methods and other services that help
// with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminOrgControllerOrgService] method instead.
type SuperAdminOrgControllerOrgService struct {
	Options []option.RequestOption
}

// NewSuperAdminOrgControllerOrgService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSuperAdminOrgControllerOrgService(opts ...option.RequestOption) (r *SuperAdminOrgControllerOrgService) {
	r = &SuperAdminOrgControllerOrgService{}
	r.Options = opts
	return
}

// Get organization or proto organization info. (SuperAdmin privileges required)
func (r *SuperAdminOrgControllerOrgService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/admin/org/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update org information and settings. Superadmin privileges required
func (r *SuperAdminOrgControllerOrgService) Update(ctx context.Context, orgName string, body SuperAdminOrgControllerOrgUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/admin/org/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SuperAdminOrgControllerOrgUpdateParams struct {
	// Org Owner Alternate Contact
	AlternateContact param.Field[SuperAdminOrgControllerOrgUpdateParamsAlternateContact] `json:"alternateContact"`
	// Name of the company
	CompanyName param.Field[string] `json:"companyName"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identity Provider ID.
	IdpID param.Field[string] `json:"idpId"`
	// Infinity manager setting definition
	InfinityManagerSettings param.Field[SuperAdminOrgControllerOrgUpdateParamsInfinityManagerSettings] `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled param.Field[bool] `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled param.Field[bool] `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled param.Field[bool] `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled param.Field[bool] `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled param.Field[bool] `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in Base Command Platform job
	// telemetry
	IsSeparateInfluxDBUsed param.Field[bool] `json:"isSeparateInfluxDbUsed"`
	// Org owner.
	OrgOwner param.Field[SuperAdminOrgControllerOrgUpdateParamsOrgOwner] `json:"orgOwner"`
	// Org owners
	OrgOwners param.Field[[]SuperAdminOrgControllerOrgUpdateParamsOrgOwner] `json:"orgOwners"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID            param.Field[string]                                                      `json:"pecSfdcId"`
	ProductEnablements   param.Field[[]SuperAdminOrgControllerOrgUpdateParamsProductEnablement]   `json:"productEnablements"`
	ProductSubscriptions param.Field[[]SuperAdminOrgControllerOrgUpdateParamsProductSubscription] `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings param.Field[SuperAdminOrgControllerOrgUpdateParamsRepoScanSettings] `json:"repoScanSettings"`
	Type             param.Field[SuperAdminOrgControllerOrgUpdateParamsType]             `json:"type"`
}

func (r SuperAdminOrgControllerOrgUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org Owner Alternate Contact
type SuperAdminOrgControllerOrgUpdateParamsAlternateContact struct {
	// Alternate contact's email.
	Email param.Field[string] `json:"email"`
	// Full name of the alternate contact.
	FullName param.Field[string] `json:"fullName"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsAlternateContact) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Infinity manager setting definition
type SuperAdminOrgControllerOrgUpdateParamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled param.Field[bool] `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride param.Field[bool] `json:"infinityManagerEnableTeamOverride"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsInfinityManagerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type SuperAdminOrgControllerOrgUpdateParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `json:"lastLoginDate"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type SuperAdminOrgControllerOrgUpdateParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                                             `json:"expirationDate"`
	PoDetails      param.Field[[]SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType string

const (
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminEval       SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminNfr        SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_NFR"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminCommercial SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsEval            SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "EMS_EVAL"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsNfr             SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "EMS_NFR"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsCommercial      SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "EMS_COMMERCIAL"
	SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminDeveloper  SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r SuperAdminOrgControllerOrgUpdateParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminEval, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminNfr, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminCommercial, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsEval, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsNfr, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeEmsCommercial, SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type SuperAdminOrgControllerOrgUpdateParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType] `json:"type"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType string

const (
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval       SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval, SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType string

const (
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval       SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminNfr        SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminCommercial SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval, SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminNfr, SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type SuperAdminOrgControllerOrgUpdateParamsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride param.Field[bool] `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault param.Field[bool] `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled param.Field[bool] `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications param.Field[bool] `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride param.Field[bool] `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults param.Field[bool] `json:"repoScanShowResults"`
}

func (r SuperAdminOrgControllerOrgUpdateParamsRepoScanSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperAdminOrgControllerOrgUpdateParamsType string

const (
	SuperAdminOrgControllerOrgUpdateParamsTypeUnknown    SuperAdminOrgControllerOrgUpdateParamsType = "UNKNOWN"
	SuperAdminOrgControllerOrgUpdateParamsTypeCloud      SuperAdminOrgControllerOrgUpdateParamsType = "CLOUD"
	SuperAdminOrgControllerOrgUpdateParamsTypeEnterprise SuperAdminOrgControllerOrgUpdateParamsType = "ENTERPRISE"
	SuperAdminOrgControllerOrgUpdateParamsTypeIndividual SuperAdminOrgControllerOrgUpdateParamsType = "INDIVIDUAL"
)

func (r SuperAdminOrgControllerOrgUpdateParamsType) IsKnown() bool {
	switch r {
	case SuperAdminOrgControllerOrgUpdateParamsTypeUnknown, SuperAdminOrgControllerOrgUpdateParamsTypeCloud, SuperAdminOrgControllerOrgUpdateParamsTypeEnterprise, SuperAdminOrgControllerOrgUpdateParamsTypeIndividual:
		return true
	}
	return false
}
