package swagger

import (
	"context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	ctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: "admin",
		Password: "Harbor12345",
	})
	config := NewConfigurationWithContext("http://harbor.hchenc.com:5088/api/v2.0", ctx)

	client := NewAPIClient(config)

	projects, resp, err := client.ProjectApi.ListProjects(&ProjectApiListProjectsOpts{})
	users, resp, err := client.UserApi.ListUsers(&UserApiListUsersOpts{})
	defer resp.Body.Close()
	fmt.Println(projects,users, resp, err)
}