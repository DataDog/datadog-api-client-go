# HostMapWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **[]string** | List of tag prefixes to group by. | [optional] 
**NoGroupHosts** | Pointer to **bool** | Whether to show the hosts that don’t fit in a group. | [optional] 
**NoMetricHosts** | Pointer to **bool** | Whether to show the hosts with no metrics. | [optional] 
**NodeType** | Pointer to [**WidgetNodeType**](WidgetNodeType.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to [**HostMapWidgetDefinitionRequests**](HostMapWidgetDefinition_requests.md) |  | 
**Scope** | Pointer to **[]string** | List of tags used to filter the map. | [optional] 
**Style** | Pointer to [**HostMapWidgetDefinitionStyle**](HostMapWidgetDefinition_style.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "hostmap"]

## Methods

### NewHostMapWidgetDefinition

`func NewHostMapWidgetDefinition(requests HostMapWidgetDefinitionRequests, type_ string, ) *HostMapWidgetDefinition`

NewHostMapWidgetDefinition instantiates a new HostMapWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMapWidgetDefinitionWithDefaults

`func NewHostMapWidgetDefinitionWithDefaults() *HostMapWidgetDefinition`

NewHostMapWidgetDefinitionWithDefaults instantiates a new HostMapWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *HostMapWidgetDefinition) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *HostMapWidgetDefinition) GetGroupOk() ([]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroup

`func (o *HostMapWidgetDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroup

`func (o *HostMapWidgetDefinition) SetGroup(v []string)`

SetGroup gets a reference to the given []string and assigns it to the Group field.

### GetNoGroupHosts

`func (o *HostMapWidgetDefinition) GetNoGroupHosts() bool`

GetNoGroupHosts returns the NoGroupHosts field if non-nil, zero value otherwise.

### GetNoGroupHostsOk

`func (o *HostMapWidgetDefinition) GetNoGroupHostsOk() (bool, bool)`

GetNoGroupHostsOk returns a tuple with the NoGroupHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNoGroupHosts

`func (o *HostMapWidgetDefinition) HasNoGroupHosts() bool`

HasNoGroupHosts returns a boolean if a field has been set.

### SetNoGroupHosts

`func (o *HostMapWidgetDefinition) SetNoGroupHosts(v bool)`

SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.

### GetNoMetricHosts

`func (o *HostMapWidgetDefinition) GetNoMetricHosts() bool`

GetNoMetricHosts returns the NoMetricHosts field if non-nil, zero value otherwise.

### GetNoMetricHostsOk

`func (o *HostMapWidgetDefinition) GetNoMetricHostsOk() (bool, bool)`

GetNoMetricHostsOk returns a tuple with the NoMetricHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNoMetricHosts

`func (o *HostMapWidgetDefinition) HasNoMetricHosts() bool`

HasNoMetricHosts returns a boolean if a field has been set.

### SetNoMetricHosts

`func (o *HostMapWidgetDefinition) SetNoMetricHosts(v bool)`

SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.

### GetNodeType

`func (o *HostMapWidgetDefinition) GetNodeType() WidgetNodeType`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *HostMapWidgetDefinition) GetNodeTypeOk() (WidgetNodeType, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNodeType

`func (o *HostMapWidgetDefinition) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### SetNodeType

`func (o *HostMapWidgetDefinition) SetNodeType(v WidgetNodeType)`

SetNodeType gets a reference to the given WidgetNodeType and assigns it to the NodeType field.

### GetNotes

`func (o *HostMapWidgetDefinition) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *HostMapWidgetDefinition) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *HostMapWidgetDefinition) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *HostMapWidgetDefinition) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### GetRequests

`func (o *HostMapWidgetDefinition) GetRequests() HostMapWidgetDefinitionRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HostMapWidgetDefinition) GetRequestsOk() (HostMapWidgetDefinitionRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequests

`func (o *HostMapWidgetDefinition) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequests

`func (o *HostMapWidgetDefinition) SetRequests(v HostMapWidgetDefinitionRequests)`

SetRequests gets a reference to the given HostMapWidgetDefinitionRequests and assigns it to the Requests field.

### GetScope

`func (o *HostMapWidgetDefinition) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HostMapWidgetDefinition) GetScopeOk() ([]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *HostMapWidgetDefinition) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *HostMapWidgetDefinition) SetScope(v []string)`

SetScope gets a reference to the given []string and assigns it to the Scope field.

### GetStyle

`func (o *HostMapWidgetDefinition) GetStyle() HostMapWidgetDefinitionStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *HostMapWidgetDefinition) GetStyleOk() (HostMapWidgetDefinitionStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStyle

`func (o *HostMapWidgetDefinition) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyle

`func (o *HostMapWidgetDefinition) SetStyle(v HostMapWidgetDefinitionStyle)`

SetStyle gets a reference to the given HostMapWidgetDefinitionStyle and assigns it to the Style field.

### GetTitle

`func (o *HostMapWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HostMapWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *HostMapWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *HostMapWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *HostMapWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *HostMapWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *HostMapWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *HostMapWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *HostMapWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *HostMapWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *HostMapWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *HostMapWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *HostMapWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostMapWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *HostMapWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *HostMapWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *HostMapWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of HostMapWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


