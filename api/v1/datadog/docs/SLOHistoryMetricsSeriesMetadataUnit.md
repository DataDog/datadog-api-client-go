# SLOHistoryMetricsSeriesMetadataUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | The family of metric unit, for example &#x60;bytes&#x60; is the family for &#x60;kibibyte&#x60;, &#x60;byte&#x60;, and &#x60;bit&#x60; units. | [optional] 
**Id** | Pointer to **int64** | The ID of the metric unit. | [optional] 
**Name** | Pointer to **string** | The unit of the metric, for instance &#x60;byte&#x60;. | [optional] 
**Plural** | Pointer to **NullableString** | The plural Unit of metric, for instance &#x60;bytes&#x60;. | [optional] 
**ScaleFactor** | Pointer to **float64** | The scale factor of metric unit, for instance &#x60;1.0&#x60;. | [optional] 
**ShortName** | Pointer to **NullableString** | A shorter and abbreviated version of the metric unit, for instance &#x60;B&#x60;. | [optional] 

## Methods

### NewSLOHistoryMetricsSeriesMetadataUnit

`func NewSLOHistoryMetricsSeriesMetadataUnit() *SLOHistoryMetricsSeriesMetadataUnit`

NewSLOHistoryMetricsSeriesMetadataUnit instantiates a new SLOHistoryMetricsSeriesMetadataUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOHistoryMetricsSeriesMetadataUnitWithDefaults

`func NewSLOHistoryMetricsSeriesMetadataUnitWithDefaults() *SLOHistoryMetricsSeriesMetadataUnit`

NewSLOHistoryMetricsSeriesMetadataUnitWithDefaults instantiates a new SLOHistoryMetricsSeriesMetadataUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetId

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlural

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetPluralOk() (*string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlural

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetPlural(v string)`

SetPlural sets Plural field to given value.

### HasPlural

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasPlural() bool`

HasPlural returns a boolean if a field has been set.

### SetPluralNil

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetPluralNil(b bool)`

 SetPluralNil sets the value for Plural to be an explicit nil

### UnsetPlural
`func (o *SLOHistoryMetricsSeriesMetadataUnit) UnsetPlural()`

UnsetPlural ensures that no value is present for Plural, not even an explicit nil
### GetScaleFactor

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetScaleFactor() float64`

GetScaleFactor returns the ScaleFactor field if non-nil, zero value otherwise.

### GetScaleFactorOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetScaleFactorOk() (*float64, bool)`

GetScaleFactorOk returns a tuple with the ScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleFactor

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetScaleFactor(v float64)`

SetScaleFactor sets ScaleFactor field to given value.

### HasScaleFactor

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasScaleFactor() bool`

HasScaleFactor returns a boolean if a field has been set.

### GetShortName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *SLOHistoryMetricsSeriesMetadataUnit) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *SLOHistoryMetricsSeriesMetadataUnit) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *SLOHistoryMetricsSeriesMetadataUnit) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *SLOHistoryMetricsSeriesMetadataUnit) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


