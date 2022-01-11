# DashboardListItemRequest

## Properties

| Name     | Type                                  | Description          | Notes |
| -------- | ------------------------------------- | -------------------- | ----- |
| **Id**   | **string**                            | ID of the dashboard. |
| **Type** | [**DashboardType**](DashboardType.md) |                      |

## Methods

### NewDashboardListItemRequest

`func NewDashboardListItemRequest(id string, type_ DashboardType) *DashboardListItemRequest`

NewDashboardListItemRequest instantiates a new DashboardListItemRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardListItemRequestWithDefaults

`func NewDashboardListItemRequestWithDefaults() *DashboardListItemRequest`

NewDashboardListItemRequestWithDefaults instantiates a new DashboardListItemRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *DashboardListItemRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardListItemRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardListItemRequest) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *DashboardListItemRequest) GetType() DashboardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardListItemRequest) GetTypeOk() (*DashboardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardListItemRequest) SetType(v DashboardType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
