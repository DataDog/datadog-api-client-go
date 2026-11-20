// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type HostmapBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Group []string `json:"group,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoGroupHosts *bool `json:"no_group_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NoMetricHosts *bool `json:"no_metric_hosts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NodeType *HostmapBackendDefinitionNodeType `json:"node_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Requests HostmapBackendDefinition20270101Requests `json:"requests"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scope []string `json:"scope,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *HostmapBackendDefinitionStyle `json:"style,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *WidgetTime `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type HostmapBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapBackendDefinition instantiates a new HostmapBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapBackendDefinition(requests HostmapBackendDefinition20270101Requests, typeVar HostmapBackendDefinitionType) *HostmapBackendDefinition {
	this := HostmapBackendDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewHostmapBackendDefinitionWithDefaults instantiates a new HostmapBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapBackendDefinitionWithDefaults() *HostmapBackendDefinition {
	this := HostmapBackendDefinition{}
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *HostmapBackendDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HostmapBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetGroup() []string {
	if o == nil || o.Group == nil {
		var ret []string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetGroupOk() (*[]string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return &o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given []string and assigns it to the Group field.
func (o *HostmapBackendDefinition) SetGroup(v []string) {
	o.Group = v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostmapBackendDefinition) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostmapBackendDefinition) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetNodeType() HostmapBackendDefinitionNodeType {
	if o == nil || o.NodeType == nil {
		var ret HostmapBackendDefinitionNodeType
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetNodeTypeOk() (*HostmapBackendDefinitionNodeType, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given HostmapBackendDefinitionNodeType and assigns it to the NodeType field.
func (o *HostmapBackendDefinition) SetNodeType(v HostmapBackendDefinitionNodeType) {
	o.NodeType = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *HostmapBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *HostmapBackendDefinition) GetRequests() HostmapBackendDefinition20270101Requests {
	if o == nil {
		var ret HostmapBackendDefinition20270101Requests
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetRequestsOk() (*HostmapBackendDefinition20270101Requests, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *HostmapBackendDefinition) SetRequests(v HostmapBackendDefinition20270101Requests) {
	o.Requests = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetScopeOk() (*[]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *HostmapBackendDefinition) SetScope(v []string) {
	o.Scope = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetStyle() HostmapBackendDefinitionStyle {
	if o == nil || o.Style == nil {
		var ret HostmapBackendDefinitionStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetStyleOk() (*HostmapBackendDefinitionStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostmapBackendDefinitionStyle and assigns it to the Style field.
func (o *HostmapBackendDefinition) SetStyle(v HostmapBackendDefinitionStyle) {
	o.Style = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *HostmapBackendDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HostmapBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *HostmapBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *HostmapBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *HostmapBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *HostmapBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *HostmapBackendDefinition) GetType() HostmapBackendDefinitionType {
	if o == nil {
		var ret HostmapBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostmapBackendDefinition) GetTypeOk() (*HostmapBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HostmapBackendDefinition) SetType(v HostmapBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapBackendDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.NoGroupHosts != nil {
		toSerialize["no_group_hosts"] = o.NoGroupHosts
	}
	if o.NoMetricHosts != nil {
		toSerialize["no_metric_hosts"] = o.NoMetricHosts
	}
	if o.NodeType != nil {
		toSerialize["node_type"] = o.NodeType
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostmapBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks   []WidgetCustomLink                        `json:"custom_links,omitempty"`
		Description   *string                                   `json:"description,omitempty"`
		Group         []string                                  `json:"group,omitempty"`
		NoGroupHosts  *bool                                     `json:"no_group_hosts,omitempty"`
		NoMetricHosts *bool                                     `json:"no_metric_hosts,omitempty"`
		NodeType      *HostmapBackendDefinitionNodeType         `json:"node_type,omitempty"`
		Notes         *string                                   `json:"notes,omitempty"`
		Requests      *HostmapBackendDefinition20270101Requests `json:"requests"`
		Scope         []string                                  `json:"scope,omitempty"`
		Style         *HostmapBackendDefinitionStyle            `json:"style,omitempty"`
		Time          *WidgetTime                               `json:"time,omitempty"`
		Title         *string                                   `json:"title,omitempty"`
		TitleAlign    *WidgetTextAlign                          `json:"title_align,omitempty"`
		TitleSize     *string                                   `json:"title_size,omitempty"`
		Type          *HostmapBackendDefinitionType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.Group = all.Group
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if all.NodeType != nil && !all.NodeType.IsValid() {
		hasInvalidField = true
	} else {
		o.NodeType = all.NodeType
	}
	o.Notes = all.Notes
	o.Requests = *all.Requests
	o.Scope = all.Scope
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style
	o.Time = all.Time
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
