// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

// OrgTeamService contains methods and other services that help with interacting
// with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamService] method instead.
type OrgTeamService struct {
	Options []option.RequestOption
	Users   *OrgTeamUserService
}

// NewOrgTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgTeamService(opts ...option.RequestOption) (r *OrgTeamService) {
	r = &OrgTeamService{}
	r.Options = opts
	r.Users = NewOrgTeamUserService(opts...)
	return
}

// Get Team by name
func (r *OrgTeamService) Get(ctx context.Context, orgName string, teamName string, opts ...option.RequestOption) (res *Team, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams/%s", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit a Team
func (r *OrgTeamService) Update(ctx context.Context, orgName string, teamName string, body OrgTeamUpdateParams, opts ...option.RequestOption) (res *Team, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams/%s", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Teams
func (r *OrgTeamService) List(ctx context.Context, orgName string, query OrgTeamListParams, opts ...option.RequestOption) (res *TeamList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Team
func (r *OrgTeamService) Delete(ctx context.Context, orgName string, teamName string, opts ...option.RequestOption) (res *OrgTeamDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams/%s", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OrgTeamDeleteResponse struct {
	RequestStatus OrgTeamDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgTeamDeleteResponseJSON          `json:"-"`
}

// orgTeamDeleteResponseJSON contains the JSON metadata for the struct
// [OrgTeamDeleteResponse]
type orgTeamDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgTeamDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgTeamDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgTeamDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                       `json:"statusDescription"`
	JSON              orgTeamDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgTeamDeleteResponseRequestStatusJSON contains the JSON metadata for the struct
// [OrgTeamDeleteResponseRequestStatus]
type orgTeamDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgTeamDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgTeamDeleteResponseRequestStatusStatusCode string

const (
	OrgTeamDeleteResponseRequestStatusStatusCodeUnknown                    OrgTeamDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgTeamDeleteResponseRequestStatusStatusCodeSuccess                    OrgTeamDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgTeamDeleteResponseRequestStatusStatusCodeUnauthorized               OrgTeamDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgTeamDeleteResponseRequestStatusStatusCodePaymentRequired            OrgTeamDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgTeamDeleteResponseRequestStatusStatusCodeForbidden                  OrgTeamDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgTeamDeleteResponseRequestStatusStatusCodeTimeout                    OrgTeamDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgTeamDeleteResponseRequestStatusStatusCodeExists                     OrgTeamDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgTeamDeleteResponseRequestStatusStatusCodeNotFound                   OrgTeamDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgTeamDeleteResponseRequestStatusStatusCodeInternalError              OrgTeamDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgTeamDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgTeamDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgTeamDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgTeamDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgTeamDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgTeamDeleteResponseRequestStatusStatusCodeConflict                   OrgTeamDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgTeamDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgTeamDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgTeamDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgTeamDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgTeamDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgTeamDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgTeamDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgTeamDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgTeamDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgTeamDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgTeamDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgTeamDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgTeamDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgTeamDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgTeamDeleteResponseRequestStatusStatusCodeBadGateway                 OrgTeamDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgTeamDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgTeamDeleteResponseRequestStatusStatusCodeUnknown, OrgTeamDeleteResponseRequestStatusStatusCodeSuccess, OrgTeamDeleteResponseRequestStatusStatusCodeUnauthorized, OrgTeamDeleteResponseRequestStatusStatusCodePaymentRequired, OrgTeamDeleteResponseRequestStatusStatusCodeForbidden, OrgTeamDeleteResponseRequestStatusStatusCodeTimeout, OrgTeamDeleteResponseRequestStatusStatusCodeExists, OrgTeamDeleteResponseRequestStatusStatusCodeNotFound, OrgTeamDeleteResponseRequestStatusStatusCodeInternalError, OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgTeamDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgTeamDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgTeamDeleteResponseRequestStatusStatusCodeConflict, OrgTeamDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgTeamDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgTeamDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgTeamDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgTeamDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgTeamDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgTeamDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgTeamDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgTeamUpdateParams struct {
	// description of the team
	Description param.Field[string] `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings param.Field[OrgTeamUpdateParamsInfinityManagerSettings] `json:"infinityManagerSettings"`
	// Repo scan setting definition
	RepoScanSettings param.Field[OrgTeamUpdateParamsRepoScanSettings] `json:"repoScanSettings"`
}

func (r OrgTeamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Infinity manager setting definition
type OrgTeamUpdateParamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled param.Field[bool] `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride param.Field[bool] `json:"infinityManagerEnableTeamOverride"`
}

func (r OrgTeamUpdateParamsInfinityManagerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Repo scan setting definition
type OrgTeamUpdateParamsRepoScanSettings struct {
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

func (r OrgTeamUpdateParamsRepoScanSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgTeamListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
}

// URLQuery serializes [OrgTeamListParams]'s query parameters as `url.Values`.
func (r OrgTeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
