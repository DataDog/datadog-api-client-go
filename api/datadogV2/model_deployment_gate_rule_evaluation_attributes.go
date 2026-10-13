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

// DeploymentGateRuleEvaluationAttributes Attributes of a deployment gate rule evaluation.
type DeploymentGateRuleEvaluationAttributes struct {
	// Evaluated rule configuration. Fields depend on rule type and unset fields are omitted.
	// Monitor rules can include `duration`, `query`, `monitor_ids`, `warmup`, `fail_on_no_groups_found`, and `fail_on_no_data`.
	// Faulty deployment detection rules can include `duration`, `allowed_resources`, and `excluded_resources`.
	Configuration DeploymentGateRuleEvaluationConfiguration `json:"configuration"`
	// Whether this rule is non-enforcing. A failed dry-run rule is ignored when computing the gate outcome. Independent of `gate_dry_run`.
	DryRun bool `json:"dry_run"`
	// Rule evaluation duration in seconds. Null while it is in progress.
	DurationSeconds datadog.NullableInt64 `json:"duration_seconds"`
	// Evaluated environment.
	Env string `json:"env"`
	// Rule evaluation UUID. Matches the resource `id`.
	EvaluationId uuid.UUID `json:"evaluation_id"`
	// Rule failure details.
	Failures DeploymentGateRuleFailures `json:"failures"`
	// Time the rule evaluation finished. Null while it is in progress.
	FinishedAt datadog.NullableTime `json:"finished_at"`
	// Whether the parent gate is dry-run. A failed dry-run gate blocks but does not stop deployment. Independent of rule-level `dry_run`.
	GateDryRun bool `json:"gate_dry_run"`
	// Deployment gate evaluation UUID.
	GateEvaluationId uuid.UUID `json:"gate_evaluation_id"`
	// Configured deployment gate UUID. Null for just-in-time evaluations.
	GateId datadog.NullableUUID `json:"gate_id"`
	// Deployment gate identifier.
	Identifier string `json:"identifier"`
	// Rule name.
	Name string `json:"name"`
	// Reason for the rule result.
	Reason string `json:"reason"`
	// Configured deployment rule UUID. Null for just-in-time rules.
	RuleId datadog.NullableUUID `json:"rule_id"`
	// Evaluated service.
	Service string `json:"service"`
	// Time the rule evaluation started.
	StartedAt time.Time `json:"started_at"`
	// The recorded result of a gate or rule evaluation.
	// - `in_progress`: The evaluation is still running.
	// - `pass`: All rules passed successfully.
	// - `fail`: One or more rules did not pass.
	Status DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status"`
	// Type of deployment gate rule.
	Type DeploymentGateRuleEvaluationType `json:"type"`
	// Evaluated deployment version. Empty when no version was provided.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateRuleEvaluationAttributes instantiates a new DeploymentGateRuleEvaluationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateRuleEvaluationAttributes(configuration DeploymentGateRuleEvaluationConfiguration, dryRun bool, durationSeconds datadog.NullableInt64, env string, evaluationId uuid.UUID, failures DeploymentGateRuleFailures, finishedAt datadog.NullableTime, gateDryRun bool, gateEvaluationId uuid.UUID, gateId datadog.NullableUUID, identifier string, name string, reason string, ruleId datadog.NullableUUID, service string, startedAt time.Time, status DeploymentGatesEvaluationResultResponseAttributesGateStatus, typeVar DeploymentGateRuleEvaluationType, version string) *DeploymentGateRuleEvaluationAttributes {
	this := DeploymentGateRuleEvaluationAttributes{}
	this.Configuration = configuration
	this.DryRun = dryRun
	this.DurationSeconds = durationSeconds
	this.Env = env
	this.EvaluationId = evaluationId
	this.Failures = failures
	this.FinishedAt = finishedAt
	this.GateDryRun = gateDryRun
	this.GateEvaluationId = gateEvaluationId
	this.GateId = gateId
	this.Identifier = identifier
	this.Name = name
	this.Reason = reason
	this.RuleId = ruleId
	this.Service = service
	this.StartedAt = startedAt
	this.Status = status
	this.Type = typeVar
	this.Version = version
	return &this
}

