# \GCPIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGCPIntegration**](GCPIntegrationApi.md#CreateGCPIntegration) | **Post** /api/v1/integration/gcp | Add a GCP integration to your Datadog account.
[**DeleteGCPIntegration**](GCPIntegrationApi.md#DeleteGCPIntegration) | **Delete** /api/v1/integration/gcp | Delete a GCP Integration from your Datadog account.
[**ListGCPIntegration**](GCPIntegrationApi.md#ListGCPIntegration) | **Get** /api/v1/integration/gcp | List configured GCP integrations.
[**UpdateGCPIntegration**](GCPIntegrationApi.md#UpdateGCPIntegration) | **Put** /api/v1/integration/gcp | Update a GCP integration in your Datadog account.



## CreateGCPIntegration

> interface{} CreateGCPIntegration(ctx, gcpAccount)

Add a GCP integration to your Datadog account.

### Overview
Create a Datadog-GCP integration.
### Arguments
* **`type`** [*required*, *default* = **None**]: The value for service_account found in your JSON service account key.
* **`project_id`** [*required*, *default* = **None**]: Your Google Cloud project ID found in your JSON service account key.
* **`private_key_id`** [*required*, *default* = **None**]: Your private key ID found in your JSON service account key.
* **`private_key`** [*required*, *default* = **None**]: Your private key name found in your JSON service account key.
* **`client_email`** [*required*, *default* = **None**]: Your email found in your JSON service account key.
* **`client_id`** [*required*, *default* = **None**]: Your ID found in your JSON service account key.
* **`auth_uri`** [*required*, *default* = **None**]: Should be https://accounts.google.com/o/oauth2/auth.
* **`token_uri`** [*required*, *default* = **None**]: Should be https://accounts.google.com/o/oauth2/token.
* **`auth_provider_x509_cert_url`** [*required*, *default* = **None**]: Should be https://www.googleapis.com/oauth2/v1/certs.
* **`client_x509_cert_url`** [*required*, *default* = **None**]: Should be https://www.googleapis.com/robot/v1/metadata/x509/<CLIENT_EMAIL> where <CLIENT_EMAIL> is the email found in your JSON service account key.
* **`host_filters`** [*optional*, *default* = **None**]: Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.
* **`automute`** [*optional*, *default* = false]: Silence monitors for expected GCE instance shutdowns.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gcpAccount** | [**GcpAccount**](GcpAccount.md)| Create a Datadog-Azure integration. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGCPIntegration

> interface{} DeleteGCPIntegration(ctx, gcpAccount)

Delete a GCP Integration from your Datadog account.

### Overview
Delete a given Datadog-GCP integration.
### Arguments
* **`project_id`** [*required*, *default* = **None**]: Your Google Cloud project ID to be deleted.
* **`client_email`** [*required*, *default* = **None**]: Your client email associated with the integration to be deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gcpAccount** | [**GcpAccount**](GcpAccount.md)| Delete a given Datadog-GCP integration. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGCPIntegration

> []GcpAccount ListGCPIntegration(ctx, )

List configured GCP integrations.

### Overview
List all Datadog-GCP integrations configured in your Datadog account.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GcpAccount**](GCPAccount.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGCPIntegration

> interface{} UpdateGCPIntegration(ctx, gcpAccount)

Update a GCP integration in your Datadog account.

### Overview
Update a Datadog-GCP integrations host_filters and/or automute. Requires a project_id and client_email, however these fields cannot be updated. If you need to update these fields please delete and use the create (POST) endpoint. The unspecified fields will keep their original values.
### Arguments
* **`project_id`** [*required*, *default* = **None**]: Your Google Cloud project ID associated with the integration to be updated.
* **`client_email`** [*required*, *default* = **None**]: Your client email associated with the integration to be updated.
* **`host_filters`** [*optional*, *default* = **None**]: Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.
* **`automute`** [*optional*, *default* = false]: Silence monitors for expected GCE instance shutdowns.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gcpAccount** | [**GcpAccount**](GcpAccount.md)| Update a Datadog-GCP integration. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

