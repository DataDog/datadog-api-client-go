package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestPermissionList(t *testing.T) {
	c := NewClientWithRecording(t)
	defer c.Close()

	psr, httpresp, err := c.Client.RolesApi.ListPermissions(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, psr.HasData())
	assert.NotEmpty(t, psr.GetData())
}
