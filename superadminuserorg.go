// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// SuperAdminUserOrgService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminUserOrgService] method instead.
type SuperAdminUserOrgService struct {
	Options []option.RequestOption
	Users   *SuperAdminUserOrgUserService
	Teams   *SuperAdminUserOrgTeamService
}

// NewSuperAdminUserOrgService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperAdminUserOrgService(opts ...option.RequestOption) (r *SuperAdminUserOrgService) {
	r = &SuperAdminUserOrgService{}
	r.Options = opts
	r.Users = NewSuperAdminUserOrgUserService(opts...)
	r.Teams = NewSuperAdminUserOrgTeamService(opts...)
	return
}

// Backfill the org owner for users
func (r *SuperAdminUserOrgService) OrgOwnerBackfill(ctx context.Context, orgName string, opts ...option.RequestOption) (res *SuperAdminUserOrgOrgOwnerBackfillResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/org-owner-backfill", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SuperAdminUserOrgOrgOwnerBackfillResponse struct {
	RequestStatus SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatus `json:"requestStatus"`
	JSON          superAdminUserOrgOrgOwnerBackfillResponseJSON          `json:"-"`
}

// superAdminUserOrgOrgOwnerBackfillResponseJSON contains the JSON metadata for the
// struct [SuperAdminUserOrgOrgOwnerBackfillResponse]
type superAdminUserOrgOrgOwnerBackfillResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SuperAdminUserOrgOrgOwnerBackfillResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserOrgOrgOwnerBackfillResponseJSON) RawJSON() string {
	return r.raw
}

type SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                           `json:"statusDescription"`
	JSON              superAdminUserOrgOrgOwnerBackfillResponseRequestStatusJSON       `json:"-"`
}

// superAdminUserOrgOrgOwnerBackfillResponseRequestStatusJSON contains the JSON
// metadata for the struct [SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatus]
type superAdminUserOrgOrgOwnerBackfillResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserOrgOrgOwnerBackfillResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode string

const (
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown                    SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNKNOWN"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess                    SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "SUCCESS"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized               SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAUTHORIZED"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired            SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden                  SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "FORBIDDEN"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout                    SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "TIMEOUT"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeExists                     SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "EXISTS"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound                   SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_FOUND"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError              SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest             SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion      SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData         SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed           SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeConflict                   SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "CONFLICT"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity        SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests            SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage        SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable         SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge            SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable              SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway                 SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeExists, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeConflict, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons, SuperAdminUserOrgOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
