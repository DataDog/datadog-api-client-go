# LogsMetricCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**LogsMetricCreateAttributes**](LogsMetricCreateAttributes.md) |  | 
**Id** | **string** | The name of the log-based metric. | 
**Type** | [**LogsMetricType**](LogsMetricType.md) |  | [default to "logs_metrics"]

## Methods

### NewLogsMetricCreateData

`func NewLogsMetricCreateData(attributes LogsMetricCreateAttributes, id string, type_ LogsMetricType, ) *LogsMetricCreateData`

NewLogsMetricCreateData instantiates a new LogsMetricCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricCreateDataWithDefaults

`func NewLogsMetricCreateDataWithDefaults() *LogsMetricCreateData`

NewLogsMetricCreateDataWithDefaults instantiates a new LogsMetricCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogsMetricCreateData) GetAttributes() LogsMetricCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsMetricCreateData) GetAttributesOk() (*LogsMetricCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsMetricCreateData) SetAttributes(v LogsMetricCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *LogsMetricCreateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsMetricCreateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsMetricCreateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LogsMetricCreateData) GetType() LogsMetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsMetricCreateData) GetTypeOk() (*LogsMetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsMetricCreateData) SetType(v LogsMetricType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


