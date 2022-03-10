# AuditLogsResponseMetadata

## Properties

| Name          | Type                                                                 | Description                                                                                                            | Notes      |
| ------------- | -------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Elapsed**   | Pointer to **int64**                                                 | Time elapsed in milliseconds.                                                                                          | [optional] |
| **Page**      | Pointer to [**AuditLogsResponsePage**](AuditLogsResponsePage.md)     |                                                                                                                        | [optional] |
| **RequestId** | Pointer to **string**                                                | The identifier of the request.                                                                                         | [optional] |
| **Status**    | Pointer to [**AuditLogsResponseStatus**](AuditLogsResponseStatus.md) |                                                                                                                        | [optional] |
| **Warnings**  | Pointer to [**[]AuditLogsWarning**](AuditLogsWarning.md)             | A list of warnings (non-fatal errors) encountered. Partial results may return if warnings are present in the response. | [optional] |

## Methods

### NewAuditLogsResponseMetadata

`func NewAuditLogsResponseMetadata() *AuditLogsResponseMetadata`

NewAuditLogsResponseMetadata instantiates a new AuditLogsResponseMetadata object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsResponseMetadataWithDefaults

`func NewAuditLogsResponseMetadataWithDefaults() *AuditLogsResponseMetadata`

NewAuditLogsResponseMetadataWithDefaults instantiates a new AuditLogsResponseMetadata object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetElapsed

`func (o *AuditLogsResponseMetadata) GetElapsed() int64`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *AuditLogsResponseMetadata) GetElapsedOk() (*int64, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *AuditLogsResponseMetadata) SetElapsed(v int64)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *AuditLogsResponseMetadata) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetPage

`func (o *AuditLogsResponseMetadata) GetPage() AuditLogsResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AuditLogsResponseMetadata) GetPageOk() (*AuditLogsResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AuditLogsResponseMetadata) SetPage(v AuditLogsResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *AuditLogsResponseMetadata) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRequestId

`func (o *AuditLogsResponseMetadata) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AuditLogsResponseMetadata) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AuditLogsResponseMetadata) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AuditLogsResponseMetadata) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *AuditLogsResponseMetadata) GetStatus() AuditLogsResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditLogsResponseMetadata) GetStatusOk() (*AuditLogsResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditLogsResponseMetadata) SetStatus(v AuditLogsResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditLogsResponseMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarnings

`func (o *AuditLogsResponseMetadata) GetWarnings() []AuditLogsWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AuditLogsResponseMetadata) GetWarningsOk() (*[]AuditLogsWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AuditLogsResponseMetadata) SetWarnings(v []AuditLogsWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AuditLogsResponseMetadata) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
