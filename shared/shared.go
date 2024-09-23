// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/ngc-go/internal/apijson"
)

// This API is invoked by monitoring tools, other services and infrastructure to
// retrieve health status the targeted service, this is unprotected method
type Health struct {
	// object that describes health of the service
	Health        HealthHealth        `json:"health"`
	RequestStatus HealthRequestStatus `json:"requestStatus"`
	JSON          healthJSON          `json:"-"`
}

// healthJSON contains the JSON metadata for the struct [Health]
type healthJSON struct {
	Health        apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Health) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthJSON) RawJSON() string {
	return r.raw
}

// object that describes health of the service
type HealthHealth struct {
	// Enum that describes health of the service
	HealthCode HealthHealthHealthCode `json:"healthCode"`
	// Human readable description
	HealthCodeDescription string                 `json:"healthCodeDescription"`
	MetaData              []HealthHealthMetaData `json:"metaData"`
	JSON                  healthHealthJSON       `json:"-"`
}

// healthHealthJSON contains the JSON metadata for the struct [HealthHealth]
type healthHealthJSON struct {
	HealthCode            apijson.Field
	HealthCodeDescription apijson.Field
	MetaData              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *HealthHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthJSON) RawJSON() string {
	return r.raw
}

// Enum that describes health of the service
type HealthHealthHealthCode string

const (
	HealthHealthHealthCodeUnknown          HealthHealthHealthCode = "UNKNOWN"
	HealthHealthHealthCodeOk               HealthHealthHealthCode = "OK"
	HealthHealthHealthCodeUnderMaintenance HealthHealthHealthCode = "UNDER_MAINTENANCE"
	HealthHealthHealthCodeWarning          HealthHealthHealthCode = "WARNING"
	HealthHealthHealthCodeFailed           HealthHealthHealthCode = "FAILED"
)

func (r HealthHealthHealthCode) IsKnown() bool {
	switch r {
	case HealthHealthHealthCodeUnknown, HealthHealthHealthCodeOk, HealthHealthHealthCodeUnderMaintenance, HealthHealthHealthCodeWarning, HealthHealthHealthCodeFailed:
		return true
	}
	return false
}

type HealthHealthMetaData struct {
	Key   string                   `json:"key"`
	Value string                   `json:"value"`
	JSON  healthHealthMetaDataJSON `json:"-"`
}

// healthHealthMetaDataJSON contains the JSON metadata for the struct
// [HealthHealthMetaData]
type healthHealthMetaDataJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthHealthMetaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthMetaDataJSON) RawJSON() string {
	return r.raw
}

type HealthRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        HealthRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                        `json:"statusDescription"`
	JSON              healthRequestStatusJSON       `json:"-"`
}

// healthRequestStatusJSON contains the JSON metadata for the struct
// [HealthRequestStatus]
type healthRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HealthRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type HealthRequestStatusStatusCode string

