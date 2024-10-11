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

func TestOrgNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.New(context.TODO(), nvidiagpucloud.OrgNewParams{
		Country:     nvidiagpucloud.F("country"),
		Description: nvidiagpucloud.F("description"),
		DisplayName: nvidiagpucloud.F("x"),
		Initiator:   nvidiagpucloud.F("initiator"),
		IsInternal:  nvidiagpucloud.F(true),
		Name:        nvidiagpucloud.F("xx"),
		NcaID:       nvidiagpucloud.F("ncaId"),
		NcaNumber:   nvidiagpucloud.F("ncaNumber"),
		OrgOwner: nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsOrgOwner{
			Email:       nvidiagpucloud.F("email"),
			FullName:    nvidiagpucloud.F("x"),
			IdpID:       nvidiagpucloud.F("idpId"),
			StarfleetID: nvidiagpucloud.F("starfleetId"),
		}),
		PecName:   nvidiagpucloud.F("pecName"),
		PecSfdcID: nvidiagpucloud.F("pecSfdcId"),
		ProductEnablements: nvidiagpucloud.F([]nvidiagpucloud.OrgNewParamsProductEnablement{{
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.OrgNewParamsProductSubscription{{
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		ProtoOrgID:                nvidiagpucloud.F("protoOrgId"),
		SalesforceAccountIndustry: nvidiagpucloud.F("salesforceAccountIndustry"),
		SendEmail:                 nvidiagpucloud.F(true),
		Type:                      nvidiagpucloud.F(nvidiagpucloud.OrgNewParamsTypeUnknown),
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
}

func TestOrgGet(t *testing.T) {
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
	_, err := client.Orgs.Get(context.TODO(), "org-name")
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrgUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Update(
		context.TODO(),
		"org-name",
		nvidiagpucloud.OrgUpdateParams{
			Description: nvidiagpucloud.F("description"),
			DisplayName: nvidiagpucloud.F("x"),
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

func TestOrgListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.List(context.TODO(), nvidiagpucloud.OrgListParams{
		FilterUsingOrgDisplayName: nvidiagpucloud.F("Filter using org display name"),
		FilterUsingOrgName:        nvidiagpucloud.F("Filter using org name"),
		FilterUsingOrgOwnerEmail:  nvidiagpucloud.F("Filter using org owner email"),
		FilterUsingOrgOwnerName:   nvidiagpucloud.F("Filter using org owner name"),
		FilterUsingPecID:          nvidiagpucloud.F("Filter using PEC ID"),
		PageNumber:                nvidiagpucloud.F(int64(0)),
		PageSize:                  nvidiagpucloud.F(int64(0)),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
