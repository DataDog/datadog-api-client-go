# ScatterplotTableRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Formulas** | Pointer to [**[]ScatterplotWidgetFormula**](ScatterplotWidgetFormula.md) | List of Scatterplot formulas that operate on queries. **This feature is currently in beta.** | [optional] 
**Queries** | Pointer to [**[]FormulaAndFunctionQueryDefinition**](FormulaAndFunctionQueryDefinition.md) | List of queries that can be returned directly or used in formulas. **This feature is currently in beta.** | [optional] 
**ResponseFormat** | Pointer to [**FormulaAndFunctionResponseFormat**](FormulaAndFunctionResponseFormat.md) |  | [optional] 

## Methods

### NewScatterplotTableRequest

`func NewScatterplotTableRequest() *ScatterplotTableRequest`

NewScatterplotTableRequest instantiates a new ScatterplotTableRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewScatterplotTableRequestWithDefaults

`func NewScatterplotTableRequestWithDefaults() *ScatterplotTableRequest`

NewScatterplotTableRequestWithDefaults instantiates a new ScatterplotTableRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFormulas

`func (o *ScatterplotTableRequest) GetFormulas() []ScatterplotWidgetFormula`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *ScatterplotTableRequest) GetFormulasOk() (*[]ScatterplotWidgetFormula, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *ScatterplotTableRequest) SetFormulas(v []ScatterplotWidgetFormula)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *ScatterplotTableRequest) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### GetQueries

`func (o *ScatterplotTableRequest) GetQueries() []FormulaAndFunctionQueryDefinition`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *ScatterplotTableRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *ScatterplotTableRequest) SetQueries(v []FormulaAndFunctionQueryDefinition)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *ScatterplotTableRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetResponseFormat

`func (o *ScatterplotTableRequest) GetResponseFormat() FormulaAndFunctionResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ScatterplotTableRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ScatterplotTableRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ScatterplotTableRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


