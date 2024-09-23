// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brevdev/ngc-go"
	"github.com/brevdev/ngc-go/internal/testutil"
	"github.com/brevdev/ngc-go/option"
)

func TestOrgMeteringListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Orgs.Metering.List(
		context.TODO(),
		"org-name",
		ngc.OrgMeteringListParams{
			Q: ngc.F(ngc.OrgMeteringListParamsQ{
				Measurements: ngc.F([]ngc.OrgMeteringListParamsQMeasurement{{
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringListParamsQMeasurementsTypeEgxGPUUtilizationDaily),
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
}
