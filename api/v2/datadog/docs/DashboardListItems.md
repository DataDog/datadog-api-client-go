# DashboardListItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]DashboardListItem**](DashboardListItem.md) | List of dashboards in the dashboard list. | 
**Total** | Pointer to **int64** | Number of dashboards in the dashboard list. | [optional] [readonly] 

## Methods

### NewDashboardListItems

`func NewDashboardListItems(dashboards []DashboardListItem, ) *DashboardListItems`

NewDashboardListItems instantiates a new DashboardListItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListItemsWithDefaults

`func NewDashboardListItemsWithDefaults() *DashboardListItems`

NewDashboardListItemsWithDefaults instantiates a new DashboardListItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *DashboardListItems) GetDashboards() []DashboardListItem`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardListItems) GetDashboardsOk() ([]DashboardListItem, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDashboards

`func (o *DashboardListItems) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### SetDashboards

`func (o *DashboardListItems) SetDashboards(v []DashboardListItem)`

SetDashboards gets a reference to the given []DashboardListItem and assigns it to the Dashboards field.

### GetTotal

`func (o *DashboardListItems) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DashboardListItems) GetTotalOk() (int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *DashboardListItems) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *DashboardListItems) SetTotal(v int64)`

SetTotal gets a reference to the given int64 and assigns it to the Total field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


