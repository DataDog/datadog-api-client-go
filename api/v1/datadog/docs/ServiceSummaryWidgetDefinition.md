# ServiceSummaryWidgetDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DisplayFormat** | Pointer to [**WidgetServiceSummaryDisplayFormat**](WidgetServiceSummaryDisplayFormat.md) |  | [optional] 
**Env** | **string** | APM environment. | 
**Service** | **string** | APM service. | 
**ShowBreakdown** | Pointer to **bool** | Whether to show the latency breakdown or not. | [optional] 
**ShowDistribution** | Pointer to **bool** | Whether to show the latency distribution or not. | [optional] 
**ShowErrors** | Pointer to **bool** | Whether to show the error metrics or not. | [optional] 
**ShowHits** | Pointer to **bool** | Whether to show the hits metrics or not. | [optional] 
**ShowLatency** | Pointer to **bool** | Whether to show the latency metrics or not. | [optional] 
**ShowResourceList** | Pointer to **bool** | Whether to show the resource list or not. | [optional] 
**SizeFormat** | Pointer to [**WidgetSizeFormat**](WidgetSizeFormat.md) |  | [optional] 
**SpanName** | **string** | APM span name. | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**ServiceSummaryWidgetDefinitionType**](ServiceSummaryWidgetDefinitionType.md) |  | [default to SERVICESUMMARYWIDGETDEFINITIONTYPE_TRACE_SERVICE]

## Methods

### NewServiceSummaryWidgetDefinition

`func NewServiceSummaryWidgetDefinition(env string, service string, spanName string, type_ ServiceSummaryWidgetDefinitionType, ) *ServiceSummaryWidgetDefinition`

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

`func (o *ServiceSummaryWidgetDefinition) GetDisplayFormatOk() (*WidgetServiceSummaryDisplayFormat, bool)`

GetDisplayFormatOk returns a tuple with the DisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFormat

`func (o *ServiceSummaryWidgetDefinition) SetDisplayFormat(v WidgetServiceSummaryDisplayFormat)`

SetDisplayFormat sets DisplayFormat field to given value.

### HasDisplayFormat

`func (o *ServiceSummaryWidgetDefinition) HasDisplayFormat() bool`

HasDisplayFormat returns a boolean if a field has been set.

### GetEnv

`func (o *ServiceSummaryWidgetDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServiceSummaryWidgetDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServiceSummaryWidgetDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetService

`func (o *ServiceSummaryWidgetDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceSummaryWidgetDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceSummaryWidgetDefinition) SetService(v string)`

SetService sets Service field to given value.


### GetShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) GetShowBreakdown() bool`

GetShowBreakdown returns the ShowBreakdown field if non-nil, zero value otherwise.

### GetShowBreakdownOk

`func (o *ServiceSummaryWidgetDefinition) GetShowBreakdownOk() (*bool, bool)`

