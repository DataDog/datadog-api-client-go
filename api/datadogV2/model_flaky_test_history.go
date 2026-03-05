// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestHistory A single history entry representing a status change for a flaky test.
type FlakyTestHistory struct {
	// The commit SHA associated with this status change. Will be an empty string if the commit SHA is not available.
	CommitSha string `json:"commit_sha"`
	// The test status at this point in history.
	Status string `json:"status"`
	// Unix timestamp in milliseconds when this status change occurred.
	Timestamp int64 `json:"timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestHistory instantiates a new FlakyTestHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestHistory(commitSha string, status string, timestamp int64) *FlakyTestHistory {
	this := FlakyTestHistory{}
	this.CommitSha = commitSha
	this.Status = status
	this.Timestamp = timestamp
	return &this
}

// NewFlakyTestHistoryWithDefaults instantiates a new FlakyTestHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestHistoryWithDefaults() *FlakyTestHistory {
	this := FlakyTestHistory{}
	return &this
}

// GetCommitSha returns the CommitSha field value.
func (o *FlakyTestHistory) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *FlakyTestHistory) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value.
func (o *FlakyTestHistory) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetStatus returns the Status field value.
func (o *FlakyTestHistory) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FlakyTestHistory) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *FlakyTestHistory) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value.
func (o *FlakyTestHistory) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *FlakyTestHistory) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *FlakyTestHistory) SetTimestamp(v int64) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commit_sha"] = o.CommitSha
	toSerialize["status"] = o.Status
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitSha *string `json:"commit_sha"`
		Status    *string `json:"status"`
		Timestamp *int64  `json:"timestamp"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitSha == nil {
		return fmt.Errorf("required field commit_sha missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_sha", "status", "timestamp"})
	} else {
		return err
	}
	o.CommitSha = *all.CommitSha
	o.Status = *all.Status
	o.Timestamp = *all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
