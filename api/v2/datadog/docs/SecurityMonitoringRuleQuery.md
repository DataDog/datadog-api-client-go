# SecurityMonitoringRuleQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByFields** | Pointer to **[]string** | Fields to group by. | [optional] 
**Name** | Pointer to **string** | Name of the query | [optional] 
**Query** | Pointer to **string** | Query to run on logs | [optional] 

## Methods

### NewSecurityMonitoringRuleQuery

`func NewSecurityMonitoringRuleQuery() *SecurityMonitoringRuleQuery`

NewSecurityMonitoringRuleQuery instantiates a new SecurityMonitoringRuleQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleQueryWithDefaults

`func NewSecurityMonitoringRuleQueryWithDefaults() *SecurityMonitoringRuleQuery`

NewSecurityMonitoringRuleQueryWithDefaults instantiates a new SecurityMonitoringRuleQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByFields

`func (o *SecurityMonitoringRuleQuery) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *SecurityMonitoringRuleQuery) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *SecurityMonitoringRuleQuery) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *SecurityMonitoringRuleQuery) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityMonitoringRuleQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityMonitoringRuleQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityMonitoringRuleQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SecurityMonitoringRuleQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


