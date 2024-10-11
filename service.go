// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/apiquery"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// ServiceService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// Used to get the latest version number of a given package.
func (r *ServiceService) Version(ctx context.Context, query ServiceVersionParams, opts ...option.RequestOption) (res *ServiceVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// an array of versions
type ServiceVersionResponse struct {
	RequestStatus ServiceVersionResponseRequestStatus `json:"requestStatus"`
	Versions      []ServiceVersionResponseVersion     `json:"versions"`
	JSON          serviceVersionResponseJSON          `json:"-"`
}

// serviceVersionResponseJSON contains the JSON metadata for the struct
// [ServiceVersionResponse]
type serviceVersionResponseJSON struct {
	RequestStatus apijson.Field
	Versions      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceVersionResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceVersionResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        ServiceVersionResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                        `json:"statusDescription"`
	JSON              serviceVersionResponseRequestStatusJSON       `json:"-"`
}

// serviceVersionResponseRequestStatusJSON contains the JSON metadata for the
// struct [ServiceVersionResponseRequestStatus]
type serviceVersionResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ServiceVersionResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceVersionResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type ServiceVersionResponseRequestStatusStatusCode string

const (
	ServiceVersionResponseRequestStatusStatusCodeUnknown                    ServiceVersionResponseRequestStatusStatusCode = "UNKNOWN"
	ServiceVersionResponseRequestStatusStatusCodeSuccess                    ServiceVersionResponseRequestStatusStatusCode = "SUCCESS"
	ServiceVersionResponseRequestStatusStatusCodeUnauthorized               ServiceVersionResponseRequestStatusStatusCode = "UNAUTHORIZED"
	ServiceVersionResponseRequestStatusStatusCodePaymentRequired            ServiceVersionResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	ServiceVersionResponseRequestStatusStatusCodeForbidden                  ServiceVersionResponseRequestStatusStatusCode = "FORBIDDEN"
	ServiceVersionResponseRequestStatusStatusCodeTimeout                    ServiceVersionResponseRequestStatusStatusCode = "TIMEOUT"
	ServiceVersionResponseRequestStatusStatusCodeExists                     ServiceVersionResponseRequestStatusStatusCode = "EXISTS"
	ServiceVersionResponseRequestStatusStatusCodeNotFound                   ServiceVersionResponseRequestStatusStatusCode = "NOT_FOUND"
	ServiceVersionResponseRequestStatusStatusCodeInternalError              ServiceVersionResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	ServiceVersionResponseRequestStatusStatusCodeInvalidRequest             ServiceVersionResponseRequestStatusStatusCode = "INVALID_REQUEST"
	ServiceVersionResponseRequestStatusStatusCodeInvalidRequestVersion      ServiceVersionResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	ServiceVersionResponseRequestStatusStatusCodeInvalidRequestData         ServiceVersionResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	ServiceVersionResponseRequestStatusStatusCodeMethodNotAllowed           ServiceVersionResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	ServiceVersionResponseRequestStatusStatusCodeConflict                   ServiceVersionResponseRequestStatusStatusCode = "CONFLICT"
	ServiceVersionResponseRequestStatusStatusCodeUnprocessableEntity        ServiceVersionResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	ServiceVersionResponseRequestStatusStatusCodeTooManyRequests            ServiceVersionResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	ServiceVersionResponseRequestStatusStatusCodeInsufficientStorage        ServiceVersionResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	ServiceVersionResponseRequestStatusStatusCodeServiceUnavailable         ServiceVersionResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	ServiceVersionResponseRequestStatusStatusCodePayloadTooLarge            ServiceVersionResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	ServiceVersionResponseRequestStatusStatusCodeNotAcceptable              ServiceVersionResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	ServiceVersionResponseRequestStatusStatusCodeUnavailableForLegalReasons ServiceVersionResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	ServiceVersionResponseRequestStatusStatusCodeBadGateway                 ServiceVersionResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r ServiceVersionResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case ServiceVersionResponseRequestStatusStatusCodeUnknown, ServiceVersionResponseRequestStatusStatusCodeSuccess, ServiceVersionResponseRequestStatusStatusCodeUnauthorized, ServiceVersionResponseRequestStatusStatusCodePaymentRequired, ServiceVersionResponseRequestStatusStatusCodeForbidden, ServiceVersionResponseRequestStatusStatusCodeTimeout, ServiceVersionResponseRequestStatusStatusCodeExists, ServiceVersionResponseRequestStatusStatusCodeNotFound, ServiceVersionResponseRequestStatusStatusCodeInternalError, ServiceVersionResponseRequestStatusStatusCodeInvalidRequest, ServiceVersionResponseRequestStatusStatusCodeInvalidRequestVersion, ServiceVersionResponseRequestStatusStatusCodeInvalidRequestData, ServiceVersionResponseRequestStatusStatusCodeMethodNotAllowed, ServiceVersionResponseRequestStatusStatusCodeConflict, ServiceVersionResponseRequestStatusStatusCodeUnprocessableEntity, ServiceVersionResponseRequestStatusStatusCodeTooManyRequests, ServiceVersionResponseRequestStatusStatusCodeInsufficientStorage, ServiceVersionResponseRequestStatusStatusCodeServiceUnavailable, ServiceVersionResponseRequestStatusStatusCodePayloadTooLarge, ServiceVersionResponseRequestStatusStatusCodeNotAcceptable, ServiceVersionResponseRequestStatusStatusCodeUnavailableForLegalReasons, ServiceVersionResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Latest version information for a specific package
type ServiceVersionResponseVersion struct {
	// The creation date of the version
	CreatedAt string `json:"createdAt"`
	// Human readable description of the package
	Description string `json:"description"`
	// The name of the package
	Name string `json:"name"`
	// The version number
	Version string                            `json:"version"`
	JSON    serviceVersionResponseVersionJSON `json:"-"`
}

// serviceVersionResponseVersionJSON contains the JSON metadata for the struct
// [ServiceVersionResponseVersion]
type serviceVersionResponseVersionJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceVersionResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceVersionResponseVersionJSON) RawJSON() string {
	return r.raw
}

type ServiceVersionParams struct {
	// name of package
	Package param.Field[string] `query:"package"`
	// repository of package
	Repo param.Field[string] `query:"repo"`
}

// URLQuery serializes [ServiceVersionParams]'s query parameters as `url.Values`.
func (r ServiceVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
