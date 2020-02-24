# DashboardListDeleteItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedDashboardsFromList** | Pointer to [**[]DashboardListItem**](DashboardListItem.md) | List of dashboards deleted from the dashboard list | [optional] 

## Methods

### NewDashboardListDeleteItemsResponse

`func NewDashboardListDeleteItemsResponse() *DashboardListDeleteItemsResponse`

NewDashboardListDeleteItemsResponse instantiates a new DashboardListDeleteItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardListDeleteItemsResponseWithDefaults

`func NewDashboardListDeleteItemsResponseWithDefaults() *DashboardListDeleteItemsResponse`

NewDashboardListDeleteItemsResponseWithDefaults instantiates a new DashboardListDeleteItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedDashboardsFromList

`func (o *DashboardListDeleteItemsResponse) GetDeletedDashboardsFromList() []DashboardListItem`

GetDeletedDashboardsFromList returns the DeletedDashboardsFromList field if non-nil, zero value otherwise.

### GetDeletedDashboardsFromListOk

`func (o *DashboardListDeleteItemsResponse) GetDeletedDashboardsFromListOk() ([]DashboardListItem, bool)`

GetDeletedDashboardsFromListOk returns a tuple with the DeletedDashboardsFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDashboardsFromList

`func (o *DashboardListDeleteItemsResponse) HasDeletedDashboardsFromList() bool`

HasDeletedDashboardsFromList returns a boolean if a field has been set.

### SetDeletedDashboardsFromList

`func (o *DashboardListDeleteItemsResponse) SetDeletedDashboardsFromList(v []DashboardListItem)`

SetDeletedDashboardsFromList gets a reference to the given []DashboardListItem and assigns it to the DeletedDashboardsFromList field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


