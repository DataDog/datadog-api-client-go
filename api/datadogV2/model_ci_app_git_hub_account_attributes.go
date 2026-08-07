// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppGitHubAccountAttributes Attributes describing a GitHub account's CI Visibility opt-in status.
type CIAppGitHubAccountAttributes struct {
	// The GitHub account (organization or user) name.
	Account *string `json:"account,omitempty"`
	// Whether CI Visibility is enabled at the account level.
	Enabled *bool `json:"enabled,omitempty"`
	// The GitHub host (`github.com` or a GHES hostname) this account belongs to.
	Host *string `json:"host,omitempty"`
	// The number of repositories known for this account.
	RepoCount *int64 `json:"repo_count,omitempty"`
	// The repositories belonging to this account, with their individual opt-in status.
	Repositories []CIAppGitHubAccountRepository `json:"repositories,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppGitHubAccountAttributes instantiates a new CIAppGitHubAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppGitHubAccountAttributes() *CIAppGitHubAccountAttributes {
	this := CIAppGitHubAccountAttributes{}
	return &this
}

// NewCIAppGitHubAccountAttributesWithDefaults instantiates a new CIAppGitHubAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppGitHubAccountAttributesWithDefaults() *CIAppGitHubAccountAttributes {
	this := CIAppGitHubAccountAttributes{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CIAppGitHubAccountAttributes) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountAttributes) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CIAppGitHubAccountAttributes) HasAccount() bool {
	return o != nil && o.Account != nil
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *CIAppGitHubAccountAttributes) SetAccount(v string) {
	o.Account = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CIAppGitHubAccountAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CIAppGitHubAccountAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CIAppGitHubAccountAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CIAppGitHubAccountAttributes) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountAttributes) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CIAppGitHubAccountAttributes) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CIAppGitHubAccountAttributes) SetHost(v string) {
	o.Host = &v
}

// GetRepoCount returns the RepoCount field value if set, zero value otherwise.
func (o *CIAppGitHubAccountAttributes) GetRepoCount() int64 {
	if o == nil || o.RepoCount == nil {
		var ret int64
		return ret
	}
	return *o.RepoCount
}

// GetRepoCountOk returns a tuple with the RepoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountAttributes) GetRepoCountOk() (*int64, bool) {
	if o == nil || o.RepoCount == nil {
		return nil, false
	}
	return o.RepoCount, true
}

// HasRepoCount returns a boolean if a field has been set.
func (o *CIAppGitHubAccountAttributes) HasRepoCount() bool {
	return o != nil && o.RepoCount != nil
}

// SetRepoCount gets a reference to the given int64 and assigns it to the RepoCount field.
func (o *CIAppGitHubAccountAttributes) SetRepoCount(v int64) {
	o.RepoCount = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *CIAppGitHubAccountAttributes) GetRepositories() []CIAppGitHubAccountRepository {
	if o == nil || o.Repositories == nil {
		var ret []CIAppGitHubAccountRepository
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppGitHubAccountAttributes) GetRepositoriesOk() (*[]CIAppGitHubAccountRepository, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return &o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *CIAppGitHubAccountAttributes) HasRepositories() bool {
	return o != nil && o.Repositories != nil
}

// SetRepositories gets a reference to the given []CIAppGitHubAccountRepository and assigns it to the Repositories field.
func (o *CIAppGitHubAccountAttributes) SetRepositories(v []CIAppGitHubAccountRepository) {
	o.Repositories = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppGitHubAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.RepoCount != nil {
		toSerialize["repo_count"] = o.RepoCount
	}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppGitHubAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Account      *string                        `json:"account,omitempty"`
		Enabled      *bool                          `json:"enabled,omitempty"`
		Host         *string                        `json:"host,omitempty"`
		RepoCount    *int64                         `json:"repo_count,omitempty"`
		Repositories []CIAppGitHubAccountRepository `json:"repositories,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account", "enabled", "host", "repo_count", "repositories"})
	} else {
		return err
	}
	o.Account = all.Account
	o.Enabled = all.Enabled
	o.Host = all.Host
	o.RepoCount = all.RepoCount
	o.Repositories = all.Repositories

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
