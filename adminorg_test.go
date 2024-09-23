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

func TestAdminOrgNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.Admin.Orgs.New(context.TODO(), ngc.AdminOrgNewParams{
		Country:     ngc.F("country"),
		Description: ngc.F("description"),
		DisplayName: ngc.F("x"),
		Initiator:   ngc.F("initiator"),
		IsInternal:  ngc.F(true),
		Name:        ngc.F("xx"),
		NcaID:       ngc.F("ncaId"),
		NcaNumber:   ngc.F("ncaNumber"),
		OrgOwner: ngc.F(ngc.AdminOrgNewParamsOrgOwner{
			Email:       ngc.F("email"),
			FullName:    ngc.F("x"),
			IdpID:       ngc.F("idpId"),
			StarfleetID: ngc.F("starfleetId"),
		}),
		PecName:   ngc.F("pecName"),
		PecSfdcID: ngc.F("pecSfdcId"),
		ProductEnablements: ngc.F([]ngc.AdminOrgNewParamsProductEnablement{{
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: ngc.F([]ngc.AdminOrgNewParamsProductSubscription{{
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		ProtoOrgID:                ngc.F("protoOrgId"),
		SalesforceAccountIndustry: ngc.F("salesforceAccountIndustry"),
		SendEmail:                 ngc.F(true),
		Type:                      ngc.F(ngc.AdminOrgNewParamsTypeUnknown),
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

func TestAdminOrgListWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.Admin.Orgs.List(context.TODO(), ngc.AdminOrgListParams{
		FilterUsingOrgDisplayName: ngc.F("Filter using org display name"),
		FilterUsingOrgOwnerEmail: ngc.F(ngc.AdminOrgListParamsFilterUsingOrgOwnerEmail{
			EmailShouldBeBase64Encoded: ngc.F(" Email should be base-64-encoded"),
		}),
		FilterUsingOrgOwnerName: ngc.F("Filter using org owner name"),
		OrgDesc:                 ngc.F("org-desc"),
		OrgName:                 ngc.F("org-name"),
		OrgType:                 ngc.F(ngc.AdminOrgListParamsOrgTypeUnknown),
		PageNumber:              ngc.F(int64(0)),
		PageSize:                ngc.F(int64(0)),
		PecID:                   ngc.F("pec-id"),
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
