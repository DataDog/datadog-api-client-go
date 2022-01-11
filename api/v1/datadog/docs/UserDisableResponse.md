# UserDisableResponse

## Properties

| Name        | Type                  | Description                                                         | Notes      |
| ----------- | --------------------- | ------------------------------------------------------------------- | ---------- |
| **Message** | Pointer to **string** | Information pertaining to a user disabled for a given organization. | [optional] |

## Methods

### NewUserDisableResponse

`func NewUserDisableResponse() *UserDisableResponse`

NewUserDisableResponse instantiates a new UserDisableResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserDisableResponseWithDefaults

`func NewUserDisableResponseWithDefaults() *UserDisableResponse`

NewUserDisableResponseWithDefaults instantiates a new UserDisableResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMessage

`func (o *UserDisableResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserDisableResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserDisableResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserDisableResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
