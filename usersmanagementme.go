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
	"github.com/stainless-sdks/ngc-go/shared"
)

// UsersManagementMeService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsersManagementMeService] method instead.
type UsersManagementMeService struct {
	Options []option.RequestOption
}

// NewUsersManagementMeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUsersManagementMeService(opts ...option.RequestOption) (r *UsersManagementMeService) {
	r = &UsersManagementMeService{}
	r.Options = opts
	return
}

// What am I?
func (r *UsersManagementMeService) Get(ctx context.Context, query UsersManagementMeGetParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edit current user profile
func (r *UsersManagementMeService) Update(ctx context.Context, body UsersManagementMeUpdateParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type UsersManagementMeGetParams struct {
	OrgName param.Field[string] `query:"org-name"`
}

// URLQuery serializes [UsersManagementMeGetParams]'s query parameters as
// `url.Values`.
func (r UsersManagementMeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsersManagementMeUpdateParams struct {
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
	UserMetadata param.Field[UsersManagementMeUpdateParamsUserMetadata] `json:"userMetadata"`
}

func (r UsersManagementMeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata information about the user.
type UsersManagementMeUpdateParamsUserMetadata struct {
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

func (r UsersManagementMeUpdateParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
