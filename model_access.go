/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Access struct {
	// The resource of the access
	Resource string `json:"resource,omitempty"`
	// The action of the access
	Action string `json:"action,omitempty"`
	// The effect of the access
	Effect string `json:"effect,omitempty"`
}
