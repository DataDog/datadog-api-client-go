# OrganizationSettingsSamlStrictMode

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Enabled** | Pointer to **bool** | Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. Learn more on the [SAML Strict documentation](https://docs.datadoghq.com/account_management/saml/#saml-strict). | [optional] 

## Methods

### NewOrganizationSettingsSamlStrictMode

`func NewOrganizationSettingsSamlStrictMode() *OrganizationSettingsSamlStrictMode`

NewOrganizationSettingsSamlStrictMode instantiates a new OrganizationSettingsSamlStrictMode object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationSettingsSamlStrictModeWithDefaults

`func NewOrganizationSettingsSamlStrictModeWithDefaults() *OrganizationSettingsSamlStrictMode`

NewOrganizationSettingsSamlStrictModeWithDefaults instantiates a new OrganizationSettingsSamlStrictMode object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEnabled

`func (o *OrganizationSettingsSamlStrictMode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationSettingsSamlStrictMode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationSettingsSamlStrictMode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationSettingsSamlStrictMode) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


