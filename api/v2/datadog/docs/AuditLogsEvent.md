# AuditLogsEvent

## Properties

| Name           | Type                                                                   | Description             | Notes                                            |
| -------------- | ---------------------------------------------------------------------- | ----------------------- | ------------------------------------------------ |
| **Attributes** | Pointer to [**AuditLogsEventAttributes**](AuditLogsEventAttributes.md) |                         | [optional]                                       |
| **Id**         | Pointer to **string**                                                  | Unique ID of the event. | [optional]                                       |
| **Type**       | Pointer to [**AuditLogsEventType**](AuditLogsEventType.md)             |                         | [optional] [default to AUDITLOGSEVENTTYPE_Audit] |

## Methods

### NewAuditLogsEvent

`func NewAuditLogsEvent() *AuditLogsEvent`

NewAuditLogsEvent instantiates a new AuditLogsEvent object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsEventWithDefaults

`func NewAuditLogsEventWithDefaults() *AuditLogsEvent`

NewAuditLogsEventWithDefaults instantiates a new AuditLogsEvent object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *AuditLogsEvent) GetAttributes() AuditLogsEventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuditLogsEvent) GetAttributesOk() (*AuditLogsEventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuditLogsEvent) SetAttributes(v AuditLogsEventAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuditLogsEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *AuditLogsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AuditLogsEvent) GetType() AuditLogsEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLogsEvent) GetTypeOk() (*AuditLogsEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLogsEvent) SetType(v AuditLogsEventType)`

SetType sets Type field to given value.

### HasType

`func (o *AuditLogsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
