// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/brevdev/ngc-go"
	"github.com/brevdev/ngc-go/internal/testutil"
	"github.com/brevdev/ngc-go/option"
)

func TestOrgMeteringGpupeakListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Metering.Gpupeak.List(
		context.TODO(),
		"org-name",
		ngc.OrgMeteringGpupeakListParams{
			TheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS: ngc.F(ngc.OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS{
				Sssz: ngc.F(time.Now()),
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
