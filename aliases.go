// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiagpucloud

import (
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/internal/apierror"
	"github.com/stainless-sdks/nvidia-gpu-cloud-go/shared"
)

type Error = apierror.Error

// response containing a list of all metering queries results
//
// This is an alias to an internal type.
type MeteringResultList = shared.MeteringResultList

// result of a single measurement query
//
// This is an alias to an internal type.
type MeteringResultListMeasurement = shared.MeteringResultListMeasurement

// object for a single series in the measurement
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSery = shared.MeteringResultListMeasurementsSery

// object for measurement tags which identifies a measuurement series
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSeriesTag = shared.MeteringResultListMeasurementsSeriesTag

// object for the measurement values
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSeriesValue = shared.MeteringResultListMeasurementsSeriesValue

// This is an alias to an internal type.
type MeteringResultListRequestStatus = shared.MeteringResultListRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type MeteringResultListRequestStatusStatusCode = shared.MeteringResultListRequestStatusStatusCode

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnknown = shared.MeteringResultListRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeSuccess = shared.MeteringResultListRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnauthorized = shared.MeteringResultListRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodePaymentRequired = shared.MeteringResultListRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeForbidden = shared.MeteringResultListRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeTimeout = shared.MeteringResultListRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeExists = shared.MeteringResultListRequestStatusStatusCodeExists

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeNotFound = shared.MeteringResultListRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInternalError = shared.MeteringResultListRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequest = shared.MeteringResultListRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequestVersion = shared.MeteringResultListRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequestData = shared.MeteringResultListRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeMethodNotAllowed = shared.MeteringResultListRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeConflict = shared.MeteringResultListRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnprocessableEntity = shared.MeteringResultListRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeTooManyRequests = shared.MeteringResultListRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInsufficientStorage = shared.MeteringResultListRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeServiceUnavailable = shared.MeteringResultListRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodePayloadTooLarge = shared.MeteringResultListRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeNotAcceptable = shared.MeteringResultListRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons = shared.MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeBadGateway = shared.MeteringResultListRequestStatusStatusCodeBadGateway

// about one user
//
// This is an alias to an internal type.
type User = shared.User

// NCA role
//
// This is an alias to an internal type.
type UserNcaRole = shared.UserNcaRole

// This is an alias to an internal value.
const UserNcaRoleUnknown = shared.UserNcaRoleUnknown

// This is an alias to an internal value.
const UserNcaRoleAdministrator = shared.UserNcaRoleAdministrator

// This is an alias to an internal value.
const UserNcaRoleMember = shared.UserNcaRoleMember

// This is an alias to an internal value.
const UserNcaRoleOwner = shared.UserNcaRoleOwner

// This is an alias to an internal value.
const UserNcaRolePending = shared.UserNcaRolePending

// This is an alias to an internal type.
type UserRequestStatus = shared.UserRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserRequestStatusStatusCode = shared.UserRequestStatusStatusCode

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnknown = shared.UserRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserRequestStatusStatusCodeSuccess = shared.UserRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnauthorized = shared.UserRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserRequestStatusStatusCodePaymentRequired = shared.UserRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserRequestStatusStatusCodeForbidden = shared.UserRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserRequestStatusStatusCodeTimeout = shared.UserRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserRequestStatusStatusCodeExists = shared.UserRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserRequestStatusStatusCodeNotFound = shared.UserRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInternalError = shared.UserRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequest = shared.UserRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequestVersion = shared.UserRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequestData = shared.UserRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserRequestStatusStatusCodeMethodNotAllowed = shared.UserRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserRequestStatusStatusCodeConflict = shared.UserRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnprocessableEntity = shared.UserRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserRequestStatusStatusCodeTooManyRequests = shared.UserRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInsufficientStorage = shared.UserRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserRequestStatusStatusCodeServiceUnavailable = shared.UserRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserRequestStatusStatusCodePayloadTooLarge = shared.UserRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserRequestStatusStatusCodeNotAcceptable = shared.UserRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserRequestStatusStatusCodeBadGateway = shared.UserRequestStatusStatusCodeBadGateway

// information about the user
//
// This is an alias to an internal type.
type UserUser = shared.UserUser

// Type of IDP, Identity Provider. Used for login.
//
// This is an alias to an internal type.
type UserUserIdpType = shared.UserUserIdpType

// This is an alias to an internal value.
const UserUserIdpTypeNvidia = shared.UserUserIdpTypeNvidia

// This is an alias to an internal value.
const UserUserIdpTypeEnterprise = shared.UserUserIdpTypeEnterprise

// List of roles that the user have
//
// This is an alias to an internal type.
type UserUserRole = shared.UserUserRole

// Information about the Organization
//
// This is an alias to an internal type.
type UserUserRolesOrg = shared.UserUserRolesOrg

// Org Owner Alternate Contact
//
// This is an alias to an internal type.
type UserUserRolesOrgAlternateContact = shared.UserUserRolesOrgAlternateContact

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserUserRolesOrgInfinityManagerSettings = shared.UserUserRolesOrgInfinityManagerSettings

