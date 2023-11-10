// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetDefinition The top list visualization enables you to display a list of Tag value like hostname or service with the most or least of any metric value, such as highest consumers of CPU, hosts with the least disk space, etc.
type ToplistWidgetDefinition struct {
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// List of top list widget requests.
	Requests []ToplistWidgetRequest `json:"requests"`
	// Style customization for a top list widget.
	Style *ToplistWidgetStyle `json:"style,omitempty"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the top list widget.
	Type ToplistWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewToplistWidgetDefinition instantiates a new ToplistWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistWidgetDefinition(requests []ToplistWidgetRequest, typeVar ToplistWidgetDefinitionType) *ToplistWidgetDefinition {
	this := ToplistWidgetDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewToplistWidgetDefinitionWithDefaults instantiates a new ToplistWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistWidgetDefinitionWithDefaults() *ToplistWidgetDefinition {
	this := ToplistWidgetDefinition{}
	var typeVar ToplistWidgetDefinitionType = TOPLISTWIDGETDEFINITIONTYPE_TOPLIST
	this.Type = typeVar
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *ToplistWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetRequests returns the Requests field value.
func (o *ToplistWidgetDefinition) GetRequests() []ToplistWidgetRequest {
	if o == nil {
		var ret []ToplistWidgetRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetRequestsOk() (*[]ToplistWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *ToplistWidgetDefinition) SetRequests(v []ToplistWidgetRequest) {
	o.Requests = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetStyle() ToplistWidgetStyle {
	if o == nil || o.Style == nil {
		var ret ToplistWidgetStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetStyleOk() (*ToplistWidgetStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given ToplistWidgetStyle and assigns it to the Style field.
func (o *ToplistWidgetDefinition) SetStyle(v ToplistWidgetStyle) {
	o.Style = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *ToplistWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ToplistWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ToplistWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ToplistWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ToplistWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ToplistWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *ToplistWidgetDefinition) GetType() ToplistWidgetDefinitionType {
	if o == nil {
		var ret ToplistWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ToplistWidgetDefinition) GetTypeOk() (*ToplistWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ToplistWidgetDefinition) SetType(v ToplistWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	toSerialize["requests"] = o.Requests
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks []WidgetCustomLink           `json:"custom_links,omitempty"`
		Requests    *[]ToplistWidgetRequest      `json:"requests"`
		Style       *ToplistWidgetStyle          `json:"style,omitempty"`
		Time        *WidgetTime                  `json:"time,omitempty"`
		Title       *string                      `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign             `json:"title_align,omitempty"`
		TitleSize   *string                      `json:"title_size,omitempty"`
		Type        *ToplistWidgetDefinitionType `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_links", "requests", "style", "time", "title", "title_align", "title_size", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Requests = *all.Requests
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
