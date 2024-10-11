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

func TestAdminOrgTeamUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Org.Team.Users.New(
		context.TODO(),
		"org-name",
		"team-name",
		nvidiagpucloud.AdminOrgTeamUserNewParams{
			Email:                    nvidiagpucloud.F("xxxxxx"),
			IdpID:                    nvidiagpucloud.F("idp-id"),
			SendEmail:                nvidiagpucloud.F(true),
			EmailOptIn:               nvidiagpucloud.F(true),
			EulaAccepted:             nvidiagpucloud.F(true),
			Name:                     nvidiagpucloud.F("x"),
			RoleType:                 nvidiagpucloud.F("roleType"),
			RoleTypes:                nvidiagpucloud.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: nvidiagpucloud.F("salesforceContactJobRole"),
			UserMetadata: nvidiagpucloud.F(nvidiagpucloud.AdminOrgTeamUserNewParamsUserMetadata{
				Company:    nvidiagpucloud.F("company"),
				CompanyURL: nvidiagpucloud.F("companyUrl"),
				Country:    nvidiagpucloud.F("country"),
				FirstName:  nvidiagpucloud.F("firstName"),
				Industry:   nvidiagpucloud.F("industry"),
				Interest:   nvidiagpucloud.F([]string{"string", "string", "string"}),
				LastName:   nvidiagpucloud.F("lastName"),
				Role:       nvidiagpucloud.F("role"),
			}),
			Ncid:      nvidiagpucloud.F("ncid"),
			VisitorID: nvidiagpucloud.F("VisitorID"),
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

func TestAdminOrgTeamUserAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Org.Team.Users.Add(
		context.TODO(),
		"org-name",
		"team-name",
		"id",
		nvidiagpucloud.AdminOrgTeamUserAddParams{
			UserRoleDefaultsToRegistryRead: nvidiagpucloud.F("user role, defaults to REGISTRY_READ"),
			Ncid:                           nvidiagpucloud.F("ncid"),
			VisitorID:                      nvidiagpucloud.F("VisitorID"),
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

func TestAdminOrgTeamUserAddRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Org.Team.Users.AddRole(
		context.TODO(),
		"org-name",
		"team-name",
		"id",
		nvidiagpucloud.AdminOrgTeamUserAddRoleParams{
			Roles: nvidiagpucloud.F([]string{"string", "string", "string"}),
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
