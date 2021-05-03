# UsageCustomReportsResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]UsageCustomReportsData**](UsageCustomReportsData.md) | An array of available custom reports. | [optional] 
**Meta** | Pointer to [**UsageCustomReportsMeta**](UsageCustomReportsMeta.md) |  | [optional] 

## Methods

### NewUsageCustomReportsResponse

`func NewUsageCustomReportsResponse() *UsageCustomReportsResponse`

NewUsageCustomReportsResponse instantiates a new UsageCustomReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageCustomReportsResponseWithDefaults

`func NewUsageCustomReportsResponseWithDefaults() *UsageCustomReportsResponse`

NewUsageCustomReportsResponseWithDefaults instantiates a new UsageCustomReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsageCustomReportsResponse) GetData() []UsageCustomReportsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageCustomReportsResponse) GetDataOk() (*[]UsageCustomReportsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageCustomReportsResponse) SetData(v []UsageCustomReportsData)`

SetData sets Data field to given value.

### HasData

`func (o *UsageCustomReportsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *UsageCustomReportsResponse) GetMeta() UsageCustomReportsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsageCustomReportsResponse) GetMetaOk() (*UsageCustomReportsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsageCustomReportsResponse) SetMeta(v UsageCustomReportsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsageCustomReportsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


