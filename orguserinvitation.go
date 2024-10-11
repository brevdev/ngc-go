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

// OrgUserInvitationService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserInvitationService] method instead.
type OrgUserInvitationService struct {
	Options []option.RequestOption
}

// NewOrgUserInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgUserInvitationService(opts ...option.RequestOption) (r *OrgUserInvitationService) {
	r = &OrgUserInvitationService{}
	r.Options = opts
	return
}

// List invitations in an org. (Org User Admin privileges required)
func (r *OrgUserInvitationService) List(ctx context.Context, orgName string, query OrgUserInvitationListParams, opts ...option.RequestOption) (res *shared.UserInvitationList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/invitations", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Resend email of a specific invitation in an org (Org User Admin privileges
// required).
func (r *OrgUserInvitationService) ResendInvitationEmail(ctx context.Context, orgName string, id string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/invitations/%s/resend-invitation-email", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrgUserInvitationListParams struct {
	OrderBy param.Field[OrgUserInvitationListParamsOrderBy] `query:"orderBy"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters
	Q param.Field[OrgUserInvitationListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgUserInvitationListParams]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsOrderBy string

const (
	OrgUserInvitationListParamsOrderByNameAsc  OrgUserInvitationListParamsOrderBy = "NAME_ASC"
	OrgUserInvitationListParamsOrderByNameDesc OrgUserInvitationListParamsOrderBy = "NAME_DESC"
)

func (r OrgUserInvitationListParamsOrderBy) IsKnown() bool {
	switch r {
	case OrgUserInvitationListParamsOrderByNameAsc, OrgUserInvitationListParamsOrderByNameDesc:
		return true
	}
	return false
}

// User Search Parameters
type OrgUserInvitationListParamsQ struct {
	Fields      param.Field[[]string]                              `query:"fields"`
	Filters     param.Field[[]OrgUserInvitationListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                                `query:"groupBy"`
	OrderBy     param.Field[[]OrgUserInvitationListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                                 `query:"page"`
	PageSize    param.Field[int64]                                 `query:"pageSize"`
	Query       param.Field[string]                                `query:"query"`
	QueryFields param.Field[[]string]                              `query:"queryFields"`
	ScoredSize  param.Field[int64]                                 `query:"scoredSize"`
}

// URLQuery serializes [OrgUserInvitationListParamsQ]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgUserInvitationListParamsQFilter]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQOrderBy struct {
	Field param.Field[string]                                   `query:"field"`
	Value param.Field[OrgUserInvitationListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgUserInvitationListParamsQOrderBy]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQOrderByValue string

const (
	OrgUserInvitationListParamsQOrderByValueAsc  OrgUserInvitationListParamsQOrderByValue = "ASC"
	OrgUserInvitationListParamsQOrderByValueDesc OrgUserInvitationListParamsQOrderByValue = "DESC"
)

func (r OrgUserInvitationListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgUserInvitationListParamsQOrderByValueAsc, OrgUserInvitationListParamsQOrderByValueDesc:
		return true
	}
	return false
}
