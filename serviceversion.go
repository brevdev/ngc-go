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

// ServiceVersionService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceVersionService] method instead.
type ServiceVersionService struct {
	Options []option.RequestOption
}

// NewServiceVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServiceVersionService(opts ...option.RequestOption) (r *ServiceVersionService) {
	r = &ServiceVersionService{}
	r.Options = opts
	return
}

// Used to get the latest version number of a given package.
func (r *ServiceVersionService) Get(ctx context.Context, query ServiceVersionGetParams, opts ...option.RequestOption) (res *PackageVersionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// an array of versions
type PackageVersionList struct {
	RequestStatus PackageVersionListRequestStatus `json:"requestStatus"`
	Versions      []PackageVersionListVersion     `json:"versions"`
	JSON          packageVersionListJSON          `json:"-"`
}

// packageVersionListJSON contains the JSON metadata for the struct
// [PackageVersionList]
type packageVersionListJSON struct {
	RequestStatus apijson.Field
	Versions      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PackageVersionList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r packageVersionListJSON) RawJSON() string {
	return r.raw
}

type PackageVersionListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        PackageVersionListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                    `json:"statusDescription"`
	JSON              packageVersionListRequestStatusJSON       `json:"-"`
}

// packageVersionListRequestStatusJSON contains the JSON metadata for the struct
// [PackageVersionListRequestStatus]
type packageVersionListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PackageVersionListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r packageVersionListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type PackageVersionListRequestStatusStatusCode string

const (
	PackageVersionListRequestStatusStatusCodeUnknown                    PackageVersionListRequestStatusStatusCode = "UNKNOWN"
	PackageVersionListRequestStatusStatusCodeSuccess                    PackageVersionListRequestStatusStatusCode = "SUCCESS"
	PackageVersionListRequestStatusStatusCodeUnauthorized               PackageVersionListRequestStatusStatusCode = "UNAUTHORIZED"
	PackageVersionListRequestStatusStatusCodePaymentRequired            PackageVersionListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	PackageVersionListRequestStatusStatusCodeForbidden                  PackageVersionListRequestStatusStatusCode = "FORBIDDEN"
	PackageVersionListRequestStatusStatusCodeTimeout                    PackageVersionListRequestStatusStatusCode = "TIMEOUT"
	PackageVersionListRequestStatusStatusCodeExists                     PackageVersionListRequestStatusStatusCode = "EXISTS"
	PackageVersionListRequestStatusStatusCodeNotFound                   PackageVersionListRequestStatusStatusCode = "NOT_FOUND"
	PackageVersionListRequestStatusStatusCodeInternalError              PackageVersionListRequestStatusStatusCode = "INTERNAL_ERROR"
	PackageVersionListRequestStatusStatusCodeInvalidRequest             PackageVersionListRequestStatusStatusCode = "INVALID_REQUEST"
	PackageVersionListRequestStatusStatusCodeInvalidRequestVersion      PackageVersionListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	PackageVersionListRequestStatusStatusCodeInvalidRequestData         PackageVersionListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	PackageVersionListRequestStatusStatusCodeMethodNotAllowed           PackageVersionListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	PackageVersionListRequestStatusStatusCodeConflict                   PackageVersionListRequestStatusStatusCode = "CONFLICT"
	PackageVersionListRequestStatusStatusCodeUnprocessableEntity        PackageVersionListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	PackageVersionListRequestStatusStatusCodeTooManyRequests            PackageVersionListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	PackageVersionListRequestStatusStatusCodeInsufficientStorage        PackageVersionListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	PackageVersionListRequestStatusStatusCodeServiceUnavailable         PackageVersionListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	PackageVersionListRequestStatusStatusCodePayloadTooLarge            PackageVersionListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	PackageVersionListRequestStatusStatusCodeNotAcceptable              PackageVersionListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	PackageVersionListRequestStatusStatusCodeUnavailableForLegalReasons PackageVersionListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	PackageVersionListRequestStatusStatusCodeBadGateway                 PackageVersionListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r PackageVersionListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case PackageVersionListRequestStatusStatusCodeUnknown, PackageVersionListRequestStatusStatusCodeSuccess, PackageVersionListRequestStatusStatusCodeUnauthorized, PackageVersionListRequestStatusStatusCodePaymentRequired, PackageVersionListRequestStatusStatusCodeForbidden, PackageVersionListRequestStatusStatusCodeTimeout, PackageVersionListRequestStatusStatusCodeExists, PackageVersionListRequestStatusStatusCodeNotFound, PackageVersionListRequestStatusStatusCodeInternalError, PackageVersionListRequestStatusStatusCodeInvalidRequest, PackageVersionListRequestStatusStatusCodeInvalidRequestVersion, PackageVersionListRequestStatusStatusCodeInvalidRequestData, PackageVersionListRequestStatusStatusCodeMethodNotAllowed, PackageVersionListRequestStatusStatusCodeConflict, PackageVersionListRequestStatusStatusCodeUnprocessableEntity, PackageVersionListRequestStatusStatusCodeTooManyRequests, PackageVersionListRequestStatusStatusCodeInsufficientStorage, PackageVersionListRequestStatusStatusCodeServiceUnavailable, PackageVersionListRequestStatusStatusCodePayloadTooLarge, PackageVersionListRequestStatusStatusCodeNotAcceptable, PackageVersionListRequestStatusStatusCodeUnavailableForLegalReasons, PackageVersionListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Latest version information for a specific package
type PackageVersionListVersion struct {
	// The creation date of the version
	CreatedAt string `json:"createdAt"`
	// Human readable description of the package
	Description string `json:"description"`
	// The name of the package
	Name string `json:"name"`
	// The version number
	Version string                        `json:"version"`
	JSON    packageVersionListVersionJSON `json:"-"`
}

// packageVersionListVersionJSON contains the JSON metadata for the struct
// [PackageVersionListVersion]
type packageVersionListVersionJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PackageVersionListVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r packageVersionListVersionJSON) RawJSON() string {
	return r.raw
}

type ServiceVersionGetParams struct {
	// name of package
	Package param.Field[string] `query:"package"`
	// repository of package
	Repo param.Field[string] `query:"repo"`
}

// URLQuery serializes [ServiceVersionGetParams]'s query parameters as
// `url.Values`.
func (r ServiceVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
