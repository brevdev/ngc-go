// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

// AdminTeamUserService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminTeamUserService] method instead.
type AdminTeamUserService struct {
	Options []option.RequestOption
}

// NewAdminTeamUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminTeamUserService(opts ...option.RequestOption) (r *AdminTeamUserService) {
	r = &AdminTeamUserService{}
	r.Options = opts
	return
}

// Remove user role in team (if the last role is removed, the user is removed from
// the team).
func (r *AdminTeamUserService) RemoveRole(ctx context.Context, orgName string, teamName string, id string, body AdminTeamUserRemoveRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
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
	path := fmt.Sprintf("v2/admin/org/%s/team/%s/users/%s/remove-role", orgName, teamName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AdminTeamUserRemoveRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [AdminTeamUserRemoveRoleParams]'s query parameters as
// `url.Values`.
func (r AdminTeamUserRemoveRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
