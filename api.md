# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#MeteringResultList">MeteringResultList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#UserInvitationList">UserInvitationList</a>

# Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Org">Org</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgList">OrgList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Team">Team</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamList">TeamList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserList">UserList</a>

Methods:

- <code title="post /v3/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgNewParams">OrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUpdateParams">OrgUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgListParams">OrgListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgList">OrgList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserNewParams">OrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/orgs/{org-name}/users/{user-email-or-id}">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserListParams">OrgUserListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserList">UserList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{id}">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserDeleteParams">OrgUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/users/{id}/add-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserAddRoleParams">OrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/orgs/{org-name}/users/nca-invitations">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.NcaInvitations">NcaInvitations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserNcaInvitationsParams">OrgUserNcaInvitationsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{user-email-or-id}/remove-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserRemoveRoleParams">OrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamDeleteResponse">OrgTeamDeleteResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/teams/{team-name}">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/teams/{team-name}">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUpdateParams">OrgTeamUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/teams">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamListParams">OrgTeamListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamList">TeamList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/teams/{team-name}">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamDeleteResponse">OrgTeamDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>

Methods:

- <code title="get /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{id}">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserDeleteParams">OrgTeamUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/team/{team-name}/users/{id}/add-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserAddRoleParams">OrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/orgs/{org-name}/teams/{team-name}/users/nca-invitations">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.NcaInvitations">NcaInvitations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserNcaInvitationsParams">OrgTeamUserNcaInvitationsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}/remove-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserRemoveRoleParams">OrgTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/team/{team-name}/users/{id}/update-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserUpdateRoleParams">OrgTeamUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProtoOrg

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgInvitation">OrgInvitation</a>

Methods:

- <code title="post /v3/orgs/proto-org">client.Orgs.ProtoOrg.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgProtoOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgProtoOrgNewParams">OrgProtoOrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/orgs/proto-org/validate">client.Orgs.ProtoOrg.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgProtoOrgService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgProtoOrgValidateParams">OrgProtoOrgValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgInvitation">OrgInvitation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Credits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#CreditsHistory">CreditsHistory</a>

Methods:

- <code title="get /v2/orgs/{org-name}/credits">client.Orgs.Credits.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgCreditService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#CreditsHistory">CreditsHistory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AuditLogs">AuditLogs</a>

# Admin

Methods:

- <code title="post /v2/admin/backfill-orgs-to-kratos">client.Admin.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminService.BackfillOrgsToKratos">BackfillOrgsToKratos</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orgs

Methods:

- <code title="post /v2/admin/orgs">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgNewParams">AdminOrgNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org/{org-name}">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUpdateParams">AdminOrgUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/admin/org">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgListParams">AdminOrgListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/admin/orgs/ncaIds">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.NcaIDs">NcaIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgNcaIDsParams">AdminOrgNcaIDsParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ProtoOrg

Methods:

- <code title="post /v3/admin/orgs">client.Admin.Orgs.ProtoOrg.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgProtoOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgProtoOrgNewParams">AdminOrgProtoOrgNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Methods:

- <code title="patch /v2/admin/org/{org-name}/users/{id}/update-role">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserUpdateRoleParams">AdminOrgUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Offboarded

Methods:

- <code title="get /v2/admin/orgs/offboarded">client.Admin.Orgs.Offboarded.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOffboardedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOffboardedListParams">AdminOrgOffboardedListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserCRMSyncResponse">AdminUserCRMSyncResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserOrgOwnerBackfillResponse">AdminUserOrgOwnerBackfillResponse</a>

Methods:

- <code title="get /v2/admin/users/me">client.Admin.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserGetParams">AdminUserGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{id}/crm-sync">client.Admin.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserService.CRMSync">CRMSync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserCRMSyncResponse">AdminUserCRMSyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/users/invite">client.Admin.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserInviteParams">AdminUserInviteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{id}/migrate-deprecated-roles">client.Admin.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserService.MigrateDeprecatedRoles">MigrateDeprecatedRoles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{user-id}/org-owner-backfill">client.Admin.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserService.OrgOwnerBackfill">OrgOwnerBackfill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminUserOrgOwnerBackfillResponse">AdminUserOrgOwnerBackfillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Org

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOrgOwnerBackfillResponse">AdminOrgOrgOwnerBackfillResponse</a>

Methods:

