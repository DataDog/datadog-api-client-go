# DashboardTemplateVariable

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AvailableValues** | Pointer to **[]string** | The list of values that the template variable drop-down is limited to. | [optional] 
**Default** | Pointer to **NullableString** | The default value for the template variable on dashboard load. | [optional] 
**Name** | **string** | The name of the variable. | 
**Prefix** | Pointer to **NullableString** | The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down. | [optional] 

## Methods

### NewDashboardTemplateVariable

`func NewDashboardTemplateVariable(name string) *DashboardTemplateVariable`

NewDashboardTemplateVariable instantiates a new DashboardTemplateVariable object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardTemplateVariableWithDefaults

`func NewDashboardTemplateVariableWithDefaults() *DashboardTemplateVariable`

NewDashboardTemplateVariableWithDefaults instantiates a new DashboardTemplateVariable object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAvailableValues

`func (o *DashboardTemplateVariable) GetAvailableValues() []string`

GetAvailableValues returns the AvailableValues field if non-nil, zero value otherwise.

### GetAvailableValuesOk

`func (o *DashboardTemplateVariable) GetAvailableValuesOk() (*[]string, bool)`

GetAvailableValuesOk returns a tuple with the AvailableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableValues

`func (o *DashboardTemplateVariable) SetAvailableValues(v []string)`

SetAvailableValues sets AvailableValues field to given value.

### HasAvailableValues

`func (o *DashboardTemplateVariable) HasAvailableValues() bool`

HasAvailableValues returns a boolean if a field has been set.

### SetAvailableValuesNil

`func (o *DashboardTemplateVariable) SetAvailableValuesNil(b bool)`

 SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil

### UnsetAvailableValues
`func (o *DashboardTemplateVariable) UnsetAvailableValues()`

UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil
### GetDefault

`func (o *DashboardTemplateVariable) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DashboardTemplateVariable) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DashboardTemplateVariable) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DashboardTemplateVariable) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *DashboardTemplateVariable) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *DashboardTemplateVariable) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetName

`func (o *DashboardTemplateVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardTemplateVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardTemplateVariable) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *DashboardTemplateVariable) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DashboardTemplateVariable) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DashboardTemplateVariable) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DashboardTemplateVariable) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *DashboardTemplateVariable) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *DashboardTemplateVariable) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


