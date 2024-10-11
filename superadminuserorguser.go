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
	"github.com/brevdev/ngc-go/shared"
)

// SuperAdminUserOrgUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminUserOrgUserService] method instead.
type SuperAdminUserOrgUserService struct {
	Options []option.RequestOption
}

// NewSuperAdminUserOrgUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperAdminUserOrgUserService(opts ...option.RequestOption) (r *SuperAdminUserOrgUserService) {
	r = &SuperAdminUserOrgUserService{}
	r.Options = opts
	return
}

// Create an user in an Organization (Super Admin privileges required)
func (r *SuperAdminUserOrgUserService) New(ctx context.Context, orgName string, params SuperAdminUserOrgUserNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add existing User to an Org
func (r *SuperAdminUserOrgUserService) Add(ctx context.Context, orgName string, id string, params SuperAdminUserOrgUserAddParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove User from org.
func (r *SuperAdminUserOrgUserService) Remove(ctx context.Context, orgName string, id string, opts ...option.RequestOption) (res *SuperAdminUserOrgUserRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SuperAdminUserOrgUserRemoveResponse struct {
	RequestStatus SuperAdminUserOrgUserRemoveResponseRequestStatus `json:"requestStatus"`
	JSON          superAdminUserOrgUserRemoveResponseJSON          `json:"-"`
}

// superAdminUserOrgUserRemoveResponseJSON contains the JSON metadata for the
// struct [SuperAdminUserOrgUserRemoveResponse]
type superAdminUserOrgUserRemoveResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SuperAdminUserOrgUserRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserOrgUserRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type SuperAdminUserOrgUserRemoveResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                     `json:"statusDescription"`
	JSON              superAdminUserOrgUserRemoveResponseRequestStatusJSON       `json:"-"`
}

// superAdminUserOrgUserRemoveResponseRequestStatusJSON contains the JSON metadata
// for the struct [SuperAdminUserOrgUserRemoveResponseRequestStatus]
type superAdminUserOrgUserRemoveResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuperAdminUserOrgUserRemoveResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superAdminUserOrgUserRemoveResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode string

const (
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnknown                    SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "UNKNOWN"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeSuccess                    SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "SUCCESS"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnauthorized               SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "UNAUTHORIZED"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodePaymentRequired            SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeForbidden                  SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "FORBIDDEN"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeTimeout                    SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "TIMEOUT"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeExists                     SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "EXISTS"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeNotFound                   SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "NOT_FOUND"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInternalError              SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequest             SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestVersion      SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestData         SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeMethodNotAllowed           SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeConflict                   SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "CONFLICT"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnprocessableEntity        SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeTooManyRequests            SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInsufficientStorage        SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeServiceUnavailable         SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodePayloadTooLarge            SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeNotAcceptable              SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnavailableForLegalReasons SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeBadGateway                 SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnknown, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeSuccess, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnauthorized, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodePaymentRequired, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeForbidden, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeTimeout, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeExists, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeNotFound, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInternalError, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequest, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestVersion, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestData, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeMethodNotAllowed, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeConflict, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnprocessableEntity, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeTooManyRequests, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeInsufficientStorage, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeServiceUnavailable, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodePayloadTooLarge, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeNotAcceptable, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeUnavailableForLegalReasons, SuperAdminUserOrgUserRemoveResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type SuperAdminUserOrgUserNewParams struct {
	// Email address of the user. This should be unique.
	Email param.Field[string] `json:"email,required"`
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization
	IdpID param.Field[string] `query:"idp-id"`
	// Boolean to send email notification, default is true
	SendEmail param.Field[bool] `query:"send-email"`
	// indicates if user has opt in to nvidia emails
	EmailOptIn param.Field[bool] `json:"emailOptIn"`
	// indicates if user has accepted EULA
	EulaAccepted param.Field[bool] `json:"eulaAccepted"`
	// user name
	Name param.Field[string] `json:"name"`
	// DEPRECATED - use roleTypes which allows multiple roles
	RoleType param.Field[string] `json:"roleType"`
	// feature roles to give to the user
	RoleTypes param.Field[[]string] `json:"roleTypes"`
	// user job role
	SalesforceContactJobRole param.Field[string] `json:"salesforceContactJobRole"`
	// Metadata information about the user.
	UserMetadata param.Field[SuperAdminUserOrgUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                                     `cookie:"ncid"`
	VisitorID    param.Field[string]                                     `cookie:"VisitorID"`
}

func (r SuperAdminUserOrgUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SuperAdminUserOrgUserNewParams]'s query parameters as
// `url.Values`.
func (r SuperAdminUserOrgUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type SuperAdminUserOrgUserNewParamsUserMetadata struct {
	// Name of the company
	Company param.Field[string] `json:"company"`
	// Company URL
	CompanyURL param.Field[string] `json:"companyUrl"`
	// Country of the user
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"firstName"`
	// Industry segment
	Industry param.Field[string] `json:"industry"`
	// List of development areas that user has interest
	Interest param.Field[[]string] `json:"interest"`
	// User's last name
	LastName param.Field[string] `json:"lastName"`
	// Role of the user in the company
	Role param.Field[string] `json:"role"`
}

func (r SuperAdminUserOrgUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperAdminUserOrgUserAddParams struct {
	UserRoleDefaultsToRegistryRead param.Field[string] `query:"user role, defaults to REGISTRY_READ"`
	Ncid                           param.Field[string] `cookie:"ncid"`
	VisitorID                      param.Field[string] `cookie:"VisitorID"`
}

// URLQuery serializes [SuperAdminUserOrgUserAddParams]'s query parameters as
// `url.Values`.
func (r SuperAdminUserOrgUserAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
