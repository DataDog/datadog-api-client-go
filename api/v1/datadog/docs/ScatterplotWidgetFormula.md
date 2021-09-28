# ScatterplotWidgetFormula

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Alias** | Pointer to **string** | Expression alias. | [optional] 
**Dimension** | [**ScatterplotDimension**](ScatterplotDimension.md) |  | 
**Formula** | **string** | String expression built from queries, formulas, and functions. | 

## Methods

### NewScatterplotWidgetFormula

`func NewScatterplotWidgetFormula(dimension ScatterplotDimension, formula string) *ScatterplotWidgetFormula`

NewScatterplotWidgetFormula instantiates a new ScatterplotWidgetFormula object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewScatterplotWidgetFormulaWithDefaults

`func NewScatterplotWidgetFormulaWithDefaults() *ScatterplotWidgetFormula`

NewScatterplotWidgetFormulaWithDefaults instantiates a new ScatterplotWidgetFormula object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAlias

`func (o *ScatterplotWidgetFormula) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ScatterplotWidgetFormula) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ScatterplotWidgetFormula) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ScatterplotWidgetFormula) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDimension

`func (o *ScatterplotWidgetFormula) GetDimension() ScatterplotDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ScatterplotWidgetFormula) GetDimensionOk() (*ScatterplotDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ScatterplotWidgetFormula) SetDimension(v ScatterplotDimension)`

SetDimension sets Dimension field to given value.


### GetFormula

`func (o *ScatterplotWidgetFormula) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *ScatterplotWidgetFormula) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *ScatterplotWidgetFormula) SetFormula(v string)`

SetFormula sets Formula field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


