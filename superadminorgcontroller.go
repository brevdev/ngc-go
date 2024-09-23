// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/ngc-go/internal/requestconfig"
	"github.com/stainless-sdks/ngc-go/option"
)

// SuperAdminOrgControllerService contains methods and other services that help
// with interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminOrgControllerService] method instead.
type SuperAdminOrgControllerService struct {
	Options []option.RequestOption
	Org     *SuperAdminOrgControllerOrgService
}

// NewSuperAdminOrgControllerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperAdminOrgControllerService(opts ...option.RequestOption) (r *SuperAdminOrgControllerService) {
	r = &SuperAdminOrgControllerService{}
	r.Options = opts
	r.Org = NewSuperAdminOrgControllerOrgService(opts...)
	return
}

// Backfill Orgs to Kratos
func (r *SuperAdminOrgControllerService) BackfillOrgsToKratos(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/backfill-orgs-to-kratos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
