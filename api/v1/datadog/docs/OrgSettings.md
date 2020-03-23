# OrgSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateWidgetShare** | Pointer to **bool** | TODO. | [optional] 
**Saml** | Pointer to [**OrgSettingsSaml**](Org_settings_saml.md) |  | [optional] 
**SamlAutocreateAccessRole** | Pointer to [**AccessRole**](AccessRole.md) |  | [optional] 
**SamlAutocreateUsersDomains** | Pointer to [**OrgSettingsSamlAutocreateUsersDomains**](Org_settings_saml_autocreate_users_domains.md) |  | [optional] 
**SamlCanBeEnabled** | Pointer to **bool** | TODO. | [optional] 
**SamlIdpEndpoint** | Pointer to **string** | TODO. | [optional] 
**SamlIdpInitiatedLogin** | Pointer to [**OrgSettingsSamlIdpInitiatedLogin**](Org_settings_saml_idp_initiated_login.md) |  | [optional] 
**SamlIdpMetadataUploaded** | Pointer to **bool** | TODO. | [optional] 
**SamlLoginUrl** | Pointer to **string** | TODO. | [optional] 
**SamlStrictMode** | Pointer to [**OrgSettingsSamlIdpInitiatedLogin**](Org_settings_saml_idp_initiated_login.md) |  | [optional] 

## Methods

### NewOrgSettings

`func NewOrgSettings() *OrgSettings`

NewOrgSettings instantiates a new OrgSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingsWithDefaults

`func NewOrgSettingsWithDefaults() *OrgSettings`

NewOrgSettingsWithDefaults instantiates a new OrgSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateWidgetShare

`func (o *OrgSettings) GetPrivateWidgetShare() bool`

GetPrivateWidgetShare returns the PrivateWidgetShare field if non-nil, zero value otherwise.

### GetPrivateWidgetShareOk

`func (o *OrgSettings) GetPrivateWidgetShareOk() (bool, bool)`

GetPrivateWidgetShareOk returns a tuple with the PrivateWidgetShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateWidgetShare

`func (o *OrgSettings) HasPrivateWidgetShare() bool`

HasPrivateWidgetShare returns a boolean if a field has been set.

### SetPrivateWidgetShare

`func (o *OrgSettings) SetPrivateWidgetShare(v bool)`

SetPrivateWidgetShare gets a reference to the given bool and assigns it to the PrivateWidgetShare field.

### GetSaml

`func (o *OrgSettings) GetSaml() OrgSettingsSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *OrgSettings) GetSamlOk() (OrgSettingsSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaml

`func (o *OrgSettings) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### SetSaml

`func (o *OrgSettings) SetSaml(v OrgSettingsSaml)`

SetSaml gets a reference to the given OrgSettingsSaml and assigns it to the Saml field.

### GetSamlAutocreateAccessRole

`func (o *OrgSettings) GetSamlAutocreateAccessRole() AccessRole`

GetSamlAutocreateAccessRole returns the SamlAutocreateAccessRole field if non-nil, zero value otherwise.

### GetSamlAutocreateAccessRoleOk

`func (o *OrgSettings) GetSamlAutocreateAccessRoleOk() (AccessRole, bool)`

GetSamlAutocreateAccessRoleOk returns a tuple with the SamlAutocreateAccessRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlAutocreateAccessRole

`func (o *OrgSettings) HasSamlAutocreateAccessRole() bool`

HasSamlAutocreateAccessRole returns a boolean if a field has been set.

### SetSamlAutocreateAccessRole

`func (o *OrgSettings) SetSamlAutocreateAccessRole(v AccessRole)`

SetSamlAutocreateAccessRole gets a reference to the given AccessRole and assigns it to the SamlAutocreateAccessRole field.

### GetSamlAutocreateUsersDomains

`func (o *OrgSettings) GetSamlAutocreateUsersDomains() OrgSettingsSamlAutocreateUsersDomains`

GetSamlAutocreateUsersDomains returns the SamlAutocreateUsersDomains field if non-nil, zero value otherwise.

### GetSamlAutocreateUsersDomainsOk

`func (o *OrgSettings) GetSamlAutocreateUsersDomainsOk() (OrgSettingsSamlAutocreateUsersDomains, bool)`

GetSamlAutocreateUsersDomainsOk returns a tuple with the SamlAutocreateUsersDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlAutocreateUsersDomains

`func (o *OrgSettings) HasSamlAutocreateUsersDomains() bool`

HasSamlAutocreateUsersDomains returns a boolean if a field has been set.

### SetSamlAutocreateUsersDomains

`func (o *OrgSettings) SetSamlAutocreateUsersDomains(v OrgSettingsSamlAutocreateUsersDomains)`

SetSamlAutocreateUsersDomains gets a reference to the given OrgSettingsSamlAutocreateUsersDomains and assigns it to the SamlAutocreateUsersDomains field.

### GetSamlCanBeEnabled

`func (o *OrgSettings) GetSamlCanBeEnabled() bool`

GetSamlCanBeEnabled returns the SamlCanBeEnabled field if non-nil, zero value otherwise.

### GetSamlCanBeEnabledOk

