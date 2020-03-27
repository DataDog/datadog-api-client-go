# Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to [**WidgetDefinition**](WidgetDefinition.md) |  | 
**Id** | Pointer to **int64** | ID of the widget | [optional] 
**Layout** | Pointer to [**WidgetLayout**](WidgetLayout.md) |  | [optional] 

## Methods

### NewWidget

`func NewWidget(definition WidgetDefinition, ) *Widget`

NewWidget instantiates a new Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetWithDefaults

`func NewWidgetWithDefaults() *Widget`

NewWidgetWithDefaults instantiates a new Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *Widget) GetDefinition() WidgetDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Widget) GetDefinitionOk() (*WidgetDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Widget) SetDefinition(v WidgetDefinition)`

SetDefinition sets Definition field to given value.


### GetId

`func (o *Widget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Widget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Widget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Widget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLayout

`func (o *Widget) GetLayout() WidgetLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Widget) GetLayoutOk() (*WidgetLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Widget) SetLayout(v WidgetLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Widget) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


