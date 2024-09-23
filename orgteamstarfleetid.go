// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
	"github.com/stainless-sdks/ngc-go/shared"
)

// OrgTeamStarfleetIDService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamStarfleetIDService] method instead.
type OrgTeamStarfleetIDService struct {
	Options []option.RequestOption
}

// NewOrgTeamStarfleetIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgTeamStarfleetIDService(opts ...option.RequestOption) (r *OrgTeamStarfleetIDService) {
	r = &OrgTeamStarfleetIDService{}
	r.Options = opts
	return
}

// Get User details in team by starfleet Id
func (r *OrgTeamStarfleetIDService) Get(ctx context.Context, orgName string, teamName string, starfleetID string, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if starfleetID == "" {
		err = errors.New("missing required starfleet-id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/team/%s/starfleetIds/%s", orgName, teamName, starfleetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
