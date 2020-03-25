# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | TODO. | [optional] [readonly] 
**Creator** | Pointer to [**Creator**](Creator.md) |  | [optional] 
**Deleted** | Pointer to [**NullableTime**](time.Time.md) | TODO. | [optional] [readonly] 
**Id** | Pointer to **int64** | ID of this monitor. | [optional] [readonly] 
**Message** | Pointer to **string** | A message to include with notifications for this monitor. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | TODO. | [optional] [readonly] 
**Multi** | Pointer to **bool** | TODO. | [optional] [readonly] 
**Name** | Pointer to **string** | TODO. | [optional] 
**Options** | Pointer to [**MonitorOptions**](MonitorOptions.md) |  | [optional] 
**OverallState** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**MonitorState**](MonitorState.md) |  | [optional] 
**Tags** | Pointer to **[]string** | TODO. | [optional] 
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

`func (o *Monitor) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Monitor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Monitor) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCreator

`func (o *Monitor) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Monitor) GetCreatorOk() (Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreator

`func (o *Monitor) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreator

`func (o *Monitor) SetCreator(v Creator)`

SetCreator gets a reference to the given Creator and assigns it to the Creator field.

### GetDeleted

`func (o *Monitor) GetDeleted() NullableTime`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Monitor) GetDeletedOk() (NullableTime, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleted

`func (o *Monitor) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeleted

`func (o *Monitor) SetDeleted(v NullableTime)`

SetDeleted gets a reference to the given NullableTime and assigns it to the Deleted field.

### SetDeletedExplicitNull

`func (o *Monitor) SetDeletedExplicitNull(b bool)`

SetDeletedExplicitNull (un)sets Deleted to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Deleted value is set to nil even if false is passed
### GetId

`func (o *Monitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Monitor) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Monitor) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetMessage

`func (o *Monitor) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Monitor) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *Monitor) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *Monitor) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetModified

`func (o *Monitor) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Monitor) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *Monitor) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *Monitor) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetMulti

`func (o *Monitor) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *Monitor) GetMultiOk() (bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMulti

`func (o *Monitor) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### SetMulti

`func (o *Monitor) SetMulti(v bool)`

SetMulti gets a reference to the given bool and assigns it to the Multi field.

### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Monitor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetOptions

`func (o *Monitor) GetOptions() MonitorOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Monitor) GetOptionsOk() (MonitorOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOptions

`func (o *Monitor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptions

`func (o *Monitor) SetOptions(v MonitorOptions)`

SetOptions gets a reference to the given MonitorOptions and assigns it to the Options field.

### GetOverallState

`func (o *Monitor) GetOverallState() MonitorOverallStates`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *Monitor) GetOverallStateOk() (MonitorOverallStates, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverallState

`func (o *Monitor) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### SetOverallState

`func (o *Monitor) SetOverallState(v MonitorOverallStates)`

SetOverallState gets a reference to the given MonitorOverallStates and assigns it to the OverallState field.

### GetQuery

`func (o *Monitor) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Monitor) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *Monitor) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *Monitor) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetState

`func (o *Monitor) GetState() MonitorState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Monitor) GetStateOk() (MonitorState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Monitor) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Monitor) SetState(v MonitorState)`

SetState gets a reference to the given MonitorState and assigns it to the State field.

### GetTags

`func (o *Monitor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Monitor) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Monitor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Monitor) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetType

`func (o *Monitor) GetType() MonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Monitor) GetTypeOk() (MonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Monitor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Monitor) SetType(v MonitorType)`

SetType gets a reference to the given MonitorType and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


