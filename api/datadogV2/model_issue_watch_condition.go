// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchCondition Condition that triggers the watch action. The fields that are set depend on `kind`:
// `OCCURRENCES` uses `threshold`, `CARDINALITY` uses `threshold` and `projection`,
// `ELAPSED_TIME` uses `duration`, and `ANY` and `ALL` use `operands`.
type IssueWatchCondition struct {
	// Time that must elapse before the condition is met, as a Go duration string.
	Duration *string `json:"duration,omitempty"`
	// Kind of watch condition.
	Kind IssueWatchConditionKind `json:"kind"`
	// Nested conditions combined by `ANY` or `ALL`.
	Operands []IssueWatchCondition `json:"operands,omitempty"`
	// Attributes whose distinct values are counted for `CARDINALITY`.
	Projection []string `json:"projection,omitempty"`
	// Number of occurrences, or of distinct values for `CARDINALITY`, that triggers the condition.
	Threshold *int64 `json:"threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueWatchCondition instantiates a new IssueWatchCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueWatchCondition(kind IssueWatchConditionKind) *IssueWatchCondition {
	this := IssueWatchCondition{}
	this.Kind = kind
	return &this
}

// NewIssueWatchConditionWithDefaults instantiates a new IssueWatchCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueWatchConditionWithDefaults() *IssueWatchCondition {
	this := IssueWatchCondition{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IssueWatchCondition) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchCondition) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IssueWatchCondition) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *IssueWatchCondition) SetDuration(v string) {
	o.Duration = &v
}

// GetKind returns the Kind field value.
func (o *IssueWatchCondition) GetKind() IssueWatchConditionKind {
	if o == nil {
		var ret IssueWatchConditionKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *IssueWatchCondition) GetKindOk() (*IssueWatchConditionKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *IssueWatchCondition) SetKind(v IssueWatchConditionKind) {
	o.Kind = v
}

// GetOperands returns the Operands field value if set, zero value otherwise.
func (o *IssueWatchCondition) GetOperands() []IssueWatchCondition {
	if o == nil || o.Operands == nil {
		var ret []IssueWatchCondition
		return ret
	}
	return o.Operands
}

// GetOperandsOk returns a tuple with the Operands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchCondition) GetOperandsOk() (*[]IssueWatchCondition, bool) {
	if o == nil || o.Operands == nil {
		return nil, false
	}
	return &o.Operands, true
}

// HasOperands returns a boolean if a field has been set.
func (o *IssueWatchCondition) HasOperands() bool {
	return o != nil && o.Operands != nil
}

// SetOperands gets a reference to the given []IssueWatchCondition and assigns it to the Operands field.
func (o *IssueWatchCondition) SetOperands(v []IssueWatchCondition) {
	o.Operands = v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *IssueWatchCondition) GetProjection() []string {
	if o == nil || o.Projection == nil {
		var ret []string
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchCondition) GetProjectionOk() (*[]string, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return &o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *IssueWatchCondition) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given []string and assigns it to the Projection field.
func (o *IssueWatchCondition) SetProjection(v []string) {
	o.Projection = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *IssueWatchCondition) GetThreshold() int64 {
	if o == nil || o.Threshold == nil {
		var ret int64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchCondition) GetThresholdOk() (*int64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *IssueWatchCondition) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given int64 and assigns it to the Threshold field.
func (o *IssueWatchCondition) SetThreshold(v int64) {
	o.Threshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueWatchCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["kind"] = o.Kind
	if o.Operands != nil {
		toSerialize["operands"] = o.Operands
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueWatchCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration   *string                  `json:"duration,omitempty"`
		Kind       *IssueWatchConditionKind `json:"kind"`
		Operands   []IssueWatchCondition    `json:"operands,omitempty"`
		Projection []string                 `json:"projection,omitempty"`
		Threshold  *int64                   `json:"threshold,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "kind", "operands", "projection", "threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Duration = all.Duration
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	o.Operands = all.Operands
	o.Projection = all.Projection
	o.Threshold = all.Threshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
