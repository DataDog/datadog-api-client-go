# RUMQueryFilter

## Properties

| Name      | Type                  | Description                                                                                               | Notes                             |
| --------- | --------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------- |
| **From**  | Pointer to **string** | The minimum time for the requested events; supports date, math, and regular timestamps (in milliseconds). | [optional] [default to "now-15m"] |
| **Query** | Pointer to **string** | The search query following the RUM search syntax.                                                         | [optional] [default to "*"]       |
| **To**    | Pointer to **string** | The maximum time for the requested events; supports date, math, and regular timestamps (in milliseconds). | [optional] [default to "now"]     |

## Methods

### NewRUMQueryFilter

`func NewRUMQueryFilter() *RUMQueryFilter`

NewRUMQueryFilter instantiates a new RUMQueryFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMQueryFilterWithDefaults

`func NewRUMQueryFilterWithDefaults() *RUMQueryFilter`

NewRUMQueryFilterWithDefaults instantiates a new RUMQueryFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFrom

`func (o *RUMQueryFilter) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RUMQueryFilter) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RUMQueryFilter) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RUMQueryFilter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQuery

`func (o *RUMQueryFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RUMQueryFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RUMQueryFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RUMQueryFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTo

`func (o *RUMQueryFilter) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RUMQueryFilter) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RUMQueryFilter) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RUMQueryFilter) HasTo() bool`

HasTo returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
