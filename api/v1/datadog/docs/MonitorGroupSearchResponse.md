# MonitorGroupSearchResponse

## Properties

| Name         | Type                                                                                   | Description                       | Notes                 |
| ------------ | -------------------------------------------------------------------------------------- | --------------------------------- | --------------------- |
| **Counts**   | Pointer to [**MonitorGroupSearchResponseCounts**](MonitorGroupSearchResponseCounts.md) |                                   | [optional]            |
| **Groups**   | Pointer to [**[]MonitorGroupSearchResult**](MonitorGroupSearchResult.md)               | The list of found monitor groups. | [optional] [readonly] |
| **Metadata** | Pointer to [**MonitorSearchResponseMetadata**](MonitorSearchResponseMetadata.md)       |                                   | [optional]            |

## Methods

### NewMonitorGroupSearchResponse

`func NewMonitorGroupSearchResponse() *MonitorGroupSearchResponse`

NewMonitorGroupSearchResponse instantiates a new MonitorGroupSearchResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorGroupSearchResponseWithDefaults

`func NewMonitorGroupSearchResponseWithDefaults() *MonitorGroupSearchResponse`

NewMonitorGroupSearchResponseWithDefaults instantiates a new MonitorGroupSearchResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCounts

`func (o *MonitorGroupSearchResponse) GetCounts() MonitorGroupSearchResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MonitorGroupSearchResponse) GetCountsOk() (*MonitorGroupSearchResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MonitorGroupSearchResponse) SetCounts(v MonitorGroupSearchResponseCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MonitorGroupSearchResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetGroups

`func (o *MonitorGroupSearchResponse) GetGroups() []MonitorGroupSearchResult`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MonitorGroupSearchResponse) GetGroupsOk() (*[]MonitorGroupSearchResult, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MonitorGroupSearchResponse) SetGroups(v []MonitorGroupSearchResult)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MonitorGroupSearchResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMetadata

`func (o *MonitorGroupSearchResponse) GetMetadata() MonitorSearchResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MonitorGroupSearchResponse) GetMetadataOk() (*MonitorSearchResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MonitorGroupSearchResponse) SetMetadata(v MonitorSearchResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MonitorGroupSearchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
