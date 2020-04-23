/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	api "github.com/DataDog/datadog-api-client-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// IsRecording returns true if the recording mode is enabled
func IsRecording() bool {
	return os.Getenv("RECORD") == "true"
}

// Retry calls the call function for count times every interval while it returns false
func Retry(interval time.Duration, count int, call func() bool) error {
	for i := 0; i < count; i++ {
		if call() {
			return nil
		}
		if IsRecording() {
			time.Sleep(interval)
		}
	}
	return fmt.Errorf("Retry error: failed to satisfy the condition after %d times", count)
}

// ReadFixture opens the file at path and returns the contents as a string
func ReadFixture(path string) (string, error) {
	fixturePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get fixture file path: %v", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		return "", fmt.Errorf("failed to open fixture file: %v", err)
	}
	return string(data), nil
}

// ConfigureTracer starts the tracer.
func ConfigureTracer(m *testing.M) {
	tracer.Start(
		tracer.WithService("datadog-api-client-go"),
		tracer.WithServiceVersion(api.Version),
	)
	code := m.Run()

	tracer.Stop()
	os.Exit(code)
}
