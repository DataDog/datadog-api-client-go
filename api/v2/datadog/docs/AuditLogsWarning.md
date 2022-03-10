# AuditLogsWarning

## Properties

| Name       | Type                  | Description                                    | Notes      |
| ---------- | --------------------- | ---------------------------------------------- | ---------- |
| **Code**   | Pointer to **string** | Unique code for this type of warning.          | [optional] |
| **Detail** | Pointer to **string** | Detailed explanation of this specific warning. | [optional] |
| **Title**  | Pointer to **string** | Short human-readable summary of the warning.   | [optional] |

## Methods

### NewAuditLogsWarning

`func NewAuditLogsWarning() *AuditLogsWarning`

NewAuditLogsWarning instantiates a new AuditLogsWarning object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsWarningWithDefaults

`func NewAuditLogsWarningWithDefaults() *AuditLogsWarning`

NewAuditLogsWarningWithDefaults instantiates a new AuditLogsWarning object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCode

`func (o *AuditLogsWarning) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuditLogsWarning) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuditLogsWarning) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuditLogsWarning) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *AuditLogsWarning) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AuditLogsWarning) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AuditLogsWarning) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AuditLogsWarning) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetTitle

`func (o *AuditLogsWarning) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuditLogsWarning) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuditLogsWarning) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuditLogsWarning) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
