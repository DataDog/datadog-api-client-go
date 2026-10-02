// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2ResponseDataAttributes Details of the experiment.
type ExperimentsPatchExperimentV2ResponseDataAttributes struct {
	// End of the time window for experiment assignments.
	AssignmentsEndDate datadog.NullableTime `json:"assignments_end_date,omitempty"`
	// Start of the time window for experiment assignments.
	AssignmentsStartDate datadog.NullableTime `json:"assignments_start_date,omitempty"`
	// Time when the experiment was concluded.
	ConcludedAt datadog.NullableTime `json:"concluded_at,omitempty"`
	// Outcome and supporting text recorded when the experiment is concluded.
	Conclusion *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion `json:"conclusion,omitempty"`
	// Time when this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Feature flag, environment, and targeting configuration for the experiment.
	DatadogFlagConfiguration NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration `json:"datadog_flag_configuration,omitempty"`
	// Metrics used to decide the experiment outcome.
	DecisionMetrics []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems `json:"decision_metrics,omitempty"`
	// Key of the variant selected in the experiment decision.
	DecisionVariantKey *string `json:"decision_variant_key,omitempty"`
	// End of the time window for metric events.
	EventsEndDate datadog.NullableTime `json:"events_end_date,omitempty"`
	// Start of the time window for metric events.
	EventsStartDate datadog.NullableTime `json:"events_start_date,omitempty"`
	// Kind of experiment. STANDARD is an ordinary experiment. Other values, such as CANARY and HOLDOUT, identify experiments owned by another workflow. New kinds may be added; treat unknown values as non-standard.
	ExperimentType *string `json:"experiment_type,omitempty"`
	// Expected effect that the experiment is intended to test.
	Hypothesis *string `json:"hypothesis,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment.
	Name *string `json:"name,omitempty"`
	// Suffix used to identify the experiment's pipeline output table.
	PipelineTableSuffix *string `json:"pipeline_table_suffix,omitempty"`
	// ID of the protocol associated with the experiment.
	ProtocolId *string `json:"protocol_id,omitempty"`
	// Links to supporting material for the experiment.
	RelatedLinks []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems `json:"related_links,omitempty"`
	// Time of the most recent update to the experiment's results.
	ResultsLastUpdated datadog.NullableTime `json:"results_last_updated,omitempty"`
	// Properties used to split the experiment results into groups.
	SplitByProperties []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems `json:"split_by_properties,omitempty"`
	// Current stage in the experiment lifecycle.
	Status *ExperimentsExperimentV2DTODataAttributesStatus `json:"status,omitempty"`
	// Values of structured metadata fields attached to the experiment.
	StructuredMetadata []ExperimentsStructuredMetadataResponse `json:"structured_metadata,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId datadog.NullableString `json:"subject_type_id,omitempty"`
	// Summary text recorded for the experiment.
	Summary *string `json:"summary,omitempty"`
	// Tags attached to the experiment.
	Tags []string `json:"tags,omitempty"`
	// Teams associated with the experiment.
	Teams []string `json:"teams,omitempty"`
	// Traffic exposure fraction or schedule configured for the experiment.
	TrafficExposure *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure `json:"traffic_exposure,omitempty"`
	// Time when this resource was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Variants configured for the experiment.
	Variants []ExperimentsExperimentV2DTODataAttributesVariantsItems `json:"variants,omitempty"`
	// Warehouse exposure model and settings used to identify experiment assignments.
	WarehouseExposureConfiguration NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration `json:"warehouse_exposure_configuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2ResponseDataAttributes instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2ResponseDataAttributes() *ExperimentsPatchExperimentV2ResponseDataAttributes {
	this := ExperimentsPatchExperimentV2ResponseDataAttributes{}
	return &this
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesWithDefaults instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2ResponseDataAttributesWithDefaults() *ExperimentsPatchExperimentV2ResponseDataAttributes {
	this := ExperimentsPatchExperimentV2ResponseDataAttributes{}
	return &this
}

// GetAssignmentsEndDate returns the AssignmentsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetAssignmentsEndDate() time.Time {
	if o == nil || o.AssignmentsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsEndDate.Get()
}

// GetAssignmentsEndDateOk returns a tuple with the AssignmentsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetAssignmentsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsEndDate.Get(), o.AssignmentsEndDate.IsSet()
}

// HasAssignmentsEndDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasAssignmentsEndDate() bool {
	return o != nil && o.AssignmentsEndDate.IsSet()
}

// SetAssignmentsEndDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsEndDate field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetAssignmentsEndDate(v time.Time) {
	o.AssignmentsEndDate.Set(&v)
}

// SetAssignmentsEndDateNil sets the value for AssignmentsEndDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetAssignmentsEndDateNil() {
	o.AssignmentsEndDate.Set(nil)
}

// UnsetAssignmentsEndDate ensures that no value is present for AssignmentsEndDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetAssignmentsEndDate() {
	o.AssignmentsEndDate.Unset()
}

// GetAssignmentsStartDate returns the AssignmentsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetAssignmentsStartDate() time.Time {
	if o == nil || o.AssignmentsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsStartDate.Get()
}

// GetAssignmentsStartDateOk returns a tuple with the AssignmentsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetAssignmentsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsStartDate.Get(), o.AssignmentsStartDate.IsSet()
}

// HasAssignmentsStartDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasAssignmentsStartDate() bool {
	return o != nil && o.AssignmentsStartDate.IsSet()
}

// SetAssignmentsStartDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsStartDate field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetAssignmentsStartDate(v time.Time) {
	o.AssignmentsStartDate.Set(&v)
}

// SetAssignmentsStartDateNil sets the value for AssignmentsStartDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetAssignmentsStartDateNil() {
	o.AssignmentsStartDate.Set(nil)
}

// UnsetAssignmentsStartDate ensures that no value is present for AssignmentsStartDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetAssignmentsStartDate() {
	o.AssignmentsStartDate.Unset()
}

// GetConcludedAt returns the ConcludedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetConcludedAt() time.Time {
	if o == nil || o.ConcludedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ConcludedAt.Get()
}

// GetConcludedAtOk returns a tuple with the ConcludedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetConcludedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConcludedAt.Get(), o.ConcludedAt.IsSet()
}

// HasConcludedAt returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasConcludedAt() bool {
	return o != nil && o.ConcludedAt.IsSet()
}

// SetConcludedAt gets a reference to the given datadog.NullableTime and assigns it to the ConcludedAt field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetConcludedAt(v time.Time) {
	o.ConcludedAt.Set(&v)
}

// SetConcludedAtNil sets the value for ConcludedAt to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetConcludedAtNil() {
	o.ConcludedAt.Set(nil)
}

// UnsetConcludedAt ensures that no value is present for ConcludedAt, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetConcludedAt() {
	o.ConcludedAt.Unset()
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetConclusion() ExperimentsPatchExperimentV2ResponseDataAttributesConclusion {
	if o == nil || o.Conclusion == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesConclusion
		return ret
	}
	return *o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetConclusionOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesConclusion, bool) {
	if o == nil || o.Conclusion == nil {
		return nil, false
	}
	return o.Conclusion, true
}

// HasConclusion returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasConclusion() bool {
	return o != nil && o.Conclusion != nil
}

// SetConclusion gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesConclusion and assigns it to the Conclusion field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetConclusion(v ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) {
	o.Conclusion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatadogFlagConfiguration returns the DatadogFlagConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDatadogFlagConfiguration() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration {
	if o == nil || o.DatadogFlagConfiguration.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration
		return ret
	}
	return *o.DatadogFlagConfiguration.Get()
}

// GetDatadogFlagConfigurationOk returns a tuple with the DatadogFlagConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDatadogFlagConfigurationOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogFlagConfiguration.Get(), o.DatadogFlagConfiguration.IsSet()
}

// HasDatadogFlagConfiguration returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasDatadogFlagConfiguration() bool {
	return o != nil && o.DatadogFlagConfiguration.IsSet()
}

// SetDatadogFlagConfiguration gets a reference to the given NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration and assigns it to the DatadogFlagConfiguration field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetDatadogFlagConfiguration(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) {
	o.DatadogFlagConfiguration.Set(&v)
}

// SetDatadogFlagConfigurationNil sets the value for DatadogFlagConfiguration to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetDatadogFlagConfigurationNil() {
	o.DatadogFlagConfiguration.Set(nil)
}

// UnsetDatadogFlagConfiguration ensures that no value is present for DatadogFlagConfiguration, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetDatadogFlagConfiguration() {
	o.DatadogFlagConfiguration.Unset()
}

// GetDecisionMetrics returns the DecisionMetrics field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDecisionMetrics() []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems {
	if o == nil || o.DecisionMetrics == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems
		return ret
	}
	return o.DecisionMetrics
}

// GetDecisionMetricsOk returns a tuple with the DecisionMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDecisionMetricsOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems, bool) {
	if o == nil || o.DecisionMetrics == nil {
		return nil, false
	}
	return &o.DecisionMetrics, true
}

// HasDecisionMetrics returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasDecisionMetrics() bool {
	return o != nil && o.DecisionMetrics != nil
}

// SetDecisionMetrics gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems and assigns it to the DecisionMetrics field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetDecisionMetrics(v []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) {
	o.DecisionMetrics = v
}

// GetDecisionVariantKey returns the DecisionVariantKey field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDecisionVariantKey() string {
	if o == nil || o.DecisionVariantKey == nil {
		var ret string
		return ret
	}
	return *o.DecisionVariantKey
}

// GetDecisionVariantKeyOk returns a tuple with the DecisionVariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetDecisionVariantKeyOk() (*string, bool) {
	if o == nil || o.DecisionVariantKey == nil {
		return nil, false
	}
	return o.DecisionVariantKey, true
}

// HasDecisionVariantKey returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasDecisionVariantKey() bool {
	return o != nil && o.DecisionVariantKey != nil
}

// SetDecisionVariantKey gets a reference to the given string and assigns it to the DecisionVariantKey field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetDecisionVariantKey(v string) {
	o.DecisionVariantKey = &v
}

// GetEventsEndDate returns the EventsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetEventsEndDate() time.Time {
	if o == nil || o.EventsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsEndDate.Get()
}

// GetEventsEndDateOk returns a tuple with the EventsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetEventsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsEndDate.Get(), o.EventsEndDate.IsSet()
}

// HasEventsEndDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasEventsEndDate() bool {
	return o != nil && o.EventsEndDate.IsSet()
}

// SetEventsEndDate gets a reference to the given datadog.NullableTime and assigns it to the EventsEndDate field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetEventsEndDate(v time.Time) {
	o.EventsEndDate.Set(&v)
}

// SetEventsEndDateNil sets the value for EventsEndDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetEventsEndDateNil() {
	o.EventsEndDate.Set(nil)
}

// UnsetEventsEndDate ensures that no value is present for EventsEndDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetEventsEndDate() {
	o.EventsEndDate.Unset()
}

// GetEventsStartDate returns the EventsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetEventsStartDate() time.Time {
	if o == nil || o.EventsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsStartDate.Get()
}

// GetEventsStartDateOk returns a tuple with the EventsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetEventsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsStartDate.Get(), o.EventsStartDate.IsSet()
}

// HasEventsStartDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasEventsStartDate() bool {
	return o != nil && o.EventsStartDate.IsSet()
}

// SetEventsStartDate gets a reference to the given datadog.NullableTime and assigns it to the EventsStartDate field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetEventsStartDate(v time.Time) {
	o.EventsStartDate.Set(&v)
}

// SetEventsStartDateNil sets the value for EventsStartDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetEventsStartDateNil() {
	o.EventsStartDate.Set(nil)
}

// UnsetEventsStartDate ensures that no value is present for EventsStartDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetEventsStartDate() {
	o.EventsStartDate.Unset()
}

// GetExperimentType returns the ExperimentType field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetExperimentType() string {
	if o == nil || o.ExperimentType == nil {
		var ret string
		return ret
	}
	return *o.ExperimentType
}

// GetExperimentTypeOk returns a tuple with the ExperimentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetExperimentTypeOk() (*string, bool) {
	if o == nil || o.ExperimentType == nil {
		return nil, false
	}
	return o.ExperimentType, true
}

// HasExperimentType returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasExperimentType() bool {
	return o != nil && o.ExperimentType != nil
}

// SetExperimentType gets a reference to the given string and assigns it to the ExperimentType field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetExperimentType(v string) {
	o.ExperimentType = &v
}

// GetHypothesis returns the Hypothesis field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetHypothesis() string {
	if o == nil || o.Hypothesis == nil {
		var ret string
		return ret
	}
	return *o.Hypothesis
}

// GetHypothesisOk returns a tuple with the Hypothesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetHypothesisOk() (*string, bool) {
	if o == nil || o.Hypothesis == nil {
		return nil, false
	}
	return o.Hypothesis, true
}

// HasHypothesis returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasHypothesis() bool {
	return o != nil && o.Hypothesis != nil
}

// SetHypothesis gets a reference to the given string and assigns it to the Hypothesis field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetHypothesis(v string) {
	o.Hypothesis = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPipelineTableSuffix returns the PipelineTableSuffix field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetPipelineTableSuffix() string {
	if o == nil || o.PipelineTableSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineTableSuffix
}

// GetPipelineTableSuffixOk returns a tuple with the PipelineTableSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetPipelineTableSuffixOk() (*string, bool) {
	if o == nil || o.PipelineTableSuffix == nil {
		return nil, false
	}
	return o.PipelineTableSuffix, true
}

// HasPipelineTableSuffix returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasPipelineTableSuffix() bool {
	return o != nil && o.PipelineTableSuffix != nil
}

// SetPipelineTableSuffix gets a reference to the given string and assigns it to the PipelineTableSuffix field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetPipelineTableSuffix(v string) {
	o.PipelineTableSuffix = &v
}

// GetProtocolId returns the ProtocolId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetProtocolId() string {
	if o == nil || o.ProtocolId == nil {
		var ret string
		return ret
	}
	return *o.ProtocolId
}

// GetProtocolIdOk returns a tuple with the ProtocolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetProtocolIdOk() (*string, bool) {
	if o == nil || o.ProtocolId == nil {
		return nil, false
	}
	return o.ProtocolId, true
}

// HasProtocolId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasProtocolId() bool {
	return o != nil && o.ProtocolId != nil
}

// SetProtocolId gets a reference to the given string and assigns it to the ProtocolId field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetProtocolId(v string) {
	o.ProtocolId = &v
}

// GetRelatedLinks returns the RelatedLinks field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetRelatedLinks() []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems {
	if o == nil || o.RelatedLinks == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems
		return ret
	}
	return o.RelatedLinks
}

// GetRelatedLinksOk returns a tuple with the RelatedLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetRelatedLinksOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems, bool) {
	if o == nil || o.RelatedLinks == nil {
		return nil, false
	}
	return &o.RelatedLinks, true
}

// HasRelatedLinks returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasRelatedLinks() bool {
	return o != nil && o.RelatedLinks != nil
}

// SetRelatedLinks gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems and assigns it to the RelatedLinks field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetRelatedLinks(v []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) {
	o.RelatedLinks = v
}

// GetResultsLastUpdated returns the ResultsLastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetResultsLastUpdated() time.Time {
	if o == nil || o.ResultsLastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResultsLastUpdated.Get()
}

// GetResultsLastUpdatedOk returns a tuple with the ResultsLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetResultsLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsLastUpdated.Get(), o.ResultsLastUpdated.IsSet()
}

// HasResultsLastUpdated returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasResultsLastUpdated() bool {
	return o != nil && o.ResultsLastUpdated.IsSet()
}

// SetResultsLastUpdated gets a reference to the given datadog.NullableTime and assigns it to the ResultsLastUpdated field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetResultsLastUpdated(v time.Time) {
	o.ResultsLastUpdated.Set(&v)
}

// SetResultsLastUpdatedNil sets the value for ResultsLastUpdated to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetResultsLastUpdatedNil() {
	o.ResultsLastUpdated.Set(nil)
}

// UnsetResultsLastUpdated ensures that no value is present for ResultsLastUpdated, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetResultsLastUpdated() {
	o.ResultsLastUpdated.Unset()
}

// GetSplitByProperties returns the SplitByProperties field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSplitByProperties() []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems {
	if o == nil || o.SplitByProperties == nil {
		var ret []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems
		return ret
	}
	return o.SplitByProperties
}

// GetSplitByPropertiesOk returns a tuple with the SplitByProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSplitByPropertiesOk() (*[]ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems, bool) {
	if o == nil || o.SplitByProperties == nil {
		return nil, false
	}
	return &o.SplitByProperties, true
}

// HasSplitByProperties returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasSplitByProperties() bool {
	return o != nil && o.SplitByProperties != nil
}

// SetSplitByProperties gets a reference to the given []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems and assigns it to the SplitByProperties field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetSplitByProperties(v []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems) {
	o.SplitByProperties = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetStatus() ExperimentsExperimentV2DTODataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret ExperimentsExperimentV2DTODataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetStatusOk() (*ExperimentsExperimentV2DTODataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given ExperimentsExperimentV2DTODataAttributesStatus and assigns it to the Status field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetStatus(v ExperimentsExperimentV2DTODataAttributesStatus) {
	o.Status = &v
}

// GetStructuredMetadata returns the StructuredMetadata field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetStructuredMetadata() []ExperimentsStructuredMetadataResponse {
	if o == nil || o.StructuredMetadata == nil {
		var ret []ExperimentsStructuredMetadataResponse
		return ret
	}
	return o.StructuredMetadata
}

// GetStructuredMetadataOk returns a tuple with the StructuredMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetStructuredMetadataOk() (*[]ExperimentsStructuredMetadataResponse, bool) {
	if o == nil || o.StructuredMetadata == nil {
		return nil, false
	}
	return &o.StructuredMetadata, true
}

// HasStructuredMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasStructuredMetadata() bool {
	return o != nil && o.StructuredMetadata != nil
}

// SetStructuredMetadata gets a reference to the given []ExperimentsStructuredMetadataResponse and assigns it to the StructuredMetadata field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetStructuredMetadata(v []ExperimentsStructuredMetadataResponse) {
	o.StructuredMetadata = v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId gets a reference to the given datadog.NullableString and assigns it to the SubjectTypeId field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId.Set(&v)
}

// SetSubjectTypeIdNil sets the value for SubjectTypeId to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetSubjectTypeIdNil() {
	o.SubjectTypeId.Set(nil)
}

// UnsetSubjectTypeId ensures that no value is present for SubjectTypeId, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetSubjectTypeId() {
	o.SubjectTypeId.Unset()
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetSummary(v string) {
	o.Summary = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetTrafficExposure returns the TrafficExposure field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTrafficExposure() ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure {
	if o == nil || o.TrafficExposure == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure
		return ret
	}
	return *o.TrafficExposure
}

// GetTrafficExposureOk returns a tuple with the TrafficExposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetTrafficExposureOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure, bool) {
	if o == nil || o.TrafficExposure == nil {
		return nil, false
	}
	return o.TrafficExposure, true
}

// HasTrafficExposure returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasTrafficExposure() bool {
	return o != nil && o.TrafficExposure != nil
}

// SetTrafficExposure gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure and assigns it to the TrafficExposure field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetTrafficExposure(v ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) {
	o.TrafficExposure = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetVariants() []ExperimentsExperimentV2DTODataAttributesVariantsItems {
	if o == nil || o.Variants == nil {
		var ret []ExperimentsExperimentV2DTODataAttributesVariantsItems
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetVariantsOk() (*[]ExperimentsExperimentV2DTODataAttributesVariantsItems, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []ExperimentsExperimentV2DTODataAttributesVariantsItems and assigns it to the Variants field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetVariants(v []ExperimentsExperimentV2DTODataAttributesVariantsItems) {
	o.Variants = v
}

// GetWarehouseExposureConfiguration returns the WarehouseExposureConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetWarehouseExposureConfiguration() ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration {
	if o == nil || o.WarehouseExposureConfiguration.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration
		return ret
	}
	return *o.WarehouseExposureConfiguration.Get()
}

// GetWarehouseExposureConfigurationOk returns a tuple with the WarehouseExposureConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) GetWarehouseExposureConfigurationOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehouseExposureConfiguration.Get(), o.WarehouseExposureConfiguration.IsSet()
}

// HasWarehouseExposureConfiguration returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) HasWarehouseExposureConfiguration() bool {
	return o != nil && o.WarehouseExposureConfiguration.IsSet()
}

// SetWarehouseExposureConfiguration gets a reference to the given NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration and assigns it to the WarehouseExposureConfiguration field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetWarehouseExposureConfiguration(v ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration) {
	o.WarehouseExposureConfiguration.Set(&v)
}

// SetWarehouseExposureConfigurationNil sets the value for WarehouseExposureConfiguration to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) SetWarehouseExposureConfigurationNil() {
	o.WarehouseExposureConfiguration.Set(nil)
}

// UnsetWarehouseExposureConfiguration ensures that no value is present for WarehouseExposureConfiguration, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnsetWarehouseExposureConfiguration() {
	o.WarehouseExposureConfiguration.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentsEndDate.IsSet() {
		toSerialize["assignments_end_date"] = o.AssignmentsEndDate.Get()
	}
	if o.AssignmentsStartDate.IsSet() {
		toSerialize["assignments_start_date"] = o.AssignmentsStartDate.Get()
	}
	if o.ConcludedAt.IsSet() {
		toSerialize["concluded_at"] = o.ConcludedAt.Get()
	}
	if o.Conclusion != nil {
		toSerialize["conclusion"] = o.Conclusion
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DatadogFlagConfiguration.IsSet() {
		toSerialize["datadog_flag_configuration"] = o.DatadogFlagConfiguration.Get()
	}
	if o.DecisionMetrics != nil {
		toSerialize["decision_metrics"] = o.DecisionMetrics
	}
	if o.DecisionVariantKey != nil {
		toSerialize["decision_variant_key"] = o.DecisionVariantKey
	}
	if o.EventsEndDate.IsSet() {
		toSerialize["events_end_date"] = o.EventsEndDate.Get()
	}
	if o.EventsStartDate.IsSet() {
		toSerialize["events_start_date"] = o.EventsStartDate.Get()
	}
	if o.ExperimentType != nil {
		toSerialize["experiment_type"] = o.ExperimentType
	}
	if o.Hypothesis != nil {
		toSerialize["hypothesis"] = o.Hypothesis
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PipelineTableSuffix != nil {
		toSerialize["pipeline_table_suffix"] = o.PipelineTableSuffix
	}
	if o.ProtocolId != nil {
		toSerialize["protocol_id"] = o.ProtocolId
	}
	if o.RelatedLinks != nil {
		toSerialize["related_links"] = o.RelatedLinks
	}
	if o.ResultsLastUpdated.IsSet() {
		toSerialize["results_last_updated"] = o.ResultsLastUpdated.Get()
	}
	if o.SplitByProperties != nil {
		toSerialize["split_by_properties"] = o.SplitByProperties
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StructuredMetadata != nil {
		toSerialize["structured_metadata"] = o.StructuredMetadata
	}
	if o.SubjectTypeId.IsSet() {
		toSerialize["subject_type_id"] = o.SubjectTypeId.Get()
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.TrafficExposure != nil {
		toSerialize["traffic_exposure"] = o.TrafficExposure
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.WarehouseExposureConfiguration.IsSet() {
		toSerialize["warehouse_exposure_configuration"] = o.WarehouseExposureConfiguration.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentsEndDate             datadog.NullableTime                                                                     `json:"assignments_end_date,omitempty"`
		AssignmentsStartDate           datadog.NullableTime                                                                     `json:"assignments_start_date,omitempty"`
		ConcludedAt                    datadog.NullableTime                                                                     `json:"concluded_at,omitempty"`
		Conclusion                     *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion                            `json:"conclusion,omitempty"`
		CreatedAt                      *time.Time                                                                               `json:"created_at,omitempty"`
		DatadogFlagConfiguration       NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration       `json:"datadog_flag_configuration,omitempty"`
		DecisionMetrics                []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems                 `json:"decision_metrics,omitempty"`
		DecisionVariantKey             *string                                                                                  `json:"decision_variant_key,omitempty"`
		EventsEndDate                  datadog.NullableTime                                                                     `json:"events_end_date,omitempty"`
		EventsStartDate                datadog.NullableTime                                                                     `json:"events_start_date,omitempty"`
		ExperimentType                 *string                                                                                  `json:"experiment_type,omitempty"`
		Hypothesis                     *string                                                                                  `json:"hypothesis,omitempty"`
		MigrationMetadata              interface{}                                                                              `json:"migration_metadata,omitempty"`
		Name                           *string                                                                                  `json:"name,omitempty"`
		PipelineTableSuffix            *string                                                                                  `json:"pipeline_table_suffix,omitempty"`
		ProtocolId                     *string                                                                                  `json:"protocol_id,omitempty"`
		RelatedLinks                   []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems                    `json:"related_links,omitempty"`
		ResultsLastUpdated             datadog.NullableTime                                                                     `json:"results_last_updated,omitempty"`
		SplitByProperties              []ExperimentsExperimentV2DTODataAttributesSplitByPropertiesItems                         `json:"split_by_properties,omitempty"`
		Status                         *ExperimentsExperimentV2DTODataAttributesStatus                                          `json:"status,omitempty"`
		StructuredMetadata             []ExperimentsStructuredMetadataResponse                                                  `json:"structured_metadata,omitempty"`
		SubjectTypeId                  datadog.NullableString                                                                   `json:"subject_type_id,omitempty"`
		Summary                        *string                                                                                  `json:"summary,omitempty"`
		Tags                           []string                                                                                 `json:"tags,omitempty"`
		Teams                          []string                                                                                 `json:"teams,omitempty"`
		TrafficExposure                *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure                       `json:"traffic_exposure,omitempty"`
		UpdatedAt                      *time.Time                                                                               `json:"updated_at,omitempty"`
		Variants                       []ExperimentsExperimentV2DTODataAttributesVariantsItems                                  `json:"variants,omitempty"`
		WarehouseExposureConfiguration NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfiguration `json:"warehouse_exposure_configuration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignments_end_date", "assignments_start_date", "concluded_at", "conclusion", "created_at", "datadog_flag_configuration", "decision_metrics", "decision_variant_key", "events_end_date", "events_start_date", "experiment_type", "hypothesis", "migration_metadata", "name", "pipeline_table_suffix", "protocol_id", "related_links", "results_last_updated", "split_by_properties", "status", "structured_metadata", "subject_type_id", "summary", "tags", "teams", "traffic_exposure", "updated_at", "variants", "warehouse_exposure_configuration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignmentsEndDate = all.AssignmentsEndDate
	o.AssignmentsStartDate = all.AssignmentsStartDate
	o.ConcludedAt = all.ConcludedAt
	if all.Conclusion != nil && all.Conclusion.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Conclusion = all.Conclusion
	o.CreatedAt = all.CreatedAt
	o.DatadogFlagConfiguration = all.DatadogFlagConfiguration
	o.DecisionMetrics = all.DecisionMetrics
	o.DecisionVariantKey = all.DecisionVariantKey
	o.EventsEndDate = all.EventsEndDate
	o.EventsStartDate = all.EventsStartDate
	o.ExperimentType = all.ExperimentType
	o.Hypothesis = all.Hypothesis
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.PipelineTableSuffix = all.PipelineTableSuffix
	o.ProtocolId = all.ProtocolId
	o.RelatedLinks = all.RelatedLinks
	o.ResultsLastUpdated = all.ResultsLastUpdated
	o.SplitByProperties = all.SplitByProperties
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.StructuredMetadata = all.StructuredMetadata
	o.SubjectTypeId = all.SubjectTypeId
	o.Summary = all.Summary
	o.Tags = all.Tags
	o.Teams = all.Teams
	if all.TrafficExposure != nil && all.TrafficExposure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrafficExposure = all.TrafficExposure
	o.UpdatedAt = all.UpdatedAt
	o.Variants = all.Variants
	o.WarehouseExposureConfiguration = all.WarehouseExposureConfiguration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
