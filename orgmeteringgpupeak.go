// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apiquery"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/param"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

// OrgMeteringGpupeakService contains methods and other services that help with
// interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgMeteringGpupeakService] method instead.
type OrgMeteringGpupeakService struct {
	Options []option.RequestOption
}

// NewOrgMeteringGpupeakService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgMeteringGpupeakService(opts ...option.RequestOption) (r *OrgMeteringGpupeakService) {
	r = &OrgMeteringGpupeakService{}
	r.Options = opts
	return
}

// Returns GPU Peak Usage as measurement series. Requires admin privileges for
// organization.
func (r *OrgMeteringGpupeakService) List(ctx context.Context, orgName string, query OrgMeteringGpupeakListParams, opts ...option.RequestOption) (res *shared.MeteringResultList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/meterings/gpupeak", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrgMeteringGpupeakListParams struct {
	TheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS param.Field[OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS] `query:"The to date in ISO 8601 format including time zone information (yyyy-MM-dd'T'HH:mm:ss"`
}

// URLQuery serializes [OrgMeteringGpupeakListParams]'s query parameters as
// `url.Values`.
func (r OrgMeteringGpupeakListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS struct {
	Sssz param.Field[time.Time] `query:"SSSZ)" format:"date-time"`
}

// URLQuery serializes
// [OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS]'s
// query parameters as `url.Values`.
func (r OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
