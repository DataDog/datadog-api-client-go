# OrganizationSettingsSamlAutocreateUsersDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | List of domains where the SAML automated user creation is enabled. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the automated user creation based on SAML domain is enabled. | [optional] 

## Methods

### NewOrganizationSettingsSamlAutocreateUsersDomains

`func NewOrganizationSettingsSamlAutocreateUsersDomains() *OrganizationSettingsSamlAutocreateUsersDomains`

NewOrganizationSettingsSamlAutocreateUsersDomains instantiates a new OrganizationSettingsSamlAutocreateUsersDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsSamlAutocreateUsersDomainsWithDefaults

`func NewOrganizationSettingsSamlAutocreateUsersDomainsWithDefaults() *OrganizationSettingsSamlAutocreateUsersDomains`

NewOrganizationSettingsSamlAutocreateUsersDomainsWithDefaults instantiates a new OrganizationSettingsSamlAutocreateUsersDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationSettingsSamlAutocreateUsersDomains) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


