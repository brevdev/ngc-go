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

// AdminUserService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUserService] method instead.
type AdminUserService struct {
	Options []option.RequestOption
}

// NewAdminUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminUserService(opts ...option.RequestOption) (r *AdminUserService) {
	r = &AdminUserService{}
	r.Options = opts
	return
}

// Get User details
func (r *AdminUserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Invite an existing user again (Super Admin privileges required)
func (r *AdminUserService) Invite(ctx context.Context, body AdminUserInviteParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/admin/users/invite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// What am I? Admin version, shows more info than regular endpoint
func (r *AdminUserService) Me(ctx context.Context, query AdminUserMeParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/admin/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Backfill the org owner for individual users
func (r *AdminUserService) OrgOwnerBackfill(ctx context.Context, userID string, opts ...option.RequestOption) (res *AdminUserOrgOwnerBackfillResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user-id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/users/%s/org-owner-backfill", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AdminUserOrgOwnerBackfillResponse struct {
	RequestStatus AdminUserOrgOwnerBackfillResponseRequestStatus `json:"requestStatus"`
	JSON          adminUserOrgOwnerBackfillResponseJSON          `json:"-"`
}

// adminUserOrgOwnerBackfillResponseJSON contains the JSON metadata for the struct
// [AdminUserOrgOwnerBackfillResponse]
type adminUserOrgOwnerBackfillResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminUserOrgOwnerBackfillResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserOrgOwnerBackfillResponseJSON) RawJSON() string {
	return r.raw
}

type AdminUserOrgOwnerBackfillResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                   `json:"statusDescription"`
	JSON              adminUserOrgOwnerBackfillResponseRequestStatusJSON       `json:"-"`
}

// adminUserOrgOwnerBackfillResponseRequestStatusJSON contains the JSON metadata
// for the struct [AdminUserOrgOwnerBackfillResponseRequestStatus]
type adminUserOrgOwnerBackfillResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AdminUserOrgOwnerBackfillResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserOrgOwnerBackfillResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode string

const (
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown                    AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "UNKNOWN"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess                    AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "SUCCESS"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized               AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired            AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden                  AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "FORBIDDEN"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout                    AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "TIMEOUT"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeExists                     AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "EXISTS"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound                   AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_FOUND"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError              AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest             AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion      AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData         AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed           AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeConflict                   AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "CONFLICT"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity        AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests            AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage        AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable         AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge            AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable              AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway                 AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AdminUserOrgOwnerBackfillResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeExists, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeConflict, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons, AdminUserOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type AdminUserInviteParams struct {
	Email param.Field[string] `query:"email,required"`
	// Boolean to send email notification, default is true
	SendEmail param.Field[bool] `query:"send-email"`
}

// URLQuery serializes [AdminUserInviteParams]'s query parameters as `url.Values`.
func (r AdminUserInviteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminUserMeParams struct {
	OrgName param.Field[string] `query:"org-name"`
}

// URLQuery serializes [AdminUserMeParams]'s query parameters as `url.Values`.
func (r AdminUserMeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
