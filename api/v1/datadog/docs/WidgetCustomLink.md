# WidgetCustomLink

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Label** | **string** | The label for the custom link URL. Keep the label short and descriptive. Use metrics and tags as variables. | 
**Link** | **string** | The URL of the custom link. URL must include &#x60;http&#x60; or &#x60;https&#x60;. A relative URL must start with &#x60;/&#x60;. | 

## Methods

### NewWidgetCustomLink

`func NewWidgetCustomLink(label string, link string, ) *WidgetCustomLink`

NewWidgetCustomLink instantiates a new WidgetCustomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetCustomLinkWithDefaults

`func NewWidgetCustomLinkWithDefaults() *WidgetCustomLink`

NewWidgetCustomLinkWithDefaults instantiates a new WidgetCustomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


