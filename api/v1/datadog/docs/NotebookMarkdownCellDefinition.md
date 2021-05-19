# NotebookMarkdownCellDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Text** | **string** | The markdown content. | 
**Type** | [**NotebookMarkdownCellDefinitionType**](NotebookMarkdownCellDefinitionType.md) |  | [default to NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN]

## Methods

### NewNotebookMarkdownCellDefinition

`func NewNotebookMarkdownCellDefinition(text string, type_ NotebookMarkdownCellDefinitionType, ) *NotebookMarkdownCellDefinition`

NewNotebookMarkdownCellDefinition instantiates a new NotebookMarkdownCellDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookMarkdownCellDefinitionWithDefaults

`func NewNotebookMarkdownCellDefinitionWithDefaults() *NotebookMarkdownCellDefinition`

NewNotebookMarkdownCellDefinitionWithDefaults instantiates a new NotebookMarkdownCellDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *NotebookMarkdownCellDefinition) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NotebookMarkdownCellDefinition) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NotebookMarkdownCellDefinition) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *NotebookMarkdownCellDefinition) GetType() NotebookMarkdownCellDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookMarkdownCellDefinition) GetTypeOk() (*NotebookMarkdownCellDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookMarkdownCellDefinition) SetType(v NotebookMarkdownCellDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


