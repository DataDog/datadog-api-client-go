# TreeMapWidgetRequest

## Properties

| Name               | Type                                                                                       | Description                                                                                               | Notes      |
| ------------------ | ------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------- | ---------- |
| **Formulas**       | Pointer to [**[]WidgetFormula**](WidgetFormula.md)                                         | List of formulas that operate on queries. **This feature is currently in beta.**                          | [optional] |
| **Q**              | Pointer to **string**                                                                      | The widget metrics query.                                                                                 | [optional] |
| **Queries**        | Pointer to [**[]FormulaAndFunctionQueryDefinition**](FormulaAndFunctionQueryDefinition.md) | List of queries that can be returned directly or used in formulas. **This feature is currently in beta.** | [optional] |
| **ResponseFormat** | Pointer to [**FormulaAndFunctionResponseFormat**](FormulaAndFunctionResponseFormat.md)     |                                                                                                           | [optional] |

## Methods

### NewTreeMapWidgetRequest

`func NewTreeMapWidgetRequest() *TreeMapWidgetRequest`

NewTreeMapWidgetRequest instantiates a new TreeMapWidgetRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewTreeMapWidgetRequestWithDefaults

`func NewTreeMapWidgetRequestWithDefaults() *TreeMapWidgetRequest`

NewTreeMapWidgetRequestWithDefaults instantiates a new TreeMapWidgetRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFormulas

`func (o *TreeMapWidgetRequest) GetFormulas() []WidgetFormula`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *TreeMapWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *TreeMapWidgetRequest) SetFormulas(v []WidgetFormula)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *TreeMapWidgetRequest) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### GetQ

`func (o *TreeMapWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *TreeMapWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *TreeMapWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *TreeMapWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetQueries

`func (o *TreeMapWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TreeMapWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TreeMapWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TreeMapWidgetRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetResponseFormat

`func (o *TreeMapWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *TreeMapWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *TreeMapWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *TreeMapWidgetRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
