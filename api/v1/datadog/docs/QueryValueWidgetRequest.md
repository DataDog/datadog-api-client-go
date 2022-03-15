# QueryValueWidgetRequest

## Properties

| Name                    | Type                                                                                       | Description                                                                                               | Notes      |
| ----------------------- | ------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------- | ---------- |
| **Aggregator**          | Pointer to [**WidgetAggregator**](WidgetAggregator.md)                                     |                                                                                                           | [optional] |
| **ApmQuery**            | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **AuditQuery**          | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **ConditionalFormats**  | Pointer to [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md)                     | List of conditional formats.                                                                              | [optional] |
| **EventQuery**          | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **Formulas**            | Pointer to [**[]WidgetFormula**](WidgetFormula.md)                                         | List of formulas that operate on queries. **This feature is currently in beta.**                          | [optional] |
| **LogQuery**            | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **NetworkQuery**        | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **ProcessQuery**        | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md)                         |                                                                                                           | [optional] |
| **ProfileMetricsQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **Q**                   | Pointer to **string**                                                                      | TODO.                                                                                                     | [optional] |
| **Queries**             | Pointer to [**[]FormulaAndFunctionQueryDefinition**](FormulaAndFunctionQueryDefinition.md) | List of queries that can be returned directly or used in formulas. **This feature is currently in beta.** | [optional] |
| **ResponseFormat**      | Pointer to [**FormulaAndFunctionResponseFormat**](FormulaAndFunctionResponseFormat.md)     |                                                                                                           | [optional] |
| **RumQuery**            | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |
| **SecurityQuery**       | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md)                                 |                                                                                                           | [optional] |

## Methods

### NewQueryValueWidgetRequest

`func NewQueryValueWidgetRequest() *QueryValueWidgetRequest`

NewQueryValueWidgetRequest instantiates a new QueryValueWidgetRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewQueryValueWidgetRequestWithDefaults

`func NewQueryValueWidgetRequestWithDefaults() *QueryValueWidgetRequest`

NewQueryValueWidgetRequestWithDefaults instantiates a new QueryValueWidgetRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregator

`func (o *QueryValueWidgetRequest) GetAggregator() WidgetAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *QueryValueWidgetRequest) GetAggregatorOk() (*WidgetAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *QueryValueWidgetRequest) SetAggregator(v WidgetAggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *QueryValueWidgetRequest) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetApmQuery

`func (o *QueryValueWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *QueryValueWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *QueryValueWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *QueryValueWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetAuditQuery

`func (o *QueryValueWidgetRequest) GetAuditQuery() LogQueryDefinition`

GetAuditQuery returns the AuditQuery field if non-nil, zero value otherwise.

### GetAuditQueryOk

`func (o *QueryValueWidgetRequest) GetAuditQueryOk() (*LogQueryDefinition, bool)`

GetAuditQueryOk returns a tuple with the AuditQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditQuery

`func (o *QueryValueWidgetRequest) SetAuditQuery(v LogQueryDefinition)`

SetAuditQuery sets AuditQuery field to given value.

### HasAuditQuery

`func (o *QueryValueWidgetRequest) HasAuditQuery() bool`

HasAuditQuery returns a boolean if a field has been set.

### GetConditionalFormats

`func (o *QueryValueWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *QueryValueWidgetRequest) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormats

`func (o *QueryValueWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat)`

SetConditionalFormats sets ConditionalFormats field to given value.

### HasConditionalFormats

`func (o *QueryValueWidgetRequest) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

### GetEventQuery

`func (o *QueryValueWidgetRequest) GetEventQuery() LogQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *QueryValueWidgetRequest) GetEventQueryOk() (*LogQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *QueryValueWidgetRequest) SetEventQuery(v LogQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *QueryValueWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetFormulas

`func (o *QueryValueWidgetRequest) GetFormulas() []WidgetFormula`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *QueryValueWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *QueryValueWidgetRequest) SetFormulas(v []WidgetFormula)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *QueryValueWidgetRequest) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### GetLogQuery

`func (o *QueryValueWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *QueryValueWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *QueryValueWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *QueryValueWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *QueryValueWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *QueryValueWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *QueryValueWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *QueryValueWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *QueryValueWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *QueryValueWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *QueryValueWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *QueryValueWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetProfileMetricsQuery

`func (o *QueryValueWidgetRequest) GetProfileMetricsQuery() LogQueryDefinition`

GetProfileMetricsQuery returns the ProfileMetricsQuery field if non-nil, zero value otherwise.

### GetProfileMetricsQueryOk

`func (o *QueryValueWidgetRequest) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool)`

GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMetricsQuery

`func (o *QueryValueWidgetRequest) SetProfileMetricsQuery(v LogQueryDefinition)`

SetProfileMetricsQuery sets ProfileMetricsQuery field to given value.

### HasProfileMetricsQuery

`func (o *QueryValueWidgetRequest) HasProfileMetricsQuery() bool`

HasProfileMetricsQuery returns a boolean if a field has been set.

### GetQ

`func (o *QueryValueWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *QueryValueWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *QueryValueWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *QueryValueWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetQueries

`func (o *QueryValueWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *QueryValueWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *QueryValueWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *QueryValueWidgetRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetResponseFormat

`func (o *QueryValueWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *QueryValueWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *QueryValueWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *QueryValueWidgetRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetRumQuery

`func (o *QueryValueWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *QueryValueWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *QueryValueWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *QueryValueWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetSecurityQuery

`func (o *QueryValueWidgetRequest) GetSecurityQuery() LogQueryDefinition`

GetSecurityQuery returns the SecurityQuery field if non-nil, zero value otherwise.

### GetSecurityQueryOk

`func (o *QueryValueWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool)`

GetSecurityQueryOk returns a tuple with the SecurityQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuery

`func (o *QueryValueWidgetRequest) SetSecurityQuery(v LogQueryDefinition)`

SetSecurityQuery sets SecurityQuery field to given value.

### HasSecurityQuery

`func (o *QueryValueWidgetRequest) HasSecurityQuery() bool`

HasSecurityQuery returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
