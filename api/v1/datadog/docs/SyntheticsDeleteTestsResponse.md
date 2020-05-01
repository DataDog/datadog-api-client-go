# SyntheticsDeleteTestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedTests** | Pointer to [**[]SyntheticsDeleteTestsResponseDeletedTests**](SyntheticsDeleteTestsResponse_deleted_tests.md) | Array of objects containing a deleted Synthetic test ID with the associated deletion timestamp. | [optional] 

## Methods

### NewSyntheticsDeleteTestsResponse

`func NewSyntheticsDeleteTestsResponse() *SyntheticsDeleteTestsResponse`

NewSyntheticsDeleteTestsResponse instantiates a new SyntheticsDeleteTestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsDeleteTestsResponseWithDefaults

`func NewSyntheticsDeleteTestsResponseWithDefaults() *SyntheticsDeleteTestsResponse`

NewSyntheticsDeleteTestsResponseWithDefaults instantiates a new SyntheticsDeleteTestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedTests

`func (o *SyntheticsDeleteTestsResponse) GetDeletedTests() []SyntheticsDeleteTestsResponseDeletedTests`

GetDeletedTests returns the DeletedTests field if non-nil, zero value otherwise.

### GetDeletedTestsOk

`func (o *SyntheticsDeleteTestsResponse) GetDeletedTestsOk() (*[]SyntheticsDeleteTestsResponseDeletedTests, bool)`

GetDeletedTestsOk returns a tuple with the DeletedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTests

`func (o *SyntheticsDeleteTestsResponse) SetDeletedTests(v []SyntheticsDeleteTestsResponseDeletedTests)`

SetDeletedTests sets DeletedTests field to given value.

### HasDeletedTests

`func (o *SyntheticsDeleteTestsResponse) HasDeletedTests() bool`

HasDeletedTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


