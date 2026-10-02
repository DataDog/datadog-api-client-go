// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsSavedFilterCondition A condition that uses a saved filter. Inline fields must be omitted or null.
type ExperimentsSavedFilterCondition struct {
	// Saved-filter UUID.
	SavedFilterId uuid.UUID `json:"saved_filter_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsSavedFilterCondition instantiates a new ExperimentsSavedFilterCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsSavedFilterCondition(savedFilterId uuid.UUID) *ExperimentsSavedFilterCondition {
	this := ExperimentsSavedFilterCondition{}
	this.SavedFilterId = savedFilterId
	return &this
}

// NewExperimentsSavedFilterConditionWithDefaults instantiates a new ExperimentsSavedFilterCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsSavedFilterConditionWithDefaults() *ExperimentsSavedFilterCondition {
	this := ExperimentsSavedFilterCondition{}
	return &this
}

// GetSavedFilterId returns the SavedFilterId field value.
func (o *ExperimentsSavedFilterCondition) GetSavedFilterId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SavedFilterId
}

// GetSavedFilterIdOk returns a tuple with the SavedFilterId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsSavedFilterCondition) GetSavedFilterIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedFilterId, true
}

// SetSavedFilterId sets field value.
func (o *ExperimentsSavedFilterCondition) SetSavedFilterId(v uuid.UUID) {
	o.SavedFilterId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsSavedFilterCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["saved_filter_id"] = o.SavedFilterId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsSavedFilterCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SavedFilterId *uuid.UUID `json:"saved_filter_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SavedFilterId == nil {
		return fmt.Errorf("required field saved_filter_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"saved_filter_id"})
	} else {
		return err
	}
	o.SavedFilterId = *all.SavedFilterId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
