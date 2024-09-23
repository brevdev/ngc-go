// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/apiquery"
	"github.com/stainless-sdks/ngc-go/internal/param"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
	"github.com/stainless-sdks/ngc-go/shared"
)

// SuperAdminUserOrgTeamUserService contains methods and other services that help
// with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminUserOrgTeamUserService] method instead.
type SuperAdminUserOrgTeamUserService struct {
	Options []option.RequestOption
}

// NewSuperAdminUserOrgTeamUserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSuperAdminUserOrgTeamUserService(opts ...option.RequestOption) (r *SuperAdminUserOrgTeamUserService) {
	r = &SuperAdminUserOrgTeamUserService{}
	r.Options = opts
	return
}

// Create an Org-Admin User (Super Admin privileges required)
func (r *SuperAdminUserOrgTeamUserService) New(ctx context.Context, orgName string, teamName string, params SuperAdminUserOrgTeamUserNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/team/%s/users", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add existing User to an Team
func (r *SuperAdminUserOrgTeamUserService) Add(ctx context.Context, orgName string, teamName string, id string, params SuperAdminUserOrgTeamUserAddParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/team/%s/users/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SuperAdminUserOrgTeamUserNewParams struct {
	// Email address of the user. This should be unique.
	Email param.Field[string] `json:"email,required"`
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization
	IdpID param.Field[string] `query:"idp-id"`
	// Boolean to send email notification, default is true
	SendEmail param.Field[bool] `query:"send-email"`
	// indicates if user has opt in to nvidia emails
	EmailOptIn param.Field[bool] `json:"emailOptIn"`
	// indicates if user has accepted EULA
	EulaAccepted param.Field[bool] `json:"eulaAccepted"`
	// user name
	Name param.Field[string] `json:"name"`
	// DEPRECATED - use roleTypes which allows multiple roles
	RoleType param.Field[string] `json:"roleType"`
	// feature roles to give to the user
	RoleTypes param.Field[[]string] `json:"roleTypes"`
	// user job role
	SalesforceContactJobRole param.Field[string] `json:"salesforceContactJobRole"`
	// Metadata information about the user.
	UserMetadata param.Field[SuperAdminUserOrgTeamUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                                         `cookie:"ncid"`
	VisitorID    param.Field[string]                                         `cookie:"VisitorID"`
}

func (r SuperAdminUserOrgTeamUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SuperAdminUserOrgTeamUserNewParams]'s query parameters as
// `url.Values`.
func (r SuperAdminUserOrgTeamUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type SuperAdminUserOrgTeamUserNewParamsUserMetadata struct {
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

func (r SuperAdminUserOrgTeamUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperAdminUserOrgTeamUserAddParams struct {
	UserRoleDefaultsToRegistryRead param.Field[string] `query:"user role, defaults to REGISTRY_READ"`
	Ncid                           param.Field[string] `cookie:"ncid"`
	VisitorID                      param.Field[string] `cookie:"VisitorID"`
}

// URLQuery serializes [SuperAdminUserOrgTeamUserAddParams]'s query parameters as
// `url.Values`.
func (r SuperAdminUserOrgTeamUserAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
