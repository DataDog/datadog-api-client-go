# FunnelStep

## Properties

| Name      | Type       | Description            | Notes |
| --------- | ---------- | ---------------------- | ----- |
| **Facet** | **string** | The facet of the step. |
| **Value** | **string** | The value of the step. |

## Methods

### NewFunnelStep

`func NewFunnelStep(facet string, value string) *FunnelStep`

NewFunnelStep instantiates a new FunnelStep object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFunnelStepWithDefaults

`func NewFunnelStepWithDefaults() *FunnelStep`

NewFunnelStepWithDefaults instantiates a new FunnelStep object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFacet

`func (o *FunnelStep) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *FunnelStep) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *FunnelStep) SetFacet(v string)`

SetFacet sets Facet field to given value.

### GetValue

`func (o *FunnelStep) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FunnelStep) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FunnelStep) SetValue(v string)`

SetValue sets Value field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
