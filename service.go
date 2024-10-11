// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// ServiceService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
	Version *ServiceVersionService
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	r.Version = NewServiceVersionService(opts...)
	return
}

// Used to check the health status of the web service only
func (r *ServiceService) Health(ctx context.Context, opts ...option.RequestOption) (res *Health, err error) {
	opts = append(r.Options[:], opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Used to get health status of all services
func (r *ServiceService) HealthAll(ctx context.Context, query ServiceHealthAllParams, opts ...option.RequestOption) (res *Health, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This API is invoked by monitoring tools, other services and infrastructure to
// retrieve health status the targeted service, this is unprotected method
type Health struct {
	// object that describes health of the service
	Health        HealthHealth        `json:"health"`
	RequestStatus HealthRequestStatus `json:"requestStatus"`
	JSON          healthJSON          `json:"-"`
}

// healthJSON contains the JSON metadata for the struct [Health]
type healthJSON struct {
	Health        apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Health) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthJSON) RawJSON() string {
	return r.raw
}

// object that describes health of the service
type HealthHealth struct {
	// Enum that describes health of the service
	HealthCode HealthHealthHealthCode `json:"healthCode"`
	// Human readable description
	HealthCodeDescription string                 `json:"healthCodeDescription"`
	MetaData              []HealthHealthMetaData `json:"metaData"`
	JSON                  healthHealthJSON       `json:"-"`
}

// healthHealthJSON contains the JSON metadata for the struct [HealthHealth]
type healthHealthJSON struct {
	HealthCode            apijson.Field
	HealthCodeDescription apijson.Field
	MetaData              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *HealthHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthJSON) RawJSON() string {
	return r.raw
}

// Enum that describes health of the service
type HealthHealthHealthCode string

const (
	HealthHealthHealthCodeUnknown          HealthHealthHealthCode = "UNKNOWN"
	HealthHealthHealthCodeOk               HealthHealthHealthCode = "OK"
	HealthHealthHealthCodeUnderMaintenance HealthHealthHealthCode = "UNDER_MAINTENANCE"
	HealthHealthHealthCodeWarning          HealthHealthHealthCode = "WARNING"
	HealthHealthHealthCodeFailed           HealthHealthHealthCode = "FAILED"
)

func (r HealthHealthHealthCode) IsKnown() bool {
	switch r {
	case HealthHealthHealthCodeUnknown, HealthHealthHealthCodeOk, HealthHealthHealthCodeUnderMaintenance, HealthHealthHealthCodeWarning, HealthHealthHealthCodeFailed:
		return true
	}
	return false
}

type HealthHealthMetaData struct {
	Key   string                   `json:"key"`
	Value string                   `json:"value"`
	JSON  healthHealthMetaDataJSON `json:"-"`
}

// healthHealthMetaDataJSON contains the JSON metadata for the struct
// [HealthHealthMetaData]
type healthHealthMetaDataJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthHealthMetaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthMetaDataJSON) RawJSON() string {
	return r.raw
}

type HealthRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        HealthRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                        `json:"statusDescription"`
	JSON              healthRequestStatusJSON       `json:"-"`
}

// healthRequestStatusJSON contains the JSON metadata for the struct
// [HealthRequestStatus]
type healthRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HealthRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type HealthRequestStatusStatusCode string

const (
	HealthRequestStatusStatusCodeUnknown                    HealthRequestStatusStatusCode = "UNKNOWN"
	HealthRequestStatusStatusCodeSuccess                    HealthRequestStatusStatusCode = "SUCCESS"
	HealthRequestStatusStatusCodeUnauthorized               HealthRequestStatusStatusCode = "UNAUTHORIZED"
	HealthRequestStatusStatusCodePaymentRequired            HealthRequestStatusStatusCode = "PAYMENT_REQUIRED"
	HealthRequestStatusStatusCodeForbidden                  HealthRequestStatusStatusCode = "FORBIDDEN"
	HealthRequestStatusStatusCodeTimeout                    HealthRequestStatusStatusCode = "TIMEOUT"
	HealthRequestStatusStatusCodeExists                     HealthRequestStatusStatusCode = "EXISTS"
	HealthRequestStatusStatusCodeNotFound                   HealthRequestStatusStatusCode = "NOT_FOUND"
	HealthRequestStatusStatusCodeInternalError              HealthRequestStatusStatusCode = "INTERNAL_ERROR"
	HealthRequestStatusStatusCodeInvalidRequest             HealthRequestStatusStatusCode = "INVALID_REQUEST"
	HealthRequestStatusStatusCodeInvalidRequestVersion      HealthRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	HealthRequestStatusStatusCodeInvalidRequestData         HealthRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	HealthRequestStatusStatusCodeMethodNotAllowed           HealthRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	HealthRequestStatusStatusCodeConflict                   HealthRequestStatusStatusCode = "CONFLICT"
	HealthRequestStatusStatusCodeUnprocessableEntity        HealthRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	HealthRequestStatusStatusCodeTooManyRequests            HealthRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	HealthRequestStatusStatusCodeInsufficientStorage        HealthRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	HealthRequestStatusStatusCodeServiceUnavailable         HealthRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	HealthRequestStatusStatusCodePayloadTooLarge            HealthRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	HealthRequestStatusStatusCodeNotAcceptable              HealthRequestStatusStatusCode = "NOT_ACCEPTABLE"
	HealthRequestStatusStatusCodeUnavailableForLegalReasons HealthRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HealthRequestStatusStatusCodeBadGateway                 HealthRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r HealthRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case HealthRequestStatusStatusCodeUnknown, HealthRequestStatusStatusCodeSuccess, HealthRequestStatusStatusCodeUnauthorized, HealthRequestStatusStatusCodePaymentRequired, HealthRequestStatusStatusCodeForbidden, HealthRequestStatusStatusCodeTimeout, HealthRequestStatusStatusCodeExists, HealthRequestStatusStatusCodeNotFound, HealthRequestStatusStatusCodeInternalError, HealthRequestStatusStatusCodeInvalidRequest, HealthRequestStatusStatusCodeInvalidRequestVersion, HealthRequestStatusStatusCodeInvalidRequestData, HealthRequestStatusStatusCodeMethodNotAllowed, HealthRequestStatusStatusCodeConflict, HealthRequestStatusStatusCodeUnprocessableEntity, HealthRequestStatusStatusCodeTooManyRequests, HealthRequestStatusStatusCodeInsufficientStorage, HealthRequestStatusStatusCodeServiceUnavailable, HealthRequestStatusStatusCodePayloadTooLarge, HealthRequestStatusStatusCodeNotAcceptable, HealthRequestStatusStatusCodeUnavailableForLegalReasons, HealthRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type ServiceHealthAllParams struct {
	// secret value that validates the call in order to show details
	Secret param.Field[string] `query:"secret"`
	// boolean value to indicating to show details or not
	ShowDetails param.Field[bool] `query:"showDetails"`
}

// URLQuery serializes [ServiceHealthAllParams]'s query parameters as `url.Values`.
func (r ServiceHealthAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
