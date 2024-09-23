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
	"github.com/stainless-sdks/ngc-go/shared"
)

// V2AdminOrgUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AdminOrgUserService] method instead.
type V2AdminOrgUserService struct {
	Options []option.RequestOption
}

// NewV2AdminOrgUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AdminOrgUserService(opts ...option.RequestOption) (r *V2AdminOrgUserService) {
	r = &V2AdminOrgUserService{}
	r.Options = opts
	return
}

// Add user role in org.
func (r *V2AdminOrgUserService) AddRole(ctx context.Context, orgName string, id string, body V2AdminOrgUserAddRoleParams, opts ...option.RequestOption) (res *shared.UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s/add-role", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type V2AdminOrgUserAddRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [V2AdminOrgUserAddRoleParams]'s query parameters as
// `url.Values`.
func (r V2AdminOrgUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
