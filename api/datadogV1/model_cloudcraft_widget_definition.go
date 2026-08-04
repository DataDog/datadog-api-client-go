// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudcraftWidgetDefinition This widget displays a Cloudcraft topology of cloud resources for the selected provider.
type CloudcraftWidgetDefinition struct {
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// The description of the widget.
	Description *string `json:"description,omitempty"`
	// List of tags or attributes used to group the cloud resources in the widget.
	GroupBy []string `json:"group_by"`
	// Search query that visually highlights matching resources in the diagram.
	Highlighted string `json:"highlighted"`
	// Overlay applied on top of the Cloudcraft topology.
	Overlay CloudcraftWidgetDefinitionOverlay `json:"overlay"`
	// Filter applied to the selected overlay.
	OverlayFilter string `json:"overlay_filter"`
	// Projection used to render the Cloudcraft topology.
	Projection CloudcraftWidgetDefinitionProjection `json:"projection"`
	// Cloud provider for the Cloudcraft widget.
	Provider CloudcraftWidgetDefinitionProvider `json:"provider"`
	// Query string used to filter the cloud resources displayed in the widget.
	QueryString string `json:"query_string"`
	// Whether to show empty outline groups in the diagram.
	ShowEmptyGroups bool `json:"show_empty_groups"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of Cloudcraft widget.
	Type CloudcraftWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudcraftWidgetDefinition instantiates a new CloudcraftWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudcraftWidgetDefinition(groupBy []string, highlighted string, overlay CloudcraftWidgetDefinitionOverlay, overlayFilter string, projection CloudcraftWidgetDefinitionProjection, provider CloudcraftWidgetDefinitionProvider, queryString string, showEmptyGroups bool, typeVar CloudcraftWidgetDefinitionType) *CloudcraftWidgetDefinition {
	this := CloudcraftWidgetDefinition{}
	this.GroupBy = groupBy
	this.Highlighted = highlighted
	this.Overlay = overlay
	this.OverlayFilter = overlayFilter
	this.Projection = projection
	this.Provider = provider
	this.QueryString = queryString
	this.ShowEmptyGroups = showEmptyGroups
	this.Type = typeVar
	return &this
}

// NewCloudcraftWidgetDefinitionWithDefaults instantiates a new CloudcraftWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudcraftWidgetDefinitionWithDefaults() *CloudcraftWidgetDefinition {
	this := CloudcraftWidgetDefinition{}
	var typeVar CloudcraftWidgetDefinitionType = CLOUDCRAFTWIDGETDEFINITIONTYPE_CLOUDCRAFT
	this.Type = typeVar
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *CloudcraftWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *CloudcraftWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *CloudcraftWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudcraftWidgetDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudcraftWidgetDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudcraftWidgetDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetGroupBy returns the GroupBy field value.
func (o *CloudcraftWidgetDefinition) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetGroupByOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *CloudcraftWidgetDefinition) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetHighlighted returns the Highlighted field value.
func (o *CloudcraftWidgetDefinition) GetHighlighted() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Highlighted
}

// GetHighlightedOk returns a tuple with the Highlighted field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetHighlightedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Highlighted, true
}

// SetHighlighted sets field value.
func (o *CloudcraftWidgetDefinition) SetHighlighted(v string) {
	o.Highlighted = v
}

// GetOverlay returns the Overlay field value.
func (o *CloudcraftWidgetDefinition) GetOverlay() CloudcraftWidgetDefinitionOverlay {
	if o == nil {
		var ret CloudcraftWidgetDefinitionOverlay
		return ret
	}
	return o.Overlay
}

// GetOverlayOk returns a tuple with the Overlay field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetOverlayOk() (*CloudcraftWidgetDefinitionOverlay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overlay, true
}

// SetOverlay sets field value.
func (o *CloudcraftWidgetDefinition) SetOverlay(v CloudcraftWidgetDefinitionOverlay) {
	o.Overlay = v
}

// GetOverlayFilter returns the OverlayFilter field value.
func (o *CloudcraftWidgetDefinition) GetOverlayFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OverlayFilter
}

// GetOverlayFilterOk returns a tuple with the OverlayFilter field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetOverlayFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverlayFilter, true
}

// SetOverlayFilter sets field value.
func (o *CloudcraftWidgetDefinition) SetOverlayFilter(v string) {
	o.OverlayFilter = v
}

// GetProjection returns the Projection field value.
func (o *CloudcraftWidgetDefinition) GetProjection() CloudcraftWidgetDefinitionProjection {
	if o == nil {
		var ret CloudcraftWidgetDefinitionProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetProjectionOk() (*CloudcraftWidgetDefinitionProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *CloudcraftWidgetDefinition) SetProjection(v CloudcraftWidgetDefinitionProjection) {
	o.Projection = v
}

// GetProvider returns the Provider field value.
func (o *CloudcraftWidgetDefinition) GetProvider() CloudcraftWidgetDefinitionProvider {
	if o == nil {
		var ret CloudcraftWidgetDefinitionProvider
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetProviderOk() (*CloudcraftWidgetDefinitionProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *CloudcraftWidgetDefinition) SetProvider(v CloudcraftWidgetDefinitionProvider) {
	o.Provider = v
}

// GetQueryString returns the QueryString field value.
func (o *CloudcraftWidgetDefinition) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *CloudcraftWidgetDefinition) SetQueryString(v string) {
	o.QueryString = v
}

// GetShowEmptyGroups returns the ShowEmptyGroups field value.
func (o *CloudcraftWidgetDefinition) GetShowEmptyGroups() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ShowEmptyGroups
}

// GetShowEmptyGroupsOk returns a tuple with the ShowEmptyGroups field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetShowEmptyGroupsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowEmptyGroups, true
}

// SetShowEmptyGroups sets field value.
func (o *CloudcraftWidgetDefinition) SetShowEmptyGroups(v bool) {
	o.ShowEmptyGroups = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CloudcraftWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CloudcraftWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CloudcraftWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *CloudcraftWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *CloudcraftWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *CloudcraftWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *CloudcraftWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *CloudcraftWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *CloudcraftWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *CloudcraftWidgetDefinition) GetType() CloudcraftWidgetDefinitionType {
	if o == nil {
		var ret CloudcraftWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudcraftWidgetDefinition) GetTypeOk() (*CloudcraftWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CloudcraftWidgetDefinition) SetType(v CloudcraftWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudcraftWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["group_by"] = o.GroupBy
	toSerialize["highlighted"] = o.Highlighted
	toSerialize["overlay"] = o.Overlay
	toSerialize["overlay_filter"] = o.OverlayFilter
	toSerialize["projection"] = o.Projection
	toSerialize["provider"] = o.Provider
	toSerialize["query_string"] = o.QueryString
	toSerialize["show_empty_groups"] = o.ShowEmptyGroups
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudcraftWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks     []WidgetCustomLink                    `json:"custom_links,omitempty"`
		Description     *string                               `json:"description,omitempty"`
		GroupBy         *[]string                             `json:"group_by"`
		Highlighted     *string                               `json:"highlighted"`
		Overlay         *CloudcraftWidgetDefinitionOverlay    `json:"overlay"`
		OverlayFilter   *string                               `json:"overlay_filter"`
		Projection      *CloudcraftWidgetDefinitionProjection `json:"projection"`
		Provider        *CloudcraftWidgetDefinitionProvider   `json:"provider"`
		QueryString     *string                               `json:"query_string"`
		ShowEmptyGroups *bool                                 `json:"show_empty_groups"`
		Title           *string                               `json:"title,omitempty"`
		TitleAlign      *WidgetTextAlign                      `json:"title_align,omitempty"`
		TitleSize       *string                               `json:"title_size,omitempty"`
		Type            *CloudcraftWidgetDefinitionType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.Highlighted == nil {
		return fmt.Errorf("required field highlighted missing")
	}
	if all.Overlay == nil {
		return fmt.Errorf("required field overlay missing")
	}
	if all.OverlayFilter == nil {
		return fmt.Errorf("required field overlay_filter missing")
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}
	if all.ShowEmptyGroups == nil {
		return fmt.Errorf("required field show_empty_groups missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_links", "description", "group_by", "highlighted", "overlay", "overlay_filter", "projection", "provider", "query_string", "show_empty_groups", "title", "title_align", "title_size", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.GroupBy = *all.GroupBy
	o.Highlighted = *all.Highlighted
	if !all.Overlay.IsValid() {
		hasInvalidField = true
	} else {
		o.Overlay = *all.Overlay
	}
	o.OverlayFilter = *all.OverlayFilter
	if !all.Projection.IsValid() {
		hasInvalidField = true
	} else {
		o.Projection = *all.Projection
	}
	if !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = *all.Provider
	}
	o.QueryString = *all.QueryString
	o.ShowEmptyGroups = *all.ShowEmptyGroups
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
