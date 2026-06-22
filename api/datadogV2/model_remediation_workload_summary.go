// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationWorkloadSummary Aggregated health across all tasks in a workload (service or daemon).
type RemediationWorkloadSummary struct {
	// ECS deployment state, populated only for deployment failures.
	DeploymentRolloutState *RemediationDeploymentRolloutState `json:"deployment_rollout_state,omitempty"`
	// Expected task count (64-bit integer encoded as a string).
	DesiredCount *string `json:"desired_count,omitempty"`
	// Tasks in a problematic state (64-bit integer encoded as a string).
	FailedCount *string `json:"failed_count,omitempty"`
	// Percentage of desired count that is failed (64-bit integer encoded as a string).
	FailedPercent *string `json:"failed_percent,omitempty"`
	// Tasks currently pending (64-bit integer encoded as a string).
	PendingCount *string `json:"pending_count,omitempty"`
	// Previous deployment's task definition, as family:revision or a full task definition ARN. Empty when no rollback target is visible.
	PreviousTaskDefinition *string `json:"previous_task_definition,omitempty"`
	// Tasks currently running (64-bit integer encoded as a string).
	RunningCount *string `json:"running_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationWorkloadSummary instantiates a new RemediationWorkloadSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationWorkloadSummary() *RemediationWorkloadSummary {
	this := RemediationWorkloadSummary{}
	return &this
}

// NewRemediationWorkloadSummaryWithDefaults instantiates a new RemediationWorkloadSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationWorkloadSummaryWithDefaults() *RemediationWorkloadSummary {
	this := RemediationWorkloadSummary{}
	return &this
}

// GetDeploymentRolloutState returns the DeploymentRolloutState field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetDeploymentRolloutState() RemediationDeploymentRolloutState {
	if o == nil || o.DeploymentRolloutState == nil {
		var ret RemediationDeploymentRolloutState
		return ret
	}
	return *o.DeploymentRolloutState
}

// GetDeploymentRolloutStateOk returns a tuple with the DeploymentRolloutState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetDeploymentRolloutStateOk() (*RemediationDeploymentRolloutState, bool) {
	if o == nil || o.DeploymentRolloutState == nil {
		return nil, false
	}
	return o.DeploymentRolloutState, true
}

// HasDeploymentRolloutState returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasDeploymentRolloutState() bool {
	return o != nil && o.DeploymentRolloutState != nil
}

// SetDeploymentRolloutState gets a reference to the given RemediationDeploymentRolloutState and assigns it to the DeploymentRolloutState field.
func (o *RemediationWorkloadSummary) SetDeploymentRolloutState(v RemediationDeploymentRolloutState) {
	o.DeploymentRolloutState = &v
}

// GetDesiredCount returns the DesiredCount field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetDesiredCount() string {
	if o == nil || o.DesiredCount == nil {
		var ret string
		return ret
	}
	return *o.DesiredCount
}

// GetDesiredCountOk returns a tuple with the DesiredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetDesiredCountOk() (*string, bool) {
	if o == nil || o.DesiredCount == nil {
		return nil, false
	}
	return o.DesiredCount, true
}

// HasDesiredCount returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasDesiredCount() bool {
	return o != nil && o.DesiredCount != nil
}

// SetDesiredCount gets a reference to the given string and assigns it to the DesiredCount field.
func (o *RemediationWorkloadSummary) SetDesiredCount(v string) {
	o.DesiredCount = &v
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetFailedCount() string {
	if o == nil || o.FailedCount == nil {
		var ret string
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetFailedCountOk() (*string, bool) {
	if o == nil || o.FailedCount == nil {
		return nil, false
	}
	return o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasFailedCount() bool {
	return o != nil && o.FailedCount != nil
}

// SetFailedCount gets a reference to the given string and assigns it to the FailedCount field.
func (o *RemediationWorkloadSummary) SetFailedCount(v string) {
	o.FailedCount = &v
}

// GetFailedPercent returns the FailedPercent field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetFailedPercent() string {
	if o == nil || o.FailedPercent == nil {
		var ret string
		return ret
	}
	return *o.FailedPercent
}

// GetFailedPercentOk returns a tuple with the FailedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetFailedPercentOk() (*string, bool) {
	if o == nil || o.FailedPercent == nil {
		return nil, false
	}
	return o.FailedPercent, true
}

// HasFailedPercent returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasFailedPercent() bool {
	return o != nil && o.FailedPercent != nil
}

// SetFailedPercent gets a reference to the given string and assigns it to the FailedPercent field.
func (o *RemediationWorkloadSummary) SetFailedPercent(v string) {
	o.FailedPercent = &v
}

// GetPendingCount returns the PendingCount field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetPendingCount() string {
	if o == nil || o.PendingCount == nil {
		var ret string
		return ret
	}
	return *o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetPendingCountOk() (*string, bool) {
	if o == nil || o.PendingCount == nil {
		return nil, false
	}
	return o.PendingCount, true
}

// HasPendingCount returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasPendingCount() bool {
	return o != nil && o.PendingCount != nil
}

// SetPendingCount gets a reference to the given string and assigns it to the PendingCount field.
func (o *RemediationWorkloadSummary) SetPendingCount(v string) {
	o.PendingCount = &v
}

// GetPreviousTaskDefinition returns the PreviousTaskDefinition field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetPreviousTaskDefinition() string {
	if o == nil || o.PreviousTaskDefinition == nil {
		var ret string
		return ret
	}
	return *o.PreviousTaskDefinition
}

// GetPreviousTaskDefinitionOk returns a tuple with the PreviousTaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetPreviousTaskDefinitionOk() (*string, bool) {
	if o == nil || o.PreviousTaskDefinition == nil {
		return nil, false
	}
	return o.PreviousTaskDefinition, true
}

// HasPreviousTaskDefinition returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasPreviousTaskDefinition() bool {
	return o != nil && o.PreviousTaskDefinition != nil
}

// SetPreviousTaskDefinition gets a reference to the given string and assigns it to the PreviousTaskDefinition field.
func (o *RemediationWorkloadSummary) SetPreviousTaskDefinition(v string) {
	o.PreviousTaskDefinition = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *RemediationWorkloadSummary) GetRunningCount() string {
	if o == nil || o.RunningCount == nil {
		var ret string
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationWorkloadSummary) GetRunningCountOk() (*string, bool) {
	if o == nil || o.RunningCount == nil {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *RemediationWorkloadSummary) HasRunningCount() bool {
	return o != nil && o.RunningCount != nil
}

// SetRunningCount gets a reference to the given string and assigns it to the RunningCount field.
func (o *RemediationWorkloadSummary) SetRunningCount(v string) {
	o.RunningCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationWorkloadSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeploymentRolloutState != nil {
		toSerialize["deployment_rollout_state"] = o.DeploymentRolloutState
	}
	if o.DesiredCount != nil {
		toSerialize["desired_count"] = o.DesiredCount
	}
	if o.FailedCount != nil {
		toSerialize["failed_count"] = o.FailedCount
	}
	if o.FailedPercent != nil {
		toSerialize["failed_percent"] = o.FailedPercent
	}
	if o.PendingCount != nil {
		toSerialize["pending_count"] = o.PendingCount
	}
	if o.PreviousTaskDefinition != nil {
		toSerialize["previous_task_definition"] = o.PreviousTaskDefinition
	}
	if o.RunningCount != nil {
		toSerialize["running_count"] = o.RunningCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationWorkloadSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentRolloutState *RemediationDeploymentRolloutState `json:"deployment_rollout_state,omitempty"`
		DesiredCount           *string                            `json:"desired_count,omitempty"`
		FailedCount            *string                            `json:"failed_count,omitempty"`
		FailedPercent          *string                            `json:"failed_percent,omitempty"`
		PendingCount           *string                            `json:"pending_count,omitempty"`
		PreviousTaskDefinition *string                            `json:"previous_task_definition,omitempty"`
		RunningCount           *string                            `json:"running_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deployment_rollout_state", "desired_count", "failed_count", "failed_percent", "pending_count", "previous_task_definition", "running_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DeploymentRolloutState != nil && !all.DeploymentRolloutState.IsValid() {
		hasInvalidField = true
	} else {
		o.DeploymentRolloutState = all.DeploymentRolloutState
	}
	o.DesiredCount = all.DesiredCount
	o.FailedCount = all.FailedCount
	o.FailedPercent = all.FailedPercent
	o.PendingCount = all.PendingCount
	o.PreviousTaskDefinition = all.PreviousTaskDefinition
	o.RunningCount = all.RunningCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
