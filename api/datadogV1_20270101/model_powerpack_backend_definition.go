// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type PowerpackBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BackgroundColor *string `json:"background_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	BannerImg *string `json:"banner_img,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PowerpackId string `json:"powerpack_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowTitle *bool `json:"show_title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	State *PowerpackState `json:"state,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TemplateVariables *PowerpackBackendTemplateVariables `json:"template_variables,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Title *string `json:"title,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type PowerpackBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPowerpackBackendDefinition instantiates a new PowerpackBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackBackendDefinition(powerpackId string, typeVar PowerpackBackendDefinitionType) *PowerpackBackendDefinition {
	this := PowerpackBackendDefinition{}
	this.PowerpackId = powerpackId
	this.Type = typeVar
	return &this
}

// NewPowerpackBackendDefinitionWithDefaults instantiates a new PowerpackBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackBackendDefinitionWithDefaults() *PowerpackBackendDefinition {
	this := PowerpackBackendDefinition{}
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasBackgroundColor() bool {
	return o != nil && o.BackgroundColor != nil
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *PowerpackBackendDefinition) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetBannerImg returns the BannerImg field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetBannerImg() string {
	if o == nil || o.BannerImg == nil {
		var ret string
		return ret
	}
	return *o.BannerImg
}

// GetBannerImgOk returns a tuple with the BannerImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetBannerImgOk() (*string, bool) {
	if o == nil || o.BannerImg == nil {
		return nil, false
	}
	return o.BannerImg, true
}

// HasBannerImg returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasBannerImg() bool {
	return o != nil && o.BannerImg != nil
}

// SetBannerImg gets a reference to the given string and assigns it to the BannerImg field.
func (o *PowerpackBackendDefinition) SetBannerImg(v string) {
	o.BannerImg = &v
}

// GetPowerpackId returns the PowerpackId field value.
func (o *PowerpackBackendDefinition) GetPowerpackId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PowerpackId
}

// GetPowerpackIdOk returns a tuple with the PowerpackId field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetPowerpackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerpackId, true
}

// SetPowerpackId sets field value.
func (o *PowerpackBackendDefinition) SetPowerpackId(v string) {
	o.PowerpackId = v
}

// GetShowTitle returns the ShowTitle field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetShowTitle() bool {
	if o == nil || o.ShowTitle == nil {
		var ret bool
		return ret
	}
	return *o.ShowTitle
}

// GetShowTitleOk returns a tuple with the ShowTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetShowTitleOk() (*bool, bool) {
	if o == nil || o.ShowTitle == nil {
		return nil, false
	}
	return o.ShowTitle, true
}

// HasShowTitle returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasShowTitle() bool {
	return o != nil && o.ShowTitle != nil
}

// SetShowTitle gets a reference to the given bool and assigns it to the ShowTitle field.
func (o *PowerpackBackendDefinition) SetShowTitle(v bool) {
	o.ShowTitle = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetState() PowerpackState {
	if o == nil || o.State == nil {
		var ret PowerpackState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetStateOk() (*PowerpackState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given PowerpackState and assigns it to the State field.
func (o *PowerpackBackendDefinition) SetState(v PowerpackState) {
	o.State = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetTemplateVariables() PowerpackBackendTemplateVariables {
	if o == nil || o.TemplateVariables == nil {
		var ret PowerpackBackendTemplateVariables
		return ret
	}
	return *o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetTemplateVariablesOk() (*PowerpackBackendTemplateVariables, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given PowerpackBackendTemplateVariables and assigns it to the TemplateVariables field.
func (o *PowerpackBackendDefinition) SetTemplateVariables(v PowerpackBackendTemplateVariables) {
	o.TemplateVariables = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PowerpackBackendDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PowerpackBackendDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PowerpackBackendDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *PowerpackBackendDefinition) GetType() PowerpackBackendDefinitionType {
	if o == nil {
		var ret PowerpackBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendDefinition) GetTypeOk() (*PowerpackBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PowerpackBackendDefinition) SetType(v PowerpackBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackBackendDefinition) MarshalJSON() ([]byte, error) {
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
	toSerialize["powerpack_id"] = o.PowerpackId
	if o.ShowTitle != nil {
		toSerialize["show_title"] = o.ShowTitle
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackgroundColor   *string                            `json:"background_color,omitempty"`
		BannerImg         *string                            `json:"banner_img,omitempty"`
		PowerpackId       *string                            `json:"powerpack_id"`
		ShowTitle         *bool                              `json:"show_title,omitempty"`
		State             *PowerpackState                    `json:"state,omitempty"`
		TemplateVariables *PowerpackBackendTemplateVariables `json:"template_variables,omitempty"`
		Title             *string                            `json:"title,omitempty"`
		Type              *PowerpackBackendDefinitionType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PowerpackId == nil {
		return fmt.Errorf("required field powerpack_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.BackgroundColor = all.BackgroundColor
	o.BannerImg = all.BannerImg
	o.PowerpackId = *all.PowerpackId
	o.ShowTitle = all.ShowTitle
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	if all.TemplateVariables != nil && all.TemplateVariables.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TemplateVariables = all.TemplateVariables
	o.Title = all.Title
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
