// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/ngc-go/internal/apiquery"
	"github.com/stainless-sdks/ngc-go/internal/param"
	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// AdminOrgTeamService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgTeamService] method instead.
type AdminOrgTeamService struct {
	Options []option.RequestOption
}

// NewAdminOrgTeamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgTeamService(opts ...option.RequestOption) (r *AdminOrgTeamService) {
	r = &AdminOrgTeamService{}
	r.Options = opts
	return
}

// List all Teams
func (r *AdminOrgTeamService) List(ctx context.Context, orgName string, query AdminOrgTeamListParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/teams", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgTeamListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// Get all team that has scan allow override only
	ScanAllowOverride param.Field[bool] `query:"scan-allow-override"`
	// Get all team that has scan by default only
	ScanByDefault param.Field[bool] `query:"scan-by-default"`
	// Get all team that has scan show results only
	ScanShowResults param.Field[bool] `query:"scan-show-results"`
	// Team name to search
	TeamName param.Field[string] `query:"team-name"`
}

// URLQuery serializes [AdminOrgTeamListParams]'s query parameters as `url.Values`.
func (r AdminOrgTeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
