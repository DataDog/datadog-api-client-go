# FunnelWidgetRequest

## Properties

| Name            | Type                                          | Description | Notes |
| --------------- | --------------------------------------------- | ----------- | ----- |
| **Query**       | [**FunnelQuery**](FunnelQuery.md)             |             |
| **RequestType** | [**FunnelRequestType**](FunnelRequestType.md) |             |

## Methods

### NewFunnelWidgetRequest

`func NewFunnelWidgetRequest(query FunnelQuery, requestType FunnelRequestType) *FunnelWidgetRequest`

NewFunnelWidgetRequest instantiates a new FunnelWidgetRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFunnelWidgetRequestWithDefaults

`func NewFunnelWidgetRequestWithDefaults() *FunnelWidgetRequest`

NewFunnelWidgetRequestWithDefaults instantiates a new FunnelWidgetRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetQuery

`func (o *FunnelWidgetRequest) GetQuery() FunnelQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FunnelWidgetRequest) GetQueryOk() (*FunnelQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FunnelWidgetRequest) SetQuery(v FunnelQuery)`

SetQuery sets Query field to given value.

### GetRequestType

`func (o *FunnelWidgetRequest) GetRequestType() FunnelRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *FunnelWidgetRequest) GetRequestTypeOk() (*FunnelRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *FunnelWidgetRequest) SetRequestType(v FunnelRequestType)`

SetRequestType sets RequestType field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
