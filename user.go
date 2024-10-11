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
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	APIKey  *UserAPIKeyService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.APIKey = NewUserAPIKeyService(opts...)
	return
}

// What am I?
func (r *UserService) Get(ctx context.Context, query UserGetParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edit current user profile
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// response to a request to access key such as docker token
type UserKey struct {
	Key              string               `json:"key,required"`
	CloudNfsKey      string               `json:"cloudNfsKey"`
	CloudNfsKeyPwd   string               `json:"cloudNfsKeyPwd"`
	CloudNfsUserName string               `json:"cloudNfsUserName"`
	RequestStatus    UserKeyRequestStatus `json:"requestStatus"`
	JSON             userKeyJSON          `json:"-"`
}

// userKeyJSON contains the JSON metadata for the struct [UserKey]
type userKeyJSON struct {
	Key              apijson.Field
	CloudNfsKey      apijson.Field
	CloudNfsKeyPwd   apijson.Field
	CloudNfsUserName apijson.Field
	RequestStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyJSON) RawJSON() string {
	return r.raw
}

type UserKeyRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserKeyRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                         `json:"statusDescription"`
	JSON              userKeyRequestStatusJSON       `json:"-"`
}

// userKeyRequestStatusJSON contains the JSON metadata for the struct
// [UserKeyRequestStatus]
type userKeyRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserKeyRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserKeyRequestStatusStatusCode string

const (
	UserKeyRequestStatusStatusCodeUnknown                    UserKeyRequestStatusStatusCode = "UNKNOWN"
	UserKeyRequestStatusStatusCodeSuccess                    UserKeyRequestStatusStatusCode = "SUCCESS"
	UserKeyRequestStatusStatusCodeUnauthorized               UserKeyRequestStatusStatusCode = "UNAUTHORIZED"
	UserKeyRequestStatusStatusCodePaymentRequired            UserKeyRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserKeyRequestStatusStatusCodeForbidden                  UserKeyRequestStatusStatusCode = "FORBIDDEN"
	UserKeyRequestStatusStatusCodeTimeout                    UserKeyRequestStatusStatusCode = "TIMEOUT"
	UserKeyRequestStatusStatusCodeExists                     UserKeyRequestStatusStatusCode = "EXISTS"
	UserKeyRequestStatusStatusCodeNotFound                   UserKeyRequestStatusStatusCode = "NOT_FOUND"
	UserKeyRequestStatusStatusCodeInternalError              UserKeyRequestStatusStatusCode = "INTERNAL_ERROR"
	UserKeyRequestStatusStatusCodeInvalidRequest             UserKeyRequestStatusStatusCode = "INVALID_REQUEST"
	UserKeyRequestStatusStatusCodeInvalidRequestVersion      UserKeyRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserKeyRequestStatusStatusCodeInvalidRequestData         UserKeyRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserKeyRequestStatusStatusCodeMethodNotAllowed           UserKeyRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserKeyRequestStatusStatusCodeConflict                   UserKeyRequestStatusStatusCode = "CONFLICT"
	UserKeyRequestStatusStatusCodeUnprocessableEntity        UserKeyRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserKeyRequestStatusStatusCodeTooManyRequests            UserKeyRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserKeyRequestStatusStatusCodeInsufficientStorage        UserKeyRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserKeyRequestStatusStatusCodeServiceUnavailable         UserKeyRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserKeyRequestStatusStatusCodePayloadTooLarge            UserKeyRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserKeyRequestStatusStatusCodeNotAcceptable              UserKeyRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserKeyRequestStatusStatusCodeUnavailableForLegalReasons UserKeyRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserKeyRequestStatusStatusCodeBadGateway                 UserKeyRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserKeyRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserKeyRequestStatusStatusCodeUnknown, UserKeyRequestStatusStatusCodeSuccess, UserKeyRequestStatusStatusCodeUnauthorized, UserKeyRequestStatusStatusCodePaymentRequired, UserKeyRequestStatusStatusCodeForbidden, UserKeyRequestStatusStatusCodeTimeout, UserKeyRequestStatusStatusCodeExists, UserKeyRequestStatusStatusCodeNotFound, UserKeyRequestStatusStatusCodeInternalError, UserKeyRequestStatusStatusCodeInvalidRequest, UserKeyRequestStatusStatusCodeInvalidRequestVersion, UserKeyRequestStatusStatusCodeInvalidRequestData, UserKeyRequestStatusStatusCodeMethodNotAllowed, UserKeyRequestStatusStatusCodeConflict, UserKeyRequestStatusStatusCodeUnprocessableEntity, UserKeyRequestStatusStatusCodeTooManyRequests, UserKeyRequestStatusStatusCodeInsufficientStorage, UserKeyRequestStatusStatusCodeServiceUnavailable, UserKeyRequestStatusStatusCodePayloadTooLarge, UserKeyRequestStatusStatusCodeNotAcceptable, UserKeyRequestStatusStatusCodeUnavailableForLegalReasons, UserKeyRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type UserGetParams struct {
	OrgName param.Field[string] `query:"org-name"`
}

// URLQuery serializes [UserGetParams]'s query parameters as `url.Values`.
func (r UserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserUpdateParams struct {
	// indicates if user has opt in to nvidia emails
	HasEmailOptIn param.Field[bool] `json:"hasEmailOptIn"`
	// indicates if user has accepted AI Foundry Partnerships End User License
	// Agreement.
	HasSignedAIFoundryPartnershipsEula param.Field[bool] `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command EULA
	HasSignedBaseCommandEula param.Field[bool] `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula param.Field[bool] `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula param.Field[bool] `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula param.Field[bool] `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt End User License Agreement.
	HasSignedCuOptEula param.Field[bool] `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 End User License Agreement.
	HasSignedEarth2Eula param.Field[bool] `json:"hasSignedEarth2EULA"`
	// indicates if user has accepted EGX EULA
	HasSignedEgxEula param.Field[bool] `json:"hasSignedEgxEULA"`
	// indicates if user has accepted NGC EULA
	HasSignedEula param.Field[bool] `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula param.Field[bool] `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula param.Field[bool] `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula param.Field[bool] `json:"hasSignedNVAIEEULA"`
	// indicates if user has accepted NVIDIA EULA
	HasSignedNvidiaEula param.Field[bool] `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula param.Field[bool] `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula param.Field[bool] `json:"hasSignedOmniverseEULA"`
	// indicates if the user has signed the Privacy Policy
	HasSignedPrivacyPolicy param.Field[bool] `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula param.Field[bool] `json:"hasSignedThirdPartyRegistryShareEULA"`
	// user name
	Name param.Field[string] `json:"name"`
	// Metadata information about the user.
	UserMetadata param.Field[UserUpdateParamsUserMetadata] `json:"userMetadata"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata information about the user.
type UserUpdateParamsUserMetadata struct {
	// Name of the company
	Company param.Field[string] `json:"company"`
	// Company URL
	CompanyURL param.Field[string] `json:"companyUrl"`
	// Country of the user
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"firstName"`
	// Industry segment
	Industry param.Field[string] `json:"industry"`
	// List of development areas that user has interest
	Interest param.Field[[]string] `json:"interest"`
	// User's last name
	LastName param.Field[string] `json:"lastName"`
	// Role of the user in the company
	Role param.Field[string] `json:"role"`
}

func (r UserUpdateParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
