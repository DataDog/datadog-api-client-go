package test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestPermissionList(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	permissions, httpresp, err := TestAPIClient.RolesApi.ListPermissions(TestAuth).Execute()
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	assert.Assert(t, permissions.HasData())
	assert.Assert(t, len(permissions.GetData()) > 0)
}
