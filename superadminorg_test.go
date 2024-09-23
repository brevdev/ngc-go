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

	"github.com/stainless-sdks/ngc-go"
	"github.com/stainless-sdks/ngc-go/option"
)

func TestSuperAdminOrgNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.SuperAdminOrg.New(context.TODO(), ngc.SuperAdminOrgNewParams{
		OrgOwner: ngc.F(ngc.SuperAdminOrgNewParamsOrgOwner{
			Email:         ngc.F("email"),
			FullName:      ngc.F("x"),
			LastLoginDate: ngc.F("lastLoginDate"),
		}),
		Country:     ngc.F("country"),
		Description: ngc.F("description"),
		DisplayName: ngc.F("x"),
		IdpID:       ngc.F("idpId"),
		IsInternal:  ngc.F(true),
		Name:        ngc.F("xx"),
		PecName:     ngc.F("pecName"),
		PecSfdcID:   ngc.F("pecSfdcId"),
		ProductEnablements: ngc.F([]ngc.SuperAdminOrgNewParamsProductEnablement{{
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.SuperAdminOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.SuperAdminOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.SuperAdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.SuperAdminOrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: ngc.F([]ngc.SuperAdminOrgNewParamsProductSubscription{{
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.SuperAdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		SalesforceAccountIndustry: ngc.F("salesforceAccountIndustry"),
		SendEmail:                 ngc.F(true),
		Type:                      ngc.F(ngc.SuperAdminOrgNewParamsTypeUnknown),
		Ncid:                      ngc.F("ncid"),
		VisitorID:                 ngc.F("VisitorID"),
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