// NewDeploymentGateRuleEvaluationAttributesWithDefaults instantiates a new DeploymentGateRuleEvaluationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateRuleEvaluationAttributesWithDefaults() *DeploymentGateRuleEvaluationAttributes {
	this := DeploymentGateRuleEvaluationAttributes{}
	return &this
}

// GetConfiguration returns the Configuration field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetConfiguration() DeploymentGateRuleEvaluationConfiguration {
	if o == nil {
		var ret DeploymentGateRuleEvaluationConfiguration
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetConfigurationOk() (*DeploymentGateRuleEvaluationConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetConfiguration(v DeploymentGateRuleEvaluationConfiguration) {
	o.Configuration = v
}

// GetDryRun returns the DryRun field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetDurationSeconds returns the DurationSeconds field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// SetDurationSeconds sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetDurationSeconds(v int64) {
	o.DurationSeconds.Set(&v)
}

// GetEnv returns the Env field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetEnv(v string) {
	o.Env = v
}

// GetEvaluationId returns the EvaluationId field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetEvaluationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetEvaluationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationId, true
}

// SetEvaluationId sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetEvaluationId(v uuid.UUID) {
	o.EvaluationId = v
}

// GetFailures returns the Failures field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetFailures() DeploymentGateRuleFailures {
	if o == nil {
		var ret DeploymentGateRuleFailures
		return ret
	}
	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetFailuresOk() (*DeploymentGateRuleFailures, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failures, true
}

// SetFailures sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetFailures(v DeploymentGateRuleFailures) {
	o.Failures = v
}

// GetFinishedAt returns the FinishedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// SetFinishedAt sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}

// GetGateDryRun returns the GateDryRun field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.GateDryRun
}

// GetGateDryRunOk returns a tuple with the GateDryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GateDryRun, true
}

// SetGateDryRun sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetGateDryRun(v bool) {
	o.GateDryRun = v
}

// GetGateEvaluationId returns the GateEvaluationId field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateEvaluationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.GateEvaluationId
}

// GetGateEvaluationIdOk returns a tuple with the GateEvaluationId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateEvaluationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GateEvaluationId, true
}

// SetGateEvaluationId sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetGateEvaluationId(v uuid.UUID) {
	o.GateEvaluationId = v
}

// GetGateId returns the GateId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateId() uuid.UUID {
	if o == nil || o.GateId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.GateId.Get()
}

// GetGateIdOk returns a tuple with the GateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetGateIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.GateId.Get(), o.GateId.IsSet()
}

// SetGateId sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetGateId(v uuid.UUID) {
	o.GateId.Set(&v)
}

// GetIdentifier returns the Identifier field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetReason(v string) {
	o.Reason = v
}

// GetRuleId returns the RuleId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetRuleId() uuid.UUID {
	if o == nil || o.RuleId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.RuleId.Get()
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DeploymentGateRuleEvaluationAttributes) GetRuleIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleId.Get(), o.RuleId.IsSet()
}

// SetRuleId sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetRuleId(v uuid.UUID) {
	o.RuleId.Set(&v)
}

// GetService returns the Service field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetService(v string) {
	o.Service = v
}

