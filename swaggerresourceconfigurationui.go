// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// SwaggerResourceConfigurationUiService contains methods and other services that
// help with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwaggerResourceConfigurationUiService] method instead.
type SwaggerResourceConfigurationUiService struct {
	Options []option.RequestOption
}

// NewSwaggerResourceConfigurationUiService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSwaggerResourceConfigurationUiService(opts ...option.RequestOption) (r *SwaggerResourceConfigurationUiService) {
	r = &SwaggerResourceConfigurationUiService{}
	r.Options = opts
	return
}

func (r *SwaggerResourceConfigurationUiService) New(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/ui"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationUiService) Get(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/ui"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationUiService) Update(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/ui"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationUiService) Delete(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/ui"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
