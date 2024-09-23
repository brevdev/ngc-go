// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/ngc-go/internal/apijson"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
	"github.com/brevdev/ngc-go/shared"
)

// OrgV3TeamUserNcaInvitationService contains methods and other services that help
// with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV3TeamUserNcaInvitationService] method instead.
type OrgV3TeamUserNcaInvitationService struct {
	Options []option.RequestOption
}

// NewOrgV3TeamUserNcaInvitationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrgV3TeamUserNcaInvitationService(opts ...option.RequestOption) (r *OrgV3TeamUserNcaInvitationService) {
	r = &OrgV3TeamUserNcaInvitationService{}
	r.Options = opts
	return
}

// Invites and creates a User in team
func (r *OrgV3TeamUserNcaInvitationService) New(ctx context.Context, orgName string, teamName string, body OrgV3TeamUserNcaInvitationNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/nca-invitations", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrgV3TeamUserNcaInvitationNewParams struct {
	// Is the user email
	Email param.Field[string] `json:"email"`
	// Is the numbers of days the invitation will expire
	InvitationExpirationIn param.Field[int64] `json:"invitationExpirationIn"`
	// Nca allow users to be invited as Admin and as Member
	InviteAs param.Field[OrgV3TeamUserNcaInvitationNewParamsInviteAs] `json:"inviteAs"`
	// Is a message to the new user
	Message param.Field[string] `json:"message"`
}

func (r OrgV3TeamUserNcaInvitationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nca allow users to be invited as Admin and as Member
type OrgV3TeamUserNcaInvitationNewParamsInviteAs string

const (
	OrgV3TeamUserNcaInvitationNewParamsInviteAsAdmin  OrgV3TeamUserNcaInvitationNewParamsInviteAs = "ADMIN"
	OrgV3TeamUserNcaInvitationNewParamsInviteAsMember OrgV3TeamUserNcaInvitationNewParamsInviteAs = "MEMBER"
)

func (r OrgV3TeamUserNcaInvitationNewParamsInviteAs) IsKnown() bool {
	switch r {
	case OrgV3TeamUserNcaInvitationNewParamsInviteAsAdmin, OrgV3TeamUserNcaInvitationNewParamsInviteAsMember:
		return true
	}
	return false
}
