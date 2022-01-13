# IncidentNotificationHandle

## Properties

| Name            | Type                  | Description                                  | Notes      |
| --------------- | --------------------- | -------------------------------------------- | ---------- |
| **DisplayName** | Pointer to **string** | The name of the notified handle.             | [optional] |
| **Handle**      | Pointer to **string** | The email address used for the notification. | [optional] |

## Methods

### NewIncidentNotificationHandle

`func NewIncidentNotificationHandle() *IncidentNotificationHandle`

NewIncidentNotificationHandle instantiates a new IncidentNotificationHandle object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentNotificationHandleWithDefaults

`func NewIncidentNotificationHandleWithDefaults() *IncidentNotificationHandle`

NewIncidentNotificationHandleWithDefaults instantiates a new IncidentNotificationHandle object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplayName

`func (o *IncidentNotificationHandle) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IncidentNotificationHandle) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IncidentNotificationHandle) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IncidentNotificationHandle) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *IncidentNotificationHandle) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *IncidentNotificationHandle) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *IncidentNotificationHandle) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *IncidentNotificationHandle) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
