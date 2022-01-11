# HTTPLogError

## Properties

| Name       | Type                  | Description    | Notes      |
| ---------- | --------------------- | -------------- | ---------- |
| **Detail** | Pointer to **string** | Error message. | [optional] |
| **Status** | Pointer to **string** | Error code.    | [optional] |
| **Title**  | Pointer to **string** | Error title.   | [optional] |

## Methods

### NewHTTPLogError

`func NewHTTPLogError() *HTTPLogError`

NewHTTPLogError instantiates a new HTTPLogError object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHTTPLogErrorWithDefaults

`func NewHTTPLogErrorWithDefaults() *HTTPLogError`

NewHTTPLogErrorWithDefaults instantiates a new HTTPLogError object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDetail

`func (o *HTTPLogError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HTTPLogError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HTTPLogError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *HTTPLogError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatus

`func (o *HTTPLogError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HTTPLogError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HTTPLogError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HTTPLogError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *HTTPLogError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HTTPLogError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HTTPLogError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HTTPLogError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
