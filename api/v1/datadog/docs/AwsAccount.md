# AWSAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountSpecificNamespaceRules** | Pointer to **map[string]bool** |  | [optional] 
**FilterTags** | Pointer to **[]string** |  | [optional] 
**HostTags** | Pointer to **[]string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewAWSAccount

`func NewAWSAccount() *AWSAccount`

NewAWSAccount instantiates a new AWSAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSAccountWithDefaults

`func NewAWSAccountWithDefaults() *AWSAccount`

NewAWSAccountWithDefaults instantiates a new AWSAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AWSAccount) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSAccount) GetAccessKeyIdOk() (string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessKeyId

`func (o *AWSAccount) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyId

`func (o *AWSAccount) SetAccessKeyId(v string)`

SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.

### GetAccountId

`func (o *AWSAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSAccount) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *AWSAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *AWSAccount) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetAccountSpecificNamespaceRules

`func (o *AWSAccount) GetAccountSpecificNamespaceRules() map[string]bool`

GetAccountSpecificNamespaceRules returns the AccountSpecificNamespaceRules field if non-nil, zero value otherwise.

### GetAccountSpecificNamespaceRulesOk

`func (o *AWSAccount) GetAccountSpecificNamespaceRulesOk() (map[string]bool, bool)`

GetAccountSpecificNamespaceRulesOk returns a tuple with the AccountSpecificNamespaceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountSpecificNamespaceRules

`func (o *AWSAccount) HasAccountSpecificNamespaceRules() bool`

HasAccountSpecificNamespaceRules returns a boolean if a field has been set.

### SetAccountSpecificNamespaceRules

`func (o *AWSAccount) SetAccountSpecificNamespaceRules(v map[string]bool)`

SetAccountSpecificNamespaceRules gets a reference to the given map[string]bool and assigns it to the AccountSpecificNamespaceRules field.

### GetFilterTags

`func (o *AWSAccount) GetFilterTags() []string`

GetFilterTags returns the FilterTags field if non-nil, zero value otherwise.

### GetFilterTagsOk

`func (o *AWSAccount) GetFilterTagsOk() ([]string, bool)`

GetFilterTagsOk returns a tuple with the FilterTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilterTags

`func (o *AWSAccount) HasFilterTags() bool`

HasFilterTags returns a boolean if a field has been set.

### SetFilterTags

`func (o *AWSAccount) SetFilterTags(v []string)`

SetFilterTags gets a reference to the given []string and assigns it to the FilterTags field.

### GetHostTags

`func (o *AWSAccount) GetHostTags() []string`

GetHostTags returns the HostTags field if non-nil, zero value otherwise.

### GetHostTagsOk

`func (o *AWSAccount) GetHostTagsOk() ([]string, bool)`

GetHostTagsOk returns a tuple with the HostTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostTags

`func (o *AWSAccount) HasHostTags() bool`

HasHostTags returns a boolean if a field has been set.

### SetHostTags

`func (o *AWSAccount) SetHostTags(v []string)`

SetHostTags gets a reference to the given []string and assigns it to the HostTags field.

### GetRoleName

`func (o *AWSAccount) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSAccount) GetRoleNameOk() (string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleName

`func (o *AWSAccount) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleName

`func (o *AWSAccount) SetRoleName(v string)`

SetRoleName gets a reference to the given string and assigns it to the RoleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


