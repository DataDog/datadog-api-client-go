// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributes Diagnostic check results and their evaluation state.
type ExperimentsExperimentDiagnosticsV2DTODataAttributes struct {
	// Results of individual diagnostic checks.
	Diagnostics []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems `json:"diagnostics"`
	// Time when the diagnostic checks were evaluated.
	EvaluatedAt datadog.NullableTime `json:"evaluated_at,omitempty"`
	// Overall result of the experiment diagnostic checks.
	Result NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult `json:"result,omitempty"`
	// Current state of the diagnostic evaluation.
	State ExperimentsExperimentDiagnosticsV2DTODataAttributesState `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributes instantiates a new ExperimentsExperimentDiagnosticsV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributes(diagnostics []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems, state ExperimentsExperimentDiagnosticsV2DTODataAttributesState) *ExperimentsExperimentDiagnosticsV2DTODataAttributes {
	this := ExperimentsExperimentDiagnosticsV2DTODataAttributes{}
	this.Diagnostics = diagnostics
	this.State = state
	return &this
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesWithDefaults instantiates a new ExperimentsExperimentDiagnosticsV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesWithDefaults() *ExperimentsExperimentDiagnosticsV2DTODataAttributes {
	this := ExperimentsExperimentDiagnosticsV2DTODataAttributes{}
	return &this
}

// GetDiagnostics returns the Diagnostics field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetDiagnostics() []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems {
	if o == nil {
		var ret []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems
		return ret
	}
	return o.Diagnostics
}

// GetDiagnosticsOk returns a tuple with the Diagnostics field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetDiagnosticsOk() (*[]ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diagnostics, true
}

// SetDiagnostics sets field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetDiagnostics(v []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) {
	o.Diagnostics = v
}

// GetEvaluatedAt returns the EvaluatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetEvaluatedAt() time.Time {
	if o == nil || o.EvaluatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EvaluatedAt.Get()
}

// GetEvaluatedAtOk returns a tuple with the EvaluatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetEvaluatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvaluatedAt.Get(), o.EvaluatedAt.IsSet()
}

// HasEvaluatedAt returns a boolean if a field has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) HasEvaluatedAt() bool {
	return o != nil && o.EvaluatedAt.IsSet()
}

// SetEvaluatedAt gets a reference to the given datadog.NullableTime and assigns it to the EvaluatedAt field.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetEvaluatedAt(v time.Time) {
	o.EvaluatedAt.Set(&v)
}

// SetEvaluatedAtNil sets the value for EvaluatedAt to be an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetEvaluatedAtNil() {
	o.EvaluatedAt.Set(nil)
}

// UnsetEvaluatedAt ensures that no value is present for EvaluatedAt, not even an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) UnsetEvaluatedAt() {
	o.EvaluatedAt.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetResult() ExperimentsExperimentDiagnosticsV2DTODataAttributesResult {
	if o == nil || o.Result.Get() == nil {
		var ret ExperimentsExperimentDiagnosticsV2DTODataAttributesResult
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetResultOk() (*ExperimentsExperimentDiagnosticsV2DTODataAttributesResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) HasResult() bool {
	return o != nil && o.Result.IsSet()
}

// SetResult gets a reference to the given NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult and assigns it to the Result field.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetResult(v ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) {
	o.Result.Set(&v)
}

// SetResultNil sets the value for Result to be an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) UnsetResult() {
	o.Result.Unset()
}

// GetState returns the State field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetState() ExperimentsExperimentDiagnosticsV2DTODataAttributesState {
	if o == nil {
		var ret ExperimentsExperimentDiagnosticsV2DTODataAttributesState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) GetStateOk() (*ExperimentsExperimentDiagnosticsV2DTODataAttributesState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) SetState(v ExperimentsExperimentDiagnosticsV2DTODataAttributesState) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentDiagnosticsV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["diagnostics"] = o.Diagnostics
	if o.EvaluatedAt.IsSet() {
		toSerialize["evaluated_at"] = o.EvaluatedAt.Get()
	}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Diagnostics *[]ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems `json:"diagnostics"`
		EvaluatedAt datadog.NullableTime                                                   `json:"evaluated_at,omitempty"`
		Result      NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult      `json:"result,omitempty"`
		State       *ExperimentsExperimentDiagnosticsV2DTODataAttributesState              `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Diagnostics == nil {
		return fmt.Errorf("required field diagnostics missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"diagnostics", "evaluated_at", "result", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Diagnostics = *all.Diagnostics
	o.EvaluatedAt = all.EvaluatedAt
	if all.Result.Get() != nil && !all.Result.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Result = all.Result
	}
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
