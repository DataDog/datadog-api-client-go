// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemTemplateAttributesRequest Attributes for creating or updating a postmortem template.
type PostmortemTemplateAttributesRequest struct {
	// Settings for a postmortem template stored in Confluence. Required when `location` is `confluence`.
	ConfluencePostmortemSettings *ConfluencePostmortemSettings `json:"confluence_postmortem_settings,omitempty"`
	// The templated content of the postmortem, supporting Markdown and incident template variables.
	Content *string `json:"content,omitempty"`
	// Settings for a postmortem template stored in Google Docs. Required when `location` is `google_docs`.
	GoogleDocsPostmortemSettings *GoogleDocsPostmortemSettings `json:"google_docs_postmortem_settings,omitempty"`
	// When set, marks this template as a default. The effective default for an incident type is the template with the most recent `is_default` timestamp. Set to `null` to unset.
	IsDefault datadog.NullableTime `json:"is_default,omitempty"`
	// The location where the postmortem is created and stored.
	Location *PostmortemTemplateLocation `json:"location,omitempty"`
	// The name of the template.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostmortemTemplateAttributesRequest instantiates a new PostmortemTemplateAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostmortemTemplateAttributesRequest(name string) *PostmortemTemplateAttributesRequest {
	this := PostmortemTemplateAttributesRequest{}
	var location PostmortemTemplateLocation = POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS
	this.Location = &location
	this.Name = name
	return &this
}

// NewPostmortemTemplateAttributesRequestWithDefaults instantiates a new PostmortemTemplateAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostmortemTemplateAttributesRequestWithDefaults() *PostmortemTemplateAttributesRequest {
	this := PostmortemTemplateAttributesRequest{}
	var location PostmortemTemplateLocation = POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS
	this.Location = &location
	return &this
}

// GetConfluencePostmortemSettings returns the ConfluencePostmortemSettings field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesRequest) GetConfluencePostmortemSettings() ConfluencePostmortemSettings {
	if o == nil || o.ConfluencePostmortemSettings == nil {
		var ret ConfluencePostmortemSettings
		return ret
	}
	return *o.ConfluencePostmortemSettings
}

// GetConfluencePostmortemSettingsOk returns a tuple with the ConfluencePostmortemSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesRequest) GetConfluencePostmortemSettingsOk() (*ConfluencePostmortemSettings, bool) {
	if o == nil || o.ConfluencePostmortemSettings == nil {
		return nil, false
	}
	return o.ConfluencePostmortemSettings, true
}

// HasConfluencePostmortemSettings returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesRequest) HasConfluencePostmortemSettings() bool {
	return o != nil && o.ConfluencePostmortemSettings != nil
}

// SetConfluencePostmortemSettings gets a reference to the given ConfluencePostmortemSettings and assigns it to the ConfluencePostmortemSettings field.
func (o *PostmortemTemplateAttributesRequest) SetConfluencePostmortemSettings(v ConfluencePostmortemSettings) {
	o.ConfluencePostmortemSettings = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesRequest) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesRequest) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesRequest) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PostmortemTemplateAttributesRequest) SetContent(v string) {
	o.Content = &v
}

// GetGoogleDocsPostmortemSettings returns the GoogleDocsPostmortemSettings field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesRequest) GetGoogleDocsPostmortemSettings() GoogleDocsPostmortemSettings {
	if o == nil || o.GoogleDocsPostmortemSettings == nil {
		var ret GoogleDocsPostmortemSettings
		return ret
	}
	return *o.GoogleDocsPostmortemSettings
}

// GetGoogleDocsPostmortemSettingsOk returns a tuple with the GoogleDocsPostmortemSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesRequest) GetGoogleDocsPostmortemSettingsOk() (*GoogleDocsPostmortemSettings, bool) {
	if o == nil || o.GoogleDocsPostmortemSettings == nil {
		return nil, false
	}
	return o.GoogleDocsPostmortemSettings, true
}

// HasGoogleDocsPostmortemSettings returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesRequest) HasGoogleDocsPostmortemSettings() bool {
	return o != nil && o.GoogleDocsPostmortemSettings != nil
}

// SetGoogleDocsPostmortemSettings gets a reference to the given GoogleDocsPostmortemSettings and assigns it to the GoogleDocsPostmortemSettings field.
func (o *PostmortemTemplateAttributesRequest) SetGoogleDocsPostmortemSettings(v GoogleDocsPostmortemSettings) {
	o.GoogleDocsPostmortemSettings = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostmortemTemplateAttributesRequest) GetIsDefault() time.Time {
	if o == nil || o.IsDefault.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostmortemTemplateAttributesRequest) GetIsDefaultOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesRequest) HasIsDefault() bool {
	return o != nil && o.IsDefault.IsSet()
}

// SetIsDefault gets a reference to the given datadog.NullableTime and assigns it to the IsDefault field.
func (o *PostmortemTemplateAttributesRequest) SetIsDefault(v time.Time) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil.
func (o *PostmortemTemplateAttributesRequest) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil.
func (o *PostmortemTemplateAttributesRequest) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesRequest) GetLocation() PostmortemTemplateLocation {
	if o == nil || o.Location == nil {
		var ret PostmortemTemplateLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesRequest) GetLocationOk() (*PostmortemTemplateLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesRequest) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given PostmortemTemplateLocation and assigns it to the Location field.
func (o *PostmortemTemplateAttributesRequest) SetLocation(v PostmortemTemplateLocation) {
	o.Location = &v
}

// GetName returns the Name field value.
func (o *PostmortemTemplateAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PostmortemTemplateAttributesRequest) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostmortemTemplateAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfluencePostmortemSettings != nil {
		toSerialize["confluence_postmortem_settings"] = o.ConfluencePostmortemSettings
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.GoogleDocsPostmortemSettings != nil {
		toSerialize["google_docs_postmortem_settings"] = o.GoogleDocsPostmortemSettings
	}
	if o.IsDefault.IsSet() {
		toSerialize["is_default"] = o.IsDefault.Get()
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostmortemTemplateAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfluencePostmortemSettings *ConfluencePostmortemSettings `json:"confluence_postmortem_settings,omitempty"`
		Content                      *string                       `json:"content,omitempty"`
		GoogleDocsPostmortemSettings *GoogleDocsPostmortemSettings `json:"google_docs_postmortem_settings,omitempty"`
		IsDefault                    datadog.NullableTime          `json:"is_default,omitempty"`
		Location                     *PostmortemTemplateLocation   `json:"location,omitempty"`
		Name                         *string                       `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confluence_postmortem_settings", "content", "google_docs_postmortem_settings", "is_default", "location", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConfluencePostmortemSettings != nil && all.ConfluencePostmortemSettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfluencePostmortemSettings = all.ConfluencePostmortemSettings
	o.Content = all.Content
	if all.GoogleDocsPostmortemSettings != nil && all.GoogleDocsPostmortemSettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GoogleDocsPostmortemSettings = all.GoogleDocsPostmortemSettings
	o.IsDefault = all.IsDefault
	if all.Location != nil && !all.Location.IsValid() {
		hasInvalidField = true
	} else {
		o.Location = all.Location
	}
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
