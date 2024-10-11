// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

func TestRegistryMeteringDownsamplingDownsampleWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.RegistryMeteringDownsampling.Downsample(
		context.TODO(),
		"org-name",
		nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParams{
			Q: nvidiagpucloud.F(nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParamsQ{
				Measurements: nvidiagpucloud.F([]nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParamsQMeasurement{{
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            nvidiagpucloud.F(0.000000),
					FromDate:        nvidiagpucloud.F("fromDate"),
					GroupBy:         nvidiagpucloud.F([]string{"string", "string", "string"}),
					PeriodInSeconds: nvidiagpucloud.F(0.000000),
					ToDate:          nvidiagpucloud.F("toDate"),
					Type:            nvidiagpucloud.F(nvidiagpucloud.RegistryMeteringDownsamplingDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
