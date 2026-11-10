// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountCreateAttributes Writable attributes used to create a GitHub Costs integration account.
type GitHubCostsIntegrationAccountCreateAttributes struct {
	// Authentication for creating the GitHub Costs integration account. Exactly one method is set.
	Authentication GitHubCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Human-readable name of the GitHub Costs integration account.
	Name string `json:"name"`
	// Settings for creating the GitHub Costs integration account.
	Settings GitHubCostsIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountCreateAttributes instantiates a new GitHubCostsIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountCreateAttributes(authentication GitHubCostsIntegrationAccountAuthenticationRequest, name string, settings GitHubCostsIntegrationAccountSettingsRequest) *GitHubCostsIntegrationAccountCreateAttributes {
	this := GitHubCostsIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewGitHubCostsIntegrationAccountCreateAttributesWithDefaults instantiates a new GitHubCostsIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountCreateAttributesWithDefaults() *GitHubCostsIntegrationAccountCreateAttributes {
	this := GitHubCostsIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetAuthentication() GitHubCostsIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret GitHubCostsIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetAuthenticationOk() (*GitHubCostsIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) SetAuthentication(v GitHubCostsIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetName returns the Name field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetSettings() GitHubCostsIntegrationAccountSettingsRequest {
	if o == nil {
		var ret GitHubCostsIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountCreateAttributes) GetSettingsOk() (*GitHubCostsIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *GitHubCostsIntegrationAccountCreateAttributes) SetSettings(v GitHubCostsIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *GitHubCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
		Name           *string                                             `json:"name"`
		Settings       *GitHubCostsIntegrationAccountSettingsRequest       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Authentication == nil {
		return fmt.Errorf("required field authentication missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = *all.Authentication
	o.Name = *all.Name
	if all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
