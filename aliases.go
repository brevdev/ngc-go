// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/stainless-sdks/ngc-go/internal/apierror"
	"github.com/stainless-sdks/ngc-go/shared"
)

type Error = apierror.Error

// This API is invoked by monitoring tools, other services and infrastructure to
// retrieve health status the targeted service, this is unprotected method
//
// This is an alias to an internal type.
type Health = shared.Health

// object that describes health of the service
//
// This is an alias to an internal type.
type HealthHealth = shared.HealthHealth

// Enum that describes health of the service
//
// This is an alias to an internal type.
type HealthHealthHealthCode = shared.HealthHealthHealthCode

// This is an alias to an internal value.
const HealthHealthHealthCodeUnknown = shared.HealthHealthHealthCodeUnknown

// This is an alias to an internal value.
const HealthHealthHealthCodeOk = shared.HealthHealthHealthCodeOk

// This is an alias to an internal value.
const HealthHealthHealthCodeUnderMaintenance = shared.HealthHealthHealthCodeUnderMaintenance

// This is an alias to an internal value.
const HealthHealthHealthCodeWarning = shared.HealthHealthHealthCodeWarning

// This is an alias to an internal value.
const HealthHealthHealthCodeFailed = shared.HealthHealthHealthCodeFailed

// This is an alias to an internal type.
type HealthHealthMetaData = shared.HealthHealthMetaData

// This is an alias to an internal type.
type HealthRequestStatus = shared.HealthRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type HealthRequestStatusStatusCode = shared.HealthRequestStatusStatusCode

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnknown = shared.HealthRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeSuccess = shared.HealthRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnauthorized = shared.HealthRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const HealthRequestStatusStatusCodePaymentRequired = shared.HealthRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeForbidden = shared.HealthRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeTimeout = shared.HealthRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeExists = shared.HealthRequestStatusStatusCodeExists

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeNotFound = shared.HealthRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInternalError = shared.HealthRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequest = shared.HealthRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequestVersion = shared.HealthRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequestData = shared.HealthRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeMethodNotAllowed = shared.HealthRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeConflict = shared.HealthRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnprocessableEntity = shared.HealthRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeTooManyRequests = shared.HealthRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInsufficientStorage = shared.HealthRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeServiceUnavailable = shared.HealthRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const HealthRequestStatusStatusCodePayloadTooLarge = shared.HealthRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeNotAcceptable = shared.HealthRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnavailableForLegalReasons = shared.HealthRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeBadGateway = shared.HealthRequestStatusStatusCodeBadGateway

// response containing a list of all metering queries results
//
// This is an alias to an internal type.
type MeteringResult = shared.MeteringResult

// result of a single measurement query
//
// This is an alias to an internal type.
type MeteringResultMeasurement = shared.MeteringResultMeasurement

// object for a single series in the measurement
//
// This is an alias to an internal type.
type MeteringResultMeasurementsSery = shared.MeteringResultMeasurementsSery

// object for measurement tags which identifies a measuurement series
//
// This is an alias to an internal type.
type MeteringResultMeasurementsSeriesTag = shared.MeteringResultMeasurementsSeriesTag

// object for the measurement values
//
// This is an alias to an internal type.
type MeteringResultMeasurementsSeriesValue = shared.MeteringResultMeasurementsSeriesValue

// This is an alias to an internal type.
type MeteringResultRequestStatus = shared.MeteringResultRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type MeteringResultRequestStatusStatusCode = shared.MeteringResultRequestStatusStatusCode

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeUnknown = shared.MeteringResultRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeSuccess = shared.MeteringResultRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeUnauthorized = shared.MeteringResultRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodePaymentRequired = shared.MeteringResultRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeForbidden = shared.MeteringResultRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeTimeout = shared.MeteringResultRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeExists = shared.MeteringResultRequestStatusStatusCodeExists

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeNotFound = shared.MeteringResultRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeInternalError = shared.MeteringResultRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeInvalidRequest = shared.MeteringResultRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeInvalidRequestVersion = shared.MeteringResultRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeInvalidRequestData = shared.MeteringResultRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeMethodNotAllowed = shared.MeteringResultRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeConflict = shared.MeteringResultRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeUnprocessableEntity = shared.MeteringResultRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeTooManyRequests = shared.MeteringResultRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeInsufficientStorage = shared.MeteringResultRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeServiceUnavailable = shared.MeteringResultRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodePayloadTooLarge = shared.MeteringResultRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeNotAcceptable = shared.MeteringResultRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeUnavailableForLegalReasons = shared.MeteringResultRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const MeteringResultRequestStatusStatusCodeBadGateway = shared.MeteringResultRequestStatusStatusCodeBadGateway

// listing of all teams
//
// This is an alias to an internal type.
type TeamListResponse = shared.TeamListResponse

