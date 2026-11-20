// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type GroupBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BackgroundColor *string `json:"background_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	BannerImg *string `json:"banner_img,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DefaultState *GroupBackendDefinitionDefaultState `json:"default_state,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Description *string `json:"description,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LayoutType LayoutType `json:"layout_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowTitle *bool `json:"show_title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type GroupBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Widgets []DashboardBackendWidget `json:"widgets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGroupBackendDefinition instantiates a new GroupBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGroupBackendDefinition(layoutType LayoutType, typeVar GroupBackendDefinitionType, widgets []DashboardBackendWidget) *GroupBackendDefinition {
	this := GroupBackendDefinition{}
	this.LayoutType = layoutType
	this.Type = typeVar
	this.Widgets = widgets
	return &this
}

// NewGroupBackendDefinitionWithDefaults instantiates a new GroupBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGroupBackendDefinitionWithDefaults() *GroupBackendDefinition {
	this := GroupBackendDefinition{}
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasBackgroundColor() bool {
	return o != nil && o.BackgroundColor != nil
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *GroupBackendDefinition) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetBannerImg returns the BannerImg field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetBannerImg() string {
	if o == nil || o.BannerImg == nil {
		var ret string
		return ret
	}
	return *o.BannerImg
}

// GetBannerImgOk returns a tuple with the BannerImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetBannerImgOk() (*string, bool) {
	if o == nil || o.BannerImg == nil {
		return nil, false
	}
	return o.BannerImg, true
}

// HasBannerImg returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasBannerImg() bool {
	return o != nil && o.BannerImg != nil
}

// SetBannerImg gets a reference to the given string and assigns it to the BannerImg field.
func (o *GroupBackendDefinition) SetBannerImg(v string) {
	o.BannerImg = &v
}

// GetDefaultState returns the DefaultState field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetDefaultState() GroupBackendDefinitionDefaultState {
	if o == nil || o.DefaultState == nil {
		var ret GroupBackendDefinitionDefaultState
		return ret
	}
	return *o.DefaultState
}

// GetDefaultStateOk returns a tuple with the DefaultState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetDefaultStateOk() (*GroupBackendDefinitionDefaultState, bool) {
	if o == nil || o.DefaultState == nil {
		return nil, false
	}
	return o.DefaultState, true
}

// HasDefaultState returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasDefaultState() bool {
	return o != nil && o.DefaultState != nil
}

// SetDefaultState gets a reference to the given GroupBackendDefinitionDefaultState and assigns it to the DefaultState field.
func (o *GroupBackendDefinition) SetDefaultState(v GroupBackendDefinitionDefaultState) {
	o.DefaultState = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupBackendDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLayoutType returns the LayoutType field value.
func (o *GroupBackendDefinition) GetLayoutType() LayoutType {
	if o == nil {
		var ret LayoutType
		return ret
	}
	return o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetLayoutTypeOk() (*LayoutType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayoutType, true
}

// SetLayoutType sets field value.
func (o *GroupBackendDefinition) SetLayoutType(v LayoutType) {
	o.LayoutType = v
}

// GetShowTitle returns the ShowTitle field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetShowTitle() bool {
	if o == nil || o.ShowTitle == nil {
		var ret bool
		return ret
	}
	return *o.ShowTitle
}

// GetShowTitleOk returns a tuple with the ShowTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetShowTitleOk() (*bool, bool) {
	if o == nil || o.ShowTitle == nil {
		return nil, false
	}
	return o.ShowTitle, true
}

// HasShowTitle returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasShowTitle() bool {
	return o != nil && o.ShowTitle != nil
}

// SetShowTitle gets a reference to the given bool and assigns it to the ShowTitle field.
func (o *GroupBackendDefinition) SetShowTitle(v bool) {
	o.ShowTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *GroupBackendDefinition) GetType() GroupBackendDefinitionType {
	if o == nil {
		var ret GroupBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetTypeOk() (*GroupBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GroupBackendDefinition) SetType(v GroupBackendDefinitionType) {
	o.Type = v
}

// GetWidgets returns the Widgets field value.
func (o *GroupBackendDefinition) GetWidgets() []DashboardBackendWidget {
	if o == nil {
		var ret []DashboardBackendWidget
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value
// and a boolean to check if the value has been set.
func (o *GroupBackendDefinition) GetWidgetsOk() (*[]DashboardBackendWidget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Widgets, true
}

// SetWidgets sets field value.
func (o *GroupBackendDefinition) SetWidgets(v []DashboardBackendWidget) {
	o.Widgets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GroupBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BackgroundColor != nil {
		toSerialize["background_color"] = o.BackgroundColor
	}
	if o.BannerImg != nil {
		toSerialize["banner_img"] = o.BannerImg
	}
	if o.DefaultState != nil {
		toSerialize["default_state"] = o.DefaultState
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["layout_type"] = o.LayoutType
	if o.ShowTitle != nil {
		toSerialize["show_title"] = o.ShowTitle
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type
	toSerialize["widgets"] = o.Widgets
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GroupBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackgroundColor *string                             `json:"background_color,omitempty"`
		BannerImg       *string                             `json:"banner_img,omitempty"`
		DefaultState    *GroupBackendDefinitionDefaultState `json:"default_state,omitempty"`
		Description     *string                             `json:"description,omitempty"`
		LayoutType      *LayoutType                         `json:"layout_type"`
		ShowTitle       *bool                               `json:"show_title,omitempty"`
		Title           *string                             `json:"title,omitempty"`
		Type            *GroupBackendDefinitionType         `json:"type"`
		Widgets         *[]DashboardBackendWidget           `json:"widgets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LayoutType == nil {
		return fmt.Errorf("required field layout_type missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Widgets == nil {
		return fmt.Errorf("required field widgets missing")
	}

	hasInvalidField := false
	o.BackgroundColor = all.BackgroundColor
	o.BannerImg = all.BannerImg
	if all.DefaultState != nil && !all.DefaultState.IsValid() {
		hasInvalidField = true
	} else {
		o.DefaultState = all.DefaultState
	}
	o.Description = all.Description
	if !all.LayoutType.IsValid() {
		hasInvalidField = true
	} else {
		o.LayoutType = *all.LayoutType
	}
	o.ShowTitle = all.ShowTitle
	o.Title = all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Widgets = *all.Widgets

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
