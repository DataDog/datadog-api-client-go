# OrganizationSettings

## Properties

| Name                           | Type                                                                                                               | Description                                                                                     | Notes                                       |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **PrivateWidgetShare**         | Pointer to **bool**                                                                                                | Whether or not the organization users can share widgets outside of Datadog.                     | [optional]                                  |
| **Saml**                       | Pointer to [**OrganizationSettingsSaml**](OrganizationSettingsSaml.md)                                             |                                                                                                 | [optional]                                  |
| **SamlAutocreateAccessRole**   | Pointer to [**AccessRole**](AccessRole.md)                                                                         |                                                                                                 | [optional] [default to ACCESSROLE_STANDARD] |
| **SamlAutocreateUsersDomains** | Pointer to [**OrganizationSettingsSamlAutocreateUsersDomains**](OrganizationSettingsSamlAutocreateUsersDomains.md) |                                                                                                 | [optional]                                  |
| **SamlCanBeEnabled**           | Pointer to **bool**                                                                                                | Whether or not SAML can be enabled for this organization.                                       | [optional]                                  |
| **SamlIdpEndpoint**            | Pointer to **string**                                                                                              | Identity provider endpoint for SAML authentication.                                             | [optional]                                  |
| **SamlIdpInitiatedLogin**      | Pointer to [**OrganizationSettingsSamlIdpInitiatedLogin**](OrganizationSettingsSamlIdpInitiatedLogin.md)           |                                                                                                 | [optional]                                  |
| **SamlIdpMetadataUploaded**    | Pointer to **bool**                                                                                                | Whether or not a SAML identity provider metadata file was provided to the Datadog organization. | [optional]                                  |
| **SamlLoginUrl**               | Pointer to **string**                                                                                              | URL for SAML logging.                                                                           | [optional]                                  |
| **SamlStrictMode**             | Pointer to [**OrganizationSettingsSamlStrictMode**](OrganizationSettingsSamlStrictMode.md)                         |                                                                                                 | [optional]                                  |

## Methods

### NewOrganizationSettings

`func NewOrganizationSettings() *OrganizationSettings`

NewOrganizationSettings instantiates a new OrganizationSettings object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationSettingsWithDefaults

`func NewOrganizationSettingsWithDefaults() *OrganizationSettings`

NewOrganizationSettingsWithDefaults instantiates a new OrganizationSettings object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPrivateWidgetShare

`func (o *OrganizationSettings) GetPrivateWidgetShare() bool`

GetPrivateWidgetShare returns the PrivateWidgetShare field if non-nil, zero value otherwise.

### GetPrivateWidgetShareOk

`func (o *OrganizationSettings) GetPrivateWidgetShareOk() (*bool, bool)`

GetPrivateWidgetShareOk returns a tuple with the PrivateWidgetShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateWidgetShare

`func (o *OrganizationSettings) SetPrivateWidgetShare(v bool)`

SetPrivateWidgetShare sets PrivateWidgetShare field to given value.

### HasPrivateWidgetShare

`func (o *OrganizationSettings) HasPrivateWidgetShare() bool`

HasPrivateWidgetShare returns a boolean if a field has been set.

### GetSaml

