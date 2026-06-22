// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationCodeSession A linked code session (for example, a pull request) for the investigation.
type RemediationCodeSession struct {
	// Code session ID.
	Id *string `json:"id,omitempty"`
	// Pull request status for a linked code session.
	PullRequestStatus *RemediationPullRequestStatus `json:"pull_request_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationCodeSession instantiates a new RemediationCodeSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationCodeSession() *RemediationCodeSession {
	this := RemediationCodeSession{}
	return &this
}

// NewRemediationCodeSessionWithDefaults instantiates a new RemediationCodeSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationCodeSessionWithDefaults() *RemediationCodeSession {
	this := RemediationCodeSession{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemediationCodeSession) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationCodeSession) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemediationCodeSession) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemediationCodeSession) SetId(v string) {
	o.Id = &v
}

// GetPullRequestStatus returns the PullRequestStatus field value if set, zero value otherwise.
func (o *RemediationCodeSession) GetPullRequestStatus() RemediationPullRequestStatus {
	if o == nil || o.PullRequestStatus == nil {
		var ret RemediationPullRequestStatus
		return ret
	}
	return *o.PullRequestStatus
}

// GetPullRequestStatusOk returns a tuple with the PullRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationCodeSession) GetPullRequestStatusOk() (*RemediationPullRequestStatus, bool) {
	if o == nil || o.PullRequestStatus == nil {
		return nil, false
	}
	return o.PullRequestStatus, true
}

// HasPullRequestStatus returns a boolean if a field has been set.
func (o *RemediationCodeSession) HasPullRequestStatus() bool {
	return o != nil && o.PullRequestStatus != nil
}

// SetPullRequestStatus gets a reference to the given RemediationPullRequestStatus and assigns it to the PullRequestStatus field.
func (o *RemediationCodeSession) SetPullRequestStatus(v RemediationPullRequestStatus) {
	o.PullRequestStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationCodeSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PullRequestStatus != nil {
		toSerialize["pull_request_status"] = o.PullRequestStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationCodeSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string                       `json:"id,omitempty"`
		PullRequestStatus *RemediationPullRequestStatus `json:"pull_request_status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "pull_request_status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.PullRequestStatus != nil && !all.PullRequestStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.PullRequestStatus = all.PullRequestStatus
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
