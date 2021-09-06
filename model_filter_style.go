/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The style of the resource filter
type FilterStyle struct {
	// The filter type
	Type_ string `json:"type,omitempty"`
	// The filter style
	Style string `json:"style,omitempty"`
	// The filter values
	Values []string `json:"values,omitempty"`
}