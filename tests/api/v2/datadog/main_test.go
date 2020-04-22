/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"os"
	"testing"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// TestMain starts the tracer.
func TestMain(m *testing.M) {
	tracer.Start(tracer.WithServiceName("datadog-api-client-go"))
	code := m.Run()
	tracer.Stop()
	os.Exit(code)
}
