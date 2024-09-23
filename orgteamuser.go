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

// OrgTeamUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamUserService] method instead.
type OrgTeamUserService struct {
	Options []option.RequestOption
}

// NewOrgTeamUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgTeamUserService(opts ...option.RequestOption) (r *OrgTeamUserService) {
	r = &OrgTeamUserService{}
	r.Options = opts
	return
}

// Remove User from team.
func (r *OrgTeamUserService) Delete(ctx context.Context, orgName string, teamName string, id string, body OrgTeamUserDeleteParams, opts ...option.RequestOption) (res *OrgTeamUserDeleteResponse, err error) {
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
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Invite if user does not exist, otherwise add role in team
func (r *OrgTeamUserService) AddRole(ctx context.Context, orgName string, teamName string, userEmailOrID string, params OrgTeamUserAddRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/%s/add-role", orgName, teamName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Remove role in team if user exists, otherwise remove invitation
func (r *OrgTeamUserService) RemoveRole(ctx context.Context, orgName string, teamName string, userEmailOrID string, body OrgTeamUserRemoveRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/%s/remove-role", orgName, teamName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type OrgTeamUserDeleteResponse struct {
	RequestStatus OrgTeamUserDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgTeamUserDeleteResponseJSON          `json:"-"`
}

// orgTeamUserDeleteResponseJSON contains the JSON metadata for the struct
// [OrgTeamUserDeleteResponse]
type orgTeamUserDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgTeamUserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamUserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgTeamUserDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgTeamUserDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                           `json:"statusDescription"`
	JSON              orgTeamUserDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgTeamUserDeleteResponseRequestStatusJSON contains the JSON metadata for the
// struct [OrgTeamUserDeleteResponseRequestStatus]
type orgTeamUserDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgTeamUserDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamUserDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgTeamUserDeleteResponseRequestStatusStatusCode string

const (
	OrgTeamUserDeleteResponseRequestStatusStatusCodeUnknown                    OrgTeamUserDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeSuccess                    OrgTeamUserDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeUnauthorized               OrgTeamUserDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgTeamUserDeleteResponseRequestStatusStatusCodePaymentRequired            OrgTeamUserDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeForbidden                  OrgTeamUserDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeTimeout                    OrgTeamUserDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeExists                     OrgTeamUserDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeNotFound                   OrgTeamUserDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeInternalError              OrgTeamUserDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgTeamUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgTeamUserDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeConflict                   OrgTeamUserDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgTeamUserDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgTeamUserDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgTeamUserDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgTeamUserDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgTeamUserDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgTeamUserDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgTeamUserDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgTeamUserDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgTeamUserDeleteResponseRequestStatusStatusCodeBadGateway                 OrgTeamUserDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgTeamUserDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgTeamUserDeleteResponseRequestStatusStatusCodeUnknown, OrgTeamUserDeleteResponseRequestStatusStatusCodeSuccess, OrgTeamUserDeleteResponseRequestStatusStatusCodeUnauthorized, OrgTeamUserDeleteResponseRequestStatusStatusCodePaymentRequired, OrgTeamUserDeleteResponseRequestStatusStatusCodeForbidden, OrgTeamUserDeleteResponseRequestStatusStatusCodeTimeout, OrgTeamUserDeleteResponseRequestStatusStatusCodeExists, OrgTeamUserDeleteResponseRequestStatusStatusCodeNotFound, OrgTeamUserDeleteResponseRequestStatusStatusCodeInternalError, OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgTeamUserDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgTeamUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgTeamUserDeleteResponseRequestStatusStatusCodeConflict, OrgTeamUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgTeamUserDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgTeamUserDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgTeamUserDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgTeamUserDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgTeamUserDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgTeamUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgTeamUserDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgTeamUserDeleteParams struct {
	// If anonymize is true, then org owner permission is required.
	Anonymize param.Field[bool] `query:"anonymize"`
}

// URLQuery serializes [OrgTeamUserDeleteParams]'s query parameters as
// `url.Values`.
func (r OrgTeamUserDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserAddRoleParams struct {
	Roles     param.Field[[]string] `query:"roles,required"`
	Ncid      param.Field[string]   `cookie:"ncid"`
	VisitorID param.Field[string]   `cookie:"VisitorID"`
}

// URLQuery serializes [OrgTeamUserAddRoleParams]'s query parameters as
// `url.Values`.
func (r OrgTeamUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgTeamUserRemoveRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgTeamUserRemoveRoleParams]'s query parameters as
// `url.Values`.
func (r OrgTeamUserRemoveRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
