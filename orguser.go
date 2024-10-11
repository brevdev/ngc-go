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

// OrgUserService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserService] method instead.
type OrgUserService struct {
	Options []option.RequestOption
}

// NewOrgUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgUserService(opts ...option.RequestOption) (r *OrgUserService) {
	r = &OrgUserService{}
	r.Options = opts
	return
}

// Creates a User
func (r *OrgUserService) New(ctx context.Context, orgName string, params OrgUserNewParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get info and role/invitation in an org by email or id
func (r *OrgUserService) Get(ctx context.Context, orgName string, userEmailOrID string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get list of users in organization. (User Admin in org privileges required)
func (r *OrgUserService) List(ctx context.Context, orgName string, query OrgUserListParams, opts ...option.RequestOption) (res *UserList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove User from org.
func (r *OrgUserService) Delete(ctx context.Context, orgName string, id string, body OrgUserDeleteParams, opts ...option.RequestOption) (res *OrgUserDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Add User Role in org.
func (r *OrgUserService) AddRole(ctx context.Context, orgName string, id string, body OrgUserAddRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/%s/add-role", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Invites and creates a User in org
func (r *OrgUserService) NcaInvitations(ctx context.Context, orgName string, body OrgUserNcaInvitationsParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/nca-invitations", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove role in org if user exists, otherwise remove invitation
func (r *OrgUserService) RemoveRole(ctx context.Context, orgName string, userEmailOrID string, body OrgUserRemoveRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s/remove-role", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type OrgUserDeleteResponse struct {
	RequestStatus OrgUserDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgUserDeleteResponseJSON          `json:"-"`
}

// orgUserDeleteResponseJSON contains the JSON metadata for the struct
// [OrgUserDeleteResponse]
type orgUserDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgUserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgUserDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgUserDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                       `json:"statusDescription"`
	JSON              orgUserDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgUserDeleteResponseRequestStatusJSON contains the JSON metadata for the struct
// [OrgUserDeleteResponseRequestStatus]
type orgUserDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgUserDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgUserDeleteResponseRequestStatusStatusCode string

const (
	OrgUserDeleteResponseRequestStatusStatusCodeUnknown                    OrgUserDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgUserDeleteResponseRequestStatusStatusCodeSuccess                    OrgUserDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgUserDeleteResponseRequestStatusStatusCodeUnauthorized               OrgUserDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgUserDeleteResponseRequestStatusStatusCodePaymentRequired            OrgUserDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgUserDeleteResponseRequestStatusStatusCodeForbidden                  OrgUserDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgUserDeleteResponseRequestStatusStatusCodeTimeout                    OrgUserDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgUserDeleteResponseRequestStatusStatusCodeExists                     OrgUserDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgUserDeleteResponseRequestStatusStatusCodeNotFound                   OrgUserDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgUserDeleteResponseRequestStatusStatusCodeInternalError              OrgUserDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgUserDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgUserDeleteResponseRequestStatusStatusCodeConflict                   OrgUserDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgUserDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgUserDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgUserDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgUserDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgUserDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgUserDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgUserDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgUserDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgUserDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgUserDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgUserDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgUserDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgUserDeleteResponseRequestStatusStatusCodeBadGateway                 OrgUserDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgUserDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgUserDeleteResponseRequestStatusStatusCodeUnknown, OrgUserDeleteResponseRequestStatusStatusCodeSuccess, OrgUserDeleteResponseRequestStatusStatusCodeUnauthorized, OrgUserDeleteResponseRequestStatusStatusCodePaymentRequired, OrgUserDeleteResponseRequestStatusStatusCodeForbidden, OrgUserDeleteResponseRequestStatusStatusCodeTimeout, OrgUserDeleteResponseRequestStatusStatusCodeExists, OrgUserDeleteResponseRequestStatusStatusCodeNotFound, OrgUserDeleteResponseRequestStatusStatusCodeInternalError, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgUserDeleteResponseRequestStatusStatusCodeConflict, OrgUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgUserDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgUserDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgUserDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgUserDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgUserDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgUserDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgUserNewParams struct {
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
	UserMetadata param.Field[OrgUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                       `cookie:"ncid"`
	VisitorID    param.Field[string]                       `cookie:"VisitorID"`
}

func (r OrgUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrgUserNewParams]'s query parameters as `url.Values`.
func (r OrgUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type OrgUserNewParamsUserMetadata struct {
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

func (r OrgUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgUserListParams struct {
	// Name of team to exclude members from
	ExcludeFromTeam param.Field[string] `query:"exclude-from-team"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
	// implemented
	Q param.Field[OrgUserListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgUserListParams]'s query parameters as `url.Values`.
func (r OrgUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
// implemented
type OrgUserListParamsQ struct {
	Fields      param.Field[[]string]                    `query:"fields"`
	Filters     param.Field[[]OrgUserListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                      `query:"groupBy"`
	OrderBy     param.Field[[]OrgUserListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                       `query:"page"`
	PageSize    param.Field[int64]                       `query:"pageSize"`
	Query       param.Field[string]                      `query:"query"`
	QueryFields param.Field[[]string]                    `query:"queryFields"`
	ScoredSize  param.Field[int64]                       `query:"scoredSize"`
}

// URLQuery serializes [OrgUserListParamsQ]'s query parameters as `url.Values`.
func (r OrgUserListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgUserListParamsQFilter]'s query parameters as
// `url.Values`.
func (r OrgUserListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQOrderBy struct {
	Field param.Field[string]                         `query:"field"`
	Value param.Field[OrgUserListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgUserListParamsQOrderBy]'s query parameters as
// `url.Values`.
func (r OrgUserListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQOrderByValue string

const (
	OrgUserListParamsQOrderByValueAsc  OrgUserListParamsQOrderByValue = "ASC"
	OrgUserListParamsQOrderByValueDesc OrgUserListParamsQOrderByValue = "DESC"
)

func (r OrgUserListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgUserListParamsQOrderByValueAsc, OrgUserListParamsQOrderByValueDesc:
		return true
	}
	return false
}

type OrgUserDeleteParams struct {
	// If anonymize is true, then org owner permission is required.
	Anonymize param.Field[bool] `query:"anonymize"`
}

// URLQuery serializes [OrgUserDeleteParams]'s query parameters as `url.Values`.
func (r OrgUserDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserAddRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgUserAddRoleParams]'s query parameters as `url.Values`.
func (r OrgUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserNcaInvitationsParams struct {
	// Is the user email
	Email param.Field[string] `json:"email"`
	// Is the numbers of days the invitation will expire
	InvitationExpirationIn param.Field[int64] `json:"invitationExpirationIn"`
	// Nca allow users to be invited as Admin and as Member
	InviteAs param.Field[OrgUserNcaInvitationsParamsInviteAs] `json:"inviteAs"`
	// Is a message to the new user
	Message param.Field[string] `json:"message"`
}

func (r OrgUserNcaInvitationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nca allow users to be invited as Admin and as Member
type OrgUserNcaInvitationsParamsInviteAs string

const (
	OrgUserNcaInvitationsParamsInviteAsAdmin  OrgUserNcaInvitationsParamsInviteAs = "ADMIN"
	OrgUserNcaInvitationsParamsInviteAsMember OrgUserNcaInvitationsParamsInviteAs = "MEMBER"
)

func (r OrgUserNcaInvitationsParamsInviteAs) IsKnown() bool {
	switch r {
	case OrgUserNcaInvitationsParamsInviteAsAdmin, OrgUserNcaInvitationsParamsInviteAsMember:
		return true
	}
	return false
}

type OrgUserRemoveRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgUserRemoveRoleParams]'s query parameters as
// `url.Values`.
func (r OrgUserRemoveRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