// object that describes the pagination information
//
// This is an alias to an internal type.
type TeamListResponsePaginationInfo = shared.TeamListResponsePaginationInfo

// This is an alias to an internal type.
type TeamListResponseRequestStatus = shared.TeamListResponseRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type TeamListResponseRequestStatusStatusCode = shared.TeamListResponseRequestStatusStatusCode

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeUnknown = shared.TeamListResponseRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeSuccess = shared.TeamListResponseRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeUnauthorized = shared.TeamListResponseRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodePaymentRequired = shared.TeamListResponseRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeForbidden = shared.TeamListResponseRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeTimeout = shared.TeamListResponseRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeExists = shared.TeamListResponseRequestStatusStatusCodeExists

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeNotFound = shared.TeamListResponseRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeInternalError = shared.TeamListResponseRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeInvalidRequest = shared.TeamListResponseRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeInvalidRequestVersion = shared.TeamListResponseRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeInvalidRequestData = shared.TeamListResponseRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeMethodNotAllowed = shared.TeamListResponseRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeConflict = shared.TeamListResponseRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeUnprocessableEntity = shared.TeamListResponseRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeTooManyRequests = shared.TeamListResponseRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeInsufficientStorage = shared.TeamListResponseRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeServiceUnavailable = shared.TeamListResponseRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodePayloadTooLarge = shared.TeamListResponseRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeNotAcceptable = shared.TeamListResponseRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeUnavailableForLegalReasons = shared.TeamListResponseRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const TeamListResponseRequestStatusStatusCodeBadGateway = shared.TeamListResponseRequestStatusStatusCodeBadGateway

// Information about the team
//
// This is an alias to an internal type.
type TeamListResponseTeam = shared.TeamListResponseTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type TeamListResponseTeamsInfinityManagerSettings = shared.TeamListResponseTeamsInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type TeamListResponseTeamsRepoScanSettings = shared.TeamListResponseTeamsRepoScanSettings

// Response for List User reponse
//
// This is an alias to an internal type.
type UserListResponse = shared.UserListResponse

// object that describes the pagination information
//
// This is an alias to an internal type.
type UserListResponsePaginationInfo = shared.UserListResponsePaginationInfo

// This is an alias to an internal type.
type UserListResponseRequestStatus = shared.UserListResponseRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserListResponseRequestStatusStatusCode = shared.UserListResponseRequestStatusStatusCode

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeUnknown = shared.UserListResponseRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeSuccess = shared.UserListResponseRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeUnauthorized = shared.UserListResponseRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodePaymentRequired = shared.UserListResponseRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeForbidden = shared.UserListResponseRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeTimeout = shared.UserListResponseRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeExists = shared.UserListResponseRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeNotFound = shared.UserListResponseRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeInternalError = shared.UserListResponseRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeInvalidRequest = shared.UserListResponseRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeInvalidRequestVersion = shared.UserListResponseRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeInvalidRequestData = shared.UserListResponseRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeMethodNotAllowed = shared.UserListResponseRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeConflict = shared.UserListResponseRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeUnprocessableEntity = shared.UserListResponseRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeTooManyRequests = shared.UserListResponseRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeInsufficientStorage = shared.UserListResponseRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeServiceUnavailable = shared.UserListResponseRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodePayloadTooLarge = shared.UserListResponseRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeNotAcceptable = shared.UserListResponseRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserListResponseRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserListResponseRequestStatusStatusCodeBadGateway = shared.UserListResponseRequestStatusStatusCodeBadGateway

// information about the user
//
// This is an alias to an internal type.
type UserListResponseUser = shared.UserListResponseUser

// Type of IDP, Identity Provider. Used for login.
//
// This is an alias to an internal type.
type UserListResponseUsersIdpType = shared.UserListResponseUsersIdpType

// This is an alias to an internal value.
const UserListResponseUsersIdpTypeNvidia = shared.UserListResponseUsersIdpTypeNvidia

// This is an alias to an internal value.
const UserListResponseUsersIdpTypeEnterprise = shared.UserListResponseUsersIdpTypeEnterprise

// List of roles that the user have
//
// This is an alias to an internal type.
type UserListResponseUsersRole = shared.UserListResponseUsersRole

// Information about the Organization
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrg = shared.UserListResponseUsersRolesOrg

// Org Owner Alternate Contact
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgAlternateContact = shared.UserListResponseUsersRolesOrgAlternateContact

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgInfinityManagerSettings = shared.UserListResponseUsersRolesOrgInfinityManagerSettings

// Org owner.
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgOrgOwner = shared.UserListResponseUsersRolesOrgOrgOwner

// Product Enablement
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductEnablement = shared.UserListResponseUsersRolesOrgProductEnablement

