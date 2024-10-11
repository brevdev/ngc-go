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

func TestAdminOrgTeamGet(t *testing.T) {
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
	resp, err := client.Admin.Org.Teams.Get(
		context.TODO(),
		"org-name",
		"team-name",
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

func TestAdminOrgTeamUpdateWithOptionalParams(t *testing.T) {
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
	resp, err := client.Admin.Org.Teams.Update(
		context.TODO(),
		"org-name",
		"team-name",
		nvidiagpucloud.AdminOrgTeamUpdateParams{
			Description: nvidiagpucloud.F("description"),
			InfinityManagerSettings: nvidiagpucloud.F(nvidiagpucloud.AdminOrgTeamUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            nvidiagpucloud.F(true),
				InfinityManagerEnableTeamOverride: nvidiagpucloud.F(true),
			}),
			RepoScanSettings: nvidiagpucloud.F(nvidiagpucloud.AdminOrgTeamUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       nvidiagpucloud.F(true),
				RepoScanByDefault:           nvidiagpucloud.F(true),
				RepoScanEnabled:             nvidiagpucloud.F(true),
				RepoScanEnableNotifications: nvidiagpucloud.F(true),
				RepoScanEnableTeamOverride:  nvidiagpucloud.F(true),
				RepoScanShowResults:         nvidiagpucloud.F(true),
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

func TestAdminOrgTeamListWithOptionalParams(t *testing.T) {
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
	resp, err := client.Admin.Org.Teams.List(
		context.TODO(),
		"org-name",
		nvidiagpucloud.AdminOrgTeamListParams{
			PageNumber:        nvidiagpucloud.F(int64(0)),
			PageSize:          nvidiagpucloud.F(int64(0)),
			ScanAllowOverride: nvidiagpucloud.F(true),
			ScanByDefault:     nvidiagpucloud.F(true),
			ScanShowResults:   nvidiagpucloud.F(true),
			TeamName:          nvidiagpucloud.F("team-name"),
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
