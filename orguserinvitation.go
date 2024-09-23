// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/option"
)

// OrgUserInvitationService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserInvitationService] method instead.
type OrgUserInvitationService struct {
	Options []option.RequestOption
}

// NewOrgUserInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgUserInvitationService(opts ...option.RequestOption) (r *OrgUserInvitationService) {
	r = &OrgUserInvitationService{}
	r.Options = opts
	return
}

// Response for a list of user invitations.
type UserInvitationList struct {
	// List of invitations.
	Invitations []UserInvitationListInvitation `json:"invitations"`
	// object that describes the pagination information
	PaginationInfo UserInvitationListPaginationInfo `json:"paginationInfo"`
	RequestStatus  UserInvitationListRequestStatus  `json:"requestStatus"`
	JSON           userInvitationListJSON           `json:"-"`
}

// userInvitationListJSON contains the JSON metadata for the struct
// [UserInvitationList]
type userInvitationListJSON struct {
	Invitations    apijson.Field
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserInvitationList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInvitationListJSON) RawJSON() string {
	return r.raw
}

// User invitation to an NGC org or team
type UserInvitationListInvitation struct {
	// Unique invitation ID
	ID string `json:"id"`
	// Date on which the invitation was created. (ISO-8601 format)
	CreatedDate string `json:"createdDate"`
	// Email address of the user.
	Email string `json:"email"`
	// Flag indicating if the invitation has already been accepted by the user.
	IsProcessed bool `json:"isProcessed"`
	// user name
	Name string `json:"name"`
	// Org to which a user was invited.
	Org string `json:"org"`
	// List of roles that the user have.
	Roles []string `json:"roles"`
	// Team to which a user was invited.
	Team string `json:"team"`
	// Type of invitation. The invitation is either to an organization or to a team
	// within organization.
	Type UserInvitationListInvitationsType `json:"type"`
	JSON userInvitationListInvitationJSON  `json:"-"`
}

// userInvitationListInvitationJSON contains the JSON metadata for the struct
// [UserInvitationListInvitation]
type userInvitationListInvitationJSON struct {
	ID          apijson.Field
	CreatedDate apijson.Field
	Email       apijson.Field
	IsProcessed apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	Roles       apijson.Field
	Team        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInvitationListInvitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInvitationListInvitationJSON) RawJSON() string {
	return r.raw
}

// Type of invitation. The invitation is either to an organization or to a team
// within organization.
type UserInvitationListInvitationsType string

const (
	UserInvitationListInvitationsTypeOrganization UserInvitationListInvitationsType = "ORGANIZATION"
	UserInvitationListInvitationsTypeTeam         UserInvitationListInvitationsType = "TEAM"
)

func (r UserInvitationListInvitationsType) IsKnown() bool {
	switch r {
	case UserInvitationListInvitationsTypeOrganization, UserInvitationListInvitationsTypeTeam:
		return true
	}
	return false
}

// object that describes the pagination information
type UserInvitationListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                                `json:"totalResults"`
	JSON         userInvitationListPaginationInfoJSON `json:"-"`
}

// userInvitationListPaginationInfoJSON contains the JSON metadata for the struct
// [UserInvitationListPaginationInfo]
type userInvitationListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserInvitationListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInvitationListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type UserInvitationListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserInvitationListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                    `json:"statusDescription"`
	JSON              userInvitationListRequestStatusJSON       `json:"-"`
}

// userInvitationListRequestStatusJSON contains the JSON metadata for the struct
// [UserInvitationListRequestStatus]
type userInvitationListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserInvitationListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInvitationListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserInvitationListRequestStatusStatusCode string

const (
	UserInvitationListRequestStatusStatusCodeUnknown                    UserInvitationListRequestStatusStatusCode = "UNKNOWN"
	UserInvitationListRequestStatusStatusCodeSuccess                    UserInvitationListRequestStatusStatusCode = "SUCCESS"
	UserInvitationListRequestStatusStatusCodeUnauthorized               UserInvitationListRequestStatusStatusCode = "UNAUTHORIZED"
	UserInvitationListRequestStatusStatusCodePaymentRequired            UserInvitationListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserInvitationListRequestStatusStatusCodeForbidden                  UserInvitationListRequestStatusStatusCode = "FORBIDDEN"
	UserInvitationListRequestStatusStatusCodeTimeout                    UserInvitationListRequestStatusStatusCode = "TIMEOUT"
	UserInvitationListRequestStatusStatusCodeExists                     UserInvitationListRequestStatusStatusCode = "EXISTS"
	UserInvitationListRequestStatusStatusCodeNotFound                   UserInvitationListRequestStatusStatusCode = "NOT_FOUND"
	UserInvitationListRequestStatusStatusCodeInternalError              UserInvitationListRequestStatusStatusCode = "INTERNAL_ERROR"
	UserInvitationListRequestStatusStatusCodeInvalidRequest             UserInvitationListRequestStatusStatusCode = "INVALID_REQUEST"
	UserInvitationListRequestStatusStatusCodeInvalidRequestVersion      UserInvitationListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserInvitationListRequestStatusStatusCodeInvalidRequestData         UserInvitationListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserInvitationListRequestStatusStatusCodeMethodNotAllowed           UserInvitationListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserInvitationListRequestStatusStatusCodeConflict                   UserInvitationListRequestStatusStatusCode = "CONFLICT"
	UserInvitationListRequestStatusStatusCodeUnprocessableEntity        UserInvitationListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserInvitationListRequestStatusStatusCodeTooManyRequests            UserInvitationListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserInvitationListRequestStatusStatusCodeInsufficientStorage        UserInvitationListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserInvitationListRequestStatusStatusCodeServiceUnavailable         UserInvitationListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserInvitationListRequestStatusStatusCodePayloadTooLarge            UserInvitationListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserInvitationListRequestStatusStatusCodeNotAcceptable              UserInvitationListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserInvitationListRequestStatusStatusCodeUnavailableForLegalReasons UserInvitationListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserInvitationListRequestStatusStatusCodeBadGateway                 UserInvitationListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserInvitationListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserInvitationListRequestStatusStatusCodeUnknown, UserInvitationListRequestStatusStatusCodeSuccess, UserInvitationListRequestStatusStatusCodeUnauthorized, UserInvitationListRequestStatusStatusCodePaymentRequired, UserInvitationListRequestStatusStatusCodeForbidden, UserInvitationListRequestStatusStatusCodeTimeout, UserInvitationListRequestStatusStatusCodeExists, UserInvitationListRequestStatusStatusCodeNotFound, UserInvitationListRequestStatusStatusCodeInternalError, UserInvitationListRequestStatusStatusCodeInvalidRequest, UserInvitationListRequestStatusStatusCodeInvalidRequestVersion, UserInvitationListRequestStatusStatusCodeInvalidRequestData, UserInvitationListRequestStatusStatusCodeMethodNotAllowed, UserInvitationListRequestStatusStatusCodeConflict, UserInvitationListRequestStatusStatusCodeUnprocessableEntity, UserInvitationListRequestStatusStatusCodeTooManyRequests, UserInvitationListRequestStatusStatusCodeInsufficientStorage, UserInvitationListRequestStatusStatusCodeServiceUnavailable, UserInvitationListRequestStatusStatusCodePayloadTooLarge, UserInvitationListRequestStatusStatusCodeNotAcceptable, UserInvitationListRequestStatusStatusCodeUnavailableForLegalReasons, UserInvitationListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
