# LogsResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | The time elapsed in milliseconds | [optional] 
**Page** | Pointer to [**LogsResponseMetadataPage**](LogsResponseMetadata_page.md) |  | [optional] 
**RequestId** | Pointer to **string** | The identifier of the request | [optional] 
**Status** | Pointer to [**LogsAggregateResponseStatus**](LogsAggregateResponseStatus.md) |  | [optional] 
**Warnings** | Pointer to [**[]LogsWarning**](LogsWarning.md) | A list of warnings (non fatal errors) encountered, partial results might be returned if warnings are present in the response. | [optional] 

## Methods

### NewLogsResponseMetadata

`func NewLogsResponseMetadata() *LogsResponseMetadata`

NewLogsResponseMetadata instantiates a new LogsResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseMetadataWithDefaults

`func NewLogsResponseMetadataWithDefaults() *LogsResponseMetadata`

NewLogsResponseMetadataWithDefaults instantiates a new LogsResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *LogsResponseMetadata) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *LogsResponseMetadata) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *LogsResponseMetadata) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *LogsResponseMetadata) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetPage

`func (o *LogsResponseMetadata) GetPage() LogsResponseMetadataPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LogsResponseMetadata) GetPageOk() (*LogsResponseMetadataPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LogsResponseMetadata) SetPage(v LogsResponseMetadataPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *LogsResponseMetadata) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRequestId

`func (o *LogsResponseMetadata) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LogsResponseMetadata) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LogsResponseMetadata) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *LogsResponseMetadata) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *LogsResponseMetadata) GetStatus() LogsAggregateResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogsResponseMetadata) GetStatusOk() (*LogsAggregateResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogsResponseMetadata) SetStatus(v LogsAggregateResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogsResponseMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarnings

`func (o *LogsResponseMetadata) GetWarnings() []LogsWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *LogsResponseMetadata) GetWarningsOk() (*[]LogsWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *LogsResponseMetadata) SetWarnings(v []LogsWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *LogsResponseMetadata) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


