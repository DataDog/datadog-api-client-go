// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

const stsWebIdentityResponse = `<AssumeRoleWithWebIdentityResponse xmlns="https://sts.amazonaws.com/doc/2011-06-15/">
  <AssumeRoleWithWebIdentityResult>
    <Credentials>
      <AccessKeyId>ASIAIOSFODNN7EXAMPLE</AccessKeyId>
      <SecretAccessKey>wJalrXUtnFEMI/K7MDENG/bPxRfiCYzEXAMPLEKEY</SecretAccessKey>
      <SessionToken>AQoXnyc4lcK4w4OIWen8ORnVZYIkFuyfAc/EXAMPLE</SessionToken>
      <Expiration>2026-06-09T23:00:00Z</Expiration>
    </Credentials>
  </AssumeRoleWithWebIdentityResult>
  <ResponseMetadata>
    <RequestId>ad4156e9-bce1-11e2-82e6-6b6efEXAMPLE</RequestId>
  </ResponseMetadata>
</AssumeRoleWithWebIdentityResponse>`

func TestGetCredentials_EnvVars(t *testing.T) {
	t.Setenv(AWSAccessKeyIdName, "AKIAIOSFODNN7EXAMPLE")
	t.Setenv(AWSSecretAccessKeyName, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYzEXAMPLEKEY")
	t.Setenv(AWSSessionTokenName, "token123")

	a := &AWSAuth{}
	creds := a.GetCredentials(context.Background())

	if creds.AccessKeyID != "AKIAIOSFODNN7EXAMPLE" {
		t.Errorf("unexpected AccessKeyID: %s", creds.AccessKeyID)
	}
	if creds.SessionToken != "token123" {
		t.Errorf("unexpected SessionToken: %s", creds.SessionToken)
	}
}

func TestGetCredentials_WebIdentity(t *testing.T) {
	// Spin up a mock STS server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if r.FormValue("Action") != "AssumeRoleWithWebIdentity" {
			http.Error(w, "unexpected action", http.StatusBadRequest)
			return
		}
		if r.FormValue("WebIdentityToken") != "test-oidc-token" {
			http.Error(w, "unexpected token", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprint(w, stsWebIdentityResponse)
	}))
	defer srv.Close()

	// Write a temp token file
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("test-oidc-token"), 0600); err != nil {
		t.Fatal(err)
	}

	t.Setenv(AWSWebIdentityTokenFileEnvVar, tokenFile)
	t.Setenv(AWSRoleARNEnvVar, "arn:aws:iam::123456789012:role/TestRole")
	// Clear static creds so we fall through to web identity
	t.Setenv(AWSAccessKeyIdName, "")
	t.Setenv(AWSSecretAccessKeyName, "")
	t.Setenv(AWSSessionTokenName, "")

	// Point the auth at our mock STS server
	a := &AWSAuth{stsEndpointOverride: srv.URL}
	creds := a.GetCredentials(context.Background())

	if creds.AccessKeyID != "ASIAIOSFODNN7EXAMPLE" {
		t.Errorf("unexpected AccessKeyID: %s", creds.AccessKeyID)
	}
	if creds.SessionToken != "AQoXnyc4lcK4w4OIWen8ORnVZYIkFuyfAc/EXAMPLE" {
		t.Errorf("unexpected SessionToken: %s", creds.SessionToken)
	}
}

func TestGetCredentials_WebIdentityMissingEnvVars(t *testing.T) {
	t.Setenv(AWSAccessKeyIdName, "")
	t.Setenv(AWSSecretAccessKeyName, "")
	t.Setenv(AWSSessionTokenName, "")
	t.Setenv(AWSWebIdentityTokenFileEnvVar, "")
	t.Setenv(AWSRoleARNEnvVar, "")

	a := &AWSAuth{}
	creds := a.GetCredentials(context.Background())

	// Should return empty credentials without panicking
	if creds == nil {
		t.Fatal("expected non-nil credentials")
	}
	if creds.SessionToken != "" {
		t.Errorf("expected empty session token, got: %s", creds.SessionToken)
	}
}

func TestGetCredentials_ContextOverridesWebIdentity(t *testing.T) {
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("should-not-be-used"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(AWSWebIdentityTokenFileEnvVar, tokenFile)
	t.Setenv(AWSRoleARNEnvVar, "arn:aws:iam::123456789012:role/TestRole")

	ctx := context.WithValue(context.Background(), ContextAWSVariables, map[string]string{
		AWSAccessKeyIdName:     "EXPLICIT_KEY",
		AWSSecretAccessKeyName: "EXPLICIT_SECRET",
		AWSSessionTokenName:    "EXPLICIT_TOKEN",
	})

	a := &AWSAuth{}
	creds := a.GetCredentials(ctx)

	if creds.AccessKeyID != "EXPLICIT_KEY" {
		t.Errorf("context credentials should take precedence, got: %s", creds.AccessKeyID)
	}
	if creds.SessionToken != "EXPLICIT_TOKEN" {
		t.Errorf("context credentials should take precedence, got: %s", creds.SessionToken)
	}
}
