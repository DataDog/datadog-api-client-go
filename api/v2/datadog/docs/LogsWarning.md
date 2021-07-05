# LogsWarning

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Code** | Pointer to **string** | A unique code for this type of warning | [optional] 
**Detail** | Pointer to **string** | A detailed explanation of this specific warning | [optional] 
**Title** | Pointer to **string** | A short human-readable summary of the warning | [optional] 

## Methods

### NewLogsWarning

`func NewLogsWarning() *LogsWarning`

NewLogsWarning instantiates a new LogsWarning object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsWarningWithDefaults

`func NewLogsWarningWithDefaults() *LogsWarning`

NewLogsWarningWithDefaults instantiates a new LogsWarning object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCode

`func (o *LogsWarning) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LogsWarning) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LogsWarning) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LogsWarning) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *LogsWarning) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LogsWarning) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LogsWarning) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LogsWarning) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetTitle

`func (o *LogsWarning) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogsWarning) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogsWarning) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LogsWarning) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


