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

func TestOrganizationTeamUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Teams.Users.New(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.OrganizationTeamUserNewParams{
			Email:                    ngc.F("xxxxxx"),
			IdpID:                    ngc.F("idp-id"),
			SendEmail:                ngc.F(true),
			EmailOptIn:               ngc.F(true),
			EulaAccepted:             ngc.F(true),
			Name:                     ngc.F("x"),
			RoleType:                 ngc.F("roleType"),
			RoleTypes:                ngc.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: ngc.F("salesforceContactJobRole"),
			UserMetadata: ngc.F(ngc.OrganizationTeamUserNewParamsUserMetadata{
				Company:    ngc.F("company"),
				CompanyURL: ngc.F("companyUrl"),
				Country:    ngc.F("country"),
				FirstName:  ngc.F("firstName"),
				Industry:   ngc.F("industry"),
				Interest:   ngc.F([]string{"string", "string", "string"}),
				LastName:   ngc.F("lastName"),
				Role:       ngc.F("role"),
			}),
			Ncid:      ngc.F("ncid"),
			VisitorID: ngc.F("VisitorID"),
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

func TestOrganizationTeamUserGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Teams.Users.Get(
		context.TODO(),
		"org-name",
		"team-name",
		"id",
		ngc.OrganizationTeamUserGetParams{
			IdpID: ngc.F("idp-id"),
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

func TestOrganizationTeamUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Teams.Users.List(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.OrganizationTeamUserListParams{
			PageNumber: ngc.F(int64(0)),
			PageSize:   ngc.F(int64(0)),
			Q: ngc.F(ngc.OrganizationTeamUserListParamsQ{
				Fields: ngc.F([]string{"string", "string", "string"}),
				Filters: ngc.F([]ngc.OrganizationTeamUserListParamsQFilter{{
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}}),
				GroupBy: ngc.F("groupBy"),
				OrderBy: ngc.F([]ngc.OrganizationTeamUserListParamsQOrderBy{{
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrganizationTeamUserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrganizationTeamUserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrganizationTeamUserListParamsQOrderByValueAsc),
				}}),
				Page:        ngc.F(int64(0)),
				PageSize:    ngc.F(int64(0)),
				Query:       ngc.F("query"),
				QueryFields: ngc.F([]string{"string", "string", "string"}),
				ScoredSize:  ngc.F(int64(0)),
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

func TestOrganizationTeamUserDelete(t *testing.T) {
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
	_, err := client.Organizations.Teams.Users.Delete(
		context.TODO(),
		"org-name",
		"team-name",
		"id",
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
