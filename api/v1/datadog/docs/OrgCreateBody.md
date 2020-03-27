# OrgCreateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to [**OrgBilling**](Org_billing.md) |  | 
**Name** | Pointer to **string** | The name of the new child-organization, limited to 32 characters. | 
**Subscription** | Pointer to [**OrgSubscription**](Org_subscription.md) |  | 

## Methods

### NewOrgCreateBody

`func NewOrgCreateBody(billing OrgBilling, name string, subscription OrgSubscription, ) *OrgCreateBody`

NewOrgCreateBody instantiates a new OrgCreateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCreateBodyWithDefaults

`func NewOrgCreateBodyWithDefaults() *OrgCreateBody`

NewOrgCreateBodyWithDefaults instantiates a new OrgCreateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *OrgCreateBody) GetBilling() OrgBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *OrgCreateBody) GetBillingOk() (*OrgBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *OrgCreateBody) SetBilling(v OrgBilling)`

SetBilling sets Billing field to given value.


### GetName

`func (o *OrgCreateBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgCreateBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgCreateBody) SetName(v string)`

SetName sets Name field to given value.


### GetSubscription

`func (o *OrgCreateBody) GetSubscription() OrgSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *OrgCreateBody) GetSubscriptionOk() (*OrgSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *OrgCreateBody) SetSubscription(v OrgSubscription)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


