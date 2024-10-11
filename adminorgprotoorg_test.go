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

func TestAdminOrgProtoOrgNewWithOptionalParams(t *testing.T) {
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
	resp, err := client.Admin.Orgs.ProtoOrg.New(context.TODO(), nvidiagpucloud.AdminOrgProtoOrgNewParams{
		Country:     nvidiagpucloud.F("country"),
		Description: nvidiagpucloud.F("description"),
		DisplayName: nvidiagpucloud.F("x"),
		Initiator:   nvidiagpucloud.F("initiator"),
		IsInternal:  nvidiagpucloud.F(true),
		Name:        nvidiagpucloud.F("xx"),
		NcaID:       nvidiagpucloud.F("ncaId"),
		NcaNumber:   nvidiagpucloud.F("ncaNumber"),
		OrgOwner: nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsOrgOwner{
			Email:       nvidiagpucloud.F("email"),
			FullName:    nvidiagpucloud.F("x"),
			IdpID:       nvidiagpucloud.F("idpId"),
			StarfleetID: nvidiagpucloud.F("starfleetId"),
		}),
		PecName:   nvidiagpucloud.F("pecName"),
		PecSfdcID: nvidiagpucloud.F("pecSfdcId"),
		ProductEnablements: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablement{{
			ProductName:    nvidiagpucloud.F("productName"),
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: nvidiagpucloud.F("expirationDate"),
			PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgProtoOrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscription{{
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        nvidiagpucloud.F("productName"),
			ID:                 nvidiagpucloud.F("id"),
			EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     nvidiagpucloud.F("expirationDate"),
			StartDate:          nvidiagpucloud.F("startDate"),
			Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		ProtoOrgID:                nvidiagpucloud.F("protoOrgId"),
		SalesforceAccountIndustry: nvidiagpucloud.F("salesforceAccountIndustry"),
		SendEmail:                 nvidiagpucloud.F(true),
		Type:                      nvidiagpucloud.F(nvidiagpucloud.AdminOrgProtoOrgNewParamsTypeUnknown),
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
