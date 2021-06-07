# MonitorSearchResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Counts** | Pointer to [**MonitorSearchResponseCounts**](MonitorSearchResponseCounts.md) |  | [optional] 
**Metadata** | Pointer to [**MonitorSearchResponseMetadata**](MonitorSearchResponseMetadata.md) |  | [optional] 
**Monitors** | Pointer to [**[]MonitorSearchResult**](MonitorSearchResult.md) | The list of found monitors. | [optional] [readonly] 

## Methods

### NewMonitorSearchResponse

`func NewMonitorSearchResponse() *MonitorSearchResponse`

NewMonitorSearchResponse instantiates a new MonitorSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSearchResponseWithDefaults

`func NewMonitorSearchResponseWithDefaults() *MonitorSearchResponse`

NewMonitorSearchResponseWithDefaults instantiates a new MonitorSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *MonitorSearchResponse) GetCounts() MonitorSearchResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MonitorSearchResponse) GetCountsOk() (*MonitorSearchResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MonitorSearchResponse) SetCounts(v MonitorSearchResponseCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MonitorSearchResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetMetadata

`func (o *MonitorSearchResponse) GetMetadata() MonitorSearchResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MonitorSearchResponse) GetMetadataOk() (*MonitorSearchResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MonitorSearchResponse) SetMetadata(v MonitorSearchResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MonitorSearchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitors

`func (o *MonitorSearchResponse) GetMonitors() []MonitorSearchResult`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *MonitorSearchResponse) GetMonitorsOk() (*[]MonitorSearchResult, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *MonitorSearchResponse) SetMonitors(v []MonitorSearchResult)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *MonitorSearchResponse) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


