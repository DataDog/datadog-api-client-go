# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the monitor creation. | [optional] [readonly] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Deleted** | Pointer to [**NullableTime**](time.Time.md) | Whether or not the monitor is deleted. (Always &#x60;null&#x60;) | [optional] [readonly] 
**Id** | Pointer to **int64** | ID of this monitor. | [optional] [readonly] 
**Message** | Pointer to **string** | A message to include with notifications for this monitor. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Last timestamp when the monitor was edited. | [optional] [readonly] 
**Multi** | Pointer to **bool** | Whether or not the monitor is broken down on different groups. | [optional] [readonly] 
**Name** | Pointer to **string** | The monitor name. | [optional] 
**Options** | Pointer to [**MonitorOptions**](MonitorOptions.md) |  | [optional] 
**OverallState** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**Query** | Pointer to **string** | The monitor query. | [optional] 
**State** | Pointer to [**MonitorState**](MonitorState.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated to your monitor. | [optional] 
**Type** | Pointer to [**MonitorType**](MonitorType.md) |  | [optional] 

## Methods

### NewMonitor

`func NewMonitor() *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Monitor) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Monitor) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Monitor) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Monitor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *Monitor) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Monitor) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Monitor) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Monitor) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDeleted

`func (o *Monitor) GetDeleted() time.Time`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Monitor) GetDeletedOk() (*time.Time, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Monitor) SetDeleted(v time.Time)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Monitor) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *Monitor) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *Monitor) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetId

`func (o *Monitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Monitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *Monitor) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Monitor) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Monitor) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Monitor) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModified

`func (o *Monitor) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Monitor) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Monitor) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Monitor) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetMulti

`func (o *Monitor) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *Monitor) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *Monitor) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *Monitor) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Monitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *Monitor) GetOptions() MonitorOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Monitor) GetOptionsOk() (*MonitorOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Monitor) SetOptions(v MonitorOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Monitor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOverallState

`func (o *Monitor) GetOverallState() MonitorOverallStates`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *Monitor) GetOverallStateOk() (*MonitorOverallStates, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *Monitor) SetOverallState(v MonitorOverallStates)`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *Monitor) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### GetQuery

`func (o *Monitor) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Monitor) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Monitor) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Monitor) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetState

`func (o *Monitor) GetState() MonitorState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Monitor) GetStateOk() (*MonitorState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Monitor) SetState(v MonitorState)`

SetState sets State field to given value.

### HasState

`func (o *Monitor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *Monitor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Monitor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Monitor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Monitor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Monitor) GetType() MonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Monitor) GetTypeOk() (*MonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Monitor) SetType(v MonitorType)`

SetType sets Type field to given value.

### HasType

`func (o *Monitor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