`func (o *OrgSettings) GetSamlCanBeEnabledOk() (bool, bool)`

GetSamlCanBeEnabledOk returns a tuple with the SamlCanBeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlCanBeEnabled

`func (o *OrgSettings) HasSamlCanBeEnabled() bool`

HasSamlCanBeEnabled returns a boolean if a field has been set.

### SetSamlCanBeEnabled

`func (o *OrgSettings) SetSamlCanBeEnabled(v bool)`

SetSamlCanBeEnabled gets a reference to the given bool and assigns it to the SamlCanBeEnabled field.

### GetSamlIdpEndpoint

`func (o *OrgSettings) GetSamlIdpEndpoint() string`

GetSamlIdpEndpoint returns the SamlIdpEndpoint field if non-nil, zero value otherwise.

### GetSamlIdpEndpointOk

`func (o *OrgSettings) GetSamlIdpEndpointOk() (string, bool)`

GetSamlIdpEndpointOk returns a tuple with the SamlIdpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlIdpEndpoint

`func (o *OrgSettings) HasSamlIdpEndpoint() bool`

HasSamlIdpEndpoint returns a boolean if a field has been set.

### SetSamlIdpEndpoint

`func (o *OrgSettings) SetSamlIdpEndpoint(v string)`

SetSamlIdpEndpoint gets a reference to the given string and assigns it to the SamlIdpEndpoint field.

### GetSamlIdpInitiatedLogin

`func (o *OrgSettings) GetSamlIdpInitiatedLogin() OrgSettingsSamlIdpInitiatedLogin`

GetSamlIdpInitiatedLogin returns the SamlIdpInitiatedLogin field if non-nil, zero value otherwise.

### GetSamlIdpInitiatedLoginOk

`func (o *OrgSettings) GetSamlIdpInitiatedLoginOk() (OrgSettingsSamlIdpInitiatedLogin, bool)`

GetSamlIdpInitiatedLoginOk returns a tuple with the SamlIdpInitiatedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlIdpInitiatedLogin

`func (o *OrgSettings) HasSamlIdpInitiatedLogin() bool`

HasSamlIdpInitiatedLogin returns a boolean if a field has been set.

### SetSamlIdpInitiatedLogin

`func (o *OrgSettings) SetSamlIdpInitiatedLogin(v OrgSettingsSamlIdpInitiatedLogin)`

SetSamlIdpInitiatedLogin gets a reference to the given OrgSettingsSamlIdpInitiatedLogin and assigns it to the SamlIdpInitiatedLogin field.

### GetSamlIdpMetadataUploaded

`func (o *OrgSettings) GetSamlIdpMetadataUploaded() bool`

GetSamlIdpMetadataUploaded returns the SamlIdpMetadataUploaded field if non-nil, zero value otherwise.

### GetSamlIdpMetadataUploadedOk

`func (o *OrgSettings) GetSamlIdpMetadataUploadedOk() (bool, bool)`

GetSamlIdpMetadataUploadedOk returns a tuple with the SamlIdpMetadataUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlIdpMetadataUploaded

`func (o *OrgSettings) HasSamlIdpMetadataUploaded() bool`

HasSamlIdpMetadataUploaded returns a boolean if a field has been set.

### SetSamlIdpMetadataUploaded

`func (o *OrgSettings) SetSamlIdpMetadataUploaded(v bool)`

SetSamlIdpMetadataUploaded gets a reference to the given bool and assigns it to the SamlIdpMetadataUploaded field.

### GetSamlLoginUrl

`func (o *OrgSettings) GetSamlLoginUrl() string`

GetSamlLoginUrl returns the SamlLoginUrl field if non-nil, zero value otherwise.

### GetSamlLoginUrlOk

`func (o *OrgSettings) GetSamlLoginUrlOk() (string, bool)`

GetSamlLoginUrlOk returns a tuple with the SamlLoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlLoginUrl

`func (o *OrgSettings) HasSamlLoginUrl() bool`

HasSamlLoginUrl returns a boolean if a field has been set.

### SetSamlLoginUrl

`func (o *OrgSettings) SetSamlLoginUrl(v string)`

SetSamlLoginUrl gets a reference to the given string and assigns it to the SamlLoginUrl field.

### GetSamlStrictMode

`func (o *OrgSettings) GetSamlStrictMode() OrgSettingsSamlIdpInitiatedLogin`

GetSamlStrictMode returns the SamlStrictMode field if non-nil, zero value otherwise.

### GetSamlStrictModeOk

`func (o *OrgSettings) GetSamlStrictModeOk() (OrgSettingsSamlIdpInitiatedLogin, bool)`

GetSamlStrictModeOk returns a tuple with the SamlStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamlStrictMode

`func (o *OrgSettings) HasSamlStrictMode() bool`

HasSamlStrictMode returns a boolean if a field has been set.

### SetSamlStrictMode

`func (o *OrgSettings) SetSamlStrictMode(v OrgSettingsSamlIdpInitiatedLogin)`

SetSamlStrictMode gets a reference to the given OrgSettingsSamlIdpInitiatedLogin and assigns it to the SamlStrictMode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


