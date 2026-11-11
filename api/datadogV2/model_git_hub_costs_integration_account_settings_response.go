// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountSettingsResponse Settings configured on the GitHub Costs integration account.
type GitHubCostsIntegrationAccountSettingsResponse struct {
	// Your GitHub enterprise slug, found on the enterprise settings page (e.g. enterprise-name), or an organization login (e.g. organization-name). An enterprise already covers every organization it owns.
	EnterpriseName string `json:"enterprise_name"`
	// Comma-separated tags added to the collected costs, for example `env:prod,team:example`.
	Tags *string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountSettingsResponse instantiates a new GitHubCostsIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountSettingsResponse(enterpriseName string) *GitHubCostsIntegrationAccountSettingsResponse {
	this := GitHubCostsIntegrationAccountSettingsResponse{}
	this.EnterpriseName = enterpriseName
	return &this
}

// NewGitHubCostsIntegrationAccountSettingsResponseWithDefaults instantiates a new GitHubCostsIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountSettingsResponseWithDefaults() *GitHubCostsIntegrationAccountSettingsResponse {
	this := GitHubCostsIntegrationAccountSettingsResponse{}
	return &this
}

// GetEnterpriseName returns the EnterpriseName field value.
func (o *GitHubCostsIntegrationAccountSettingsResponse) GetEnterpriseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountSettingsResponse) GetEnterpriseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnterpriseName, true
}

// SetEnterpriseName sets field value.
func (o *GitHubCostsIntegrationAccountSettingsResponse) SetEnterpriseName(v string) {
	o.EnterpriseName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountSettingsResponse) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountSettingsResponse) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountSettingsResponse) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *GitHubCostsIntegrationAccountSettingsResponse) SetTags(v string) {
	o.Tags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enterprise_name"] = o.EnterpriseName
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnterpriseName *string `json:"enterprise_name"`
		Tags           *string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnterpriseName == nil {
		return fmt.Errorf("required field enterprise_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enterprise_name", "tags"})
	} else {
		return err
	}
	o.EnterpriseName = *all.EnterpriseName
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
