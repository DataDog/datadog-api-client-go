# WidgetCustomLink

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsHidden** | Pointer to **bool** | The flag for toggling context menu link visibility. | [optional] 
**Label** | Pointer to **string** | The label for the custom link URL. Keep the label short and descriptive. Use metrics and tags as variables. | [optional] 
**Link** | Pointer to **string** | The URL of the custom link. URL must include &#x60;http&#x60; or &#x60;https&#x60;. A relative URL must start with &#x60;/&#x60;. | [optional] 
**OverrideLabel** | Pointer to **string** | The label ID that refers to a context menu link. Can be &#x60;logs&#x60;, &#x60;hosts&#x60;, &#x60;traces&#x60;, &#x60;profiles&#x60;, &#x60;processes&#x60;, &#x60;containers&#x60;, or &#x60;rum&#x60;. | [optional] 

## Methods

### NewWidgetCustomLink

`func NewWidgetCustomLink() *WidgetCustomLink`

NewWidgetCustomLink instantiates a new WidgetCustomLink object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWidgetCustomLinkWithDefaults

`func NewWidgetCustomLinkWithDefaults() *WidgetCustomLink`

NewWidgetCustomLinkWithDefaults instantiates a new WidgetCustomLink object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsHidden

`func (o *WidgetCustomLink) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *WidgetCustomLink) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *WidgetCustomLink) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *WidgetCustomLink) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLabel

`func (o *WidgetCustomLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WidgetCustomLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WidgetCustomLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WidgetCustomLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLink

`func (o *WidgetCustomLink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WidgetCustomLink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WidgetCustomLink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *WidgetCustomLink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetOverrideLabel

`func (o *WidgetCustomLink) GetOverrideLabel() string`

GetOverrideLabel returns the OverrideLabel field if non-nil, zero value otherwise.

### GetOverrideLabelOk

`func (o *WidgetCustomLink) GetOverrideLabelOk() (*string, bool)`

GetOverrideLabelOk returns a tuple with the OverrideLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLabel

`func (o *WidgetCustomLink) SetOverrideLabel(v string)`

SetOverrideLabel sets OverrideLabel field to given value.

### HasOverrideLabel

`func (o *WidgetCustomLink) HasOverrideLabel() bool`

HasOverrideLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


