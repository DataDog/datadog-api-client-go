# IncidentTimelineCellMarkdownCreateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CellType** | [**IncidentTimelineCellMarkdownContentType**](IncidentTimelineCellMarkdownContentType.md) |  | [default to INCIDENTTIMELINECELLMARKDOWNCONTENTTYPE_MARKDOWN]
**Content** | [**IncidentTimelineCellMarkdownCreateAttributesContent**](IncidentTimelineCellMarkdownCreateAttributesContent.md) |  | 
**Important** | Pointer to **bool** | A flag indicating whether the timeline cell is important and should be highlighted. | [optional] [default to false]

## Methods

### NewIncidentTimelineCellMarkdownCreateAttributes

`func NewIncidentTimelineCellMarkdownCreateAttributes(cellType IncidentTimelineCellMarkdownContentType, content IncidentTimelineCellMarkdownCreateAttributesContent, ) *IncidentTimelineCellMarkdownCreateAttributes`

NewIncidentTimelineCellMarkdownCreateAttributes instantiates a new IncidentTimelineCellMarkdownCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults

`func NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults() *IncidentTimelineCellMarkdownCreateAttributes`

NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults instantiates a new IncidentTimelineCellMarkdownCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellType

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellType() IncidentTimelineCellMarkdownContentType`

GetCellType returns the CellType field if non-nil, zero value otherwise.

### GetCellTypeOk

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellTypeOk() (*IncidentTimelineCellMarkdownContentType, bool)`

GetCellTypeOk returns a tuple with the CellType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellType

`func (o *IncidentTimelineCellMarkdownCreateAttributes) SetCellType(v IncidentTimelineCellMarkdownContentType)`

SetCellType sets CellType field to given value.


### GetContent

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContent() IncidentTimelineCellMarkdownCreateAttributesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContentOk() (*IncidentTimelineCellMarkdownCreateAttributesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentTimelineCellMarkdownCreateAttributes) SetContent(v IncidentTimelineCellMarkdownCreateAttributesContent)`

SetContent sets Content field to given value.


### GetImportant

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportant() bool`

GetImportant returns the Important field if non-nil, zero value otherwise.

### GetImportantOk

`func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportantOk() (*bool, bool)`

GetImportantOk returns a tuple with the Important field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportant

`func (o *IncidentTimelineCellMarkdownCreateAttributes) SetImportant(v bool)`

SetImportant sets Important field to given value.

### HasImportant

`func (o *IncidentTimelineCellMarkdownCreateAttributes) HasImportant() bool`

HasImportant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


