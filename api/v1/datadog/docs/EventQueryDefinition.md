# EventQueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** |  | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or | 

## Methods

### NewEventQueryDefinition

`func NewEventQueryDefinition(search string, tagsExecution string, ) *EventQueryDefinition`

NewEventQueryDefinition instantiates a new EventQueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryDefinitionWithDefaults

`func NewEventQueryDefinitionWithDefaults() *EventQueryDefinition`

NewEventQueryDefinitionWithDefaults instantiates a new EventQueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *EventQueryDefinition) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *EventQueryDefinition) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *EventQueryDefinition) SetSearch(v string)`

SetSearch sets Search field to given value.


### GetTagsExecution

`func (o *EventQueryDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *EventQueryDefinition) GetTagsExecutionOk() (*string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsExecution

`func (o *EventQueryDefinition) SetTagsExecution(v string)`

SetTagsExecution sets TagsExecution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


