# ListStreamWidgetRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Columns** | [**[]ListStreamColumn**](ListStreamColumn.md) | Widget columns. | 
**Query** | [**ListStreamQuery**](ListStreamQuery.md) |  | 
**ResponseFormat** | [**ListStreamResponseFormat**](ListStreamResponseFormat.md) |  | 

## Methods

### NewListStreamWidgetRequest

`func NewListStreamWidgetRequest(columns []ListStreamColumn, query ListStreamQuery, responseFormat ListStreamResponseFormat) *ListStreamWidgetRequest`

NewListStreamWidgetRequest instantiates a new ListStreamWidgetRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListStreamWidgetRequestWithDefaults

`func NewListStreamWidgetRequestWithDefaults() *ListStreamWidgetRequest`

NewListStreamWidgetRequestWithDefaults instantiates a new ListStreamWidgetRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetColumns

`func (o *ListStreamWidgetRequest) GetColumns() []ListStreamColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ListStreamWidgetRequest) GetColumnsOk() (*[]ListStreamColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ListStreamWidgetRequest) SetColumns(v []ListStreamColumn)`

SetColumns sets Columns field to given value.


### GetQuery

`func (o *ListStreamWidgetRequest) GetQuery() ListStreamQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListStreamWidgetRequest) GetQueryOk() (*ListStreamQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListStreamWidgetRequest) SetQuery(v ListStreamQuery)`

SetQuery sets Query field to given value.


### GetResponseFormat

`func (o *ListStreamWidgetRequest) GetResponseFormat() ListStreamResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ListStreamWidgetRequest) GetResponseFormatOk() (*ListStreamResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ListStreamWidgetRequest) SetResponseFormat(v ListStreamResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


