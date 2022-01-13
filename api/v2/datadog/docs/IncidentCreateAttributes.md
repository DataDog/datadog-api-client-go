# IncidentCreateAttributes

## Properties

| Name                    | Type                                                                                             | Description                                                                                | Notes      |
| ----------------------- | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ---------- |
| **CustomerImpacted**    | **bool**                                                                                         | A flag indicating whether the incident caused customer impact.                             |
| **Fields**              | Pointer to [**map[string]IncidentFieldAttributes**](IncidentFieldAttributes.md)                  | A condensed view of the user-defined fields for which to create initial selections.        | [optional] |
| **InitialCells**        | Pointer to [**[]IncidentTimelineCellCreateAttributes**](IncidentTimelineCellCreateAttributes.md) | An array of initial timeline cells to be placed at the beginning of the incident timeline. | [optional] |
| **NotificationHandles** | Pointer to [**[]IncidentNotificationHandle**](IncidentNotificationHandle.md)                     | Notification handles that will be notified of the incident at creation.                    | [optional] |
| **Title**               | **string**                                                                                       | The title of the incident, which summarizes what happened.                                 |

## Methods

### NewIncidentCreateAttributes

`func NewIncidentCreateAttributes(customerImpacted bool, title string) *IncidentCreateAttributes`

NewIncidentCreateAttributes instantiates a new IncidentCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentCreateAttributesWithDefaults

`func NewIncidentCreateAttributesWithDefaults() *IncidentCreateAttributes`

NewIncidentCreateAttributesWithDefaults instantiates a new IncidentCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCustomerImpacted

`func (o *IncidentCreateAttributes) GetCustomerImpacted() bool`

GetCustomerImpacted returns the CustomerImpacted field if non-nil, zero value otherwise.

### GetCustomerImpactedOk

`func (o *IncidentCreateAttributes) GetCustomerImpactedOk() (*bool, bool)`

GetCustomerImpactedOk returns a tuple with the CustomerImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpacted

`func (o *IncidentCreateAttributes) SetCustomerImpacted(v bool)`

SetCustomerImpacted sets CustomerImpacted field to given value.

### GetFields

`func (o *IncidentCreateAttributes) GetFields() map[string]IncidentFieldAttributes`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IncidentCreateAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IncidentCreateAttributes) SetFields(v map[string]IncidentFieldAttributes)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IncidentCreateAttributes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetInitialCells

`func (o *IncidentCreateAttributes) GetInitialCells() []IncidentTimelineCellCreateAttributes`

GetInitialCells returns the InitialCells field if non-nil, zero value otherwise.

### GetInitialCellsOk

`func (o *IncidentCreateAttributes) GetInitialCellsOk() (*[]IncidentTimelineCellCreateAttributes, bool)`

GetInitialCellsOk returns a tuple with the InitialCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCells

`func (o *IncidentCreateAttributes) SetInitialCells(v []IncidentTimelineCellCreateAttributes)`

SetInitialCells sets InitialCells field to given value.

### HasInitialCells

`func (o *IncidentCreateAttributes) HasInitialCells() bool`

HasInitialCells returns a boolean if a field has been set.

### GetNotificationHandles

`func (o *IncidentCreateAttributes) GetNotificationHandles() []IncidentNotificationHandle`

GetNotificationHandles returns the NotificationHandles field if non-nil, zero value otherwise.

### GetNotificationHandlesOk

`func (o *IncidentCreateAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool)`

GetNotificationHandlesOk returns a tuple with the NotificationHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationHandles

`func (o *IncidentCreateAttributes) SetNotificationHandles(v []IncidentNotificationHandle)`

SetNotificationHandles sets NotificationHandles field to given value.

### HasNotificationHandles

`func (o *IncidentCreateAttributes) HasNotificationHandles() bool`

HasNotificationHandles returns a boolean if a field has been set.

### GetTitle

`func (o *IncidentCreateAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentCreateAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentCreateAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
