/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

// TestMain starts the tracer.
func TestMain(m *testing.M) {
	stopServer, err := startGeneratedTestServer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if testServerEnabled() {
		fmt.Println("=== Using Generated Test Runner ===")
	}
	code := tests.RunWithTracer(m)
	stopServer()
	os.Exit(code)
}

func startGeneratedTestServer() (func(), error) {
	if !generatedTestsEnabled() {
		return func() {}, nil
	}
	if os.Getenv("RECORD") != "" && os.Getenv("RECORD") != "false" {
		return func() {}, nil
	}
	root := "generated-test"
	server := filepath.Join(root, "test-server")
	if _, err := os.Stat(server); err != nil {
		return func() {}, nil
	}
	if os.Getenv("DD_TEST_RUNNER_DATA") == "" {
		if err := os.Setenv("DD_TEST_RUNNER_DATA", filepath.Join(root, "test-runner-data")); err != nil {
			return nil, err
		}
	}
	if os.Getenv("DD_TEST_SERVER_URL") != "" {
		return func() {}, nil
	}

	port := os.Getenv("DD_TEST_SERVER_PORT")
	if port == "" {
		port = "18087"
	}
	serverURL := "http://127.0.0.1:" + port
	if err := os.Setenv("DD_TEST_SERVER_URL", serverURL); err != nil {
		return nil, err
	}
	logPath := os.Getenv("DD_TEST_SERVER_LOG")
	if logPath == "" {
		logPath = filepath.Join(os.TempDir(), "datadog-go-test-server.log")
	}
	logFile, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	command := exec.Command(server, "--port", port)
	command.Stdout = logFile
	command.Stderr = logFile
	if err := command.Start(); err != nil {
		logFile.Close()
		return nil, err
	}
	done := make(chan error, 1)
	go func() { done <- command.Wait() }()
	stop := func() {
		if command.Process != nil {
			_ = command.Process.Kill()
		}
		select {
		case <-done:
		case <-time.After(5 * time.Second):
		}
		_ = logFile.Close()
	}

	client := &http.Client{Timeout: time.Second}
	health := serverURL + "/__openapi_transformer__/health"
	for attempt := 0; attempt < 50; attempt++ {
		response, requestErr := client.Get(health)
		if requestErr == nil {
			response.Body.Close()
			if response.StatusCode >= 200 && response.StatusCode < 300 {
				return stop, nil
			}
		}
		select {
		case serverErr := <-done:
			logFile.Close()
			return nil, fmt.Errorf("generated test server exited early (%v); see %s", serverErr, logPath)
		default:
		}
		time.Sleep(100 * time.Millisecond)
	}
	stop()
	return nil, fmt.Errorf("generated test server failed to start; see %s", logPath)
}
