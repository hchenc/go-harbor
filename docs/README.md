# Go API client for swagger

These APIs provide services for manipulating Harbor project.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/v2.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ArtifactApi* | [**AddLabel**](docs/ArtifactApi.md#addlabel) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels | Add label to artifact
*ArtifactApi* | [**CopyArtifact**](docs/ArtifactApi.md#copyartifact) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts | Copy artifact
*ArtifactApi* | [**CreateTag**](docs/ArtifactApi.md#createtag) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags | Create tag
*ArtifactApi* | [**DeleteArtifact**](docs/ArtifactApi.md#deleteartifact) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference} | Delete the specific artifact
*ArtifactApi* | [**DeleteTag**](docs/ArtifactApi.md#deletetag) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name} | Delete tag
*ArtifactApi* | [**GetAddition**](docs/ArtifactApi.md#getaddition) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition} | Get the addition of the specific artifact
*ArtifactApi* | [**GetArtifact**](docs/ArtifactApi.md#getartifact) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference} | Get the specific artifact
*ArtifactApi* | [**GetVulnerabilitiesAddition**](docs/ArtifactApi.md#getvulnerabilitiesaddition) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities | Get the vulnerabilities addition of the specific artifact
*ArtifactApi* | [**ListArtifacts**](docs/ArtifactApi.md#listartifacts) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts | List artifacts
*ArtifactApi* | [**ListTags**](docs/ArtifactApi.md#listtags) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags | List tags
*ArtifactApi* | [**RemoveLabel**](docs/ArtifactApi.md#removelabel) | **Delete** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id} | Remove label from artifact
*AuditlogApi* | [**ListAuditLogs**](docs/AuditlogApi.md#listauditlogs) | **Get** /audit-logs | Get recent logs of the projects which the user is a member of
*ConfigureApi* | [**GetConfigurations**](docs/ConfigureApi.md#getconfigurations) | **Get** /configurations | Get system configurations.
*ConfigureApi* | [**GetInternalconfig**](docs/ConfigureApi.md#getinternalconfig) | **Get** /internalconfig | Get internal configurations.
*ConfigureApi* | [**UpdateConfigurations**](docs/ConfigureApi.md#updateconfigurations) | **Put** /configurations | Modify system configurations.
*GcApi* | [**CreateGCSchedule**](docs/GcApi.md#creategcschedule) | **Post** /system/gc/schedule | Create a gc schedule.
*GcApi* | [**GetGC**](docs/GcApi.md#getgc) | **Get** /system/gc/{gc_id} | Get gc status.
*GcApi* | [**GetGCHistory**](docs/GcApi.md#getgchistory) | **Get** /system/gc | Get gc results.
*GcApi* | [**GetGCLog**](docs/GcApi.md#getgclog) | **Get** /system/gc/{gc_id}/log | Get gc job log.
*GcApi* | [**GetGCSchedule**](docs/GcApi.md#getgcschedule) | **Get** /system/gc/schedule | Get gc&#39;s schedule.
*GcApi* | [**UpdateGCSchedule**](docs/GcApi.md#updategcschedule) | **Put** /system/gc/schedule | Update gc&#39;s schedule.
*HealthApi* | [**GetHealth**](docs/HealthApi.md#gethealth) | **Get** /health | Check the status of Harbor components
*IconApi* | [**GetIcon**](docs/IconApi.md#geticon) | **Get** /icons/{digest} | Get artifact icon
*ImmutableApi* | [**CreateImmuRule**](docs/ImmutableApi.md#createimmurule) | **Post** /projects/{project_name_or_id}/immutabletagrules | Add an immutable tag rule to current project
*ImmutableApi* | [**DeleteImmuRule**](docs/ImmutableApi.md#deleteimmurule) | **Delete** /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id} | Delete the immutable tag rule.
*ImmutableApi* | [**ListImmuRules**](docs/ImmutableApi.md#listimmurules) | **Get** /projects/{project_name_or_id}/immutabletagrules | List all immutable tag rules of current project
*ImmutableApi* | [**UpdateImmuRule**](docs/ImmutableApi.md#updateimmurule) | **Put** /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id} | Update the immutable tag rule or enable or disable the rule
*LabelApi* | [**CreateLabel**](docs/LabelApi.md#createlabel) | **Post** /labels | Post creates a label
*LabelApi* | [**DeleteLabel**](docs/LabelApi.md#deletelabel) | **Delete** /labels/{label_id} | Delete the label specified by ID.
*LabelApi* | [**GetLabelByID**](docs/LabelApi.md#getlabelbyid) | **Get** /labels/{label_id} | Get the label specified by ID.
*LabelApi* | [**ListLabels**](docs/LabelApi.md#listlabels) | **Get** /labels | List labels according to the query strings.
*LabelApi* | [**UpdateLabel**](docs/LabelApi.md#updatelabel) | **Put** /labels/{label_id} | Update the label properties.
*LdapApi* | [**ImportLdapUser**](docs/LdapApi.md#importldapuser) | **Post** /ldap/users/import | Import selected available ldap users.
*LdapApi* | [**PingLdap**](docs/LdapApi.md#pingldap) | **Post** /ldap/ping | Ping available ldap service.
*LdapApi* | [**SearchLdapGroup**](docs/LdapApi.md#searchldapgroup) | **Get** /ldap/groups/search | Search available ldap groups.
*LdapApi* | [**SearchLdapUser**](docs/LdapApi.md#searchldapuser) | **Get** /ldap/users/search | Search available ldap users.
*MemberApi* | [**CreateProjectMember**](docs/MemberApi.md#createprojectmember) | **Post** /projects/{project_name_or_id}/members | Create project member
*MemberApi* | [**DeleteProjectMember**](docs/MemberApi.md#deleteprojectmember) | **Delete** /projects/{project_name_or_id}/members/{mid} | Delete project member
*MemberApi* | [**GetProjectMember**](docs/MemberApi.md#getprojectmember) | **Get** /projects/{project_name_or_id}/members/{mid} | Get the project member information
*MemberApi* | [**ListProjectMembers**](docs/MemberApi.md#listprojectmembers) | **Get** /projects/{project_name_or_id}/members | Get all project member information
*MemberApi* | [**UpdateProjectMember**](docs/MemberApi.md#updateprojectmember) | **Put** /projects/{project_name_or_id}/members/{mid} | Update project member
*OidcApi* | [**PingOIDC**](docs/OidcApi.md#pingoidc) | **Post** /system/oidc/ping | Test the OIDC endpoint.
*PingApi* | [**GetPing**](docs/PingApi.md#getping) | **Get** /ping | Ping Harbor to check if it&#39;s alive.
*PreheatApi* | [**CreateInstance**](docs/PreheatApi.md#createinstance) | **Post** /p2p/preheat/instances | Create p2p provider instances
*PreheatApi* | [**CreatePolicy**](docs/PreheatApi.md#createpolicy) | **Post** /projects/{project_name}/preheat/policies | Create a preheat policy under a project
*PreheatApi* | [**DeleteInstance**](docs/PreheatApi.md#deleteinstance) | **Delete** /p2p/preheat/instances/{preheat_instance_name} | Delete the specified P2P provider instance
*PreheatApi* | [**DeletePolicy**](docs/PreheatApi.md#deletepolicy) | **Delete** /projects/{project_name}/preheat/policies/{preheat_policy_name} | Delete a preheat policy
*PreheatApi* | [**GetExecution**](docs/PreheatApi.md#getexecution) | **Get** /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id} | Get a execution detail by id
*PreheatApi* | [**GetInstance**](docs/PreheatApi.md#getinstance) | **Get** /p2p/preheat/instances/{preheat_instance_name} | Get a P2P provider instance
*PreheatApi* | [**GetPolicy**](docs/PreheatApi.md#getpolicy) | **Get** /projects/{project_name}/preheat/policies/{preheat_policy_name} | Get a preheat policy
*PreheatApi* | [**GetPreheatLog**](docs/PreheatApi.md#getpreheatlog) | **Get** /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs | Get the log text stream of the specified task for the given execution
*PreheatApi* | [**ListExecutions**](docs/PreheatApi.md#listexecutions) | **Get** /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions | List executions for the given policy
*PreheatApi* | [**ListInstances**](docs/PreheatApi.md#listinstances) | **Get** /p2p/preheat/instances | List P2P provider instances
*PreheatApi* | [**ListPolicies**](docs/PreheatApi.md#listpolicies) | **Get** /projects/{project_name}/preheat/policies | List preheat policies
*PreheatApi* | [**ListProviders**](docs/PreheatApi.md#listproviders) | **Get** /p2p/preheat/providers | List P2P providers
*PreheatApi* | [**ListProvidersUnderProject**](docs/PreheatApi.md#listprovidersunderproject) | **Get** /projects/{project_name}/preheat/providers | Get all providers at project level
*PreheatApi* | [**ListTasks**](docs/PreheatApi.md#listtasks) | **Get** /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks | List all the related tasks for the given execution
*PreheatApi* | [**ManualPreheat**](docs/PreheatApi.md#manualpreheat) | **Post** /projects/{project_name}/preheat/policies/{preheat_policy_name} | Manual preheat
*PreheatApi* | [**PingInstances**](docs/PreheatApi.md#pinginstances) | **Post** /p2p/preheat/instances/ping | Ping status of a instance.
*PreheatApi* | [**StopExecution**](docs/PreheatApi.md#stopexecution) | **Patch** /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id} | Stop a execution
*PreheatApi* | [**UpdateInstance**](docs/PreheatApi.md#updateinstance) | **Put** /p2p/preheat/instances/{preheat_instance_name} | Update the specified P2P provider instance
*PreheatApi* | [**UpdatePolicy**](docs/PreheatApi.md#updatepolicy) | **Put** /projects/{project_name}/preheat/policies/{preheat_policy_name} | Update preheat policy
*ProjectApi* | [**CreateProject**](docs/ProjectApi.md#createproject) | **Post** /projects | Create a new project.
*ProjectApi* | [**DeleteProject**](docs/ProjectApi.md#deleteproject) | **Delete** /projects/{project_name_or_id} | Delete project by projectID
*ProjectApi* | [**GetLogs**](docs/ProjectApi.md#getlogs) | **Get** /projects/{project_name}/logs | Get recent logs of the projects
*ProjectApi* | [**GetProject**](docs/ProjectApi.md#getproject) | **Get** /projects/{project_name_or_id} | Return specific project detail information
*ProjectApi* | [**GetProjectDeletable**](docs/ProjectApi.md#getprojectdeletable) | **Get** /projects/{project_name_or_id}/_deletable | Get the deletable status of the project
*ProjectApi* | [**GetProjectSummary**](docs/ProjectApi.md#getprojectsummary) | **Get** /projects/{project_name_or_id}/summary | Get summary of the project.
*ProjectApi* | [**GetScannerOfProject**](docs/ProjectApi.md#getscannerofproject) | **Get** /projects/{project_name_or_id}/scanner | Get project level scanner
*ProjectApi* | [**HeadProject**](docs/ProjectApi.md#headproject) | **Head** /projects | Check if the project name user provided already exists.
*ProjectApi* | [**ListProjects**](docs/ProjectApi.md#listprojects) | **Get** /projects | List projects
*ProjectApi* | [**ListScannerCandidatesOfProject**](docs/ProjectApi.md#listscannercandidatesofproject) | **Get** /projects/{project_name_or_id}/scanner/candidates | Get scanner registration candidates for configurating project level scanner
*ProjectApi* | [**SetScannerOfProject**](docs/ProjectApi.md#setscannerofproject) | **Put** /projects/{project_name_or_id}/scanner | Configure scanner for the specified project
*ProjectApi* | [**UpdateProject**](docs/ProjectApi.md#updateproject) | **Put** /projects/{project_name_or_id} | Update properties for a selected project.
*ProjectMetadataApi* | [**AddProjectMetadatas**](docs/ProjectMetadataApi.md#addprojectmetadatas) | **Post** /projects/{project_name_or_id}/metadatas/ | Add metadata for the specific project
*ProjectMetadataApi* | [**DeleteProjectMetadata**](docs/ProjectMetadataApi.md#deleteprojectmetadata) | **Delete** /projects/{project_name_or_id}/metadatas/{meta_name} | Delete the specific metadata for the specific project
*ProjectMetadataApi* | [**GetProjectMetadata**](docs/ProjectMetadataApi.md#getprojectmetadata) | **Get** /projects/{project_name_or_id}/metadatas/{meta_name} | Get the specific metadata of the specific project
*ProjectMetadataApi* | [**ListProjectMetadatas**](docs/ProjectMetadataApi.md#listprojectmetadatas) | **Get** /projects/{project_name_or_id}/metadatas/ | Get the metadata of the specific project
*ProjectMetadataApi* | [**UpdateProjectMetadata**](docs/ProjectMetadataApi.md#updateprojectmetadata) | **Put** /projects/{project_name_or_id}/metadatas/{meta_name} | Update the specific metadata for the specific project
*QuotaApi* | [**GetQuota**](docs/QuotaApi.md#getquota) | **Get** /quotas/{id} | Get the specified quota
*QuotaApi* | [**ListQuotas**](docs/QuotaApi.md#listquotas) | **Get** /quotas | List quotas
*QuotaApi* | [**UpdateQuota**](docs/QuotaApi.md#updatequota) | **Put** /quotas/{id} | Update the specified quota
*RegistryApi* | [**CreateRegistry**](docs/RegistryApi.md#createregistry) | **Post** /registries | Create a registry
*RegistryApi* | [**DeleteRegistry**](docs/RegistryApi.md#deleteregistry) | **Delete** /registries/{id} | Delete the specific registry
*RegistryApi* | [**GetRegistry**](docs/RegistryApi.md#getregistry) | **Get** /registries/{id} | Get the specific registry
*RegistryApi* | [**GetRegistryInfo**](docs/RegistryApi.md#getregistryinfo) | **Get** /registries/{id}/info | Get the registry info
*RegistryApi* | [**ListRegistries**](docs/RegistryApi.md#listregistries) | **Get** /registries | List the registries
*RegistryApi* | [**ListRegistryProviderInfos**](docs/RegistryApi.md#listregistryproviderinfos) | **Get** /replication/adapterinfos | List all registered registry provider information
*RegistryApi* | [**ListRegistryProviderTypes**](docs/RegistryApi.md#listregistryprovidertypes) | **Get** /replication/adapters | List registry adapters
*RegistryApi* | [**PingRegistry**](docs/RegistryApi.md#pingregistry) | **Post** /registries/ping | Check status of a registry
*RegistryApi* | [**UpdateRegistry**](docs/RegistryApi.md#updateregistry) | **Put** /registries/{id} | Update the registry
*ReplicationApi* | [**CreateReplicationPolicy**](docs/ReplicationApi.md#createreplicationpolicy) | **Post** /replication/policies | Create a replication policy
*ReplicationApi* | [**DeleteReplicationPolicy**](docs/ReplicationApi.md#deletereplicationpolicy) | **Delete** /replication/policies/{id} | Delete the specific replication policy
*ReplicationApi* | [**GetReplicationExecution**](docs/ReplicationApi.md#getreplicationexecution) | **Get** /replication/executions/{id} | Get the specific replication execution
*ReplicationApi* | [**GetReplicationLog**](docs/ReplicationApi.md#getreplicationlog) | **Get** /replication/executions/{id}/tasks/{task_id}/log | Get the log of the specific replication task
*ReplicationApi* | [**GetReplicationPolicy**](docs/ReplicationApi.md#getreplicationpolicy) | **Get** /replication/policies/{id} | Get the specific replication policy
*ReplicationApi* | [**ListReplicationExecutions**](docs/ReplicationApi.md#listreplicationexecutions) | **Get** /replication/executions | List replication executions
*ReplicationApi* | [**ListReplicationPolicies**](docs/ReplicationApi.md#listreplicationpolicies) | **Get** /replication/policies | List replication policies
*ReplicationApi* | [**ListReplicationTasks**](docs/ReplicationApi.md#listreplicationtasks) | **Get** /replication/executions/{id}/tasks | List replication tasks for a specific execution
*ReplicationApi* | [**StartReplication**](docs/ReplicationApi.md#startreplication) | **Post** /replication/executions | Start one replication execution
*ReplicationApi* | [**StopReplication**](docs/ReplicationApi.md#stopreplication) | **Put** /replication/executions/{id} | Stop the specific replication execution
*ReplicationApi* | [**UpdateReplicationPolicy**](docs/ReplicationApi.md#updatereplicationpolicy) | **Put** /replication/policies/{id} | Update the replication policy
*RepositoryApi* | [**DeleteRepository**](docs/RepositoryApi.md#deleterepository) | **Delete** /projects/{project_name}/repositories/{repository_name} | Delete repository
*RepositoryApi* | [**GetRepository**](docs/RepositoryApi.md#getrepository) | **Get** /projects/{project_name}/repositories/{repository_name} | Get repository
*RepositoryApi* | [**ListAllRepositories**](docs/RepositoryApi.md#listallrepositories) | **Get** /repositories | List all authorized repositories
*RepositoryApi* | [**ListRepositories**](docs/RepositoryApi.md#listrepositories) | **Get** /projects/{project_name}/repositories | List repositories
*RepositoryApi* | [**UpdateRepository**](docs/RepositoryApi.md#updaterepository) | **Put** /projects/{project_name}/repositories/{repository_name} | Update repository
*RetentionApi* | [**CreateRetention**](docs/RetentionApi.md#createretention) | **Post** /retentions | Create Retention Policy
*RetentionApi* | [**DeleteRetention**](docs/RetentionApi.md#deleteretention) | **Delete** /retentions/{id} | Delete Retention Policy
*RetentionApi* | [**GetRentenitionMetadata**](docs/RetentionApi.md#getrentenitionmetadata) | **Get** /retentions/metadatas | Get Retention Metadatas
*RetentionApi* | [**GetRetention**](docs/RetentionApi.md#getretention) | **Get** /retentions/{id} | Get Retention Policy
*RetentionApi* | [**GetRetentionTaskLog**](docs/RetentionApi.md#getretentiontasklog) | **Get** /retentions/{id}/executions/{eid}/tasks/{tid} | Get Retention job task log
*RetentionApi* | [**ListRetentionExecutions**](docs/RetentionApi.md#listretentionexecutions) | **Get** /retentions/{id}/executions | Get Retention executions
*RetentionApi* | [**ListRetentionTasks**](docs/RetentionApi.md#listretentiontasks) | **Get** /retentions/{id}/executions/{eid}/tasks | Get Retention tasks
*RetentionApi* | [**OperateRetentionExecution**](docs/RetentionApi.md#operateretentionexecution) | **Patch** /retentions/{id}/executions/{eid} | Stop a Retention execution
*RetentionApi* | [**TriggerRetentionExecution**](docs/RetentionApi.md#triggerretentionexecution) | **Post** /retentions/{id}/executions | Trigger a Retention Execution
*RetentionApi* | [**UpdateRetention**](docs/RetentionApi.md#updateretention) | **Put** /retentions/{id} | Update Retention Policy
*RobotApi* | [**CreateRobot**](docs/RobotApi.md#createrobot) | **Post** /robots | Create a robot account
*RobotApi* | [**DeleteRobot**](docs/RobotApi.md#deleterobot) | **Delete** /robots/{robot_id} | Delete a robot account
*RobotApi* | [**GetRobotByID**](docs/RobotApi.md#getrobotbyid) | **Get** /robots/{robot_id} | Get a robot account
*RobotApi* | [**ListRobot**](docs/RobotApi.md#listrobot) | **Get** /robots | Get robot account
*RobotApi* | [**RefreshSec**](docs/RobotApi.md#refreshsec) | **Patch** /robots/{robot_id} | Refresh the robot secret
*RobotApi* | [**UpdateRobot**](docs/RobotApi.md#updaterobot) | **Put** /robots/{robot_id} | Update a robot account
*Robotv1Api* | [**CreateRobotV1**](docs/Robotv1Api.md#createrobotv1) | **Post** /projects/{project_name_or_id}/robots | Create a robot account
*Robotv1Api* | [**DeleteRobotV1**](docs/Robotv1Api.md#deleterobotv1) | **Delete** /projects/{project_name_or_id}/robots/{robot_id} | Delete a robot account
*Robotv1Api* | [**GetRobotByIDV1**](docs/Robotv1Api.md#getrobotbyidv1) | **Get** /projects/{project_name_or_id}/robots/{robot_id} | Get a robot account
*Robotv1Api* | [**ListRobotV1**](docs/Robotv1Api.md#listrobotv1) | **Get** /projects/{project_name_or_id}/robots | Get all robot accounts of specified project
*Robotv1Api* | [**UpdateRobotV1**](docs/Robotv1Api.md#updaterobotv1) | **Put** /projects/{project_name_or_id}/robots/{robot_id} | Update status of robot account.
*ScanApi* | [**GetReportLog**](docs/ScanApi.md#getreportlog) | **Get** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log | Get the log of the scan report
*ScanApi* | [**ScanArtifact**](docs/ScanApi.md#scanartifact) | **Post** /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan | Scan the artifact
*ScanAllApi* | [**CreateScanAllSchedule**](docs/ScanAllApi.md#createscanallschedule) | **Post** /system/scanAll/schedule | Create a schedule or a manual trigger for the scan all job.
*ScanAllApi* | [**GetLatestScanAllMetrics**](docs/ScanAllApi.md#getlatestscanallmetrics) | **Get** /scans/all/metrics | Get the metrics of the latest scan all process
*ScanAllApi* | [**GetLatestScheduledScanAllMetrics**](docs/ScanAllApi.md#getlatestscheduledscanallmetrics) | **Get** /scans/schedule/metrics | Get the metrics of the latest scheduled scan all process
*ScanAllApi* | [**GetScanAllSchedule**](docs/ScanAllApi.md#getscanallschedule) | **Get** /system/scanAll/schedule | Get scan all&#39;s schedule.
*ScanAllApi* | [**UpdateScanAllSchedule**](docs/ScanAllApi.md#updatescanallschedule) | **Put** /system/scanAll/schedule | Update scan all&#39;s schedule.
*ScannerApi* | [**CreateScanner**](docs/ScannerApi.md#createscanner) | **Post** /scanners | Create a scanner registration
*ScannerApi* | [**DeleteScanner**](docs/ScannerApi.md#deletescanner) | **Delete** /scanners/{registration_id} | Delete a scanner registration
*ScannerApi* | [**GetScanner**](docs/ScannerApi.md#getscanner) | **Get** /scanners/{registration_id} | Get a scanner registration details
*ScannerApi* | [**GetScannerMetadata**](docs/ScannerApi.md#getscannermetadata) | **Get** /scanners/{registration_id}/metadata | Get the metadata of the specified scanner registration
*ScannerApi* | [**ListScanners**](docs/ScannerApi.md#listscanners) | **Get** /scanners | List scanner registrations
*ScannerApi* | [**PingScanner**](docs/ScannerApi.md#pingscanner) | **Post** /scanners/ping | Tests scanner registration settings
*ScannerApi* | [**SetScannerAsDefault**](docs/ScannerApi.md#setscannerasdefault) | **Patch** /scanners/{registration_id} | Set system default scanner registration
*ScannerApi* | [**UpdateScanner**](docs/ScannerApi.md#updatescanner) | **Put** /scanners/{registration_id} | Update a scanner registration
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **Get** /search | Search for projects, repositories and helm charts
*StatisticApi* | [**GetStatistic**](docs/StatisticApi.md#getstatistic) | **Get** /statistics | Get the statistic information about the projects and repositories
*SystemCVEAllowlistApi* | [**GetSystemCVEAllowlist**](docs/SystemCVEAllowlistApi.md#getsystemcveallowlist) | **Get** /system/CVEAllowlist | Get the system level allowlist of CVE.
*SystemCVEAllowlistApi* | [**PutSystemCVEAllowlist**](docs/SystemCVEAllowlistApi.md#putsystemcveallowlist) | **Put** /system/CVEAllowlist | Update the system level allowlist of CVE.
*SysteminfoApi* | [**GetCert**](docs/SysteminfoApi.md#getcert) | **Get** /systeminfo/getcert | Get default root certificate.
*SysteminfoApi* | [**GetSystemInfo**](docs/SysteminfoApi.md#getsysteminfo) | **Get** /systeminfo | Get general system info
*SysteminfoApi* | [**GetVolumes**](docs/SysteminfoApi.md#getvolumes) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
*UserApi* | [**CreateUser**](docs/UserApi.md#createuser) | **Post** /users | Create a local user.
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /users/{user_id} | Mark a registered user as be removed.
*UserApi* | [**GetCurrentUserInfo**](docs/UserApi.md#getcurrentuserinfo) | **Get** /users/current | Get current user info.
*UserApi* | [**GetCurrentUserPermissions**](docs/UserApi.md#getcurrentuserpermissions) | **Get** /users/current/permissions | Get current user permissions.
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /users/{user_id} | Get a user&#39;s profile.
*UserApi* | [**ListUsers**](docs/UserApi.md#listusers) | **Get** /users | List users
*UserApi* | [**SearchUsers**](docs/UserApi.md#searchusers) | **Get** /users/search | Search users by username
*UserApi* | [**SetCliSecret**](docs/UserApi.md#setclisecret) | **Put** /users/{user_id}/cli_secret | Set CLI secret for a user.
*UserApi* | [**SetUserSysAdmin**](docs/UserApi.md#setusersysadmin) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.
*UserApi* | [**UpdateUserPassword**](docs/UserApi.md#updateuserpassword) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
*UserApi* | [**UpdateUserProfile**](docs/UserApi.md#updateuserprofile) | **Put** /users/{user_id} | Update user&#39;s profile.
*UsergroupApi* | [**CreateUserGroup**](docs/UsergroupApi.md#createusergroup) | **Post** /usergroups | Create user group
*UsergroupApi* | [**DeleteUserGroup**](docs/UsergroupApi.md#deleteusergroup) | **Delete** /usergroups/{group_id} | Delete user group
*UsergroupApi* | [**GetUserGroup**](docs/UsergroupApi.md#getusergroup) | **Get** /usergroups/{group_id} | Get user group information
*UsergroupApi* | [**ListUserGroups**](docs/UsergroupApi.md#listusergroups) | **Get** /usergroups | Get all user groups information
*UsergroupApi* | [**SearchUserGroups**](docs/UsergroupApi.md#searchusergroups) | **Get** /usergroups/search | Search groups by groupname
*UsergroupApi* | [**UpdateUserGroup**](docs/UsergroupApi.md#updateusergroup) | **Put** /usergroups/{group_id} | Update group information
*WebhookApi* | [**CreateWebhookPolicyOfProject**](docs/WebhookApi.md#createwebhookpolicyofproject) | **Post** /projects/{project_name_or_id}/webhook/policies | Create project webhook policy.
*WebhookApi* | [**DeleteWebhookPolicyOfProject**](docs/WebhookApi.md#deletewebhookpolicyofproject) | **Delete** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Delete webhook policy of a project
*WebhookApi* | [**GetSupportedEventTypes**](docs/WebhookApi.md#getsupportedeventtypes) | **Get** /projects/{project_name_or_id}/webhook/events | Get supported event types and notify types.
*WebhookApi* | [**GetWebhookPolicyOfProject**](docs/WebhookApi.md#getwebhookpolicyofproject) | **Get** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Get project webhook policy
*WebhookApi* | [**LastTrigger**](docs/WebhookApi.md#lasttrigger) | **Get** /projects/{project_name_or_id}/webhook/lasttrigger | Get project webhook policy last trigger info
*WebhookApi* | [**ListWebhookPoliciesOfProject**](docs/WebhookApi.md#listwebhookpoliciesofproject) | **Get** /projects/{project_name_or_id}/webhook/policies | List project webhook policies.
*WebhookApi* | [**UpdateWebhookPolicyOfProject**](docs/WebhookApi.md#updatewebhookpolicyofproject) | **Put** /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} | Update webhook policy of a project.
*WebhookjobApi* | [**ListWebhookJobs**](docs/WebhookjobApi.md#listwebhookjobs) | **Get** /projects/{project_name_or_id}/webhook/jobs | List project webhook jobs


## Documentation For Models

 - [Access](docs/Access.md)
 - [AdditionLink](docs/AdditionLink.md)
 - [AdditionLinks](docs/AdditionLinks.md)
 - [Annotations](docs/Annotations.md)
 - [Artifact](docs/Artifact.md)
 - [AuditLog](docs/AuditLog.md)
 - [AuthproxySetting](docs/AuthproxySetting.md)
 - [BoolConfigItem](docs/BoolConfigItem.md)
 - [ChartMetadata](docs/ChartMetadata.md)
 - [ChartVersion](docs/ChartVersion.md)
 - [ComponentHealthStatus](docs/ComponentHealthStatus.md)
 - [Configurations](docs/Configurations.md)
 - [ConfigurationsResponse](docs/ConfigurationsResponse.md)
 - [CveAllowlist](docs/CveAllowlist.md)
 - [CveAllowlistItem](docs/CveAllowlistItem.md)
 - [Errors](docs/Errors.md)
 - [EventType](docs/EventType.md)
 - [Execution](docs/Execution.md)
 - [ExtraAttrs](docs/ExtraAttrs.md)
 - [FilterStyle](docs/FilterStyle.md)
 - [GcHistory](docs/GcHistory.md)
 - [GeneralInfo](docs/GeneralInfo.md)
 - [Icon](docs/Icon.md)
 - [ImmutableRule](docs/ImmutableRule.md)
 - [ImmutableSelector](docs/ImmutableSelector.md)
 - [Instance](docs/Instance.md)
 - [IntegerConfigItem](docs/IntegerConfigItem.md)
 - [InternalConfigurationsResponse](docs/InternalConfigurationsResponse.md)
 - [IsDefault](docs/IsDefault.md)
 - [Label](docs/Label.md)
 - [LdapConf](docs/LdapConf.md)
 - [LdapFailedImportUser](docs/LdapFailedImportUser.md)
 - [LdapImportUsers](docs/LdapImportUsers.md)
 - [LdapPingResult](docs/LdapPingResult.md)
 - [LdapUser](docs/LdapUser.md)
 - [Metadata](docs/Metadata.md)
 - [Metrics](docs/Metrics.md)
 - [ModelError](docs/ModelError.md)
 - [NativeReportSummary](docs/NativeReportSummary.md)
 - [NotifyType](docs/NotifyType.md)
 - [OidcCliSecretReq](docs/OidcCliSecretReq.md)
 - [OidcUserInfo](docs/OidcUserInfo.md)
 - [OverallHealthStatus](docs/OverallHealthStatus.md)
 - [PasswordReq](docs/PasswordReq.md)
 - [Permission](docs/Permission.md)
 - [Platform](docs/Platform.md)
 - [PreheatPolicy](docs/PreheatPolicy.md)
 - [Project](docs/Project.md)
 - [ProjectDeletable](docs/ProjectDeletable.md)
 - [ProjectMember](docs/ProjectMember.md)
 - [ProjectMemberEntity](docs/ProjectMemberEntity.md)
 - [ProjectMetadata](docs/ProjectMetadata.md)
 - [ProjectReq](docs/ProjectReq.md)
 - [ProjectScanner](docs/ProjectScanner.md)
 - [ProjectSummary](docs/ProjectSummary.md)
 - [ProjectSummaryQuota](docs/ProjectSummaryQuota.md)
 - [ProviderUnderProject](docs/ProviderUnderProject.md)
 - [Quota](docs/Quota.md)
 - [QuotaRefObject](docs/QuotaRefObject.md)
 - [QuotaUpdateReq](docs/QuotaUpdateReq.md)
 - [Reference](docs/Reference.md)
 - [Registry](docs/Registry.md)
 - [RegistryCredential](docs/RegistryCredential.md)
 - [RegistryEndpoint](docs/RegistryEndpoint.md)
 - [RegistryInfo](docs/RegistryInfo.md)
 - [RegistryPing](docs/RegistryPing.md)
 - [RegistryProviderCredentialPattern](docs/RegistryProviderCredentialPattern.md)
 - [RegistryProviderEndpointPattern](docs/RegistryProviderEndpointPattern.md)
 - [RegistryProviderInfo](docs/RegistryProviderInfo.md)
 - [RegistryUpdate](docs/RegistryUpdate.md)
 - [ReplicationExecution](docs/ReplicationExecution.md)
 - [ReplicationFilter](docs/ReplicationFilter.md)
 - [ReplicationPolicy](docs/ReplicationPolicy.md)
 - [ReplicationTask](docs/ReplicationTask.md)
 - [ReplicationTrigger](docs/ReplicationTrigger.md)
 - [ReplicationTriggerSettings](docs/ReplicationTriggerSettings.md)
 - [Repository](docs/Repository.md)
 - [ResourceList](docs/ResourceList.md)
 - [RetentionExecution](docs/RetentionExecution.md)
 - [RetentionExecutionTask](docs/RetentionExecutionTask.md)
 - [RetentionMetadata](docs/RetentionMetadata.md)
 - [RetentionPolicy](docs/RetentionPolicy.md)
 - [RetentionPolicyScope](docs/RetentionPolicyScope.md)
 - [RetentionRule](docs/RetentionRule.md)
 - [RetentionRuleMetadata](docs/RetentionRuleMetadata.md)
 - [RetentionRuleParamMetadata](docs/RetentionRuleParamMetadata.md)
 - [RetentionRuleTrigger](docs/RetentionRuleTrigger.md)
 - [RetentionSelector](docs/RetentionSelector.md)
 - [RetentionSelectorMetadata](docs/RetentionSelectorMetadata.md)
 - [Robot](docs/Robot.md)
 - [RobotCreate](docs/RobotCreate.md)
 - [RobotCreateV1](docs/RobotCreateV1.md)
 - [RobotCreated](docs/RobotCreated.md)
 - [RobotPermission](docs/RobotPermission.md)
 - [RobotSec](docs/RobotSec.md)
 - [RoleRequest](docs/RoleRequest.md)
 - [ScanOverview](docs/ScanOverview.md)
 - [Scanner](docs/Scanner.md)
 - [ScannerAdapterMetadata](docs/ScannerAdapterMetadata.md)
 - [ScannerCapability](docs/ScannerCapability.md)
 - [ScannerRegistration](docs/ScannerRegistration.md)
 - [ScannerRegistrationReq](docs/ScannerRegistrationReq.md)
 - [ScannerRegistrationSettings](docs/ScannerRegistrationSettings.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleObj](docs/ScheduleObj.md)
 - [Search](docs/Search.md)
 - [SearchRepository](docs/SearchRepository.md)
 - [SearchResult](docs/SearchResult.md)
 - [StartReplicationExecution](docs/StartReplicationExecution.md)
 - [Statistic](docs/Statistic.md)
 - [Stats](docs/Stats.md)
 - [Storage](docs/Storage.md)
 - [StringConfigItem](docs/StringConfigItem.md)
 - [SupportedWebhookEventTypes](docs/SupportedWebhookEventTypes.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [Tag](docs/Tag.md)
 - [Task](docs/Task.md)
 - [UserCreationReq](docs/UserCreationReq.md)
 - [UserEntity](docs/UserEntity.md)
 - [UserGroup](docs/UserGroup.md)
 - [UserGroupSearchItem](docs/UserGroupSearchItem.md)
 - [UserProfile](docs/UserProfile.md)
 - [UserResp](docs/UserResp.md)
 - [UserSearch](docs/UserSearch.md)
 - [UserSearchRespItem](docs/UserSearchRespItem.md)
 - [UserSysAdminFlag](docs/UserSysAdminFlag.md)
 - [VulnerabilitySummary](docs/VulnerabilitySummary.md)
 - [WebhookJob](docs/WebhookJob.md)
 - [WebhookLastTrigger](docs/WebhookLastTrigger.md)
 - [WebhookPolicy](docs/WebhookPolicy.md)
 - [WebhookTargetObject](docs/WebhookTargetObject.md)


## Documentation For Authorization

## basic
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author


