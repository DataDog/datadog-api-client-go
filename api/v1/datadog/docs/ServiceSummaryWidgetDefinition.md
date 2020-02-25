# ServiceSummaryWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayFormat** | Pointer to [**WidgetServiceSummaryDisplayFormat**](WidgetServiceSummaryDisplayFormat.md) |  | [optional] 
**Env** | Pointer to **string** | APM environment | 
**Service** | Pointer to **string** | APM service | 
**ShowBreakdown** | Pointer to **bool** | Whether to show the latency breakdown or not | [optional] 
**ShowDistribution** | Pointer to **bool** | Whether to show the latency distribution or not | [optional] 
**ShowErrors** | Pointer to **bool** | Whether to show the error metrics or not | [optional] 
**ShowHits** | Pointer to **bool** | Whether to show the hits metrics or not | [optional] 
**ShowLatency** | Pointer to **bool** | Whether to show the latency metrics or not | [optional] 
**ShowResourceList** | Pointer to **bool** | Whether to show the resource list or not | [optional] 
**SizeFormat** | Pointer to [**WidgetSizeFormat**](WidgetSizeFormat.md) |  | [optional] 
**SpanName** | Pointer to **string** | APM span name | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "trace_service"]

## Methods

### NewServiceSummaryWidgetDefinition

`func NewServiceSummaryWidgetDefinition(env string, service string, spanName string, type_ string, ) *ServiceSummaryWidgetDefinition`

NewServiceSummaryWidgetDefinition instantiates a new ServiceSummaryWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSummaryWidgetDefinitionWithDefaults

`func NewServiceSummaryWidgetDefinitionWithDefaults() *ServiceSummaryWidgetDefinition`

NewServiceSummaryWidgetDefinitionWithDefaults instantiates a new ServiceSummaryWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayFormat

`func (o *ServiceSummaryWidgetDefinition) GetDisplayFormat() WidgetServiceSummaryDisplayFormat`

GetDisplayFormat returns the DisplayFormat field if non-nil, zero value otherwise.

### GetDisplayFormatOk

`func (o *ServiceSummaryWidgetDefinition) GetDisplayFormatOk() (WidgetServiceSummaryDisplayFormat, bool)`

GetDisplayFormatOk returns a tuple with the DisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayFormat

`func (o *ServiceSummaryWidgetDefinition) HasDisplayFormat() bool`

HasDisplayFormat returns a boolean if a field has been set.

### SetDisplayFormat

`func (o *ServiceSummaryWidgetDefinition) SetDisplayFormat(v WidgetServiceSummaryDisplayFormat)`

SetDisplayFormat gets a reference to the given WidgetServiceSummaryDisplayFormat and assigns it to the DisplayFormat field.

### GetEnv

`func (o *ServiceSummaryWidgetDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServiceSummaryWidgetDefinition) GetEnvOk() (string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnv

`func (o *ServiceSummaryWidgetDefinition) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnv

`func (o *ServiceSummaryWidgetDefinition) SetEnv(v string)`

SetEnv gets a reference to the given string and assigns it to the Env field.

### GetService

`func (o *ServiceSummaryWidgetDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceSummaryWidgetDefinition) GetServiceOk() (string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasService

`func (o *ServiceSummaryWidgetDefinition) HasService() bool`

HasService returns a boolean if a field has been set.

### SetService

`func (o *ServiceSummaryWidgetDefinition) SetService(v string)`

SetService gets a reference to the given string and assigns it to the Service field.

### GetShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) GetShowBreakdown() bool`

GetShowBreakdown returns the ShowBreakdown field if non-nil, zero value otherwise.

### GetShowBreakdownOk

`func (o *ServiceSummaryWidgetDefinition) GetShowBreakdownOk() (bool, bool)`

GetShowBreakdownOk returns a tuple with the ShowBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) HasShowBreakdown() bool`

HasShowBreakdown returns a boolean if a field has been set.

### SetShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) SetShowBreakdown(v bool)`

SetShowBreakdown gets a reference to the given bool and assigns it to the ShowBreakdown field.

### GetShowDistribution

`func (o *ServiceSummaryWidgetDefinition) GetShowDistribution() bool`

GetShowDistribution returns the ShowDistribution field if non-nil, zero value otherwise.

### GetShowDistributionOk

`func (o *ServiceSummaryWidgetDefinition) GetShowDistributionOk() (bool, bool)`

GetShowDistributionOk returns a tuple with the ShowDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowDistribution

`func (o *ServiceSummaryWidgetDefinition) HasShowDistribution() bool`

HasShowDistribution returns a boolean if a field has been set.

### SetShowDistribution

`func (o *ServiceSummaryWidgetDefinition) SetShowDistribution(v bool)`

SetShowDistribution gets a reference to the given bool and assigns it to the ShowDistribution field.

