# DashboardTemplateVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **NullableString** | The default value for the template variable on dashboard load. | [optional] 
**Name** | Pointer to **string** | The name of the variable. | 
**Prefix** | Pointer to **NullableString** | The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down. | [optional] 

## Methods

### NewDashboardTemplateVariables

`func NewDashboardTemplateVariables(name string, ) *DashboardTemplateVariables`

NewDashboardTemplateVariables instantiates a new DashboardTemplateVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTemplateVariablesWithDefaults

`func NewDashboardTemplateVariablesWithDefaults() *DashboardTemplateVariables`

NewDashboardTemplateVariablesWithDefaults instantiates a new DashboardTemplateVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *DashboardTemplateVariables) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DashboardTemplateVariables) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DashboardTemplateVariables) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DashboardTemplateVariables) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *DashboardTemplateVariables) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *DashboardTemplateVariables) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetName

`func (o *DashboardTemplateVariables) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardTemplateVariables) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardTemplateVariables) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *DashboardTemplateVariables) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DashboardTemplateVariables) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DashboardTemplateVariables) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DashboardTemplateVariables) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *DashboardTemplateVariables) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *DashboardTemplateVariables) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


