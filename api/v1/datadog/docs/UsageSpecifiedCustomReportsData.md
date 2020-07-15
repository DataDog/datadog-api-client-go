# UsageSpecifiedCustomReportsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UsageSpecifiedCustomReportsAttributes**](UsageSpecifiedCustomReportsAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The date for specified custom reports. | [optional] 
**Type** | Pointer to [**UsageReportsType**](UsageReportsType.md) |  | [optional] [default to "reports"]

## Methods

### NewUsageSpecifiedCustomReportsData

`func NewUsageSpecifiedCustomReportsData() *UsageSpecifiedCustomReportsData`

NewUsageSpecifiedCustomReportsData instantiates a new UsageSpecifiedCustomReportsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSpecifiedCustomReportsDataWithDefaults

`func NewUsageSpecifiedCustomReportsDataWithDefaults() *UsageSpecifiedCustomReportsData`

NewUsageSpecifiedCustomReportsDataWithDefaults instantiates a new UsageSpecifiedCustomReportsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UsageSpecifiedCustomReportsData) GetAttributes() UsageSpecifiedCustomReportsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UsageSpecifiedCustomReportsData) GetAttributesOk() (*UsageSpecifiedCustomReportsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UsageSpecifiedCustomReportsData) SetAttributes(v UsageSpecifiedCustomReportsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UsageSpecifiedCustomReportsData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UsageSpecifiedCustomReportsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageSpecifiedCustomReportsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageSpecifiedCustomReportsData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageSpecifiedCustomReportsData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UsageSpecifiedCustomReportsData) GetType() UsageReportsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsageSpecifiedCustomReportsData) GetTypeOk() (*UsageReportsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsageSpecifiedCustomReportsData) SetType(v UsageReportsType)`

SetType sets Type field to given value.

### HasType

`func (o *UsageSpecifiedCustomReportsData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


