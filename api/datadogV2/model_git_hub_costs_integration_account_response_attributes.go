// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountResponseAttributes Attributes of a GitHub Costs integration account returned in responses.
type GitHubCostsIntegrationAccountResponseAttributes struct {
	// Authentication configured on the GitHub Costs integration account.
	Authentication *GitHubCostsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from GitHub Costs, keyed by dataflow id.
	Dataflows *GitHubCostsIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the GitHub Costs integration account.
	Name string `json:"name"`
	// Settings configured on the GitHub Costs integration account.
	Settings GitHubCostsIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountResponseAttributes instantiates a new GitHubCostsIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountResponseAttributes(name string, settings GitHubCostsIntegrationAccountSettingsResponse) *GitHubCostsIntegrationAccountResponseAttributes {
	this := GitHubCostsIntegrationAccountResponseAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewGitHubCostsIntegrationAccountResponseAttributesWithDefaults instantiates a new GitHubCostsIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountResponseAttributesWithDefaults() *GitHubCostsIntegrationAccountResponseAttributes {
	this := GitHubCostsIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetAuthentication() GitHubCostsIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret GitHubCostsIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetAuthenticationOk() (*GitHubCostsIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given GitHubCostsIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *GitHubCostsIntegrationAccountResponseAttributes) SetAuthentication(v GitHubCostsIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetDataflows() GitHubCostsIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret GitHubCostsIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetDataflowsOk() (*GitHubCostsIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given GitHubCostsIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *GitHubCostsIntegrationAccountResponseAttributes) SetDataflows(v GitHubCostsIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GitHubCostsIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetSettings() GitHubCostsIntegrationAccountSettingsResponse {
	if o == nil {
		var ret GitHubCostsIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountResponseAttributes) GetSettingsOk() (*GitHubCostsIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *GitHubCostsIntegrationAccountResponseAttributes) SetSettings(v GitHubCostsIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *GitHubCostsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *GitHubCostsIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                              `json:"name"`
		Settings       *GitHubCostsIntegrationAccountSettingsResponse       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	if all.Dataflows != nil && all.Dataflows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataflows = all.Dataflows
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
