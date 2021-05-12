# DistributionWidgetXAxis

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IncludeZero** | Pointer to **bool** | True includes zero. | [optional] 
**Max** | Pointer to **string** | Specifies maximum value to show on the x-axis. It takes a number, percentile (p90 &#x3D;&#x3D;&#x3D; 90th percentile), or auto for default behavior. | [optional] [default to "auto"]
**Min** | Pointer to **string** | Specifies minimum value to show on the x-axis. It takes a number, percentile (p90 &#x3D;&#x3D;&#x3D; 90th percentile), or auto for default behavior. | [optional] [default to "auto"]
**Scale** | Pointer to **string** | Specifies the scale type. Possible values are &#x60;linear&#x60;. | [optional] [default to "linear"]

## Methods

### NewDistributionWidgetXAxis

`func NewDistributionWidgetXAxis() *DistributionWidgetXAxis`

NewDistributionWidgetXAxis instantiates a new DistributionWidgetXAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWidgetXAxisWithDefaults

`func NewDistributionWidgetXAxisWithDefaults() *DistributionWidgetXAxis`

NewDistributionWidgetXAxisWithDefaults instantiates a new DistributionWidgetXAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeZero

`func (o *DistributionWidgetXAxis) GetIncludeZero() bool`

GetIncludeZero returns the IncludeZero field if non-nil, zero value otherwise.

### GetIncludeZeroOk

`func (o *DistributionWidgetXAxis) GetIncludeZeroOk() (*bool, bool)`

GetIncludeZeroOk returns a tuple with the IncludeZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZero

`func (o *DistributionWidgetXAxis) SetIncludeZero(v bool)`

SetIncludeZero sets IncludeZero field to given value.

### HasIncludeZero

`func (o *DistributionWidgetXAxis) HasIncludeZero() bool`

HasIncludeZero returns a boolean if a field has been set.

### GetMax

`func (o *DistributionWidgetXAxis) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DistributionWidgetXAxis) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DistributionWidgetXAxis) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *DistributionWidgetXAxis) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *DistributionWidgetXAxis) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DistributionWidgetXAxis) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DistributionWidgetXAxis) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *DistributionWidgetXAxis) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetScale

`func (o *DistributionWidgetXAxis) GetScale() string`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *DistributionWidgetXAxis) GetScaleOk() (*string, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *DistributionWidgetXAxis) SetScale(v string)`

SetScale sets Scale field to given value.

### HasScale

`func (o *DistributionWidgetXAxis) HasScale() bool`

HasScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


