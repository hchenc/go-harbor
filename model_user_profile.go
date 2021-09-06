/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UserProfile struct {
	Email string `json:"email,omitempty"`
	Realname string `json:"realname,omitempty"`
	Comment string `json:"comment,omitempty"`
}
