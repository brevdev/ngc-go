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

func TestSuperAdminUserOrgUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SuperAdminUser.Orgs.Users.New(
		context.TODO(),
		"org-name",
		ngc.SuperAdminUserOrgUserNewParams{
			Email:                    ngc.F("xxxxxx"),
			IdpID:                    ngc.F("idp-id"),
			SendEmail:                ngc.F(true),
			EmailOptIn:               ngc.F(true),
			EulaAccepted:             ngc.F(true),
			Name:                     ngc.F("x"),
			RoleType:                 ngc.F("roleType"),
			RoleTypes:                ngc.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: ngc.F("salesforceContactJobRole"),
			UserMetadata: ngc.F(ngc.SuperAdminUserOrgUserNewParamsUserMetadata{
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

func TestSuperAdminUserOrgUserAddWithOptionalParams(t *testing.T) {
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
	_, err := client.SuperAdminUser.Orgs.Users.Add(
		context.TODO(),
		"org-name",
		"id",
		ngc.SuperAdminUserOrgUserAddParams{
			UserRoleDefaultsToRegistryRead: ngc.F("user role, defaults to REGISTRY_READ"),
			Ncid:                           ngc.F("ncid"),
			VisitorID:                      ngc.F("VisitorID"),
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

func TestSuperAdminUserOrgUserRemove(t *testing.T) {
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
	_, err := client.SuperAdminUser.Orgs.Users.Remove(
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
