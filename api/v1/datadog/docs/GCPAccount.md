# GcpAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProviderX509CertUrl** | **string** | Should be &#x60;https://www.googleapis.com/oauth2/v1/certs&#x60;. | [optional] 
**AuthUri** | **string** | Should be &#x60;https://accounts.google.com/o/oauth2/auth&#x60;. | [optional] 
**Automute** | **bool** | Silence monitors for expected GCE instance shutdowns. | [optional] 
**ClientEmail** | **string** | Your email found in your JSON service account key. | [optional] 
**ClientId** | **string** | Your ID found in your JSON service account key. | [optional] 
**ClientX509CertUrl** | **string** | Should be &#x60;https://www.googleapis.com/robot/v1/metadata/x509/&lt;CLIENT_EMAIL&gt;&#x60; where &#x60;&lt;CLIENT_EMAIL&gt;&#x60; is the email found in your JSON service account key. | [optional] 
**Errors** | **[]string** | An array of errors. | [optional] 
**HostFilters** | **string** | Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. | [optional] 
**PrivateKey** | **string** | Your private key name found in your JSON service account key. | [optional] 
**PrivateKeyId** | **string** | Your private key ID found in your JSON service account key. | [optional] 
**ProjectId** | **string** | Your Google Cloud project ID found in your JSON service account key. | [optional] 
**TokenUri** | **string** | Should be &#x60;https://accounts.google.com/o/oauth2/token&#x60;. | [optional] 
**Type** | **string** | The value for service_account found in your JSON service account key. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


