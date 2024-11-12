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
)

// V2AdminOrgTeamService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AdminOrgTeamService] method instead.
type V2AdminOrgTeamService struct {
	Options []option.RequestOption
}

// NewV2AdminOrgTeamService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AdminOrgTeamService(opts ...option.RequestOption) (r *V2AdminOrgTeamService) {
	r = &V2AdminOrgTeamService{}
	r.Options = opts
	return
}

// Get one team
func (r *V2AdminOrgTeamService) Get(ctx context.Context, orgName string, teamName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/teams/%s", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit a Team
func (r *V2AdminOrgTeamService) Update(ctx context.Context, orgName string, teamName string, body V2AdminOrgTeamUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/teams/%s", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type V2AdminOrgTeamUpdateParams struct {
	// description of the team
	Description param.Field[string] `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings param.Field[V2AdminOrgTeamUpdateParamsInfinityManagerSettings] `json:"infinityManagerSettings"`
	// Repo scan setting definition
	RepoScanSettings param.Field[V2AdminOrgTeamUpdateParamsRepoScanSettings] `json:"repoScanSettings"`
}

func (r V2AdminOrgTeamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Infinity manager setting definition
type V2AdminOrgTeamUpdateParamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled param.Field[bool] `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride param.Field[bool] `json:"infinityManagerEnableTeamOverride"`
}

func (r V2AdminOrgTeamUpdateParamsInfinityManagerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Repo scan setting definition
type V2AdminOrgTeamUpdateParamsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride param.Field[bool] `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault param.Field[bool] `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled param.Field[bool] `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications param.Field[bool] `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride param.Field[bool] `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults param.Field[bool] `json:"repoScanShowResults"`
}

func (r V2AdminOrgTeamUpdateParamsRepoScanSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
