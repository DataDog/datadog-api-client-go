# DashboardListUpdateItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]DashboardListItemResponse**](DashboardListItemResponse.md) | List of dashboards in the dashboard list. | [optional] 

## Methods

### NewDashboardListUpdateItemsResponse

`func NewDashboardListUpdateItemsResponse() *DashboardListUpdateItemsResponse`

NewDashboardListUpdateItemsResponse instantiates a new DashboardListUpdateItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListUpdateItemsResponseWithDefaults

`func NewDashboardListUpdateItemsResponseWithDefaults() *DashboardListUpdateItemsResponse`

NewDashboardListUpdateItemsResponseWithDefaults instantiates a new DashboardListUpdateItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *DashboardListUpdateItemsResponse) GetDashboards() []DashboardListItemResponse`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardListUpdateItemsResponse) GetDashboardsOk() (*[]DashboardListItemResponse, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardListUpdateItemsResponse) SetDashboards(v []DashboardListItemResponse)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *DashboardListUpdateItemsResponse) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


