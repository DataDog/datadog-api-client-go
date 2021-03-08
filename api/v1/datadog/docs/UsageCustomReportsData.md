# UsageCustomReportsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UsageCustomReportsAttributes**](UsageCustomReportsAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The date for specified custom reports. | [optional] 
**Type** | Pointer to [**UsageReportsType**](UsageReportsType.md) |  | [optional] [default to USAGEREPORTSTYPE_REPORTS]

## Methods

### NewUsageCustomReportsData

`func NewUsageCustomReportsData() *UsageCustomReportsData`

NewUsageCustomReportsData instantiates a new UsageCustomReportsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageCustomReportsDataWithDefaults

`func NewUsageCustomReportsDataWithDefaults() *UsageCustomReportsData`

NewUsageCustomReportsDataWithDefaults instantiates a new UsageCustomReportsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UsageCustomReportsData) GetAttributes() UsageCustomReportsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UsageCustomReportsData) GetAttributesOk() (*UsageCustomReportsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UsageCustomReportsData) SetAttributes(v UsageCustomReportsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UsageCustomReportsData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UsageCustomReportsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageCustomReportsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageCustomReportsData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageCustomReportsData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UsageCustomReportsData) GetType() UsageReportsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsageCustomReportsData) GetTypeOk() (*UsageReportsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsageCustomReportsData) SetType(v UsageReportsType)`

SetType sets Type field to given value.

### HasType

`func (o *UsageCustomReportsData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


