// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionStep A single step in a workflow execution.
type ExecutionStep struct {
	// Timestamp when the step completed. Null if not yet completed.
	CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
	// Error message if the step failed.
	Error *string `json:"error,omitempty"`
	// The unique identifier of the execution step.
	Id string `json:"id"`
	// The name of the step.
	Name string `json:"name"`
	// Timestamp when the step started. Null if not yet started.
	StartedAt datadog.NullableTime `json:"started_at,omitempty"`
	// The current status of the step.
	Status ExecutionStepStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionStep instantiates a new ExecutionStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionStep(id string, name string, status ExecutionStepStatus) *ExecutionStep {
	this := ExecutionStep{}
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewExecutionStepWithDefaults instantiates a new ExecutionStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionStepWithDefaults() *ExecutionStep {
	this := ExecutionStep{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionStep) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExecutionStep) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ExecutionStep) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given datadog.NullableTime and assigns it to the CompletedAt field.
func (o *ExecutionStep) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *ExecutionStep) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *ExecutionStep) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ExecutionStep) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStep) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ExecutionStep) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ExecutionStep) SetError(v string) {
	o.Error = &v
}

// GetId returns the Id field value.
func (o *ExecutionStep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExecutionStep) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExecutionStep) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ExecutionStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExecutionStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExecutionStep) SetName(v string) {
	o.Name = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionStep) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExecutionStep) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ExecutionStep) HasStartedAt() bool {
	return o != nil && o.StartedAt.IsSet()
}

// SetStartedAt gets a reference to the given datadog.NullableTime and assigns it to the StartedAt field.
func (o *ExecutionStep) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil.
func (o *ExecutionStep) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil.
func (o *ExecutionStep) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value.
func (o *ExecutionStep) GetStatus() ExecutionStepStatus {
	if o == nil {
		var ret ExecutionStepStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExecutionStep) GetStatusOk() (*ExecutionStepStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExecutionStep) SetStatus(v ExecutionStepStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
		Error       *string              `json:"error,omitempty"`
		Id          *string              `json:"id"`
		Name        *string              `json:"name"`
		StartedAt   datadog.NullableTime `json:"started_at,omitempty"`
		Status      *ExecutionStepStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "error", "id", "name", "started_at", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompletedAt = all.CompletedAt
	o.Error = all.Error
	o.Id = *all.Id
	o.Name = *all.Name
	o.StartedAt = all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
