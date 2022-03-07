# RUMResponseMetadata

## Properties

| Name          | Type                                                     | Description                                                                                                            | Notes      |
| ------------- | -------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Elapsed**   | Pointer to **int64**                                     | The time elapsed in milliseconds.                                                                                      | [optional] |
| **Page**      | Pointer to [**RUMResponsePage**](RUMResponsePage.md)     |                                                                                                                        | [optional] |
| **RequestId** | Pointer to **string**                                    | The identifier of the request.                                                                                         | [optional] |
| **Status**    | Pointer to [**RUMResponseStatus**](RUMResponseStatus.md) |                                                                                                                        | [optional] |
| **Warnings**  | Pointer to [**[]RUMWarning**](RUMWarning.md)             | A list of warnings (non-fatal errors) encountered. Partial results may return if warnings are present in the response. | [optional] |

## Methods

### NewRUMResponseMetadata

`func NewRUMResponseMetadata() *RUMResponseMetadata`

NewRUMResponseMetadata instantiates a new RUMResponseMetadata object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMResponseMetadataWithDefaults

`func NewRUMResponseMetadataWithDefaults() *RUMResponseMetadata`

NewRUMResponseMetadataWithDefaults instantiates a new RUMResponseMetadata object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetElapsed

`func (o *RUMResponseMetadata) GetElapsed() int64`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *RUMResponseMetadata) GetElapsedOk() (*int64, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *RUMResponseMetadata) SetElapsed(v int64)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *RUMResponseMetadata) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetPage

`func (o *RUMResponseMetadata) GetPage() RUMResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RUMResponseMetadata) GetPageOk() (*RUMResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RUMResponseMetadata) SetPage(v RUMResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *RUMResponseMetadata) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRequestId

`func (o *RUMResponseMetadata) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RUMResponseMetadata) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RUMResponseMetadata) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RUMResponseMetadata) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *RUMResponseMetadata) GetStatus() RUMResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RUMResponseMetadata) GetStatusOk() (*RUMResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RUMResponseMetadata) SetStatus(v RUMResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RUMResponseMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarnings

`func (o *RUMResponseMetadata) GetWarnings() []RUMWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RUMResponseMetadata) GetWarningsOk() (*[]RUMWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RUMResponseMetadata) SetWarnings(v []RUMWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RUMResponseMetadata) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
