# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to [**OrgBilling**](Org_billing.md) |  | [optional] 
**Created** | Pointer to **string** | TODO. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the organization. | [optional] 
**Name** | Pointer to **string** | The name of the new child-organization, limited to 32 characters. | [optional] 
**PublicId** | Pointer to **string** | The public_id of the org you are operating within. | [optional] 
**Settings** | Pointer to [**OrgSettings**](Org_settings.md) |  | [optional] 
**Subscription** | Pointer to [**OrgSubscription**](Org_subscription.md) |  | [optional] 

## Methods

### NewOrg

`func NewOrg() *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *Org) GetBilling() OrgBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Org) GetBillingOk() (OrgBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBilling

`func (o *Org) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### SetBilling

`func (o *Org) SetBilling(v OrgBilling)`

SetBilling gets a reference to the given OrgBilling and assigns it to the Billing field.

### GetCreated

`func (o *Org) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Org) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Org) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Org) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetDescription

`func (o *Org) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Org) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Org) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Org) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Org) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPublicId

`func (o *Org) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Org) GetPublicIdOk() (string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicId

`func (o *Org) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### SetPublicId

`func (o *Org) SetPublicId(v string)`

SetPublicId gets a reference to the given string and assigns it to the PublicId field.

### GetSettings

`func (o *Org) GetSettings() OrgSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Org) GetSettingsOk() (OrgSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *Org) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *Org) SetSettings(v OrgSettings)`

SetSettings gets a reference to the given OrgSettings and assigns it to the Settings field.

### GetSubscription

`func (o *Org) GetSubscription() OrgSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Org) GetSubscriptionOk() (OrgSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscription

`func (o *Org) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### SetSubscription

`func (o *Org) SetSubscription(v OrgSubscription)`

SetSubscription gets a reference to the given OrgSubscription and assigns it to the Subscription field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


