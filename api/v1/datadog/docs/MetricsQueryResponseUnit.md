# MetricsQueryResponseUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | Unit family, allows for conversion between units of the same family, for scaling. | [optional] [readonly] 
**Name** | Pointer to **string** | Unit name | [optional] [readonly] 
**Plural** | Pointer to **string** | Plural form of the unit name | [optional] [readonly] 
**ScaleFactor** | Pointer to **float64** | Factor for scaling between units of the same family. | [optional] [readonly] 
**ShortName** | Pointer to **string** | Abbreviation of the unit | [optional] [readonly] 

## Methods

### GetFamily

`func (o *MetricsQueryResponseUnit) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *MetricsQueryResponseUnit) GetFamilyOk() (string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFamily

`func (o *MetricsQueryResponseUnit) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamily

`func (o *MetricsQueryResponseUnit) SetFamily(v string)`

SetFamily gets a reference to the given string and assigns it to the Family field.

### GetName

`func (o *MetricsQueryResponseUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsQueryResponseUnit) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MetricsQueryResponseUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MetricsQueryResponseUnit) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPlural

`func (o *MetricsQueryResponseUnit) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *MetricsQueryResponseUnit) GetPluralOk() (string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlural

`func (o *MetricsQueryResponseUnit) HasPlural() bool`

HasPlural returns a boolean if a field has been set.

### SetPlural

`func (o *MetricsQueryResponseUnit) SetPlural(v string)`

SetPlural gets a reference to the given string and assigns it to the Plural field.

### GetScaleFactor

`func (o *MetricsQueryResponseUnit) GetScaleFactor() float64`

GetScaleFactor returns the ScaleFactor field if non-nil, zero value otherwise.

### GetScaleFactorOk

`func (o *MetricsQueryResponseUnit) GetScaleFactorOk() (float64, bool)`

GetScaleFactorOk returns a tuple with the ScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScaleFactor

`func (o *MetricsQueryResponseUnit) HasScaleFactor() bool`

HasScaleFactor returns a boolean if a field has been set.

### SetScaleFactor

`func (o *MetricsQueryResponseUnit) SetScaleFactor(v float64)`

SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.

### GetShortName

`func (o *MetricsQueryResponseUnit) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *MetricsQueryResponseUnit) GetShortNameOk() (string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShortName

`func (o *MetricsQueryResponseUnit) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortName

`func (o *MetricsQueryResponseUnit) SetShortName(v string)`

SetShortName gets a reference to the given string and assigns it to the ShortName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


