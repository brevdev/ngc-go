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

// OrgMeteringService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgMeteringService] method instead.
type OrgMeteringService struct {
	Options []option.RequestOption
	Gpupeak *OrgMeteringGpupeakService
}

// NewOrgMeteringService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgMeteringService(opts ...option.RequestOption) (r *OrgMeteringService) {
	r = &OrgMeteringService{}
	r.Options = opts
	r.Gpupeak = NewOrgMeteringGpupeakService(opts...)
	return
}

// Returns Private Registry / EGX resources usage metering as measurement series.
// Requires admin privileges for organization.
func (r *OrgMeteringService) List(ctx context.Context, orgName string, query OrgMeteringListParams, opts ...option.RequestOption) (res *shared.MeteringResult, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/metering", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrgMeteringListParams struct {
	// request params for getting metering usage
	Q param.Field[OrgMeteringListParamsQ] `query:"q,required"`
}

// URLQuery serializes [OrgMeteringListParams]'s query parameters as `url.Values`.
func (r OrgMeteringListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// request params for getting metering usage
type OrgMeteringListParamsQ struct {
	Measurements param.Field[[]OrgMeteringListParamsQMeasurement] `query:"measurements"`
}

// URLQuery serializes [OrgMeteringListParamsQ]'s query parameters as `url.Values`.
func (r OrgMeteringListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// object used for sending metering query parameter request
type OrgMeteringListParamsQMeasurement struct {
	// this replaces all null values in an output stream with a non-null value that is
	// provided.
	Fill param.Field[float64] `query:"fill"`
	// end time range for the data, in ISO formate, yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	FromDate param.Field[string] `query:"fromDate"`
	// group by specific tags
	GroupBy param.Field[[]string] `query:"groupBy"`
	// time period to aggregate the data over with, in seconds. If none provided, raw
	// data will be returned.
	PeriodInSeconds param.Field[float64] `query:"periodInSeconds"`
	// start time range for the data, in ISO formate, yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	ToDate param.Field[string]                                 `query:"toDate"`
	Type   param.Field[OrgMeteringListParamsQMeasurementsType] `query:"type"`
}

// URLQuery serializes [OrgMeteringListParamsQMeasurement]'s query parameters as
// `url.Values`.
func (r OrgMeteringListParamsQMeasurement) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgMeteringListParamsQMeasurementsType string

const (
	OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily                   OrgMeteringListParamsQMeasurementsType = "EGX_GPU_UTILIZATION_DAILY"
	OrgMeteringListParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily          OrgMeteringListParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_DAILY"
	OrgMeteringListParamsQMeasurementsTypeEgxLogStorageUtilizationDaily            OrgMeteringListParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_DAILY"
	OrgMeteringListParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily   OrgMeteringListParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_DAILY"
	OrgMeteringListParamsQMeasurementsTypeRegistryStorageUtilizationDaily          OrgMeteringListParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_DAILY"
	OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationMonthly                 OrgMeteringListParamsQMeasurementsType = "EGX_GPU_UTILIZATION_MONTHLY"
	OrgMeteringListParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly        OrgMeteringListParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_MONTHLY"
	OrgMeteringListParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly          OrgMeteringListParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_MONTHLY"
	OrgMeteringListParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly OrgMeteringListParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_MONTHLY"
	OrgMeteringListParamsQMeasurementsTypeRegistryStorageUtilizationMonthly        OrgMeteringListParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_MONTHLY"
)

func (r OrgMeteringListParamsQMeasurementsType) IsKnown() bool {
	switch r {
	case OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily, OrgMeteringListParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily, OrgMeteringListParamsQMeasurementsTypeEgxLogStorageUtilizationDaily, OrgMeteringListParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily, OrgMeteringListParamsQMeasurementsTypeRegistryStorageUtilizationDaily, OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationMonthly, OrgMeteringListParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly, OrgMeteringListParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly, OrgMeteringListParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly, OrgMeteringListParamsQMeasurementsTypeRegistryStorageUtilizationMonthly:
		return true
	}
	return false
}
