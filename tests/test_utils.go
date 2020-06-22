/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	api "github.com/DataDog/datadog-api-client-go"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// RecordingMode defines valid usage of cassette recorder
type RecordingMode string

// Valid usage modes for cassette recorder
const (
	ModeIgnore    RecordingMode = "none"
	ModeReplaying RecordingMode = "false"
	ModeRecording RecordingMode = "true"
)

var testFiles2EndpointTags = map[string]map[string]string{
	"tests/api/v1/datadog": {
		"api_authentication_test": "validation",
		"api_aws_integration_test": "integration-aws",
		"api_aws_logs_integration_test": "integration-aws",
		"api_azure_integration_test": "integration-azure",
		"api_dashboard_lists_test": "dashboard-lists",
		"api_dashboards_test": "dashboards",
		"api_downtimes_test": "downtimes",
		"api_events_test": "events",
		"api_gcp_integration_test": "integration-gcp",
		"api_hosts_test": "hosts",
		"api_ip_ranges_test": "ip-ranges",
		"api_key_management_test": "key-management",
		"api_logs_indexes_test": "logs-indexes",
		"api_logs_pipelines_test": "logs-pipelines",
		"api_logs_test": "logs",
		"api_metrics_test": "metrics",
		"api_monitors_test": "monitors",
		"api_organizations_test": "organizations",
		"api_pager_duty_integration_test": "integration-pagerduty",
		"api_service_level_objectives_test": "service-level-objectives",
		"api_snapshots_test": "snapshots",
		"api_synthetics_test": "synthetics",
		"api_tags_test": "tags",
		"api_usage_metering_test": "usage-metering",
		"api_users_test": "users",
		"telemetry_test": "telemetry",
	},
	"tests/api/v2/datadog": {
		"api_dashboard_lists_test": "dashboard-lists",
		"api_logs_archives_test": "logs-archives",
		"api_logs_test": "logs",
		"api_permissions_test": "permissions",
		"api_roles_test": "roles",
		"api_users_test": "users",
		"security_monitoring_test": "security-monitoring",
		"telemetry_test": "telemetry",
	},
}

// GetRecording returns the value of RECORD environment variable
func GetRecording() RecordingMode {
	if value, exists := os.LookupEnv("RECORD"); exists {
		switch value {
		case string(ModeIgnore):
			return ModeIgnore
		case string(ModeRecording):
			return ModeRecording
		default:
			return ModeReplaying
		}
	} else {
		return ModeReplaying
	}
}

// IsCIRun returns true if the CI environment variable is set to "true"
func IsCIRun() bool {
	return os.Getenv("CI") == "true"
}

