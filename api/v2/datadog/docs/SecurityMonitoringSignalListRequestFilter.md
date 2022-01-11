# SecurityMonitoringSignalListRequestFilter

## Properties

| Name      | Type                     | Description                                           | Notes      |
| --------- | ------------------------ | ----------------------------------------------------- | ---------- |
| **From**  | Pointer to **time.Time** | The minimum timestamp for requested security signals. | [optional] |
| **Query** | Pointer to **string**    | Search query for listing security signals.            | [optional] |
| **To**    | Pointer to **time.Time** | The maximum timestamp for requested security signals. | [optional] |

## Methods

### NewSecurityMonitoringSignalListRequestFilter

`func NewSecurityMonitoringSignalListRequestFilter() *SecurityMonitoringSignalListRequestFilter`

NewSecurityMonitoringSignalListRequestFilter instantiates a new SecurityMonitoringSignalListRequestFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringSignalListRequestFilterWithDefaults

`func NewSecurityMonitoringSignalListRequestFilterWithDefaults() *SecurityMonitoringSignalListRequestFilter`

NewSecurityMonitoringSignalListRequestFilterWithDefaults instantiates a new SecurityMonitoringSignalListRequestFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFrom

`func (o *SecurityMonitoringSignalListRequestFilter) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SecurityMonitoringSignalListRequestFilter) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SecurityMonitoringSignalListRequestFilter) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SecurityMonitoringSignalListRequestFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQuery

`func (o *SecurityMonitoringSignalListRequestFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityMonitoringSignalListRequestFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityMonitoringSignalListRequestFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SecurityMonitoringSignalListRequestFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTo

`func (o *SecurityMonitoringSignalListRequestFilter) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SecurityMonitoringSignalListRequestFilter) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SecurityMonitoringSignalListRequestFilter) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *SecurityMonitoringSignalListRequestFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
