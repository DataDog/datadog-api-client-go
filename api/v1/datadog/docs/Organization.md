# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to [**OrganizationBilling**](Organization_billing.md) |  | [optional] 
**Created** | Pointer to **string** | Date of the organization creation. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the organization. | [optional] 
**Name** | Pointer to **string** | The name of the new child-organization, limited to 32 characters. | [optional] 
**PublicId** | Pointer to **string** | The &#x60;public_id&#x60; of the org you are operating within. | [optional] 
**Settings** | Pointer to [**OrganizationSettings**](Organization_settings.md) |  | [optional] 
**Subscription** | Pointer to [**OrganizationSubscription**](Organization_subscription.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *Organization) GetBilling() OrganizationBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Organization) GetBillingOk() (*OrganizationBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Organization) SetBilling(v OrganizationBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Organization) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreated

`func (o *Organization) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Organization) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Organization) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Organization) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicId

`func (o *Organization) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Organization) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Organization) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *Organization) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetSettings

`func (o *Organization) GetSettings() OrganizationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Organization) GetSettingsOk() (*OrganizationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Organization) SetSettings(v OrganizationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Organization) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSubscription

`func (o *Organization) GetSubscription() OrganizationSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Organization) GetSubscriptionOk() (*OrganizationSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Organization) SetSubscription(v OrganizationSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Organization) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


