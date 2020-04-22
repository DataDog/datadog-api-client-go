# LogStreamWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | Which columns to display on the widget. | [optional] 
**Indexes** | Pointer to **[]string** | An array of index names to query in the stream. Use [] to query all indexes at once. | [optional] 
**MessageDisplay** | Pointer to [**WidgetMessageDisplay**](WidgetMessageDisplay.md) |  | [optional] 
**Query** | Pointer to **string** | Query to filter the log stream with. | [optional] 
**ShowDateColumn** | Pointer to **bool** | Whether to show the date column or not | [optional] 
**ShowMessageColumn** | Pointer to **bool** | Whether to show the message column or not | [optional] 
**Sort** | Pointer to [**WidgetFieldSort**](WidgetFieldSort.md) |  | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | Pointer to **string** | Type of the widget. | [readonly] [default to "log_stream"]

## Methods

### NewLogStreamWidgetDefinition

`func NewLogStreamWidgetDefinition(type_ string, ) *LogStreamWidgetDefinition`

NewLogStreamWidgetDefinition instantiates a new LogStreamWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamWidgetDefinitionWithDefaults

`func NewLogStreamWidgetDefinitionWithDefaults() *LogStreamWidgetDefinition`

NewLogStreamWidgetDefinitionWithDefaults instantiates a new LogStreamWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *LogStreamWidgetDefinition) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *LogStreamWidgetDefinition) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *LogStreamWidgetDefinition) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *LogStreamWidgetDefinition) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetIndexes

`func (o *LogStreamWidgetDefinition) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *LogStreamWidgetDefinition) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *LogStreamWidgetDefinition) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *LogStreamWidgetDefinition) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetMessageDisplay

`func (o *LogStreamWidgetDefinition) GetMessageDisplay() WidgetMessageDisplay`

GetMessageDisplay returns the MessageDisplay field if non-nil, zero value otherwise.

### GetMessageDisplayOk

`func (o *LogStreamWidgetDefinition) GetMessageDisplayOk() (*WidgetMessageDisplay, bool)`

GetMessageDisplayOk returns a tuple with the MessageDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDisplay

`func (o *LogStreamWidgetDefinition) SetMessageDisplay(v WidgetMessageDisplay)`

SetMessageDisplay sets MessageDisplay field to given value.

### HasMessageDisplay

`func (o *LogStreamWidgetDefinition) HasMessageDisplay() bool`

HasMessageDisplay returns a boolean if a field has been set.

### GetQuery

`func (o *LogStreamWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogStreamWidgetDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogStreamWidgetDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogStreamWidgetDefinition) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetShowDateColumn

`func (o *LogStreamWidgetDefinition) GetShowDateColumn() bool`

GetShowDateColumn returns the ShowDateColumn field if non-nil, zero value otherwise.

### GetShowDateColumnOk

`func (o *LogStreamWidgetDefinition) GetShowDateColumnOk() (*bool, bool)`

GetShowDateColumnOk returns a tuple with the ShowDateColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDateColumn

`func (o *LogStreamWidgetDefinition) SetShowDateColumn(v bool)`

SetShowDateColumn sets ShowDateColumn field to given value.

### HasShowDateColumn

`func (o *LogStreamWidgetDefinition) HasShowDateColumn() bool`

HasShowDateColumn returns a boolean if a field has been set.

### GetShowMessageColumn

`func (o *LogStreamWidgetDefinition) GetShowMessageColumn() bool`

GetShowMessageColumn returns the ShowMessageColumn field if non-nil, zero value otherwise.

### GetShowMessageColumnOk

`func (o *LogStreamWidgetDefinition) GetShowMessageColumnOk() (*bool, bool)`

GetShowMessageColumnOk returns a tuple with the ShowMessageColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMessageColumn

`func (o *LogStreamWidgetDefinition) SetShowMessageColumn(v bool)`

SetShowMessageColumn sets ShowMessageColumn field to given value.

### HasShowMessageColumn

`func (o *LogStreamWidgetDefinition) HasShowMessageColumn() bool`

HasShowMessageColumn returns a boolean if a field has been set.

### GetSort

`func (o *LogStreamWidgetDefinition) GetSort() WidgetFieldSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogStreamWidgetDefinition) GetSortOk() (*WidgetFieldSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LogStreamWidgetDefinition) SetSort(v WidgetFieldSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LogStreamWidgetDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTime

`func (o *LogStreamWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogStreamWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LogStreamWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *LogStreamWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *LogStreamWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogStreamWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogStreamWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LogStreamWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *LogStreamWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *LogStreamWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *LogStreamWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *LogStreamWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *LogStreamWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *LogStreamWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *LogStreamWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *LogStreamWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *LogStreamWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogStreamWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogStreamWidgetDefinition) SetType(v string)`

SetType sets Type field to given value.



### AsWidgetDefinition

`func (s *LogStreamWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of LogStreamWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


