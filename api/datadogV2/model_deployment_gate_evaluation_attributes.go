// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateEvaluationAttributes Attributes of a deployment gate evaluation.
type DeploymentGateEvaluationAttributes struct {
	// Whether this evaluation used gate-level dry run.
	DryRun bool `json:"dry_run"`
	// Evaluation duration in seconds. Null while it is in progress.
	DurationSeconds datadog.NullableInt64 `json:"duration_seconds"`
	// Deployment environment evaluated by the gate.
	Env string `json:"env"`
	// Gate evaluation UUID. Matches the resource `id`.
	EvaluationId uuid.UUID `json:"evaluation_id"`
	// Time the evaluation finished. Null while it is in progress.
	FinishedAt datadog.NullableTime `json:"finished_at"`
	// Configured deployment gate UUID. Null for just-in-time evaluations.
	GateId datadog.NullableUUID `json:"gate_id"`
	// Deployment gate identifier.
	Identifier string `json:"identifier"`
	// Service evaluated by the deployment gate.
	Service string `json:"service"`
	// Time the evaluation started.
	StartedAt time.Time `json:"started_at"`
	// The recorded result of a gate or rule evaluation.
	// - `in_progress`: The evaluation is still running.
	// - `pass`: All rules passed successfully.
	// - `fail`: One or more rules did not pass.
	Status DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status"`
	// Evaluated deployment version. Empty when no version was provided.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateEvaluationAttributes instantiates a new DeploymentGateEvaluationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateEvaluationAttributes(dryRun bool, durationSeconds datadog.NullableInt64, env string, evaluationId uuid.UUID, finishedAt datadog.NullableTime, gateId datadog.NullableUUID, identifier string, service string, startedAt time.Time, status DeploymentGatesEvaluationResultResponseAttributesGateStatus, version string) *DeploymentGateEvaluationAttributes {
	this := DeploymentGateEvaluationAttributes{}
	this.DryRun = dryRun
	this.DurationSeconds = durationSeconds
	this.Env = env
	this.EvaluationId = evaluationId
	this.FinishedAt = finishedAt
	this.GateId = gateId
	this.Identifier = identifier
	this.Service = service
	this.StartedAt = startedAt
	this.Status = status
	this.Version = version
	return &this
}

// NewDeploymentGateEvaluationAttributesWithDefaults instantiates a new DeploymentGateEvaluationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateEvaluationAttributesWithDefaults() *DeploymentGateEvaluationAttributes {
	this := DeploymentGateEvaluationAttributes{}
	return &this
}

// GetDryRun returns the DryRun field value.
func (o *DeploymentGateEvaluationAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *DeploymentGateEvaluationAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetDurationSeconds returns the DurationSeconds field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DeploymentGateEvaluationAttributes) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateEvaluationAttributes) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// SetDurationSeconds sets field value.
func (o *DeploymentGateEvaluationAttributes) SetDurationSeconds(v int64) {
	o.DurationSeconds.Set(&v)
}

// GetEnv returns the Env field value.
func (o *DeploymentGateEvaluationAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *DeploymentGateEvaluationAttributes) SetEnv(v string) {
	o.Env = v
}

// GetEvaluationId returns the EvaluationId field value.
func (o *DeploymentGateEvaluationAttributes) GetEvaluationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetEvaluationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationId, true
}

// SetEvaluationId sets field value.
func (o *DeploymentGateEvaluationAttributes) SetEvaluationId(v uuid.UUID) {
	o.EvaluationId = v
}

// GetFinishedAt returns the FinishedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DeploymentGateEvaluationAttributes) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateEvaluationAttributes) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// SetFinishedAt sets field value.
func (o *DeploymentGateEvaluationAttributes) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}

// GetGateId returns the GateId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *DeploymentGateEvaluationAttributes) GetGateId() uuid.UUID {
	if o == nil || o.GateId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.GateId.Get()
}

// GetGateIdOk returns a tuple with the GateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateEvaluationAttributes) GetGateIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.GateId.Get(), o.GateId.IsSet()
}

// SetGateId sets field value.
func (o *DeploymentGateEvaluationAttributes) SetGateId(v uuid.UUID) {
	o.GateId.Set(&v)
}

// GetIdentifier returns the Identifier field value.
func (o *DeploymentGateEvaluationAttributes) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *DeploymentGateEvaluationAttributes) SetIdentifier(v string) {
	o.Identifier = v
}

// GetService returns the Service field value.
func (o *DeploymentGateEvaluationAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DeploymentGateEvaluationAttributes) SetService(v string) {
	o.Service = v
}

// GetStartedAt returns the StartedAt field value.
func (o *DeploymentGateEvaluationAttributes) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DeploymentGateEvaluationAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *DeploymentGateEvaluationAttributes) GetStatus() DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	if o == nil {
		var ret DeploymentGatesEvaluationResultResponseAttributesGateStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetStatusOk() (*DeploymentGatesEvaluationResultResponseAttributesGateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DeploymentGateEvaluationAttributes) SetStatus(v DeploymentGatesEvaluationResultResponseAttributesGateStatus) {
	o.Status = v
}

// GetVersion returns the Version field value.
func (o *DeploymentGateEvaluationAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *DeploymentGateEvaluationAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateEvaluationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dry_run"] = o.DryRun
	toSerialize["duration_seconds"] = o.DurationSeconds.Get()
	toSerialize["env"] = o.Env
	toSerialize["evaluation_id"] = o.EvaluationId
	toSerialize["finished_at"] = o.FinishedAt.Get()
	toSerialize["gate_id"] = o.GateId.Get()
	toSerialize["identifier"] = o.Identifier
	toSerialize["service"] = o.Service
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateEvaluationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun          *bool                                                        `json:"dry_run"`
		DurationSeconds datadog.NullableInt64                                        `json:"duration_seconds"`
		Env             *string                                                      `json:"env"`
		EvaluationId    *uuid.UUID                                                   `json:"evaluation_id"`
		FinishedAt      datadog.NullableTime                                         `json:"finished_at"`
		GateId          datadog.NullableUUID                                         `json:"gate_id"`
		Identifier      *string                                                      `json:"identifier"`
		Service         *string                                                      `json:"service"`
		StartedAt       *time.Time                                                   `json:"started_at"`
		Status          *DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status"`
		Version         *string                                                      `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DryRun == nil {
		return fmt.Errorf("required field dry_run missing")
	}
	if !all.DurationSeconds.IsSet() {
		return fmt.Errorf("required field duration_seconds missing")
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.EvaluationId == nil {
		return fmt.Errorf("required field evaluation_id missing")
	}
	if !all.FinishedAt.IsSet() {
		return fmt.Errorf("required field finished_at missing")
	}
	if !all.GateId.IsSet() {
		return fmt.Errorf("required field gate_id missing")
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "duration_seconds", "env", "evaluation_id", "finished_at", "gate_id", "identifier", "service", "started_at", "status", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DryRun = *all.DryRun
	o.DurationSeconds = all.DurationSeconds
	o.Env = *all.Env
	o.EvaluationId = *all.EvaluationId
	o.FinishedAt = all.FinishedAt
	o.GateId = all.GateId
	o.Identifier = *all.Identifier
	o.Service = *all.Service
	o.StartedAt = *all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
