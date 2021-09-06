package swagger

import (
	"context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	config := NewConfiguration("http://harbor.hchenc.com:5088/api/v2.0")
	client := NewAPIClient(config)
	ctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: "admin",
		Password: "Harbor12345",
	})
	projects, resp, err := client.ProjectApi.ListProjects(ctx, &ProjectApiListProjectsOpts{})
	defer resp.Body.Close()
	fmt.Println(projects, resp, err)
}