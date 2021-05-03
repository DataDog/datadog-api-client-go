# AWSLogsServicesRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AccountId** | **string** | Your AWS Account ID without dashes. | 
**Services** | **[]string** | Array of services IDs set to enable automatic log collection. Discover the list of available services with the get list of AWS log ready services API endpoint. | 

## Methods

### NewAWSLogsServicesRequest

`func NewAWSLogsServicesRequest(accountId string, services []string, ) *AWSLogsServicesRequest`

NewAWSLogsServicesRequest instantiates a new AWSLogsServicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSLogsServicesRequestWithDefaults

`func NewAWSLogsServicesRequestWithDefaults() *AWSLogsServicesRequest`

NewAWSLogsServicesRequestWithDefaults instantiates a new AWSLogsServicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AWSLogsServicesRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSLogsServicesRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSLogsServicesRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetServices

`func (o *AWSLogsServicesRequest) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AWSLogsServicesRequest) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AWSLogsServicesRequest) SetServices(v []string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


