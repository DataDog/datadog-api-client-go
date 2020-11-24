# AWSLogsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | [optional] 
**Lambdas** | Pointer to [**[]AWSLogsListResponseLambdas**](AWSLogsListResponseLambdas.md) | List of ARNs configured in your Datadog account. | [optional] 
**Services** | Pointer to **[]string** | Array of services IDs. | [optional] 

## Methods

### NewAWSLogsListResponse

`func NewAWSLogsListResponse() *AWSLogsListResponse`

NewAWSLogsListResponse instantiates a new AWSLogsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSLogsListResponseWithDefaults

`func NewAWSLogsListResponseWithDefaults() *AWSLogsListResponse`

NewAWSLogsListResponseWithDefaults instantiates a new AWSLogsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AWSLogsListResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSLogsListResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSLogsListResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSLogsListResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLambdas

`func (o *AWSLogsListResponse) GetLambdas() []AWSLogsListResponseLambdas`

GetLambdas returns the Lambdas field if non-nil, zero value otherwise.

### GetLambdasOk

`func (o *AWSLogsListResponse) GetLambdasOk() (*[]AWSLogsListResponseLambdas, bool)`

GetLambdasOk returns a tuple with the Lambdas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdas

`func (o *AWSLogsListResponse) SetLambdas(v []AWSLogsListResponseLambdas)`

SetLambdas sets Lambdas field to given value.

### HasLambdas

`func (o *AWSLogsListResponse) HasLambdas() bool`

HasLambdas returns a boolean if a field has been set.

### GetServices

`func (o *AWSLogsListResponse) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AWSLogsListResponse) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AWSLogsListResponse) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AWSLogsListResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


