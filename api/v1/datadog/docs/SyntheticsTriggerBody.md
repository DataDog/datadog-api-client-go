# SyntheticsTriggerBody

## Properties

| Name      | Type                                                    | Description                 | Notes |
| --------- | ------------------------------------------------------- | --------------------------- | ----- |
| **Tests** | [**[]SyntheticsTriggerTest**](SyntheticsTriggerTest.md) | Individual synthetics test. |

## Methods

### NewSyntheticsTriggerBody

`func NewSyntheticsTriggerBody(tests []SyntheticsTriggerTest) *SyntheticsTriggerBody`

NewSyntheticsTriggerBody instantiates a new SyntheticsTriggerBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTriggerBodyWithDefaults

`func NewSyntheticsTriggerBodyWithDefaults() *SyntheticsTriggerBody`

NewSyntheticsTriggerBodyWithDefaults instantiates a new SyntheticsTriggerBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTests

`func (o *SyntheticsTriggerBody) GetTests() []SyntheticsTriggerTest`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *SyntheticsTriggerBody) GetTestsOk() (*[]SyntheticsTriggerTest, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *SyntheticsTriggerBody) SetTests(v []SyntheticsTriggerTest)`

SetTests sets Tests field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
