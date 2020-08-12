# ApmResourcesQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | Column names used by front end for display. | [optional] 
**Env** | **string** | Environment name. | 
**Name** | **string** | Operation name associated with service. | 
**Resource** | Pointer to **string** | Resource name. | [optional] 
**Service** | **string** | Service name. | 

## Methods

### NewApmResourcesQueryDefinition

`func NewApmResourcesQueryDefinition(env string, name string, service string, ) *ApmResourcesQueryDefinition`

NewApmResourcesQueryDefinition instantiates a new ApmResourcesQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApmResourcesQueryDefinitionWithDefaults

`func NewApmResourcesQueryDefinitionWithDefaults() *ApmResourcesQueryDefinition`

NewApmResourcesQueryDefinitionWithDefaults instantiates a new ApmResourcesQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ApmResourcesQueryDefinition) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ApmResourcesQueryDefinition) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ApmResourcesQueryDefinition) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ApmResourcesQueryDefinition) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEnv

`func (o *ApmResourcesQueryDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ApmResourcesQueryDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ApmResourcesQueryDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetName

`func (o *ApmResourcesQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApmResourcesQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApmResourcesQueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetResource

`func (o *ApmResourcesQueryDefinition) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApmResourcesQueryDefinition) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApmResourcesQueryDefinition) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApmResourcesQueryDefinition) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetService

`func (o *ApmResourcesQueryDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ApmResourcesQueryDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ApmResourcesQueryDefinition) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


