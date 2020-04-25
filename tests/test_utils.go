/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	api "github.com/DataDog/datadog-api-client-go"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonboulle/clockwork"
	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
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

// WithTestSpan starts new span with test information.
func WithTestSpan(ctx context.Context, t *testing.T) (context.Context, func()) {
	span, ctx := tracer.StartSpanFromContext(ctx, t.Name())
	return tracer.ContextWithSpan(ctx, span), func() {
		span.SetTag(ext.Error, t.Failed())
		span.Finish()
	}
}

func createWithDir(path string) (*os.File, error) {
	dirName := filepath.Dir(path)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

// SetClock stores current time in .freeze file.
func SetClock(t *testing.T) clockwork.FakeClock {
	os.MkdirAll("cassettes", 0755)

	f, err := createWithDir(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not set clock: %v", err)
	}
	defer f.Close()
	now := clockwork.NewRealClock().Now()
	f.WriteString(now.Format(time.RFC3339Nano))
	return clockwork.NewFakeClockAt(now)
}

// RestoreClock restore current time from .freeze file.
func RestoreClock(t *testing.T) clockwork.FakeClock {
	data, err := ioutil.ReadFile(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not load clock: %v", err)
	}
	now, err := time.Parse(time.RFC3339Nano, string(data))
	if err != nil {
		t.Fatalf("Could not parse clock date: %v", err)
	}
	return clockwork.NewFakeClockAt(now)
}

func removeURLSecrets(u *url.URL) string {
	q := u.Query()
	q.Del("api_key")
	q.Del("application_key")
	u.RawQuery = q.Encode()
	return u.String()
}

// MatchInteraction checks if the request matches a store request in the given cassette.
func MatchInteraction(r *http.Request, i cassette.Request) bool {
	return r.Method == i.Method && removeURLSecrets(r.URL) == i.URL
}

// FilterInteraction removes secret arguments from the URL.
func FilterInteraction(i *cassette.Interaction) error {
	u, err := url.Parse(i.URL)
	if err != nil {
		return err
	}
	i.URL = removeURLSecrets(u)
	i.Request.Headers.Del("Dd-Api-Key")
	i.Request.Headers.Del("Dd-Application-Key")
	return nil
}

// Recorder intercepts HTTP requests.
func Recorder(ctx context.Context, t *testing.T) (*recorder.Recorder, error) {
	// Configure recorder
	var mode recorder.Mode
	if IsRecording() {
		mode = recorder.ModeRecording
	} else {
		mode = recorder.ModeReplaying
	}

	r, err := recorder.NewAsMode(fmt.Sprintf("cassettes/%s", t.Name()), mode, nil)
	if err != nil {
		return nil, err
	}

	if span, ok := tracer.SpanFromContext(ctx); ok {
		span.SetTag("recorder.mode", mode)
	}

	r.SetMatcher(MatchInteraction)
	r.AddFilter(FilterInteraction)

	return r, nil
}

// WrapRoundTripper includes tracing information.
func WrapRoundTripper(rt http.RoundTripper, opts ...ddhttp.RoundTripperOption) http.RoundTripper {
	return ddhttp.WrapRoundTripper(rt, ddhttp.WithBefore(func(r *http.Request, span ddtrace.Span) {
		span.SetTag(ext.SpanName, r.Header.Get("DD-OPERATION-ID"))
	}))
}
