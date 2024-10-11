// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"

	"github.com/brevdev/ngc-go/internal/requestconfig"
	"github.com/brevdev/ngc-go/option"
)

// SwaggerResourceConfigurationSecurityService contains methods and other services
// that help with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwaggerResourceConfigurationSecurityService] method instead.
type SwaggerResourceConfigurationSecurityService struct {
	Options []option.RequestOption
}

// NewSwaggerResourceConfigurationSecurityService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSwaggerResourceConfigurationSecurityService(opts ...option.RequestOption) (r *SwaggerResourceConfigurationSecurityService) {
	r = &SwaggerResourceConfigurationSecurityService{}
	r.Options = opts
	return
}

func (r *SwaggerResourceConfigurationSecurityService) New(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationSecurityService) Get(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationSecurityService) Update(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

func (r *SwaggerResourceConfigurationSecurityService) Delete(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "swagger-resources/configuration/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
