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

func TestOrgTeamGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Org.Teams.Get(
		context.TODO(),
		"org-name",
		"team-name",
	)
	if err != nil {
		var apierr *ngc.Error
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
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Org.Teams.Update(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.OrgTeamUpdateParams{
			Description: ngc.F("description"),
			InfinityManagerSettings: ngc.F(ngc.OrgTeamUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            ngc.F(true),
				InfinityManagerEnableTeamOverride: ngc.F(true),
			}),
			RepoScanSettings: ngc.F(ngc.OrgTeamUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       ngc.F(true),
				RepoScanByDefault:           ngc.F(true),
				RepoScanEnabled:             ngc.F(true),
				RepoScanEnableNotifications: ngc.F(true),
				RepoScanEnableTeamOverride:  ngc.F(true),
				RepoScanShowResults:         ngc.F(true),
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

func TestOrgTeamDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Org.Teams.Delete(
		context.TODO(),
		"org-name",
		"team-name",
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
