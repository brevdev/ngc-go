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

func TestAdminOrgEnablementNewWithOptionalParams(t *testing.T) {
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
	resp, err := client.Admin.Org.Enablement.New(
		context.TODO(),
		"org-name",
		nvidiagpucloud.AdminOrgEnablementNewParams{
			CreateSubscription: nvidiagpucloud.F(true),
			ProductEnablement: nvidiagpucloud.F(nvidiagpucloud.AdminOrgEnablementNewParamsProductEnablement{
				ProductName:    nvidiagpucloud.F("productName"),
				Type:           nvidiagpucloud.F(nvidiagpucloud.AdminOrgEnablementNewParamsProductEnablementTypeNgcAdminEval),
				ExpirationDate: nvidiagpucloud.F("expirationDate"),
				PoDetails: nvidiagpucloud.F([]nvidiagpucloud.AdminOrgEnablementNewParamsProductEnablementPoDetail{{
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}, {
					EntitlementID: nvidiagpucloud.F("entitlementId"),
					PkID:          nvidiagpucloud.F("pkId"),
				}}),
			}),
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
