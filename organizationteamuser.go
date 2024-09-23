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
	"github.com/brevdev/ngc-go/shared"
)

// OrganizationTeamUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationTeamUserService] method instead.
type OrganizationTeamUserService struct {
	Options []option.RequestOption
}

// NewOrganizationTeamUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationTeamUserService(opts ...option.RequestOption) (r *OrganizationTeamUserService) {
	r = &OrganizationTeamUserService{}
	r.Options = opts
	return
}

// Creates a User and add them to a team within the org.
func (r *OrganizationTeamUserService) New(ctx context.Context, orgName string, teamName string, params OrganizationTeamUserNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/team/%s/users", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get User details within team context
func (r *OrganizationTeamUserService) Get(ctx context.Context, orgName string, teamName string, id string, query OrganizationTeamUserGetParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
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
	path := fmt.Sprintf("v2/org/%s/team/%s/users/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get list of users in team. (User Admin in team privileges required)
func (r *OrganizationTeamUserService) List(ctx context.Context, orgName string, teamName string, query OrganizationTeamUserListParams, opts ...option.RequestOption) (res *shared.UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/team/%s/users", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes User from team. (Org Admin or Team Admin Privileges Required).
func (r *OrganizationTeamUserService) Delete(ctx context.Context, orgName string, teamName string, id string, opts ...option.RequestOption) (res *OrganizationTeamUserDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/org/%s/team/%s/users/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OrganizationTeamUserDeleteResponse struct {
	RequestStatus OrganizationTeamUserDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          organizationTeamUserDeleteResponseJSON          `json:"-"`
}

// organizationTeamUserDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationTeamUserDeleteResponse]
type organizationTeamUserDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationTeamUserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTeamUserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationTeamUserDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrganizationTeamUserDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                    `json:"statusDescription"`
	JSON              organizationTeamUserDeleteResponseRequestStatusJSON       `json:"-"`
}

// organizationTeamUserDeleteResponseRequestStatusJSON contains the JSON metadata
// for the struct [OrganizationTeamUserDeleteResponseRequestStatus]
type organizationTeamUserDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationTeamUserDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTeamUserDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrganizationTeamUserDeleteResponseRequestStatusStatusCode string

const (
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnknown                    OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeSuccess                    OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnauthorized               OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodePaymentRequired            OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeForbidden                  OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeTimeout                    OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeExists                     OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeNotFound                   OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInternalError              OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequest             OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeConflict                   OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeTooManyRequests            OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeNotAcceptable              OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrganizationTeamUserDeleteResponseRequestStatusStatusCodeBadGateway                 OrganizationTeamUserDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrganizationTeamUserDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnknown, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeSuccess, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnauthorized, OrganizationTeamUserDeleteResponseRequestStatusStatusCodePaymentRequired, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeForbidden, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeTimeout, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeExists, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeNotFound, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInternalError, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequest, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeConflict, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeTooManyRequests, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrganizationTeamUserDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeNotAcceptable, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrganizationTeamUserDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrganizationTeamUserNewParams struct {
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
	UserMetadata param.Field[OrganizationTeamUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                                    `cookie:"ncid"`
	VisitorID    param.Field[string]                                    `cookie:"VisitorID"`
}

func (r OrganizationTeamUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationTeamUserNewParams]'s query parameters as
// `url.Values`.
func (r OrganizationTeamUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type OrganizationTeamUserNewParamsUserMetadata struct {
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

func (r OrganizationTeamUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationTeamUserGetParams struct {
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization. If no IDP is configured for the organization, then IDP is guessed
	// based on the email domain
	IdpID param.Field[string] `query:"idp-id"`
}

// URLQuery serializes [OrganizationTeamUserGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationTeamUserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationTeamUserListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
	// implemented
	Q param.Field[OrganizationTeamUserListParamsQ] `query:"q"`
}

// URLQuery serializes [OrganizationTeamUserListParams]'s query parameters as
// `url.Values`.
func (r OrganizationTeamUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
// implemented
type OrganizationTeamUserListParamsQ struct {
	Fields      param.Field[[]string]                                 `query:"fields"`
	Filters     param.Field[[]OrganizationTeamUserListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                                   `query:"groupBy"`
	OrderBy     param.Field[[]OrganizationTeamUserListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                                    `query:"page"`
	PageSize    param.Field[int64]                                    `query:"pageSize"`
	Query       param.Field[string]                                   `query:"query"`
	QueryFields param.Field[[]string]                                 `query:"queryFields"`
	ScoredSize  param.Field[int64]                                    `query:"scoredSize"`
}

// URLQuery serializes [OrganizationTeamUserListParamsQ]'s query parameters as
// `url.Values`.
func (r OrganizationTeamUserListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationTeamUserListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrganizationTeamUserListParamsQFilter]'s query parameters
// as `url.Values`.
func (r OrganizationTeamUserListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationTeamUserListParamsQOrderBy struct {
	Field param.Field[string]                                      `query:"field"`
	Value param.Field[OrganizationTeamUserListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrganizationTeamUserListParamsQOrderBy]'s query parameters
// as `url.Values`.
func (r OrganizationTeamUserListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationTeamUserListParamsQOrderByValue string

const (
	OrganizationTeamUserListParamsQOrderByValueAsc  OrganizationTeamUserListParamsQOrderByValue = "ASC"
	OrganizationTeamUserListParamsQOrderByValueDesc OrganizationTeamUserListParamsQOrderByValue = "DESC"
)

func (r OrganizationTeamUserListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrganizationTeamUserListParamsQOrderByValueAsc, OrganizationTeamUserListParamsQOrderByValueDesc:
		return true
	}
	return false
}
