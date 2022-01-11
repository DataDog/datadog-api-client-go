# SyntheticsCITestBody

## Properties

| Name      | Type                                                     | Description                 | Notes      |
| --------- | -------------------------------------------------------- | --------------------------- | ---------- |
| **Tests** | Pointer to [**[]SyntheticsCITest**](SyntheticsCITest.md) | Individual synthetics test. | [optional] |

## Methods

### NewSyntheticsCITestBody

`func NewSyntheticsCITestBody() *SyntheticsCITestBody`

NewSyntheticsCITestBody instantiates a new SyntheticsCITestBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsCITestBodyWithDefaults

`func NewSyntheticsCITestBodyWithDefaults() *SyntheticsCITestBody`

NewSyntheticsCITestBodyWithDefaults instantiates a new SyntheticsCITestBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTests

`func (o *SyntheticsCITestBody) GetTests() []SyntheticsCITest`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *SyntheticsCITestBody) GetTestsOk() (*[]SyntheticsCITest, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *SyntheticsCITestBody) SetTests(v []SyntheticsCITest)`

SetTests sets Tests field to given value.

### HasTests

`func (o *SyntheticsCITestBody) HasTests() bool`

HasTests returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
