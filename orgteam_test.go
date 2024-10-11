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

func TestOrgTeamGet(t *testing.T) {
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
	_, err := client.Orgs.Teams.Get(
		context.TODO(),
		"org-name",
		"team-name",
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrgTeamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Teams.Update(
		context.TODO(),
		"org-name",
		"team-name",
		nvidiagpucloud.OrgTeamUpdateParams{
			Description: nvidiagpucloud.F("description"),
			InfinityManagerSettings: nvidiagpucloud.F(nvidiagpucloud.OrgTeamUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            nvidiagpucloud.F(true),
				InfinityManagerEnableTeamOverride: nvidiagpucloud.F(true),
			}),
			RepoScanSettings: nvidiagpucloud.F(nvidiagpucloud.OrgTeamUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       nvidiagpucloud.F(true),
				RepoScanByDefault:           nvidiagpucloud.F(true),
				RepoScanEnabled:             nvidiagpucloud.F(true),
				RepoScanEnableNotifications: nvidiagpucloud.F(true),
				RepoScanEnableTeamOverride:  nvidiagpucloud.F(true),
				RepoScanShowResults:         nvidiagpucloud.F(true),
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

func TestOrgTeamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Teams.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgTeamListParams{
			PageNumber: nvidiagpucloud.F(int64(0)),
			PageSize:   nvidiagpucloud.F(int64(0)),
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

func TestOrgTeamDelete(t *testing.T) {
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
	_, err := client.Orgs.Teams.Delete(
		context.TODO(),
		"org-name",
		"team-name",
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