const (
	HealthRequestStatusStatusCodeUnknown                    HealthRequestStatusStatusCode = "UNKNOWN"
	HealthRequestStatusStatusCodeSuccess                    HealthRequestStatusStatusCode = "SUCCESS"
	HealthRequestStatusStatusCodeUnauthorized               HealthRequestStatusStatusCode = "UNAUTHORIZED"
	HealthRequestStatusStatusCodePaymentRequired            HealthRequestStatusStatusCode = "PAYMENT_REQUIRED"
	HealthRequestStatusStatusCodeForbidden                  HealthRequestStatusStatusCode = "FORBIDDEN"
	HealthRequestStatusStatusCodeTimeout                    HealthRequestStatusStatusCode = "TIMEOUT"
	HealthRequestStatusStatusCodeExists                     HealthRequestStatusStatusCode = "EXISTS"
	HealthRequestStatusStatusCodeNotFound                   HealthRequestStatusStatusCode = "NOT_FOUND"
	HealthRequestStatusStatusCodeInternalError              HealthRequestStatusStatusCode = "INTERNAL_ERROR"
	HealthRequestStatusStatusCodeInvalidRequest             HealthRequestStatusStatusCode = "INVALID_REQUEST"
	HealthRequestStatusStatusCodeInvalidRequestVersion      HealthRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	HealthRequestStatusStatusCodeInvalidRequestData         HealthRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	HealthRequestStatusStatusCodeMethodNotAllowed           HealthRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	HealthRequestStatusStatusCodeConflict                   HealthRequestStatusStatusCode = "CONFLICT"
	HealthRequestStatusStatusCodeUnprocessableEntity        HealthRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	HealthRequestStatusStatusCodeTooManyRequests            HealthRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	HealthRequestStatusStatusCodeInsufficientStorage        HealthRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	HealthRequestStatusStatusCodeServiceUnavailable         HealthRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	HealthRequestStatusStatusCodePayloadTooLarge            HealthRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	HealthRequestStatusStatusCodeNotAcceptable              HealthRequestStatusStatusCode = "NOT_ACCEPTABLE"
	HealthRequestStatusStatusCodeUnavailableForLegalReasons HealthRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HealthRequestStatusStatusCodeBadGateway                 HealthRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r HealthRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case HealthRequestStatusStatusCodeUnknown, HealthRequestStatusStatusCodeSuccess, HealthRequestStatusStatusCodeUnauthorized, HealthRequestStatusStatusCodePaymentRequired, HealthRequestStatusStatusCodeForbidden, HealthRequestStatusStatusCodeTimeout, HealthRequestStatusStatusCodeExists, HealthRequestStatusStatusCodeNotFound, HealthRequestStatusStatusCodeInternalError, HealthRequestStatusStatusCodeInvalidRequest, HealthRequestStatusStatusCodeInvalidRequestVersion, HealthRequestStatusStatusCodeInvalidRequestData, HealthRequestStatusStatusCodeMethodNotAllowed, HealthRequestStatusStatusCodeConflict, HealthRequestStatusStatusCodeUnprocessableEntity, HealthRequestStatusStatusCodeTooManyRequests, HealthRequestStatusStatusCodeInsufficientStorage, HealthRequestStatusStatusCodeServiceUnavailable, HealthRequestStatusStatusCodePayloadTooLarge, HealthRequestStatusStatusCodeNotAcceptable, HealthRequestStatusStatusCodeUnavailableForLegalReasons, HealthRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// response containing a list of all metering queries results
type MeteringResult struct {
	Measurements  []MeteringResultMeasurement `json:"measurements"`
	RequestStatus MeteringResultRequestStatus `json:"requestStatus"`
	JSON          meteringResultJSON          `json:"-"`
}

// meteringResultJSON contains the JSON metadata for the struct [MeteringResult]
type meteringResultJSON struct {
	Measurements  apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeteringResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultJSON) RawJSON() string {
	return r.raw
}

// result of a single measurement query
type MeteringResultMeasurement struct {
	// array of series within a measurement
	Series []MeteringResultMeasurementsSery `json:"series"`
	JSON   meteringResultMeasurementJSON    `json:"-"`
}

// meteringResultMeasurementJSON contains the JSON metadata for the struct
// [MeteringResultMeasurement]
type meteringResultMeasurementJSON struct {
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultMeasurement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultMeasurementJSON) RawJSON() string {
	return r.raw
}

// object for a single series in the measurement
type MeteringResultMeasurementsSery struct {
	// list of columns, in order, for the series.
	Columns []string `json:"columns"`
	// name for the measurement
	Name string `json:"name"`
	// list of tags identifying the series.
	Tags []MeteringResultMeasurementsSeriesTag `json:"tags"`
	// array of values, in the same order as the columns, for the series.
	Values []MeteringResultMeasurementsSeriesValue `json:"values"`
	JSON   meteringResultMeasurementsSeryJSON      `json:"-"`
}

// meteringResultMeasurementsSeryJSON contains the JSON metadata for the struct
// [MeteringResultMeasurementsSery]
type meteringResultMeasurementsSeryJSON struct {
	Columns     apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultMeasurementsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultMeasurementsSeryJSON) RawJSON() string {
	return r.raw
}

// object for measurement tags which identifies a measuurement series
type MeteringResultMeasurementsSeriesTag struct {
	// key for the tag, ie)host, job_id, gpu_id
	TagKey string `json:"tagKey"`
	// value for the tag, ie)host=foo, job_id=bar, gpu_id=racecar
	TagValue string                                  `json:"tagValue"`
	JSON     meteringResultMeasurementsSeriesTagJSON `json:"-"`
}

// meteringResultMeasurementsSeriesTagJSON contains the JSON metadata for the
// struct [MeteringResultMeasurementsSeriesTag]
type meteringResultMeasurementsSeriesTagJSON struct {
	TagKey      apijson.Field
	TagValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultMeasurementsSeriesTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultMeasurementsSeriesTagJSON) RawJSON() string {
	return r.raw
}

// object for the measurement values
type MeteringResultMeasurementsSeriesValue struct {
	// list of values, in the same order as columns
	Value []string                                  `json:"value"`
	JSON  meteringResultMeasurementsSeriesValueJSON `json:"-"`
}

// meteringResultMeasurementsSeriesValueJSON contains the JSON metadata for the
// struct [MeteringResultMeasurementsSeriesValue]
type meteringResultMeasurementsSeriesValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultMeasurementsSeriesValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultMeasurementsSeriesValueJSON) RawJSON() string {
	return r.raw
}

type MeteringResultRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        MeteringResultRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                `json:"statusDescription"`
	JSON              meteringResultRequestStatusJSON       `json:"-"`
}

// meteringResultRequestStatusJSON contains the JSON metadata for the struct
// [MeteringResultRequestStatus]
type meteringResultRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MeteringResultRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type MeteringResultRequestStatusStatusCode string

const (
	MeteringResultRequestStatusStatusCodeUnknown                    MeteringResultRequestStatusStatusCode = "UNKNOWN"
	MeteringResultRequestStatusStatusCodeSuccess                    MeteringResultRequestStatusStatusCode = "SUCCESS"
	MeteringResultRequestStatusStatusCodeUnauthorized               MeteringResultRequestStatusStatusCode = "UNAUTHORIZED"
	MeteringResultRequestStatusStatusCodePaymentRequired            MeteringResultRequestStatusStatusCode = "PAYMENT_REQUIRED"
	MeteringResultRequestStatusStatusCodeForbidden                  MeteringResultRequestStatusStatusCode = "FORBIDDEN"
	MeteringResultRequestStatusStatusCodeTimeout                    MeteringResultRequestStatusStatusCode = "TIMEOUT"
	MeteringResultRequestStatusStatusCodeExists                     MeteringResultRequestStatusStatusCode = "EXISTS"
	MeteringResultRequestStatusStatusCodeNotFound                   MeteringResultRequestStatusStatusCode = "NOT_FOUND"
	MeteringResultRequestStatusStatusCodeInternalError              MeteringResultRequestStatusStatusCode = "INTERNAL_ERROR"
	MeteringResultRequestStatusStatusCodeInvalidRequest             MeteringResultRequestStatusStatusCode = "INVALID_REQUEST"
	MeteringResultRequestStatusStatusCodeInvalidRequestVersion      MeteringResultRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	MeteringResultRequestStatusStatusCodeInvalidRequestData         MeteringResultRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	MeteringResultRequestStatusStatusCodeMethodNotAllowed           MeteringResultRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	MeteringResultRequestStatusStatusCodeConflict                   MeteringResultRequestStatusStatusCode = "CONFLICT"
	MeteringResultRequestStatusStatusCodeUnprocessableEntity        MeteringResultRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	MeteringResultRequestStatusStatusCodeTooManyRequests            MeteringResultRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	MeteringResultRequestStatusStatusCodeInsufficientStorage        MeteringResultRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	MeteringResultRequestStatusStatusCodeServiceUnavailable         MeteringResultRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	MeteringResultRequestStatusStatusCodePayloadTooLarge            MeteringResultRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	MeteringResultRequestStatusStatusCodeNotAcceptable              MeteringResultRequestStatusStatusCode = "NOT_ACCEPTABLE"
	MeteringResultRequestStatusStatusCodeUnavailableForLegalReasons MeteringResultRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	MeteringResultRequestStatusStatusCodeBadGateway                 MeteringResultRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r MeteringResultRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case MeteringResultRequestStatusStatusCodeUnknown, MeteringResultRequestStatusStatusCodeSuccess, MeteringResultRequestStatusStatusCodeUnauthorized, MeteringResultRequestStatusStatusCodePaymentRequired, MeteringResultRequestStatusStatusCodeForbidden, MeteringResultRequestStatusStatusCodeTimeout, MeteringResultRequestStatusStatusCodeExists, MeteringResultRequestStatusStatusCodeNotFound, MeteringResultRequestStatusStatusCodeInternalError, MeteringResultRequestStatusStatusCodeInvalidRequest, MeteringResultRequestStatusStatusCodeInvalidRequestVersion, MeteringResultRequestStatusStatusCodeInvalidRequestData, MeteringResultRequestStatusStatusCodeMethodNotAllowed, MeteringResultRequestStatusStatusCodeConflict, MeteringResultRequestStatusStatusCodeUnprocessableEntity, MeteringResultRequestStatusStatusCodeTooManyRequests, MeteringResultRequestStatusStatusCodeInsufficientStorage, MeteringResultRequestStatusStatusCodeServiceUnavailable, MeteringResultRequestStatusStatusCodePayloadTooLarge, MeteringResultRequestStatusStatusCodeNotAcceptable, MeteringResultRequestStatusStatusCodeUnavailableForLegalReasons, MeteringResultRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// listing of all teams
type TeamListResponse struct {
	// object that describes the pagination information
	PaginationInfo TeamListResponsePaginationInfo `json:"paginationInfo"`
	RequestStatus  TeamListResponseRequestStatus  `json:"requestStatus"`
	Teams          []TeamListResponseTeam         `json:"teams"`
	JSON           teamListResponseJSON           `json:"-"`
}

// teamListResponseJSON contains the JSON metadata for the struct
// [TeamListResponse]
type teamListResponseJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Teams          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TeamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type TeamListResponsePaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                              `json:"totalResults"`
	JSON         teamListResponsePaginationInfoJSON `json:"-"`
}

