# MetricsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Message indicating the errors if status is not &#x60;ok&#x60; | [optional] [readonly] 
**FromDate** | Pointer to **int64** | Start of requested time window, milliseconds since Unix epoch | [optional] [readonly] 
**GroupBy** | Pointer to **[]string** | List of tag keys on which to group | [optional] [readonly] 
**Message** | Pointer to **string** | Message indicating &#x60;success&#x60; if status is &#x60;ok&#x60; | [optional] [readonly] 
**Query** | Pointer to **string** | Query string | [optional] [readonly] 
**ResType** | Pointer to **string** | Type of response | [optional] [readonly] 
**Series** | Pointer to [**[]MetricsQueryResponseSeries**](MetricsQueryResponse_series.md) | List of timeseries queried | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the query | [optional] [readonly] 
**ToDate** | Pointer to **int64** | End of requested time window, milliseconds since Unix epoch | [optional] [readonly] 

## Methods

### NewMetricsQueryResponse

`func NewMetricsQueryResponse() *MetricsQueryResponse`

NewMetricsQueryResponse instantiates a new MetricsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResponseWithDefaults

`func NewMetricsQueryResponseWithDefaults() *MetricsQueryResponse`

NewMetricsQueryResponseWithDefaults instantiates a new MetricsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MetricsQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MetricsQueryResponse) GetErrorOk() (string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MetricsQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MetricsQueryResponse) SetError(v string)`

SetError gets a reference to the given string and assigns it to the Error field.

### GetFromDate

`func (o *MetricsQueryResponse) GetFromDate() int64`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *MetricsQueryResponse) GetFromDateOk() (int64, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromDate

`func (o *MetricsQueryResponse) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDate

`func (o *MetricsQueryResponse) SetFromDate(v int64)`

SetFromDate gets a reference to the given int64 and assigns it to the FromDate field.

### GetGroupBy

`func (o *MetricsQueryResponse) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricsQueryResponse) GetGroupByOk() ([]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupBy

`func (o *MetricsQueryResponse) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupBy

`func (o *MetricsQueryResponse) SetGroupBy(v []string)`

SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.

### GetMessage

`func (o *MetricsQueryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricsQueryResponse) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MetricsQueryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MetricsQueryResponse) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetQuery

`func (o *MetricsQueryResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsQueryResponse) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *MetricsQueryResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *MetricsQueryResponse) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetResType

`func (o *MetricsQueryResponse) GetResType() string`

GetResType returns the ResType field if non-nil, zero value otherwise.

### GetResTypeOk

`func (o *MetricsQueryResponse) GetResTypeOk() (string, bool)`

GetResTypeOk returns a tuple with the ResType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResType

`func (o *MetricsQueryResponse) HasResType() bool`

HasResType returns a boolean if a field has been set.

### SetResType

`func (o *MetricsQueryResponse) SetResType(v string)`

SetResType gets a reference to the given string and assigns it to the ResType field.

### GetSeries

`func (o *MetricsQueryResponse) GetSeries() []MetricsQueryResponseSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MetricsQueryResponse) GetSeriesOk() ([]MetricsQueryResponseSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeries

`func (o *MetricsQueryResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeries

`func (o *MetricsQueryResponse) SetSeries(v []MetricsQueryResponseSeries)`

SetSeries gets a reference to the given []MetricsQueryResponseSeries and assigns it to the Series field.

### GetStatus

`func (o *MetricsQueryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricsQueryResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MetricsQueryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MetricsQueryResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetToDate

`func (o *MetricsQueryResponse) GetToDate() int64`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *MetricsQueryResponse) GetToDateOk() (int64, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToDate

`func (o *MetricsQueryResponse) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDate

`func (o *MetricsQueryResponse) SetToDate(v int64)`

SetToDate gets a reference to the given int64 and assigns it to the ToDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


