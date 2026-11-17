// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountUpdateAttributes Writable attributes used to update a GitHub Costs integration account. Every field is optional; only the fields provided are changed.
type GitHubCostsIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the GitHub Costs integration account. Exactly one method is set.
	Authentication *GitHubCostsIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Human-readable name of the GitHub Costs integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the GitHub Costs integration account. Only the fields provided are changed.
	Settings *GitHubCostsIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountUpdateAttributes instantiates a new GitHubCostsIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountUpdateAttributes() *GitHubCostsIntegrationAccountUpdateAttributes {
	this := GitHubCostsIntegrationAccountUpdateAttributes{}
	return &this
}

// NewGitHubCostsIntegrationAccountUpdateAttributesWithDefaults instantiates a new GitHubCostsIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountUpdateAttributesWithDefaults() *GitHubCostsIntegrationAccountUpdateAttributes {
	this := GitHubCostsIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetAuthentication() GitHubCostsIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret GitHubCostsIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*GitHubCostsIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given GitHubCostsIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) SetAuthentication(v GitHubCostsIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetSettings() GitHubCostsIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret GitHubCostsIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) GetSettingsOk() (*GitHubCostsIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given GitHubCostsIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) SetSettings(v GitHubCostsIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *GitHubCostsIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Name           *string                                            `json:"name,omitempty"`
		Settings       *GitHubCostsIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	o.Name = all.Name
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
