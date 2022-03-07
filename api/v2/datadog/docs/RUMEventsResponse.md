# RUMEventsResponse

## Properties

| Name      | Type                                                         | Description                           | Notes      |
| --------- | ------------------------------------------------------------ | ------------------------------------- | ---------- |
| **Data**  | Pointer to [**[]RUMEvent**](RUMEvent.md)                     | Array of events matching the request. | [optional] |
| **Links** | Pointer to [**RUMResponseLinks**](RUMResponseLinks.md)       |                                       | [optional] |
| **Meta**  | Pointer to [**RUMResponseMetadata**](RUMResponseMetadata.md) |                                       | [optional] |

## Methods

### NewRUMEventsResponse

`func NewRUMEventsResponse() *RUMEventsResponse`

NewRUMEventsResponse instantiates a new RUMEventsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMEventsResponseWithDefaults

`func NewRUMEventsResponseWithDefaults() *RUMEventsResponse`

NewRUMEventsResponseWithDefaults instantiates a new RUMEventsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *RUMEventsResponse) GetData() []RUMEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RUMEventsResponse) GetDataOk() (*[]RUMEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RUMEventsResponse) SetData(v []RUMEvent)`

SetData sets Data field to given value.

### HasData

`func (o *RUMEventsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RUMEventsResponse) GetLinks() RUMResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RUMEventsResponse) GetLinksOk() (*RUMResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RUMEventsResponse) SetLinks(v RUMResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RUMEventsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *RUMEventsResponse) GetMeta() RUMResponseMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RUMEventsResponse) GetMetaOk() (*RUMResponseMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RUMEventsResponse) SetMeta(v RUMResponseMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RUMEventsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
