// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRenderTemplateDataAttributesRequest Attributes for rendering a template.
type IncidentRenderTemplateDataAttributesRequest struct {
	// The template content to render.
	Content string `json:"content"`
	// The date-time format to use for rendering.
	DatetimeFormat *string `json:"datetime_format,omitempty"`
	// The timezone to use for rendering.
	Timezone *string `json:"timezone,omitempty"`
	// Whether to validate links in the rendered template.
	ValidateLinks datadog.NullableBool `json:"validate_links,omitempty"`
	// Whether to validate variables in the template.
	ValidateVariables datadog.NullableBool `json:"validate_variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRenderTemplateDataAttributesRequest instantiates a new IncidentRenderTemplateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRenderTemplateDataAttributesRequest(content string) *IncidentRenderTemplateDataAttributesRequest {
	this := IncidentRenderTemplateDataAttributesRequest{}
	this.Content = content
	return &this
}

// NewIncidentRenderTemplateDataAttributesRequestWithDefaults instantiates a new IncidentRenderTemplateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRenderTemplateDataAttributesRequestWithDefaults() *IncidentRenderTemplateDataAttributesRequest {
	this := IncidentRenderTemplateDataAttributesRequest{}
	return &this
}

// GetContent returns the Content field value.
func (o *IncidentRenderTemplateDataAttributesRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentRenderTemplateDataAttributesRequest) SetContent(v string) {
	o.Content = v
}

// GetDatetimeFormat returns the DatetimeFormat field value if set, zero value otherwise.
func (o *IncidentRenderTemplateDataAttributesRequest) GetDatetimeFormat() string {
	if o == nil || o.DatetimeFormat == nil {
		var ret string
		return ret
	}
	return *o.DatetimeFormat
}

// GetDatetimeFormatOk returns a tuple with the DatetimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) GetDatetimeFormatOk() (*string, bool) {
	if o == nil || o.DatetimeFormat == nil {
		return nil, false
	}
	return o.DatetimeFormat, true
}

// HasDatetimeFormat returns a boolean if a field has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) HasDatetimeFormat() bool {
	return o != nil && o.DatetimeFormat != nil
}

// SetDatetimeFormat gets a reference to the given string and assigns it to the DatetimeFormat field.
func (o *IncidentRenderTemplateDataAttributesRequest) SetDatetimeFormat(v string) {
	o.DatetimeFormat = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *IncidentRenderTemplateDataAttributesRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *IncidentRenderTemplateDataAttributesRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetValidateLinks returns the ValidateLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRenderTemplateDataAttributesRequest) GetValidateLinks() bool {
	if o == nil || o.ValidateLinks.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ValidateLinks.Get()
}

// GetValidateLinksOk returns a tuple with the ValidateLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRenderTemplateDataAttributesRequest) GetValidateLinksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidateLinks.Get(), o.ValidateLinks.IsSet()
}

// HasValidateLinks returns a boolean if a field has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) HasValidateLinks() bool {
	return o != nil && o.ValidateLinks.IsSet()
}

// SetValidateLinks gets a reference to the given datadog.NullableBool and assigns it to the ValidateLinks field.
func (o *IncidentRenderTemplateDataAttributesRequest) SetValidateLinks(v bool) {
	o.ValidateLinks.Set(&v)
}

// SetValidateLinksNil sets the value for ValidateLinks to be an explicit nil.
func (o *IncidentRenderTemplateDataAttributesRequest) SetValidateLinksNil() {
	o.ValidateLinks.Set(nil)
}

// UnsetValidateLinks ensures that no value is present for ValidateLinks, not even an explicit nil.
func (o *IncidentRenderTemplateDataAttributesRequest) UnsetValidateLinks() {
	o.ValidateLinks.Unset()
}

// GetValidateVariables returns the ValidateVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRenderTemplateDataAttributesRequest) GetValidateVariables() bool {
	if o == nil || o.ValidateVariables.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ValidateVariables.Get()
}

// GetValidateVariablesOk returns a tuple with the ValidateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRenderTemplateDataAttributesRequest) GetValidateVariablesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidateVariables.Get(), o.ValidateVariables.IsSet()
}

// HasValidateVariables returns a boolean if a field has been set.
func (o *IncidentRenderTemplateDataAttributesRequest) HasValidateVariables() bool {
	return o != nil && o.ValidateVariables.IsSet()
}

// SetValidateVariables gets a reference to the given datadog.NullableBool and assigns it to the ValidateVariables field.
func (o *IncidentRenderTemplateDataAttributesRequest) SetValidateVariables(v bool) {
	o.ValidateVariables.Set(&v)
}

// SetValidateVariablesNil sets the value for ValidateVariables to be an explicit nil.
func (o *IncidentRenderTemplateDataAttributesRequest) SetValidateVariablesNil() {
	o.ValidateVariables.Set(nil)
}

// UnsetValidateVariables ensures that no value is present for ValidateVariables, not even an explicit nil.
func (o *IncidentRenderTemplateDataAttributesRequest) UnsetValidateVariables() {
	o.ValidateVariables.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRenderTemplateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content
	if o.DatetimeFormat != nil {
		toSerialize["datetime_format"] = o.DatetimeFormat
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.ValidateLinks.IsSet() {
		toSerialize["validate_links"] = o.ValidateLinks.Get()
	}
	if o.ValidateVariables.IsSet() {
		toSerialize["validate_variables"] = o.ValidateVariables.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRenderTemplateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content           *string              `json:"content"`
		DatetimeFormat    *string              `json:"datetime_format,omitempty"`
		Timezone          *string              `json:"timezone,omitempty"`
		ValidateLinks     datadog.NullableBool `json:"validate_links,omitempty"`
		ValidateVariables datadog.NullableBool `json:"validate_variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "datetime_format", "timezone", "validate_links", "validate_variables"})
	} else {
		return err
	}
	o.Content = *all.Content
	o.DatetimeFormat = all.DatetimeFormat
	o.Timezone = all.Timezone
	o.ValidateLinks = all.ValidateLinks
	o.ValidateVariables = all.ValidateVariables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
