# SunburstWidgetRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ApmQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**AuditQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**EventQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Formulas** | Pointer to [**[]WidgetFormula**](WidgetFormula.md) | List of formulas that operate on queries. **This feature is currently in beta.** | [optional] 
**LogQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**NetworkQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**ProcessQuery** | Pointer to [**ProcessQueryDefinition**](ProcessQueryDefinition.md) |  | [optional] 
**ProfileMetricsQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**Q** | Pointer to **string** | Widget query. | [optional] 
**Queries** | Pointer to [**[]FormulaAndFunctionQueryDefinition**](FormulaAndFunctionQueryDefinition.md) | List of queries that can be returned directly or used in formulas. **This feature is currently in beta.** | [optional] 
**ResponseFormat** | Pointer to [**FormulaAndFunctionResponseFormat**](FormulaAndFunctionResponseFormat.md) |  | [optional] 
**RumQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 
**SecurityQuery** | Pointer to [**LogQueryDefinition**](LogQueryDefinition.md) |  | [optional] 

## Methods

### NewSunburstWidgetRequest

`func NewSunburstWidgetRequest() *SunburstWidgetRequest`

NewSunburstWidgetRequest instantiates a new SunburstWidgetRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSunburstWidgetRequestWithDefaults

`func NewSunburstWidgetRequestWithDefaults() *SunburstWidgetRequest`

NewSunburstWidgetRequestWithDefaults instantiates a new SunburstWidgetRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApmQuery

`func (o *SunburstWidgetRequest) GetApmQuery() LogQueryDefinition`

GetApmQuery returns the ApmQuery field if non-nil, zero value otherwise.

### GetApmQueryOk

`func (o *SunburstWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool)`

GetApmQueryOk returns a tuple with the ApmQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmQuery

`func (o *SunburstWidgetRequest) SetApmQuery(v LogQueryDefinition)`

SetApmQuery sets ApmQuery field to given value.

### HasApmQuery

`func (o *SunburstWidgetRequest) HasApmQuery() bool`

HasApmQuery returns a boolean if a field has been set.

### GetAuditQuery

`func (o *SunburstWidgetRequest) GetAuditQuery() LogQueryDefinition`

GetAuditQuery returns the AuditQuery field if non-nil, zero value otherwise.

### GetAuditQueryOk

`func (o *SunburstWidgetRequest) GetAuditQueryOk() (*LogQueryDefinition, bool)`

GetAuditQueryOk returns a tuple with the AuditQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditQuery

`func (o *SunburstWidgetRequest) SetAuditQuery(v LogQueryDefinition)`

SetAuditQuery sets AuditQuery field to given value.

### HasAuditQuery

`func (o *SunburstWidgetRequest) HasAuditQuery() bool`

HasAuditQuery returns a boolean if a field has been set.

### GetEventQuery

`func (o *SunburstWidgetRequest) GetEventQuery() LogQueryDefinition`

GetEventQuery returns the EventQuery field if non-nil, zero value otherwise.

### GetEventQueryOk

`func (o *SunburstWidgetRequest) GetEventQueryOk() (*LogQueryDefinition, bool)`

GetEventQueryOk returns a tuple with the EventQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventQuery

`func (o *SunburstWidgetRequest) SetEventQuery(v LogQueryDefinition)`

SetEventQuery sets EventQuery field to given value.

### HasEventQuery

`func (o *SunburstWidgetRequest) HasEventQuery() bool`

HasEventQuery returns a boolean if a field has been set.

### GetFormulas

`func (o *SunburstWidgetRequest) GetFormulas() []WidgetFormula`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *SunburstWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *SunburstWidgetRequest) SetFormulas(v []WidgetFormula)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *SunburstWidgetRequest) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### GetLogQuery

`func (o *SunburstWidgetRequest) GetLogQuery() LogQueryDefinition`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *SunburstWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *SunburstWidgetRequest) SetLogQuery(v LogQueryDefinition)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *SunburstWidgetRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetNetworkQuery

`func (o *SunburstWidgetRequest) GetNetworkQuery() LogQueryDefinition`

GetNetworkQuery returns the NetworkQuery field if non-nil, zero value otherwise.

### GetNetworkQueryOk

`func (o *SunburstWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool)`

GetNetworkQueryOk returns a tuple with the NetworkQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuery

`func (o *SunburstWidgetRequest) SetNetworkQuery(v LogQueryDefinition)`

SetNetworkQuery sets NetworkQuery field to given value.

### HasNetworkQuery

`func (o *SunburstWidgetRequest) HasNetworkQuery() bool`

HasNetworkQuery returns a boolean if a field has been set.

### GetProcessQuery

`func (o *SunburstWidgetRequest) GetProcessQuery() ProcessQueryDefinition`

GetProcessQuery returns the ProcessQuery field if non-nil, zero value otherwise.

### GetProcessQueryOk

`func (o *SunburstWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool)`

GetProcessQueryOk returns a tuple with the ProcessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessQuery

`func (o *SunburstWidgetRequest) SetProcessQuery(v ProcessQueryDefinition)`

SetProcessQuery sets ProcessQuery field to given value.

### HasProcessQuery

`func (o *SunburstWidgetRequest) HasProcessQuery() bool`

HasProcessQuery returns a boolean if a field has been set.

### GetProfileMetricsQuery

`func (o *SunburstWidgetRequest) GetProfileMetricsQuery() LogQueryDefinition`

GetProfileMetricsQuery returns the ProfileMetricsQuery field if non-nil, zero value otherwise.

### GetProfileMetricsQueryOk

`func (o *SunburstWidgetRequest) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool)`

GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMetricsQuery

`func (o *SunburstWidgetRequest) SetProfileMetricsQuery(v LogQueryDefinition)`

SetProfileMetricsQuery sets ProfileMetricsQuery field to given value.

### HasProfileMetricsQuery

`func (o *SunburstWidgetRequest) HasProfileMetricsQuery() bool`

HasProfileMetricsQuery returns a boolean if a field has been set.

### GetQ

`func (o *SunburstWidgetRequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *SunburstWidgetRequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *SunburstWidgetRequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *SunburstWidgetRequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetQueries

`func (o *SunburstWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SunburstWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SunburstWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *SunburstWidgetRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetResponseFormat

`func (o *SunburstWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *SunburstWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *SunburstWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *SunburstWidgetRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetRumQuery

`func (o *SunburstWidgetRequest) GetRumQuery() LogQueryDefinition`

GetRumQuery returns the RumQuery field if non-nil, zero value otherwise.

### GetRumQueryOk

`func (o *SunburstWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool)`

GetRumQueryOk returns a tuple with the RumQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumQuery

`func (o *SunburstWidgetRequest) SetRumQuery(v LogQueryDefinition)`

SetRumQuery sets RumQuery field to given value.

### HasRumQuery

`func (o *SunburstWidgetRequest) HasRumQuery() bool`

HasRumQuery returns a boolean if a field has been set.

### GetSecurityQuery

`func (o *SunburstWidgetRequest) GetSecurityQuery() LogQueryDefinition`

GetSecurityQuery returns the SecurityQuery field if non-nil, zero value otherwise.

### GetSecurityQueryOk

`func (o *SunburstWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool)`

GetSecurityQueryOk returns a tuple with the SecurityQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuery

`func (o *SunburstWidgetRequest) SetSecurityQuery(v LogQueryDefinition)`

SetSecurityQuery sets SecurityQuery field to given value.

### HasSecurityQuery

`func (o *SunburstWidgetRequest) HasSecurityQuery() bool`

HasSecurityQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


