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

func TestAdminOrgOrganizationValidate(t *testing.T) {
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
	resp, err := client.AdminOrg.Organizations.Validate(context.TODO(), nvidiagpucloud.AdminOrgOrganizationValidateParams{
		Q: nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQ{
			OrgOwner: nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQOrgOwner{
				Email:         nvidiagpucloud.F("email"),
				FullName:      nvidiagpucloud.F("x"),
				LastLoginDate: nvidiagpucloud.F("lastLoginDate"),
			}),
			PecSfdcID: nvidiagpucloud.F("pecSfdcId"),
			ProductSubscriptions: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscription{{
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        nvidiagpucloud.F("productName"),
				ID:                 nvidiagpucloud.F("id"),
				EmsEntitlementType: nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     nvidiagpucloud.F("expirationDate"),
				StartDate:          nvidiagpucloud.F("startDate"),
				Type:               nvidiagpucloud.F(nvidiagpucloud.AdminOrgOrganizationValidateParamsQProductSubscriptionsTypeNgcAdminEval),
			}}),
		}),
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
