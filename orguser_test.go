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

func TestOrgUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.New(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgUserNewParams{
			Email:                    nvidiagpucloud.F("xxxxxx"),
			IdpID:                    nvidiagpucloud.F("idp-id"),
			SendEmail:                nvidiagpucloud.F(true),
			EmailOptIn:               nvidiagpucloud.F(true),
			EulaAccepted:             nvidiagpucloud.F(true),
			Name:                     nvidiagpucloud.F("x"),
			RoleType:                 nvidiagpucloud.F("roleType"),
			RoleTypes:                nvidiagpucloud.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: nvidiagpucloud.F("salesforceContactJobRole"),
			UserMetadata: nvidiagpucloud.F(nvidiagpucloud.OrgUserNewParamsUserMetadata{
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

func TestOrgUserGet(t *testing.T) {
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
	_, err := client.Orgs.Users.Get(
		context.TODO(),
		"org-name",
		"user-email-or-id",
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrgUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgUserListParams{
			ExcludeFromTeam: nvidiagpucloud.F("exclude-from-team"),
			PageNumber:      nvidiagpucloud.F(int64(0)),
			PageSize:        nvidiagpucloud.F(int64(0)),
			Q: nvidiagpucloud.F(nvidiagpucloud.OrgUserListParamsQ{
				Fields: nvidiagpucloud.F([]string{"string", "string", "string"}),
				Filters: nvidiagpucloud.F([]nvidiagpucloud.OrgUserListParamsQFilter{{
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
				OrderBy: nvidiagpucloud.F([]nvidiagpucloud.OrgUserListParamsQOrderBy{{
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserListParamsQOrderByValueAsc),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserListParamsQOrderByValueAsc),
				}, {
					Field: nvidiagpucloud.F("field"),
					Value: nvidiagpucloud.F(nvidiagpucloud.OrgUserListParamsQOrderByValueAsc),
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

func TestOrgUserDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.Delete(
		context.TODO(),
		"org-name",
		"id",
		nvidiagpucloud.OrgUserDeleteParams{
			Anonymize: nvidiagpucloud.F(true),
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

func TestOrgUserAddRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.AddRole(
		context.TODO(),
		"org-name",
		"id",
		nvidiagpucloud.OrgUserAddRoleParams{
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

func TestOrgUserNcaInvitationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.NcaInvitations(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgUserNcaInvitationsParams{
			Email:                  nvidiagpucloud.F("xxxxxxx"),
			InvitationExpirationIn: nvidiagpucloud.F(int64(0)),
			InviteAs:               nvidiagpucloud.F(nvidiagpucloud.OrgUserNcaInvitationsParamsInviteAsAdmin),
			Message:                nvidiagpucloud.F("message"),
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

func TestOrgUserRemoveRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.RemoveRole(
		context.TODO(),
		"org-name",
		"user-email-or-id",
		nvidiagpucloud.OrgUserRemoveRoleParams{
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
