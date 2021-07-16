# DashboardBulkDeleteRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | [**[]DashboardBulkActionData**](DashboardBulkActionData.md) | List of dashboard bulk action request data objects. | 

## Methods

### NewDashboardBulkDeleteRequest

`func NewDashboardBulkDeleteRequest(data []DashboardBulkActionData) *DashboardBulkDeleteRequest`

NewDashboardBulkDeleteRequest instantiates a new DashboardBulkDeleteRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardBulkDeleteRequestWithDefaults

`func NewDashboardBulkDeleteRequestWithDefaults() *DashboardBulkDeleteRequest`

NewDashboardBulkDeleteRequestWithDefaults instantiates a new DashboardBulkDeleteRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *DashboardBulkDeleteRequest) GetData() []DashboardBulkActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardBulkDeleteRequest) GetDataOk() (*[]DashboardBulkActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardBulkDeleteRequest) SetData(v []DashboardBulkActionData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


