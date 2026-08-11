// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppGitHubAccountUpdateRequestAttributes Attributes for updating a GitHub account's CI Visibility opt-in status.
// At least one of `enabled` or `repository.enabled` must be provided.
type CIAppGitHubAccountUpdateRequestAttributes struct {
	// The GitHub account (organization or user) name to update, identified by name.
	Account string `json:"account"`
	// Whether to enable or disable CI Visibility at the account level.
	Enabled *bool `json:"enabled,omitempty"`
	// The GitHub host (`github.com` or a GHES hostname) the account belongs to. Required to disambiguate
	// when the same account name exists on more than one host.
	Host *string `json:"host,omitempty"`
	// Repository-level opt-in change to apply, identified by name.
	Repository *CIAppGitHubAccountUpdateRequestRepository `json:"repository,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppGitHubAccountUpdateRequestAttributes instantiates a new CIAppGitHubAccountUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppGitHubAccountUpdateRequestAttributes(account string) *CIAppGitHubAccountUpdateRequestAttributes {
	this := CIAppGitHubAccountUpdateRequestAttributes{}
	this.Account = account
	return &this
}

// NewCIAppGitHubAccountUpdateRequestAttributesWithDefaults instantiates a new CIAppGitHubAccountUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppGitHubAccountUpdateRequestAttributesWithDefaults() *CIAppGitHubAccountUpdateRequestAttributes {
	this := CIAppGitHubAccountUpdateRequestAttributes{}
	return &this
}

// GetAccount returns the Account field value.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value.
func (o *CIAppGitHubAccountUpdateRequestAttributes) SetAccount(v string) {
	o.Account = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CIAppGitHubAccountUpdateRequestAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CIAppGitHubAccountUpdateRequestAttributes) SetHost(v string) {
	o.Host = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetRepository() CIAppGitHubAccountUpdateRequestRepository {
	if o == nil || o.Repository == nil {
		var ret CIAppGitHubAccountUpdateRequestRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) GetRepositoryOk() (*CIAppGitHubAccountUpdateRequestRepository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CIAppGitHubAccountUpdateRequestAttributes) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given CIAppGitHubAccountUpdateRequestRepository and assigns it to the Repository field.
func (o *CIAppGitHubAccountUpdateRequestAttributes) SetRepository(v CIAppGitHubAccountUpdateRequestRepository) {
	o.Repository = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppGitHubAccountUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account"] = o.Account
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppGitHubAccountUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Account    *string                                    `json:"account"`
		Enabled    *bool                                      `json:"enabled,omitempty"`
		Host       *string                                    `json:"host,omitempty"`
		Repository *CIAppGitHubAccountUpdateRequestRepository `json:"repository,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Account == nil {
		return fmt.Errorf("required field account missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account", "enabled", "host", "repository"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Account = *all.Account
	o.Enabled = all.Enabled
	o.Host = all.Host
	if all.Repository != nil && all.Repository.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Repository = all.Repository

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
