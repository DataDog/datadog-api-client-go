// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStalenessCodeReference A repository and its files that reference the feature flag.
type FeatureFlagStalenessCodeReference struct {
	// Paths of files that reference the feature flag in this repository.
	Files []string `json:"files,omitempty"`
	// The URL of the source code repository, when available.
	RepoUrl *string `json:"repo_url,omitempty"`
	// The identifier of the source code repository.
	ScmRepositoryId *string `json:"scm_repository_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagStalenessCodeReference instantiates a new FeatureFlagStalenessCodeReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagStalenessCodeReference() *FeatureFlagStalenessCodeReference {
	this := FeatureFlagStalenessCodeReference{}
	return &this
}

// NewFeatureFlagStalenessCodeReferenceWithDefaults instantiates a new FeatureFlagStalenessCodeReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagStalenessCodeReferenceWithDefaults() *FeatureFlagStalenessCodeReference {
	this := FeatureFlagStalenessCodeReference{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FeatureFlagStalenessCodeReference) GetFiles() []string {
	if o == nil || o.Files == nil {
		var ret []string
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessCodeReference) GetFilesOk() (*[]string, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FeatureFlagStalenessCodeReference) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given []string and assigns it to the Files field.
func (o *FeatureFlagStalenessCodeReference) SetFiles(v []string) {
	o.Files = v
}

// GetRepoUrl returns the RepoUrl field value if set, zero value otherwise.
func (o *FeatureFlagStalenessCodeReference) GetRepoUrl() string {
	if o == nil || o.RepoUrl == nil {
		var ret string
		return ret
	}
	return *o.RepoUrl
}

// GetRepoUrlOk returns a tuple with the RepoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessCodeReference) GetRepoUrlOk() (*string, bool) {
	if o == nil || o.RepoUrl == nil {
		return nil, false
	}
	return o.RepoUrl, true
}

// HasRepoUrl returns a boolean if a field has been set.
func (o *FeatureFlagStalenessCodeReference) HasRepoUrl() bool {
	return o != nil && o.RepoUrl != nil
}

// SetRepoUrl gets a reference to the given string and assigns it to the RepoUrl field.
func (o *FeatureFlagStalenessCodeReference) SetRepoUrl(v string) {
	o.RepoUrl = &v
}

// GetScmRepositoryId returns the ScmRepositoryId field value if set, zero value otherwise.
func (o *FeatureFlagStalenessCodeReference) GetScmRepositoryId() string {
	if o == nil || o.ScmRepositoryId == nil {
		var ret string
		return ret
	}
	return *o.ScmRepositoryId
}

// GetScmRepositoryIdOk returns a tuple with the ScmRepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessCodeReference) GetScmRepositoryIdOk() (*string, bool) {
	if o == nil || o.ScmRepositoryId == nil {
		return nil, false
	}
	return o.ScmRepositoryId, true
}

// HasScmRepositoryId returns a boolean if a field has been set.
func (o *FeatureFlagStalenessCodeReference) HasScmRepositoryId() bool {
	return o != nil && o.ScmRepositoryId != nil
}

// SetScmRepositoryId gets a reference to the given string and assigns it to the ScmRepositoryId field.
func (o *FeatureFlagStalenessCodeReference) SetScmRepositoryId(v string) {
	o.ScmRepositoryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagStalenessCodeReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.RepoUrl != nil {
		toSerialize["repo_url"] = o.RepoUrl
	}
	if o.ScmRepositoryId != nil {
		toSerialize["scm_repository_id"] = o.ScmRepositoryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagStalenessCodeReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Files           []string `json:"files,omitempty"`
		RepoUrl         *string  `json:"repo_url,omitempty"`
		ScmRepositoryId *string  `json:"scm_repository_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"files", "repo_url", "scm_repository_id"})
	} else {
		return err
	}
	o.Files = all.Files
	o.RepoUrl = all.RepoUrl
	o.ScmRepositoryId = all.ScmRepositoryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