// Product Enablement Types
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductEnablementsType = shared.UserListResponseUsersRolesOrgProductEnablementsType

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminEval = shared.UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminEval

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminNfr = shared.UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminCommercial = shared.UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminCommercial

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeEmsEval = shared.UserListResponseUsersRolesOrgProductEnablementsTypeEmsEval

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeEmsNfr = shared.UserListResponseUsersRolesOrgProductEnablementsTypeEmsNfr

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeEmsCommercial = shared.UserListResponseUsersRolesOrgProductEnablementsTypeEmsCommercial

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper = shared.UserListResponseUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper

// Purchase Order.
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductEnablementsPoDetail = shared.UserListResponseUsersRolesOrgProductEnablementsPoDetail

// Product Subscription
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductSubscription = shared.UserListResponseUsersRolesOrgProductSubscription

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType = shared.UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementType

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval = shared.UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr = shared.UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical = shared.UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial = shared.UserListResponseUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgProductSubscriptionsType = shared.UserListResponseUsersRolesOrgProductSubscriptionsType

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminEval = shared.UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminEval

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr = shared.UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial = shared.UserListResponseUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgRepoScanSettings = shared.UserListResponseUsersRolesOrgRepoScanSettings

// This is an alias to an internal type.
type UserListResponseUsersRolesOrgType = shared.UserListResponseUsersRolesOrgType

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgTypeUnknown = shared.UserListResponseUsersRolesOrgTypeUnknown

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgTypeCloud = shared.UserListResponseUsersRolesOrgTypeCloud

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgTypeEnterprise = shared.UserListResponseUsersRolesOrgTypeEnterprise

// This is an alias to an internal value.
const UserListResponseUsersRolesOrgTypeIndividual = shared.UserListResponseUsersRolesOrgTypeIndividual

// Users information.
//
// This is an alias to an internal type.
type UserListResponseUsersRolesOrgUsersInfo = shared.UserListResponseUsersRolesOrgUsersInfo

// Information about the user who is attempting to run the job
//
// This is an alias to an internal type.
type UserListResponseUsersRolesTargetSystemUserIdentifier = shared.UserListResponseUsersRolesTargetSystemUserIdentifier

// Information about the team
//
// This is an alias to an internal type.
type UserListResponseUsersRolesTeam = shared.UserListResponseUsersRolesTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserListResponseUsersRolesTeamInfinityManagerSettings = shared.UserListResponseUsersRolesTeamInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserListResponseUsersRolesTeamRepoScanSettings = shared.UserListResponseUsersRolesTeamRepoScanSettings

// represents user storage quota for a given ace and available unused storage
//
// This is an alias to an internal type.
type UserListResponseUsersStorageQuota = shared.UserListResponseUsersStorageQuota

// Metadata information about the user.
//
// This is an alias to an internal type.
type UserListResponseUsersUserMetadata = shared.UserListResponseUsersUserMetadata

// about one user
//
// This is an alias to an internal type.
type UserResponse = shared.UserResponse

// NCA role
//
// This is an alias to an internal type.
type UserResponseNcaRole = shared.UserResponseNcaRole

// This is an alias to an internal value.
const UserResponseNcaRoleUnknown = shared.UserResponseNcaRoleUnknown

// This is an alias to an internal value.
const UserResponseNcaRoleAdministrator = shared.UserResponseNcaRoleAdministrator

// This is an alias to an internal value.
const UserResponseNcaRoleMember = shared.UserResponseNcaRoleMember

// This is an alias to an internal value.
const UserResponseNcaRoleOwner = shared.UserResponseNcaRoleOwner

// This is an alias to an internal value.
const UserResponseNcaRolePending = shared.UserResponseNcaRolePending

// This is an alias to an internal type.
type UserResponseRequestStatus = shared.UserResponseRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserResponseRequestStatusStatusCode = shared.UserResponseRequestStatusStatusCode

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeUnknown = shared.UserResponseRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeSuccess = shared.UserResponseRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeUnauthorized = shared.UserResponseRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodePaymentRequired = shared.UserResponseRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeForbidden = shared.UserResponseRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeTimeout = shared.UserResponseRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeExists = shared.UserResponseRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeNotFound = shared.UserResponseRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeInternalError = shared.UserResponseRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeInvalidRequest = shared.UserResponseRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeInvalidRequestVersion = shared.UserResponseRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeInvalidRequestData = shared.UserResponseRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeMethodNotAllowed = shared.UserResponseRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeConflict = shared.UserResponseRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeUnprocessableEntity = shared.UserResponseRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeTooManyRequests = shared.UserResponseRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeInsufficientStorage = shared.UserResponseRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeServiceUnavailable = shared.UserResponseRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodePayloadTooLarge = shared.UserResponseRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeNotAcceptable = shared.UserResponseRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserResponseRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserResponseRequestStatusStatusCodeBadGateway = shared.UserResponseRequestStatusStatusCodeBadGateway

