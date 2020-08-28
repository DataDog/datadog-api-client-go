# SecurityMonitoringRuleQueryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctFields** | Pointer to **[]string** | Field for which the cardinality is measured. Sent as an array. | [optional] 
**GroupByFields** | Pointer to **[]string** | Fields to group by. | [optional] 
**Name** | Pointer to **string** | Name of the query. | [optional] 
**Query** | **string** | Query to run on logs. | 

## Methods

### NewSecurityMonitoringRuleQueryCreate

`func NewSecurityMonitoringRuleQueryCreate(query string, ) *SecurityMonitoringRuleQueryCreate`

NewSecurityMonitoringRuleQueryCreate instantiates a new SecurityMonitoringRuleQueryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleQueryCreateWithDefaults

`func NewSecurityMonitoringRuleQueryCreateWithDefaults() *SecurityMonitoringRuleQueryCreate`

NewSecurityMonitoringRuleQueryCreateWithDefaults instantiates a new SecurityMonitoringRuleQueryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctFields

`func (o *SecurityMonitoringRuleQueryCreate) GetDistinctFields() []string`

GetDistinctFields returns the DistinctFields field if non-nil, zero value otherwise.

### GetDistinctFieldsOk

`func (o *SecurityMonitoringRuleQueryCreate) GetDistinctFieldsOk() (*[]string, bool)`

GetDistinctFieldsOk returns a tuple with the DistinctFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctFields

`func (o *SecurityMonitoringRuleQueryCreate) SetDistinctFields(v []string)`

SetDistinctFields sets DistinctFields field to given value.

### HasDistinctFields

`func (o *SecurityMonitoringRuleQueryCreate) HasDistinctFields() bool`

HasDistinctFields returns a boolean if a field has been set.

### GetGroupByFields

`func (o *SecurityMonitoringRuleQueryCreate) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *SecurityMonitoringRuleQueryCreate) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *SecurityMonitoringRuleQueryCreate) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *SecurityMonitoringRuleQueryCreate) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleQueryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleQueryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleQueryCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleQueryCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityMonitoringRuleQueryCreate) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityMonitoringRuleQueryCreate) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityMonitoringRuleQueryCreate) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


