// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/brevdev/ngc-go/internal/apiquery"
	"github.com/brevdev/ngc-go/internal/param"
	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// AdminOrgRegistryMeteringService contains methods and other services that help
// with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgRegistryMeteringService] method instead.
type AdminOrgRegistryMeteringService struct {
	Options []option.RequestOption
}

// NewAdminOrgRegistryMeteringService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAdminOrgRegistryMeteringService(opts ...option.RequestOption) (r *AdminOrgRegistryMeteringService) {
	r = &AdminOrgRegistryMeteringService{}
	r.Options = opts
	return
}

// Run registry metering downsample
func (r *AdminOrgRegistryMeteringService) Downsample(ctx context.Context, orgName string, query AdminOrgRegistryMeteringDownsampleParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/registry/metering/downsample", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AdminOrgRegistryMeteringDownsampleParams struct {
	// request params for getting metering usage
	Q param.Field[AdminOrgRegistryMeteringDownsampleParamsQ] `query:"q,required"`
}

// URLQuery serializes [AdminOrgRegistryMeteringDownsampleParams]'s query
// parameters as `url.Values`.
func (r AdminOrgRegistryMeteringDownsampleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// request params for getting metering usage
type AdminOrgRegistryMeteringDownsampleParamsQ struct {
	Measurements param.Field[[]AdminOrgRegistryMeteringDownsampleParamsQMeasurement] `query:"measurements"`
}

// URLQuery serializes [AdminOrgRegistryMeteringDownsampleParamsQ]'s query
// parameters as `url.Values`.
func (r AdminOrgRegistryMeteringDownsampleParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// object used for sending metering query parameter request
type AdminOrgRegistryMeteringDownsampleParamsQMeasurement struct {
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
	ToDate param.Field[string]                                                    `query:"toDate"`
	Type   param.Field[AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType] `query:"type"`
}

// URLQuery serializes [AdminOrgRegistryMeteringDownsampleParamsQMeasurement]'s
// query parameters as `url.Values`.
func (r AdminOrgRegistryMeteringDownsampleParamsQMeasurement) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType string

const (
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily                   AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "EGX_GPU_UTILIZATION_DAILY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily          AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_DAILY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationDaily            AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_DAILY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily   AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_DAILY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationDaily          AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_DAILY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationMonthly                 AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "EGX_GPU_UTILIZATION_MONTHLY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly        AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_MONTHLY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly          AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_MONTHLY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_MONTHLY"
	AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationMonthly        AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_MONTHLY"
)

func (r AdminOrgRegistryMeteringDownsampleParamsQMeasurementsType) IsKnown() bool {
	switch r {
	case AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationDaily, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationDaily, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationMonthly, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly, AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationMonthly:
		return true
	}
	return false
}
