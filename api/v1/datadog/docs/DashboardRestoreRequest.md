# DashboardRestoreRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | [**[]DashboardBulkActionData**](DashboardBulkActionData.md) | List of dashboard bulk action request data objects. | 

## Methods

### NewDashboardRestoreRequest

`func NewDashboardRestoreRequest(data []DashboardBulkActionData) *DashboardRestoreRequest`

NewDashboardRestoreRequest instantiates a new DashboardRestoreRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardRestoreRequestWithDefaults

`func NewDashboardRestoreRequestWithDefaults() *DashboardRestoreRequest`

NewDashboardRestoreRequestWithDefaults instantiates a new DashboardRestoreRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *DashboardRestoreRequest) GetData() []DashboardBulkActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardRestoreRequest) GetDataOk() (*[]DashboardBulkActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardRestoreRequest) SetData(v []DashboardBulkActionData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


