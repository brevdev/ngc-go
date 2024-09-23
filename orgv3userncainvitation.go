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

// OrgV3UserNcaInvitationService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgV3UserNcaInvitationService] method instead.
type OrgV3UserNcaInvitationService struct {
	Options []option.RequestOption
}

// NewOrgV3UserNcaInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgV3UserNcaInvitationService(opts ...option.RequestOption) (r *OrgV3UserNcaInvitationService) {
	r = &OrgV3UserNcaInvitationService{}
	r.Options = opts
	return
}

// Invites and creates a User in org
func (r *OrgV3UserNcaInvitationService) New(ctx context.Context, orgName string, body OrgV3UserNcaInvitationNewParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/nca-invitations", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrgV3UserNcaInvitationNewParams struct {
	// Is the user email
	Email param.Field[string] `json:"email"`
	// Is the numbers of days the invitation will expire
	InvitationExpirationIn param.Field[int64] `json:"invitationExpirationIn"`
	// Nca allow users to be invited as Admin and as Member
	InviteAs param.Field[OrgV3UserNcaInvitationNewParamsInviteAs] `json:"inviteAs"`
	// Is a message to the new user
	Message param.Field[string] `json:"message"`
}

func (r OrgV3UserNcaInvitationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nca allow users to be invited as Admin and as Member
type OrgV3UserNcaInvitationNewParamsInviteAs string

const (
	OrgV3UserNcaInvitationNewParamsInviteAsAdmin  OrgV3UserNcaInvitationNewParamsInviteAs = "ADMIN"
	OrgV3UserNcaInvitationNewParamsInviteAsMember OrgV3UserNcaInvitationNewParamsInviteAs = "MEMBER"
)

func (r OrgV3UserNcaInvitationNewParamsInviteAs) IsKnown() bool {
	switch r {
	case OrgV3UserNcaInvitationNewParamsInviteAsAdmin, OrgV3UserNcaInvitationNewParamsInviteAsMember:
		return true
	}
	return false
}
