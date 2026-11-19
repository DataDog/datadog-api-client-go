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

// ExperimentsCreateBanditV2RequestDataAttributes The new draft's setup. Only `name` is required. Whether the draft is ready to run is checked when it starts.
type ExperimentsCreateBanditV2RequestDataAttributes struct {
	// The aggregation window: how long after an exposure the bandit keeps counting its reward, in seconds. Defaults to 3600 (1 hour).
	AggregationWindowSeconds *int64 `json:"aggregation_window_seconds,omitempty"`
	// The prior the bandit's estimates start from, in the metric's units. Omitted fields keep their current values, or their defaults on create. A change takes effect on the next refresh.
	BayesianPrior *ExperimentsBanditBayesianPriorInputV2 `json:"bayesian_prior,omitempty"`
	// The flag and environment the bandit serves on, and who it serves there. The server chooses the rule's `allocation_key`. On create, `feature_flag_id` and `environment_id` are required, and `null` is the same as omitting it. On update, omitted fields are unchanged, and `null` takes a draft off its flag, which also clears its variants and scheduled end. Moving a draft to another flag clears its variants unless the request sets them. Only a draft can change it.
	DatadogFlagConfiguration NullableExperimentsBanditFlagConfigurationInputV2 `json:"datadog_flag_configuration,omitempty"`
	// The bandit's description. Defaults to empty.
	Description *string `json:"description,omitempty"`
	// The exposure SQL model the bandit reads its exposures from.
	ExposureSqlModelId *uuid.UUID `json:"exposure_sql_model_id,omitempty"`
	// How far back each refresh reads observations, in seconds. Defaults to 2592000 (30 days).
	LookbackSeconds *int64 `json:"lookback_seconds,omitempty"`
	// The metric the bandit optimizes. It must read a warehouse metric source.
	MetricId *uuid.UUID `json:"metric_id,omitempty"`
	// The bandit's name.
	Name string `json:"name"`
	// The schedule the bandit refreshes on. Defaults to the organization's default schedule.
	RefreshScheduleId *string `json:"refresh_schedule_id,omitempty"`
	// When the bandit ends and serves its fallback. Needs a fallback variant.
	ScheduledEndAt *time.Time `json:"scheduled_end_at,omitempty"`
	// The subject type the bandit randomizes and measures by.
	SubjectTypeId *uuid.UUID `json:"subject_type_id,omitempty"`
	// Tags for the bandit. The flag's tags are added too.
	Tags []string `json:"tags,omitempty"`
	// Handles of the teams that own the bandit.
	Teams []string `json:"teams,omitempty"`
	// The flag variants the bandit uses. Needs `datadog_flag_configuration`. The server fills in each variant's `key` and `name` from the flag.
	Variants []ExperimentsBanditVariantInputV2 `json:"variants,omitempty"`
	// How long the bandit serves its arms equally after it starts, in seconds. Defaults to no warmup.
	WarmupSeconds *int64 `json:"warmup_seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateBanditV2RequestDataAttributes instantiates a new ExperimentsCreateBanditV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateBanditV2RequestDataAttributes(name string) *ExperimentsCreateBanditV2RequestDataAttributes {
	this := ExperimentsCreateBanditV2RequestDataAttributes{}
	this.Name = name
	return &this
}

// NewExperimentsCreateBanditV2RequestDataAttributesWithDefaults instantiates a new ExperimentsCreateBanditV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateBanditV2RequestDataAttributesWithDefaults() *ExperimentsCreateBanditV2RequestDataAttributes {
	this := ExperimentsCreateBanditV2RequestDataAttributes{}
	return &this
}

// GetAggregationWindowSeconds returns the AggregationWindowSeconds field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetAggregationWindowSeconds() int64 {
	if o == nil || o.AggregationWindowSeconds == nil {
		var ret int64
		return ret
	}
	return *o.AggregationWindowSeconds
}

// GetAggregationWindowSecondsOk returns a tuple with the AggregationWindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetAggregationWindowSecondsOk() (*int64, bool) {
	if o == nil || o.AggregationWindowSeconds == nil {
		return nil, false
	}
	return o.AggregationWindowSeconds, true
}

// HasAggregationWindowSeconds returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasAggregationWindowSeconds() bool {
	return o != nil && o.AggregationWindowSeconds != nil
}

// SetAggregationWindowSeconds gets a reference to the given int64 and assigns it to the AggregationWindowSeconds field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetAggregationWindowSeconds(v int64) {
	o.AggregationWindowSeconds = &v
}

// GetBayesianPrior returns the BayesianPrior field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetBayesianPrior() ExperimentsBanditBayesianPriorInputV2 {
	if o == nil || o.BayesianPrior == nil {
		var ret ExperimentsBanditBayesianPriorInputV2
		return ret
	}
	return *o.BayesianPrior
}

// GetBayesianPriorOk returns a tuple with the BayesianPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetBayesianPriorOk() (*ExperimentsBanditBayesianPriorInputV2, bool) {
	if o == nil || o.BayesianPrior == nil {
		return nil, false
	}
	return o.BayesianPrior, true
}

// HasBayesianPrior returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasBayesianPrior() bool {
	return o != nil && o.BayesianPrior != nil
}

// SetBayesianPrior gets a reference to the given ExperimentsBanditBayesianPriorInputV2 and assigns it to the BayesianPrior field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetBayesianPrior(v ExperimentsBanditBayesianPriorInputV2) {
	o.BayesianPrior = &v
}

// GetDatadogFlagConfiguration returns the DatadogFlagConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetDatadogFlagConfiguration() ExperimentsBanditFlagConfigurationInputV2 {
	if o == nil || o.DatadogFlagConfiguration.Get() == nil {
		var ret ExperimentsBanditFlagConfigurationInputV2
		return ret
	}
	return *o.DatadogFlagConfiguration.Get()
}

// GetDatadogFlagConfigurationOk returns a tuple with the DatadogFlagConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetDatadogFlagConfigurationOk() (*ExperimentsBanditFlagConfigurationInputV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogFlagConfiguration.Get(), o.DatadogFlagConfiguration.IsSet()
}

// HasDatadogFlagConfiguration returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasDatadogFlagConfiguration() bool {
	return o != nil && o.DatadogFlagConfiguration.IsSet()
}

// SetDatadogFlagConfiguration gets a reference to the given NullableExperimentsBanditFlagConfigurationInputV2 and assigns it to the DatadogFlagConfiguration field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetDatadogFlagConfiguration(v ExperimentsBanditFlagConfigurationInputV2) {
	o.DatadogFlagConfiguration.Set(&v)
}

// SetDatadogFlagConfigurationNil sets the value for DatadogFlagConfiguration to be an explicit nil.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetDatadogFlagConfigurationNil() {
	o.DatadogFlagConfiguration.Set(nil)
}

// UnsetDatadogFlagConfiguration ensures that no value is present for DatadogFlagConfiguration, not even an explicit nil.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) UnsetDatadogFlagConfiguration() {
	o.DatadogFlagConfiguration.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExposureSqlModelId returns the ExposureSqlModelId field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetExposureSqlModelId() uuid.UUID {
	if o == nil || o.ExposureSqlModelId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ExposureSqlModelId
}

// GetExposureSqlModelIdOk returns a tuple with the ExposureSqlModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetExposureSqlModelIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ExposureSqlModelId == nil {
		return nil, false
	}
	return o.ExposureSqlModelId, true
}

// HasExposureSqlModelId returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasExposureSqlModelId() bool {
	return o != nil && o.ExposureSqlModelId != nil
}

// SetExposureSqlModelId gets a reference to the given uuid.UUID and assigns it to the ExposureSqlModelId field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetExposureSqlModelId(v uuid.UUID) {
	o.ExposureSqlModelId = &v
}

// GetLookbackSeconds returns the LookbackSeconds field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetLookbackSeconds() int64 {
	if o == nil || o.LookbackSeconds == nil {
		var ret int64
		return ret
	}
	return *o.LookbackSeconds
}

// GetLookbackSecondsOk returns a tuple with the LookbackSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetLookbackSecondsOk() (*int64, bool) {
	if o == nil || o.LookbackSeconds == nil {
		return nil, false
	}
	return o.LookbackSeconds, true
}

// HasLookbackSeconds returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasLookbackSeconds() bool {
	return o != nil && o.LookbackSeconds != nil
}

// SetLookbackSeconds gets a reference to the given int64 and assigns it to the LookbackSeconds field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetLookbackSeconds(v int64) {
	o.LookbackSeconds = &v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetMetricId() uuid.UUID {
	if o == nil || o.MetricId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil || o.MetricId == nil {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasMetricId() bool {
	return o != nil && o.MetricId != nil
}

// SetMetricId gets a reference to the given uuid.UUID and assigns it to the MetricId field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetMetricId(v uuid.UUID) {
	o.MetricId = &v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRefreshScheduleId returns the RefreshScheduleId field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetRefreshScheduleId() string {
	if o == nil || o.RefreshScheduleId == nil {
		var ret string
		return ret
	}
	return *o.RefreshScheduleId
}

// GetRefreshScheduleIdOk returns a tuple with the RefreshScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetRefreshScheduleIdOk() (*string, bool) {
	if o == nil || o.RefreshScheduleId == nil {
		return nil, false
	}
	return o.RefreshScheduleId, true
}

// HasRefreshScheduleId returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasRefreshScheduleId() bool {
	return o != nil && o.RefreshScheduleId != nil
}

// SetRefreshScheduleId gets a reference to the given string and assigns it to the RefreshScheduleId field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetRefreshScheduleId(v string) {
	o.RefreshScheduleId = &v
}

// GetScheduledEndAt returns the ScheduledEndAt field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetScheduledEndAt() time.Time {
	if o == nil || o.ScheduledEndAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndAt
}

// GetScheduledEndAtOk returns a tuple with the ScheduledEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetScheduledEndAtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledEndAt == nil {
		return nil, false
	}
	return o.ScheduledEndAt, true
}

// HasScheduledEndAt returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasScheduledEndAt() bool {
	return o != nil && o.ScheduledEndAt != nil
}

// SetScheduledEndAt gets a reference to the given time.Time and assigns it to the ScheduledEndAt field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetScheduledEndAt(v time.Time) {
	o.ScheduledEndAt = &v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetSubjectTypeId() uuid.UUID {
	if o == nil || o.SubjectTypeId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetSubjectTypeIdOk() (*uuid.UUID, bool) {
	if o == nil || o.SubjectTypeId == nil {
		return nil, false
	}
	return o.SubjectTypeId, true
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId != nil
}

// SetSubjectTypeId gets a reference to the given uuid.UUID and assigns it to the SubjectTypeId field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetSubjectTypeId(v uuid.UUID) {
	o.SubjectTypeId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetVariants() []ExperimentsBanditVariantInputV2 {
	if o == nil || o.Variants == nil {
		var ret []ExperimentsBanditVariantInputV2
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetVariantsOk() (*[]ExperimentsBanditVariantInputV2, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []ExperimentsBanditVariantInputV2 and assigns it to the Variants field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetVariants(v []ExperimentsBanditVariantInputV2) {
	o.Variants = v
}

// GetWarmupSeconds returns the WarmupSeconds field value if set, zero value otherwise.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetWarmupSeconds() int64 {
	if o == nil || o.WarmupSeconds == nil {
		var ret int64
		return ret
	}
	return *o.WarmupSeconds
}

// GetWarmupSecondsOk returns a tuple with the WarmupSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) GetWarmupSecondsOk() (*int64, bool) {
	if o == nil || o.WarmupSeconds == nil {
		return nil, false
	}
	return o.WarmupSeconds, true
}

// HasWarmupSeconds returns a boolean if a field has been set.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) HasWarmupSeconds() bool {
	return o != nil && o.WarmupSeconds != nil
}

// SetWarmupSeconds gets a reference to the given int64 and assigns it to the WarmupSeconds field.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) SetWarmupSeconds(v int64) {
	o.WarmupSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateBanditV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationWindowSeconds != nil {
		toSerialize["aggregation_window_seconds"] = o.AggregationWindowSeconds
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
	if o.ExposureSqlModelId != nil {
		toSerialize["exposure_sql_model_id"] = o.ExposureSqlModelId
	}
	if o.LookbackSeconds != nil {
		toSerialize["lookback_seconds"] = o.LookbackSeconds
	}
	if o.MetricId != nil {
		toSerialize["metric_id"] = o.MetricId
	}
	toSerialize["name"] = o.Name
	if o.RefreshScheduleId != nil {
		toSerialize["refresh_schedule_id"] = o.RefreshScheduleId
	}
	if o.ScheduledEndAt != nil {
		if o.ScheduledEndAt.Nanosecond() == 0 {
			toSerialize["scheduled_end_at"] = o.ScheduledEndAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["scheduled_end_at"] = o.ScheduledEndAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SubjectTypeId != nil {
		toSerialize["subject_type_id"] = o.SubjectTypeId
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
	if o.WarmupSeconds != nil {
		toSerialize["warmup_seconds"] = o.WarmupSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateBanditV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationWindowSeconds *int64                                            `json:"aggregation_window_seconds,omitempty"`
		BayesianPrior            *ExperimentsBanditBayesianPriorInputV2            `json:"bayesian_prior,omitempty"`
		DatadogFlagConfiguration NullableExperimentsBanditFlagConfigurationInputV2 `json:"datadog_flag_configuration,omitempty"`
		Description              *string                                           `json:"description,omitempty"`
		ExposureSqlModelId       *uuid.UUID                                        `json:"exposure_sql_model_id,omitempty"`
		LookbackSeconds          *int64                                            `json:"lookback_seconds,omitempty"`
		MetricId                 *uuid.UUID                                        `json:"metric_id,omitempty"`
		Name                     *string                                           `json:"name"`
		RefreshScheduleId        *string                                           `json:"refresh_schedule_id,omitempty"`
		ScheduledEndAt           *time.Time                                        `json:"scheduled_end_at,omitempty"`
		SubjectTypeId            *uuid.UUID                                        `json:"subject_type_id,omitempty"`
		Tags                     []string                                          `json:"tags,omitempty"`
		Teams                    []string                                          `json:"teams,omitempty"`
		Variants                 []ExperimentsBanditVariantInputV2                 `json:"variants,omitempty"`
		WarmupSeconds            *int64                                            `json:"warmup_seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	o.Name = *all.Name
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
