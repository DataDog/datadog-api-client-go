# LogsListResponse

## Properties

| Name      | Type                                                             | Description                         | Notes      |
| --------- | ---------------------------------------------------------------- | ----------------------------------- | ---------- |
| **Data**  | Pointer to [**[]Log**](Log.md)                                   | Array of logs matching the request. | [optional] |
| **Links** | Pointer to [**LogsListResponseLinks**](LogsListResponseLinks.md) |                                     | [optional] |
| **Meta**  | Pointer to [**LogsResponseMetadata**](LogsResponseMetadata.md)   |                                     | [optional] |

## Methods

### NewLogsListResponse

`func NewLogsListResponse() *LogsListResponse`

NewLogsListResponse instantiates a new LogsListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsListResponseWithDefaults

`func NewLogsListResponseWithDefaults() *LogsListResponse`

NewLogsListResponseWithDefaults instantiates a new LogsListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *LogsListResponse) GetData() []Log`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogsListResponse) GetDataOk() (*[]Log, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogsListResponse) SetData(v []Log)`

SetData sets Data field to given value.

### HasData

`func (o *LogsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *LogsListResponse) GetLinks() LogsListResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogsListResponse) GetLinksOk() (*LogsListResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogsListResponse) SetLinks(v LogsListResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogsListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *LogsListResponse) GetMeta() LogsResponseMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LogsListResponse) GetMetaOk() (*LogsResponseMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LogsListResponse) SetMeta(v LogsResponseMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LogsListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
