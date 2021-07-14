# AWSAccountDeleteRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AccessKeyId** | Pointer to **string** | Your AWS access key ID. Only required if your AWS account is a GovCloud or China account. | [optional] 
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | [optional] 
**RoleName** | Pointer to **string** | Your Datadog role delegation name. | [optional] 

## Methods

### NewAWSAccountDeleteRequest

`func NewAWSAccountDeleteRequest() *AWSAccountDeleteRequest`

NewAWSAccountDeleteRequest instantiates a new AWSAccountDeleteRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSAccountDeleteRequestWithDefaults

`func NewAWSAccountDeleteRequestWithDefaults() *AWSAccountDeleteRequest`

NewAWSAccountDeleteRequestWithDefaults instantiates a new AWSAccountDeleteRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccessKeyId

`func (o *AWSAccountDeleteRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSAccountDeleteRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSAccountDeleteRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AWSAccountDeleteRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccountId

`func (o *AWSAccountDeleteRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSAccountDeleteRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSAccountDeleteRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSAccountDeleteRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoleName

`func (o *AWSAccountDeleteRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSAccountDeleteRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSAccountDeleteRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AWSAccountDeleteRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