// Retry calls the call function for count times every interval while it returns false
func Retry(interval time.Duration, count int, call func() bool) error {
	for i := 0; i < count; i++ {
		if call() {
			return nil
		}
		if GetRecording() != ModeReplaying {
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
	if GetRecording() == ModeReplaying {
		os.Exit(m.Run())
	}
	service, ok := os.LookupEnv("DD_SERVICE")
	if !ok {
		service = "datadog-api-client-go"
	}
	tracer.Start(
		tracer.WithService(service),
		tracer.WithServiceVersion(api.Version),
	)
	code := m.Run()
	tracer.Stop()
	os.Exit(code)
}

// getEndpointTagValue traverses callstack frames to find the test function that invoked this call;
// it then matches the file defining this function against testFiles2EndpointTags to figure out
// the tag value to set on span
func getEndpointTagValue(t *testing.T) (string, error) {
	var pcs [512]uintptr
	var frame runtime.Frame
	more := true
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	functionFile := ""
	for more {
		frame, more = frames.Next()
		// nested test functions like `TestAuthenticationValidate/200_Valid` will have frame.Function ending with
		// ".funcX", `e.g. datadog.TestAuthenticationValidate.func1`, so trim everything after last "/" in test name
		// and everything after last "." in frame function name
		frameFunction := frame.Function
		testName := t.Name()
		if strings.Contains(testName, "/") {
			testName = testName[:strings.LastIndex(testName, "/")]
			frameFunction = frameFunction[:strings.LastIndex(frameFunction, ".")]
		}
		if strings.HasSuffix(frameFunction, "." + testName) {
			functionFile = frame.File
			// when we find the frame with the current test function, match it against testFiles2EndpointTags
			for subdir, file2tag := range testFiles2EndpointTags {
				for file, tag := range file2tag {
					if strings.HasSuffix(frame.File, fmt.Sprintf("%s/%s.go", subdir, file)) {
						return tag, nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf(
		"Endpoint tag for test file %s not found in tests/test_utils.go, please add it to `testFiles2EndpointTags`",
		functionFile)
}

// WithTestSpan starts new span with test information.
func WithTestSpan(ctx context.Context, t *testing.T) (context.Context, func()) {
	t.Helper()
	tag, err := getEndpointTagValue(t)
	if err != nil {
		panic(err.Error())
	}
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"test",
		tracer.SpanType("test"),
		tracer.ResourceName(t.Name()),
		tracer.Tag(ext.AnalyticsEvent, true),
		tracer.Measured(),
	)
	// We need to make the tag be something that is then searchable in monitors
	// https://docs.datadoghq.com/tracing/guide/metrics_namespace/#errors
	// "version" is really the only one we can use here
	// NOTE: version is treated in slightly different way, because it's a special tag;
	// if we set it in StartSpanFromContext, it would get overwritten
	span.SetTag("version", tag)
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
	t.Helper()
	now := clockwork.NewRealClock().Now()
	if GetRecording() == ModeRecording {
		os.MkdirAll("cassettes", 0755)

		f, err := createWithDir(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
		if err != nil {
			t.Fatalf("Could not set clock: %v", err)
		}
		defer f.Close()
		f.WriteString(now.Format(time.RFC3339Nano))
	}
	return clockwork.NewFakeClockAt(now)
}

// RestoreClock restore current time from .freeze file.
func RestoreClock(t *testing.T) clockwork.FakeClock {
	t.Helper()
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

type contextKey string

var (
	clockKey = contextKey("clock")
)

// WithClock sets clock to context.
func WithClock(ctx context.Context, t *testing.T) context.Context {
	t.Helper()
	var fc clockwork.FakeClock
	if GetRecording() != ModeReplaying {
		fc = SetClock(t)
	} else {
		fc = RestoreClock(t)
	}
	return context.WithValue(ctx, clockKey, fc)
}

// UniqueEntityName will return a unique string that can be used as a title/description/summary/...
// of an API entity. When used in Azure Pipelines and RECORD=true or RECORD=none, it will include
// BuildId to enable mapping resources that weren't deleted to builds.
func UniqueEntityName(ctx context.Context, t *testing.T) *string {
	buildID, present := os.LookupEnv("BUILD_BUILDID")
	if !present || !IsCIRun() || GetRecording() == ModeReplaying {
		buildID = "local"
	}

	// NOTE: some endpoints have limits on certain fields (e.g. Roles V2 names can only be 55 chars long),
	// so we need to keep this short
	result := fmt.Sprintf("go-%s-%s-%d", t.Name(), buildID, ClockFromContext(ctx).Now().Unix())
	// In case this is used in URL, make sure we replace the slash that is added by subtests
	result = strings.ReplaceAll(result, "/", "-")
	return &result
}

// ClockFromContext returns clock or panics.
func ClockFromContext(ctx context.Context) clockwork.FakeClock {
	if ctx == nil {
		log.Fatal("ctx is required")
	}
	v := ctx.Value(clockKey)
	fc, ok := v.(clockwork.FakeClock)
	if !ok {
		log.Fatalf("invalid value %v should be clockwork.FakeClock{}", v)
	}
	return fc
}

func removeURLSecrets(u *url.URL) string {
	q := u.Query()
	q.Del("api_key")
	q.Del("application_key")
	u.RawQuery = q.Encode()
	site, ok := os.LookupEnv("DD_TEST_SITE")
	if ok {
		fmt.Println("!!! WARN: Replacing DD_TEST_SITE by 'datadoghq.com' for request matching")
		u.Host = strings.Replace(u.Host, site, "datadoghq.com", 1)
	}
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
	// Remove custom URL in report-uri from response
	i.Response.Headers.Del("Content-Security-Policy")
	return nil
}

// Recorder intercepts HTTP requests.
func Recorder(ctx context.Context, t *testing.T) (*recorder.Recorder, error) {
	t.Helper()
	// Configure recorder
	var mode recorder.Mode
	switch GetRecording() {
	case ModeReplaying:
		mode = recorder.ModeReplaying
	case ModeIgnore:
		mode = recorder.ModeDisabled
	default:
		mode = recorder.ModeRecording
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
	if GetRecording() == ModeReplaying {
		return rt
	}
	return ddhttp.WrapRoundTripper(
		rt,
		ddhttp.WithBefore(func(r *http.Request, span ddtrace.Span) {
			span.SetTag(ext.SpanName, r.Header.Get("DD-OPERATION-ID"))
		}),
		ddhttp.WithAfter(func(r *http.Response, span ddtrace.Span) {
			if 500 <= r.StatusCode && r.StatusCode < 600 {
				bodyBytes, _ := ioutil.ReadAll(r.Body)
				r.Body.Close() // must close
				r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
				span.SetTag(ext.Error, true)
				span.SetTag(ext.ErrorMsg, string(bodyBytes))
				span.SetTag(ext.ErrorDetails, r.Status)
			}
		}),
	)
}

// Assertions wrapper
type Assertions struct {
	require.Assertions
}

// TestingT keeps track of a context.
type TestingT struct {
	*testing.T
	ctx context.Context
}

// Errorf format error message.
func (t *TestingT) Errorf(format string, args ...interface{}) {
	t.T.Helper()
	span, _ := tracer.StartSpanFromContext(t.ctx, "testing.error")
	defer span.Finish()
	span.SetTag(ext.Error, fmt.Errorf(format, args...))
	t.T.Errorf(format, args...)
}

// Assert wraps context and testing object.
func Assert(ctx context.Context, t *testing.T) *Assertions {
	t.Helper()
	return &Assertions{*require.New(&TestingT{t, ctx})}
}