- <code title="get /v3/admin/org/{org-name}">client.Admin.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/admin/org/{org-name}">client.Admin.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUpdateParams">AdminOrgUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org">client.Admin.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgListParams">AdminOrgListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/org-owner-backfill">client.Admin.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgService.OrgOwnerBackfill">OrgOwnerBackfill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOrgOwnerBackfillResponse">AdminOrgOrgOwnerBackfillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserDeleteResponse">AdminOrgUserDeleteResponse</a>

Methods:

- <code title="post /v2/admin/org/{org-name}/users">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserNewParams">AdminOrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/admin/org/{org-name}/users/{id}">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserDeleteResponse">AdminOrgUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/users/{id}">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserAddParams">AdminOrgUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/users/{id}/add-role">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserAddRoleParams">AdminOrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Team

#### Users

Methods:

- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users">client.Admin.Org.Team.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserNewParams">AdminOrgTeamUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users/{id}">client.Admin.Org.Team.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserAddParams">AdminOrgTeamUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/team/{team-name}/users/{id}/add-role">client.Admin.Org.Team.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUserAddRoleParams">AdminOrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Enablement

Methods:

- <code title="post /v2/admin/org/{org-name}/enablement">client.Admin.Org.Enablement.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgEnablementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgEnablementNewParams">AdminOrgEnablementNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Teams

Methods:

- <code title="get /v2/admin/org/{org-name}/teams/{team-name}">client.Admin.Org.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/teams/{team-name}">client.Admin.Org.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamUpdateParams">AdminOrgTeamUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org/{org-name}/teams">client.Admin.Org.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgTeamListParams">AdminOrgTeamListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Entitlements

Methods:

- <code title="get /v2/admin/org/{org-name}/entitlements">client.Admin.Org.Entitlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgEntitlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgEntitlementListParams">AdminOrgEntitlementListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Entitlements

Methods:

- <code title="get /v2/admin/entitlements">client.Admin.Entitlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminEntitlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminEntitlementListParams">AdminEntitlementListParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserKey">UserKey</a>

Methods:

- <code title="get /v2/users/me">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserGetParams">UserGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/users/me">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKey

Methods:

- <code title="post /v2/users/me/api-key">client.Users.APIKey.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserAPIKeyService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserKey">UserKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Org

## Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamNewResponse">OrgTeamNewResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/teams">client.Org.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamNewParams">OrgTeamNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamNewResponse">OrgTeamNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserAddResponse">OrgTeamUserAddResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserRemoveResponse">OrgTeamUserRemoveResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/team/{team-name}/users">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserNewParams">OrgTeamUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users/{id}">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserGetParams">OrgTeamUserGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserListParams">OrgTeamUserListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserList">UserList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/org/{org-name}/team/{team-name}/users/{id}">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserAddParams">OrgTeamUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserAddResponse">OrgTeamUserAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/team/{team-name}/users/{id}">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserRemoveResponse">OrgTeamUserRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AuditLogsPresignedURL">AuditLogsPresignedURL</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRemoveResponse">OrgAuditLogRemoveResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRequestResponse">OrgAuditLogRequestResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/auditLogs/{log-id}">client.Org.AuditLogs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, logID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AuditLogsPresignedURL">AuditLogsPresignedURL</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/auditLogs">client.Org.AuditLogs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AuditLogs">AuditLogs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/auditLogs">client.Org.AuditLogs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRemoveParams">OrgAuditLogRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRemoveResponse">OrgAuditLogRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/org/{org-name}/auditLogs">client.Org.AuditLogs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRequestParams">OrgAuditLogRequestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgAuditLogRequestResponse">OrgAuditLogRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserRemoveResponse">OrgUserRemoveResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/users/{id}">client.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserGetParams">OrgUserGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users/{id}/invite">client.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInviteParams">OrgUserInviteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/users/{id}">client.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserRemoveParams">OrgUserRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserRemoveResponse">OrgUserRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/users/{id}/update-role">client.Org.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserUpdateRoleParams">OrgUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Invitations

Methods:

- <code title="get /v2/org/{org-name}/users/invitations">client.Org.Users.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationListParams">OrgUserInvitationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#UserInvitationList">UserInvitationList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users/invitations/{id}/resend-invitation-email">client.Org.Users.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationService.ResendInvitationEmail">ResendInvitationEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Team

### Users

#### Invitations

Methods:

