// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// OrgAuditLogService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
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

// List of audit logs object
type AuditLogs struct {
	// Array of auditLogs object
	AuditLogsList []AuditLogsAuditLogsList `json:"auditLogsList"`
	RequestStatus AuditLogsRequestStatus   `json:"requestStatus"`
	JSON          auditLogsJSON            `json:"-"`
}

// auditLogsJSON contains the JSON metadata for the struct [AuditLogs]
type auditLogsJSON struct {
	AuditLogsList apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuditLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsJSON) RawJSON() string {
	return r.raw
}

// Object of audit logs
type AuditLogsAuditLogsList struct {
	// Audit logs from date
	AuditLogsFrom string `json:"auditLogsFrom"`
	// Unique Id of audit logs
	AuditLogsID string `json:"auditLogsId"`
	// Status of audit logs
	AuditLogsStatus AuditLogsAuditLogsListAuditLogsStatus `json:"auditLogsStatus"`
	// Audit logs to date
	AuditLogsTo string `json:"auditLogsTo"`
	// Audit logs requested date
	RequestedDate string `json:"requestedDate"`
	// Audit logs requester email
	RequesterEmail string `json:"requesterEmail"`
	// Audit logs requester name
	RequesterName string `json:"requesterName"`
	// [DEPRECATED] Audit logs requester email
	RequsterEmail string `json:"requsterEmail"`
	// [DEPRECATED] Audit logs requester name
	RequsterName string                     `json:"requsterName"`
	JSON         auditLogsAuditLogsListJSON `json:"-"`
}

// auditLogsAuditLogsListJSON contains the JSON metadata for the struct
// [AuditLogsAuditLogsList]
type auditLogsAuditLogsListJSON struct {
	AuditLogsFrom   apijson.Field
	AuditLogsID     apijson.Field
	AuditLogsStatus apijson.Field
	AuditLogsTo     apijson.Field
	RequestedDate   apijson.Field
	RequesterEmail  apijson.Field
	RequesterName   apijson.Field
	RequsterEmail   apijson.Field
	RequsterName    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AuditLogsAuditLogsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsAuditLogsListJSON) RawJSON() string {
	return r.raw
}

// Status of audit logs
type AuditLogsAuditLogsListAuditLogsStatus string

const (
	AuditLogsAuditLogsListAuditLogsStatusUnknown   AuditLogsAuditLogsListAuditLogsStatus = "UNKNOWN"
	AuditLogsAuditLogsListAuditLogsStatusRequested AuditLogsAuditLogsListAuditLogsStatus = "REQUESTED"
	AuditLogsAuditLogsListAuditLogsStatusReady     AuditLogsAuditLogsListAuditLogsStatus = "READY"
)

func (r AuditLogsAuditLogsListAuditLogsStatus) IsKnown() bool {
	switch r {
	case AuditLogsAuditLogsListAuditLogsStatusUnknown, AuditLogsAuditLogsListAuditLogsStatusRequested, AuditLogsAuditLogsListAuditLogsStatusReady:
		return true
	}
	return false
}

type AuditLogsRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AuditLogsRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                           `json:"statusDescription"`
	JSON              auditLogsRequestStatusJSON       `json:"-"`
}

// auditLogsRequestStatusJSON contains the JSON metadata for the struct
// [AuditLogsRequestStatus]
type auditLogsRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuditLogsRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AuditLogsRequestStatusStatusCode string

const (
	AuditLogsRequestStatusStatusCodeUnknown                    AuditLogsRequestStatusStatusCode = "UNKNOWN"
	AuditLogsRequestStatusStatusCodeSuccess                    AuditLogsRequestStatusStatusCode = "SUCCESS"
	AuditLogsRequestStatusStatusCodeUnauthorized               AuditLogsRequestStatusStatusCode = "UNAUTHORIZED"
	AuditLogsRequestStatusStatusCodePaymentRequired            AuditLogsRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AuditLogsRequestStatusStatusCodeForbidden                  AuditLogsRequestStatusStatusCode = "FORBIDDEN"
	AuditLogsRequestStatusStatusCodeTimeout                    AuditLogsRequestStatusStatusCode = "TIMEOUT"
	AuditLogsRequestStatusStatusCodeExists                     AuditLogsRequestStatusStatusCode = "EXISTS"
	AuditLogsRequestStatusStatusCodeNotFound                   AuditLogsRequestStatusStatusCode = "NOT_FOUND"
	AuditLogsRequestStatusStatusCodeInternalError              AuditLogsRequestStatusStatusCode = "INTERNAL_ERROR"
	AuditLogsRequestStatusStatusCodeInvalidRequest             AuditLogsRequestStatusStatusCode = "INVALID_REQUEST"
	AuditLogsRequestStatusStatusCodeInvalidRequestVersion      AuditLogsRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AuditLogsRequestStatusStatusCodeInvalidRequestData         AuditLogsRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AuditLogsRequestStatusStatusCodeMethodNotAllowed           AuditLogsRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AuditLogsRequestStatusStatusCodeConflict                   AuditLogsRequestStatusStatusCode = "CONFLICT"
	AuditLogsRequestStatusStatusCodeUnprocessableEntity        AuditLogsRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AuditLogsRequestStatusStatusCodeTooManyRequests            AuditLogsRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AuditLogsRequestStatusStatusCodeInsufficientStorage        AuditLogsRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AuditLogsRequestStatusStatusCodeServiceUnavailable         AuditLogsRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AuditLogsRequestStatusStatusCodePayloadTooLarge            AuditLogsRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AuditLogsRequestStatusStatusCodeNotAcceptable              AuditLogsRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AuditLogsRequestStatusStatusCodeUnavailableForLegalReasons AuditLogsRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AuditLogsRequestStatusStatusCodeBadGateway                 AuditLogsRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AuditLogsRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AuditLogsRequestStatusStatusCodeUnknown, AuditLogsRequestStatusStatusCodeSuccess, AuditLogsRequestStatusStatusCodeUnauthorized, AuditLogsRequestStatusStatusCodePaymentRequired, AuditLogsRequestStatusStatusCodeForbidden, AuditLogsRequestStatusStatusCodeTimeout, AuditLogsRequestStatusStatusCodeExists, AuditLogsRequestStatusStatusCodeNotFound, AuditLogsRequestStatusStatusCodeInternalError, AuditLogsRequestStatusStatusCodeInvalidRequest, AuditLogsRequestStatusStatusCodeInvalidRequestVersion, AuditLogsRequestStatusStatusCodeInvalidRequestData, AuditLogsRequestStatusStatusCodeMethodNotAllowed, AuditLogsRequestStatusStatusCodeConflict, AuditLogsRequestStatusStatusCodeUnprocessableEntity, AuditLogsRequestStatusStatusCodeTooManyRequests, AuditLogsRequestStatusStatusCodeInsufficientStorage, AuditLogsRequestStatusStatusCodeServiceUnavailable, AuditLogsRequestStatusStatusCodePayloadTooLarge, AuditLogsRequestStatusStatusCodeNotAcceptable, AuditLogsRequestStatusStatusCodeUnavailableForLegalReasons, AuditLogsRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
