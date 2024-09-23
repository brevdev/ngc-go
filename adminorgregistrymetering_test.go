// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stainless-sdks/ngc-go"
	"github.com/stainless-sdks/ngc-go/option"
)

func TestAdminOrgRegistryMeteringDownsampleWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.Admin.Org.Registry.Metering.Downsample(
		context.TODO(),
		"org-name",
		ngc.AdminOrgRegistryMeteringDownsampleParams{
			Q: ngc.F(ngc.AdminOrgRegistryMeteringDownsampleParamsQ{
				Measurements: ngc.F([]ngc.AdminOrgRegistryMeteringDownsampleParamsQMeasurement{{
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.AdminOrgRegistryMeteringDownsampleParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}}),
			}),
		},
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