// teamListResponsePaginationInfoJSON contains the JSON metadata for the struct
// [TeamListResponsePaginationInfo]
type teamListResponsePaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamListResponsePaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponsePaginationInfoJSON) RawJSON() string {
	return r.raw
}

type TeamListResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamListResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                  `json:"statusDescription"`
	JSON              teamListResponseRequestStatusJSON       `json:"-"`
}

// teamListResponseRequestStatusJSON contains the JSON metadata for the struct
// [TeamListResponseRequestStatus]
type teamListResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamListResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamListResponseRequestStatusStatusCode string

const (
	TeamListResponseRequestStatusStatusCodeUnknown                    TeamListResponseRequestStatusStatusCode = "UNKNOWN"
	TeamListResponseRequestStatusStatusCodeSuccess                    TeamListResponseRequestStatusStatusCode = "SUCCESS"
	TeamListResponseRequestStatusStatusCodeUnauthorized               TeamListResponseRequestStatusStatusCode = "UNAUTHORIZED"
	TeamListResponseRequestStatusStatusCodePaymentRequired            TeamListResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamListResponseRequestStatusStatusCodeForbidden                  TeamListResponseRequestStatusStatusCode = "FORBIDDEN"
	TeamListResponseRequestStatusStatusCodeTimeout                    TeamListResponseRequestStatusStatusCode = "TIMEOUT"
	TeamListResponseRequestStatusStatusCodeExists                     TeamListResponseRequestStatusStatusCode = "EXISTS"
	TeamListResponseRequestStatusStatusCodeNotFound                   TeamListResponseRequestStatusStatusCode = "NOT_FOUND"
	TeamListResponseRequestStatusStatusCodeInternalError              TeamListResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamListResponseRequestStatusStatusCodeInvalidRequest             TeamListResponseRequestStatusStatusCode = "INVALID_REQUEST"
	TeamListResponseRequestStatusStatusCodeInvalidRequestVersion      TeamListResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamListResponseRequestStatusStatusCodeInvalidRequestData         TeamListResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamListResponseRequestStatusStatusCodeMethodNotAllowed           TeamListResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamListResponseRequestStatusStatusCodeConflict                   TeamListResponseRequestStatusStatusCode = "CONFLICT"
	TeamListResponseRequestStatusStatusCodeUnprocessableEntity        TeamListResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamListResponseRequestStatusStatusCodeTooManyRequests            TeamListResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamListResponseRequestStatusStatusCodeInsufficientStorage        TeamListResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamListResponseRequestStatusStatusCodeServiceUnavailable         TeamListResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamListResponseRequestStatusStatusCodePayloadTooLarge            TeamListResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamListResponseRequestStatusStatusCodeNotAcceptable              TeamListResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamListResponseRequestStatusStatusCodeUnavailableForLegalReasons TeamListResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamListResponseRequestStatusStatusCodeBadGateway                 TeamListResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamListResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamListResponseRequestStatusStatusCodeUnknown, TeamListResponseRequestStatusStatusCodeSuccess, TeamListResponseRequestStatusStatusCodeUnauthorized, TeamListResponseRequestStatusStatusCodePaymentRequired, TeamListResponseRequestStatusStatusCodeForbidden, TeamListResponseRequestStatusStatusCodeTimeout, TeamListResponseRequestStatusStatusCodeExists, TeamListResponseRequestStatusStatusCodeNotFound, TeamListResponseRequestStatusStatusCodeInternalError, TeamListResponseRequestStatusStatusCodeInvalidRequest, TeamListResponseRequestStatusStatusCodeInvalidRequestVersion, TeamListResponseRequestStatusStatusCodeInvalidRequestData, TeamListResponseRequestStatusStatusCodeMethodNotAllowed, TeamListResponseRequestStatusStatusCodeConflict, TeamListResponseRequestStatusStatusCodeUnprocessableEntity, TeamListResponseRequestStatusStatusCodeTooManyRequests, TeamListResponseRequestStatusStatusCodeInsufficientStorage, TeamListResponseRequestStatusStatusCodeServiceUnavailable, TeamListResponseRequestStatusStatusCodePayloadTooLarge, TeamListResponseRequestStatusStatusCodeNotAcceptable, TeamListResponseRequestStatusStatusCodeUnavailableForLegalReasons, TeamListResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type TeamListResponseTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamListResponseTeamsInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamListResponseTeamsRepoScanSettings `json:"repoScanSettings"`
	JSON             teamListResponseTeamJSON              `json:"-"`
}

