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

// ExperimentsBanditV2DTODataAttributes A bandit's setup, the policy it serves, and its latest refresh.
type ExperimentsBanditV2DTODataAttributes struct {
	// The policy the bandit serves. Null until the bandit starts.
	ActivePolicy NullableExperimentsBanditPolicyV2 `json:"active_policy"`
	// How long after an exposure its reward is counted, in seconds.
	AggregationWindowSeconds int64 `json:"aggregation_window_seconds"`
	// The prior the bandit's estimates start from.
	BayesianPrior ExperimentsBanditBayesianPriorV2 `json:"bayesian_prior"`
	// The version of the bandit's configuration. Each saved change adds one.
	ConfigVersion int64 `json:"config_version"`
	// When the bandit was created.
	CreatedAt time.Time `json:"created_at"`
	// The flag rule the bandit serves through. Absent until a draft is put on a flag, and when the caller lacks feature flag and environment read permissions (`meta.warnings` then says so).
	DatadogFlagConfiguration *ExperimentsBanditFlagConfigurationV2 `json:"datadog_flag_configuration,omitempty"`
	// The bandit's description.
	Description string `json:"description"`
	// When the bandit was stopped or reached its scheduled end. Null while it hasn't ended.
	EndedAt datadog.NullableTime `json:"ended_at"`
	// The exposure SQL model the bandit reads its exposures from. Null until set on a draft.
	ExposureSqlModelId datadog.NullableUUID `json:"exposure_sql_model_id"`
	// The most recently started refresh. Null until one starts.
	LatestRefresh NullableExperimentsBanditRefreshV2 `json:"latest_refresh"`
	// How far back each refresh reads observations, in seconds.
	LookbackSeconds int64 `json:"lookback_seconds"`
	// The metric the bandit optimizes. Null until set on a draft.
	MetricId datadog.NullableUUID `json:"metric_id"`
	// The bandit's name.
	Name string `json:"name"`
	// The key of the arm most likely to be best, when one arm leads and the bandit is outside warmup. Null otherwise. Matches Get bandit results.
	OptimalVariantKey datadog.NullableString `json:"optimal_variant_key"`
	// Why no optimal variant is reported. Absent when one is, or when the cause isn't known.
	OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
	// The schedule the bandit refreshes on.
	RefreshScheduleId datadog.NullableString `json:"refresh_schedule_id"`
	// When the bandit ends and serves its fallback. Null when it has no scheduled end.
	ScheduledEndAt datadog.NullableTime `json:"scheduled_end_at"`
	// Why the scheduled end couldn't run. It isn't retried until the end or fallback changes. Absent when there's no failure.
	ScheduledEndFailure *ExperimentsBanditScheduledEndFailureV2 `json:"scheduled_end_failure,omitempty"`
	// When the bandit started. Null until it starts.
	StartedAt datadog.NullableTime `json:"started_at"`
	// The bandit's status. STOPPED means the bandit was stopped and serves the variant chosen then. COMPLETED means it reached its scheduled end and serves its fallback.
	Status ExperimentsBanditStatus `json:"status"`
	// The subject type the bandit randomizes and measures by. Null until set on a draft.
	SubjectTypeId datadog.NullableUUID `json:"subject_type_id"`
	// The bandit's tags, without team tags.
	Tags []string `json:"tags"`
	// Handles of the teams that own the bandit.
	Teams []string `json:"teams"`
	// When the bandit was last changed.
	UpdatedAt time.Time `json:"updated_at"`
	// The flag variants the bandit uses: its arms in arm order, then its fallback when the fallback isn't an arm.
	Variants []ExperimentsBanditVariantV2 `json:"variants"`
	// How long after starting the bandit serves its arms equally, in seconds. Null when it has no warmup.
	WarmupSeconds datadog.NullableInt64 `json:"warmup_seconds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditV2DTODataAttributes instantiates a new ExperimentsBanditV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditV2DTODataAttributes(activePolicy NullableExperimentsBanditPolicyV2, aggregationWindowSeconds int64, bayesianPrior ExperimentsBanditBayesianPriorV2, configVersion int64, createdAt time.Time, description string, endedAt datadog.NullableTime, exposureSqlModelId datadog.NullableUUID, latestRefresh NullableExperimentsBanditRefreshV2, lookbackSeconds int64, metricId datadog.NullableUUID, name string, optimalVariantKey datadog.NullableString, refreshScheduleId datadog.NullableString, scheduledEndAt datadog.NullableTime, startedAt datadog.NullableTime, status ExperimentsBanditStatus, subjectTypeId datadog.NullableUUID, tags []string, teams []string, updatedAt time.Time, variants []ExperimentsBanditVariantV2, warmupSeconds datadog.NullableInt64) *ExperimentsBanditV2DTODataAttributes {
	this := ExperimentsBanditV2DTODataAttributes{}
	this.ActivePolicy = activePolicy
	this.AggregationWindowSeconds = aggregationWindowSeconds
	this.BayesianPrior = bayesianPrior
	this.ConfigVersion = configVersion
	this.CreatedAt = createdAt
	this.Description = description
	this.EndedAt = endedAt
	this.ExposureSqlModelId = exposureSqlModelId
	this.LatestRefresh = latestRefresh
	this.LookbackSeconds = lookbackSeconds
	this.MetricId = metricId
	this.Name = name
	this.OptimalVariantKey = optimalVariantKey
	this.RefreshScheduleId = refreshScheduleId
	this.ScheduledEndAt = scheduledEndAt
	this.StartedAt = startedAt
	this.Status = status
	this.SubjectTypeId = subjectTypeId
	this.Tags = tags
	this.Teams = teams
	this.UpdatedAt = updatedAt
	this.Variants = variants
	this.WarmupSeconds = warmupSeconds
	return &this
}

// NewExperimentsBanditV2DTODataAttributesWithDefaults instantiates a new ExperimentsBanditV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditV2DTODataAttributesWithDefaults() *ExperimentsBanditV2DTODataAttributes {
	this := ExperimentsBanditV2DTODataAttributes{}
	return &this
}

// GetActivePolicy returns the ActivePolicy field value.
// If the value is explicit nil, the zero value for ExperimentsBanditPolicyV2 will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetActivePolicy() ExperimentsBanditPolicyV2 {
	if o == nil || o.ActivePolicy.Get() == nil {
		var ret ExperimentsBanditPolicyV2
		return ret
	}
	return *o.ActivePolicy.Get()
}

// GetActivePolicyOk returns a tuple with the ActivePolicy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetActivePolicyOk() (*ExperimentsBanditPolicyV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivePolicy.Get(), o.ActivePolicy.IsSet()
}

// SetActivePolicy sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetActivePolicy(v ExperimentsBanditPolicyV2) {
	o.ActivePolicy.Set(&v)
}

// GetAggregationWindowSeconds returns the AggregationWindowSeconds field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetAggregationWindowSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AggregationWindowSeconds
}

// GetAggregationWindowSecondsOk returns a tuple with the AggregationWindowSeconds field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetAggregationWindowSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationWindowSeconds, true
}

// SetAggregationWindowSeconds sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetAggregationWindowSeconds(v int64) {
	o.AggregationWindowSeconds = v
}

// GetBayesianPrior returns the BayesianPrior field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetBayesianPrior() ExperimentsBanditBayesianPriorV2 {
	if o == nil {
		var ret ExperimentsBanditBayesianPriorV2
		return ret
	}
	return o.BayesianPrior
}

// GetBayesianPriorOk returns a tuple with the BayesianPrior field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetBayesianPriorOk() (*ExperimentsBanditBayesianPriorV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BayesianPrior, true
}

// SetBayesianPrior sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetBayesianPrior(v ExperimentsBanditBayesianPriorV2) {
	o.BayesianPrior = v
}

// GetConfigVersion returns the ConfigVersion field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetConfigVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigVersion, true
}

// SetConfigVersion sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDatadogFlagConfiguration returns the DatadogFlagConfiguration field value if set, zero value otherwise.
func (o *ExperimentsBanditV2DTODataAttributes) GetDatadogFlagConfiguration() ExperimentsBanditFlagConfigurationV2 {
	if o == nil || o.DatadogFlagConfiguration == nil {
		var ret ExperimentsBanditFlagConfigurationV2
		return ret
	}
	return *o.DatadogFlagConfiguration
}

// GetDatadogFlagConfigurationOk returns a tuple with the DatadogFlagConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetDatadogFlagConfigurationOk() (*ExperimentsBanditFlagConfigurationV2, bool) {
	if o == nil || o.DatadogFlagConfiguration == nil {
		return nil, false
	}
	return o.DatadogFlagConfiguration, true
}

// HasDatadogFlagConfiguration returns a boolean if a field has been set.
func (o *ExperimentsBanditV2DTODataAttributes) HasDatadogFlagConfiguration() bool {
	return o != nil && o.DatadogFlagConfiguration != nil
}

// SetDatadogFlagConfiguration gets a reference to the given ExperimentsBanditFlagConfigurationV2 and assigns it to the DatadogFlagConfiguration field.
func (o *ExperimentsBanditV2DTODataAttributes) SetDatadogFlagConfiguration(v ExperimentsBanditFlagConfigurationV2) {
	o.DatadogFlagConfiguration = &v
}

// GetDescription returns the Description field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndedAt returns the EndedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetEndedAt() time.Time {
	if o == nil || o.EndedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// SetEndedAt sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}

// GetExposureSqlModelId returns the ExposureSqlModelId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetExposureSqlModelId() uuid.UUID {
	if o == nil || o.ExposureSqlModelId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ExposureSqlModelId.Get()
}

// GetExposureSqlModelIdOk returns a tuple with the ExposureSqlModelId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetExposureSqlModelIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposureSqlModelId.Get(), o.ExposureSqlModelId.IsSet()
}

// SetExposureSqlModelId sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetExposureSqlModelId(v uuid.UUID) {
	o.ExposureSqlModelId.Set(&v)
}

// GetLatestRefresh returns the LatestRefresh field value.
// If the value is explicit nil, the zero value for ExperimentsBanditRefreshV2 will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetLatestRefresh() ExperimentsBanditRefreshV2 {
	if o == nil || o.LatestRefresh.Get() == nil {
		var ret ExperimentsBanditRefreshV2
		return ret
	}
	return *o.LatestRefresh.Get()
}

// GetLatestRefreshOk returns a tuple with the LatestRefresh field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetLatestRefreshOk() (*ExperimentsBanditRefreshV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestRefresh.Get(), o.LatestRefresh.IsSet()
}

// SetLatestRefresh sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetLatestRefresh(v ExperimentsBanditRefreshV2) {
	o.LatestRefresh.Set(&v)
}

// GetLookbackSeconds returns the LookbackSeconds field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetLookbackSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LookbackSeconds
}

// GetLookbackSecondsOk returns a tuple with the LookbackSeconds field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetLookbackSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookbackSeconds, true
}

// SetLookbackSeconds sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetLookbackSeconds(v int64) {
	o.LookbackSeconds = v
}

// GetMetricId returns the MetricId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetMetricId() uuid.UUID {
	if o == nil || o.MetricId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.MetricId.Get()
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricId.Get(), o.MetricId.IsSet()
}

// SetMetricId sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetMetricId(v uuid.UUID) {
	o.MetricId.Set(&v)
}

// GetName returns the Name field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetName(v string) {
	o.Name = v
}

// GetOptimalVariantKey returns the OptimalVariantKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetOptimalVariantKey() string {
	if o == nil || o.OptimalVariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.OptimalVariantKey.Get()
}

// GetOptimalVariantKeyOk returns a tuple with the OptimalVariantKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetOptimalVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimalVariantKey.Get(), o.OptimalVariantKey.IsSet()
}

// SetOptimalVariantKey sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetOptimalVariantKey(v string) {
	o.OptimalVariantKey.Set(&v)
}

// GetOptimalVariantUnavailableReason returns the OptimalVariantUnavailableReason field value if set, zero value otherwise.
func (o *ExperimentsBanditV2DTODataAttributes) GetOptimalVariantUnavailableReason() ExperimentsBanditOptimalVariantUnavailableReason {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		var ret ExperimentsBanditOptimalVariantUnavailableReason
		return ret
	}
	return *o.OptimalVariantUnavailableReason
}

// GetOptimalVariantUnavailableReasonOk returns a tuple with the OptimalVariantUnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetOptimalVariantUnavailableReasonOk() (*ExperimentsBanditOptimalVariantUnavailableReason, bool) {
	if o == nil || o.OptimalVariantUnavailableReason == nil {
		return nil, false
	}
	return o.OptimalVariantUnavailableReason, true
}

// HasOptimalVariantUnavailableReason returns a boolean if a field has been set.
func (o *ExperimentsBanditV2DTODataAttributes) HasOptimalVariantUnavailableReason() bool {
	return o != nil && o.OptimalVariantUnavailableReason != nil
}

// SetOptimalVariantUnavailableReason gets a reference to the given ExperimentsBanditOptimalVariantUnavailableReason and assigns it to the OptimalVariantUnavailableReason field.
func (o *ExperimentsBanditV2DTODataAttributes) SetOptimalVariantUnavailableReason(v ExperimentsBanditOptimalVariantUnavailableReason) {
	o.OptimalVariantUnavailableReason = &v
}

// GetRefreshScheduleId returns the RefreshScheduleId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetRefreshScheduleId() string {
	if o == nil || o.RefreshScheduleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshScheduleId.Get()
}

// GetRefreshScheduleIdOk returns a tuple with the RefreshScheduleId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetRefreshScheduleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshScheduleId.Get(), o.RefreshScheduleId.IsSet()
}

// SetRefreshScheduleId sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetRefreshScheduleId(v string) {
	o.RefreshScheduleId.Set(&v)
}

// GetScheduledEndAt returns the ScheduledEndAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetScheduledEndAt() time.Time {
	if o == nil || o.ScheduledEndAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndAt.Get()
}

// GetScheduledEndAtOk returns a tuple with the ScheduledEndAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetScheduledEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledEndAt.Get(), o.ScheduledEndAt.IsSet()
}

// SetScheduledEndAt sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetScheduledEndAt(v time.Time) {
	o.ScheduledEndAt.Set(&v)
}

// GetScheduledEndFailure returns the ScheduledEndFailure field value if set, zero value otherwise.
func (o *ExperimentsBanditV2DTODataAttributes) GetScheduledEndFailure() ExperimentsBanditScheduledEndFailureV2 {
	if o == nil || o.ScheduledEndFailure == nil {
		var ret ExperimentsBanditScheduledEndFailureV2
		return ret
	}
	return *o.ScheduledEndFailure
}

// GetScheduledEndFailureOk returns a tuple with the ScheduledEndFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetScheduledEndFailureOk() (*ExperimentsBanditScheduledEndFailureV2, bool) {
	if o == nil || o.ScheduledEndFailure == nil {
		return nil, false
	}
	return o.ScheduledEndFailure, true
}

// HasScheduledEndFailure returns a boolean if a field has been set.
func (o *ExperimentsBanditV2DTODataAttributes) HasScheduledEndFailure() bool {
	return o != nil && o.ScheduledEndFailure != nil
}

// SetScheduledEndFailure gets a reference to the given ExperimentsBanditScheduledEndFailureV2 and assigns it to the ScheduledEndFailure field.
func (o *ExperimentsBanditV2DTODataAttributes) SetScheduledEndFailure(v ExperimentsBanditScheduledEndFailureV2) {
	o.ScheduledEndFailure = &v
}

// GetStartedAt returns the StartedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetStatus returns the Status field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetStatus() ExperimentsBanditStatus {
	if o == nil {
		var ret ExperimentsBanditStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetStatusOk() (*ExperimentsBanditStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetStatus(v ExperimentsBanditStatus) {
	o.Status = v
}

// GetSubjectTypeId returns the SubjectTypeId field value.
// If the value is explicit nil, the zero value for uuid.UUID will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetSubjectTypeId() uuid.UUID {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetSubjectTypeIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetSubjectTypeId(v uuid.UUID) {
	o.SubjectTypeId.Set(&v)
}

// GetTags returns the Tags field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teams, true
}

// SetTeams sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVariants returns the Variants field value.
func (o *ExperimentsBanditV2DTODataAttributes) GetVariants() []ExperimentsBanditVariantV2 {
	if o == nil {
		var ret []ExperimentsBanditVariantV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTODataAttributes) GetVariantsOk() (*[]ExperimentsBanditVariantV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetVariants(v []ExperimentsBanditVariantV2) {
	o.Variants = v
}

// GetWarmupSeconds returns the WarmupSeconds field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetWarmupSeconds() int64 {
	if o == nil || o.WarmupSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.WarmupSeconds.Get()
}

// GetWarmupSecondsOk returns a tuple with the WarmupSeconds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditV2DTODataAttributes) GetWarmupSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarmupSeconds.Get(), o.WarmupSeconds.IsSet()
}

// SetWarmupSeconds sets field value.
func (o *ExperimentsBanditV2DTODataAttributes) SetWarmupSeconds(v int64) {
	o.WarmupSeconds.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["active_policy"] = o.ActivePolicy.Get()
	toSerialize["aggregation_window_seconds"] = o.AggregationWindowSeconds
	toSerialize["bayesian_prior"] = o.BayesianPrior
	toSerialize["config_version"] = o.ConfigVersion
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DatadogFlagConfiguration != nil {
		toSerialize["datadog_flag_configuration"] = o.DatadogFlagConfiguration
	}
	toSerialize["description"] = o.Description
	toSerialize["ended_at"] = o.EndedAt.Get()
	toSerialize["exposure_sql_model_id"] = o.ExposureSqlModelId.Get()
	toSerialize["latest_refresh"] = o.LatestRefresh.Get()
	toSerialize["lookback_seconds"] = o.LookbackSeconds
	toSerialize["metric_id"] = o.MetricId.Get()
	toSerialize["name"] = o.Name
	toSerialize["optimal_variant_key"] = o.OptimalVariantKey.Get()
	if o.OptimalVariantUnavailableReason != nil {
		toSerialize["optimal_variant_unavailable_reason"] = o.OptimalVariantUnavailableReason
	}
	toSerialize["refresh_schedule_id"] = o.RefreshScheduleId.Get()
	toSerialize["scheduled_end_at"] = o.ScheduledEndAt.Get()
	if o.ScheduledEndFailure != nil {
		toSerialize["scheduled_end_failure"] = o.ScheduledEndFailure
	}
	toSerialize["started_at"] = o.StartedAt.Get()
	toSerialize["status"] = o.Status
	toSerialize["subject_type_id"] = o.SubjectTypeId.Get()
	toSerialize["tags"] = o.Tags
	toSerialize["teams"] = o.Teams
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["variants"] = o.Variants
	toSerialize["warmup_seconds"] = o.WarmupSeconds.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActivePolicy                    NullableExperimentsBanditPolicyV2                 `json:"active_policy"`
		AggregationWindowSeconds        *int64                                            `json:"aggregation_window_seconds"`
		BayesianPrior                   *ExperimentsBanditBayesianPriorV2                 `json:"bayesian_prior"`
		ConfigVersion                   *int64                                            `json:"config_version"`
		CreatedAt                       *time.Time                                        `json:"created_at"`
		DatadogFlagConfiguration        *ExperimentsBanditFlagConfigurationV2             `json:"datadog_flag_configuration,omitempty"`
		Description                     *string                                           `json:"description"`
		EndedAt                         datadog.NullableTime                              `json:"ended_at"`
		ExposureSqlModelId              datadog.NullableUUID                              `json:"exposure_sql_model_id"`
		LatestRefresh                   NullableExperimentsBanditRefreshV2                `json:"latest_refresh"`
		LookbackSeconds                 *int64                                            `json:"lookback_seconds"`
		MetricId                        datadog.NullableUUID                              `json:"metric_id"`
		Name                            *string                                           `json:"name"`
		OptimalVariantKey               datadog.NullableString                            `json:"optimal_variant_key"`
		OptimalVariantUnavailableReason *ExperimentsBanditOptimalVariantUnavailableReason `json:"optimal_variant_unavailable_reason,omitempty"`
		RefreshScheduleId               datadog.NullableString                            `json:"refresh_schedule_id"`
		ScheduledEndAt                  datadog.NullableTime                              `json:"scheduled_end_at"`
		ScheduledEndFailure             *ExperimentsBanditScheduledEndFailureV2           `json:"scheduled_end_failure,omitempty"`
		StartedAt                       datadog.NullableTime                              `json:"started_at"`
		Status                          *ExperimentsBanditStatus                          `json:"status"`
		SubjectTypeId                   datadog.NullableUUID                              `json:"subject_type_id"`
		Tags                            *[]string                                         `json:"tags"`
		Teams                           *[]string                                         `json:"teams"`
		UpdatedAt                       *time.Time                                        `json:"updated_at"`
		Variants                        *[]ExperimentsBanditVariantV2                     `json:"variants"`
		WarmupSeconds                   datadog.NullableInt64                             `json:"warmup_seconds"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.ActivePolicy.IsSet() {
		return fmt.Errorf("required field active_policy missing")
	}
	if all.AggregationWindowSeconds == nil {
		return fmt.Errorf("required field aggregation_window_seconds missing")
	}
	if all.BayesianPrior == nil {
		return fmt.Errorf("required field bayesian_prior missing")
	}
	if all.ConfigVersion == nil {
		return fmt.Errorf("required field config_version missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.EndedAt.IsSet() {
		return fmt.Errorf("required field ended_at missing")
	}
	if !all.ExposureSqlModelId.IsSet() {
		return fmt.Errorf("required field exposure_sql_model_id missing")
	}
	if !all.LatestRefresh.IsSet() {
		return fmt.Errorf("required field latest_refresh missing")
	}
	if all.LookbackSeconds == nil {
		return fmt.Errorf("required field lookback_seconds missing")
	}
	if !all.MetricId.IsSet() {
		return fmt.Errorf("required field metric_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.OptimalVariantKey.IsSet() {
		return fmt.Errorf("required field optimal_variant_key missing")
	}
	if !all.RefreshScheduleId.IsSet() {
		return fmt.Errorf("required field refresh_schedule_id missing")
	}
	if !all.ScheduledEndAt.IsSet() {
		return fmt.Errorf("required field scheduled_end_at missing")
	}
	if !all.StartedAt.IsSet() {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.SubjectTypeId.IsSet() {
		return fmt.Errorf("required field subject_type_id missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Teams == nil {
		return fmt.Errorf("required field teams missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	if !all.WarmupSeconds.IsSet() {
		return fmt.Errorf("required field warmup_seconds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active_policy", "aggregation_window_seconds", "bayesian_prior", "config_version", "created_at", "datadog_flag_configuration", "description", "ended_at", "exposure_sql_model_id", "latest_refresh", "lookback_seconds", "metric_id", "name", "optimal_variant_key", "optimal_variant_unavailable_reason", "refresh_schedule_id", "scheduled_end_at", "scheduled_end_failure", "started_at", "status", "subject_type_id", "tags", "teams", "updated_at", "variants", "warmup_seconds"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActivePolicy = all.ActivePolicy
	o.AggregationWindowSeconds = *all.AggregationWindowSeconds
	if all.BayesianPrior.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BayesianPrior = *all.BayesianPrior
	o.ConfigVersion = *all.ConfigVersion
	o.CreatedAt = *all.CreatedAt
	if all.DatadogFlagConfiguration != nil && all.DatadogFlagConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatadogFlagConfiguration = all.DatadogFlagConfiguration
	o.Description = *all.Description
	o.EndedAt = all.EndedAt
	o.ExposureSqlModelId = all.ExposureSqlModelId
	o.LatestRefresh = all.LatestRefresh
	o.LookbackSeconds = *all.LookbackSeconds
	o.MetricId = all.MetricId
	o.Name = *all.Name
	o.OptimalVariantKey = all.OptimalVariantKey
	if all.OptimalVariantUnavailableReason != nil && !all.OptimalVariantUnavailableReason.IsValid() {
		hasInvalidField = true
	} else {
		o.OptimalVariantUnavailableReason = all.OptimalVariantUnavailableReason
	}
	o.RefreshScheduleId = all.RefreshScheduleId
	o.ScheduledEndAt = all.ScheduledEndAt
	if all.ScheduledEndFailure != nil && all.ScheduledEndFailure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScheduledEndFailure = all.ScheduledEndFailure
	o.StartedAt = all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.SubjectTypeId = all.SubjectTypeId
	o.Tags = *all.Tags
	o.Teams = *all.Teams
	o.UpdatedAt = *all.UpdatedAt
	o.Variants = *all.Variants
	o.WarmupSeconds = all.WarmupSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
