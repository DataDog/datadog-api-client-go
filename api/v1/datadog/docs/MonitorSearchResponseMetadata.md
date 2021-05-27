# MonitorSearchResponseMetadata

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Page** | Pointer to **int64** | The page to start paginating from. | [optional] [readonly] 
**PageCount** | Pointer to **int64** | The number of pages. | [optional] [readonly] 
**PerPage** | Pointer to **int64** | The number of monitors to return per page. | [optional] [readonly] 
**TotalCount** | Pointer to **int64** | The total number of monitors. | [optional] [readonly] 

## Methods

### NewMonitorSearchResponseMetadata

`func NewMonitorSearchResponseMetadata() *MonitorSearchResponseMetadata`

NewMonitorSearchResponseMetadata instantiates a new MonitorSearchResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSearchResponseMetadataWithDefaults

`func NewMonitorSearchResponseMetadataWithDefaults() *MonitorSearchResponseMetadata`

NewMonitorSearchResponseMetadataWithDefaults instantiates a new MonitorSearchResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *MonitorSearchResponseMetadata) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MonitorSearchResponseMetadata) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MonitorSearchResponseMetadata) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *MonitorSearchResponseMetadata) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageCount

`func (o *MonitorSearchResponseMetadata) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *MonitorSearchResponseMetadata) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *MonitorSearchResponseMetadata) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *MonitorSearchResponseMetadata) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPerPage

`func (o *MonitorSearchResponseMetadata) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *MonitorSearchResponseMetadata) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *MonitorSearchResponseMetadata) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *MonitorSearchResponseMetadata) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotalCount

`func (o *MonitorSearchResponseMetadata) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MonitorSearchResponseMetadata) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MonitorSearchResponseMetadata) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *MonitorSearchResponseMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


