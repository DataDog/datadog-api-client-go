# DashboardTemplateVariablePreset

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Name** | Pointer to **string** | The name of the variable. | [optional] 
**TemplateVariables** | Pointer to [**[]DashboardTemplateVariablePresetValue**](DashboardTemplateVariablePresetValue.md) | List of variables. | [optional] 

## Methods

### NewDashboardTemplateVariablePreset

`func NewDashboardTemplateVariablePreset() *DashboardTemplateVariablePreset`

NewDashboardTemplateVariablePreset instantiates a new DashboardTemplateVariablePreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTemplateVariablePresetWithDefaults

`func NewDashboardTemplateVariablePresetWithDefaults() *DashboardTemplateVariablePreset`

NewDashboardTemplateVariablePresetWithDefaults instantiates a new DashboardTemplateVariablePreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DashboardTemplateVariablePreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardTemplateVariablePreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardTemplateVariablePreset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DashboardTemplateVariablePreset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateVariables

`func (o *DashboardTemplateVariablePreset) GetTemplateVariables() []DashboardTemplateVariablePresetValue`

GetTemplateVariables returns the TemplateVariables field if non-nil, zero value otherwise.

### GetTemplateVariablesOk

`func (o *DashboardTemplateVariablePreset) GetTemplateVariablesOk() (*[]DashboardTemplateVariablePresetValue, bool)`

GetTemplateVariablesOk returns a tuple with the TemplateVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVariables

`func (o *DashboardTemplateVariablePreset) SetTemplateVariables(v []DashboardTemplateVariablePresetValue)`

SetTemplateVariables sets TemplateVariables field to given value.

### HasTemplateVariables

`func (o *DashboardTemplateVariablePreset) HasTemplateVariables() bool`

HasTemplateVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


