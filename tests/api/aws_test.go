package api

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func TestGetCredentials(t *testing.T) {
	type testCase struct {
		name             string
		awsAccessKeyName string
		awsSecretKeyName string
		awsSessionToken  string
	}

	testCases := []testCase{
		{
			name:             "Valid AWS Credentials",
			awsAccessKeyName: "test",
			awsSecretKeyName: "test",
			awsSessionToken:  "test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variables
			t.Setenv(datadog.AWSAccessKeyName, tc.awsAccessKeyName)
			t.Setenv(datadog.AWSSecretKeyName, tc.awsSecretKeyName)
			t.Setenv(datadog.AWSSessionToken, tc.awsSessionToken)

			awsAuth := datadog.AWSAuth{}
			creds := awsAuth.GetCredentials()

			if creds.AccessKeyID != tc.awsAccessKeyName {
				t.Errorf("expected AccessKeyID %s, got %s", tc.awsAccessKeyName, creds.AccessKeyID)
			}
			if creds.SecretAccessKey != tc.awsSecretKeyName {
				t.Errorf("expected SecretAccessKey %s, got %s", tc.awsSecretKeyName, creds.SecretAccessKey)
			}
			if creds.SessionToken != tc.awsSessionToken {
				t.Errorf("expected SessionToken %s, got %s", tc.awsSessionToken, creds.SessionToken)
			}
		})
	}
}

func TestGenerateAWSAuthData(t *testing.T) {
	type testCase struct {
		name        string
		AWSAuth     datadog.AWSAuth
		creds       *datadog.Credentials
		orgUUID     string
		expectedErr bool
		expectedUrl string
	}

	testCases := []testCase{
		{
			name:    "Valid AWS Auth Data Generation",
			AWSAuth: datadog.AWSAuth{AwsRegion: "us-east-1"},
			creds: &datadog.Credentials{
				AccessKeyID:     "testAccessKey",
				SecretAccessKey: "testSecretKey",
				SessionToken:    "testSessionToken",
			},
			orgUUID:     "123456789012",
			expectedErr: false,
			expectedUrl: "https://sts.us-east-1.amazonaws.com",
		},
		{
			name:    "Invalid AWS Auth Data Generation with empty orgUUID",
			AWSAuth: datadog.AWSAuth{AwsRegion: "us-east-1"},
			creds: &datadog.Credentials{
				AccessKeyID:     "testAccessKey",
				SecretAccessKey: "testSecretKey",
				SessionToken:    "testSessionToken",
			},
			orgUUID:     "",
			expectedErr: true,
			expectedUrl: "https://sts.us-east-1.amazonaws.com",
		},
		{
			name:    "Valid AWS Auth Data Generation with empty region",
			AWSAuth: datadog.AWSAuth{},
			creds: &datadog.Credentials{
				AccessKeyID:     "testAccessKey",
				SecretAccessKey: "testSecretKey",
				SessionToken:    "testSessionToken",
			},
			orgUUID:     "123456789012",
			expectedErr: false,
			expectedUrl: "https://sts.amazonaws.com",
		},
		{
			name:        "Invalid AWS Auth Data Generation with nil credentials",
			AWSAuth:     datadog.AWSAuth{AwsRegion: "us-east-1"},
			creds:       nil,
			orgUUID:     "123456789012",
			expectedErr: true,
			expectedUrl: "https://sts.us-east-1.amazonaws.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := tc.AWSAuth.GenerateAwsAuthData(tc.orgUUID, tc.creds)
			if (err != nil) != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			if !tc.expectedErr && data == nil {
				t.Error("expected non-nil SigningData")
			}

			if (err != nil) && tc.expectedErr {
				// If we expect an error, we should not proceed further
				return
			}

			url, err := url.Parse(tc.expectedUrl)
			if err != nil {
				t.Errorf("failed to parse expected URL %s: %v", tc.expectedUrl, err)
			}

			if data.BodyEncoded == "" || data.HeadersEncoded == "" || data.Method == "" || data.URLEncoded == "" {
				t.Error("expected non-empty fields in SigningData")
			}

			urlDecoded, err := base64.StdEncoding.DecodeString(data.URLEncoded)
			if err != nil {
				t.Errorf("failed to decode URLEncoded: %v", err)
			}
			if string(urlDecoded) != tc.expectedUrl {
				t.Errorf("expected URL %s, got %s", tc.expectedUrl, urlDecoded)
			}

			headers := make(map[string][]string)
			headersDecoded, err := base64.StdEncoding.DecodeString(data.HeadersEncoded)
			if err != nil {
				t.Errorf("failed to decode HeadersEncoded: %v", err)
			}

			err = json.Unmarshal(headersDecoded, &headers)
			if err != nil {
				t.Errorf("failed to unmarshal HeadersEncoded: %v", err)
			}
			if len(headers) == 0 {
				t.Error("expected non-empty headers in SigningData")
			}

			if headers["x-ddog-org-id"][0] != tc.orgUUID {
				t.Errorf("expected org UUID %s in headers, got %s", tc.orgUUID, headers["x-ddog-org-id"][0])
			}
			if headers["host"][0] != url.Host {
				t.Errorf("expected host %s in headers, got %s", url.Host, headers["host"][0])

			}
			if headers["X-Amz-Security-Token"][0] != tc.creds.SessionToken {
				t.Errorf("expected SessionToken %s in headers, got %s", tc.creds.SessionToken, headers["X-Amz-Security-Token"][0])
			}
			if headers["X-Amz-Date"][0] == "" {
				t.Error("expected non-empty X-Amz-Date header in SigningData")
			}
			authorization := headers["Authorization"][0]
			if !strings.Contains(authorization, tc.AWSAuth.AwsRegion) {
				t.Errorf("expected region %s in Authorization header, got %s", tc.AWSAuth.AwsRegion, headers["Authorization"][0])
			}

			splitAuthorization := strings.Split(authorization, ", ")
			if len(splitAuthorization) < 3 {
				t.Error("expected at least 3 components in Authorization header")
			}

			if splitAuthorization[1] != "SignedHeaders=content-length;content-type;host;x-amz-date;x-amz-security-token;x-ddog-org-id" {
				t.Errorf("SignedHeaders not expected got %s", splitAuthorization[1])
			}
		})
	}
}
