# NotebookAbsoluteTime

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**End** | **time.Time** | The end time. | 
**Live** | Pointer to **bool** | Indicates whether the timeframe should be shifted to end at the current time. | [optional] 
**Start** | **time.Time** | The start time. | 

## Methods

### NewNotebookAbsoluteTime

`func NewNotebookAbsoluteTime(end time.Time, start time.Time) *NotebookAbsoluteTime`

NewNotebookAbsoluteTime instantiates a new NotebookAbsoluteTime object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookAbsoluteTimeWithDefaults

`func NewNotebookAbsoluteTimeWithDefaults() *NotebookAbsoluteTime`

NewNotebookAbsoluteTimeWithDefaults instantiates a new NotebookAbsoluteTime object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEnd

`func (o *NotebookAbsoluteTime) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *NotebookAbsoluteTime) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *NotebookAbsoluteTime) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetLive

`func (o *NotebookAbsoluteTime) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *NotebookAbsoluteTime) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *NotebookAbsoluteTime) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *NotebookAbsoluteTime) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetStart

`func (o *NotebookAbsoluteTime) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *NotebookAbsoluteTime) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *NotebookAbsoluteTime) SetStart(v time.Time)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


