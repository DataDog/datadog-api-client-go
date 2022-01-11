# SlackIntegrationChannel

## Properties

| Name        | Type                                                                               | Description        | Notes      |
| ----------- | ---------------------------------------------------------------------------------- | ------------------ | ---------- |
| **Display** | Pointer to [**SlackIntegrationChannelDisplay**](SlackIntegrationChannelDisplay.md) |                    | [optional] |
| **Name**    | Pointer to **string**                                                              | Your channel name. | [optional] |

## Methods

### NewSlackIntegrationChannel

`func NewSlackIntegrationChannel() *SlackIntegrationChannel`

NewSlackIntegrationChannel instantiates a new SlackIntegrationChannel object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSlackIntegrationChannelWithDefaults

`func NewSlackIntegrationChannelWithDefaults() *SlackIntegrationChannel`

NewSlackIntegrationChannelWithDefaults instantiates a new SlackIntegrationChannel object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplay

`func (o *SlackIntegrationChannel) GetDisplay() SlackIntegrationChannelDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SlackIntegrationChannel) GetDisplayOk() (*SlackIntegrationChannelDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SlackIntegrationChannel) SetDisplay(v SlackIntegrationChannelDisplay)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *SlackIntegrationChannel) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *SlackIntegrationChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackIntegrationChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackIntegrationChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlackIntegrationChannel) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
