# SecurityMonitoringSignalListRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Filter** | Pointer to [**SecurityMonitoringSignalListRequestFilter**](SecurityMonitoringSignalListRequestFilter.md) |  | [optional] 
**Page** | Pointer to [**SecurityMonitoringSignalListRequestPage**](SecurityMonitoringSignalListRequestPage.md) |  | [optional] 
**Sort** | Pointer to [**SecurityMonitoringSignalsSort**](SecurityMonitoringSignalsSort.md) |  | [optional] 

## Methods

### NewSecurityMonitoringSignalListRequest

`func NewSecurityMonitoringSignalListRequest() *SecurityMonitoringSignalListRequest`

NewSecurityMonitoringSignalListRequest instantiates a new SecurityMonitoringSignalListRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringSignalListRequestWithDefaults

`func NewSecurityMonitoringSignalListRequestWithDefaults() *SecurityMonitoringSignalListRequest`

NewSecurityMonitoringSignalListRequestWithDefaults instantiates a new SecurityMonitoringSignalListRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilter

`func (o *SecurityMonitoringSignalListRequest) GetFilter() SecurityMonitoringSignalListRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SecurityMonitoringSignalListRequest) GetFilterOk() (*SecurityMonitoringSignalListRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SecurityMonitoringSignalListRequest) SetFilter(v SecurityMonitoringSignalListRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SecurityMonitoringSignalListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPage

`func (o *SecurityMonitoringSignalListRequest) GetPage() SecurityMonitoringSignalListRequestPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SecurityMonitoringSignalListRequest) GetPageOk() (*SecurityMonitoringSignalListRequestPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SecurityMonitoringSignalListRequest) SetPage(v SecurityMonitoringSignalListRequestPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *SecurityMonitoringSignalListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSort

`func (o *SecurityMonitoringSignalListRequest) GetSort() SecurityMonitoringSignalsSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SecurityMonitoringSignalListRequest) GetSortOk() (*SecurityMonitoringSignalsSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SecurityMonitoringSignalListRequest) SetSort(v SecurityMonitoringSignalsSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SecurityMonitoringSignalListRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


