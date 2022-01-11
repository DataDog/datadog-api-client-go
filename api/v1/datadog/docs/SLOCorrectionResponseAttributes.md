# SLOCorrectionResponseAttributes

## Properties

| Name            | Type                                                                                                         | Description                                                                                                                                                                             | Notes      |
| --------------- | ------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Category**    | Pointer to [**SLOCorrectionCategory**](SLOCorrectionCategory.md)                                             |                                                                                                                                                                                         | [optional] |
| **CreatedAt**   | Pointer to **int64**                                                                                         | The epoch timestamp of when the correction was created at                                                                                                                               | [optional] |
| **Creator**     | Pointer to [**Creator**](Creator.md)                                                                         |                                                                                                                                                                                         | [optional] |
| **Description** | Pointer to **string**                                                                                        | Description of the correction being made.                                                                                                                                               | [optional] |
| **Duration**    | Pointer to **NullableInt64**                                                                                 | Length of time (in seconds) for a specified &#x60;rrule&#x60; recurring SLO correction.                                                                                                 | [optional] |
| **End**         | Pointer to **int64**                                                                                         | Ending time of the correction in epoch seconds.                                                                                                                                         | [optional] |
| **ModifiedAt**  | Pointer to **int64**                                                                                         | The epoch timestamp of when the correction was modified at                                                                                                                              | [optional] |
| **Modifier**    | Pointer to [**NullableSLOCorrectionResponseAttributesModifier**](SLOCorrectionResponseAttributesModifier.md) |                                                                                                                                                                                         | [optional] |
| **Rrule**       | Pointer to **NullableString**                                                                                | The recurrence rules as defined in the iCalendar RFC 5545. The supported rules for SLO corrections are &#x60;FREQ&#x60;, &#x60;INTERVAL&#x60;, &#x60;COUNT&#x60; and &#x60;UNTIL&#x60;. | [optional] |
| **SloId**       | Pointer to **string**                                                                                        | ID of the SLO that this correction will be applied to.                                                                                                                                  | [optional] |
| **Start**       | Pointer to **int64**                                                                                         | Starting time of the correction in epoch seconds.                                                                                                                                       | [optional] |
| **Timezone**    | Pointer to **string**                                                                                        | The timezone to display in the UI for the correction times (defaults to \&quot;UTC\&quot;).                                                                                             | [optional] |

## Methods

### NewSLOCorrectionResponseAttributes

`func NewSLOCorrectionResponseAttributes() *SLOCorrectionResponseAttributes`

NewSLOCorrectionResponseAttributes instantiates a new SLOCorrectionResponseAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOCorrectionResponseAttributesWithDefaults

`func NewSLOCorrectionResponseAttributesWithDefaults() *SLOCorrectionResponseAttributes`

NewSLOCorrectionResponseAttributesWithDefaults instantiates a new SLOCorrectionResponseAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCategory

`func (o *SLOCorrectionResponseAttributes) GetCategory() SLOCorrectionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SLOCorrectionResponseAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SLOCorrectionResponseAttributes) SetCategory(v SLOCorrectionCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SLOCorrectionResponseAttributes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SLOCorrectionResponseAttributes) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SLOCorrectionResponseAttributes) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SLOCorrectionResponseAttributes) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SLOCorrectionResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *SLOCorrectionResponseAttributes) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SLOCorrectionResponseAttributes) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SLOCorrectionResponseAttributes) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SLOCorrectionResponseAttributes) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *SLOCorrectionResponseAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLOCorrectionResponseAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLOCorrectionResponseAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLOCorrectionResponseAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *SLOCorrectionResponseAttributes) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SLOCorrectionResponseAttributes) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SLOCorrectionResponseAttributes) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SLOCorrectionResponseAttributes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *SLOCorrectionResponseAttributes) SetDurationNil(b bool)`

SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration

`func (o *SLOCorrectionResponseAttributes) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

### GetEnd

`func (o *SLOCorrectionResponseAttributes) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SLOCorrectionResponseAttributes) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SLOCorrectionResponseAttributes) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SLOCorrectionResponseAttributes) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SLOCorrectionResponseAttributes) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SLOCorrectionResponseAttributes) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SLOCorrectionResponseAttributes) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SLOCorrectionResponseAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifier

`func (o *SLOCorrectionResponseAttributes) GetModifier() SLOCorrectionResponseAttributesModifier`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *SLOCorrectionResponseAttributes) GetModifierOk() (*SLOCorrectionResponseAttributesModifier, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *SLOCorrectionResponseAttributes) SetModifier(v SLOCorrectionResponseAttributesModifier)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *SLOCorrectionResponseAttributes) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### SetModifierNil

`func (o *SLOCorrectionResponseAttributes) SetModifierNil(b bool)`

SetModifierNil sets the value for Modifier to be an explicit nil

### UnsetModifier

`func (o *SLOCorrectionResponseAttributes) UnsetModifier()`

UnsetModifier ensures that no value is present for Modifier, not even an explicit nil

### GetRrule

`func (o *SLOCorrectionResponseAttributes) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *SLOCorrectionResponseAttributes) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *SLOCorrectionResponseAttributes) SetRrule(v string)`

SetRrule sets Rrule field to given value.

### HasRrule

`func (o *SLOCorrectionResponseAttributes) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### SetRruleNil

`func (o *SLOCorrectionResponseAttributes) SetRruleNil(b bool)`

SetRruleNil sets the value for Rrule to be an explicit nil

### UnsetRrule

`func (o *SLOCorrectionResponseAttributes) UnsetRrule()`

UnsetRrule ensures that no value is present for Rrule, not even an explicit nil

### GetSloId

`func (o *SLOCorrectionResponseAttributes) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *SLOCorrectionResponseAttributes) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *SLOCorrectionResponseAttributes) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *SLOCorrectionResponseAttributes) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetStart

`func (o *SLOCorrectionResponseAttributes) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SLOCorrectionResponseAttributes) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SLOCorrectionResponseAttributes) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *SLOCorrectionResponseAttributes) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimezone

`func (o *SLOCorrectionResponseAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SLOCorrectionResponseAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SLOCorrectionResponseAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SLOCorrectionResponseAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
