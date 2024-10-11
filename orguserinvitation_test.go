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

func TestOrgUserInvitationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Org.Users.Invitations.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgUserInvitationListParams{
			OrderBy:    nvidiagpucloud.F(nvidiagpucloud.OrgUserInvitationListParamsOrderByNameAsc),
			PageNumber: nvidiagpucloud.F(int64(0)),
			PageSize:   nvidiagpucloud.F(int64(0)),
			Q: nvidiagpucloud.F(nvidiagpucloud.OrgUserInvitationListParamsQ{
				Fields: nvidiagpucloud.F([]string{"string", "string", "string"}),
				Filters: nvidiagpucloud.F([]nvidiagpucloud.OrgUserInvitationListParamsQFilter{{
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F("value"),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F("value"),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F("value"),
				}}),
				GroupBy: nvidiagpucloud.F("groupBy"),
				OrderBy: nvidiagpucloud.F([]nvidiagpucloud.OrgUserInvitationListParamsQOrderBy{{
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserInvitationListParamsQOrderByValueAsc),
				}}),
				Page:        nvidiagpucloud.F(int64(0)),
				PageSize:    nvidiagpucloud.F(int64(0)),
				Query:       nvidiagpucloud.F("query"),
				QueryFields: nvidiagpucloud.F([]string{"string", "string", "string"}),
				ScoredSize:  nvidiagpucloud.F(int64(0)),
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

func TestOrgUserInvitationResendInvitationEmail(t *testing.T) {
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
	_, err := client.Org.Users.Invitations.ResendInvitationEmail(
		context.TODO(),
		"org-name",
		"id",
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
