// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/testutil"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

func TestOrgMeteringGpupeakListWithOptionalParams(t *testing.T) {
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
	_, err := client.Org.Meterings.Gpupeak.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgMeteringGpupeakListParams{
			TheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS: nvidiagpucloud.F(nvidiagpucloud.OrgMeteringGpupeakListParamsTheToDateInISO8601FormatIncludingTimeZoneInformationYyyyMmDdTHhMmSS{
				Sssz: nvidiagpucloud.F(time.Now()),
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
