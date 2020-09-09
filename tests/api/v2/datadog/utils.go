/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/publicsuffix"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

// WithFakeAuth avoids issue of API returning `text/html` instead of `application/json`
func WithFakeAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: "FAKE_KEY",
			},
			"appKeyAuth": {
				Key: "FAKE_KEY",
			},
		},
	)
}

// WithTestAuth returns authenticated context.
func WithTestAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
}

// NewConfiguration return configuration with known options.
func NewConfiguration() *datadog.Configuration {
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"

	// Set test site as default
	testSite, ok := os.LookupEnv("DD_TEST_SITE")
	if ok {
		server := config.Servers[0]
		site := server.Variables["site"]
		site.DefaultValue = testSite
		site.EnumValues = append(site.EnumValues, testSite)
		server.Variables["site"] = site
		config.Servers[0] = server

		for operationID, servers := range config.OperationServers {
			server := servers[0]
			site := server.Variables["site"]
			site.DefaultValue = testSite
			site.EnumValues = append(site.EnumValues, testSite)
			server.Variables["site"] = site
			servers[0] = server
			config.OperationServers[operationID] = servers
		}
	}
	return config
}

type contextKey string

var (
	clientKey = contextKey("client")
)

// WithClient sets client for unit tests in context.
func WithClient(ctx context.Context, t *testing.T) (context.Context, func()) {
	ctx, finish := tests.WithTestSpan(ctx, t)
	ctx = context.WithValue(ctx, clientKey, datadog.NewAPIClient(NewConfiguration()))
	return ctx, finish
}

// ClientFromContext returns client and indication if it was successful.
func ClientFromContext(ctx context.Context) (*datadog.APIClient, bool) {
	if ctx == nil {
		return nil, false
	}
	v := ctx.Value(clientKey)
	if c, ok := v.(*datadog.APIClient); ok {
		return c, true
	}
	return nil, false
}

// Client returns client from context.
func Client(ctx context.Context) *datadog.APIClient {
	c, ok := ClientFromContext(ctx)
	if !ok {
		log.Fatal("client is not configured")
	}
	return c
}

// WithRecorder configures client with recorder.
func WithRecorder(ctx context.Context, t *testing.T) (context.Context, func()) {
	ctx, finish := WithClient(ctx, t)
	client := Client(ctx)

	ctx, err := tests.WithClock(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		log.Fatal(err)
	}

	r, err := tests.Recorder(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		log.Fatal(err)
	}
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	return ctx, func() {
		r.Stop()
		finish()
	}
}

func GetTestDomain(ctx context.Context, client *datadog.APIClient) (string, error) {
	baseUrl, err := client.GetConfig().ServerURLWithContext(ctx, "")
	if err != nil {
		return "", fmt.Errorf("could not generate base url: %v", err)
	}

	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return "", fmt.Errorf("could not parse base url: %v", err)
	}

	host, err := publicsuffix.EffectiveTLDPlusOne(parsedUrl.Host)
	if err != nil {
		return "", fmt.Errorf("could not parse TLD+1: %v", err)
	}
	return host, nil
}

// SendRequest sends request to endpoints without specification.
func SendRequest(ctx context.Context, method, url string, payload []byte) (*http.Response, []byte, error) {
	baseURL := ""
	if !strings.HasPrefix(url, "https://") {
		var err error
		baseURL, err = Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
		if err != nil {
			return nil, []byte{}, fmt.Errorf("Failed to get base URL for Datadog API: %s", err.Error())
		}
	}

	request, err := http.NewRequest(method, baseURL+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Failed to create request for Datadog API: %s", err.Error())
	}
	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	request.Header.Add("DD-API-KEY", keys["apiKeyAuth"].Key)
	request.Header.Add("DD-APPLICATION-KEY", keys["appKeyAuth"].Key)
	request.Header.Set("Content-Type", "application/json")

	resp, respErr := Client(ctx).GetConfig().HTTPClient.Do(request)
	body, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		respErr = fmt.Errorf("Failed reading response body: %s", rerr.Error())
	}
	return resp, body, respErr
}

func setIncidentsUnstableOperations(client *datadog.APIClient, enable bool) {
	client.GetConfig().SetUnstableOperationEnabled("CreateUserDefinedField", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetUserDefinedFields", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchUserDefinedField", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteUserDefinedField", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetUserDefinedField", enable)
	client.GetConfig().SetUnstableOperationEnabled("CreateUserDefinedFieldChoice", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetUserDefinedFieldChoices", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchUserDefinedFieldChoice", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteUserDefinedFieldChoice", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetUserDefinedFieldChoice", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidents", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("SearchIncidents", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateIncidentPostmortem", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentPostmortems", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchIncidentPostmortem", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteIncidentPostmortem", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentPostmortem", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateUserDefinedFieldSelection", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetUserDefinedFieldSelections", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchUserDefinedFieldSelection", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteUserDefinedFieldSelection", enable)

	client.GetConfig().SetUnstableOperationEnabled("AddServiceToIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetServicesForIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("RemoveServiceFromIncident", enable)

	client.GetConfig().SetUnstableOperationEnabled("AddTeamToIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetTeamsForIncident", enable)
	client.GetConfig().SetUnstableOperationEnabled("RemoveTeamFromIncident", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateIncidentTimelineCell", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentTimelineCells", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchIncidentTimelineCell", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteIncidentTimelineCell", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentTimelineCell", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateIncidentTodo", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentTodos", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchIncidentTodo", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteIncidentTodo", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetIncidentTodo", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateTeam", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetTeams", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchTeam", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteTeam", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetTeam", enable)

	client.GetConfig().SetUnstableOperationEnabled("CreateService", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetServices", enable)
	client.GetConfig().SetUnstableOperationEnabled("PatchService", enable)
	client.GetConfig().SetUnstableOperationEnabled("DeleteService", enable)
	client.GetConfig().SetUnstableOperationEnabled("GetService", enable)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string drawn from the provided character set
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString generates a random string from the set of alphabetical characters
func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
