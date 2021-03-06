/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LdapFailedImportUser struct {
	// the uid can't add to system.
	Uid string `json:"uid,omitempty"`
	// fail reason.
	Error_ string `json:"error,omitempty"`
}
