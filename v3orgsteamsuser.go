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

// V3OrgsTeamsUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3OrgsTeamsUserService] method instead.
type V3OrgsTeamsUserService struct {
	Options []option.RequestOption
}

// NewV3OrgsTeamsUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3OrgsTeamsUserService(opts ...option.RequestOption) (r *V3OrgsTeamsUserService) {
	r = &V3OrgsTeamsUserService{}
	r.Options = opts
	return
}

// Get info and role/invitation in a team by email or id
func (r *V3OrgsTeamsUserService) Get(ctx context.Context, orgName string, teamName string, userEmailOrID string, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/%s", orgName, teamName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
