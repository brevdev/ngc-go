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

func TestSuperAdminOrgOrgStatusNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.SuperAdminOrg.OrgStatus.New(
		context.TODO(),
		"org-name",
		ngc.SuperAdminOrgOrgStatusNewParams{
			CreateSubscription: ngc.F(true),
			ProductEnablement: ngc.F(ngc.SuperAdminOrgOrgStatusNewParamsProductEnablement{
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.SuperAdminOrgOrgStatusNewParamsProductEnablementTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.SuperAdminOrgOrgStatusNewParamsProductEnablementPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
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
