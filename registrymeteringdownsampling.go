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
)

// RegistryMeteringDownsamplingService contains methods and other services that
// help with interacting with the nvidia-gpu-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryMeteringDownsamplingService] method instead.
type RegistryMeteringDownsamplingService struct {
	Options []option.RequestOption
}

// NewRegistryMeteringDownsamplingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRegistryMeteringDownsamplingService(opts ...option.RequestOption) (r *RegistryMeteringDownsamplingService) {
	r = &RegistryMeteringDownsamplingService{}
	r.Options = opts
	return
}

// Run registry metering downsample
func (r *RegistryMeteringDownsamplingService) Downsample(ctx context.Context, orgName string, query RegistryMeteringDownsamplingDownsampleParams, opts ...option.RequestOption) (res *http.Response, err error) {
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

type RegistryMeteringDownsamplingDownsampleParams struct {
	// request params for getting metering usage
	Q param.Field[RegistryMeteringDownsamplingDownsampleParamsQ] `query:"q,required"`
}

// URLQuery serializes [RegistryMeteringDownsamplingDownsampleParams]'s query
// parameters as `url.Values`.
func (r RegistryMeteringDownsamplingDownsampleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// request params for getting metering usage
type RegistryMeteringDownsamplingDownsampleParamsQ struct {
	Measurements param.Field[[]RegistryMeteringDownsamplingDownsampleParamsQMeasurement] `query:"measurements"`
}

// URLQuery serializes [RegistryMeteringDownsamplingDownsampleParamsQ]'s query
// parameters as `url.Values`.
func (r RegistryMeteringDownsamplingDownsampleParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// object used for sending metering query parameter request
type RegistryMeteringDownsamplingDownsampleParamsQMeasurement struct {
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
	ToDate param.Field[string]                                                        `query:"toDate"`
	Type   param.Field[RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType] `query:"type"`
}

// URLQuery serializes [RegistryMeteringDownsamplingDownsampleParamsQMeasurement]'s
// query parameters as `url.Values`.
func (r RegistryMeteringDownsamplingDownsampleParamsQMeasurement) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType string

const (
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily                   RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "EGX_GPU_UTILIZATION_DAILY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily          RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_DAILY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationDaily            RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_DAILY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily   RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_DAILY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationDaily          RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_DAILY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationMonthly                 RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "EGX_GPU_UTILIZATION_MONTHLY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly        RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_MONTHLY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly          RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_MONTHLY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_MONTHLY"
	RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationMonthly        RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_MONTHLY"
)

func (r RegistryMeteringDownsamplingDownsampleParamsQMeasurementsType) IsKnown() bool {
	switch r {
	case RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationDaily, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationDaily, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationMonthly, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly, RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeRegistryStorageUtilizationMonthly:
		return true
	}
	return false
}
