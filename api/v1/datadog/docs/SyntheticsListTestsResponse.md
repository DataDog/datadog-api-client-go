# SyntheticsListTestsResponse

## Properties

| Name      | Type                                                               | Description                             | Notes      |
| --------- | ------------------------------------------------------------------ | --------------------------------------- | ---------- |
| **Tests** | Pointer to [**[]SyntheticsTestDetails**](SyntheticsTestDetails.md) | Array of Synthetic tests configuration. | [optional] |

## Methods

### NewSyntheticsListTestsResponse

`func NewSyntheticsListTestsResponse() *SyntheticsListTestsResponse`

NewSyntheticsListTestsResponse instantiates a new SyntheticsListTestsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsListTestsResponseWithDefaults

`func NewSyntheticsListTestsResponseWithDefaults() *SyntheticsListTestsResponse`

NewSyntheticsListTestsResponseWithDefaults instantiates a new SyntheticsListTestsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTests

`func (o *SyntheticsListTestsResponse) GetTests() []SyntheticsTestDetails`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *SyntheticsListTestsResponse) GetTestsOk() (*[]SyntheticsTestDetails, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *SyntheticsListTestsResponse) SetTests(v []SyntheticsTestDetails)`

SetTests sets Tests field to given value.

### HasTests

`func (o *SyntheticsListTestsResponse) HasTests() bool`

HasTests returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
