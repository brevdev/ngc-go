// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/apiquery"
	"github.com/stainless-sdks/ngc-go/internal/param"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// RoleService contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r *RoleService) {
	r = &RoleService{}
	r.Options = opts
	return
}

// List of roles in NGC and their scopes
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *UserRoleDefinitions, err error) {
	opts = append(r.Options[:], opts...)
	path := "roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response containing all the roles defined in NGC and their allowed actions
type UserRoleDefinitions struct {
	RequestStatus UserRoleDefinitionsRequestStatus `json:"requestStatus"`
	// List of roles
	Roles []UserRoleDefinitionsRole `json:"roles"`
	JSON  userRoleDefinitionsJSON   `json:"-"`
}

// userRoleDefinitionsJSON contains the JSON metadata for the struct
// [UserRoleDefinitions]
type userRoleDefinitionsJSON struct {
	RequestStatus apijson.Field
	Roles         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserRoleDefinitions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleDefinitionsJSON) RawJSON() string {
	return r.raw
}

type UserRoleDefinitionsRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserRoleDefinitionsRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                     `json:"statusDescription"`
	JSON              userRoleDefinitionsRequestStatusJSON       `json:"-"`
}

// userRoleDefinitionsRequestStatusJSON contains the JSON metadata for the struct
// [UserRoleDefinitionsRequestStatus]
type userRoleDefinitionsRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserRoleDefinitionsRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleDefinitionsRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserRoleDefinitionsRequestStatusStatusCode string

