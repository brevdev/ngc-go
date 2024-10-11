// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/testutil"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

func TestOrgMeteringListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Org.Meterings.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgMeteringListParams{
			Q: nvidiagpucloud.F(nvidiagpucloud.OrgMeteringListParamsQ{
				Measurements: nvidiagpucloud.F([]nvidiagpucloud.OrgMeteringListParamsQMeasurement{{
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}}),
			}),
		},
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
