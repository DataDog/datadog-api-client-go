# SecurityMonitoringSignalsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityMonitoringSignal**](SecurityMonitoringSignal.md) | Array of security signals matching the request. | [optional] 
**Links** | Pointer to [**SecurityMonitoringSignalsListResponseLinks**](SecurityMonitoringSignalsListResponse_links.md) |  | [optional] 
**Meta** | Pointer to [**LogsListResponseMeta**](LogsListResponse_meta.md) |  | [optional] 

## Methods

### NewSecurityMonitoringSignalsListResponse

`func NewSecurityMonitoringSignalsListResponse() *SecurityMonitoringSignalsListResponse`

NewSecurityMonitoringSignalsListResponse instantiates a new SecurityMonitoringSignalsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringSignalsListResponseWithDefaults

`func NewSecurityMonitoringSignalsListResponseWithDefaults() *SecurityMonitoringSignalsListResponse`

NewSecurityMonitoringSignalsListResponseWithDefaults instantiates a new SecurityMonitoringSignalsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityMonitoringSignalsListResponse) GetData() []SecurityMonitoringSignal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityMonitoringSignalsListResponse) GetDataOk() (*[]SecurityMonitoringSignal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityMonitoringSignalsListResponse) SetData(v []SecurityMonitoringSignal)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityMonitoringSignalsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityMonitoringSignalsListResponse) GetLinks() SecurityMonitoringSignalsListResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityMonitoringSignalsListResponse) GetLinksOk() (*SecurityMonitoringSignalsListResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityMonitoringSignalsListResponse) SetLinks(v SecurityMonitoringSignalsListResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityMonitoringSignalsListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityMonitoringSignalsListResponse) GetMeta() LogsListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityMonitoringSignalsListResponse) GetMetaOk() (*LogsListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityMonitoringSignalsListResponse) SetMeta(v LogsListResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityMonitoringSignalsListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


