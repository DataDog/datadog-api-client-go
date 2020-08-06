# SecurityMonitoringSignalListRequestPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | A list of results using the cursor provided in the previous query. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of security signals in the response. | [optional] [default to 10]

## Methods

### NewSecurityMonitoringSignalListRequestPage

`func NewSecurityMonitoringSignalListRequestPage() *SecurityMonitoringSignalListRequestPage`

NewSecurityMonitoringSignalListRequestPage instantiates a new SecurityMonitoringSignalListRequestPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringSignalListRequestPageWithDefaults

`func NewSecurityMonitoringSignalListRequestPageWithDefaults() *SecurityMonitoringSignalListRequestPage`

NewSecurityMonitoringSignalListRequestPageWithDefaults instantiates a new SecurityMonitoringSignalListRequestPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *SecurityMonitoringSignalListRequestPage) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SecurityMonitoringSignalListRequestPage) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SecurityMonitoringSignalListRequestPage) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SecurityMonitoringSignalListRequestPage) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetLimit

`func (o *SecurityMonitoringSignalListRequestPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SecurityMonitoringSignalListRequestPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SecurityMonitoringSignalListRequestPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SecurityMonitoringSignalListRequestPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


