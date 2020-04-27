/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"
)

// TestMain starts the tracer.
func TestMain(m *testing.M) {
	tests.ConfigureTracer(m)
}
