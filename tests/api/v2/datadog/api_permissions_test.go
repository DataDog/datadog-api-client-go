package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestPermissionList(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	psr, httpresp, err := TestAPIClient.RolesApi.ListPermissions(TestAuth).Execute()
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, psr.HasData())
	assert.NotEmpty(t, psr.GetData())
}
