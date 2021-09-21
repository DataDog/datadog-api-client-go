# FunnelQuery

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DataSource** | [**FunnelSource**](FunnelSource.md) |  | [default to FUNNELSOURCE_RUM]
**QueryString** | **string** | The widget query. | 
**Steps** | **[]interface{}** | List of funnel steps. | 

## Methods

### NewFunnelQuery

`func NewFunnelQuery(dataSource FunnelSource, queryString string, steps []interface{}) *FunnelQuery`

NewFunnelQuery instantiates a new FunnelQuery object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFunnelQueryWithDefaults

`func NewFunnelQueryWithDefaults() *FunnelQuery`

NewFunnelQueryWithDefaults instantiates a new FunnelQuery object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDataSource

`func (o *FunnelQuery) GetDataSource() FunnelSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FunnelQuery) GetDataSourceOk() (*FunnelSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FunnelQuery) SetDataSource(v FunnelSource)`

SetDataSource sets DataSource field to given value.


### GetQueryString

`func (o *FunnelQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *FunnelQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *FunnelQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetSteps

`func (o *FunnelQuery) GetSteps() []interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FunnelQuery) GetStepsOk() (*[]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FunnelQuery) SetSteps(v []interface{})`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


