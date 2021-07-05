# DashboardListItemResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Id** | **string** | ID of the dashboard. | [readonly] 
**Type** | [**DashboardType**](DashboardType.md) |  | 

## Methods

### NewDashboardListItemResponse

`func NewDashboardListItemResponse(id string, type_ DashboardType) *DashboardListItemResponse`

NewDashboardListItemResponse instantiates a new DashboardListItemResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardListItemResponseWithDefaults

`func NewDashboardListItemResponseWithDefaults() *DashboardListItemResponse`

NewDashboardListItemResponseWithDefaults instantiates a new DashboardListItemResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *DashboardListItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardListItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardListItemResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DashboardListItemResponse) GetType() DashboardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardListItemResponse) GetTypeOk() (*DashboardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardListItemResponse) SetType(v DashboardType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


