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

// V3OrgService contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3OrgService] method instead.
type V3OrgService struct {
	Options []option.RequestOption
}

// NewV3OrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3OrgService(opts ...option.RequestOption) (r *V3OrgService) {
	r = &V3OrgService{}
	r.Options = opts
	return
}

// Validate org creation from proto org
func (r *V3OrgService) Validate(ctx context.Context, query V3OrgValidateParams, opts ...option.RequestOption) (res *OrgInvitation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/orgs/proto-org/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Invitation Validation Response.
type OrgInvitation struct {
	// Org invitation to NGC
	OrgInvitation OrgInvitationOrgInvitation `json:"orgInvitation"`
	RequestStatus OrgInvitationRequestStatus `json:"requestStatus"`
	JSON          orgInvitationJSON          `json:"-"`
}

// orgInvitationJSON contains the JSON metadata for the struct [OrgInvitation]
type orgInvitationJSON struct {
	OrgInvitation apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgInvitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationJSON) RawJSON() string {
	return r.raw
}

// Org invitation to NGC
type OrgInvitationOrgInvitation struct {
	// Email address of the user.
	Email string `json:"email"`
	// Proto Org identifier.
	ProtoOrgID string                         `json:"protoOrgId"`
	JSON       orgInvitationOrgInvitationJSON `json:"-"`
}

// orgInvitationOrgInvitationJSON contains the JSON metadata for the struct
// [OrgInvitationOrgInvitation]
type orgInvitationOrgInvitationJSON struct {
	Email       apijson.Field
	ProtoOrgID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgInvitationOrgInvitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationOrgInvitationJSON) RawJSON() string {
	return r.raw
}

type OrgInvitationRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgInvitationRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                               `json:"statusDescription"`
	JSON              orgInvitationRequestStatusJSON       `json:"-"`
}

// orgInvitationRequestStatusJSON contains the JSON metadata for the struct
// [OrgInvitationRequestStatus]
type orgInvitationRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgInvitationRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgInvitationRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgInvitationRequestStatusStatusCode string

const (
	OrgInvitationRequestStatusStatusCodeUnknown                    OrgInvitationRequestStatusStatusCode = "UNKNOWN"
	OrgInvitationRequestStatusStatusCodeSuccess                    OrgInvitationRequestStatusStatusCode = "SUCCESS"
	OrgInvitationRequestStatusStatusCodeUnauthorized               OrgInvitationRequestStatusStatusCode = "UNAUTHORIZED"
	OrgInvitationRequestStatusStatusCodePaymentRequired            OrgInvitationRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgInvitationRequestStatusStatusCodeForbidden                  OrgInvitationRequestStatusStatusCode = "FORBIDDEN"
	OrgInvitationRequestStatusStatusCodeTimeout                    OrgInvitationRequestStatusStatusCode = "TIMEOUT"
	OrgInvitationRequestStatusStatusCodeExists                     OrgInvitationRequestStatusStatusCode = "EXISTS"
	OrgInvitationRequestStatusStatusCodeNotFound                   OrgInvitationRequestStatusStatusCode = "NOT_FOUND"
	OrgInvitationRequestStatusStatusCodeInternalError              OrgInvitationRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgInvitationRequestStatusStatusCodeInvalidRequest             OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST"
	OrgInvitationRequestStatusStatusCodeInvalidRequestVersion      OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgInvitationRequestStatusStatusCodeInvalidRequestData         OrgInvitationRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgInvitationRequestStatusStatusCodeMethodNotAllowed           OrgInvitationRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgInvitationRequestStatusStatusCodeConflict                   OrgInvitationRequestStatusStatusCode = "CONFLICT"
	OrgInvitationRequestStatusStatusCodeUnprocessableEntity        OrgInvitationRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgInvitationRequestStatusStatusCodeTooManyRequests            OrgInvitationRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgInvitationRequestStatusStatusCodeInsufficientStorage        OrgInvitationRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgInvitationRequestStatusStatusCodeServiceUnavailable         OrgInvitationRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgInvitationRequestStatusStatusCodePayloadTooLarge            OrgInvitationRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgInvitationRequestStatusStatusCodeNotAcceptable              OrgInvitationRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgInvitationRequestStatusStatusCodeUnavailableForLegalReasons OrgInvitationRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgInvitationRequestStatusStatusCodeBadGateway                 OrgInvitationRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgInvitationRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgInvitationRequestStatusStatusCodeUnknown, OrgInvitationRequestStatusStatusCodeSuccess, OrgInvitationRequestStatusStatusCodeUnauthorized, OrgInvitationRequestStatusStatusCodePaymentRequired, OrgInvitationRequestStatusStatusCodeForbidden, OrgInvitationRequestStatusStatusCodeTimeout, OrgInvitationRequestStatusStatusCodeExists, OrgInvitationRequestStatusStatusCodeNotFound, OrgInvitationRequestStatusStatusCodeInternalError, OrgInvitationRequestStatusStatusCodeInvalidRequest, OrgInvitationRequestStatusStatusCodeInvalidRequestVersion, OrgInvitationRequestStatusStatusCodeInvalidRequestData, OrgInvitationRequestStatusStatusCodeMethodNotAllowed, OrgInvitationRequestStatusStatusCodeConflict, OrgInvitationRequestStatusStatusCodeUnprocessableEntity, OrgInvitationRequestStatusStatusCodeTooManyRequests, OrgInvitationRequestStatusStatusCodeInsufficientStorage, OrgInvitationRequestStatusStatusCodeServiceUnavailable, OrgInvitationRequestStatusStatusCodePayloadTooLarge, OrgInvitationRequestStatusStatusCodeNotAcceptable, OrgInvitationRequestStatusStatusCodeUnavailableForLegalReasons, OrgInvitationRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type V3OrgValidateParams struct {
	// JWT that contains org owner email and proto org identifier
	InvitationToken param.Field[string] `query:"invitation_token,required"`
}

// URLQuery serializes [V3OrgValidateParams]'s query parameters as `url.Values`.
func (r V3OrgValidateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
