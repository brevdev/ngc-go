// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brevdev/ngc-go"
	"github.com/brevdev/ngc-go/option"
)

func TestSuperAdminOrgControllerOrgGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.SuperAdminOrgControllers.Org.Get(context.TODO(), "org-name")
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestSuperAdminOrgControllerOrgUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.SuperAdminOrgControllers.Org.Update(
		context.TODO(),
		"org-name",
		ngc.SuperAdminOrgControllerOrgUpdateParams{
			AlternateContact: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsAlternateContact{
				Email:    ngc.F("xxxxxx"),
				FullName: ngc.F("fullName"),
			}),
			CompanyName: ngc.F("companyName"),
			Description: ngc.F("description"),
			DisplayName: ngc.F("x"),
			IdpID:       ngc.F("idpId"),
			InfinityManagerSettings: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            ngc.F(true),
				InfinityManagerEnableTeamOverride: ngc.F(true),
			}),
			IsDatasetServiceEnabled:                 ngc.F(true),
			IsInternal:                              ngc.F(true),
			IsQuickStartEnabled:                     ngc.F(true),
			IsRegistrySseEnabled:                    ngc.F(true),
			IsSecretsManagerServiceEnabled:          ngc.F(true),
			IsSecureCredentialSharingServiceEnabled: ngc.F(true),
			IsSeparateInfluxDBUsed:                  ngc.F(true),
			OrgOwner: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsOrgOwner{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}),
			OrgOwners: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsOrgOwner{{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}, {
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}, {
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}}),
			PecName:   ngc.F("pecName"),
			PecSfdcID: ngc.F("pecSfdcId"),
			ProductEnablements: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablement{{
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}, {
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}, {
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}}),
			ProductSubscriptions: ngc.F([]ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscription{{
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}}),
			RepoScanSettings: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       ngc.F(true),
				RepoScanByDefault:           ngc.F(true),
				RepoScanEnabled:             ngc.F(true),
				RepoScanEnableNotifications: ngc.F(true),
				RepoScanEnableTeamOverride:  ngc.F(true),
				RepoScanShowResults:         ngc.F(true),
			}),
			Type: ngc.F(ngc.SuperAdminOrgControllerOrgUpdateParamsTypeUnknown),
		},
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
