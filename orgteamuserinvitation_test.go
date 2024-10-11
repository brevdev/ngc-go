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

func TestOrgTeamUserInvitationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Org.Team.Users.Invitations.List(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.OrgTeamUserInvitationListParams{
			OrderBy:    ngc.F(ngc.OrgTeamUserInvitationListParamsOrderByNameAsc),
			PageNumber: ngc.F(int64(0)),
			PageSize:   ngc.F(int64(0)),
			Q: ngc.F(ngc.OrgTeamUserInvitationListParamsQ{
				Fields: ngc.F([]string{"string", "string", "string"}),
				Filters: ngc.F([]ngc.OrgTeamUserInvitationListParamsQFilter{{
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
				OrderBy: ngc.F([]ngc.OrgTeamUserInvitationListParamsQOrderBy{{
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgTeamUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgTeamUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgTeamUserInvitationListParamsQOrderByValueAsc),
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

func TestOrgTeamUserInvitationDelete(t *testing.T) {
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
	_, err := client.Org.Team.Users.Invitations.Delete(
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

func TestOrgTeamUserInvitationInviteResend(t *testing.T) {
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
	_, err := client.Org.Team.Users.Invitations.InviteResend(
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
