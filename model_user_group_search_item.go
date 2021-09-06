/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UserGroupSearchItem struct {
	// The ID of the user group
	Id int32 `json:"id,omitempty"`
	// The name of the user group
	GroupName string `json:"group_name,omitempty"`
	// The group type, 1 for LDAP group, 2 for HTTP group.
	GroupType int32 `json:"group_type,omitempty"`
}
