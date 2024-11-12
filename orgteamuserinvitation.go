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

// OrgTeamUserInvitationService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamUserInvitationService] method instead.
type OrgTeamUserInvitationService struct {
	Options []option.RequestOption
}

// NewOrgTeamUserInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgTeamUserInvitationService(opts ...option.RequestOption) (r *OrgTeamUserInvitationService) {
	r = &OrgTeamUserInvitationService{}
	r.Options = opts
	return
}

// List invitations in a team. (Team User Admin privileges required)
func (r *OrgTeamUserInvitationService) List(ctx context.Context, orgName string, teamName string, query OrgTeamUserInvitationListParams, opts ...option.RequestOption) (res *UserInvitationList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/team/%s/users/invitations", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific invitation in an team. (Org Admin or Team User Admin
// privileges required)
func (r *OrgTeamUserInvitationService) Delete(ctx context.Context, orgName string, teamName string, id string, opts ...option.RequestOption) (res *OrgTeamUserInvitationDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/org/%s/team/%s/users/invitations/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Resend email of a specific invitation in a team (Org or Team User Admin
// privileges required).
func (r *OrgTeamUserInvitationService) InviteResend(ctx context.Context, orgName string, teamName string, id string, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
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
	path := fmt.Sprintf("v2/org/%s/team/%s/users/invitations/%s/resend-invitation-email", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrgTeamUserInvitationDeleteResponse struct {
	RequestStatus OrgTeamUserInvitationDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgTeamUserInvitationDeleteResponseJSON          `json:"-"`
}

// orgTeamUserInvitationDeleteResponseJSON contains the JSON metadata for the
// struct [OrgTeamUserInvitationDeleteResponse]
type orgTeamUserInvitationDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgTeamUserInvitationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamUserInvitationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgTeamUserInvitationDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                     `json:"statusDescription"`
	JSON              orgTeamUserInvitationDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgTeamUserInvitationDeleteResponseRequestStatusJSON contains the JSON metadata
// for the struct [OrgTeamUserInvitationDeleteResponseRequestStatus]
type orgTeamUserInvitationDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgTeamUserInvitationDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamUserInvitationDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode string

const (
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnknown                    OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeSuccess                    OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized               OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired            OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeForbidden                  OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeTimeout                    OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeExists                     OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeNotFound                   OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInternalError              OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeConflict                   OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway                 OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgTeamUserInvitationDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnknown, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeSuccess, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeForbidden, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeTimeout, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeExists, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeNotFound, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInternalError, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeConflict, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgTeamUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgTeamUserInvitationListParams struct {
	OrderBy param.Field[OrgTeamUserInvitationListParamsOrderBy] `query:"orderBy"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters
	Q param.Field[OrgTeamUserInvitationListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgTeamUserInvitationListParams]'s query parameters as
// `url.Values`.
func (r OrgTeamUserInvitationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserInvitationListParamsOrderBy string

const (
	OrgTeamUserInvitationListParamsOrderByNameAsc  OrgTeamUserInvitationListParamsOrderBy = "NAME_ASC"
	OrgTeamUserInvitationListParamsOrderByNameDesc OrgTeamUserInvitationListParamsOrderBy = "NAME_DESC"
)

func (r OrgTeamUserInvitationListParamsOrderBy) IsKnown() bool {
	switch r {
	case OrgTeamUserInvitationListParamsOrderByNameAsc, OrgTeamUserInvitationListParamsOrderByNameDesc:
		return true
	}
	return false
}

// User Search Parameters
type OrgTeamUserInvitationListParamsQ struct {
	Fields      param.Field[[]string]                                  `query:"fields"`
	Filters     param.Field[[]OrgTeamUserInvitationListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                                    `query:"groupBy"`
	OrderBy     param.Field[[]OrgTeamUserInvitationListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                                     `query:"page"`
	PageSize    param.Field[int64]                                     `query:"pageSize"`
	Query       param.Field[string]                                    `query:"query"`
	QueryFields param.Field[[]string]                                  `query:"queryFields"`
	ScoredSize  param.Field[int64]                                     `query:"scoredSize"`
}

// URLQuery serializes [OrgTeamUserInvitationListParamsQ]'s query parameters as
// `url.Values`.
func (r OrgTeamUserInvitationListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserInvitationListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgTeamUserInvitationListParamsQFilter]'s query parameters
// as `url.Values`.
func (r OrgTeamUserInvitationListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserInvitationListParamsQOrderBy struct {
	Field param.Field[string]                                       `query:"field"`
	Value param.Field[OrgTeamUserInvitationListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgTeamUserInvitationListParamsQOrderBy]'s query parameters
// as `url.Values`.
func (r OrgTeamUserInvitationListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserInvitationListParamsQOrderByValue string

const (
	OrgTeamUserInvitationListParamsQOrderByValueAsc  OrgTeamUserInvitationListParamsQOrderByValue = "ASC"
	OrgTeamUserInvitationListParamsQOrderByValueDesc OrgTeamUserInvitationListParamsQOrderByValue = "DESC"
)

func (r OrgTeamUserInvitationListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgTeamUserInvitationListParamsQOrderByValueAsc, OrgTeamUserInvitationListParamsQOrderByValueDesc:
		return true
	}
	return false
}