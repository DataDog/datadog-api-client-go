# TreeMapWidgetDefinition

## Properties

| Name         | Type                                                              | Description                       | Notes                                            |
| ------------ | ----------------------------------------------------------------- | --------------------------------- | ------------------------------------------------ |
| **ColorBy**  | Pointer to [**TreeMapColorBy**](TreeMapColorBy.md)                |                                   | [optional] [default to TREEMAPCOLORBY_USER]      |
| **GroupBy**  | Pointer to [**TreeMapGroupBy**](TreeMapGroupBy.md)                |                                   | [optional]                                       |
| **Requests** | [**[]TreeMapWidgetRequest**](TreeMapWidgetRequest.md)             | List of top list widget requests. |
| **SizeBy**   | Pointer to [**TreeMapSizeBy**](TreeMapSizeBy.md)                  |                                   | [optional]                                       |
| **Title**    | Pointer to **string**                                             | Title of your widget.             | [optional]                                       |
| **Type**     | [**TreeMapWidgetDefinitionType**](TreeMapWidgetDefinitionType.md) |                                   | [default to TREEMAPWIDGETDEFINITIONTYPE_TREEMAP] |

## Methods

### NewTreeMapWidgetDefinition

`func NewTreeMapWidgetDefinition(requests []TreeMapWidgetRequest, type_ TreeMapWidgetDefinitionType) *TreeMapWidgetDefinition`

NewTreeMapWidgetDefinition instantiates a new TreeMapWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewTreeMapWidgetDefinitionWithDefaults

`func NewTreeMapWidgetDefinitionWithDefaults() *TreeMapWidgetDefinition`

NewTreeMapWidgetDefinitionWithDefaults instantiates a new TreeMapWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetColorBy

`func (o *TreeMapWidgetDefinition) GetColorBy() TreeMapColorBy`

GetColorBy returns the ColorBy field if non-nil, zero value otherwise.

### GetColorByOk

`func (o *TreeMapWidgetDefinition) GetColorByOk() (*TreeMapColorBy, bool)`

GetColorByOk returns a tuple with the ColorBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorBy

`func (o *TreeMapWidgetDefinition) SetColorBy(v TreeMapColorBy)`

SetColorBy sets ColorBy field to given value.

### HasColorBy

`func (o *TreeMapWidgetDefinition) HasColorBy() bool`

HasColorBy returns a boolean if a field has been set.

### GetGroupBy

`func (o *TreeMapWidgetDefinition) GetGroupBy() TreeMapGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *TreeMapWidgetDefinition) GetGroupByOk() (*TreeMapGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *TreeMapWidgetDefinition) SetGroupBy(v TreeMapGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *TreeMapWidgetDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetRequests

`func (o *TreeMapWidgetDefinition) GetRequests() []TreeMapWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TreeMapWidgetDefinition) GetRequestsOk() (*[]TreeMapWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TreeMapWidgetDefinition) SetRequests(v []TreeMapWidgetRequest)`

SetRequests sets Requests field to given value.

### GetSizeBy

`func (o *TreeMapWidgetDefinition) GetSizeBy() TreeMapSizeBy`

GetSizeBy returns the SizeBy field if non-nil, zero value otherwise.

### GetSizeByOk

`func (o *TreeMapWidgetDefinition) GetSizeByOk() (*TreeMapSizeBy, bool)`

GetSizeByOk returns a tuple with the SizeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBy

`func (o *TreeMapWidgetDefinition) SetSizeBy(v TreeMapSizeBy)`

SetSizeBy sets SizeBy field to given value.

### HasSizeBy

`func (o *TreeMapWidgetDefinition) HasSizeBy() bool`

HasSizeBy returns a boolean if a field has been set.

### GetTitle

`func (o *TreeMapWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TreeMapWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TreeMapWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TreeMapWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *TreeMapWidgetDefinition) GetType() TreeMapWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TreeMapWidgetDefinition) GetTypeOk() (*TreeMapWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TreeMapWidgetDefinition) SetType(v TreeMapWidgetDefinitionType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
