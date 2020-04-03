# AWSAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | Your AWS access key ID. Only required if your AWS account is a GovCloud or China account. | [optional] 
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | [optional] 
**AccountSpecificNamespaceRules** | Pointer to **map[string]bool** | An object (in the form {\&quot;namespace1\&quot;:true/false, \&quot;namespace2\&quot;:true/false}) that enables or disables metric collection for specific AWS namespaces for this AWS account only. A list of namespaces can be found at the /v1/integration/aws/available_namespace_rules endpoint. | [optional] 
**ExcludedRegions** | Pointer to **[]string** | An array of AWS regions to exclude from metrics collection. | [optional] 
**FilterTags** | Pointer to **[]string** | The array of EC2 tags (in the form key:value) defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as ? (for single characters) and * (for multiple characters) can also be used. Only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given tag can also be excluded by adding ! before the tag. For example, &#x60;env:production,instance-type:c1.*,!region:us-east-1&#x60; | [optional] 
**HostTags** | Pointer to **[]string** | Array of tags (in the form key:value) to add to all hosts and metrics reporting through this integration. | [optional] 
**RoleName** | Pointer to **string** | Your Datadog role delegation name. | [optional] 
**SecretAccessKey** | Pointer to **string** | Your AWS secret access key. Only required if your AWS account is a GovCloud or China account. | [optional] 

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

`func (o *AWSAccount) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSAccount) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AWSAccount) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccountId

`func (o *AWSAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountSpecificNamespaceRules

`func (o *AWSAccount) GetAccountSpecificNamespaceRules() map[string]bool`

GetAccountSpecificNamespaceRules returns the AccountSpecificNamespaceRules field if non-nil, zero value otherwise.

### GetAccountSpecificNamespaceRulesOk

`func (o *AWSAccount) GetAccountSpecificNamespaceRulesOk() (*map[string]bool, bool)`

GetAccountSpecificNamespaceRulesOk returns a tuple with the AccountSpecificNamespaceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSpecificNamespaceRules

`func (o *AWSAccount) SetAccountSpecificNamespaceRules(v map[string]bool)`

SetAccountSpecificNamespaceRules sets AccountSpecificNamespaceRules field to given value.

### HasAccountSpecificNamespaceRules

`func (o *AWSAccount) HasAccountSpecificNamespaceRules() bool`

HasAccountSpecificNamespaceRules returns a boolean if a field has been set.

### GetExcludedRegions

`func (o *AWSAccount) GetExcludedRegions() []string`

GetExcludedRegions returns the ExcludedRegions field if non-nil, zero value otherwise.

### GetExcludedRegionsOk

`func (o *AWSAccount) GetExcludedRegionsOk() (*[]string, bool)`

GetExcludedRegionsOk returns a tuple with the ExcludedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedRegions

`func (o *AWSAccount) SetExcludedRegions(v []string)`

SetExcludedRegions sets ExcludedRegions field to given value.

### HasExcludedRegions

`func (o *AWSAccount) HasExcludedRegions() bool`

HasExcludedRegions returns a boolean if a field has been set.

### GetFilterTags

`func (o *AWSAccount) GetFilterTags() []string`

GetFilterTags returns the FilterTags field if non-nil, zero value otherwise.

### GetFilterTagsOk

`func (o *AWSAccount) GetFilterTagsOk() (*[]string, bool)`

GetFilterTagsOk returns a tuple with the FilterTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTags

`func (o *AWSAccount) SetFilterTags(v []string)`

SetFilterTags sets FilterTags field to given value.

### HasFilterTags

`func (o *AWSAccount) HasFilterTags() bool`

HasFilterTags returns a boolean if a field has been set.

### GetHostTags

`func (o *AWSAccount) GetHostTags() []string`

GetHostTags returns the HostTags field if non-nil, zero value otherwise.

### GetHostTagsOk

`func (o *AWSAccount) GetHostTagsOk() (*[]string, bool)`

GetHostTagsOk returns a tuple with the HostTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTags

`func (o *AWSAccount) SetHostTags(v []string)`

SetHostTags sets HostTags field to given value.

### HasHostTags

`func (o *AWSAccount) HasHostTags() bool`

HasHostTags returns a boolean if a field has been set.

### GetRoleName

`func (o *AWSAccount) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSAccount) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSAccount) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AWSAccount) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *AWSAccount) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSAccount) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSAccount) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AWSAccount) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