`func (o *OrganizationSettings) GetSaml() OrganizationSettingsSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *OrganizationSettings) GetSamlOk() (*OrganizationSettingsSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *OrganizationSettings) SetSaml(v OrganizationSettingsSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *OrganizationSettings) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetSamlAutocreateAccessRole

`func (o *OrganizationSettings) GetSamlAutocreateAccessRole() AccessRole`

GetSamlAutocreateAccessRole returns the SamlAutocreateAccessRole field if non-nil, zero value otherwise.

### GetSamlAutocreateAccessRoleOk

`func (o *OrganizationSettings) GetSamlAutocreateAccessRoleOk() (*AccessRole, bool)`

GetSamlAutocreateAccessRoleOk returns a tuple with the SamlAutocreateAccessRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAutocreateAccessRole

`func (o *OrganizationSettings) SetSamlAutocreateAccessRole(v AccessRole)`

SetSamlAutocreateAccessRole sets SamlAutocreateAccessRole field to given value.

### HasSamlAutocreateAccessRole

`func (o *OrganizationSettings) HasSamlAutocreateAccessRole() bool`

HasSamlAutocreateAccessRole returns a boolean if a field has been set.

### GetSamlAutocreateUsersDomains

`func (o *OrganizationSettings) GetSamlAutocreateUsersDomains() OrganizationSettingsSamlAutocreateUsersDomains`

GetSamlAutocreateUsersDomains returns the SamlAutocreateUsersDomains field if non-nil, zero value otherwise.

### GetSamlAutocreateUsersDomainsOk

`func (o *OrganizationSettings) GetSamlAutocreateUsersDomainsOk() (*OrganizationSettingsSamlAutocreateUsersDomains, bool)`

GetSamlAutocreateUsersDomainsOk returns a tuple with the SamlAutocreateUsersDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAutocreateUsersDomains

`func (o *OrganizationSettings) SetSamlAutocreateUsersDomains(v OrganizationSettingsSamlAutocreateUsersDomains)`

SetSamlAutocreateUsersDomains sets SamlAutocreateUsersDomains field to given value.

### HasSamlAutocreateUsersDomains

`func (o *OrganizationSettings) HasSamlAutocreateUsersDomains() bool`

HasSamlAutocreateUsersDomains returns a boolean if a field has been set.

### GetSamlCanBeEnabled

`func (o *OrganizationSettings) GetSamlCanBeEnabled() bool`

GetSamlCanBeEnabled returns the SamlCanBeEnabled field if non-nil, zero value otherwise.

### GetSamlCanBeEnabledOk

`func (o *OrganizationSettings) GetSamlCanBeEnabledOk() (*bool, bool)`

GetSamlCanBeEnabledOk returns a tuple with the SamlCanBeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCanBeEnabled

`func (o *OrganizationSettings) SetSamlCanBeEnabled(v bool)`

SetSamlCanBeEnabled sets SamlCanBeEnabled field to given value.

### HasSamlCanBeEnabled

`func (o *OrganizationSettings) HasSamlCanBeEnabled() bool`

HasSamlCanBeEnabled returns a boolean if a field has been set.

### GetSamlIdpEndpoint

`func (o *OrganizationSettings) GetSamlIdpEndpoint() string`

GetSamlIdpEndpoint returns the SamlIdpEndpoint field if non-nil, zero value otherwise.

### GetSamlIdpEndpointOk

`func (o *OrganizationSettings) GetSamlIdpEndpointOk() (*string, bool)`

GetSamlIdpEndpointOk returns a tuple with the SamlIdpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpEndpoint

`func (o *OrganizationSettings) SetSamlIdpEndpoint(v string)`

SetSamlIdpEndpoint sets SamlIdpEndpoint field to given value.

### HasSamlIdpEndpoint

`func (o *OrganizationSettings) HasSamlIdpEndpoint() bool`

HasSamlIdpEndpoint returns a boolean if a field has been set.

### GetSamlIdpInitiatedLogin

`func (o *OrganizationSettings) GetSamlIdpInitiatedLogin() OrganizationSettingsSamlIdpInitiatedLogin`

GetSamlIdpInitiatedLogin returns the SamlIdpInitiatedLogin field if non-nil, zero value otherwise.

### GetSamlIdpInitiatedLoginOk

`func (o *OrganizationSettings) GetSamlIdpInitiatedLoginOk() (*OrganizationSettingsSamlIdpInitiatedLogin, bool)`

GetSamlIdpInitiatedLoginOk returns a tuple with the SamlIdpInitiatedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpInitiatedLogin

`func (o *OrganizationSettings) SetSamlIdpInitiatedLogin(v OrganizationSettingsSamlIdpInitiatedLogin)`

SetSamlIdpInitiatedLogin sets SamlIdpInitiatedLogin field to given value.

### HasSamlIdpInitiatedLogin

`func (o *OrganizationSettings) HasSamlIdpInitiatedLogin() bool`

HasSamlIdpInitiatedLogin returns a boolean if a field has been set.

### GetSamlIdpMetadataUploaded

`func (o *OrganizationSettings) GetSamlIdpMetadataUploaded() bool`

GetSamlIdpMetadataUploaded returns the SamlIdpMetadataUploaded field if non-nil, zero value otherwise.

### GetSamlIdpMetadataUploadedOk

`func (o *OrganizationSettings) GetSamlIdpMetadataUploadedOk() (*bool, bool)`

GetSamlIdpMetadataUploadedOk returns a tuple with the SamlIdpMetadataUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpMetadataUploaded

`func (o *OrganizationSettings) SetSamlIdpMetadataUploaded(v bool)`

SetSamlIdpMetadataUploaded sets SamlIdpMetadataUploaded field to given value.

### HasSamlIdpMetadataUploaded

`func (o *OrganizationSettings) HasSamlIdpMetadataUploaded() bool`

HasSamlIdpMetadataUploaded returns a boolean if a field has been set.

### GetSamlLoginUrl

`func (o *OrganizationSettings) GetSamlLoginUrl() string`

GetSamlLoginUrl returns the SamlLoginUrl field if non-nil, zero value otherwise.

### GetSamlLoginUrlOk

`func (o *OrganizationSettings) GetSamlLoginUrlOk() (*string, bool)`

GetSamlLoginUrlOk returns a tuple with the SamlLoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLoginUrl

`func (o *OrganizationSettings) SetSamlLoginUrl(v string)`

SetSamlLoginUrl sets SamlLoginUrl field to given value.

### HasSamlLoginUrl

`func (o *OrganizationSettings) HasSamlLoginUrl() bool`

HasSamlLoginUrl returns a boolean if a field has been set.

### GetSamlStrictMode

`func (o *OrganizationSettings) GetSamlStrictMode() OrganizationSettingsSamlStrictMode`

GetSamlStrictMode returns the SamlStrictMode field if non-nil, zero value otherwise.

### GetSamlStrictModeOk

`func (o *OrganizationSettings) GetSamlStrictModeOk() (*OrganizationSettingsSamlStrictMode, bool)`

GetSamlStrictModeOk returns a tuple with the SamlStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlStrictMode

`func (o *OrganizationSettings) SetSamlStrictMode(v OrganizationSettingsSamlStrictMode)`

SetSamlStrictMode sets SamlStrictMode field to given value.

### HasSamlStrictMode

`func (o *OrganizationSettings) HasSamlStrictMode() bool`

HasSamlStrictMode returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
