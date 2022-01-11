# SyntheticsDeletedTest

## Properties

| Name          | Type                     | Description                                  | Notes      |
| ------------- | ------------------------ | -------------------------------------------- | ---------- |
| **DeletedAt** | Pointer to **time.Time** | Deletion timestamp of the Synthetic test ID. | [optional] |
| **PublicId**  | Pointer to **string**    | The Synthetic test ID deleted.               | [optional] |

## Methods

### NewSyntheticsDeletedTest

`func NewSyntheticsDeletedTest() *SyntheticsDeletedTest`

NewSyntheticsDeletedTest instantiates a new SyntheticsDeletedTest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsDeletedTestWithDefaults

`func NewSyntheticsDeletedTestWithDefaults() *SyntheticsDeletedTest`

NewSyntheticsDeletedTestWithDefaults instantiates a new SyntheticsDeletedTest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDeletedAt

`func (o *SyntheticsDeletedTest) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SyntheticsDeletedTest) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SyntheticsDeletedTest) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SyntheticsDeletedTest) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsDeletedTest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsDeletedTest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsDeletedTest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsDeletedTest) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