// GetStartedAt returns the StartedAt field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetStatus() DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	if o == nil {
		var ret DeploymentGatesEvaluationResultResponseAttributesGateStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetStatusOk() (*DeploymentGatesEvaluationResultResponseAttributesGateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetStatus(v DeploymentGatesEvaluationResultResponseAttributesGateStatus) {
	o.Status = v
}

// GetType returns the Type field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetType() DeploymentGateRuleEvaluationType {
	if o == nil {
		var ret DeploymentGateRuleEvaluationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetTypeOk() (*DeploymentGateRuleEvaluationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetType(v DeploymentGateRuleEvaluationType) {
	o.Type = v
}

// GetVersion returns the Version field value.
func (o *DeploymentGateRuleEvaluationAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *DeploymentGateRuleEvaluationAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateRuleEvaluationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration"] = o.Configuration
	toSerialize["dry_run"] = o.DryRun
	toSerialize["duration_seconds"] = o.DurationSeconds.Get()
	toSerialize["env"] = o.Env
	toSerialize["evaluation_id"] = o.EvaluationId
	toSerialize["failures"] = o.Failures
	toSerialize["finished_at"] = o.FinishedAt.Get()
	toSerialize["gate_dry_run"] = o.GateDryRun
	toSerialize["gate_evaluation_id"] = o.GateEvaluationId
	toSerialize["gate_id"] = o.GateId.Get()
	toSerialize["identifier"] = o.Identifier
	toSerialize["name"] = o.Name
	toSerialize["reason"] = o.Reason
	toSerialize["rule_id"] = o.RuleId.Get()
	toSerialize["service"] = o.Service
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateRuleEvaluationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configuration    *DeploymentGateRuleEvaluationConfiguration                   `json:"configuration"`
		DryRun           *bool                                                        `json:"dry_run"`
		DurationSeconds  datadog.NullableInt64                                        `json:"duration_seconds"`
		Env              *string                                                      `json:"env"`
		EvaluationId     *uuid.UUID                                                   `json:"evaluation_id"`
		Failures         *DeploymentGateRuleFailures                                  `json:"failures"`
		FinishedAt       datadog.NullableTime                                         `json:"finished_at"`
		GateDryRun       *bool                                                        `json:"gate_dry_run"`
		GateEvaluationId *uuid.UUID                                                   `json:"gate_evaluation_id"`
		GateId           datadog.NullableUUID                                         `json:"gate_id"`
		Identifier       *string                                                      `json:"identifier"`
		Name             *string                                                      `json:"name"`
		Reason           *string                                                      `json:"reason"`
		RuleId           datadog.NullableUUID                                         `json:"rule_id"`
		Service          *string                                                      `json:"service"`
		StartedAt        *time.Time                                                   `json:"started_at"`
		Status           *DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status"`
		Type             *DeploymentGateRuleEvaluationType                            `json:"type"`
		Version          *string                                                      `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Configuration == nil {
		return fmt.Errorf("required field configuration missing")
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
	if all.Failures == nil {
		return fmt.Errorf("required field failures missing")
	}
	if !all.FinishedAt.IsSet() {
		return fmt.Errorf("required field finished_at missing")
	}
	if all.GateDryRun == nil {
		return fmt.Errorf("required field gate_dry_run missing")
	}
	if all.GateEvaluationId == nil {
		return fmt.Errorf("required field gate_evaluation_id missing")
	}
	if !all.GateId.IsSet() {
		return fmt.Errorf("required field gate_id missing")
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if !all.RuleId.IsSet() {
		return fmt.Errorf("required field rule_id missing")
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration", "dry_run", "duration_seconds", "env", "evaluation_id", "failures", "finished_at", "gate_dry_run", "gate_evaluation_id", "gate_id", "identifier", "name", "reason", "rule_id", "service", "started_at", "status", "type", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Configuration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Configuration = *all.Configuration
	o.DryRun = *all.DryRun
	o.DurationSeconds = all.DurationSeconds
	o.Env = *all.Env
	o.EvaluationId = *all.EvaluationId
	if all.Failures.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Failures = *all.Failures
	o.FinishedAt = all.FinishedAt
	o.GateDryRun = *all.GateDryRun
	o.GateEvaluationId = *all.GateEvaluationId
	o.GateId = all.GateId
	o.Identifier = *all.Identifier
	o.Name = *all.Name
	o.Reason = *all.Reason
	o.RuleId = all.RuleId
	o.Service = *all.Service
	o.StartedAt = *all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
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
