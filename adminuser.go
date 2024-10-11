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

// AdminUserService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
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

// What am I? Admin version, shows more info than regular endpoint
func (r *AdminUserService) Get(ctx context.Context, query AdminUserGetParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/admin/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Sync crm id with user email (Super Admin privileges required)
func (r *AdminUserService) CRMSync(ctx context.Context, id int64, opts ...option.RequestOption) (res *AdminUserCRMSyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v2/admin/users/%v/crm-sync", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Invite an existing user again (Super Admin privileges required)
func (r *AdminUserService) Invite(ctx context.Context, body AdminUserInviteParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/admin/users/invite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Migrate User Deprecated Roles.
func (r *AdminUserService) MigrateDeprecatedRoles(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/users/%s/migrate-deprecated-roles", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
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

type AdminUserCRMSyncResponse struct {
	RequestStatus AdminUserCRMSyncResponseRequestStatus `json:"requestStatus"`
	JSON          adminUserCRMSyncResponseJSON          `json:"-"`
}

// adminUserCRMSyncResponseJSON contains the JSON metadata for the struct
// [AdminUserCRMSyncResponse]
type adminUserCRMSyncResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminUserCRMSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserCRMSyncResponseJSON) RawJSON() string {
	return r.raw
}

type AdminUserCRMSyncResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AdminUserCRMSyncResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                          `json:"statusDescription"`
	JSON              adminUserCRMSyncResponseRequestStatusJSON       `json:"-"`
}

// adminUserCRMSyncResponseRequestStatusJSON contains the JSON metadata for the
// struct [AdminUserCRMSyncResponseRequestStatus]
type adminUserCRMSyncResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AdminUserCRMSyncResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserCRMSyncResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AdminUserCRMSyncResponseRequestStatusStatusCode string

const (
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnknown                    AdminUserCRMSyncResponseRequestStatusStatusCode = "UNKNOWN"
	AdminUserCRMSyncResponseRequestStatusStatusCodeSuccess                    AdminUserCRMSyncResponseRequestStatusStatusCode = "SUCCESS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized               AdminUserCRMSyncResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired            AdminUserCRMSyncResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AdminUserCRMSyncResponseRequestStatusStatusCodeForbidden                  AdminUserCRMSyncResponseRequestStatusStatusCode = "FORBIDDEN"
	AdminUserCRMSyncResponseRequestStatusStatusCodeTimeout                    AdminUserCRMSyncResponseRequestStatusStatusCode = "TIMEOUT"
	AdminUserCRMSyncResponseRequestStatusStatusCodeExists                     AdminUserCRMSyncResponseRequestStatusStatusCode = "EXISTS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeNotFound                   AdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_FOUND"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInternalError              AdminUserCRMSyncResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest             AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion      AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData         AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed           AdminUserCRMSyncResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AdminUserCRMSyncResponseRequestStatusStatusCodeConflict                   AdminUserCRMSyncResponseRequestStatusStatusCode = "CONFLICT"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity        AdminUserCRMSyncResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests            AdminUserCRMSyncResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage        AdminUserCRMSyncResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable         AdminUserCRMSyncResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge            AdminUserCRMSyncResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable              AdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons AdminUserCRMSyncResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway                 AdminUserCRMSyncResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AdminUserCRMSyncResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AdminUserCRMSyncResponseRequestStatusStatusCodeUnknown, AdminUserCRMSyncResponseRequestStatusStatusCodeSuccess, AdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized, AdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired, AdminUserCRMSyncResponseRequestStatusStatusCodeForbidden, AdminUserCRMSyncResponseRequestStatusStatusCodeTimeout, AdminUserCRMSyncResponseRequestStatusStatusCodeExists, AdminUserCRMSyncResponseRequestStatusStatusCodeNotFound, AdminUserCRMSyncResponseRequestStatusStatusCodeInternalError, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData, AdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed, AdminUserCRMSyncResponseRequestStatusStatusCodeConflict, AdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity, AdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests, AdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage, AdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable, AdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge, AdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable, AdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons, AdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
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

type AdminUserGetParams struct {
	OrgName param.Field[string] `query:"org-name"`
}

// URLQuery serializes [AdminUserGetParams]'s query parameters as `url.Values`.
func (r AdminUserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
