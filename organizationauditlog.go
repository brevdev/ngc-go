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
)

// OrganizationAuditLogService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationAuditLogService] method instead.
type OrganizationAuditLogService struct {
	Options []option.RequestOption
}

// NewOrganizationAuditLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationAuditLogService(opts ...option.RequestOption) (r *OrganizationAuditLogService) {
	r = &OrganizationAuditLogService{}
	r.Options = opts
	return
}

// List audit logs
func (r *OrganizationAuditLogService) List(ctx context.Context, orgName string, opts ...option.RequestOption) (res *AuditLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/auditLogs", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete audit logs
func (r *OrganizationAuditLogService) Delete(ctx context.Context, orgName string, body OrganizationAuditLogDeleteParams, opts ...option.RequestOption) (res *OrganizationAuditLogDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/auditLogs", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Request audit logs
func (r *OrganizationAuditLogService) Request(ctx context.Context, orgName string, body OrganizationAuditLogRequestParams, opts ...option.RequestOption) (res *OrganizationAuditLogRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/auditLogs", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List of audit logs object
type AuditLogsResponse struct {
	// Array of auditLogs object
	AuditLogsList []AuditLogsResponseAuditLogsList `json:"auditLogsList"`
	RequestStatus AuditLogsResponseRequestStatus   `json:"requestStatus"`
	JSON          auditLogsResponseJSON            `json:"-"`
}

// auditLogsResponseJSON contains the JSON metadata for the struct
// [AuditLogsResponse]
type auditLogsResponseJSON struct {
	AuditLogsList apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuditLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsResponseJSON) RawJSON() string {
	return r.raw
}

// Object of audit logs
type AuditLogsResponseAuditLogsList struct {
	// Audit logs from date
	AuditLogsFrom string `json:"auditLogsFrom"`
	// Unique Id of audit logs
	AuditLogsID string `json:"auditLogsId"`
	// Status of audit logs
	AuditLogsStatus AuditLogsResponseAuditLogsListAuditLogsStatus `json:"auditLogsStatus"`
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
	RequsterName string                             `json:"requsterName"`
	JSON         auditLogsResponseAuditLogsListJSON `json:"-"`
}

// auditLogsResponseAuditLogsListJSON contains the JSON metadata for the struct
// [AuditLogsResponseAuditLogsList]
type auditLogsResponseAuditLogsListJSON struct {
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

func (r *AuditLogsResponseAuditLogsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsResponseAuditLogsListJSON) RawJSON() string {
	return r.raw
}

// Status of audit logs
type AuditLogsResponseAuditLogsListAuditLogsStatus string

const (
	AuditLogsResponseAuditLogsListAuditLogsStatusUnknown   AuditLogsResponseAuditLogsListAuditLogsStatus = "UNKNOWN"
	AuditLogsResponseAuditLogsListAuditLogsStatusRequested AuditLogsResponseAuditLogsListAuditLogsStatus = "REQUESTED"
	AuditLogsResponseAuditLogsListAuditLogsStatusReady     AuditLogsResponseAuditLogsListAuditLogsStatus = "READY"
)

func (r AuditLogsResponseAuditLogsListAuditLogsStatus) IsKnown() bool {
	switch r {
	case AuditLogsResponseAuditLogsListAuditLogsStatusUnknown, AuditLogsResponseAuditLogsListAuditLogsStatusRequested, AuditLogsResponseAuditLogsListAuditLogsStatusReady:
		return true
	}
	return false
}

type AuditLogsResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AuditLogsResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                   `json:"statusDescription"`
	JSON              auditLogsResponseRequestStatusJSON       `json:"-"`
}

// auditLogsResponseRequestStatusJSON contains the JSON metadata for the struct
// [AuditLogsResponseRequestStatus]
type auditLogsResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuditLogsResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogsResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AuditLogsResponseRequestStatusStatusCode string

const (
	AuditLogsResponseRequestStatusStatusCodeUnknown                    AuditLogsResponseRequestStatusStatusCode = "UNKNOWN"
	AuditLogsResponseRequestStatusStatusCodeSuccess                    AuditLogsResponseRequestStatusStatusCode = "SUCCESS"
	AuditLogsResponseRequestStatusStatusCodeUnauthorized               AuditLogsResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AuditLogsResponseRequestStatusStatusCodePaymentRequired            AuditLogsResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AuditLogsResponseRequestStatusStatusCodeForbidden                  AuditLogsResponseRequestStatusStatusCode = "FORBIDDEN"
	AuditLogsResponseRequestStatusStatusCodeTimeout                    AuditLogsResponseRequestStatusStatusCode = "TIMEOUT"
	AuditLogsResponseRequestStatusStatusCodeExists                     AuditLogsResponseRequestStatusStatusCode = "EXISTS"
	AuditLogsResponseRequestStatusStatusCodeNotFound                   AuditLogsResponseRequestStatusStatusCode = "NOT_FOUND"
	AuditLogsResponseRequestStatusStatusCodeInternalError              AuditLogsResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AuditLogsResponseRequestStatusStatusCodeInvalidRequest             AuditLogsResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AuditLogsResponseRequestStatusStatusCodeInvalidRequestVersion      AuditLogsResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AuditLogsResponseRequestStatusStatusCodeInvalidRequestData         AuditLogsResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AuditLogsResponseRequestStatusStatusCodeMethodNotAllowed           AuditLogsResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AuditLogsResponseRequestStatusStatusCodeConflict                   AuditLogsResponseRequestStatusStatusCode = "CONFLICT"
	AuditLogsResponseRequestStatusStatusCodeUnprocessableEntity        AuditLogsResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AuditLogsResponseRequestStatusStatusCodeTooManyRequests            AuditLogsResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AuditLogsResponseRequestStatusStatusCodeInsufficientStorage        AuditLogsResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AuditLogsResponseRequestStatusStatusCodeServiceUnavailable         AuditLogsResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AuditLogsResponseRequestStatusStatusCodePayloadTooLarge            AuditLogsResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AuditLogsResponseRequestStatusStatusCodeNotAcceptable              AuditLogsResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AuditLogsResponseRequestStatusStatusCodeUnavailableForLegalReasons AuditLogsResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AuditLogsResponseRequestStatusStatusCodeBadGateway                 AuditLogsResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AuditLogsResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AuditLogsResponseRequestStatusStatusCodeUnknown, AuditLogsResponseRequestStatusStatusCodeSuccess, AuditLogsResponseRequestStatusStatusCodeUnauthorized, AuditLogsResponseRequestStatusStatusCodePaymentRequired, AuditLogsResponseRequestStatusStatusCodeForbidden, AuditLogsResponseRequestStatusStatusCodeTimeout, AuditLogsResponseRequestStatusStatusCodeExists, AuditLogsResponseRequestStatusStatusCodeNotFound, AuditLogsResponseRequestStatusStatusCodeInternalError, AuditLogsResponseRequestStatusStatusCodeInvalidRequest, AuditLogsResponseRequestStatusStatusCodeInvalidRequestVersion, AuditLogsResponseRequestStatusStatusCodeInvalidRequestData, AuditLogsResponseRequestStatusStatusCodeMethodNotAllowed, AuditLogsResponseRequestStatusStatusCodeConflict, AuditLogsResponseRequestStatusStatusCodeUnprocessableEntity, AuditLogsResponseRequestStatusStatusCodeTooManyRequests, AuditLogsResponseRequestStatusStatusCodeInsufficientStorage, AuditLogsResponseRequestStatusStatusCodeServiceUnavailable, AuditLogsResponseRequestStatusStatusCodePayloadTooLarge, AuditLogsResponseRequestStatusStatusCodeNotAcceptable, AuditLogsResponseRequestStatusStatusCodeUnavailableForLegalReasons, AuditLogsResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrganizationAuditLogDeleteResponse struct {
	RequestStatus OrganizationAuditLogDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          organizationAuditLogDeleteResponseJSON          `json:"-"`
}

// organizationAuditLogDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationAuditLogDeleteResponse]
type organizationAuditLogDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationAuditLogDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAuditLogDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAuditLogDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrganizationAuditLogDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                    `json:"statusDescription"`
	JSON              organizationAuditLogDeleteResponseRequestStatusJSON       `json:"-"`
}

// organizationAuditLogDeleteResponseRequestStatusJSON contains the JSON metadata
// for the struct [OrganizationAuditLogDeleteResponseRequestStatus]
type organizationAuditLogDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationAuditLogDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAuditLogDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrganizationAuditLogDeleteResponseRequestStatusStatusCode string

const (
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnknown                    OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeSuccess                    OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnauthorized               OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodePaymentRequired            OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeForbidden                  OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeTimeout                    OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeExists                     OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeNotFound                   OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInternalError              OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequest             OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeConflict                   OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeTooManyRequests            OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeNotAcceptable              OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrganizationAuditLogDeleteResponseRequestStatusStatusCodeBadGateway                 OrganizationAuditLogDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrganizationAuditLogDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnknown, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeSuccess, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnauthorized, OrganizationAuditLogDeleteResponseRequestStatusStatusCodePaymentRequired, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeForbidden, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeTimeout, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeExists, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeNotFound, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInternalError, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequest, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeConflict, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeTooManyRequests, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrganizationAuditLogDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeNotAcceptable, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrganizationAuditLogDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrganizationAuditLogRequestResponse struct {
	RequestStatus OrganizationAuditLogRequestResponseRequestStatus `json:"requestStatus"`
	JSON          organizationAuditLogRequestResponseJSON          `json:"-"`
}

// organizationAuditLogRequestResponseJSON contains the JSON metadata for the
// struct [OrganizationAuditLogRequestResponse]
type organizationAuditLogRequestResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationAuditLogRequestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAuditLogRequestResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAuditLogRequestResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrganizationAuditLogRequestResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                     `json:"statusDescription"`
	JSON              organizationAuditLogRequestResponseRequestStatusJSON       `json:"-"`
}

// organizationAuditLogRequestResponseRequestStatusJSON contains the JSON metadata
// for the struct [OrganizationAuditLogRequestResponseRequestStatus]
type organizationAuditLogRequestResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationAuditLogRequestResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAuditLogRequestResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrganizationAuditLogRequestResponseRequestStatusStatusCode string

const (
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnknown                    OrganizationAuditLogRequestResponseRequestStatusStatusCode = "UNKNOWN"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeSuccess                    OrganizationAuditLogRequestResponseRequestStatusStatusCode = "SUCCESS"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnauthorized               OrganizationAuditLogRequestResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodePaymentRequired            OrganizationAuditLogRequestResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeForbidden                  OrganizationAuditLogRequestResponseRequestStatusStatusCode = "FORBIDDEN"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeTimeout                    OrganizationAuditLogRequestResponseRequestStatusStatusCode = "TIMEOUT"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeExists                     OrganizationAuditLogRequestResponseRequestStatusStatusCode = "EXISTS"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeNotFound                   OrganizationAuditLogRequestResponseRequestStatusStatusCode = "NOT_FOUND"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeInternalError              OrganizationAuditLogRequestResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequest             OrganizationAuditLogRequestResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequestVersion      OrganizationAuditLogRequestResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequestData         OrganizationAuditLogRequestResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeMethodNotAllowed           OrganizationAuditLogRequestResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeConflict                   OrganizationAuditLogRequestResponseRequestStatusStatusCode = "CONFLICT"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnprocessableEntity        OrganizationAuditLogRequestResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeTooManyRequests            OrganizationAuditLogRequestResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeInsufficientStorage        OrganizationAuditLogRequestResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeServiceUnavailable         OrganizationAuditLogRequestResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodePayloadTooLarge            OrganizationAuditLogRequestResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeNotAcceptable              OrganizationAuditLogRequestResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnavailableForLegalReasons OrganizationAuditLogRequestResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrganizationAuditLogRequestResponseRequestStatusStatusCodeBadGateway                 OrganizationAuditLogRequestResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrganizationAuditLogRequestResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnknown, OrganizationAuditLogRequestResponseRequestStatusStatusCodeSuccess, OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnauthorized, OrganizationAuditLogRequestResponseRequestStatusStatusCodePaymentRequired, OrganizationAuditLogRequestResponseRequestStatusStatusCodeForbidden, OrganizationAuditLogRequestResponseRequestStatusStatusCodeTimeout, OrganizationAuditLogRequestResponseRequestStatusStatusCodeExists, OrganizationAuditLogRequestResponseRequestStatusStatusCodeNotFound, OrganizationAuditLogRequestResponseRequestStatusStatusCodeInternalError, OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequest, OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequestVersion, OrganizationAuditLogRequestResponseRequestStatusStatusCodeInvalidRequestData, OrganizationAuditLogRequestResponseRequestStatusStatusCodeMethodNotAllowed, OrganizationAuditLogRequestResponseRequestStatusStatusCodeConflict, OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnprocessableEntity, OrganizationAuditLogRequestResponseRequestStatusStatusCodeTooManyRequests, OrganizationAuditLogRequestResponseRequestStatusStatusCodeInsufficientStorage, OrganizationAuditLogRequestResponseRequestStatusStatusCodeServiceUnavailable, OrganizationAuditLogRequestResponseRequestStatusStatusCodePayloadTooLarge, OrganizationAuditLogRequestResponseRequestStatusStatusCodeNotAcceptable, OrganizationAuditLogRequestResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrganizationAuditLogRequestResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrganizationAuditLogDeleteParams struct {
	LogIDs param.Field[[]string] `query:"logIds,required"`
}

// URLQuery serializes [OrganizationAuditLogDeleteParams]'s query parameters as
// `url.Values`.
func (r OrganizationAuditLogDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationAuditLogRequestParams struct {
	// Audit logs from date in ISO-8601 format
	AuditLogsFrom param.Field[string] `json:"auditLogsFrom,required"`
	// Audit logs to date in ISO-8601 format
	AuditLogsTo param.Field[string] `json:"auditLogsTo,required"`
}

func (r OrganizationAuditLogRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
