// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemTemplateAttributesResponse Attributes of a postmortem template returned in a response.
type PostmortemTemplateAttributesResponse struct {
	// Settings for a postmortem template stored in Confluence. Required when `location` is `confluence`.
	ConfluencePostmortemSettings *ConfluencePostmortemSettings `json:"confluence_postmortem_settings,omitempty"`
	// The templated content of the postmortem, supporting Markdown and incident template variables.
	Content string `json:"content"`
	// When the template was created.
	CreatedAt time.Time `json:"createdAt"`
	// Settings for a postmortem template stored in Google Docs. Required when `location` is `google_docs`.
	GoogleDocsPostmortemSettings *GoogleDocsPostmortemSettings `json:"google_docs_postmortem_settings,omitempty"`
	// When set, marks this template as a default. The effective default for an incident type is the template with the most recent `is_default` timestamp.
	IsDefault datadog.NullableTime `json:"is_default"`
	// The location where the postmortem is created and stored.
	Location PostmortemTemplateLocation `json:"location"`
	// When the template was last modified.
	ModifiedAt time.Time `json:"modifiedAt"`
	// The name of the template.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostmortemTemplateAttributesResponse instantiates a new PostmortemTemplateAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostmortemTemplateAttributesResponse(content string, createdAt time.Time, isDefault datadog.NullableTime, location PostmortemTemplateLocation, modifiedAt time.Time, name string) *PostmortemTemplateAttributesResponse {
	this := PostmortemTemplateAttributesResponse{}
	this.Content = content
	this.CreatedAt = createdAt
	this.IsDefault = isDefault
	this.Location = location
	this.ModifiedAt = modifiedAt
	this.Name = name
	return &this
}

// NewPostmortemTemplateAttributesResponseWithDefaults instantiates a new PostmortemTemplateAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostmortemTemplateAttributesResponseWithDefaults() *PostmortemTemplateAttributesResponse {
	this := PostmortemTemplateAttributesResponse{}
	var location PostmortemTemplateLocation = POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS
	this.Location = location
	return &this
}

// GetConfluencePostmortemSettings returns the ConfluencePostmortemSettings field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesResponse) GetConfluencePostmortemSettings() ConfluencePostmortemSettings {
	if o == nil || o.ConfluencePostmortemSettings == nil {
		var ret ConfluencePostmortemSettings
		return ret
	}
	return *o.ConfluencePostmortemSettings
}

// GetConfluencePostmortemSettingsOk returns a tuple with the ConfluencePostmortemSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetConfluencePostmortemSettingsOk() (*ConfluencePostmortemSettings, bool) {
	if o == nil || o.ConfluencePostmortemSettings == nil {
		return nil, false
	}
	return o.ConfluencePostmortemSettings, true
}

// HasConfluencePostmortemSettings returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesResponse) HasConfluencePostmortemSettings() bool {
	return o != nil && o.ConfluencePostmortemSettings != nil
}

// SetConfluencePostmortemSettings gets a reference to the given ConfluencePostmortemSettings and assigns it to the ConfluencePostmortemSettings field.
func (o *PostmortemTemplateAttributesResponse) SetConfluencePostmortemSettings(v ConfluencePostmortemSettings) {
	o.ConfluencePostmortemSettings = &v
}

// GetContent returns the Content field value.
func (o *PostmortemTemplateAttributesResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *PostmortemTemplateAttributesResponse) SetContent(v string) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *PostmortemTemplateAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *PostmortemTemplateAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetGoogleDocsPostmortemSettings returns the GoogleDocsPostmortemSettings field value if set, zero value otherwise.
func (o *PostmortemTemplateAttributesResponse) GetGoogleDocsPostmortemSettings() GoogleDocsPostmortemSettings {
	if o == nil || o.GoogleDocsPostmortemSettings == nil {
		var ret GoogleDocsPostmortemSettings
		return ret
	}
	return *o.GoogleDocsPostmortemSettings
}

