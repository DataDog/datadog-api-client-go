/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"net/http"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestGetAllProcesses(t *testing.T) {
	// Set up the client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	response, httpResponse, err := Client(ctx).ProcessesApi.GetAllProcesses(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all processes. Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(http.StatusOK, httpResponse.StatusCode)

	meta := response.GetMeta()
	data := response.GetData()

	assert.True(meta.GetPageSize() > 0)
	assert.NotEmpty(data)
}
