// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems Refresh outcome for one experiment.
type ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems struct {
	// ID of the experiment associated with this result.
	ExperimentId *string `json:"experiment_id,omitempty"`
	// Outcome of attempting to refresh one experiment.
	Outcome *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome `json:"outcome,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems instantiates a new ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems() *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems {
	this := ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems{}
	return &this
}

// NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsWithDefaults instantiates a new ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsWithDefaults() *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems {
	this := ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems{}
	return &this
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) GetExperimentId() string {
	if o == nil || o.ExperimentId == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) GetExperimentIdOk() (*string, bool) {
	if o == nil || o.ExperimentId == nil {
		return nil, false
	}
	return o.ExperimentId, true
}

// HasExperimentId returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) HasExperimentId() bool {
	return o != nil && o.ExperimentId != nil
}

// SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) SetExperimentId(v string) {
	o.ExperimentId = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) GetOutcome() ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome {
	if o == nil || o.Outcome == nil {
		var ret ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) GetOutcomeOk() (*ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) HasOutcome() bool {
	return o != nil && o.Outcome != nil
}

// SetOutcome gets a reference to the given ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome and assigns it to the Outcome field.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) SetOutcome(v ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome) {
	o.Outcome = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExperimentId != nil {
		toSerialize["experiment_id"] = o.ExperimentId
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExperimentId *string                                                               `json:"experiment_id,omitempty"`
		Outcome      *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItemsOutcome `json:"outcome,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"experiment_id", "outcome"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExperimentId = all.ExperimentId
	if all.Outcome != nil && !all.Outcome.IsValid() {
		hasInvalidField = true
	} else {
		o.Outcome = all.Outcome
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems handles when a null is used for ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems.
type NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems struct {
	value *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) Get() *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) Set(val *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems initializes the struct as if Set has been called.
func NewNullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems(val *ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) *NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems {
	return &NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
