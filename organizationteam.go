// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/apijson"
	"github.com/stainless-sdks/ngc-go/internal/param"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// OrganizationTeamService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationTeamService] method instead.
type OrganizationTeamService struct {
	Options []option.RequestOption
	Users   *OrganizationTeamUserService
}

// NewOrganizationTeamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationTeamService(opts ...option.RequestOption) (r *OrganizationTeamService) {
	r = &OrganizationTeamService{}
	r.Options = opts
	r.Users = NewOrganizationTeamUserService(opts...)
	return
}

// Create team in org. (Org Admin Privileges Required)
func (r *OrganizationTeamService) New(ctx context.Context, orgName string, body OrganizationTeamNewParams, opts ...option.RequestOption) (res *TeamCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// response to an team creation request, includes unique team id
type TeamCreateResponse struct {
	RequestStatus TeamCreateResponseRequestStatus `json:"requestStatus"`
	// Information about the team
	Team TeamCreateResponseTeam `json:"team"`
	JSON teamCreateResponseJSON `json:"-"`
}

// teamCreateResponseJSON contains the JSON metadata for the struct
// [TeamCreateResponse]
type teamCreateResponseJSON struct {
	RequestStatus apijson.Field
	Team          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreateResponseJSON) RawJSON() string {
	return r.raw
}

type TeamCreateResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamCreateResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                    `json:"statusDescription"`
	JSON              teamCreateResponseRequestStatusJSON       `json:"-"`
}

// teamCreateResponseRequestStatusJSON contains the JSON metadata for the struct
// [TeamCreateResponseRequestStatus]
type teamCreateResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamCreateResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreateResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamCreateResponseRequestStatusStatusCode string

const (
	TeamCreateResponseRequestStatusStatusCodeUnknown                    TeamCreateResponseRequestStatusStatusCode = "UNKNOWN"
	TeamCreateResponseRequestStatusStatusCodeSuccess                    TeamCreateResponseRequestStatusStatusCode = "SUCCESS"
	TeamCreateResponseRequestStatusStatusCodeUnauthorized               TeamCreateResponseRequestStatusStatusCode = "UNAUTHORIZED"
	TeamCreateResponseRequestStatusStatusCodePaymentRequired            TeamCreateResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamCreateResponseRequestStatusStatusCodeForbidden                  TeamCreateResponseRequestStatusStatusCode = "FORBIDDEN"
	TeamCreateResponseRequestStatusStatusCodeTimeout                    TeamCreateResponseRequestStatusStatusCode = "TIMEOUT"
	TeamCreateResponseRequestStatusStatusCodeExists                     TeamCreateResponseRequestStatusStatusCode = "EXISTS"
	TeamCreateResponseRequestStatusStatusCodeNotFound                   TeamCreateResponseRequestStatusStatusCode = "NOT_FOUND"
	TeamCreateResponseRequestStatusStatusCodeInternalError              TeamCreateResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamCreateResponseRequestStatusStatusCodeInvalidRequest             TeamCreateResponseRequestStatusStatusCode = "INVALID_REQUEST"
	TeamCreateResponseRequestStatusStatusCodeInvalidRequestVersion      TeamCreateResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamCreateResponseRequestStatusStatusCodeInvalidRequestData         TeamCreateResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamCreateResponseRequestStatusStatusCodeMethodNotAllowed           TeamCreateResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamCreateResponseRequestStatusStatusCodeConflict                   TeamCreateResponseRequestStatusStatusCode = "CONFLICT"
	TeamCreateResponseRequestStatusStatusCodeUnprocessableEntity        TeamCreateResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamCreateResponseRequestStatusStatusCodeTooManyRequests            TeamCreateResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamCreateResponseRequestStatusStatusCodeInsufficientStorage        TeamCreateResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamCreateResponseRequestStatusStatusCodeServiceUnavailable         TeamCreateResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamCreateResponseRequestStatusStatusCodePayloadTooLarge            TeamCreateResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamCreateResponseRequestStatusStatusCodeNotAcceptable              TeamCreateResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamCreateResponseRequestStatusStatusCodeUnavailableForLegalReasons TeamCreateResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamCreateResponseRequestStatusStatusCodeBadGateway                 TeamCreateResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamCreateResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamCreateResponseRequestStatusStatusCodeUnknown, TeamCreateResponseRequestStatusStatusCodeSuccess, TeamCreateResponseRequestStatusStatusCodeUnauthorized, TeamCreateResponseRequestStatusStatusCodePaymentRequired, TeamCreateResponseRequestStatusStatusCodeForbidden, TeamCreateResponseRequestStatusStatusCodeTimeout, TeamCreateResponseRequestStatusStatusCodeExists, TeamCreateResponseRequestStatusStatusCodeNotFound, TeamCreateResponseRequestStatusStatusCodeInternalError, TeamCreateResponseRequestStatusStatusCodeInvalidRequest, TeamCreateResponseRequestStatusStatusCodeInvalidRequestVersion, TeamCreateResponseRequestStatusStatusCodeInvalidRequestData, TeamCreateResponseRequestStatusStatusCodeMethodNotAllowed, TeamCreateResponseRequestStatusStatusCodeConflict, TeamCreateResponseRequestStatusStatusCodeUnprocessableEntity, TeamCreateResponseRequestStatusStatusCodeTooManyRequests, TeamCreateResponseRequestStatusStatusCodeInsufficientStorage, TeamCreateResponseRequestStatusStatusCodeServiceUnavailable, TeamCreateResponseRequestStatusStatusCodePayloadTooLarge, TeamCreateResponseRequestStatusStatusCodeNotAcceptable, TeamCreateResponseRequestStatusStatusCodeUnavailableForLegalReasons, TeamCreateResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type TeamCreateResponseTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamCreateResponseTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamCreateResponseTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             teamCreateResponseTeamJSON             `json:"-"`
}

// teamCreateResponseTeamJSON contains the JSON metadata for the struct
// [TeamCreateResponseTeam]
type teamCreateResponseTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TeamCreateResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreateResponseTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamCreateResponseTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                              `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamCreateResponseTeamInfinityManagerSettingsJSON `json:"-"`
}

// teamCreateResponseTeamInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [TeamCreateResponseTeamInfinityManagerSettings]
type teamCreateResponseTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamCreateResponseTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreateResponseTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamCreateResponseTeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                       `json:"repoScanShowResults"`
	JSON                teamCreateResponseTeamRepoScanSettingsJSON `json:"-"`
}

// teamCreateResponseTeamRepoScanSettingsJSON contains the JSON metadata for the
// struct [TeamCreateResponseTeamRepoScanSettings]
type teamCreateResponseTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamCreateResponseTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreateResponseTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrganizationTeamNewParams struct {
	// team name
	Name param.Field[string] `json:"name,required"`
	// description of the team
	Description param.Field[string] `json:"description"`
}

func (r OrganizationTeamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
