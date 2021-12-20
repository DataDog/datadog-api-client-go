# MonitorSearchResult

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Classification** | Pointer to **string** | Classification of the monitor. | [optional] [readonly] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Id** | Pointer to **int64** | ID of the monitor. | [optional] [readonly] 
**LastTriggeredTs** | Pointer to **NullableInt64** | Latest timestamp the monitor triggered. | [optional] [readonly] 
**Metrics** | Pointer to **[]string** | Metrics used by the monitor. | [optional] [readonly] 
**Name** | Pointer to **string** | The monitor name. | [optional] [readonly] 
**Notifications** | Pointer to [**[]MonitorSearchResultNotification**](MonitorSearchResultNotification.md) | The notification triggered by the monitor. | [optional] [readonly] 
**OrgId** | Pointer to **int64** | The ID of the organization. | [optional] [readonly] 
**Query** | Pointer to **string** | The monitor query. | [optional] 
**Scopes** | Pointer to **[]string** | The scope(s) to which the downtime applies, for example &#x60;host:app2&#x60;. Provide multiple scopes as a comma-separated list, for example &#x60;env:dev,env:prod&#x60;. The resulting downtime applies to sources that matches ALL provided scopes (that is &#x60;env:dev AND env:prod&#x60;), NOT any of them. | [optional] 
**Status** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the monitor. | [optional] [readonly] 
**Type** | Pointer to [**MonitorType**](MonitorType.md) |  | [optional] 

## Methods

### NewMonitorSearchResult

`func NewMonitorSearchResult() *MonitorSearchResult`

NewMonitorSearchResult instantiates a new MonitorSearchResult object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorSearchResultWithDefaults

`func NewMonitorSearchResultWithDefaults() *MonitorSearchResult`

NewMonitorSearchResultWithDefaults instantiates a new MonitorSearchResult object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClassification

`func (o *MonitorSearchResult) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MonitorSearchResult) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *MonitorSearchResult) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *MonitorSearchResult) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetCreator

`func (o *MonitorSearchResult) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MonitorSearchResult) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MonitorSearchResult) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MonitorSearchResult) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetId

`func (o *MonitorSearchResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorSearchResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorSearchResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MonitorSearchResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastTriggeredTs

`func (o *MonitorSearchResult) GetLastTriggeredTs() int64`

GetLastTriggeredTs returns the LastTriggeredTs field if non-nil, zero value otherwise.

### GetLastTriggeredTsOk

`func (o *MonitorSearchResult) GetLastTriggeredTsOk() (*int64, bool)`

GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredTs

`func (o *MonitorSearchResult) SetLastTriggeredTs(v int64)`

SetLastTriggeredTs sets LastTriggeredTs field to given value.

### HasLastTriggeredTs

`func (o *MonitorSearchResult) HasLastTriggeredTs() bool`

HasLastTriggeredTs returns a boolean if a field has been set.

### SetLastTriggeredTsNil

`func (o *MonitorSearchResult) SetLastTriggeredTsNil(b bool)`

 SetLastTriggeredTsNil sets the value for LastTriggeredTs to be an explicit nil

### UnsetLastTriggeredTs
`func (o *MonitorSearchResult) UnsetLastTriggeredTs()`

UnsetLastTriggeredTs ensures that no value is present for LastTriggeredTs, not even an explicit nil
### GetMetrics

`func (o *MonitorSearchResult) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MonitorSearchResult) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MonitorSearchResult) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MonitorSearchResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetName

`func (o *MonitorSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorSearchResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorSearchResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *MonitorSearchResult) GetNotifications() []MonitorSearchResultNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MonitorSearchResult) GetNotificationsOk() (*[]MonitorSearchResultNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MonitorSearchResult) SetNotifications(v []MonitorSearchResultNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MonitorSearchResult) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrgId

`func (o *MonitorSearchResult) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MonitorSearchResult) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MonitorSearchResult) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MonitorSearchResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetQuery

`func (o *MonitorSearchResult) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorSearchResult) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorSearchResult) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MonitorSearchResult) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetScopes

`func (o *MonitorSearchResult) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *MonitorSearchResult) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *MonitorSearchResult) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *MonitorSearchResult) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *MonitorSearchResult) GetStatus() MonitorOverallStates`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorSearchResult) GetStatusOk() (*MonitorOverallStates, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorSearchResult) SetStatus(v MonitorOverallStates)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorSearchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MonitorSearchResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitorSearchResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitorSearchResult) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitorSearchResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *MonitorSearchResult) GetType() MonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorSearchResult) GetTypeOk() (*MonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorSearchResult) SetType(v MonitorType)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorSearchResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