// teamListResponseTeamJSON contains the JSON metadata for the struct
// [TeamListResponseTeam]
type teamListResponseTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TeamListResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamListResponseTeamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                             `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamListResponseTeamsInfinityManagerSettingsJSON `json:"-"`
}

// teamListResponseTeamsInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [TeamListResponseTeamsInfinityManagerSettings]
type teamListResponseTeamsInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamListResponseTeamsInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseTeamsInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamListResponseTeamsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                      `json:"repoScanShowResults"`
	JSON                teamListResponseTeamsRepoScanSettingsJSON `json:"-"`
}

// teamListResponseTeamsRepoScanSettingsJSON contains the JSON metadata for the
// struct [TeamListResponseTeamsRepoScanSettings]
type teamListResponseTeamsRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamListResponseTeamsRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseTeamsRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// Response for List User reponse
type UserListResponse struct {
	// object that describes the pagination information
	PaginationInfo UserListResponsePaginationInfo `json:"paginationInfo"`
	RequestStatus  UserListResponseRequestStatus  `json:"requestStatus"`
	// information about the user
	Users []UserListResponseUser `json:"users"`
	JSON  userListResponseJSON   `json:"-"`
}

// userListResponseJSON contains the JSON metadata for the struct
// [UserListResponse]
type userListResponseJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Users          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type UserListResponsePaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                              `json:"totalResults"`
	JSON         userListResponsePaginationInfoJSON `json:"-"`
}

// userListResponsePaginationInfoJSON contains the JSON metadata for the struct
// [UserListResponsePaginationInfo]
type userListResponsePaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListResponsePaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponsePaginationInfoJSON) RawJSON() string {
	return r.raw
}

type UserListResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserListResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                  `json:"statusDescription"`
	JSON              userListResponseRequestStatusJSON       `json:"-"`
}

// userListResponseRequestStatusJSON contains the JSON metadata for the struct
// [UserListResponseRequestStatus]
type userListResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserListResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserListResponseRequestStatusStatusCode string

const (
	UserListResponseRequestStatusStatusCodeUnknown                    UserListResponseRequestStatusStatusCode = "UNKNOWN"
	UserListResponseRequestStatusStatusCodeSuccess                    UserListResponseRequestStatusStatusCode = "SUCCESS"
	UserListResponseRequestStatusStatusCodeUnauthorized               UserListResponseRequestStatusStatusCode = "UNAUTHORIZED"
	UserListResponseRequestStatusStatusCodePaymentRequired            UserListResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserListResponseRequestStatusStatusCodeForbidden                  UserListResponseRequestStatusStatusCode = "FORBIDDEN"
	UserListResponseRequestStatusStatusCodeTimeout                    UserListResponseRequestStatusStatusCode = "TIMEOUT"
	UserListResponseRequestStatusStatusCodeExists                     UserListResponseRequestStatusStatusCode = "EXISTS"
	UserListResponseRequestStatusStatusCodeNotFound                   UserListResponseRequestStatusStatusCode = "NOT_FOUND"
	UserListResponseRequestStatusStatusCodeInternalError              UserListResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	UserListResponseRequestStatusStatusCodeInvalidRequest             UserListResponseRequestStatusStatusCode = "INVALID_REQUEST"
	UserListResponseRequestStatusStatusCodeInvalidRequestVersion      UserListResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserListResponseRequestStatusStatusCodeInvalidRequestData         UserListResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserListResponseRequestStatusStatusCodeMethodNotAllowed           UserListResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserListResponseRequestStatusStatusCodeConflict                   UserListResponseRequestStatusStatusCode = "CONFLICT"
	UserListResponseRequestStatusStatusCodeUnprocessableEntity        UserListResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserListResponseRequestStatusStatusCodeTooManyRequests            UserListResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserListResponseRequestStatusStatusCodeInsufficientStorage        UserListResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserListResponseRequestStatusStatusCodeServiceUnavailable         UserListResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserListResponseRequestStatusStatusCodePayloadTooLarge            UserListResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserListResponseRequestStatusStatusCodeNotAcceptable              UserListResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserListResponseRequestStatusStatusCodeUnavailableForLegalReasons UserListResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserListResponseRequestStatusStatusCodeBadGateway                 UserListResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserListResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserListResponseRequestStatusStatusCodeUnknown, UserListResponseRequestStatusStatusCodeSuccess, UserListResponseRequestStatusStatusCodeUnauthorized, UserListResponseRequestStatusStatusCodePaymentRequired, UserListResponseRequestStatusStatusCodeForbidden, UserListResponseRequestStatusStatusCodeTimeout, UserListResponseRequestStatusStatusCodeExists, UserListResponseRequestStatusStatusCodeNotFound, UserListResponseRequestStatusStatusCodeInternalError, UserListResponseRequestStatusStatusCodeInvalidRequest, UserListResponseRequestStatusStatusCodeInvalidRequestVersion, UserListResponseRequestStatusStatusCodeInvalidRequestData, UserListResponseRequestStatusStatusCodeMethodNotAllowed, UserListResponseRequestStatusStatusCodeConflict, UserListResponseRequestStatusStatusCodeUnprocessableEntity, UserListResponseRequestStatusStatusCodeTooManyRequests, UserListResponseRequestStatusStatusCodeInsufficientStorage, UserListResponseRequestStatusStatusCodeServiceUnavailable, UserListResponseRequestStatusStatusCodePayloadTooLarge, UserListResponseRequestStatusStatusCodeNotAcceptable, UserListResponseRequestStatusStatusCodeUnavailableForLegalReasons, UserListResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// information about the user
type UserListResponseUser struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType UserListResponseUsersIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []UserListResponseUsersRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []UserListResponseUsersStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata UserListResponseUsersUserMetadata `json:"userMetadata"`
	JSON         userListResponseUserJSON          `json:"-"`
}

// userListResponseUserJSON contains the JSON metadata for the struct
// [UserListResponseUser]
type userListResponseUserJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *UserListResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUserJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type UserListResponseUsersIdpType string

const (
	UserListResponseUsersIdpTypeNvidia     UserListResponseUsersIdpType = "NVIDIA"
	UserListResponseUsersIdpTypeEnterprise UserListResponseUsersIdpType = "ENTERPRISE"
)

func (r UserListResponseUsersIdpType) IsKnown() bool {
	switch r {
	case UserListResponseUsersIdpTypeNvidia, UserListResponseUsersIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type UserListResponseUsersRole struct {
	// Information about the Organization
	Org UserListResponseUsersRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier UserListResponseUsersRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team UserListResponseUsersRolesTeam `json:"team"`
	// List of team role types that the user have
	TeamRoles []string                      `json:"teamRoles"`
	JSON      userListResponseUsersRoleJSON `json:"-"`
}

// userListResponseUsersRoleJSON contains the JSON metadata for the struct
// [UserListResponseUsersRole]
type userListResponseUsersRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserListResponseUsersRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type UserListResponseUsersRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact UserListResponseUsersRolesOrgAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings UserListResponseUsersRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner UserListResponseUsersRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []UserListResponseUsersRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                             `json:"pecSfdcId"`
	ProductEnablements   []UserListResponseUsersRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []UserListResponseUsersRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings UserListResponseUsersRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             UserListResponseUsersRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo UserListResponseUsersRolesOrgUsersInfo `json:"usersInfo"`
	JSON      userListResponseUsersRolesOrgJSON      `json:"-"`
}

// userListResponseUsersRolesOrgJSON contains the JSON metadata for the struct
// [UserListResponseUsersRolesOrg]
type userListResponseUsersRolesOrgJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type UserListResponseUsersRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                            `json:"fullName"`
	JSON     userListResponseUsersRolesOrgAlternateContactJSON `json:"-"`
}

// userListResponseUsersRolesOrgAlternateContactJSON contains the JSON metadata for
// the struct [UserListResponseUsersRolesOrgAlternateContact]
type userListResponseUsersRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserListResponseUsersRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                     `json:"infinityManagerEnableTeamOverride"`
	JSON                              userListResponseUsersRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// userListResponseUsersRolesOrgInfinityManagerSettingsJSON contains the JSON
// metadata for the struct [UserListResponseUsersRolesOrgInfinityManagerSettings]
type userListResponseUsersRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type UserListResponseUsersRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                                    `json:"lastLoginDate"`
	JSON          userListResponseUsersRolesOrgOrgOwnerJSON `json:"-"`
}

// userListResponseUsersRolesOrgOrgOwnerJSON contains the JSON metadata for the
// struct [UserListResponseUsersRolesOrgOrgOwner]
type userListResponseUsersRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type UserListResponseUsersRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type UserListResponseUsersRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                                    `json:"expirationDate"`
	PoDetails      []UserListResponseUsersRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           userListResponseUsersRolesOrgProductEnablementJSON        `json:"-"`
}

// userListResponseUsersRolesOrgProductEnablementJSON contains the JSON metadata
// for the struct [UserListResponseUsersRolesOrgProductEnablement]
type userListResponseUsersRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type UserListResponseUsersRolesOrgProductEnablementsType string

const (
	UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminEval       UserListResponseUsersRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminNfr        UserListResponseUsersRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminCommercial UserListResponseUsersRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	UserListResponseUsersRolesOrgProductEnablementsTypeEmsEval            UserListResponseUsersRolesOrgProductEnablementsType = "EMS_EVAL"
	UserListResponseUsersRolesOrgProductEnablementsTypeEmsNfr             UserListResponseUsersRolesOrgProductEnablementsType = "EMS_NFR"
	UserListResponseUsersRolesOrgProductEnablementsTypeEmsCommercial      UserListResponseUsersRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper  UserListResponseUsersRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r UserListResponseUsersRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminEval, UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminNfr, UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminCommercial, UserListResponseUsersRolesOrgProductEnablementsTypeEmsEval, UserListResponseUsersRolesOrgProductEnablementsTypeEmsNfr, UserListResponseUsersRolesOrgProductEnablementsTypeEmsCommercial, UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type UserListResponseUsersRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                                      `json:"pkId"`
	JSON userListResponseUsersRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// userListResponseUsersRolesOrgProductEnablementsPoDetailJSON contains the JSON
// metadata for the struct
// [UserListResponseUsersRolesOrgProductEnablementsPoDetail]
type userListResponseUsersRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type UserListResponseUsersRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type UserListResponseUsersRolesOrgProductSubscriptionsType `json:"type"`
	JSON userListResponseUsersRolesOrgProductSubscriptionJSON  `json:"-"`
}

// userListResponseUsersRolesOrgProductSubscriptionJSON contains the JSON metadata
// for the struct [UserListResponseUsersRolesOrgProductSubscription]
type userListResponseUsersRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type UserListResponseUsersRolesOrgProductSubscriptionsType string

const (
	UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminEval       UserListResponseUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr        UserListResponseUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial UserListResponseUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r UserListResponseUsersRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminEval, UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr, UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type UserListResponseUsersRolesOrgRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                              `json:"repoScanShowResults"`
	JSON                userListResponseUsersRolesOrgRepoScanSettingsJSON `json:"-"`
}

// userListResponseUsersRolesOrgRepoScanSettingsJSON contains the JSON metadata for
// the struct [UserListResponseUsersRolesOrgRepoScanSettings]
type userListResponseUsersRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type UserListResponseUsersRolesOrgType string

const (
	UserListResponseUsersRolesOrgTypeUnknown    UserListResponseUsersRolesOrgType = "UNKNOWN"
	UserListResponseUsersRolesOrgTypeCloud      UserListResponseUsersRolesOrgType = "CLOUD"
	UserListResponseUsersRolesOrgTypeEnterprise UserListResponseUsersRolesOrgType = "ENTERPRISE"
	UserListResponseUsersRolesOrgTypeIndividual UserListResponseUsersRolesOrgType = "INDIVIDUAL"
)

func (r UserListResponseUsersRolesOrgType) IsKnown() bool {
	switch r {
	case UserListResponseUsersRolesOrgTypeUnknown, UserListResponseUsersRolesOrgTypeCloud, UserListResponseUsersRolesOrgTypeEnterprise, UserListResponseUsersRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type UserListResponseUsersRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                                      `json:"totalUsers"`
	JSON       userListResponseUsersRolesOrgUsersInfoJSON `json:"-"`
}

// userListResponseUsersRolesOrgUsersInfoJSON contains the JSON metadata for the
// struct [UserListResponseUsersRolesOrgUsersInfo]
type userListResponseUsersRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseUsersRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type UserListResponseUsersRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                                    `json:"userId"`
	JSON   userListResponseUsersRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// userListResponseUsersRolesTargetSystemUserIdentifierJSON contains the JSON
// metadata for the struct [UserListResponseUsersRolesTargetSystemUserIdentifier]
type userListResponseUsersRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseUsersRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// Information about the team
type UserListResponseUsersRolesTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings UserListResponseUsersRolesTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings UserListResponseUsersRolesTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             userListResponseUsersRolesTeamJSON             `json:"-"`
}

// userListResponseUsersRolesTeamJSON contains the JSON metadata for the struct
// [UserListResponseUsersRolesTeam]
type userListResponseUsersRolesTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserListResponseUsersRolesTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserListResponseUsersRolesTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                      `json:"infinityManagerEnableTeamOverride"`
	JSON                              userListResponseUsersRolesTeamInfinityManagerSettingsJSON `json:"-"`
}

// userListResponseUsersRolesTeamInfinityManagerSettingsJSON contains the JSON
// metadata for the struct [UserListResponseUsersRolesTeamInfinityManagerSettings]
type userListResponseUsersRolesTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserListResponseUsersRolesTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type UserListResponseUsersRolesTeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                               `json:"repoScanShowResults"`
	JSON                userListResponseUsersRolesTeamRepoScanSettingsJSON `json:"-"`
}

// userListResponseUsersRolesTeamRepoScanSettingsJSON contains the JSON metadata
// for the struct [UserListResponseUsersRolesTeamRepoScanSettings]
type userListResponseUsersRolesTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserListResponseUsersRolesTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersRolesTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type UserListResponseUsersStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                                 `json:"workspacesUsage"`
	JSON            userListResponseUsersStorageQuotaJSON `json:"-"`
}

// userListResponseUsersStorageQuotaJSON contains the JSON metadata for the struct
// [UserListResponseUsersStorageQuota]
type userListResponseUsersStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserListResponseUsersStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type UserListResponseUsersUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                                `json:"role"`
	JSON userListResponseUsersUserMetadataJSON `json:"-"`
}

// userListResponseUsersUserMetadataJSON contains the JSON metadata for the struct
// [UserListResponseUsersUserMetadata]
type userListResponseUsersUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseUsersUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseUsersUserMetadataJSON) RawJSON() string {
	return r.raw
}

// about one user
type UserResponse struct {
	// token needed to activate the user to enable login and other features
	ActivationToken string `json:"activationToken"`
	// NCA role
	NcaRole       UserResponseNcaRole       `json:"ncaRole"`
	RequestStatus UserResponseRequestStatus `json:"requestStatus"`
	// information about the user
	User UserResponseUser `json:"user"`
	// DEPRECATED - Please use roles inside user
	UserRoles []UserResponseUserRole `json:"userRoles"`
	JSON      userResponseJSON       `json:"-"`
}

// userResponseJSON contains the JSON metadata for the struct [UserResponse]
type userResponseJSON struct {
	ActivationToken apijson.Field
	NcaRole         apijson.Field
	RequestStatus   apijson.Field
	User            apijson.Field
	UserRoles       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseJSON) RawJSON() string {
	return r.raw
}

// NCA role
type UserResponseNcaRole string

const (
	UserResponseNcaRoleUnknown       UserResponseNcaRole = "UNKNOWN"
	UserResponseNcaRoleAdministrator UserResponseNcaRole = "ADMINISTRATOR"
	UserResponseNcaRoleMember        UserResponseNcaRole = "MEMBER"
	UserResponseNcaRoleOwner         UserResponseNcaRole = "OWNER"
	UserResponseNcaRolePending       UserResponseNcaRole = "PENDING"
)

func (r UserResponseNcaRole) IsKnown() bool {
	switch r {
	case UserResponseNcaRoleUnknown, UserResponseNcaRoleAdministrator, UserResponseNcaRoleMember, UserResponseNcaRoleOwner, UserResponseNcaRolePending:
		return true
	}
	return false
}

type UserResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                              `json:"statusDescription"`
	JSON              userResponseRequestStatusJSON       `json:"-"`
}

// userResponseRequestStatusJSON contains the JSON metadata for the struct
// [UserResponseRequestStatus]
type userResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserResponseRequestStatusStatusCode string

const (
	UserResponseRequestStatusStatusCodeUnknown                    UserResponseRequestStatusStatusCode = "UNKNOWN"
	UserResponseRequestStatusStatusCodeSuccess                    UserResponseRequestStatusStatusCode = "SUCCESS"
	UserResponseRequestStatusStatusCodeUnauthorized               UserResponseRequestStatusStatusCode = "UNAUTHORIZED"
	UserResponseRequestStatusStatusCodePaymentRequired            UserResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserResponseRequestStatusStatusCodeForbidden                  UserResponseRequestStatusStatusCode = "FORBIDDEN"
	UserResponseRequestStatusStatusCodeTimeout                    UserResponseRequestStatusStatusCode = "TIMEOUT"
	UserResponseRequestStatusStatusCodeExists                     UserResponseRequestStatusStatusCode = "EXISTS"
	UserResponseRequestStatusStatusCodeNotFound                   UserResponseRequestStatusStatusCode = "NOT_FOUND"
	UserResponseRequestStatusStatusCodeInternalError              UserResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	UserResponseRequestStatusStatusCodeInvalidRequest             UserResponseRequestStatusStatusCode = "INVALID_REQUEST"
	UserResponseRequestStatusStatusCodeInvalidRequestVersion      UserResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserResponseRequestStatusStatusCodeInvalidRequestData         UserResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserResponseRequestStatusStatusCodeMethodNotAllowed           UserResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserResponseRequestStatusStatusCodeConflict                   UserResponseRequestStatusStatusCode = "CONFLICT"
	UserResponseRequestStatusStatusCodeUnprocessableEntity        UserResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserResponseRequestStatusStatusCodeTooManyRequests            UserResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserResponseRequestStatusStatusCodeInsufficientStorage        UserResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserResponseRequestStatusStatusCodeServiceUnavailable         UserResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserResponseRequestStatusStatusCodePayloadTooLarge            UserResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserResponseRequestStatusStatusCodeNotAcceptable              UserResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserResponseRequestStatusStatusCodeUnavailableForLegalReasons UserResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserResponseRequestStatusStatusCodeBadGateway                 UserResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserResponseRequestStatusStatusCodeUnknown, UserResponseRequestStatusStatusCodeSuccess, UserResponseRequestStatusStatusCodeUnauthorized, UserResponseRequestStatusStatusCodePaymentRequired, UserResponseRequestStatusStatusCodeForbidden, UserResponseRequestStatusStatusCodeTimeout, UserResponseRequestStatusStatusCodeExists, UserResponseRequestStatusStatusCodeNotFound, UserResponseRequestStatusStatusCodeInternalError, UserResponseRequestStatusStatusCodeInvalidRequest, UserResponseRequestStatusStatusCodeInvalidRequestVersion, UserResponseRequestStatusStatusCodeInvalidRequestData, UserResponseRequestStatusStatusCodeMethodNotAllowed, UserResponseRequestStatusStatusCodeConflict, UserResponseRequestStatusStatusCodeUnprocessableEntity, UserResponseRequestStatusStatusCodeTooManyRequests, UserResponseRequestStatusStatusCodeInsufficientStorage, UserResponseRequestStatusStatusCodeServiceUnavailable, UserResponseRequestStatusStatusCodePayloadTooLarge, UserResponseRequestStatusStatusCodeNotAcceptable, UserResponseRequestStatusStatusCodeUnavailableForLegalReasons, UserResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// information about the user
type UserResponseUser struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType UserResponseUserIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []UserResponseUserRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []UserResponseUserStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata UserResponseUserUserMetadata `json:"userMetadata"`
	JSON         userResponseUserJSON         `json:"-"`
}

// userResponseUserJSON contains the JSON metadata for the struct
// [UserResponseUser]
type userResponseUserJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *UserResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type UserResponseUserIdpType string

const (
	UserResponseUserIdpTypeNvidia     UserResponseUserIdpType = "NVIDIA"
	UserResponseUserIdpTypeEnterprise UserResponseUserIdpType = "ENTERPRISE"
)

func (r UserResponseUserIdpType) IsKnown() bool {
	switch r {
	case UserResponseUserIdpTypeNvidia, UserResponseUserIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type UserResponseUserRole struct {
	// Information about the Organization
	Org UserResponseUserRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier UserResponseUserRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team UserResponseUserRolesTeam `json:"team"`
	// List of team role types that the user have
	TeamRoles []string                 `json:"teamRoles"`
	JSON      userResponseUserRoleJSON `json:"-"`
}

// userResponseUserRoleJSON contains the JSON metadata for the struct
// [UserResponseUserRole]
type userResponseUserRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserResponseUserRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type UserResponseUserRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact UserResponseUserRolesOrgAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings UserResponseUserRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner UserResponseUserRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []UserResponseUserRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                        `json:"pecSfdcId"`
	ProductEnablements   []UserResponseUserRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []UserResponseUserRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings UserResponseUserRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             UserResponseUserRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo UserResponseUserRolesOrgUsersInfo `json:"usersInfo"`
	JSON      userResponseUserRolesOrgJSON      `json:"-"`
}

// userResponseUserRolesOrgJSON contains the JSON metadata for the struct
// [UserResponseUserRolesOrg]
type userResponseUserRolesOrgJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *UserResponseUserRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type UserResponseUserRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                       `json:"fullName"`
	JSON     userResponseUserRolesOrgAlternateContactJSON `json:"-"`
}

// userResponseUserRolesOrgAlternateContactJSON contains the JSON metadata for the
// struct [UserResponseUserRolesOrgAlternateContact]
type userResponseUserRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserResponseUserRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                `json:"infinityManagerEnableTeamOverride"`
	JSON                              userResponseUserRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// userResponseUserRolesOrgInfinityManagerSettingsJSON contains the JSON metadata
// for the struct [UserResponseUserRolesOrgInfinityManagerSettings]
type userResponseUserRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type UserResponseUserRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                               `json:"lastLoginDate"`
	JSON          userResponseUserRolesOrgOrgOwnerJSON `json:"-"`
}

// userResponseUserRolesOrgOrgOwnerJSON contains the JSON metadata for the struct
// [UserResponseUserRolesOrgOrgOwner]
type userResponseUserRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type UserResponseUserRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type UserResponseUserRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                               `json:"expirationDate"`
	PoDetails      []UserResponseUserRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           userResponseUserRolesOrgProductEnablementJSON        `json:"-"`
}

// userResponseUserRolesOrgProductEnablementJSON contains the JSON metadata for the
// struct [UserResponseUserRolesOrgProductEnablement]
type userResponseUserRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type UserResponseUserRolesOrgProductEnablementsType string

const (
	UserResponseUserRolesOrgProductEnablementsTypeNgcAdminEval       UserResponseUserRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	UserResponseUserRolesOrgProductEnablementsTypeNgcAdminNfr        UserResponseUserRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	UserResponseUserRolesOrgProductEnablementsTypeNgcAdminCommercial UserResponseUserRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	UserResponseUserRolesOrgProductEnablementsTypeEmsEval            UserResponseUserRolesOrgProductEnablementsType = "EMS_EVAL"
	UserResponseUserRolesOrgProductEnablementsTypeEmsNfr             UserResponseUserRolesOrgProductEnablementsType = "EMS_NFR"
	UserResponseUserRolesOrgProductEnablementsTypeEmsCommercial      UserResponseUserRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	UserResponseUserRolesOrgProductEnablementsTypeNgcAdminDeveloper  UserResponseUserRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r UserResponseUserRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case UserResponseUserRolesOrgProductEnablementsTypeNgcAdminEval, UserResponseUserRolesOrgProductEnablementsTypeNgcAdminNfr, UserResponseUserRolesOrgProductEnablementsTypeNgcAdminCommercial, UserResponseUserRolesOrgProductEnablementsTypeEmsEval, UserResponseUserRolesOrgProductEnablementsTypeEmsNfr, UserResponseUserRolesOrgProductEnablementsTypeEmsCommercial, UserResponseUserRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type UserResponseUserRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                                 `json:"pkId"`
	JSON userResponseUserRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// userResponseUserRolesOrgProductEnablementsPoDetailJSON contains the JSON
// metadata for the struct [UserResponseUserRolesOrgProductEnablementsPoDetail]
type userResponseUserRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type UserResponseUserRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type UserResponseUserRolesOrgProductSubscriptionsType `json:"type"`
	JSON userResponseUserRolesOrgProductSubscriptionJSON  `json:"-"`
}

// userResponseUserRolesOrgProductSubscriptionJSON contains the JSON metadata for
// the struct [UserResponseUserRolesOrgProductSubscription]
type userResponseUserRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type UserResponseUserRolesOrgProductSubscriptionsType string

const (
	UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminEval       UserResponseUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminNfr        UserResponseUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial UserResponseUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r UserResponseUserRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminEval, UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminNfr, UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type UserResponseUserRolesOrgRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                         `json:"repoScanShowResults"`
	JSON                userResponseUserRolesOrgRepoScanSettingsJSON `json:"-"`
}

// userResponseUserRolesOrgRepoScanSettingsJSON contains the JSON metadata for the
// struct [UserResponseUserRolesOrgRepoScanSettings]
type userResponseUserRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type UserResponseUserRolesOrgType string

const (
	UserResponseUserRolesOrgTypeUnknown    UserResponseUserRolesOrgType = "UNKNOWN"
	UserResponseUserRolesOrgTypeCloud      UserResponseUserRolesOrgType = "CLOUD"
	UserResponseUserRolesOrgTypeEnterprise UserResponseUserRolesOrgType = "ENTERPRISE"
	UserResponseUserRolesOrgTypeIndividual UserResponseUserRolesOrgType = "INDIVIDUAL"
)

func (r UserResponseUserRolesOrgType) IsKnown() bool {
	switch r {
	case UserResponseUserRolesOrgTypeUnknown, UserResponseUserRolesOrgTypeCloud, UserResponseUserRolesOrgTypeEnterprise, UserResponseUserRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type UserResponseUserRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                                 `json:"totalUsers"`
	JSON       userResponseUserRolesOrgUsersInfoJSON `json:"-"`
}

// userResponseUserRolesOrgUsersInfoJSON contains the JSON metadata for the struct
// [UserResponseUserRolesOrgUsersInfo]
type userResponseUserRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserResponseUserRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type UserResponseUserRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                               `json:"userId"`
	JSON   userResponseUserRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// userResponseUserRolesTargetSystemUserIdentifierJSON contains the JSON metadata
// for the struct [UserResponseUserRolesTargetSystemUserIdentifier]
type userResponseUserRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserResponseUserRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// Information about the team
type UserResponseUserRolesTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings UserResponseUserRolesTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings UserResponseUserRolesTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             userResponseUserRolesTeamJSON             `json:"-"`
}

// userResponseUserRolesTeamJSON contains the JSON metadata for the struct
// [UserResponseUserRolesTeam]
type userResponseUserRolesTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserResponseUserRolesTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserResponseUserRolesTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                 `json:"infinityManagerEnableTeamOverride"`
	JSON                              userResponseUserRolesTeamInfinityManagerSettingsJSON `json:"-"`
}

// userResponseUserRolesTeamInfinityManagerSettingsJSON contains the JSON metadata
// for the struct [UserResponseUserRolesTeamInfinityManagerSettings]
type userResponseUserRolesTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserResponseUserRolesTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type UserResponseUserRolesTeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                          `json:"repoScanShowResults"`
	JSON                userResponseUserRolesTeamRepoScanSettingsJSON `json:"-"`
}

// userResponseUserRolesTeamRepoScanSettingsJSON contains the JSON metadata for the
// struct [UserResponseUserRolesTeamRepoScanSettings]
type userResponseUserRolesTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserResponseUserRolesTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserRolesTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type UserResponseUserStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                            `json:"workspacesUsage"`
	JSON            userResponseUserStorageQuotaJSON `json:"-"`
}

// userResponseUserStorageQuotaJSON contains the JSON metadata for the struct
// [UserResponseUserStorageQuota]
type userResponseUserStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserResponseUserStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type UserResponseUserUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                           `json:"role"`
	JSON userResponseUserUserMetadataJSON `json:"-"`
}

// userResponseUserUserMetadataJSON contains the JSON metadata for the struct
// [UserResponseUserUserMetadata]
type userResponseUserUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserResponseUserUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseUserUserMetadataJSON) RawJSON() string {
	return r.raw
}