GetShowBreakdownOk returns a tuple with the ShowBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) SetShowBreakdown(v bool)`

SetShowBreakdown sets ShowBreakdown field to given value.

### HasShowBreakdown

`func (o *ServiceSummaryWidgetDefinition) HasShowBreakdown() bool`

HasShowBreakdown returns a boolean if a field has been set.

### GetShowDistribution

`func (o *ServiceSummaryWidgetDefinition) GetShowDistribution() bool`

GetShowDistribution returns the ShowDistribution field if non-nil, zero value otherwise.

### GetShowDistributionOk

`func (o *ServiceSummaryWidgetDefinition) GetShowDistributionOk() (*bool, bool)`

GetShowDistributionOk returns a tuple with the ShowDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDistribution

`func (o *ServiceSummaryWidgetDefinition) SetShowDistribution(v bool)`

SetShowDistribution sets ShowDistribution field to given value.

### HasShowDistribution

`func (o *ServiceSummaryWidgetDefinition) HasShowDistribution() bool`

HasShowDistribution returns a boolean if a field has been set.

### GetShowErrors

`func (o *ServiceSummaryWidgetDefinition) GetShowErrors() bool`

GetShowErrors returns the ShowErrors field if non-nil, zero value otherwise.

### GetShowErrorsOk

`func (o *ServiceSummaryWidgetDefinition) GetShowErrorsOk() (*bool, bool)`

GetShowErrorsOk returns a tuple with the ShowErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowErrors

`func (o *ServiceSummaryWidgetDefinition) SetShowErrors(v bool)`

SetShowErrors sets ShowErrors field to given value.

### HasShowErrors

`func (o *ServiceSummaryWidgetDefinition) HasShowErrors() bool`

HasShowErrors returns a boolean if a field has been set.

### GetShowHits

`func (o *ServiceSummaryWidgetDefinition) GetShowHits() bool`

GetShowHits returns the ShowHits field if non-nil, zero value otherwise.

### GetShowHitsOk

`func (o *ServiceSummaryWidgetDefinition) GetShowHitsOk() (*bool, bool)`

GetShowHitsOk returns a tuple with the ShowHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHits

`func (o *ServiceSummaryWidgetDefinition) SetShowHits(v bool)`

SetShowHits sets ShowHits field to given value.

### HasShowHits

`func (o *ServiceSummaryWidgetDefinition) HasShowHits() bool`

HasShowHits returns a boolean if a field has been set.

### GetShowLatency

`func (o *ServiceSummaryWidgetDefinition) GetShowLatency() bool`

GetShowLatency returns the ShowLatency field if non-nil, zero value otherwise.

### GetShowLatencyOk

`func (o *ServiceSummaryWidgetDefinition) GetShowLatencyOk() (*bool, bool)`

GetShowLatencyOk returns a tuple with the ShowLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLatency

`func (o *ServiceSummaryWidgetDefinition) SetShowLatency(v bool)`

SetShowLatency sets ShowLatency field to given value.

### HasShowLatency

`func (o *ServiceSummaryWidgetDefinition) HasShowLatency() bool`

HasShowLatency returns a boolean if a field has been set.

### GetShowResourceList

`func (o *ServiceSummaryWidgetDefinition) GetShowResourceList() bool`

GetShowResourceList returns the ShowResourceList field if non-nil, zero value otherwise.

### GetShowResourceListOk

`func (o *ServiceSummaryWidgetDefinition) GetShowResourceListOk() (*bool, bool)`

GetShowResourceListOk returns a tuple with the ShowResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResourceList

`func (o *ServiceSummaryWidgetDefinition) SetShowResourceList(v bool)`

SetShowResourceList sets ShowResourceList field to given value.

### HasShowResourceList

`func (o *ServiceSummaryWidgetDefinition) HasShowResourceList() bool`

HasShowResourceList returns a boolean if a field has been set.

### GetSizeFormat

`func (o *ServiceSummaryWidgetDefinition) GetSizeFormat() WidgetSizeFormat`

GetSizeFormat returns the SizeFormat field if non-nil, zero value otherwise.

### GetSizeFormatOk

`func (o *ServiceSummaryWidgetDefinition) GetSizeFormatOk() (*WidgetSizeFormat, bool)`

GetSizeFormatOk returns a tuple with the SizeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeFormat

`func (o *ServiceSummaryWidgetDefinition) SetSizeFormat(v WidgetSizeFormat)`

SetSizeFormat sets SizeFormat field to given value.

### HasSizeFormat

`func (o *ServiceSummaryWidgetDefinition) HasSizeFormat() bool`

HasSizeFormat returns a boolean if a field has been set.

### GetSpanName

`func (o *ServiceSummaryWidgetDefinition) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *ServiceSummaryWidgetDefinition) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *ServiceSummaryWidgetDefinition) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetTime

`func (o *ServiceSummaryWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ServiceSummaryWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ServiceSummaryWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *ServiceSummaryWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *ServiceSummaryWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceSummaryWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceSummaryWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *ServiceSummaryWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *ServiceSummaryWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *ServiceSummaryWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *ServiceSummaryWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *ServiceSummaryWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *ServiceSummaryWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *ServiceSummaryWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *ServiceSummaryWidgetDefinition) GetType() ServiceSummaryWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceSummaryWidgetDefinition) GetTypeOk() (*ServiceSummaryWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceSummaryWidgetDefinition) SetType(v ServiceSummaryWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


