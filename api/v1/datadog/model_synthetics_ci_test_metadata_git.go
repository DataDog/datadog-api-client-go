/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsCITestMetadataGit Git information.
type SyntheticsCITestMetadataGit struct {
	// Branch name.
	Branch *string `json:"branch,omitempty"`
	// Commit SHA.
	CommitSha *string `json:"commit_sha,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsCITestMetadataGit instantiates a new SyntheticsCITestMetadataGit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsCITestMetadataGit() *SyntheticsCITestMetadataGit {
	this := SyntheticsCITestMetadataGit{}
	return &this
}

// NewSyntheticsCITestMetadataGitWithDefaults instantiates a new SyntheticsCITestMetadataGit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsCITestMetadataGitWithDefaults() *SyntheticsCITestMetadataGit {
	this := SyntheticsCITestMetadataGit{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *SyntheticsCITestMetadataGit) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITestMetadataGit) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *SyntheticsCITestMetadataGit) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *SyntheticsCITestMetadataGit) SetBranch(v string) {
	o.Branch = &v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *SyntheticsCITestMetadataGit) GetCommitSha() string {
	if o == nil || o.CommitSha == nil {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITestMetadataGit) GetCommitShaOk() (*string, bool) {
	if o == nil || o.CommitSha == nil {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *SyntheticsCITestMetadataGit) HasCommitSha() bool {
	if o != nil && o.CommitSha != nil {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *SyntheticsCITestMetadataGit) SetCommitSha(v string) {
	o.CommitSha = &v
}

func (o SyntheticsCITestMetadataGit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.CommitSha != nil {
		toSerialize["commit_sha"] = o.CommitSha
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsCITestMetadataGit) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Branch    *string `json:"branch,omitempty"`
		CommitSha *string `json:"commit_sha,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Branch = all.Branch
	o.CommitSha = all.CommitSha
	return nil
}

type NullableSyntheticsCITestMetadataGit struct {
	value *SyntheticsCITestMetadataGit
	isSet bool
}

func (v NullableSyntheticsCITestMetadataGit) Get() *SyntheticsCITestMetadataGit {
	return v.value
}

func (v *NullableSyntheticsCITestMetadataGit) Set(val *SyntheticsCITestMetadataGit) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsCITestMetadataGit) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsCITestMetadataGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsCITestMetadataGit(val *SyntheticsCITestMetadataGit) *NullableSyntheticsCITestMetadataGit {
	return &NullableSyntheticsCITestMetadataGit{value: val, isSet: true}
}

func (v NullableSyntheticsCITestMetadataGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsCITestMetadataGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
