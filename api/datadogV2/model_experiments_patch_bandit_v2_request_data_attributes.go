// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchBanditV2RequestDataAttributes The attributes to change. Omitted attributes are unchanged. Which ones can change depends on the bandit's status.
type ExperimentsPatchBanditV2RequestDataAttributes struct {
	// The aggregation window: how long after an exposure the bandit keeps counting its reward, in seconds. `null` restores the default of 3600 (1 hour). Only a draft can change it.
	AggregationWindowSeconds datadog.NullableInt64 `json:"aggregation_window_seconds,omitempty"`
	// The prior the bandit's estimates start from, in the metric's units. Omitted fields keep their current values, or their defaults on create. A change takes effect on the next refresh.
	BayesianPrior *ExperimentsBanditBayesianPriorInputV2 `json:"bayesian_prior,omitempty"`
	// The flag and environment the bandit serves on, and who it serves there. The server chooses the rule's `allocation_key`. On create, `feature_flag_id` and `environment_id` are required, and `null` is the same as omitting it. On update, omitted fields are unchanged, and `null` takes a draft off its flag, which also clears its variants and scheduled end. Moving a draft to another flag clears its variants unless the request sets them. Only a draft can change it.
	DatadogFlagConfiguration NullableExperimentsBanditFlagConfigurationInputV2 `json:"datadog_flag_configuration,omitempty"`
	// The bandit's description.
	Description *string `json:"description,omitempty"`
	// The exposure SQL model the bandit reads its exposures from. Only a draft can change it.
	ExposureSqlModelId datadog.NullableUUID `json:"exposure_sql_model_id,omitempty"`
	// How far back each refresh reads observations, in seconds. `null` restores the default of 2592000 (30 days). Only a draft can change it.
	LookbackSeconds datadog.NullableInt64 `json:"lookback_seconds,omitempty"`
	// The metric the bandit optimizes. It must read a warehouse metric source. Only a draft can change it.
	MetricId datadog.NullableUUID `json:"metric_id,omitempty"`
	// The bandit's name.
	Name *string `json:"name,omitempty"`
	// The schedule the bandit refreshes on.
	RefreshScheduleId *string `json:"refresh_schedule_id,omitempty"`
	// When the bandit ends and serves its fallback. Needs a fallback variant. `null` removes the scheduled end.
	ScheduledEndAt datadog.NullableTime `json:"scheduled_end_at,omitempty"`
	// The subject type the bandit randomizes and measures by. Only a draft can change it.
	SubjectTypeId datadog.NullableUUID `json:"subject_type_id,omitempty"`
	// Replaces the bandit's tags and keeps its teams. Allowed in any status.
	Tags []string `json:"tags,omitempty"`
	// Replaces the bandit's teams and keeps its tags. Allowed in any status.
	Teams []string `json:"teams,omitempty"`
	// Replaces the variants the bandit uses. Once the bandit has started, only `is_fallback` can change, so send the current arms with a different `is_fallback`. Only a draft can have an empty list.
	Variants []ExperimentsBanditVariantInputV2 `json:"variants,omitempty"`
	// How long the bandit serves its arms equally after it starts, in seconds. `null` removes the warmup. Only a draft can change it.
	WarmupSeconds datadog.NullableInt64 `json:"warmup_seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchBanditV2RequestDataAttributes instantiates a new ExperimentsPatchBanditV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchBanditV2RequestDataAttributes() *ExperimentsPatchBanditV2RequestDataAttributes {
	this := ExperimentsPatchBanditV2RequestDataAttributes{}
	return &this
}

// NewExperimentsPatchBanditV2RequestDataAttributesWithDefaults instantiates a new ExperimentsPatchBanditV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchBanditV2RequestDataAttributesWithDefaults() *ExperimentsPatchBanditV2RequestDataAttributes {
	this := ExperimentsPatchBanditV2RequestDataAttributes{}
	return &this
}

// GetAggregationWindowSeconds returns the AggregationWindowSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetAggregationWindowSeconds() int64 {
	if o == nil || o.AggregationWindowSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AggregationWindowSeconds.Get()
}

// GetAggregationWindowSecondsOk returns a tuple with the AggregationWindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetAggregationWindowSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationWindowSeconds.Get(), o.AggregationWindowSeconds.IsSet()
}

// HasAggregationWindowSeconds returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasAggregationWindowSeconds() bool {
	return o != nil && o.AggregationWindowSeconds.IsSet()
}

// SetAggregationWindowSeconds gets a reference to the given datadog.NullableInt64 and assigns it to the AggregationWindowSeconds field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetAggregationWindowSeconds(v int64) {
	o.AggregationWindowSeconds.Set(&v)
}

// SetAggregationWindowSecondsNil sets the value for AggregationWindowSeconds to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetAggregationWindowSecondsNil() {
	o.AggregationWindowSeconds.Set(nil)
}

// UnsetAggregationWindowSeconds ensures that no value is present for AggregationWindowSeconds, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetAggregationWindowSeconds() {
	o.AggregationWindowSeconds.Unset()
}

// GetBayesianPrior returns the BayesianPrior field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetBayesianPrior() ExperimentsBanditBayesianPriorInputV2 {
	if o == nil || o.BayesianPrior == nil {
		var ret ExperimentsBanditBayesianPriorInputV2
		return ret
	}
	return *o.BayesianPrior
}

// GetBayesianPriorOk returns a tuple with the BayesianPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetBayesianPriorOk() (*ExperimentsBanditBayesianPriorInputV2, bool) {
	if o == nil || o.BayesianPrior == nil {
		return nil, false
	}
	return o.BayesianPrior, true
}

// HasBayesianPrior returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasBayesianPrior() bool {
	return o != nil && o.BayesianPrior != nil
}

// SetBayesianPrior gets a reference to the given ExperimentsBanditBayesianPriorInputV2 and assigns it to the BayesianPrior field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetBayesianPrior(v ExperimentsBanditBayesianPriorInputV2) {
	o.BayesianPrior = &v
}

// GetDatadogFlagConfiguration returns the DatadogFlagConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetDatadogFlagConfiguration() ExperimentsBanditFlagConfigurationInputV2 {
	if o == nil || o.DatadogFlagConfiguration.Get() == nil {
		var ret ExperimentsBanditFlagConfigurationInputV2
		return ret
	}
	return *o.DatadogFlagConfiguration.Get()
}

// GetDatadogFlagConfigurationOk returns a tuple with the DatadogFlagConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetDatadogFlagConfigurationOk() (*ExperimentsBanditFlagConfigurationInputV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogFlagConfiguration.Get(), o.DatadogFlagConfiguration.IsSet()
}

// HasDatadogFlagConfiguration returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasDatadogFlagConfiguration() bool {
	return o != nil && o.DatadogFlagConfiguration.IsSet()
}

// SetDatadogFlagConfiguration gets a reference to the given NullableExperimentsBanditFlagConfigurationInputV2 and assigns it to the DatadogFlagConfiguration field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetDatadogFlagConfiguration(v ExperimentsBanditFlagConfigurationInputV2) {
	o.DatadogFlagConfiguration.Set(&v)
}

// SetDatadogFlagConfigurationNil sets the value for DatadogFlagConfiguration to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetDatadogFlagConfigurationNil() {
	o.DatadogFlagConfiguration.Set(nil)
}

// UnsetDatadogFlagConfiguration ensures that no value is present for DatadogFlagConfiguration, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetDatadogFlagConfiguration() {
	o.DatadogFlagConfiguration.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExposureSqlModelId returns the ExposureSqlModelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetExposureSqlModelId() uuid.UUID {
	if o == nil || o.ExposureSqlModelId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ExposureSqlModelId.Get()
}

// GetExposureSqlModelIdOk returns a tuple with the ExposureSqlModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetExposureSqlModelIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposureSqlModelId.Get(), o.ExposureSqlModelId.IsSet()
}

// HasExposureSqlModelId returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasExposureSqlModelId() bool {
	return o != nil && o.ExposureSqlModelId.IsSet()
}

// SetExposureSqlModelId gets a reference to the given datadog.NullableUUID and assigns it to the ExposureSqlModelId field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetExposureSqlModelId(v uuid.UUID) {
	o.ExposureSqlModelId.Set(&v)
}

// SetExposureSqlModelIdNil sets the value for ExposureSqlModelId to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetExposureSqlModelIdNil() {
	o.ExposureSqlModelId.Set(nil)
}

// UnsetExposureSqlModelId ensures that no value is present for ExposureSqlModelId, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetExposureSqlModelId() {
	o.ExposureSqlModelId.Unset()
}

// GetLookbackSeconds returns the LookbackSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetLookbackSeconds() int64 {
	if o == nil || o.LookbackSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LookbackSeconds.Get()
}

// GetLookbackSecondsOk returns a tuple with the LookbackSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetLookbackSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LookbackSeconds.Get(), o.LookbackSeconds.IsSet()
}

// HasLookbackSeconds returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasLookbackSeconds() bool {
	return o != nil && o.LookbackSeconds.IsSet()
}

// SetLookbackSeconds gets a reference to the given datadog.NullableInt64 and assigns it to the LookbackSeconds field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetLookbackSeconds(v int64) {
	o.LookbackSeconds.Set(&v)
}

// SetLookbackSecondsNil sets the value for LookbackSeconds to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetLookbackSecondsNil() {
	o.LookbackSeconds.Set(nil)
}

// UnsetLookbackSeconds ensures that no value is present for LookbackSeconds, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetLookbackSeconds() {
	o.LookbackSeconds.Unset()
}

// GetMetricId returns the MetricId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetMetricId() uuid.UUID {
	if o == nil || o.MetricId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.MetricId.Get()
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricId.Get(), o.MetricId.IsSet()
}

// HasMetricId returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasMetricId() bool {
	return o != nil && o.MetricId.IsSet()
}

// SetMetricId gets a reference to the given datadog.NullableUUID and assigns it to the MetricId field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetMetricId(v uuid.UUID) {
	o.MetricId.Set(&v)
}

// SetMetricIdNil sets the value for MetricId to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetMetricIdNil() {
	o.MetricId.Set(nil)
}

// UnsetMetricId ensures that no value is present for MetricId, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetMetricId() {
	o.MetricId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRefreshScheduleId returns the RefreshScheduleId field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetRefreshScheduleId() string {
	if o == nil || o.RefreshScheduleId == nil {
		var ret string
		return ret
	}
	return *o.RefreshScheduleId
}

// GetRefreshScheduleIdOk returns a tuple with the RefreshScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetRefreshScheduleIdOk() (*string, bool) {
	if o == nil || o.RefreshScheduleId == nil {
		return nil, false
	}
	return o.RefreshScheduleId, true
}

// HasRefreshScheduleId returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasRefreshScheduleId() bool {
	return o != nil && o.RefreshScheduleId != nil
}

// SetRefreshScheduleId gets a reference to the given string and assigns it to the RefreshScheduleId field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetRefreshScheduleId(v string) {
	o.RefreshScheduleId = &v
}

// GetScheduledEndAt returns the ScheduledEndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetScheduledEndAt() time.Time {
	if o == nil || o.ScheduledEndAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndAt.Get()
}

// GetScheduledEndAtOk returns a tuple with the ScheduledEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetScheduledEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledEndAt.Get(), o.ScheduledEndAt.IsSet()
}

// HasScheduledEndAt returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasScheduledEndAt() bool {
	return o != nil && o.ScheduledEndAt.IsSet()
}

// SetScheduledEndAt gets a reference to the given datadog.NullableTime and assigns it to the ScheduledEndAt field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetScheduledEndAt(v time.Time) {
	o.ScheduledEndAt.Set(&v)
}

// SetScheduledEndAtNil sets the value for ScheduledEndAt to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetScheduledEndAtNil() {
	o.ScheduledEndAt.Set(nil)
}

// UnsetScheduledEndAt ensures that no value is present for ScheduledEndAt, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetScheduledEndAt() {
	o.ScheduledEndAt.Unset()
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetSubjectTypeId() uuid.UUID {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetSubjectTypeIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId gets a reference to the given datadog.NullableUUID and assigns it to the SubjectTypeId field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetSubjectTypeId(v uuid.UUID) {
	o.SubjectTypeId.Set(&v)
}

// SetSubjectTypeIdNil sets the value for SubjectTypeId to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetSubjectTypeIdNil() {
	o.SubjectTypeId.Set(nil)
}

// UnsetSubjectTypeId ensures that no value is present for SubjectTypeId, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetSubjectTypeId() {
	o.SubjectTypeId.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetVariants() []ExperimentsBanditVariantInputV2 {
	if o == nil || o.Variants == nil {
		var ret []ExperimentsBanditVariantInputV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetVariantsOk() (*[]ExperimentsBanditVariantInputV2, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []ExperimentsBanditVariantInputV2 and assigns it to the Variants field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetVariants(v []ExperimentsBanditVariantInputV2) {
	o.Variants = v
}

// GetWarmupSeconds returns the WarmupSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetWarmupSeconds() int64 {
	if o == nil || o.WarmupSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.WarmupSeconds.Get()
}

// GetWarmupSecondsOk returns a tuple with the WarmupSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) GetWarmupSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarmupSeconds.Get(), o.WarmupSeconds.IsSet()
}

// HasWarmupSeconds returns a boolean if a field has been set.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) HasWarmupSeconds() bool {
	return o != nil && o.WarmupSeconds.IsSet()
}

// SetWarmupSeconds gets a reference to the given datadog.NullableInt64 and assigns it to the WarmupSeconds field.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetWarmupSeconds(v int64) {
	o.WarmupSeconds.Set(&v)
}

// SetWarmupSecondsNil sets the value for WarmupSeconds to be an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) SetWarmupSecondsNil() {
	o.WarmupSeconds.Set(nil)
}

// UnsetWarmupSeconds ensures that no value is present for WarmupSeconds, not even an explicit nil.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnsetWarmupSeconds() {
	o.WarmupSeconds.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchBanditV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationWindowSeconds.IsSet() {
		toSerialize["aggregation_window_seconds"] = o.AggregationWindowSeconds.Get()
	}
	if o.BayesianPrior != nil {
		toSerialize["bayesian_prior"] = o.BayesianPrior
	}
	if o.DatadogFlagConfiguration.IsSet() {
		toSerialize["datadog_flag_configuration"] = o.DatadogFlagConfiguration.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExposureSqlModelId.IsSet() {
		toSerialize["exposure_sql_model_id"] = o.ExposureSqlModelId.Get()
	}
	if o.LookbackSeconds.IsSet() {
		toSerialize["lookback_seconds"] = o.LookbackSeconds.Get()
	}
	if o.MetricId.IsSet() {
		toSerialize["metric_id"] = o.MetricId.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RefreshScheduleId != nil {
		toSerialize["refresh_schedule_id"] = o.RefreshScheduleId
	}
	if o.ScheduledEndAt.IsSet() {
		toSerialize["scheduled_end_at"] = o.ScheduledEndAt.Get()
	}
	if o.SubjectTypeId.IsSet() {
		toSerialize["subject_type_id"] = o.SubjectTypeId.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.WarmupSeconds.IsSet() {
		toSerialize["warmup_seconds"] = o.WarmupSeconds.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchBanditV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationWindowSeconds datadog.NullableInt64                             `json:"aggregation_window_seconds,omitempty"`
		BayesianPrior            *ExperimentsBanditBayesianPriorInputV2            `json:"bayesian_prior,omitempty"`
		DatadogFlagConfiguration NullableExperimentsBanditFlagConfigurationInputV2 `json:"datadog_flag_configuration,omitempty"`
		Description              *string                                           `json:"description,omitempty"`
		ExposureSqlModelId       datadog.NullableUUID                              `json:"exposure_sql_model_id,omitempty"`
		LookbackSeconds          datadog.NullableInt64                             `json:"lookback_seconds,omitempty"`
		MetricId                 datadog.NullableUUID                              `json:"metric_id,omitempty"`
		Name                     *string                                           `json:"name,omitempty"`
		RefreshScheduleId        *string                                           `json:"refresh_schedule_id,omitempty"`
		ScheduledEndAt           datadog.NullableTime                              `json:"scheduled_end_at,omitempty"`
		SubjectTypeId            datadog.NullableUUID                              `json:"subject_type_id,omitempty"`
		Tags                     []string                                          `json:"tags,omitempty"`
		Teams                    []string                                          `json:"teams,omitempty"`
		Variants                 []ExperimentsBanditVariantInputV2                 `json:"variants,omitempty"`
		WarmupSeconds            datadog.NullableInt64                             `json:"warmup_seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_window_seconds", "bayesian_prior", "datadog_flag_configuration", "description", "exposure_sql_model_id", "lookback_seconds", "metric_id", "name", "refresh_schedule_id", "scheduled_end_at", "subject_type_id", "tags", "teams", "variants", "warmup_seconds"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregationWindowSeconds = all.AggregationWindowSeconds
	if all.BayesianPrior != nil && all.BayesianPrior.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BayesianPrior = all.BayesianPrior
	o.DatadogFlagConfiguration = all.DatadogFlagConfiguration
	o.Description = all.Description
	o.ExposureSqlModelId = all.ExposureSqlModelId
	o.LookbackSeconds = all.LookbackSeconds
	o.MetricId = all.MetricId
	o.Name = all.Name
	o.RefreshScheduleId = all.RefreshScheduleId
	o.ScheduledEndAt = all.ScheduledEndAt
	o.SubjectTypeId = all.SubjectTypeId
	o.Tags = all.Tags
	o.Teams = all.Teams
	o.Variants = all.Variants
	o.WarmupSeconds = all.WarmupSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
