package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	awsauth "github.com/DataDog/datadog-api-client-go/v2/auth/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

func awsCredentialsContext(ctx context.Context, key, secret, token string) context.Context {
	return context.WithValue(ctx, datadog.ContextAWSVariables, map[string]string{
		datadog.AWSAccessKeyIdName: key, datadog.AWSSecretAccessKeyName: secret, datadog.AWSSessionTokenName: token,
	})
}

func TestAWSProviderContextCredentialsOverrideProfileAndStaticCredentials(t *testing.T) {
	for _, constructorCredentials := range []bool{false, true} {
		t.Run(fmt.Sprintf("constructor_credentials=%t", constructorCredentials), func(t *testing.T) {
			isolateAWSAuthEnvironment(t)
			directory := t.TempDir()
			configFile, credentialsFile := filepath.Join(directory, "config"), filepath.Join(directory, "credentials")
			if err := os.WriteFile(configFile, []byte("[profile selected]\nregion = us-west-1\n"), 0o600); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(credentialsFile, []byte("[selected]\naws_access_key_id = profile-key\naws_secret_access_key = profile-secret\naws_session_token = profile-token\n"), 0o600); err != nil {
				t.Fatal(err)
			}
			t.Setenv("AWS_CONFIG_FILE", configFile)
			t.Setenv("AWS_SHARED_CREDENTIALS_FILE", credentialsFile)
			t.Setenv("AWS_PROFILE", "selected")
			var proof string
			server := newDelegatedTokenServer(t, func(r *http.Request) { proof = strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated ") })
			defer server.Close()
			var options []awsauth.Option
			if constructorCredentials {
				options = append(options, awsauth.WithStaticCredentials("constructor-key", "constructor-secret", "constructor-token"))
			}
			provider, err := awsauth.New(options...)
			if err != nil {
				t.Fatal(err)
			}
			// An absent context token must not be filled from either fallback source.
			for _, token := range []string{"", "context-token"} {
				ctx := awsCredentialsContext(delegatedTokenContext(server.URL), "context-key", "context-secret", token)
				if _, err := provider.Authenticate(ctx, delegatedTokenConfig()); err != nil {
					t.Fatal(err)
				}
				headers, endpoint := decodeProof(t, proof)
				if !strings.Contains(firstHeader(headers, "Authorization"), "Credential=context-key/") || firstHeader(headers, "X-Amz-Security-Token") != token {
					t.Fatal("context credentials did not take precedence as a complete credential set")
				}
				if endpoint != "https://sts.us-west-1.amazonaws.com" {
					t.Fatalf("context credentials lost profile region: %s", endpoint)
				}
			}
		})
	}
}

func TestAWSProviderContextEmptyCredentialsUseDefaultChain(t *testing.T) {
	for _, tc := range []struct {
		name string
		keys map[string]string
	}{
		{"nil map", nil},
		{"empty map", map[string]string{}},
		{"empty fields", map[string]string{datadog.AWSAccessKeyIdName: "", datadog.AWSSecretAccessKeyName: "", datadog.AWSSessionTokenName: ""}},
		{"unrelated field", map[string]string{"unrelated": "value"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			isolateAWSAuthEnvironment(t)
			t.Setenv("AWS_ACCESS_KEY_ID", "environment-key")
			t.Setenv("AWS_SECRET_ACCESS_KEY", "environment-secret")
			t.Setenv("AWS_SESSION_TOKEN", "environment-token")
			var proof string
			server := newDelegatedTokenServer(t, func(r *http.Request) { proof = strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated ") })
			defer server.Close()
			provider, err := awsauth.New()
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(delegatedTokenContext(server.URL), datadog.ContextAWSVariables, tc.keys)
			if _, err := provider.Authenticate(ctx, delegatedTokenConfig()); err != nil {
				t.Fatal(err)
			}
			headers, _ := decodeProof(t, proof)
			if !strings.Contains(firstHeader(headers, "Authorization"), "Credential=environment-key/") || firstHeader(headers, "X-Amz-Security-Token") != "environment-token" {
				t.Fatal("empty context credentials did not fall back to the SDK chain")
			}
		})
	}
}

func TestAWSProviderContextInvalidCredentialsDoNotFallBackOrExchange(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value interface{}
	}{
		{"access key only", map[string]string{datadog.AWSAccessKeyIdName: "sensitive-access-key"}},
		{"secret only", map[string]string{datadog.AWSSecretAccessKeyName: "sensitive-secret"}},
		{"session token only", map[string]string{datadog.AWSSessionTokenName: "sensitive-token"}},
		{"access key and token", map[string]string{datadog.AWSAccessKeyIdName: "sensitive-access-key", datadog.AWSSessionTokenName: "sensitive-token"}},
		{"secret and token", map[string]string{datadog.AWSSecretAccessKeyName: "sensitive-secret", datadog.AWSSessionTokenName: "sensitive-token"}},
		{"wrong map type", map[string]interface{}{datadog.AWSAccessKeyIdName: "sensitive-access-key"}},
		{"wrong value type", "sensitive-value"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			isolateAWSAuthEnvironment(t)
			var retrievals, exchanges atomic.Int32
			server := newDelegatedTokenServer(t, func(*http.Request) { exchanges.Add(1) })
			defer server.Close()
			provider, err := awsauth.New(awsauth.WithConfigOptions(awsconfig.WithCredentialsProvider(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
				retrievals.Add(1)
				return aws.Credentials{AccessKeyID: "fallback-key", SecretAccessKey: "fallback-secret"}, nil
			}))))
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(delegatedTokenContext(server.URL), datadog.ContextAWSVariables, tc.value)
			credentials, err := provider.Authenticate(ctx, delegatedTokenConfig())
			if err == nil || credentials != nil {
				t.Fatal("invalid context credentials were accepted")
			}
			if strings.Contains(err.Error(), "sensitive-") {
				t.Fatal("error leaked credential material")
			}
			if retrievals.Load() != 0 || exchanges.Load() != 0 {
				t.Fatal("invalid context credentials reached the fallback provider or token exchange")
			}
		})
	}
}

