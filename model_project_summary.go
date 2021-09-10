/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProjectSummary struct {
	// The number of the repositories under this project.
	RepoCount int32 `json:"repo_count,omitempty"`
	// The total number of charts under this project.
	ChartCount int32 `json:"chart_count,omitempty"`
	// The total number of project admin members.
	ProjectAdminCount int32 `json:"project_admin_count,omitempty"`
	// The total number of maintainer members.
	MaintainerCount int32 `json:"maintainer_count,omitempty"`
	// The total number of developer members.
	DeveloperCount int32 `json:"developer_count,omitempty"`
	// The total number of guest members.
	GuestCount int32 `json:"guest_count,omitempty"`
	// The total number of limited guest members.
	LimitedGuestCount int32                `json:"limited_guest_count,omitempty"`
	Quota             *ProjectSummaryQuota `json:"quota,omitempty"`
	Registry          *Registry            `json:"registry,omitempty"`
}
