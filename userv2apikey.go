// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// UserV2APIKeyService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserV2APIKeyService] method instead.
type UserV2APIKeyService struct {
	Options []option.RequestOption
}

// NewUserV2APIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserV2APIKeyService(opts ...option.RequestOption) (r *UserV2APIKeyService) {
	r = &UserV2APIKeyService{}
	r.Options = opts
	return
}

// Generate API Key
func (r *UserV2APIKeyService) New(ctx context.Context, opts ...option.RequestOption) (res *UserKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me/api-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// response to a request to access key such as docker token
type UserKeyResponse struct {
	Key              string                       `json:"key,required"`
	CloudNfsKey      string                       `json:"cloudNfsKey"`
	CloudNfsKeyPwd   string                       `json:"cloudNfsKeyPwd"`
	CloudNfsUserName string                       `json:"cloudNfsUserName"`
	RequestStatus    UserKeyResponseRequestStatus `json:"requestStatus"`
	JSON             userKeyResponseJSON          `json:"-"`
}

// userKeyResponseJSON contains the JSON metadata for the struct [UserKeyResponse]
type userKeyResponseJSON struct {
	Key              apijson.Field
	CloudNfsKey      apijson.Field
	CloudNfsKeyPwd   apijson.Field
	CloudNfsUserName apijson.Field
	RequestStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyResponseJSON) RawJSON() string {
	return r.raw
}

type UserKeyResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserKeyResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                 `json:"statusDescription"`
	JSON              userKeyResponseRequestStatusJSON       `json:"-"`
}

// userKeyResponseRequestStatusJSON contains the JSON metadata for the struct
// [UserKeyResponseRequestStatus]
type userKeyResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserKeyResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserKeyResponseRequestStatusStatusCode string

const (
	UserKeyResponseRequestStatusStatusCodeUnknown                    UserKeyResponseRequestStatusStatusCode = "UNKNOWN"
	UserKeyResponseRequestStatusStatusCodeSuccess                    UserKeyResponseRequestStatusStatusCode = "SUCCESS"
	UserKeyResponseRequestStatusStatusCodeUnauthorized               UserKeyResponseRequestStatusStatusCode = "UNAUTHORIZED"
	UserKeyResponseRequestStatusStatusCodePaymentRequired            UserKeyResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserKeyResponseRequestStatusStatusCodeForbidden                  UserKeyResponseRequestStatusStatusCode = "FORBIDDEN"
	UserKeyResponseRequestStatusStatusCodeTimeout                    UserKeyResponseRequestStatusStatusCode = "TIMEOUT"
	UserKeyResponseRequestStatusStatusCodeExists                     UserKeyResponseRequestStatusStatusCode = "EXISTS"
	UserKeyResponseRequestStatusStatusCodeNotFound                   UserKeyResponseRequestStatusStatusCode = "NOT_FOUND"
	UserKeyResponseRequestStatusStatusCodeInternalError              UserKeyResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	UserKeyResponseRequestStatusStatusCodeInvalidRequest             UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST"
	UserKeyResponseRequestStatusStatusCodeInvalidRequestVersion      UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserKeyResponseRequestStatusStatusCodeInvalidRequestData         UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserKeyResponseRequestStatusStatusCodeMethodNotAllowed           UserKeyResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserKeyResponseRequestStatusStatusCodeConflict                   UserKeyResponseRequestStatusStatusCode = "CONFLICT"
	UserKeyResponseRequestStatusStatusCodeUnprocessableEntity        UserKeyResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserKeyResponseRequestStatusStatusCodeTooManyRequests            UserKeyResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserKeyResponseRequestStatusStatusCodeInsufficientStorage        UserKeyResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserKeyResponseRequestStatusStatusCodeServiceUnavailable         UserKeyResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserKeyResponseRequestStatusStatusCodePayloadTooLarge            UserKeyResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserKeyResponseRequestStatusStatusCodeNotAcceptable              UserKeyResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserKeyResponseRequestStatusStatusCodeUnavailableForLegalReasons UserKeyResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserKeyResponseRequestStatusStatusCodeBadGateway                 UserKeyResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserKeyResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserKeyResponseRequestStatusStatusCodeUnknown, UserKeyResponseRequestStatusStatusCodeSuccess, UserKeyResponseRequestStatusStatusCodeUnauthorized, UserKeyResponseRequestStatusStatusCodePaymentRequired, UserKeyResponseRequestStatusStatusCodeForbidden, UserKeyResponseRequestStatusStatusCodeTimeout, UserKeyResponseRequestStatusStatusCodeExists, UserKeyResponseRequestStatusStatusCodeNotFound, UserKeyResponseRequestStatusStatusCodeInternalError, UserKeyResponseRequestStatusStatusCodeInvalidRequest, UserKeyResponseRequestStatusStatusCodeInvalidRequestVersion, UserKeyResponseRequestStatusStatusCodeInvalidRequestData, UserKeyResponseRequestStatusStatusCodeMethodNotAllowed, UserKeyResponseRequestStatusStatusCodeConflict, UserKeyResponseRequestStatusStatusCodeUnprocessableEntity, UserKeyResponseRequestStatusStatusCodeTooManyRequests, UserKeyResponseRequestStatusStatusCodeInsufficientStorage, UserKeyResponseRequestStatusStatusCodeServiceUnavailable, UserKeyResponseRequestStatusStatusCodePayloadTooLarge, UserKeyResponseRequestStatusStatusCodeNotAcceptable, UserKeyResponseRequestStatusStatusCodeUnavailableForLegalReasons, UserKeyResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
