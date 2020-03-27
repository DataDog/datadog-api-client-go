# OrgSettingsSamlAutocreateUsersDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | List of domains where the SAML automated user creation is enabled. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the automated user creation based on SAML domain is enabled. | [optional] 

## Methods

### NewOrgSettingsSamlAutocreateUsersDomains

`func NewOrgSettingsSamlAutocreateUsersDomains() *OrgSettingsSamlAutocreateUsersDomains`

NewOrgSettingsSamlAutocreateUsersDomains instantiates a new OrgSettingsSamlAutocreateUsersDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingsSamlAutocreateUsersDomainsWithDefaults

`func NewOrgSettingsSamlAutocreateUsersDomainsWithDefaults() *OrgSettingsSamlAutocreateUsersDomains`

NewOrgSettingsSamlAutocreateUsersDomainsWithDefaults instantiates a new OrgSettingsSamlAutocreateUsersDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *OrgSettingsSamlAutocreateUsersDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrgSettingsSamlAutocreateUsersDomains) GetDomainsOk() ([]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomains

`func (o *OrgSettingsSamlAutocreateUsersDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomains

`func (o *OrgSettingsSamlAutocreateUsersDomains) SetDomains(v []string)`

SetDomains gets a reference to the given []string and assigns it to the Domains field.

### GetEnabled

`func (o *OrgSettingsSamlAutocreateUsersDomains) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrgSettingsSamlAutocreateUsersDomains) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *OrgSettingsSamlAutocreateUsersDomains) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *OrgSettingsSamlAutocreateUsersDomains) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


