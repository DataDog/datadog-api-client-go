# SecurityMonitoringFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**SecurityMonitoringFilterAction**](SecurityMonitoringFilterAction.md) |  | [optional] 
**Query** | Pointer to **string** | Query for selecting logs to apply the filtering action. | [optional] 

## Methods

### NewSecurityMonitoringFilter

`func NewSecurityMonitoringFilter() *SecurityMonitoringFilter`

NewSecurityMonitoringFilter instantiates a new SecurityMonitoringFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringFilterWithDefaults

`func NewSecurityMonitoringFilterWithDefaults() *SecurityMonitoringFilter`

NewSecurityMonitoringFilterWithDefaults instantiates a new SecurityMonitoringFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SecurityMonitoringFilter) GetAction() SecurityMonitoringFilterAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SecurityMonitoringFilter) GetActionOk() (*SecurityMonitoringFilterAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SecurityMonitoringFilter) SetAction(v SecurityMonitoringFilterAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *SecurityMonitoringFilter) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityMonitoringFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityMonitoringFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityMonitoringFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SecurityMonitoringFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


