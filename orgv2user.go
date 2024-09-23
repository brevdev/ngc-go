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

// OrgV2UserService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV2UserService] method instead.
type OrgV2UserService struct {
	Options []option.RequestOption
}

// NewOrgV2UserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgV2UserService(opts ...option.RequestOption) (r *OrgV2UserService) {
	r = &OrgV2UserService{}
	r.Options = opts
	return
}

// Creates a User
func (r *OrgV2UserService) New(ctx context.Context, orgName string, params OrgV2UserNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get list of users in organization. (User Admin in org privileges required)
func (r *OrgV2UserService) List(ctx context.Context, orgName string, query OrgV2UserListParams, opts ...option.RequestOption) (res *shared.UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrgV2UserNewParams struct {
	// Email address of the user. This should be unique.
	Email param.Field[string] `json:"email,required"`
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization
	IdpID     param.Field[string] `query:"idp-id"`
	SendEmail param.Field[bool]   `query:"send-email"`
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
	UserMetadata param.Field[OrgV2UserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                         `cookie:"ncid"`
	VisitorID    param.Field[string]                         `cookie:"VisitorID"`
}

func (r OrgV2UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrgV2UserNewParams]'s query parameters as `url.Values`.
func (r OrgV2UserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type OrgV2UserNewParamsUserMetadata struct {
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

func (r OrgV2UserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgV2UserListParams struct {
	// Name of team to exclude members from
	ExcludeFromTeam param.Field[string] `query:"exclude-from-team"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
	// implemented
	Q param.Field[OrgV2UserListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgV2UserListParams]'s query parameters as `url.Values`.
func (r OrgV2UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
// implemented
type OrgV2UserListParamsQ struct {
	Fields      param.Field[[]string]                      `query:"fields"`
	Filters     param.Field[[]OrgV2UserListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                        `query:"groupBy"`
	OrderBy     param.Field[[]OrgV2UserListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                         `query:"page"`
	PageSize    param.Field[int64]                         `query:"pageSize"`
	Query       param.Field[string]                        `query:"query"`
	QueryFields param.Field[[]string]                      `query:"queryFields"`
	ScoredSize  param.Field[int64]                         `query:"scoredSize"`
}

// URLQuery serializes [OrgV2UserListParamsQ]'s query parameters as `url.Values`.
func (r OrgV2UserListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgV2UserListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgV2UserListParamsQFilter]'s query parameters as
// `url.Values`.
func (r OrgV2UserListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgV2UserListParamsQOrderBy struct {
	Field param.Field[string]                           `query:"field"`
	Value param.Field[OrgV2UserListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgV2UserListParamsQOrderBy]'s query parameters as
// `url.Values`.
func (r OrgV2UserListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgV2UserListParamsQOrderByValue string

const (
	OrgV2UserListParamsQOrderByValueAsc  OrgV2UserListParamsQOrderByValue = "ASC"
	OrgV2UserListParamsQOrderByValueDesc OrgV2UserListParamsQOrderByValue = "DESC"
)

func (r OrgV2UserListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgV2UserListParamsQOrderByValueAsc, OrgV2UserListParamsQOrderByValueDesc:
		return true
	}
	return false
}
