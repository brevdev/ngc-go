// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// TeamUserInvitationService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamUserInvitationService] method instead.
type TeamUserInvitationService struct {
	Options []option.RequestOption
}

// NewTeamUserInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTeamUserInvitationService(opts ...option.RequestOption) (r *TeamUserInvitationService) {
	r = &TeamUserInvitationService{}
	r.Options = opts
	return
}

// Delete a specific invitation in an team. (Org Admin or Team User Admin
// privileges required)
func (r *TeamUserInvitationService) Delete(ctx context.Context, orgName string, teamName string, id string, opts ...option.RequestOption) (res *TeamUserInvitationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/team/%s/users/invitations/%s", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TeamUserInvitationDeleteResponse struct {
	RequestStatus TeamUserInvitationDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          teamUserInvitationDeleteResponseJSON          `json:"-"`
}

// teamUserInvitationDeleteResponseJSON contains the JSON metadata for the struct
// [TeamUserInvitationDeleteResponse]
type teamUserInvitationDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamUserInvitationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamUserInvitationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TeamUserInvitationDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamUserInvitationDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                  `json:"statusDescription"`
	JSON              teamUserInvitationDeleteResponseRequestStatusJSON       `json:"-"`
}

// teamUserInvitationDeleteResponseRequestStatusJSON contains the JSON metadata for
// the struct [TeamUserInvitationDeleteResponseRequestStatus]
type teamUserInvitationDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamUserInvitationDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamUserInvitationDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamUserInvitationDeleteResponseRequestStatusStatusCode string

const (
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnknown                    TeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeSuccess                    TeamUserInvitationDeleteResponseRequestStatusStatusCode = "SUCCESS"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized               TeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired            TeamUserInvitationDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeForbidden                  TeamUserInvitationDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeTimeout                    TeamUserInvitationDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeExists                     TeamUserInvitationDeleteResponseRequestStatusStatusCode = "EXISTS"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeNotFound                   TeamUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeInternalError              TeamUserInvitationDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest             TeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      TeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData         TeamUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed           TeamUserInvitationDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeConflict                   TeamUserInvitationDeleteResponseRequestStatusStatusCode = "CONFLICT"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity        TeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests            TeamUserInvitationDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage        TeamUserInvitationDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable         TeamUserInvitationDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge            TeamUserInvitationDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable              TeamUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons TeamUserInvitationDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway                 TeamUserInvitationDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamUserInvitationDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnknown, TeamUserInvitationDeleteResponseRequestStatusStatusCodeSuccess, TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized, TeamUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired, TeamUserInvitationDeleteResponseRequestStatusStatusCodeForbidden, TeamUserInvitationDeleteResponseRequestStatusStatusCodeTimeout, TeamUserInvitationDeleteResponseRequestStatusStatusCodeExists, TeamUserInvitationDeleteResponseRequestStatusStatusCodeNotFound, TeamUserInvitationDeleteResponseRequestStatusStatusCodeInternalError, TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest, TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, TeamUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData, TeamUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed, TeamUserInvitationDeleteResponseRequestStatusStatusCodeConflict, TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity, TeamUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests, TeamUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage, TeamUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable, TeamUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge, TeamUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable, TeamUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, TeamUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
