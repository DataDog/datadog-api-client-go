# HostMapWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md) | List of custom links. | [optional] 
**Group** | Pointer to **[]string** | List of tag prefixes to group by. | [optional] 
**NoGroupHosts** | Pointer to **bool** | Whether to show the hosts that don’t fit in a group. | [optional] 
**NoMetricHosts** | Pointer to **bool** | Whether to show the hosts with no metrics. | [optional] 
**NodeType** | Pointer to [**WidgetNodeType**](WidgetNodeType.md) |  | [optional] 
**Notes** | Pointer to **string** | Notes on the title. | [optional] 
**Requests** | [**HostMapWidgetDefinitionRequests**](HostMapWidgetDefinitionRequests.md) |  | 
**Scope** | Pointer to **[]string** | List of tags used to filter the map. | [optional] 
**Style** | Pointer to [**HostMapWidgetDefinitionStyle**](HostMapWidgetDefinitionStyle.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**HostMapWidgetDefinitionType**](HostMapWidgetDefinitionType.md) |  | [default to HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP]

## Methods

### NewHostMapWidgetDefinition

`func NewHostMapWidgetDefinition(requests HostMapWidgetDefinitionRequests, type_ HostMapWidgetDefinitionType, ) *HostMapWidgetDefinition`

NewHostMapWidgetDefinition instantiates a new HostMapWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMapWidgetDefinitionWithDefaults

`func NewHostMapWidgetDefinitionWithDefaults() *HostMapWidgetDefinition`

NewHostMapWidgetDefinitionWithDefaults instantiates a new HostMapWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomLinks

`func (o *HostMapWidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *HostMapWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *HostMapWidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *HostMapWidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetGroup

`func (o *HostMapWidgetDefinition) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *HostMapWidgetDefinition) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *HostMapWidgetDefinition) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *HostMapWidgetDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetNoGroupHosts

`func (o *HostMapWidgetDefinition) GetNoGroupHosts() bool`

GetNoGroupHosts returns the NoGroupHosts field if non-nil, zero value otherwise.

### GetNoGroupHostsOk

`func (o *HostMapWidgetDefinition) GetNoGroupHostsOk() (*bool, bool)`

GetNoGroupHostsOk returns a tuple with the NoGroupHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoGroupHosts

`func (o *HostMapWidgetDefinition) SetNoGroupHosts(v bool)`

SetNoGroupHosts sets NoGroupHosts field to given value.

### HasNoGroupHosts

`func (o *HostMapWidgetDefinition) HasNoGroupHosts() bool`

HasNoGroupHosts returns a boolean if a field has been set.

### GetNoMetricHosts

`func (o *HostMapWidgetDefinition) GetNoMetricHosts() bool`

GetNoMetricHosts returns the NoMetricHosts field if non-nil, zero value otherwise.

### GetNoMetricHostsOk

`func (o *HostMapWidgetDefinition) GetNoMetricHostsOk() (*bool, bool)`

GetNoMetricHostsOk returns a tuple with the NoMetricHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMetricHosts

`func (o *HostMapWidgetDefinition) SetNoMetricHosts(v bool)`

SetNoMetricHosts sets NoMetricHosts field to given value.

### HasNoMetricHosts

`func (o *HostMapWidgetDefinition) HasNoMetricHosts() bool`

HasNoMetricHosts returns a boolean if a field has been set.

### GetNodeType

`func (o *HostMapWidgetDefinition) GetNodeType() WidgetNodeType`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *HostMapWidgetDefinition) GetNodeTypeOk() (*WidgetNodeType, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *HostMapWidgetDefinition) SetNodeType(v WidgetNodeType)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *HostMapWidgetDefinition) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetNotes

`func (o *HostMapWidgetDefinition) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *HostMapWidgetDefinition) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *HostMapWidgetDefinition) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *HostMapWidgetDefinition) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRequests

`func (o *HostMapWidgetDefinition) GetRequests() HostMapWidgetDefinitionRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HostMapWidgetDefinition) GetRequestsOk() (*HostMapWidgetDefinitionRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *HostMapWidgetDefinition) SetRequests(v HostMapWidgetDefinitionRequests)`

SetRequests sets Requests field to given value.


### GetScope

`func (o *HostMapWidgetDefinition) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HostMapWidgetDefinition) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HostMapWidgetDefinition) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HostMapWidgetDefinition) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStyle

`func (o *HostMapWidgetDefinition) GetStyle() HostMapWidgetDefinitionStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *HostMapWidgetDefinition) GetStyleOk() (*HostMapWidgetDefinitionStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *HostMapWidgetDefinition) SetStyle(v HostMapWidgetDefinitionStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *HostMapWidgetDefinition) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetTitle

`func (o *HostMapWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HostMapWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HostMapWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HostMapWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *HostMapWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *HostMapWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *HostMapWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *HostMapWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *HostMapWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *HostMapWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *HostMapWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *HostMapWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *HostMapWidgetDefinition) GetType() HostMapWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostMapWidgetDefinition) GetTypeOk() (*HostMapWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostMapWidgetDefinition) SetType(v HostMapWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


