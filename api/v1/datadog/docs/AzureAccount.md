# AzureAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Azure web application ID. | [optional] 
**ClientSecret** | Pointer to **string** | Your Azure web application secret key. | [optional] 
**Errors** | Pointer to **[]string** | Errors in your configuration. | [optional] 
**HostFilters** | Pointer to **string** | Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. | [optional] 
**NewClientId** | Pointer to **string** | Your New Azure web application ID. | [optional] 
**NewTenantName** | Pointer to **string** | Your New Azure Active Directory ID. | [optional] 
**TenantName** | Pointer to **string** | Your Azure Active Directory ID. | [optional] 

## Methods

### NewAzureAccount

`func NewAzureAccount() *AzureAccount`

NewAzureAccount instantiates a new AzureAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAccountWithDefaults

`func NewAzureAccountWithDefaults() *AzureAccount`

NewAzureAccountWithDefaults instantiates a new AzureAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureAccount) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureAccount) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AzureAccount) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureAccount) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureAccount) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AzureAccount) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetErrors

`func (o *AzureAccount) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AzureAccount) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AzureAccount) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AzureAccount) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetHostFilters

`func (o *AzureAccount) GetHostFilters() string`

GetHostFilters returns the HostFilters field if non-nil, zero value otherwise.

### GetHostFiltersOk

`func (o *AzureAccount) GetHostFiltersOk() (*string, bool)`

GetHostFiltersOk returns a tuple with the HostFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFilters

`func (o *AzureAccount) SetHostFilters(v string)`

SetHostFilters sets HostFilters field to given value.

### HasHostFilters

`func (o *AzureAccount) HasHostFilters() bool`

HasHostFilters returns a boolean if a field has been set.

### GetNewClientId

`func (o *AzureAccount) GetNewClientId() string`

GetNewClientId returns the NewClientId field if non-nil, zero value otherwise.

### GetNewClientIdOk

`func (o *AzureAccount) GetNewClientIdOk() (*string, bool)`

GetNewClientIdOk returns a tuple with the NewClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientId

`func (o *AzureAccount) SetNewClientId(v string)`

SetNewClientId sets NewClientId field to given value.

### HasNewClientId

`func (o *AzureAccount) HasNewClientId() bool`

HasNewClientId returns a boolean if a field has been set.

### GetNewTenantName

`func (o *AzureAccount) GetNewTenantName() string`

GetNewTenantName returns the NewTenantName field if non-nil, zero value otherwise.

### GetNewTenantNameOk

`func (o *AzureAccount) GetNewTenantNameOk() (*string, bool)`

GetNewTenantNameOk returns a tuple with the NewTenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTenantName

`func (o *AzureAccount) SetNewTenantName(v string)`

SetNewTenantName sets NewTenantName field to given value.

### HasNewTenantName

`func (o *AzureAccount) HasNewTenantName() bool`

HasNewTenantName returns a boolean if a field has been set.

### GetTenantName

`func (o *AzureAccount) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *AzureAccount) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *AzureAccount) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *AzureAccount) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


