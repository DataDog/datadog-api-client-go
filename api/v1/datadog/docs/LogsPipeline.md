# LogsPipeline

## Properties

| Name           | Type                                               | Description                                  | Notes                 |
| -------------- | -------------------------------------------------- | -------------------------------------------- | --------------------- |
| **Filter**     | Pointer to [**LogsFilter**](LogsFilter.md)         |                                              | [optional]            |
| **Id**         | Pointer to **string**                              | ID of the pipeline.                          | [optional] [readonly] |
| **IsEnabled**  | Pointer to **bool**                                | Whether or not the pipeline is enabled.      | [optional]            |
| **IsReadOnly** | Pointer to **bool**                                | Whether or not the pipeline can be edited.   | [optional] [readonly] |
| **Name**       | **string**                                         | Name of the pipeline.                        |
| **Processors** | Pointer to [**[]LogsProcessor**](LogsProcessor.md) | Ordered list of processors in this pipeline. | [optional]            |
| **Type**       | Pointer to **string**                              | Type of pipeline.                            | [optional] [readonly] |

## Methods

### NewLogsPipeline

`func NewLogsPipeline(name string) *LogsPipeline`

NewLogsPipeline instantiates a new LogsPipeline object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsPipelineWithDefaults

`func NewLogsPipelineWithDefaults() *LogsPipeline`

NewLogsPipelineWithDefaults instantiates a new LogsPipeline object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilter

`func (o *LogsPipeline) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsPipeline) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsPipeline) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsPipeline) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *LogsPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogsPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *LogsPipeline) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsPipeline) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsPipeline) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsPipeline) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *LogsPipeline) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *LogsPipeline) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *LogsPipeline) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *LogsPipeline) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetName

`func (o *LogsPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsPipeline) SetName(v string)`

SetName sets Name field to given value.

### GetProcessors

`func (o *LogsPipeline) GetProcessors() []LogsProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *LogsPipeline) GetProcessorsOk() (*[]LogsProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *LogsPipeline) SetProcessors(v []LogsProcessor)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *LogsPipeline) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetType

`func (o *LogsPipeline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsPipeline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsPipeline) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogsPipeline) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
