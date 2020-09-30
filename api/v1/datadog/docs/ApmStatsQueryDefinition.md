# ApmStatsQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | Column names used by front end for display. | [optional] 
**Env** | **string** | Environment name. | 
**Name** | **string** | Operation name associated with service. | 
**PrimaryTag** | **string** | The organization&#39;s host group name and value. | 
**Resource** | Pointer to **string** | Resource name. | [optional] 
**RowType** | [**ApmStatsQueryRowType**](ApmStatsQueryRowType.md) |  | 
**Service** | **string** | Service name. | 

## Methods

### NewApmStatsQueryDefinition

`func NewApmStatsQueryDefinition(env string, name string, primaryTag string, rowType ApmStatsQueryRowType, service string, ) *ApmStatsQueryDefinition`

NewApmStatsQueryDefinition instantiates a new ApmStatsQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApmStatsQueryDefinitionWithDefaults

`func NewApmStatsQueryDefinitionWithDefaults() *ApmStatsQueryDefinition`

NewApmStatsQueryDefinitionWithDefaults instantiates a new ApmStatsQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ApmStatsQueryDefinition) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ApmStatsQueryDefinition) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ApmStatsQueryDefinition) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ApmStatsQueryDefinition) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEnv

`func (o *ApmStatsQueryDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ApmStatsQueryDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ApmStatsQueryDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetName

`func (o *ApmStatsQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApmStatsQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApmStatsQueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryTag

`func (o *ApmStatsQueryDefinition) GetPrimaryTag() string`

GetPrimaryTag returns the PrimaryTag field if non-nil, zero value otherwise.

### GetPrimaryTagOk

`func (o *ApmStatsQueryDefinition) GetPrimaryTagOk() (*string, bool)`

GetPrimaryTagOk returns a tuple with the PrimaryTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTag

`func (o *ApmStatsQueryDefinition) SetPrimaryTag(v string)`

SetPrimaryTag sets PrimaryTag field to given value.


### GetResource

`func (o *ApmStatsQueryDefinition) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApmStatsQueryDefinition) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApmStatsQueryDefinition) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApmStatsQueryDefinition) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRowType

`func (o *ApmStatsQueryDefinition) GetRowType() ApmStatsQueryRowType`

GetRowType returns the RowType field if non-nil, zero value otherwise.

### GetRowTypeOk

`func (o *ApmStatsQueryDefinition) GetRowTypeOk() (*ApmStatsQueryRowType, bool)`

GetRowTypeOk returns a tuple with the RowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowType

`func (o *ApmStatsQueryDefinition) SetRowType(v ApmStatsQueryRowType)`

SetRowType sets RowType field to given value.


### GetService

`func (o *ApmStatsQueryDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ApmStatsQueryDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ApmStatsQueryDefinition) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


