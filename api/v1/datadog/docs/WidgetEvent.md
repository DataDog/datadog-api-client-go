# WidgetEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Q** | Pointer to **string** |  | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or | [optional] 

## Methods

### NewWidgetEvent

`func NewWidgetEvent(q string, ) *WidgetEvent`

NewWidgetEvent instantiates a new WidgetEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetEventWithDefaults

`func NewWidgetEventWithDefaults() *WidgetEvent`

NewWidgetEventWithDefaults instantiates a new WidgetEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQ

`func (o *WidgetEvent) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *WidgetEvent) GetQOk() (string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQ

`func (o *WidgetEvent) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQ

`func (o *WidgetEvent) SetQ(v string)`

SetQ gets a reference to the given string and assigns it to the Q field.

### GetTagsExecution

`func (o *WidgetEvent) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *WidgetEvent) GetTagsExecutionOk() (string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagsExecution

`func (o *WidgetEvent) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### SetTagsExecution

`func (o *WidgetEvent) SetTagsExecution(v string)`

SetTagsExecution gets a reference to the given string and assigns it to the TagsExecution field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