// information about the user
//
// This is an alias to an internal type.
type UserResponseUser = shared.UserResponseUser

// Type of IDP, Identity Provider. Used for login.
//
// This is an alias to an internal type.
type UserResponseUserIdpType = shared.UserResponseUserIdpType

// This is an alias to an internal value.
const UserResponseUserIdpTypeNvidia = shared.UserResponseUserIdpTypeNvidia

// This is an alias to an internal value.
const UserResponseUserIdpTypeEnterprise = shared.UserResponseUserIdpTypeEnterprise

// List of roles that the user have
//
// This is an alias to an internal type.
type UserResponseUserRole = shared.UserResponseUserRole

// Information about the Organization
//
// This is an alias to an internal type.
type UserResponseUserRolesOrg = shared.UserResponseUserRolesOrg

// Org Owner Alternate Contact
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgAlternateContact = shared.UserResponseUserRolesOrgAlternateContact

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgInfinityManagerSettings = shared.UserResponseUserRolesOrgInfinityManagerSettings

// Org owner.
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgOrgOwner = shared.UserResponseUserRolesOrgOrgOwner

// Product Enablement
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductEnablement = shared.UserResponseUserRolesOrgProductEnablement

// Product Enablement Types
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductEnablementsType = shared.UserResponseUserRolesOrgProductEnablementsType

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeNgcAdminEval = shared.UserResponseUserRolesOrgProductEnablementsTypeNgcAdminEval

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeNgcAdminNfr = shared.UserResponseUserRolesOrgProductEnablementsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeNgcAdminCommercial = shared.UserResponseUserRolesOrgProductEnablementsTypeNgcAdminCommercial

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeEmsEval = shared.UserResponseUserRolesOrgProductEnablementsTypeEmsEval

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeEmsNfr = shared.UserResponseUserRolesOrgProductEnablementsTypeEmsNfr

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeEmsCommercial = shared.UserResponseUserRolesOrgProductEnablementsTypeEmsCommercial

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductEnablementsTypeNgcAdminDeveloper = shared.UserResponseUserRolesOrgProductEnablementsTypeNgcAdminDeveloper

// Purchase Order.
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductEnablementsPoDetail = shared.UserResponseUserRolesOrgProductEnablementsPoDetail

// Product Subscription
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductSubscription = shared.UserResponseUserRolesOrgProductSubscription

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType = shared.UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementType

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval = shared.UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr = shared.UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical = shared.UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial = shared.UserResponseUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgProductSubscriptionsType = shared.UserResponseUserRolesOrgProductSubscriptionsType

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminEval = shared.UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminEval

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminNfr = shared.UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial = shared.UserResponseUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgRepoScanSettings = shared.UserResponseUserRolesOrgRepoScanSettings

// This is an alias to an internal type.
type UserResponseUserRolesOrgType = shared.UserResponseUserRolesOrgType

// This is an alias to an internal value.
const UserResponseUserRolesOrgTypeUnknown = shared.UserResponseUserRolesOrgTypeUnknown

// This is an alias to an internal value.
const UserResponseUserRolesOrgTypeCloud = shared.UserResponseUserRolesOrgTypeCloud

// This is an alias to an internal value.
const UserResponseUserRolesOrgTypeEnterprise = shared.UserResponseUserRolesOrgTypeEnterprise

// This is an alias to an internal value.
const UserResponseUserRolesOrgTypeIndividual = shared.UserResponseUserRolesOrgTypeIndividual

// Users information.
//
// This is an alias to an internal type.
type UserResponseUserRolesOrgUsersInfo = shared.UserResponseUserRolesOrgUsersInfo

// Information about the user who is attempting to run the job
//
// This is an alias to an internal type.
type UserResponseUserRolesTargetSystemUserIdentifier = shared.UserResponseUserRolesTargetSystemUserIdentifier

// Information about the team
//
// This is an alias to an internal type.
type UserResponseUserRolesTeam = shared.UserResponseUserRolesTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserResponseUserRolesTeamInfinityManagerSettings = shared.UserResponseUserRolesTeamInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserResponseUserRolesTeamRepoScanSettings = shared.UserResponseUserRolesTeamRepoScanSettings

// represents user storage quota for a given ace and available unused storage
//
// This is an alias to an internal type.
type UserResponseUserStorageQuota = shared.UserResponseUserStorageQuota

// Metadata information about the user.
//
// This is an alias to an internal type.
type UserResponseUserUserMetadata = shared.UserResponseUserUserMetadata
