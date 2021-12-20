# SLOThreshold

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Target** | **float64** | The target value for the service level indicator within the corresponding timeframe. | 
**TargetDisplay** | Pointer to **string** | A string representation of the target that indicates its precision. It uses trailing zeros to show significant decimal places (for example &#x60;98.00&#x60;).  Always included in service level objective responses. Ignored in create/update requests. | [optional] 
**Timeframe** | [**SLOTimeframe**](SLOTimeframe.md) |  | 
**Warning** | Pointer to **float64** | The warning value for the service level objective. | [optional] 
**WarningDisplay** | Pointer to **string** | A string representation of the warning target (see the description of the &#x60;target_display&#x60; field for details).  Included in service level objective responses if a warning target exists. Ignored in create/update requests. | [optional] 

## Methods

### NewSLOThreshold

`func NewSLOThreshold(target float64, timeframe SLOTimeframe) *SLOThreshold`

NewSLOThreshold instantiates a new SLOThreshold object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOThresholdWithDefaults

`func NewSLOThresholdWithDefaults() *SLOThreshold`

NewSLOThresholdWithDefaults instantiates a new SLOThreshold object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTarget

`func (o *SLOThreshold) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SLOThreshold) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SLOThreshold) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetTargetDisplay

`func (o *SLOThreshold) GetTargetDisplay() string`

GetTargetDisplay returns the TargetDisplay field if non-nil, zero value otherwise.

### GetTargetDisplayOk

`func (o *SLOThreshold) GetTargetDisplayOk() (*string, bool)`

GetTargetDisplayOk returns a tuple with the TargetDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDisplay

`func (o *SLOThreshold) SetTargetDisplay(v string)`

SetTargetDisplay sets TargetDisplay field to given value.

### HasTargetDisplay

`func (o *SLOThreshold) HasTargetDisplay() bool`

HasTargetDisplay returns a boolean if a field has been set.

### GetTimeframe

`func (o *SLOThreshold) GetTimeframe() SLOTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *SLOThreshold) GetTimeframeOk() (*SLOTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *SLOThreshold) SetTimeframe(v SLOTimeframe)`

SetTimeframe sets Timeframe field to given value.


### GetWarning

`func (o *SLOThreshold) GetWarning() float64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SLOThreshold) GetWarningOk() (*float64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SLOThreshold) SetWarning(v float64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SLOThreshold) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetWarningDisplay

`func (o *SLOThreshold) GetWarningDisplay() string`

GetWarningDisplay returns the WarningDisplay field if non-nil, zero value otherwise.

### GetWarningDisplayOk

`func (o *SLOThreshold) GetWarningDisplayOk() (*string, bool)`

GetWarningDisplayOk returns a tuple with the WarningDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDisplay

`func (o *SLOThreshold) SetWarningDisplay(v string)`

SetWarningDisplay sets WarningDisplay field to given value.

### HasWarningDisplay

`func (o *SLOThreshold) HasWarningDisplay() bool`

HasWarningDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


