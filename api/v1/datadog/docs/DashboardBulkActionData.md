# DashboardBulkActionData

## Properties

| Name     | Type                                                  | Description            | Notes                                        |
| -------- | ----------------------------------------------------- | ---------------------- | -------------------------------------------- |
| **Id**   | **string**                                            | Dashboard resource ID. |
| **Type** | [**DashboardResourceType**](DashboardResourceType.md) |                        | [default to DASHBOARDRESOURCETYPE_DASHBOARD] |

## Methods

### NewDashboardBulkActionData

`func NewDashboardBulkActionData(id string, type_ DashboardResourceType) *DashboardBulkActionData`

NewDashboardBulkActionData instantiates a new DashboardBulkActionData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardBulkActionDataWithDefaults

`func NewDashboardBulkActionDataWithDefaults() *DashboardBulkActionData`

NewDashboardBulkActionDataWithDefaults instantiates a new DashboardBulkActionData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *DashboardBulkActionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardBulkActionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardBulkActionData) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *DashboardBulkActionData) GetType() DashboardResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardBulkActionData) GetTypeOk() (*DashboardResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardBulkActionData) SetType(v DashboardResourceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
