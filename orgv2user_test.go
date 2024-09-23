// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/ngc-go"
	"github.com/stainless-sdks/ngc-go/internal/testutil"
	"github.com/stainless-sdks/ngc-go/option"
)

func TestOrgV2UserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.V2.Users.New(
		context.TODO(),
		"org-name",
		ngc.OrgV2UserNewParams{
			Email:                    ngc.F("xxxxxx"),
			IdpID:                    ngc.F("idp-id"),
			SendEmail:                ngc.F(true),
			EmailOptIn:               ngc.F(true),
			EulaAccepted:             ngc.F(true),
			Name:                     ngc.F("x"),
			RoleType:                 ngc.F("roleType"),
			RoleTypes:                ngc.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: ngc.F("salesforceContactJobRole"),
			UserMetadata: ngc.F(ngc.OrgV2UserNewParamsUserMetadata{
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

func TestOrgV2UserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.V2.Users.List(
		context.TODO(),
		"org-name",
		ngc.OrgV2UserListParams{
			ExcludeFromTeam: ngc.F("exclude-from-team"),
			PageNumber:      ngc.F(int64(0)),
			PageSize:        ngc.F(int64(0)),
			Q: ngc.F(ngc.OrgV2UserListParamsQ{
				Fields: ngc.F([]string{"string", "string", "string"}),
				Filters: ngc.F([]ngc.OrgV2UserListParamsQFilter{{
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
				OrderBy: ngc.F([]ngc.OrgV2UserListParamsQOrderBy{{
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgV2UserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgV2UserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgV2UserListParamsQOrderByValueAsc),
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
