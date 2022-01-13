# OrganizationCreateBody

## Properties

| Name             | Type                                                                   | Description                                                       | Notes      |
| ---------------- | ---------------------------------------------------------------------- | ----------------------------------------------------------------- | ---------- |
| **Billing**      | Pointer to [**OrganizationBilling**](OrganizationBilling.md)           |                                                                   | [optional] |
| **Name**         | **string**                                                             | The name of the new child-organization, limited to 32 characters. |
| **Subscription** | Pointer to [**OrganizationSubscription**](OrganizationSubscription.md) |                                                                   | [optional] |

## Methods

### NewOrganizationCreateBody

`func NewOrganizationCreateBody(name string) *OrganizationCreateBody`

NewOrganizationCreateBody instantiates a new OrganizationCreateBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationCreateBodyWithDefaults

`func NewOrganizationCreateBodyWithDefaults() *OrganizationCreateBody`

NewOrganizationCreateBodyWithDefaults instantiates a new OrganizationCreateBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBilling

`func (o *OrganizationCreateBody) GetBilling() OrganizationBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *OrganizationCreateBody) GetBillingOk() (*OrganizationBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *OrganizationCreateBody) SetBilling(v OrganizationBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *OrganizationCreateBody) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetName

`func (o *OrganizationCreateBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreateBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreateBody) SetName(v string)`

SetName sets Name field to given value.

### GetSubscription

`func (o *OrganizationCreateBody) GetSubscription() OrganizationSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *OrganizationCreateBody) GetSubscriptionOk() (*OrganizationSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *OrganizationCreateBody) SetSubscription(v OrganizationSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *OrganizationCreateBody) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