const (
	UserRoleDefinitionsRequestStatusStatusCodeUnknown                    UserRoleDefinitionsRequestStatusStatusCode = "UNKNOWN"
	UserRoleDefinitionsRequestStatusStatusCodeSuccess                    UserRoleDefinitionsRequestStatusStatusCode = "SUCCESS"
	UserRoleDefinitionsRequestStatusStatusCodeUnauthorized               UserRoleDefinitionsRequestStatusStatusCode = "UNAUTHORIZED"
	UserRoleDefinitionsRequestStatusStatusCodePaymentRequired            UserRoleDefinitionsRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserRoleDefinitionsRequestStatusStatusCodeForbidden                  UserRoleDefinitionsRequestStatusStatusCode = "FORBIDDEN"
	UserRoleDefinitionsRequestStatusStatusCodeTimeout                    UserRoleDefinitionsRequestStatusStatusCode = "TIMEOUT"
	UserRoleDefinitionsRequestStatusStatusCodeExists                     UserRoleDefinitionsRequestStatusStatusCode = "EXISTS"
	UserRoleDefinitionsRequestStatusStatusCodeNotFound                   UserRoleDefinitionsRequestStatusStatusCode = "NOT_FOUND"
	UserRoleDefinitionsRequestStatusStatusCodeInternalError              UserRoleDefinitionsRequestStatusStatusCode = "INTERNAL_ERROR"
	UserRoleDefinitionsRequestStatusStatusCodeInvalidRequest             UserRoleDefinitionsRequestStatusStatusCode = "INVALID_REQUEST"
	UserRoleDefinitionsRequestStatusStatusCodeInvalidRequestVersion      UserRoleDefinitionsRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserRoleDefinitionsRequestStatusStatusCodeInvalidRequestData         UserRoleDefinitionsRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserRoleDefinitionsRequestStatusStatusCodeMethodNotAllowed           UserRoleDefinitionsRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserRoleDefinitionsRequestStatusStatusCodeConflict                   UserRoleDefinitionsRequestStatusStatusCode = "CONFLICT"
	UserRoleDefinitionsRequestStatusStatusCodeUnprocessableEntity        UserRoleDefinitionsRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserRoleDefinitionsRequestStatusStatusCodeTooManyRequests            UserRoleDefinitionsRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserRoleDefinitionsRequestStatusStatusCodeInsufficientStorage        UserRoleDefinitionsRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserRoleDefinitionsRequestStatusStatusCodeServiceUnavailable         UserRoleDefinitionsRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserRoleDefinitionsRequestStatusStatusCodePayloadTooLarge            UserRoleDefinitionsRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserRoleDefinitionsRequestStatusStatusCodeNotAcceptable              UserRoleDefinitionsRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserRoleDefinitionsRequestStatusStatusCodeUnavailableForLegalReasons UserRoleDefinitionsRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserRoleDefinitionsRequestStatusStatusCodeBadGateway                 UserRoleDefinitionsRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserRoleDefinitionsRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserRoleDefinitionsRequestStatusStatusCodeUnknown, UserRoleDefinitionsRequestStatusStatusCodeSuccess, UserRoleDefinitionsRequestStatusStatusCodeUnauthorized, UserRoleDefinitionsRequestStatusStatusCodePaymentRequired, UserRoleDefinitionsRequestStatusStatusCodeForbidden, UserRoleDefinitionsRequestStatusStatusCodeTimeout, UserRoleDefinitionsRequestStatusStatusCodeExists, UserRoleDefinitionsRequestStatusStatusCodeNotFound, UserRoleDefinitionsRequestStatusStatusCodeInternalError, UserRoleDefinitionsRequestStatusStatusCodeInvalidRequest, UserRoleDefinitionsRequestStatusStatusCodeInvalidRequestVersion, UserRoleDefinitionsRequestStatusStatusCodeInvalidRequestData, UserRoleDefinitionsRequestStatusStatusCodeMethodNotAllowed, UserRoleDefinitionsRequestStatusStatusCodeConflict, UserRoleDefinitionsRequestStatusStatusCodeUnprocessableEntity, UserRoleDefinitionsRequestStatusStatusCodeTooManyRequests, UserRoleDefinitionsRequestStatusStatusCodeInsufficientStorage, UserRoleDefinitionsRequestStatusStatusCodeServiceUnavailable, UserRoleDefinitionsRequestStatusStatusCodePayloadTooLarge, UserRoleDefinitionsRequestStatusStatusCodeNotAcceptable, UserRoleDefinitionsRequestStatusStatusCodeUnavailableForLegalReasons, UserRoleDefinitionsRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// List of roles
type UserRoleDefinitionsRole struct {
	// List of actions that this role allows
	AllowedActions []UserRoleDefinitionsRolesAllowedAction `json:"allowedActions"`
	// Display Name of the role
	DisplayName string `json:"displayName"`
	// Name of the role
	Name string `json:"name"`
	// Product information of the role
	Product UserRoleDefinitionsRolesProduct `json:"product"`
	// Short Display Name of the role
	ShortDisplayName string                      `json:"shortDisplayName"`
	JSON             userRoleDefinitionsRoleJSON `json:"-"`
}

// userRoleDefinitionsRoleJSON contains the JSON metadata for the struct
// [UserRoleDefinitionsRole]
type userRoleDefinitionsRoleJSON struct {
	AllowedActions   apijson.Field
	DisplayName      apijson.Field
	Name             apijson.Field
	Product          apijson.Field
	ShortDisplayName apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserRoleDefinitionsRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleDefinitionsRoleJSON) RawJSON() string {
	return r.raw
}

// List of actions that this role allows
type UserRoleDefinitionsRolesAllowedAction struct {
	// List of access levels that this role allows
	AccessLevels []string `json:"accessLevels"`
	// Service that this role allows
	Service string                                    `json:"service"`
	JSON    userRoleDefinitionsRolesAllowedActionJSON `json:"-"`
}

// userRoleDefinitionsRolesAllowedActionJSON contains the JSON metadata for the
// struct [UserRoleDefinitionsRolesAllowedAction]
type userRoleDefinitionsRolesAllowedActionJSON struct {
	AccessLevels apijson.Field
	Service      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserRoleDefinitionsRolesAllowedAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleDefinitionsRolesAllowedActionJSON) RawJSON() string {
	return r.raw
}

// Product information of the role
type UserRoleDefinitionsRolesProduct struct {
	// Display Name of the product from the product catalog to which this role is
	// associated with
	DisplayName string `json:"displayName"`
	// Name of the product from the product catalog to which this role is associated
	// with
	Name string                              `json:"name"`
	JSON userRoleDefinitionsRolesProductJSON `json:"-"`
}

// userRoleDefinitionsRolesProductJSON contains the JSON metadata for the struct
// [UserRoleDefinitionsRolesProduct]
type userRoleDefinitionsRolesProductJSON struct {
	DisplayName apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserRoleDefinitionsRolesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleDefinitionsRolesProductJSON) RawJSON() string {
	return r.raw
}

type RoleListParams struct {
	// flag indicate if hidden roles should be included
	ShowHidden param.Field[bool] `query:"show-hidden"`
}

// URLQuery serializes [RoleListParams]'s query parameters as `url.Values`.
func (r RoleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
