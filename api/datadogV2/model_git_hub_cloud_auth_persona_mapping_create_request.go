// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCloudAuthPersonaMappingCreateRequest Request used to create a GitHub cloud authentication persona mapping.
type GitHubCloudAuthPersonaMappingCreateRequest struct {
	// Data for creating a GitHub cloud authentication persona mapping.
	Data GitHubCloudAuthPersonaMappingCreateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGitHubCloudAuthPersonaMappingCreateRequest instantiates a new GitHubCloudAuthPersonaMappingCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCloudAuthPersonaMappingCreateRequest(data GitHubCloudAuthPersonaMappingCreateData) *GitHubCloudAuthPersonaMappingCreateRequest {
	this := GitHubCloudAuthPersonaMappingCreateRequest{}
	this.Data = data
	return &this
}

// NewGitHubCloudAuthPersonaMappingCreateRequestWithDefaults instantiates a new GitHubCloudAuthPersonaMappingCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCloudAuthPersonaMappingCreateRequestWithDefaults() *GitHubCloudAuthPersonaMappingCreateRequest {
	this := GitHubCloudAuthPersonaMappingCreateRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *GitHubCloudAuthPersonaMappingCreateRequest) GetData() GitHubCloudAuthPersonaMappingCreateData {
	if o == nil {
		var ret GitHubCloudAuthPersonaMappingCreateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GitHubCloudAuthPersonaMappingCreateRequest) GetDataOk() (*GitHubCloudAuthPersonaMappingCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GitHubCloudAuthPersonaMappingCreateRequest) SetData(v GitHubCloudAuthPersonaMappingCreateData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCloudAuthPersonaMappingCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCloudAuthPersonaMappingCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *GitHubCloudAuthPersonaMappingCreateData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
