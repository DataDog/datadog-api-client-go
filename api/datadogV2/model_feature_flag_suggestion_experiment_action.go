// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagSuggestionExperimentAction Server-generated reference to an experiment action requiring approval.
type FeatureFlagSuggestionExperimentAction struct {
	// The proposed experiment action.
	Action string `json:"action"`
	// The ID of the experiment.
	ExperimentId uuid.UUID `json:"experiment_id"`
	// The version of the experiment action reference.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagSuggestionExperimentAction instantiates a new FeatureFlagSuggestionExperimentAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagSuggestionExperimentAction(action string, experimentId uuid.UUID, version int64) *FeatureFlagSuggestionExperimentAction {
	this := FeatureFlagSuggestionExperimentAction{}
	this.Action = action
	this.ExperimentId = experimentId
	this.Version = version
	return &this
}

// NewFeatureFlagSuggestionExperimentActionWithDefaults instantiates a new FeatureFlagSuggestionExperimentAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagSuggestionExperimentActionWithDefaults() *FeatureFlagSuggestionExperimentAction {
	this := FeatureFlagSuggestionExperimentAction{}
	return &this
}

// GetAction returns the Action field value.
func (o *FeatureFlagSuggestionExperimentAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionExperimentAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *FeatureFlagSuggestionExperimentAction) SetAction(v string) {
	o.Action = v
}

// GetExperimentId returns the ExperimentId field value.
func (o *FeatureFlagSuggestionExperimentAction) GetExperimentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionExperimentAction) GetExperimentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentId, true
}

// SetExperimentId sets field value.
func (o *FeatureFlagSuggestionExperimentAction) SetExperimentId(v uuid.UUID) {
	o.ExperimentId = v
}

// GetVersion returns the Version field value.
func (o *FeatureFlagSuggestionExperimentAction) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionExperimentAction) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FeatureFlagSuggestionExperimentAction) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagSuggestionExperimentAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["experiment_id"] = o.ExperimentId
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagSuggestionExperimentAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action       *string    `json:"action"`
		ExperimentId *uuid.UUID `json:"experiment_id"`
		Version      *int64     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.ExperimentId == nil {
		return fmt.Errorf("required field experiment_id missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "experiment_id", "version"})
	} else {
		return err
	}
	o.Action = *all.Action
	o.ExperimentId = *all.ExperimentId
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
