# ProcessSummaryAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Cmdline** | Pointer to **string** | Process command line. | [optional] 
**Host** | Pointer to **string** | Host running the process. | [optional] 
**Pid** | Pointer to **int64** | Process ID. | [optional] 
**Ppid** | Pointer to **int64** | Parent process ID. | [optional] 
**Start** | Pointer to **string** | Time the process was started. | [optional] 
**Tags** | Pointer to **[]string** | List of tags associated with the process. | [optional] 
**Timestamp** | Pointer to **string** | Time the process was seen. | [optional] 
**User** | Pointer to **string** | Process owner. | [optional] 

## Methods

### NewProcessSummaryAttributes

`func NewProcessSummaryAttributes() *ProcessSummaryAttributes`

NewProcessSummaryAttributes instantiates a new ProcessSummaryAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewProcessSummaryAttributesWithDefaults

`func NewProcessSummaryAttributesWithDefaults() *ProcessSummaryAttributes`

NewProcessSummaryAttributesWithDefaults instantiates a new ProcessSummaryAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCmdline

`func (o *ProcessSummaryAttributes) GetCmdline() string`

GetCmdline returns the Cmdline field if non-nil, zero value otherwise.

### GetCmdlineOk

`func (o *ProcessSummaryAttributes) GetCmdlineOk() (*string, bool)`

GetCmdlineOk returns a tuple with the Cmdline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdline

`func (o *ProcessSummaryAttributes) SetCmdline(v string)`

SetCmdline sets Cmdline field to given value.

### HasCmdline

`func (o *ProcessSummaryAttributes) HasCmdline() bool`

HasCmdline returns a boolean if a field has been set.

### GetHost

`func (o *ProcessSummaryAttributes) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProcessSummaryAttributes) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProcessSummaryAttributes) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProcessSummaryAttributes) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPid

`func (o *ProcessSummaryAttributes) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ProcessSummaryAttributes) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ProcessSummaryAttributes) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ProcessSummaryAttributes) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPpid

`func (o *ProcessSummaryAttributes) GetPpid() int64`

GetPpid returns the Ppid field if non-nil, zero value otherwise.

### GetPpidOk

`func (o *ProcessSummaryAttributes) GetPpidOk() (*int64, bool)`

GetPpidOk returns a tuple with the Ppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpid

`func (o *ProcessSummaryAttributes) SetPpid(v int64)`

SetPpid sets Ppid field to given value.

### HasPpid

`func (o *ProcessSummaryAttributes) HasPpid() bool`

HasPpid returns a boolean if a field has been set.

### GetStart

`func (o *ProcessSummaryAttributes) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ProcessSummaryAttributes) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ProcessSummaryAttributes) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *ProcessSummaryAttributes) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *ProcessSummaryAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProcessSummaryAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProcessSummaryAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProcessSummaryAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProcessSummaryAttributes) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProcessSummaryAttributes) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProcessSummaryAttributes) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProcessSummaryAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUser

`func (o *ProcessSummaryAttributes) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProcessSummaryAttributes) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProcessSummaryAttributes) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ProcessSummaryAttributes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