// Org owner.
//
// This is an alias to an internal type.
type UserUserRolesOrgOrgOwner = shared.UserUserRolesOrgOrgOwner

// Product Enablement
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablement = shared.UserUserRolesOrgProductEnablement

// Product Enablement Types
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablementsType = shared.UserUserRolesOrgProductEnablementsType

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminEval = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminEval

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminNfr = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsEval = shared.UserUserRolesOrgProductEnablementsTypeEmsEval

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsNfr = shared.UserUserRolesOrgProductEnablementsTypeEmsNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsCommercial = shared.UserUserRolesOrgProductEnablementsTypeEmsCommercial

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper

// Purchase Order.
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablementsPoDetail = shared.UserUserRolesOrgProductEnablementsPoDetail

// Product Subscription
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscription = shared.UserUserRolesOrgProductSubscription

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscriptionsEmsEntitlementType = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementType

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscriptionsType = shared.UserUserRolesOrgProductSubscriptionsType

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserUserRolesOrgRepoScanSettings = shared.UserUserRolesOrgRepoScanSettings

// This is an alias to an internal type.
type UserUserRolesOrgType = shared.UserUserRolesOrgType

// This is an alias to an internal value.
const UserUserRolesOrgTypeUnknown = shared.UserUserRolesOrgTypeUnknown

// This is an alias to an internal value.
const UserUserRolesOrgTypeCloud = shared.UserUserRolesOrgTypeCloud

// This is an alias to an internal value.
const UserUserRolesOrgTypeEnterprise = shared.UserUserRolesOrgTypeEnterprise

// This is an alias to an internal value.
const UserUserRolesOrgTypeIndividual = shared.UserUserRolesOrgTypeIndividual

// Users information.
//
// This is an alias to an internal type.
type UserUserRolesOrgUsersInfo = shared.UserUserRolesOrgUsersInfo

// Information about the user who is attempting to run the job
//
// This is an alias to an internal type.
type UserUserRolesTargetSystemUserIdentifier = shared.UserUserRolesTargetSystemUserIdentifier

// Information about the team
//
// This is an alias to an internal type.
type UserUserRolesTeam = shared.UserUserRolesTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserUserRolesTeamInfinityManagerSettings = shared.UserUserRolesTeamInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserUserRolesTeamRepoScanSettings = shared.UserUserRolesTeamRepoScanSettings

// represents user storage quota for a given ace and available unused storage
//
// This is an alias to an internal type.
type UserUserStorageQuota = shared.UserUserStorageQuota

// Metadata information about the user.
//
// This is an alias to an internal type.
type UserUserUserMetadata = shared.UserUserUserMetadata

// Response for a list of user invitations.
//
// This is an alias to an internal type.
type UserInvitationList = shared.UserInvitationList

// User invitation to an NGC org or team
//
// This is an alias to an internal type.
type UserInvitationListInvitation = shared.UserInvitationListInvitation

// Type of invitation. The invitation is either to an organization or to a team
// within organization.
//
// This is an alias to an internal type.
type UserInvitationListInvitationsType = shared.UserInvitationListInvitationsType

// This is an alias to an internal value.
const UserInvitationListInvitationsTypeOrganization = shared.UserInvitationListInvitationsTypeOrganization

// This is an alias to an internal value.
const UserInvitationListInvitationsTypeTeam = shared.UserInvitationListInvitationsTypeTeam

// object that describes the pagination information
//
// This is an alias to an internal type.
type UserInvitationListPaginationInfo = shared.UserInvitationListPaginationInfo

// This is an alias to an internal type.
type UserInvitationListRequestStatus = shared.UserInvitationListRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserInvitationListRequestStatusStatusCode = shared.UserInvitationListRequestStatusStatusCode

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeUnknown = shared.UserInvitationListRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeSuccess = shared.UserInvitationListRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeUnauthorized = shared.UserInvitationListRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodePaymentRequired = shared.UserInvitationListRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeForbidden = shared.UserInvitationListRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeTimeout = shared.UserInvitationListRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeExists = shared.UserInvitationListRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeNotFound = shared.UserInvitationListRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeInternalError = shared.UserInvitationListRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeInvalidRequest = shared.UserInvitationListRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeInvalidRequestVersion = shared.UserInvitationListRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeInvalidRequestData = shared.UserInvitationListRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeMethodNotAllowed = shared.UserInvitationListRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeConflict = shared.UserInvitationListRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeUnprocessableEntity = shared.UserInvitationListRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeTooManyRequests = shared.UserInvitationListRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeInsufficientStorage = shared.UserInvitationListRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeServiceUnavailable = shared.UserInvitationListRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodePayloadTooLarge = shared.UserInvitationListRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeNotAcceptable = shared.UserInvitationListRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserInvitationListRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserInvitationListRequestStatusStatusCodeBadGateway = shared.UserInvitationListRequestStatusStatusCodeBadGateway
