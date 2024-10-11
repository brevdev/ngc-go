// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stainless-sdks/nvidia-gpu-cloud-go"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/option"
)

func TestAdminOrgNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Admin.Orgs.New(context.TODO(), nvidiagpucloud.AdminOrgNewParams{
		OrgOwner: nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsOrgOwner{
			Email:         nvidiagpucloud.F("email"),
			FullName:      nvidiagpucloud.F("x"),
			LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
		}),
		Country:     nvidiagpucloud.F("country"),
		Description: nvidiagpucloud.F("description"),
		DisplayName: nvidiagpucloud.F("x"),
		IdpID:       nvidiagpucloud.F("idpId"),
		IsInternal:  nvidiagpucloud.F(true),
		Name:        nvidiagpucloud.F("xx"),
		PecName:     nvidiagpucloud.F("pecName"),
		PecSfdcID:   nvidiagpucloud.F("pecSfdcId"),
		ProductEnablements: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgNewParamsProductEnablement{{
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}}),
		}, {
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}}),
		}, {
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}, {
				EntitlementID: nvidiagpucloud.F("entitlementId"),
				PkID:          nvidiagpucloud.F("pkId"),
			}}),
		}}),
		ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgNewParamsProductSubscription{{
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		SalesforceAccountIndustry: nvidiagpucloud.F("salesforceAccountIndustry"),
		SendEmail:                 nvidiagpucloud.F(true),
		Type:                      nvidiagpucloud.F(nvidiagpucloud.AdminOrgNewParamsTypeUnknown),
		Ncid:                      nvidiagpucloud.F("ncid"),
		VisitorID:                 nvidiagpucloud.F("VisitorID"),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Admin.Orgs.Get(context.TODO(), "org-name")
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Admin.Orgs.Update(
		context.TODO(),
		"org-name",
		nvidiagpucloud.AdminOrgUpdateParams{
			AlternateContact: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsAlternateContact{
				Email:    nvidiagpucloud.F("xxxxxx"),
				FullName: nvidiagpucloud.F("fullName"),
			}),
			CompanyName: nvidiagpucloud.F("companyName"),
			Description: nvidiagpucloud.F("description"),
			DisplayName: nvidiagpucloud.F("x"),
			IdpID:       nvidiagpucloud.F("idpId"),
			InfinityManagerSettings: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            nvidiagpucloud.F(true),
				InfinityManagerEnableTeamOverride: nvidiagpucloud.F(true),
			}),
			IsDatasetServiceEnabled:                 nvidiagpucloud.F(true),
			IsInternal:                              nvidiagpucloud.F(true),
			IsQuickStartEnabled:                     nvidiagpucloud.F(true),
			IsRegistrySseEnabled:                    nvidiagpucloud.F(true),
			IsSecretsManagerServiceEnabled:          nvidiagpucloud.F(true),
			IsSecureCredentialSharingServiceEnabled: nvidiagpucloud.F(true),
			IsSeparateInfluxDBUsed:                  nvidiagpucloud.F(true),
			OrgOwner: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsOrgOwner{
				Email:         nvidiagpucloud.F("email"),
				FullName:      nvidiagpucloud.F("x"),
				LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
			}),
			OrgOwners: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsOrgOwner{{
				Email:         nvidiagpucloud.F("email"),
				FullName:      nvidiagpucloud.F("x"),
				LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
			}, {
				Email:         nvidiagpucloud.F("email"),
				FullName:      nvidiagpucloud.F("x"),
				LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
			}, {
				Email:         nvidiagpucloud.F("email"),
				FullName:      nvidiagpucloud.F("x"),
				LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
			}}),
			PecName:   nvidiagpucloud.F("pecName"),
			PecSfdcID: nvidiagpucloud.F("pecSfdcId"),
			ProductEnablements: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsProductEnablement{{
				ProductName:    nvidiagpucloud.F("productName"),
				Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: nvidiagpucloud.F("expirationDate"),
				PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}}),
			}, {
				ProductName:    nvidiagpucloud.F("productName"),
				Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: nvidiagpucloud.F("expirationDate"),
				PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}}),
			}, {
				ProductName:    nvidiagpucloud.F("productName"),
				Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: nvidiagpucloud.F("expirationDate"),
				PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}}),
			}}),
			ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgUpdateParamsProductSubscription{{
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}}),
			RepoScanSettings: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       nvidiagpucloud.F(true),
				RepoScanByDefault:           nvidiagpucloud.F(true),
				RepoScanEnabled:             nvidiagpucloud.F(true),
				RepoScanEnableNotifications: nvidiagpucloud.F(true),
				RepoScanEnableTeamOverride:  nvidiagpucloud.F(true),
				RepoScanShowResults:         nvidiagpucloud.F(true),
			}),
			Type: nvidiagpucloud.F(nvidiagpucloud.AdminOrgUpdateParamsTypeUnknown),
		},
	)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgListWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Admin.Orgs.List(context.TODO(), nvidiagpucloud.AdminOrgListParams{
		FilterUsingOrgDisplayName: nvidiagpucloud.F("Filter using org display name"),
		FilterUsingOrgOwnerEmail: nvidiagpucloud.F(nvidiagpucloud.AdminOrgListParamsFilterUsingOrgOwnerEmail{
			EmailShouldBeBase64Encoded: nvidiagpucloud.F(" Email should be base-64-encoded"),
		}),
		FilterUsingOrgOwnerName: nvidiagpucloud.F("Filter using org owner name"),
		OrgDesc:                 nvidiagpucloud.F("org-desc"),
		OrgName:                 nvidiagpucloud.F("org-name"),
		OrgType:                 nvidiagpucloud.F(nvidiagpucloud.AdminOrgListParamsOrgTypeUnknown),
		PageNumber:              nvidiagpucloud.F(int64(0)),
		PageSize:                nvidiagpucloud.F(int64(0)),
		PecID:                   nvidiagpucloud.F("pec-id"),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgNcaIDsWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := nvidiagpucloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBearerToken("My Bearer Token"),
	)
	resp, err := client.Admin.Orgs.NcaIDs(context.TODO(), nvidiagpucloud.AdminOrgNcaIDsParams{
		NcaIDs:   nvidiagpucloud.F([]string{"string", "string", "string"}),
		OrgNames: nvidiagpucloud.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
