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

// OrgAuditLogService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgAuditLogService] method instead.
type OrgAuditLogService struct {
	Options []option.RequestOption
}

// NewOrgAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgAuditLogService(opts ...option.RequestOption) (r *OrgAuditLogService) {
	r = &OrgAuditLogService{}
	r.Options = opts
	return
}

// Get downloable link for audit logs
func (r *OrgAuditLogService) Get(ctx context.Context, orgName string, logID string, opts ...option.RequestOption) (res *AuditLogsPresignedURL, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if logID == "" {
		err = errors.New("missing required log-id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/auditLogs/%s", orgName, logID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of audit logs object
type AuditLogsPresignedURL struct {
	// Presign url for user to download audit logs
	AuditLogsPresignedURL string                             `json:"auditLogsPresignedUrl"`
	RequestStatus         AuditLogsPresignedURLRequestStatus `json:"requestStatus"`
	JSON                  auditLogsPresignedURLJSON          `json:"-"`
}

// auditLogsPresignedURLJSON contains the JSON metadata for the struct
// [AuditLogsPresignedURL]
type auditLogsPresignedURLJSON struct {
	AuditLogsPresignedURL apijson.Field
	RequestStatus         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AuditLogsPresignedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsPresignedURLJSON) RawJSON() string {
	return r.raw
}

type AuditLogsPresignedURLRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AuditLogsPresignedURLRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                       `json:"statusDescription"`
	JSON              auditLogsPresignedURLRequestStatusJSON       `json:"-"`
}

// auditLogsPresignedURLRequestStatusJSON contains the JSON metadata for the struct
// [AuditLogsPresignedURLRequestStatus]
type auditLogsPresignedURLRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuditLogsPresignedURLRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsPresignedURLRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AuditLogsPresignedURLRequestStatusStatusCode string

const (
	AuditLogsPresignedURLRequestStatusStatusCodeUnknown                    AuditLogsPresignedURLRequestStatusStatusCode = "UNKNOWN"
	AuditLogsPresignedURLRequestStatusStatusCodeSuccess                    AuditLogsPresignedURLRequestStatusStatusCode = "SUCCESS"
	AuditLogsPresignedURLRequestStatusStatusCodeUnauthorized               AuditLogsPresignedURLRequestStatusStatusCode = "UNAUTHORIZED"
	AuditLogsPresignedURLRequestStatusStatusCodePaymentRequired            AuditLogsPresignedURLRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AuditLogsPresignedURLRequestStatusStatusCodeForbidden                  AuditLogsPresignedURLRequestStatusStatusCode = "FORBIDDEN"
	AuditLogsPresignedURLRequestStatusStatusCodeTimeout                    AuditLogsPresignedURLRequestStatusStatusCode = "TIMEOUT"
	AuditLogsPresignedURLRequestStatusStatusCodeExists                     AuditLogsPresignedURLRequestStatusStatusCode = "EXISTS"
	AuditLogsPresignedURLRequestStatusStatusCodeNotFound                   AuditLogsPresignedURLRequestStatusStatusCode = "NOT_FOUND"
	AuditLogsPresignedURLRequestStatusStatusCodeInternalError              AuditLogsPresignedURLRequestStatusStatusCode = "INTERNAL_ERROR"
	AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequest             AuditLogsPresignedURLRequestStatusStatusCode = "INVALID_REQUEST"
	AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequestVersion      AuditLogsPresignedURLRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequestData         AuditLogsPresignedURLRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AuditLogsPresignedURLRequestStatusStatusCodeMethodNotAllowed           AuditLogsPresignedURLRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AuditLogsPresignedURLRequestStatusStatusCodeConflict                   AuditLogsPresignedURLRequestStatusStatusCode = "CONFLICT"
	AuditLogsPresignedURLRequestStatusStatusCodeUnprocessableEntity        AuditLogsPresignedURLRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AuditLogsPresignedURLRequestStatusStatusCodeTooManyRequests            AuditLogsPresignedURLRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AuditLogsPresignedURLRequestStatusStatusCodeInsufficientStorage        AuditLogsPresignedURLRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AuditLogsPresignedURLRequestStatusStatusCodeServiceUnavailable         AuditLogsPresignedURLRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AuditLogsPresignedURLRequestStatusStatusCodePayloadTooLarge            AuditLogsPresignedURLRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AuditLogsPresignedURLRequestStatusStatusCodeNotAcceptable              AuditLogsPresignedURLRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AuditLogsPresignedURLRequestStatusStatusCodeUnavailableForLegalReasons AuditLogsPresignedURLRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AuditLogsPresignedURLRequestStatusStatusCodeBadGateway                 AuditLogsPresignedURLRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AuditLogsPresignedURLRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AuditLogsPresignedURLRequestStatusStatusCodeUnknown, AuditLogsPresignedURLRequestStatusStatusCodeSuccess, AuditLogsPresignedURLRequestStatusStatusCodeUnauthorized, AuditLogsPresignedURLRequestStatusStatusCodePaymentRequired, AuditLogsPresignedURLRequestStatusStatusCodeForbidden, AuditLogsPresignedURLRequestStatusStatusCodeTimeout, AuditLogsPresignedURLRequestStatusStatusCodeExists, AuditLogsPresignedURLRequestStatusStatusCodeNotFound, AuditLogsPresignedURLRequestStatusStatusCodeInternalError, AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequest, AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequestVersion, AuditLogsPresignedURLRequestStatusStatusCodeInvalidRequestData, AuditLogsPresignedURLRequestStatusStatusCodeMethodNotAllowed, AuditLogsPresignedURLRequestStatusStatusCodeConflict, AuditLogsPresignedURLRequestStatusStatusCodeUnprocessableEntity, AuditLogsPresignedURLRequestStatusStatusCodeTooManyRequests, AuditLogsPresignedURLRequestStatusStatusCodeInsufficientStorage, AuditLogsPresignedURLRequestStatusStatusCodeServiceUnavailable, AuditLogsPresignedURLRequestStatusStatusCodePayloadTooLarge, AuditLogsPresignedURLRequestStatusStatusCodeNotAcceptable, AuditLogsPresignedURLRequestStatusStatusCodeUnavailableForLegalReasons, AuditLogsPresignedURLRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
