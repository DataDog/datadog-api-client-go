// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type NoteBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BackgroundColor *string `json:"background_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Content string `json:"content"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FontSize *string `json:"font_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasPadding *bool `json:"has_padding,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Notes *string `json:"notes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowTick *bool `json:"show_tick,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextAlign *NoteBackendDefinitionTextAlign `json:"text_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TickEdge *NoteBackendDefinitionTickEdge `json:"tick_edge,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TickPos *string `json:"tick_pos,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TitleSize *string `json:"title_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type NoteBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	VerticalAlign *NoteBackendDefinitionVerticalAlign `json:"vertical_align,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNoteBackendDefinition instantiates a new NoteBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNoteBackendDefinition(content string, typeVar NoteBackendDefinitionType) *NoteBackendDefinition {
	this := NoteBackendDefinition{}
	this.Content = content
	this.Type = typeVar
	return &this
}

// NewNoteBackendDefinitionWithDefaults instantiates a new NoteBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNoteBackendDefinitionWithDefaults() *NoteBackendDefinition {
	this := NoteBackendDefinition{}
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasBackgroundColor() bool {
	return o != nil && o.BackgroundColor != nil
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *NoteBackendDefinition) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetContent returns the Content field value.
func (o *NoteBackendDefinition) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *NoteBackendDefinition) SetContent(v string) {
	o.Content = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NoteBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetFontSize returns the FontSize field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetFontSize() string {
	if o == nil || o.FontSize == nil {
		var ret string
		return ret
	}
	return *o.FontSize
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetFontSizeOk() (*string, bool) {
	if o == nil || o.FontSize == nil {
		return nil, false
	}
	return o.FontSize, true
}

// HasFontSize returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasFontSize() bool {
	return o != nil && o.FontSize != nil
}

// SetFontSize gets a reference to the given string and assigns it to the FontSize field.
func (o *NoteBackendDefinition) SetFontSize(v string) {
	o.FontSize = &v
}

// GetHasPadding returns the HasPadding field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetHasPadding() bool {
	if o == nil || o.HasPadding == nil {
		var ret bool
		return ret
	}
	return *o.HasPadding
}

// GetHasPaddingOk returns a tuple with the HasPadding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetHasPaddingOk() (*bool, bool) {
	if o == nil || o.HasPadding == nil {
		return nil, false
	}
	return o.HasPadding, true
}

// HasHasPadding returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasHasPadding() bool {
	return o != nil && o.HasPadding != nil
}

// SetHasPadding gets a reference to the given bool and assigns it to the HasPadding field.
func (o *NoteBackendDefinition) SetHasPadding(v bool) {
	o.HasPadding = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *NoteBackendDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetShowTick returns the ShowTick field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetShowTick() bool {
	if o == nil || o.ShowTick == nil {
		var ret bool
		return ret
	}
	return *o.ShowTick
}

// GetShowTickOk returns a tuple with the ShowTick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetShowTickOk() (*bool, bool) {
	if o == nil || o.ShowTick == nil {
		return nil, false
	}
	return o.ShowTick, true
}

// HasShowTick returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasShowTick() bool {
	return o != nil && o.ShowTick != nil
}

// SetShowTick gets a reference to the given bool and assigns it to the ShowTick field.
func (o *NoteBackendDefinition) SetShowTick(v bool) {
	o.ShowTick = &v
}

// GetTextAlign returns the TextAlign field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTextAlign() NoteBackendDefinitionTextAlign {
	if o == nil || o.TextAlign == nil {
		var ret NoteBackendDefinitionTextAlign
		return ret
	}
	return *o.TextAlign
}

// GetTextAlignOk returns a tuple with the TextAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTextAlignOk() (*NoteBackendDefinitionTextAlign, bool) {
	if o == nil || o.TextAlign == nil {
		return nil, false
	}
	return o.TextAlign, true
}

// HasTextAlign returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTextAlign() bool {
	return o != nil && o.TextAlign != nil
}

// SetTextAlign gets a reference to the given NoteBackendDefinitionTextAlign and assigns it to the TextAlign field.
func (o *NoteBackendDefinition) SetTextAlign(v NoteBackendDefinitionTextAlign) {
	o.TextAlign = &v
}

// GetTickEdge returns the TickEdge field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTickEdge() NoteBackendDefinitionTickEdge {
	if o == nil || o.TickEdge == nil {
		var ret NoteBackendDefinitionTickEdge
		return ret
	}
	return *o.TickEdge
}

// GetTickEdgeOk returns a tuple with the TickEdge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTickEdgeOk() (*NoteBackendDefinitionTickEdge, bool) {
	if o == nil || o.TickEdge == nil {
		return nil, false
	}
	return o.TickEdge, true
}

// HasTickEdge returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTickEdge() bool {
	return o != nil && o.TickEdge != nil
}