func TestAWSProviderContextCredentialsAreNotCachedAcrossAuthentications(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	var retrievals atomic.Int32
	cache := aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
		retrievals.Add(1)
		return aws.Credentials{AccessKeyID: "chain-key", SecretAccessKey: "chain-secret", SessionToken: "chain-token"}, nil
	}))
	var proof string
	server := newDelegatedTokenServer(t, func(r *http.Request) { proof = strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated ") })
	defer server.Close()
	provider, err := awsauth.New(awsauth.WithConfigOptions(awsconfig.WithCredentialsProvider(cache)))
	if err != nil {
		t.Fatal(err)
	}
	base := delegatedTokenContext(server.URL)
	for _, tc := range []struct {
		ctx        context.Context
		key, token string
		retrievals int32
	}{
		{awsCredentialsContext(base, "first-key", "first-secret", "first-token"), "first-key", "first-token", 0},
		{base, "chain-key", "chain-token", 1},
		{awsCredentialsContext(base, "second-key", "second-secret", ""), "second-key", "", 1},
		{awsCredentialsContext(base, "", "", ""), "chain-key", "chain-token", 1},
	} {
		if _, err := provider.Authenticate(tc.ctx, delegatedTokenConfig()); err != nil {
			t.Fatal(err)
		}
		headers, _ := decodeProof(t, proof)
		if !strings.Contains(firstHeader(headers, "Authorization"), "Credential="+tc.key+"/") || firstHeader(headers, "X-Amz-Security-Token") != tc.token {
			t.Fatalf("authentication did not use current context credentials: want key %s", tc.key)
		}
		if got := retrievals.Load(); got != tc.retrievals {
			t.Fatalf("chain retrieval count = %d, want %d", got, tc.retrievals)
		}
	}
}

func TestAWSProviderContextCredentialsSkipRetrievalAndPreserveHTTPClient(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	var exchanges atomic.Int32
	client := &http.Client{Transport: awsAuthRoundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.String() != "https://token.invalid/api/v2/delegated-token" {
			return nil, fmt.Errorf("unexpected request: %s", r.URL)
		}
		exchanges.Add(1)
		headers, endpoint := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
		if !strings.Contains(firstHeader(headers, "Authorization"), "Credential=context-key/") || endpoint != "https://sts.eu-west-1.amazonaws.com" {
			t.Error("explicit context credentials or region were lost")
		}
		return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"data":{"attributes":{"access_token":"test-token"}}}`)), Request: r}, nil
	})}
	provider, err := awsauth.New(awsauth.WithRegion("eu-west-1"), awsauth.WithHTTPClient(client),
		awsauth.WithConfigOptions(awsconfig.WithCredentialsProvider(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			t.Error("explicit context credentials must not retrieve fallback credentials")
			return aws.Credentials{}, errors.New("fallback credentials unavailable")
		}))))
	if err != nil {
		t.Fatal(err)
	}
	ctx := awsCredentialsContext(delegatedTokenContext("https://token.invalid"), "context-key", "context-secret", "")
	if _, err := provider.Authenticate(ctx, delegatedTokenConfig()); err != nil {
		t.Fatal(err)
	}
	if exchanges.Load() != 1 {
		t.Fatal("token exchange did not use the configured HTTP client")
	}
}

func TestAWSProviderContextCredentialsPreserveConstructorFallback(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	var proof string
	server := newDelegatedTokenServer(t, func(r *http.Request) { proof = strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated ") })
	defer server.Close()
	provider, err := awsauth.New(awsauth.WithStaticCredentials("constructor-key", "constructor-secret", "constructor-token"))
	if err != nil {
		t.Fatal(err)
	}
	base := delegatedTokenContext(server.URL)
	for _, tc := range []struct {
		ctx        context.Context
		key, token string
	}{
		{awsCredentialsContext(base, "first-key", "first-secret", ""), "first-key", ""},
		{base, "constructor-key", "constructor-token"},
		{awsCredentialsContext(base, "second-key", "second-secret", "second-token"), "second-key", "second-token"},
		{awsCredentialsContext(base, "", "", ""), "constructor-key", "constructor-token"},
	} {
		if _, err := provider.Authenticate(tc.ctx, delegatedTokenConfig()); err != nil {
			t.Fatal(err)
		}
		headers, _ := decodeProof(t, proof)
		if !strings.Contains(firstHeader(headers, "Authorization"), "Credential="+tc.key+"/") || firstHeader(headers, "X-Amz-Security-Token") != tc.token {
			t.Fatalf("context credentials changed the constructor fallback: want key %s", tc.key)
		}
	}
}

func TestAWSProviderContextCredentialsRefreshThroughAPIClient(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	t.Setenv("AWS_ACCESS_KEY_ID", "chain-key")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "chain-secret")
	t.Setenv("AWS_SESSION_TOKEN", "chain-token")
	var exchanges, apiCalls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v2/delegated-token":
			n := exchanges.Add(1)
			wantKey, wantToken := "context-key", ""
			if n == 2 {
				wantKey, wantToken = "chain-key", "chain-token"
			}
			headers, _ := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
			if !strings.Contains(firstHeader(headers, "Authorization"), "Credential="+wantKey+"/") || firstHeader(headers, "X-Amz-Security-Token") != wantToken {
				t.Error("API client renewal did not use the current auth context")
			}
			fmt.Fprintf(w, `{"data":{"attributes":{"access_token":"token-%d","expires":"%d"}}}`, n, time.Now().Add(time.Hour).Unix())
		case "/api/v2/current_user":
			apiCalls.Add(1)
			if r.Header.Get("Authorization") != fmt.Sprintf("Bearer token-%d", exchanges.Load()) {
				t.Error("API request did not use the current delegated token")
			}
			io.WriteString(w, `{"data":{"id":"test-user","type":"users","attributes":{"name":"Test User"}}}`)
		default:
			t.Errorf("unexpected request path %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	provider, err := awsauth.New()
	if err != nil {
		t.Fatal(err)
	}
	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = delegatedTokenConfig()
	config.DelegatedTokenConfig.ProviderAuth = provider
	credentials := &datadog.DelegatedTokenCredentials{}
	auth := awsCredentialsContext(delegatedTokenContext(server.URL), "context-key", "context-secret", "")
	auth = context.WithValue(auth, datadog.ContextDelegatedToken, credentials)
	users := datadogV2.NewUsersApi(datadog.NewAPIClient(config))
	if _, _, err := users.GetCurrentUser(auth); err != nil {
		t.Fatal(err)
	}
	// AWS context changes are consumed on renewal, not while a valid Datadog
	// token remains cached in the standard ContextDelegatedToken entry.
	auth = awsCredentialsContext(auth, "", "", "")
	if _, _, err := users.GetCurrentUser(auth); err != nil {
		t.Fatal(err)
	}
	if exchanges.Load() != 1 {
		t.Fatal("unexpired delegated token was not reused")
	}
	credentials.Expiration = time.Now().Add(-time.Second)
	if _, _, err := users.GetCurrentUser(auth); err != nil {
		t.Fatal(err)
	}
	if exchanges.Load() != 2 || apiCalls.Load() != 3 {
		t.Fatalf("exchanges=%d calls=%d", exchanges.Load(), apiCalls.Load())
	}
}

func TestAWSProviderContextConcurrentAuthentications(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	var retrievals, exchanges atomic.Int32
	cache := aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
		retrievals.Add(1)
		return aws.Credentials{AccessKeyID: "chain-key", SecretAccessKey: "chain-secret", SessionToken: "chain-token"}, nil
	}))
	server := newDelegatedTokenServer(t, func(r *http.Request) {
		exchanges.Add(1)
		headers, _ := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
		identity := firstHeader(headers, orgIDHeader)
		if !strings.Contains(firstHeader(headers, "Authorization"), "Credential="+identity+"-key/") || firstHeader(headers, "X-Amz-Security-Token") != identity+"-token" {
			t.Errorf("concurrent authentication mixed credentials for %q", identity)
		}
	})
	defer server.Close()
	provider, err := awsauth.New(awsauth.WithConfigOptions(awsconfig.WithCredentialsProvider(cache)))
	if err != nil {
		t.Fatal(err)
	}
	base := delegatedTokenContext(server.URL)
	// NewDefaultContext populates an empty ContextAWSVariables map when AWS
	// environment credentials are absent. It must still allow the SDK chain.
	base = datadog.NewDefaultContext(base)
	base = context.WithValue(base, datadog.ContextServerVariables, delegatedTokenContext(server.URL).Value(datadog.ContextServerVariables))
	var workers sync.WaitGroup
	start := make(chan struct{})
	const count = 16
	for i := 0; i < count; i++ {
		ctx, identity := base, "chain"
		if i%2 == 0 {
			identity = fmt.Sprintf("context-%d", i)
			ctx = awsCredentialsContext(base, identity+"-key", identity+"-secret", identity+"-token")
		}
		workers.Add(1)
		go func(ctx context.Context, identity string) {
			defer workers.Done()
			<-start
			if _, err := provider.Authenticate(ctx, &datadog.DelegatedTokenConfig{OrgUUID: identity, Provider: datadog.ProviderAWS}); err != nil {
				t.Errorf("authenticating %s: %v", identity, err)
			}
		}(ctx, identity)
	}
	close(start)
	workers.Wait()
	if retrievals.Load() != 1 || exchanges.Load() != count {
		t.Fatalf("retrievals=%d exchanges=%d", retrievals.Load(), exchanges.Load())
	}
}
