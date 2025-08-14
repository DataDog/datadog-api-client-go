// Create a new Action Connection with HTTPMtlsAuth returns "Successfully created Action Connection" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CreateActionConnectionRequest{
		Data: datadogV2.ActionConnectionData{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributes{
				Name: "HTTP mTLS Connection exampleactionconnection",
				Integration: datadogV2.ActionConnectionIntegration{
					HTTPIntegration: &datadogV2.HTTPIntegration{
						Type:    datadogV2.HTTPINTEGRATIONTYPE_HTTP,
						BaseUrl: "https://api.example.com",
						Credentials: datadogV2.HTTPCredentials{
							HTTPMtlsAuth: &datadogV2.HTTPMtlsAuth{
								Type: datadogV2.HTTPMTLSAUTHTYPE_HTTPMTLSAUTH,
								Certificate: `-----BEGIN CERTIFICATE-----
MIICXjCCAUYCCQDOGcCfCHfhPzANBgkqhkiG9w0BAQsFADAzMQswCQYDVQQGEwJV
UzELMAkGA1UECAwCQ0ExFTATBgNVBAoMDERhdGFkb2cgSW5jLjAeFw0yNDA1MTQw
MDA1NThaFw0yNTA1MTQwMDA1NThaMDMxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJD
QTEVMBMGA1UECgwMRGF0YWRvZyBJbmMuMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAwQDTmLqWv2L6YhzBcjKgPEzd3kE+9dZ4hBXBCjK6HgF/3aOKOSYq
M9KPFHgJj6SjUJ+8TqX4sV6yW5xGX8dKjOpTYQfExEjGYcVrqKoOGg2k6dGkHyGm
2VzL4zKyK1C3zJ4KpJnMYK8dPPcgzJvO7jGxGKMgLVU3VNdxKGTrqKmC6RbZLQOz
M3fLp7bF2VcJ6VkGKW+yBK6vVMbQKMvjTbGqr3vIRd8SZzKRTsyIzXQDKgOv2vPn
SqYJjKFJ8vJ7JeH6zGyLjQ1cGVy9jJ3+TjJoJhCGOyGzJpBGIcXfYjFDLcSRh7KE
QIDAQABo1MwUTAdBgNVHQ4EFgQU/V8vJkPJ8b9yQnC/9bJ2kJGJ5MjoyEwHwYDVR0j
BBgwFoAU/V8vJkPJ8b9yQnC/9bJ2kJGJ5Mjo
-----END CERTIFICATE-----`,
								PrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDBaNOMV7V8T7gR
OmNcNfqGHxPrLLo1w2J7J8h6S8bVD9yCH2JKV8J5G2J8J0V8J3Jg8b3Jg8LxJgV
V8J6JgV8J9JgJg8LJg8VJgV5JgLLJgVVJg8V4Jg8VLJg8V6JgV8JqJgVV8J3JgV
V8J7JgVV8J5JgVV8J8JgVVJg8LJgJVLJgLVJgVVJgLJgVVJgVVJgVVJgLVJgVV
JgVVJgVVJgLJgVVJgLVJgLLJgVVJgVLJgVVJgVLJgVVJgVVJgVVJgVVJgVVJg
VVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJgVVJ
AgMBAAECggEBAKoJkJJJJkJJkJ
-----END PRIVATE KEY-----`,
							}},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.CreateActionConnection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.CreateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.CreateActionConnection`:\n%s\n", responseContent)
}
