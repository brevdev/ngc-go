// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// OrgCreditService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgCreditService] method instead.
type OrgCreditService struct {
	Options []option.RequestOption
}

// NewOrgCreditService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgCreditService(opts ...option.RequestOption) (r *OrgCreditService) {
	r = &OrgCreditService{}
	r.Options = opts
	return
}

// Get Organization credits
func (r *OrgCreditService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *CreditsHistory, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/orgs/%s/credits", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// showing credit balance
type CreditsHistory struct {
	CreditsHistory CreditsHistoryCreditsHistory `json:"creditsHistory"`
	RequestStatus  CreditsHistoryRequestStatus  `json:"requestStatus"`
	JSON           creditsHistoryJSON           `json:"-"`
}

// creditsHistoryJSON contains the JSON metadata for the struct [CreditsHistory]
type creditsHistoryJSON struct {
	CreditsHistory apijson.Field
	RequestStatus  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditsHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditsHistoryJSON) RawJSON() string {
	return r.raw
}

type CreditsHistoryCreditsHistory struct {
	// Latest credit balance information
	CreditBalance int64                            `json:"creditBalance"`
	JSON          creditsHistoryCreditsHistoryJSON `json:"-"`
}

// creditsHistoryCreditsHistoryJSON contains the JSON metadata for the struct
// [CreditsHistoryCreditsHistory]
type creditsHistoryCreditsHistoryJSON struct {
	CreditBalance apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CreditsHistoryCreditsHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditsHistoryCreditsHistoryJSON) RawJSON() string {
	return r.raw
}

type CreditsHistoryRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        CreditsHistoryRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                `json:"statusDescription"`
	JSON              creditsHistoryRequestStatusJSON       `json:"-"`
}

// creditsHistoryRequestStatusJSON contains the JSON metadata for the struct
// [CreditsHistoryRequestStatus]
type creditsHistoryRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CreditsHistoryRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditsHistoryRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type CreditsHistoryRequestStatusStatusCode string

const (
	CreditsHistoryRequestStatusStatusCodeUnknown                    CreditsHistoryRequestStatusStatusCode = "UNKNOWN"
	CreditsHistoryRequestStatusStatusCodeSuccess                    CreditsHistoryRequestStatusStatusCode = "SUCCESS"
	CreditsHistoryRequestStatusStatusCodeUnauthorized               CreditsHistoryRequestStatusStatusCode = "UNAUTHORIZED"
	CreditsHistoryRequestStatusStatusCodePaymentRequired            CreditsHistoryRequestStatusStatusCode = "PAYMENT_REQUIRED"
	CreditsHistoryRequestStatusStatusCodeForbidden                  CreditsHistoryRequestStatusStatusCode = "FORBIDDEN"
	CreditsHistoryRequestStatusStatusCodeTimeout                    CreditsHistoryRequestStatusStatusCode = "TIMEOUT"
	CreditsHistoryRequestStatusStatusCodeExists                     CreditsHistoryRequestStatusStatusCode = "EXISTS"
	CreditsHistoryRequestStatusStatusCodeNotFound                   CreditsHistoryRequestStatusStatusCode = "NOT_FOUND"
	CreditsHistoryRequestStatusStatusCodeInternalError              CreditsHistoryRequestStatusStatusCode = "INTERNAL_ERROR"
	CreditsHistoryRequestStatusStatusCodeInvalidRequest             CreditsHistoryRequestStatusStatusCode = "INVALID_REQUEST"
	CreditsHistoryRequestStatusStatusCodeInvalidRequestVersion      CreditsHistoryRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	CreditsHistoryRequestStatusStatusCodeInvalidRequestData         CreditsHistoryRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	CreditsHistoryRequestStatusStatusCodeMethodNotAllowed           CreditsHistoryRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	CreditsHistoryRequestStatusStatusCodeConflict                   CreditsHistoryRequestStatusStatusCode = "CONFLICT"
	CreditsHistoryRequestStatusStatusCodeUnprocessableEntity        CreditsHistoryRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	CreditsHistoryRequestStatusStatusCodeTooManyRequests            CreditsHistoryRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	CreditsHistoryRequestStatusStatusCodeInsufficientStorage        CreditsHistoryRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	CreditsHistoryRequestStatusStatusCodeServiceUnavailable         CreditsHistoryRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	CreditsHistoryRequestStatusStatusCodePayloadTooLarge            CreditsHistoryRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	CreditsHistoryRequestStatusStatusCodeNotAcceptable              CreditsHistoryRequestStatusStatusCode = "NOT_ACCEPTABLE"
	CreditsHistoryRequestStatusStatusCodeUnavailableForLegalReasons CreditsHistoryRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	CreditsHistoryRequestStatusStatusCodeBadGateway                 CreditsHistoryRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r CreditsHistoryRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case CreditsHistoryRequestStatusStatusCodeUnknown, CreditsHistoryRequestStatusStatusCodeSuccess, CreditsHistoryRequestStatusStatusCodeUnauthorized, CreditsHistoryRequestStatusStatusCodePaymentRequired, CreditsHistoryRequestStatusStatusCodeForbidden, CreditsHistoryRequestStatusStatusCodeTimeout, CreditsHistoryRequestStatusStatusCodeExists, CreditsHistoryRequestStatusStatusCodeNotFound, CreditsHistoryRequestStatusStatusCodeInternalError, CreditsHistoryRequestStatusStatusCodeInvalidRequest, CreditsHistoryRequestStatusStatusCodeInvalidRequestVersion, CreditsHistoryRequestStatusStatusCodeInvalidRequestData, CreditsHistoryRequestStatusStatusCodeMethodNotAllowed, CreditsHistoryRequestStatusStatusCodeConflict, CreditsHistoryRequestStatusStatusCodeUnprocessableEntity, CreditsHistoryRequestStatusStatusCodeTooManyRequests, CreditsHistoryRequestStatusStatusCodeInsufficientStorage, CreditsHistoryRequestStatusStatusCodeServiceUnavailable, CreditsHistoryRequestStatusStatusCodePayloadTooLarge, CreditsHistoryRequestStatusStatusCodeNotAcceptable, CreditsHistoryRequestStatusStatusCodeUnavailableForLegalReasons, CreditsHistoryRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
