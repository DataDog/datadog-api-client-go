// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppWidgetDefinition The embedded app widget displays an App Builder app on a dashboard. Exactly one of `app_id` or `template_id` must be provided; they cannot be provided together.
type EmbeddedAppWidgetDefinition struct {
	// UUID of the App Builder app to embed.
	AppId *string `json:"app_id,omitempty"`
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// The description of the widget.
	Description *string `json:"description,omitempty"`
	// Inputs passed to the embedded app.
	Inputs []EmbeddedAppWidgetInput `json:"inputs,omitempty"`
	// ID of the built-in app template to embed.
	TemplateId *string `json:"template_id,omitempty"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the embedded app widget.
	Type EmbeddedAppWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEmbeddedAppWidgetDefinition instantiates a new EmbeddedAppWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEmbeddedAppWidgetDefinition(typeVar EmbeddedAppWidgetDefinitionType) *EmbeddedAppWidgetDefinition {
	this := EmbeddedAppWidgetDefinition{}
	this.Type = typeVar
	return &this
}

// NewEmbeddedAppWidgetDefinitionWithDefaults instantiates a new EmbeddedAppWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEmbeddedAppWidgetDefinitionWithDefaults() *EmbeddedAppWidgetDefinition {
	this := EmbeddedAppWidgetDefinition{}
	var typeVar EmbeddedAppWidgetDefinitionType = EMBEDDEDAPPWIDGETDEFINITIONTYPE_EMBEDDED_APP
	this.Type = typeVar
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasAppId() bool {
	return o != nil && o.AppId != nil
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *EmbeddedAppWidgetDefinition) SetAppId(v string) {
	o.AppId = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *EmbeddedAppWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EmbeddedAppWidgetDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetInputs() []EmbeddedAppWidgetInput {
	if o == nil || o.Inputs == nil {
		var ret []EmbeddedAppWidgetInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetInputsOk() (*[]EmbeddedAppWidgetInput, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasInputs() bool {
	return o != nil && o.Inputs != nil
}

// SetInputs gets a reference to the given []EmbeddedAppWidgetInput and assigns it to the Inputs field.
func (o *EmbeddedAppWidgetDefinition) SetInputs(v []EmbeddedAppWidgetInput) {
	o.Inputs = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasTemplateId() bool {
	return o != nil && o.TemplateId != nil
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *EmbeddedAppWidgetDefinition) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *EmbeddedAppWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EmbeddedAppWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *EmbeddedAppWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *EmbeddedAppWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *EmbeddedAppWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *EmbeddedAppWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *EmbeddedAppWidgetDefinition) GetType() EmbeddedAppWidgetDefinitionType {
	if o == nil {
		var ret EmbeddedAppWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetDefinition) GetTypeOk() (*EmbeddedAppWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EmbeddedAppWidgetDefinition) SetType(v EmbeddedAppWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EmbeddedAppWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.TemplateId != nil {
		toSerialize["template_id"] = o.TemplateId
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
func (o *EmbeddedAppWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppId       *string                          `json:"app_id,omitempty"`
		CustomLinks []WidgetCustomLink               `json:"custom_links,omitempty"`
		Description *string                          `json:"description,omitempty"`
		Inputs      []EmbeddedAppWidgetInput         `json:"inputs,omitempty"`
		TemplateId  *string                          `json:"template_id,omitempty"`
		Time        *WidgetTime                      `json:"time,omitempty"`
		Title       *string                          `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize   *string                          `json:"title_size,omitempty"`
		Type        *EmbeddedAppWidgetDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.AppId = all.AppId
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	o.Inputs = all.Inputs
	o.TemplateId = all.TemplateId
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
