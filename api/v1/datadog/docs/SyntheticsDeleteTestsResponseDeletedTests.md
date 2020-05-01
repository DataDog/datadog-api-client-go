# SyntheticsDeleteTestsResponseDeletedTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Deletion timestamp of the Synthetic test ID. | [optional] 
**PublicId** | Pointer to **string** | The Synthetic test ID deleted. | [optional] 

## Methods

### NewSyntheticsDeleteTestsResponseDeletedTests

`func NewSyntheticsDeleteTestsResponseDeletedTests() *SyntheticsDeleteTestsResponseDeletedTests`

NewSyntheticsDeleteTestsResponseDeletedTests instantiates a new SyntheticsDeleteTestsResponseDeletedTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsDeleteTestsResponseDeletedTestsWithDefaults

`func NewSyntheticsDeleteTestsResponseDeletedTestsWithDefaults() *SyntheticsDeleteTestsResponseDeletedTests`

NewSyntheticsDeleteTestsResponseDeletedTestsWithDefaults instantiates a new SyntheticsDeleteTestsResponseDeletedTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedAt

`func (o *SyntheticsDeleteTestsResponseDeletedTests) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SyntheticsDeleteTestsResponseDeletedTests) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SyntheticsDeleteTestsResponseDeletedTests) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SyntheticsDeleteTestsResponseDeletedTests) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsDeleteTestsResponseDeletedTests) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsDeleteTestsResponseDeletedTests) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsDeleteTestsResponseDeletedTests) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsDeleteTestsResponseDeletedTests) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