// SetTickEdge gets a reference to the given NoteBackendDefinitionTickEdge and assigns it to the TickEdge field.
func (o *NoteBackendDefinition) SetTickEdge(v NoteBackendDefinitionTickEdge) {
	o.TickEdge = &v
}

// GetTickPos returns the TickPos field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTickPos() string {
	if o == nil || o.TickPos == nil {
		var ret string
		return ret
	}
	return *o.TickPos
}

// GetTickPosOk returns a tuple with the TickPos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTickPosOk() (*string, bool) {
	if o == nil || o.TickPos == nil {
		return nil, false
	}
	return o.TickPos, true
}

// HasTickPos returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTickPos() bool {
	return o != nil && o.TickPos != nil
}

// SetTickPos gets a reference to the given string and assigns it to the TickPos field.
func (o *NoteBackendDefinition) SetTickPos(v string) {
	o.TickPos = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NoteBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *NoteBackendDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *NoteBackendDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *NoteBackendDefinition) GetType() NoteBackendDefinitionType {
	if o == nil {
		var ret NoteBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetTypeOk() (*NoteBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NoteBackendDefinition) SetType(v NoteBackendDefinitionType) {
	o.Type = v
}

// GetVerticalAlign returns the VerticalAlign field value if set, zero value otherwise.
func (o *NoteBackendDefinition) GetVerticalAlign() NoteBackendDefinitionVerticalAlign {
	if o == nil || o.VerticalAlign == nil {
		var ret NoteBackendDefinitionVerticalAlign
		return ret
	}
	return *o.VerticalAlign
}

// GetVerticalAlignOk returns a tuple with the VerticalAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteBackendDefinition) GetVerticalAlignOk() (*NoteBackendDefinitionVerticalAlign, bool) {
	if o == nil || o.VerticalAlign == nil {
		return nil, false
	}
	return o.VerticalAlign, true
}

// HasVerticalAlign returns a boolean if a field has been set.
func (o *NoteBackendDefinition) HasVerticalAlign() bool {
	return o != nil && o.VerticalAlign != nil
}

// SetVerticalAlign gets a reference to the given NoteBackendDefinitionVerticalAlign and assigns it to the VerticalAlign field.
func (o *NoteBackendDefinition) SetVerticalAlign(v NoteBackendDefinitionVerticalAlign) {
	o.VerticalAlign = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NoteBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BackgroundColor != nil {
		toSerialize["background_color"] = o.BackgroundColor
	}
	toSerialize["content"] = o.Content
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FontSize != nil {
		toSerialize["font_size"] = o.FontSize
	}
	if o.HasPadding != nil {
		toSerialize["has_padding"] = o.HasPadding
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.ShowTick != nil {
		toSerialize["show_tick"] = o.ShowTick
	}
	if o.TextAlign != nil {
		toSerialize["text_align"] = o.TextAlign
	}
	if o.TickEdge != nil {
		toSerialize["tick_edge"] = o.TickEdge
	}
	if o.TickPos != nil {
		toSerialize["tick_pos"] = o.TickPos
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
	if o.VerticalAlign != nil {
		toSerialize["vertical_align"] = o.VerticalAlign
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NoteBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackgroundColor *string                             `json:"background_color,omitempty"`
		Content         *string                             `json:"content"`
		Description     *string                             `json:"description,omitempty"`
		FontSize        *string                             `json:"font_size,omitempty"`
		HasPadding      *bool                               `json:"has_padding,omitempty"`
		Notes           *string                             `json:"notes,omitempty"`
		ShowTick        *bool                               `json:"show_tick,omitempty"`
		TextAlign       *NoteBackendDefinitionTextAlign     `json:"text_align,omitempty"`
		TickEdge        *NoteBackendDefinitionTickEdge      `json:"tick_edge,omitempty"`
		TickPos         *string                             `json:"tick_pos,omitempty"`
		Title           *string                             `json:"title,omitempty"`
		TitleAlign      *WidgetTextAlign                    `json:"title_align,omitempty"`
		TitleSize       *string                             `json:"title_size,omitempty"`
		Type            *NoteBackendDefinitionType          `json:"type"`
		VerticalAlign   *NoteBackendDefinitionVerticalAlign `json:"vertical_align,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.BackgroundColor = all.BackgroundColor
	o.Content = *all.Content
	o.Description = all.Description
	o.FontSize = all.FontSize
	o.HasPadding = all.HasPadding
	o.Notes = all.Notes
	o.ShowTick = all.ShowTick
	if all.TextAlign != nil && !all.TextAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TextAlign = all.TextAlign
	}
	if all.TickEdge != nil && !all.TickEdge.IsValid() {
		hasInvalidField = true
	} else {
		o.TickEdge = all.TickEdge
	}
	o.TickPos = all.TickPos
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
	if all.VerticalAlign != nil && !all.VerticalAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.VerticalAlign = all.VerticalAlign
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
