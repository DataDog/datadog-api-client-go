# AzureAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Azure web application ID. | [optional] 
**ClientSecret** | Pointer to **string** | Your Azure web application secret key. | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**HostFilters** | Pointer to **string** | Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. | [optional] 
**NewClientId** | Pointer to **string** | Your New Azure web application ID. | [optional] 
**NewTenantName** | Pointer to **string** | Your New Azure Active Directory ID. | [optional] 
**TenantName** | Pointer to **string** | Your Azure Active Directory ID. | [optional] 

## Methods

### GetClientId

`func (o *AzureAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureAccount) GetClientIdOk() (string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientId

`func (o *AzureAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientId

`func (o *AzureAccount) SetClientId(v string)`

SetClientId gets a reference to the given string and assigns it to the ClientId field.

### GetClientSecret

`func (o *AzureAccount) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureAccount) GetClientSecretOk() (string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientSecret

`func (o *AzureAccount) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecret

`func (o *AzureAccount) SetClientSecret(v string)`

SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.

### GetErrors

`func (o *AzureAccount) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AzureAccount) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *AzureAccount) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *AzureAccount) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.

### GetHostFilters

`func (o *AzureAccount) GetHostFilters() string`

GetHostFilters returns the HostFilters field if non-nil, zero value otherwise.

### GetHostFiltersOk

`func (o *AzureAccount) GetHostFiltersOk() (string, bool)`

GetHostFiltersOk returns a tuple with the HostFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostFilters

`func (o *AzureAccount) HasHostFilters() bool`

HasHostFilters returns a boolean if a field has been set.

### SetHostFilters

`func (o *AzureAccount) SetHostFilters(v string)`

SetHostFilters gets a reference to the given string and assigns it to the HostFilters field.

### GetNewClientId

`func (o *AzureAccount) GetNewClientId() string`

GetNewClientId returns the NewClientId field if non-nil, zero value otherwise.

### GetNewClientIdOk

`func (o *AzureAccount) GetNewClientIdOk() (string, bool)`

GetNewClientIdOk returns a tuple with the NewClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewClientId

`func (o *AzureAccount) HasNewClientId() bool`

HasNewClientId returns a boolean if a field has been set.

### SetNewClientId

`func (o *AzureAccount) SetNewClientId(v string)`

SetNewClientId gets a reference to the given string and assigns it to the NewClientId field.

### GetNewTenantName

`func (o *AzureAccount) GetNewTenantName() string`

GetNewTenantName returns the NewTenantName field if non-nil, zero value otherwise.

### GetNewTenantNameOk

`func (o *AzureAccount) GetNewTenantNameOk() (string, bool)`

GetNewTenantNameOk returns a tuple with the NewTenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewTenantName

`func (o *AzureAccount) HasNewTenantName() bool`

HasNewTenantName returns a boolean if a field has been set.

### SetNewTenantName

`func (o *AzureAccount) SetNewTenantName(v string)`

SetNewTenantName gets a reference to the given string and assigns it to the NewTenantName field.

### GetTenantName

`func (o *AzureAccount) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *AzureAccount) GetTenantNameOk() (string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTenantName

`func (o *AzureAccount) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantName

`func (o *AzureAccount) SetTenantName(v string)`

SetTenantName gets a reference to the given string and assigns it to the TenantName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


