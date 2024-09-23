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

func TestOrgUserInvitationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Org.Users.Invitations.List(
		context.TODO(),
		"org-name",
		ngc.OrgUserInvitationListParams{
			OrderBy:    ngc.F(ngc.OrgUserInvitationListParamsOrderByNameAsc),
			PageNumber: ngc.F(int64(0)),
			PageSize:   ngc.F(int64(0)),
			Q: ngc.F(ngc.OrgUserInvitationListParamsQ{
				Fields: ngc.F([]string{"string", "string", "string"}),
				Filters: ngc.F([]ngc.OrgUserInvitationListParamsQFilter{{
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
				OrderBy: ngc.F([]ngc.OrgUserInvitationListParamsQOrderBy{{
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserInvitationListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserInvitationListParamsQOrderByValueAsc),
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

func TestOrgUserInvitationDelete(t *testing.T) {
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
	_, err := client.Org.Users.Invitations.Delete(
		context.TODO(),
		"org-name",
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

func TestOrgUserInvitationInviteResend(t *testing.T) {
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
	_, err := client.Org.Users.Invitations.InviteResend(
		context.TODO(),
		"org-name",
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