// GetGoogleDocsPostmortemSettingsOk returns a tuple with the GoogleDocsPostmortemSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetGoogleDocsPostmortemSettingsOk() (*GoogleDocsPostmortemSettings, bool) {
	if o == nil || o.GoogleDocsPostmortemSettings == nil {
		return nil, false
	}
	return o.GoogleDocsPostmortemSettings, true
}

// HasGoogleDocsPostmortemSettings returns a boolean if a field has been set.
func (o *PostmortemTemplateAttributesResponse) HasGoogleDocsPostmortemSettings() bool {
	return o != nil && o.GoogleDocsPostmortemSettings != nil
}

// SetGoogleDocsPostmortemSettings gets a reference to the given GoogleDocsPostmortemSettings and assigns it to the GoogleDocsPostmortemSettings field.
func (o *PostmortemTemplateAttributesResponse) SetGoogleDocsPostmortemSettings(v GoogleDocsPostmortemSettings) {
	o.GoogleDocsPostmortemSettings = &v
}

// GetIsDefault returns the IsDefault field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *PostmortemTemplateAttributesResponse) GetIsDefault() time.Time {
	if o == nil || o.IsDefault.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostmortemTemplateAttributesResponse) GetIsDefaultOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// SetIsDefault sets field value.
func (o *PostmortemTemplateAttributesResponse) SetIsDefault(v time.Time) {
	o.IsDefault.Set(&v)
}

// GetLocation returns the Location field value.
func (o *PostmortemTemplateAttributesResponse) GetLocation() PostmortemTemplateLocation {
	if o == nil {
		var ret PostmortemTemplateLocation
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetLocationOk() (*PostmortemTemplateLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value.
func (o *PostmortemTemplateAttributesResponse) SetLocation(v PostmortemTemplateLocation) {
	o.Location = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *PostmortemTemplateAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *PostmortemTemplateAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *PostmortemTemplateAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostmortemTemplateAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PostmortemTemplateAttributesResponse) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostmortemTemplateAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfluencePostmortemSettings != nil {
		toSerialize["confluence_postmortem_settings"] = o.ConfluencePostmortemSettings
	}
	toSerialize["content"] = o.Content
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.GoogleDocsPostmortemSettings != nil {
		toSerialize["google_docs_postmortem_settings"] = o.GoogleDocsPostmortemSettings
	}
	toSerialize["is_default"] = o.IsDefault.Get()
	toSerialize["location"] = o.Location
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostmortemTemplateAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfluencePostmortemSettings *ConfluencePostmortemSettings `json:"confluence_postmortem_settings,omitempty"`
		Content                      *string                       `json:"content"`
		CreatedAt                    *time.Time                    `json:"createdAt"`
		GoogleDocsPostmortemSettings *GoogleDocsPostmortemSettings `json:"google_docs_postmortem_settings,omitempty"`
		IsDefault                    datadog.NullableTime          `json:"is_default"`
		Location                     *PostmortemTemplateLocation   `json:"location"`
		ModifiedAt                   *time.Time                    `json:"modifiedAt"`
		Name                         *string                       `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if !all.IsDefault.IsSet() {
		return fmt.Errorf("required field is_default missing")
	}
	if all.Location == nil {
		return fmt.Errorf("required field location missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modifiedAt missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confluence_postmortem_settings", "content", "createdAt", "google_docs_postmortem_settings", "is_default", "location", "modifiedAt", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConfluencePostmortemSettings != nil && all.ConfluencePostmortemSettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfluencePostmortemSettings = all.ConfluencePostmortemSettings
	o.Content = *all.Content
	o.CreatedAt = *all.CreatedAt
	if all.GoogleDocsPostmortemSettings != nil && all.GoogleDocsPostmortemSettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GoogleDocsPostmortemSettings = all.GoogleDocsPostmortemSettings
	o.IsDefault = all.IsDefault
	if !all.Location.IsValid() {
		hasInvalidField = true
	} else {
		o.Location = *all.Location
	}
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
