// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brevdev/ngc-go"
	"github.com/brevdev/ngc-go/internal/testutil"
	"github.com/brevdev/ngc-go/option"
)

func TestOrgNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Orgs.New(context.TODO(), ngc.OrgNewParams{
		Country:     ngc.F("country"),
		Description: ngc.F("description"),
		DisplayName: ngc.F("x"),
		Initiator:   ngc.F("initiator"),
		IsInternal:  ngc.F(true),
		Name:        ngc.F("xx"),
		NcaID:       ngc.F("ncaId"),
		NcaNumber:   ngc.F("ncaNumber"),
		OrgOwner: ngc.F(ngc.OrgNewParamsOrgOwner{
			Email:       ngc.F("email"),
			FullName:    ngc.F("x"),
			IdpID:       ngc.F("idpId"),
			StarfleetID: ngc.F("starfleetId"),
		}),
		PecName:   ngc.F("pecName"),
		PecSfdcID: ngc.F("pecSfdcId"),
		ProductEnablements: ngc.F([]ngc.OrgNewParamsProductEnablement{{
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.OrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.OrgNewParamsProductEnablementsPoDetail{{
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
			Type:           ngc.F(ngc.OrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.OrgNewParamsProductEnablementsPoDetail{{
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
		ProductSubscriptions: ngc.F([]ngc.OrgNewParamsProductSubscription{{
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.OrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.OrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		ProtoOrgID:                ngc.F("protoOrgId"),
		SalesforceAccountIndustry: ngc.F("salesforceAccountIndustry"),
		SendEmail:                 ngc.F(true),
		Type:                      ngc.F(ngc.OrgNewParamsTypeUnknown),
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
}

func TestOrgGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Orgs.Get(context.TODO(), "org-name")
	if err != nil {
		var apierr *ngc.Error
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
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Orgs.Update(
		context.TODO(),
		"org-name",
		ngc.OrgUpdateParams{
			Description: ngc.F("description"),
			DisplayName: ngc.F("x"),
		},
	)
	if err != nil {
		var apierr *ngc.Error
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
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Orgs.List(context.TODO(), ngc.OrgListParams{
		FilterUsingOrgDisplayName: ngc.F("Filter using org display name"),
		FilterUsingOrgName:        ngc.F("Filter using org name"),
		FilterUsingOrgOwnerEmail:  ngc.F("Filter using org owner email"),
		FilterUsingOrgOwnerName:   ngc.F("Filter using org owner name"),
		FilterUsingPecID:          ngc.F("Filter using PEC ID"),
		PageNumber:                ngc.F(int64(0)),
		PageSize:                  ngc.F(int64(0)),
	})
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
