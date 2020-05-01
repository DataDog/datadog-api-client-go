# OrganizationSettingsSamlIdpInitiatedLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wether SAML IdP initiated login is enabled, learn more in the [SAML documentation](https://docs.datadoghq.com/account_management/saml/#idp-initiated-login). | [optional] 

## Methods

### NewOrganizationSettingsSamlIdpInitiatedLogin

`func NewOrganizationSettingsSamlIdpInitiatedLogin() *OrganizationSettingsSamlIdpInitiatedLogin`

NewOrganizationSettingsSamlIdpInitiatedLogin instantiates a new OrganizationSettingsSamlIdpInitiatedLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsSamlIdpInitiatedLoginWithDefaults

`func NewOrganizationSettingsSamlIdpInitiatedLoginWithDefaults() *OrganizationSettingsSamlIdpInitiatedLogin`

NewOrganizationSettingsSamlIdpInitiatedLoginWithDefaults instantiates a new OrganizationSettingsSamlIdpInitiatedLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OrganizationSettingsSamlIdpInitiatedLogin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationSettingsSamlIdpInitiatedLogin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationSettingsSamlIdpInitiatedLogin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationSettingsSamlIdpInitiatedLogin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