- <code title="get /v2/org/{org-name}/team/{team-name}/users/invitations">client.Org.Team.Users.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserInvitationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserInvitationListParams">OrgTeamUserInvitationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#UserInvitationList">UserInvitationList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users/invitations/{id}/resend-invitation-email">client.Org.Team.Users.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamUserInvitationService.ResendInvitationEmail">ResendInvitationEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### StarfleetIDs

Methods:

- <code title="get /v2/org/{org-name}/team/{team-name}/starfleetIds/{starfleet-id}">client.Org.Team.StarfleetIDs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgTeamStarfleetIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, starfleetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## StarfleetIDs

Methods:

- <code title="get /v2/org/{org-name}/starfleetIds/{starfleet-id}">client.Org.StarfleetIDs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgStarfleetIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, starfleetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Meterings

Methods:

- <code title="get /v2/org/{org-name}/metering">client.Org.Meterings.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgMeteringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgMeteringListParams">OrgMeteringListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#MeteringResultList">MeteringResultList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Gpupeak

Methods:

- <code title="get /v2/org/{org-name}/meterings/gpupeak">client.Org.Meterings.Gpupeak.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgMeteringGpupeakService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgMeteringGpupeakListParams">OrgMeteringGpupeakListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#MeteringResultList">MeteringResultList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Services

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Health">Health</a>

Methods:

- <code title="get /health">client.Services.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#ServiceService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Health">Health</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/all">client.Services.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#ServiceService.HealthAll">HealthAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#ServiceHealthAllParams">ServiceHealthAllParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#Health">Health</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Version

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#PackageVersionList">PackageVersionList</a>

Methods:

- <code title="get /version">client.Services.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#ServiceVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#ServiceVersionGetParams">ServiceVersionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#PackageVersionList">PackageVersionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RegistryMeteringDownsampling

Methods:

- <code title="get /v2/admin/org/{org-name}/registry/metering/downsample">client.RegistryMeteringDownsampling.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#RegistryMeteringDownsamplingService.Downsample">Downsample</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#RegistryMeteringDownsamplingDownsampleParams">RegistryMeteringDownsamplingDownsampleParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AdminOrg

## Organizations

Methods:

- <code title="get /v2/admin/org/validate">client.AdminOrg.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOrganizationService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgOrganizationValidateParams">AdminOrgOrganizationValidateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Roles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserRoleDefinitions">UserRoleDefinitions</a>

Methods:

- <code title="get /roles">client.Roles.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#RoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#RoleListParams">RoleListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#UserRoleDefinitions">UserRoleDefinitions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PublicKeys

Methods:

- <code title="get /public-keys">client.PublicKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#PublicKeyService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OrgUsers

Methods:

- <code title="delete /v2/org/{org-name}/users/{id}/remove-role">client.OrgUsers.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserRemoveRoleParams">OrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invitations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationDeleteResponse">OrgUserInvitationDeleteResponse</a>

Methods:

- <code title="delete /v2/org/{org-name}/users/invitations/{id}">client.OrgUsers.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#OrgUserInvitationDeleteResponse">OrgUserInvitationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TeamUsers

Methods:

- <code title="delete /v2/org/{org-name}/team/{team-name}/users/{id}/remove-role">client.TeamUsers.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamUserRemoveRoleParams">TeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invitations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamUserInvitationDeleteResponse">TeamUserInvitationDeleteResponse</a>

Methods:

- <code title="delete /v2/org/{org-name}/team/{team-name}/users/invitations/{id}">client.TeamUsers.Invitations.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamUserInvitationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#TeamUserInvitationDeleteResponse">TeamUserInvitationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AdminOrgUsers

Methods:

- <code title="delete /v2/admin/org/{org-name}/users/{id}/remove-role">client.AdminOrgUsers.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminOrgUserRemoveRoleParams">AdminOrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AdminTeamUsers

Methods:

- <code title="delete /v2/admin/org/{org-name}/team/{team-name}/users/{id}/remove-role">client.AdminTeamUsers.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go">nvidiagpucloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#AdminTeamUserRemoveRoleParams">AdminTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SwaggerResources

Methods:

- <code title="post /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configuration

### Ui

Methods:

- <code title="post /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationUiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationUiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationUiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationUiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Security

Methods:

- <code title="post /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationSecurityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationSecurityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationSecurityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-gpu-cloud-go#SwaggerResourceConfigurationSecurityService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
