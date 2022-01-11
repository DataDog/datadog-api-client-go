# DashboardListAddItemsRequest

## Properties

| Name           | Type                                                                     | Description                                   | Notes      |
| -------------- | ------------------------------------------------------------------------ | --------------------------------------------- | ---------- |
| **Dashboards** | Pointer to [**[]DashboardListItemRequest**](DashboardListItemRequest.md) | List of dashboards to add the dashboard list. | [optional] |

## Methods

### NewDashboardListAddItemsRequest

`func NewDashboardListAddItemsRequest() *DashboardListAddItemsRequest`

NewDashboardListAddItemsRequest instantiates a new DashboardListAddItemsRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardListAddItemsRequestWithDefaults

`func NewDashboardListAddItemsRequestWithDefaults() *DashboardListAddItemsRequest`

NewDashboardListAddItemsRequestWithDefaults instantiates a new DashboardListAddItemsRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDashboards

`func (o *DashboardListAddItemsRequest) GetDashboards() []DashboardListItemRequest`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardListAddItemsRequest) GetDashboardsOk() (*[]DashboardListItemRequest, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardListAddItemsRequest) SetDashboards(v []DashboardListItemRequest)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *DashboardListAddItemsRequest) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