### GetShowErrors

`func (o *ServiceSummaryWidgetDefinition) GetShowErrors() bool`

GetShowErrors returns the ShowErrors field if non-nil, zero value otherwise.

### GetShowErrorsOk

`func (o *ServiceSummaryWidgetDefinition) GetShowErrorsOk() (bool, bool)`

GetShowErrorsOk returns a tuple with the ShowErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowErrors

`func (o *ServiceSummaryWidgetDefinition) HasShowErrors() bool`

HasShowErrors returns a boolean if a field has been set.

### SetShowErrors

`func (o *ServiceSummaryWidgetDefinition) SetShowErrors(v bool)`

SetShowErrors gets a reference to the given bool and assigns it to the ShowErrors field.

### GetShowHits

`func (o *ServiceSummaryWidgetDefinition) GetShowHits() bool`

GetShowHits returns the ShowHits field if non-nil, zero value otherwise.

### GetShowHitsOk

`func (o *ServiceSummaryWidgetDefinition) GetShowHitsOk() (bool, bool)`

GetShowHitsOk returns a tuple with the ShowHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowHits

`func (o *ServiceSummaryWidgetDefinition) HasShowHits() bool`

HasShowHits returns a boolean if a field has been set.

### SetShowHits

`func (o *ServiceSummaryWidgetDefinition) SetShowHits(v bool)`

SetShowHits gets a reference to the given bool and assigns it to the ShowHits field.

### GetShowLatency

`func (o *ServiceSummaryWidgetDefinition) GetShowLatency() bool`

GetShowLatency returns the ShowLatency field if non-nil, zero value otherwise.

### GetShowLatencyOk

`func (o *ServiceSummaryWidgetDefinition) GetShowLatencyOk() (bool, bool)`

GetShowLatencyOk returns a tuple with the ShowLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowLatency

`func (o *ServiceSummaryWidgetDefinition) HasShowLatency() bool`

HasShowLatency returns a boolean if a field has been set.

### SetShowLatency

`func (o *ServiceSummaryWidgetDefinition) SetShowLatency(v bool)`

SetShowLatency gets a reference to the given bool and assigns it to the ShowLatency field.

### GetShowResourceList

`func (o *ServiceSummaryWidgetDefinition) GetShowResourceList() bool`

GetShowResourceList returns the ShowResourceList field if non-nil, zero value otherwise.

### GetShowResourceListOk

`func (o *ServiceSummaryWidgetDefinition) GetShowResourceListOk() (bool, bool)`

GetShowResourceListOk returns a tuple with the ShowResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowResourceList

`func (o *ServiceSummaryWidgetDefinition) HasShowResourceList() bool`

HasShowResourceList returns a boolean if a field has been set.

### SetShowResourceList

`func (o *ServiceSummaryWidgetDefinition) SetShowResourceList(v bool)`

SetShowResourceList gets a reference to the given bool and assigns it to the ShowResourceList field.

### GetSizeFormat

`func (o *ServiceSummaryWidgetDefinition) GetSizeFormat() WidgetSizeFormat`

GetSizeFormat returns the SizeFormat field if non-nil, zero value otherwise.

### GetSizeFormatOk

`func (o *ServiceSummaryWidgetDefinition) GetSizeFormatOk() (WidgetSizeFormat, bool)`

GetSizeFormatOk returns a tuple with the SizeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSizeFormat

`func (o *ServiceSummaryWidgetDefinition) HasSizeFormat() bool`

HasSizeFormat returns a boolean if a field has been set.

### SetSizeFormat

`func (o *ServiceSummaryWidgetDefinition) SetSizeFormat(v WidgetSizeFormat)`

SetSizeFormat gets a reference to the given WidgetSizeFormat and assigns it to the SizeFormat field.

### GetSpanName

`func (o *ServiceSummaryWidgetDefinition) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *ServiceSummaryWidgetDefinition) GetSpanNameOk() (string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpanName

`func (o *ServiceSummaryWidgetDefinition) HasSpanName() bool`

HasSpanName returns a boolean if a field has been set.

### SetSpanName

`func (o *ServiceSummaryWidgetDefinition) SetSpanName(v string)`

SetSpanName gets a reference to the given string and assigns it to the SpanName field.

### GetTime

`func (o *ServiceSummaryWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ServiceSummaryWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *ServiceSummaryWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *ServiceSummaryWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *ServiceSummaryWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *ServiceSummaryWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *ServiceSummaryWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *ServiceSummaryWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *ServiceSummaryWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *ServiceSummaryWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *ServiceSummaryWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *ServiceSummaryWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *ServiceSummaryWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *ServiceSummaryWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceSummaryWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ServiceSummaryWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ServiceSummaryWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *ServiceSummaryWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of ServiceSummaryWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


