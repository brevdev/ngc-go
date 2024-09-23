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

func TestOrgV3UserNcaInvitationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.V3.Users.NcaInvitations.New(
		context.TODO(),
		"org-name",
		ngc.OrgV3UserNcaInvitationNewParams{
			Email:                  ngc.F("xxxxxxx"),
			InvitationExpirationIn: ngc.F(int64(0)),
			InviteAs:               ngc.F(ngc.OrgV3UserNcaInvitationNewParamsInviteAsAdmin),
			Message:                ngc.F("message"),
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
