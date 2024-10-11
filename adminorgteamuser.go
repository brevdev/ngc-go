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
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

// AdminOrgTeamUserService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgTeamUserService] method instead.
type AdminOrgTeamUserService struct {
	Options []option.RequestOption
}

// NewAdminOrgTeamUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgTeamUserService(opts ...option.RequestOption) (r *AdminOrgTeamUserService) {
	r = &AdminOrgTeamUserService{}
	r.Options = opts
	return
}

// Create an Org-Admin User (Super Admin privileges required)
func (r *AdminOrgTeamUserService) New(ctx context.Context, orgName string, teamName string, params AdminOrgTeamUserNewParams, opts ...option.RequestOption) (res *shared.User, err error) {
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
func (r *AdminOrgTeamUserService) Add(ctx context.Context, orgName string, teamName string, id string, params AdminOrgTeamUserAddParams, opts ...option.RequestOption) (res *shared.User, err error) {
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

// Add user role in team.
func (r *AdminOrgTeamUserService) AddRole(ctx context.Context, orgName string, teamName string, id string, body AdminOrgTeamUserAddRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
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
	path := fmt.Sprintf("v2/admin/org/%s/team/%s/users/%s/add-role", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AdminOrgTeamUserNewParams struct {
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
	UserMetadata param.Field[AdminOrgTeamUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                                `cookie:"ncid"`
	VisitorID    param.Field[string]                                `cookie:"VisitorID"`
}

func (r AdminOrgTeamUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AdminOrgTeamUserNewParams]'s query parameters as
// `url.Values`.
func (r AdminOrgTeamUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type AdminOrgTeamUserNewParamsUserMetadata struct {
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

func (r AdminOrgTeamUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdminOrgTeamUserAddParams struct {
	UserRoleDefaultsToRegistryRead param.Field[string] `query:"user role, defaults to REGISTRY_READ"`
	Ncid                           param.Field[string] `cookie:"ncid"`
	VisitorID                      param.Field[string] `cookie:"VisitorID"`
}

// URLQuery serializes [AdminOrgTeamUserAddParams]'s query parameters as
// `url.Values`.
func (r AdminOrgTeamUserAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgTeamUserAddRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [AdminOrgTeamUserAddRoleParams]'s query parameters as
// `url.Values`.
func (r AdminOrgTeamUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
