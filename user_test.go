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

func TestUserGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), nvidiagpucloud.UserGetParams{
		OrgName: nvidiagpucloud.F("org-name"),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Update(context.TODO(), nvidiagpucloud.UserUpdateParams{
		HasEmailOptIn:                        nvidiagpucloud.F(true),
		HasSignedAIFoundryPartnershipsEula:   nvidiagpucloud.F(true),
		HasSignedBaseCommandEula:             nvidiagpucloud.F(true),
		HasSignedBaseCommandManagerEula:      nvidiagpucloud.F(true),
		HasSignedBioNeMoEula:                 nvidiagpucloud.F(true),
		HasSignedContainerPublishingEula:     nvidiagpucloud.F(true),
		HasSignedCuOptEula:                   nvidiagpucloud.F(true),
		HasSignedEarth2Eula:                  nvidiagpucloud.F(true),
		HasSignedEgxEula:                     nvidiagpucloud.F(true),
		HasSignedEula:                        nvidiagpucloud.F(true),
		HasSignedFleetCommandEula:            nvidiagpucloud.F(true),
		HasSignedLlmEula:                     nvidiagpucloud.F(true),
		HasSignedNvaieeula:                   nvidiagpucloud.F(true),
		HasSignedNvidiaEula:                  nvidiagpucloud.F(true),
		HasSignedNvqceula:                    nvidiagpucloud.F(true),
		HasSignedOmniverseEula:               nvidiagpucloud.F(true),
		HasSignedPrivacyPolicy:               nvidiagpucloud.F(true),
		HasSignedThirdPartyRegistryShareEula: nvidiagpucloud.F(true),
		Name:                                 nvidiagpucloud.F("x"),
		UserMetadata: nvidiagpucloud.F(nvidiagpucloud.UserUpdateParamsUserMetadata{
			Company:    nvidiagpucloud.F("company"),
			CompanyURL: nvidiagpucloud.F("companyUrl"),
			Country:    nvidiagpucloud.F("country"),
			FirstName:  nvidiagpucloud.F("firstName"),
			Industry:   nvidiagpucloud.F("industry"),
			Interest:   nvidiagpucloud.F([]string{"string", "string", "string"}),
			LastName:   nvidiagpucloud.F("lastName"),
			Role:       nvidiagpucloud.F("role"),
		}),
	})
	if err != nil {
		var apierr *nvidiagpucloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
