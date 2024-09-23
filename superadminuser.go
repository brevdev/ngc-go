// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
	"github.com/stainless-sdks/ngc-go/shared"
)

// SuperAdminUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminUserService] method instead.
type SuperAdminUserService struct {
	Options []option.RequestOption
	Orgs    *SuperAdminUserOrgService
}

// NewSuperAdminUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperAdminUserService(opts ...option.RequestOption) (r *SuperAdminUserService) {
	r = &SuperAdminUserService{}
	r.Options = opts
	r.Orgs = NewSuperAdminUserOrgService(opts...)
	return
}

// Sync crm id with user email (Super Admin privileges required)
func (r *SuperAdminUserService) CRMSync(ctx context.Context, id int64, opts ...option.RequestOption) (res *SuperAdminUserCRMSyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v2/admin/users/%v/crm-sync", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Migrate User Deprecated Roles.
func (r *SuperAdminUserService) MigrateDeprecatedRoles(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/users/%s/migrate-deprecated-roles", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SuperAdminUserCRMSyncResponse struct {
	RequestStatus SuperAdminUserCRMSyncResponseRequestStatus `json:"requestStatus"`
	JSON          superAdminUserCRMSyncResponseJSON          `json:"-"`
}

// superAdminUserCRMSyncResponseJSON contains the JSON metadata for the struct
// [SuperAdminUserCRMSyncResponse]
type superAdminUserCRMSyncResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SuperAdminUserCRMSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserCRMSyncResponseJSON) RawJSON() string {
	return r.raw
}

type SuperAdminUserCRMSyncResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        SuperAdminUserCRMSyncResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                               `json:"statusDescription"`
	JSON              superAdminUserCRMSyncResponseRequestStatusJSON       `json:"-"`
}

// superAdminUserCRMSyncResponseRequestStatusJSON contains the JSON metadata for
// the struct [SuperAdminUserCRMSyncResponseRequestStatus]
type superAdminUserCRMSyncResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuperAdminUserCRMSyncResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserCRMSyncResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type SuperAdminUserCRMSyncResponseRequestStatusStatusCode string

const (
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnknown                    SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "UNKNOWN"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeSuccess                    SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "SUCCESS"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized               SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "UNAUTHORIZED"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired            SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeForbidden                  SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "FORBIDDEN"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeTimeout                    SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "TIMEOUT"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeExists                     SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "EXISTS"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeNotFound                   SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_FOUND"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInternalError              SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest             SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion      SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData         SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed           SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeConflict                   SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "CONFLICT"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity        SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests            SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage        SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable         SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge            SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable              SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	SuperAdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway                 SuperAdminUserCRMSyncResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r SuperAdminUserCRMSyncResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnknown, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeSuccess, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized, SuperAdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeForbidden, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeTimeout, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeExists, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeNotFound, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInternalError, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeConflict, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable, SuperAdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons, SuperAdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
