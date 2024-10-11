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

func TestOrgProtoOrgNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.ProtoOrg.New(context.TODO(), nvidiagpucloud.OrgProtoOrgNewParams{
		Country:     nvidiagpucloud.F("country"),
		Description: nvidiagpucloud.F("description"),
		DisplayName: nvidiagpucloud.F("x"),
		Initiator:   nvidiagpucloud.F("initiator"),
		IsInternal:  nvidiagpucloud.F(true),
		Name:        nvidiagpucloud.F("xx"),
		NcaID:       nvidiagpucloud.F("ncaId"),
		NcaNumber:   nvidiagpucloud.F("ncaNumber"),
		OrgOwner: nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsOrgOwner{
			Email:       nvidiagpucloud.F("email"),
			FullName:    nvidiagpucloud.F("x"),
			IdpID:       nvidiagpucloud.F("idpId"),
			StarfleetID: nvidiagpucloud.F("starfleetId"),
		}),
		PecName:   nvidiagpucloud.F("pecName"),
		PecSfdcID: nvidiagpucloud.F("pecSfdcId"),
		ProductEnablements: nvidiagpucloud.F([]nvidiagpucloud.OrgProtoOrgNewParamsProductEnablement{{
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.OrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.OrgProtoOrgNewParamsProductSubscription{{
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		ProtoOrgID:                nvidiagpucloud.F("protoOrgId"),
		SalesforceAccountIndustry: nvidiagpucloud.F("salesforceAccountIndustry"),
		SendEmail:                 nvidiagpucloud.F(true),
		Type:                      nvidiagpucloud.F(nvidiagpucloud.OrgProtoOrgNewParamsTypeUnknown),
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

func TestOrgProtoOrgValidate(t *testing.T) {
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
	_, err := client.Orgs.ProtoOrg.Validate(context.TODO(), nvidiagpucloud.OrgProtoOrgValidateParams{
		InvitationToken: nvidiagpucloud.F("invitation_token"),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
