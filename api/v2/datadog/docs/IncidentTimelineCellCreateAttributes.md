# IncidentTimelineCellCreateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellType** | [**IncidentTimelineCellMarkdownContentType**](IncidentTimelineCellMarkdownContentType.md) |  | [default to "markdown"]
**Content** | [**IncidentTimelineCellMarkdownCreateAttributesContent**](IncidentTimelineCellMarkdownCreateAttributes_content.md) |  | 
**Important** | Pointer to **bool** | A flag indicating whether the timeline cell is important and should be highlighted. | [optional] [default to false]

## Methods

### NewIncidentTimelineCellCreateAttributes

`func NewIncidentTimelineCellCreateAttributes(cellType IncidentTimelineCellMarkdownContentType, content IncidentTimelineCellMarkdownCreateAttributesContent, ) *IncidentTimelineCellCreateAttributes`

NewIncidentTimelineCellCreateAttributes instantiates a new IncidentTimelineCellCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTimelineCellCreateAttributesWithDefaults

`func NewIncidentTimelineCellCreateAttributesWithDefaults() *IncidentTimelineCellCreateAttributes`

NewIncidentTimelineCellCreateAttributesWithDefaults instantiates a new IncidentTimelineCellCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellType

`func (o *IncidentTimelineCellCreateAttributes) GetCellType() IncidentTimelineCellMarkdownContentType`

GetCellType returns the CellType field if non-nil, zero value otherwise.

### GetCellTypeOk

`func (o *IncidentTimelineCellCreateAttributes) GetCellTypeOk() (*IncidentTimelineCellMarkdownContentType, bool)`

GetCellTypeOk returns a tuple with the CellType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellType

`func (o *IncidentTimelineCellCreateAttributes) SetCellType(v IncidentTimelineCellMarkdownContentType)`

SetCellType sets CellType field to given value.


### GetContent

`func (o *IncidentTimelineCellCreateAttributes) GetContent() IncidentTimelineCellMarkdownCreateAttributesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentTimelineCellCreateAttributes) GetContentOk() (*IncidentTimelineCellMarkdownCreateAttributesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentTimelineCellCreateAttributes) SetContent(v IncidentTimelineCellMarkdownCreateAttributesContent)`

SetContent sets Content field to given value.


### GetImportant

`func (o *IncidentTimelineCellCreateAttributes) GetImportant() bool`

GetImportant returns the Important field if non-nil, zero value otherwise.

### GetImportantOk

`func (o *IncidentTimelineCellCreateAttributes) GetImportantOk() (*bool, bool)`

GetImportantOk returns a tuple with the Important field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportant

`func (o *IncidentTimelineCellCreateAttributes) SetImportant(v bool)`

SetImportant sets Important field to given value.

### HasImportant

`func (o *IncidentTimelineCellCreateAttributes) HasImportant() bool`

HasImportant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


