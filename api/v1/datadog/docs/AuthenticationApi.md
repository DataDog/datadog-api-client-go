# AuthenticationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**Validate**](AuthenticationApi.md#Validate) | **Get** /api/v1/validate | Validate API key



## Validate

> AuthenticationValidationResponse Validate(ctx)

Check if the API key (not the APP key) is valid. If invalid, a 403 is returned.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.Validate(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Validate`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: AuthenticationValidationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuthenticationApi.Validate:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**AuthenticationValidationResponse**](AuthenticationValidationResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

