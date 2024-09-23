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

// OrgUserService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserService] method instead.
type OrgUserService struct {
	Options     []option.RequestOption
	Invitations *OrgUserInvitationService
}

// NewOrgUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgUserService(opts ...option.RequestOption) (r *OrgUserService) {
	r = &OrgUserService{}
	r.Options = opts
	r.Invitations = NewOrgUserInvitationService(opts...)
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

// Invite if user does not exist, otherwise add role in org
func (r *OrgUserService) AddRole(ctx context.Context, orgName string, userEmailOrID string, params OrgUserAddRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s/add-role", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Remove role in org if user exists, otherwise remove invitation
func (r *OrgUserService) RemoveRole(ctx context.Context, orgName string, userEmailOrID string, body OrgUserRemoveRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
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

// Update User Role in org
func (r *OrgUserService) UpdateRole(ctx context.Context, orgName string, id string, body OrgUserUpdateRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/%s/update-role", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
	Roles     param.Field[[]string] `query:"roles,required"`
	Ncid      param.Field[string]   `cookie:"ncid"`
	VisitorID param.Field[string]   `cookie:"VisitorID"`
}

// URLQuery serializes [OrgUserAddRoleParams]'s query parameters as `url.Values`.
func (r OrgUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type OrgUserUpdateRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgUserUpdateRoleParams]'s query parameters as
// `url.Values`.
func (r OrgUserUpdateRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
