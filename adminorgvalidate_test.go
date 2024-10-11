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

func TestAdminOrgValidateList(t *testing.T) {
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
	resp, err := client.Admin.Org.Validate.List(context.TODO(), ngc.AdminOrgValidateListParams{
		Q: ngc.F(ngc.AdminOrgValidateListParamsQ{
			OrgOwner: ngc.F(ngc.AdminOrgValidateListParamsQOrgOwner{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}),
			PecSfdcID: ngc.F("pecSfdcId"),
			ProductSubscriptions: ngc.F([]ngc.AdminOrgValidateListParamsQProductSubscription{{
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateListParamsQProductSubscriptionsTypeNgcAdminEval),
			}}),
		}),
	})
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
