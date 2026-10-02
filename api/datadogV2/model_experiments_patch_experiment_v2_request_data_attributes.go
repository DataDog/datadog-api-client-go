// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2RequestDataAttributes Fields supplied to update the experiment.
type ExperimentsPatchExperimentV2RequestDataAttributes struct {
	// End of the time window for experiment assignments.
	AssignmentsEndDate datadog.NullableTime `json:"assignments_end_date,omitempty"`
	// Start of the time window for experiment assignments.
	AssignmentsStartDate datadog.NullableTime `json:"assignments_start_date,omitempty"`
	// Feature flag, environment, and targeting configuration for the experiment.
	DatadogFlagConfiguration NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration `json:"datadog_flag_configuration,omitempty"`
	// Metrics used to decide the experiment outcome.
	DecisionMetrics []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems `json:"decision_metrics,omitempty"`
	// End of the time window for metric events.
	EventsEndDate datadog.NullableTime `json:"events_end_date,omitempty"`
	// Start of the time window for metric events.
	EventsStartDate datadog.NullableTime `json:"events_start_date,omitempty"`
	// Expected effect that the experiment is intended to test.
	Hypothesis *string `json:"hypothesis,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment.
	Name *string `json:"name,omitempty"`
	// Links to supporting material for the experiment.
	RelatedLinks []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems `json:"related_links,omitempty"`
	// Complete replacement for the Datadog split-by selection. Identify each property by column_name. Omit this field to leave the selection unchanged.
	SplitByProperties []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems `json:"split_by_properties,omitempty"`
	// Metadata values to update. Omit this field, send null, or send an empty array to leave metadata unchanged.
	StructuredMetadata []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems `json:"structured_metadata,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId *string `json:"subject_type_id,omitempty"`
	// Summary text recorded for the experiment.
	Summary datadog.NullableString `json:"summary,omitempty"`
	// Tags attached to the experiment.
	Tags []string `json:"tags,omitempty"`
	// Teams associated with the experiment.
	Teams []string `json:"teams,omitempty"`
	// Traffic exposure fraction or schedule configured for the experiment.
	TrafficExposure *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure `json:"traffic_exposure,omitempty"`
	// Variants configured for the experiment.
	Variants []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems `json:"variants,omitempty"`
	// Warehouse model and experiment key used to read assignment data.
	WarehouseExposureConfiguration NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration `json:"warehouse_exposure_configuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2RequestDataAttributes instantiates a new ExperimentsPatchExperimentV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2RequestDataAttributes() *ExperimentsPatchExperimentV2RequestDataAttributes {
	this := ExperimentsPatchExperimentV2RequestDataAttributes{}
	return &this
}

// NewExperimentsPatchExperimentV2RequestDataAttributesWithDefaults instantiates a new ExperimentsPatchExperimentV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2RequestDataAttributesWithDefaults() *ExperimentsPatchExperimentV2RequestDataAttributes {
	this := ExperimentsPatchExperimentV2RequestDataAttributes{}
	return &this
}

// GetAssignmentsEndDate returns the AssignmentsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetAssignmentsEndDate() time.Time {
	if o == nil || o.AssignmentsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsEndDate.Get()
}

// GetAssignmentsEndDateOk returns a tuple with the AssignmentsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetAssignmentsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsEndDate.Get(), o.AssignmentsEndDate.IsSet()
}

// HasAssignmentsEndDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasAssignmentsEndDate() bool {
	return o != nil && o.AssignmentsEndDate.IsSet()
}

// SetAssignmentsEndDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsEndDate field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetAssignmentsEndDate(v time.Time) {
	o.AssignmentsEndDate.Set(&v)
}

// SetAssignmentsEndDateNil sets the value for AssignmentsEndDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetAssignmentsEndDateNil() {
	o.AssignmentsEndDate.Set(nil)
}

// UnsetAssignmentsEndDate ensures that no value is present for AssignmentsEndDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetAssignmentsEndDate() {
	o.AssignmentsEndDate.Unset()
}

// GetAssignmentsStartDate returns the AssignmentsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetAssignmentsStartDate() time.Time {
	if o == nil || o.AssignmentsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsStartDate.Get()
}

// GetAssignmentsStartDateOk returns a tuple with the AssignmentsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetAssignmentsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsStartDate.Get(), o.AssignmentsStartDate.IsSet()
}

// HasAssignmentsStartDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasAssignmentsStartDate() bool {
	return o != nil && o.AssignmentsStartDate.IsSet()
}

// SetAssignmentsStartDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsStartDate field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetAssignmentsStartDate(v time.Time) {
	o.AssignmentsStartDate.Set(&v)
}

// SetAssignmentsStartDateNil sets the value for AssignmentsStartDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetAssignmentsStartDateNil() {
	o.AssignmentsStartDate.Set(nil)
}

// UnsetAssignmentsStartDate ensures that no value is present for AssignmentsStartDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetAssignmentsStartDate() {
	o.AssignmentsStartDate.Unset()
}

// GetDatadogFlagConfiguration returns the DatadogFlagConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetDatadogFlagConfiguration() ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	if o == nil || o.DatadogFlagConfiguration.Get() == nil {
		var ret ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration
		return ret
	}
	return *o.DatadogFlagConfiguration.Get()
}

// GetDatadogFlagConfigurationOk returns a tuple with the DatadogFlagConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetDatadogFlagConfigurationOk() (*ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogFlagConfiguration.Get(), o.DatadogFlagConfiguration.IsSet()
}

// HasDatadogFlagConfiguration returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasDatadogFlagConfiguration() bool {
	return o != nil && o.DatadogFlagConfiguration.IsSet()
}

// SetDatadogFlagConfiguration gets a reference to the given NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration and assigns it to the DatadogFlagConfiguration field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetDatadogFlagConfiguration(v ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) {
	o.DatadogFlagConfiguration.Set(&v)
}

// SetDatadogFlagConfigurationNil sets the value for DatadogFlagConfiguration to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetDatadogFlagConfigurationNil() {
	o.DatadogFlagConfiguration.Set(nil)
}

// UnsetDatadogFlagConfiguration ensures that no value is present for DatadogFlagConfiguration, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetDatadogFlagConfiguration() {
	o.DatadogFlagConfiguration.Unset()
}

// GetDecisionMetrics returns the DecisionMetrics field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetDecisionMetrics() []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems {
	if o == nil || o.DecisionMetrics == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems
		return ret
	}
	return o.DecisionMetrics
}

// GetDecisionMetricsOk returns a tuple with the DecisionMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetDecisionMetricsOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems, bool) {
	if o == nil || o.DecisionMetrics == nil {
		return nil, false
	}
	return &o.DecisionMetrics, true
}

// HasDecisionMetrics returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasDecisionMetrics() bool {
	return o != nil && o.DecisionMetrics != nil
}

// SetDecisionMetrics gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems and assigns it to the DecisionMetrics field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetDecisionMetrics(v []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) {
	o.DecisionMetrics = v
}

// GetEventsEndDate returns the EventsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetEventsEndDate() time.Time {
	if o == nil || o.EventsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsEndDate.Get()
}

// GetEventsEndDateOk returns a tuple with the EventsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetEventsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsEndDate.Get(), o.EventsEndDate.IsSet()
}

// HasEventsEndDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasEventsEndDate() bool {
	return o != nil && o.EventsEndDate.IsSet()
}

// SetEventsEndDate gets a reference to the given datadog.NullableTime and assigns it to the EventsEndDate field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetEventsEndDate(v time.Time) {
	o.EventsEndDate.Set(&v)
}

// SetEventsEndDateNil sets the value for EventsEndDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetEventsEndDateNil() {
	o.EventsEndDate.Set(nil)
}

// UnsetEventsEndDate ensures that no value is present for EventsEndDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetEventsEndDate() {
	o.EventsEndDate.Unset()
}

// GetEventsStartDate returns the EventsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetEventsStartDate() time.Time {
	if o == nil || o.EventsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsStartDate.Get()
}

// GetEventsStartDateOk returns a tuple with the EventsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetEventsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsStartDate.Get(), o.EventsStartDate.IsSet()
}

// HasEventsStartDate returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasEventsStartDate() bool {
	return o != nil && o.EventsStartDate.IsSet()
}

// SetEventsStartDate gets a reference to the given datadog.NullableTime and assigns it to the EventsStartDate field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetEventsStartDate(v time.Time) {
	o.EventsStartDate.Set(&v)
}

// SetEventsStartDateNil sets the value for EventsStartDate to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetEventsStartDateNil() {
	o.EventsStartDate.Set(nil)
}

// UnsetEventsStartDate ensures that no value is present for EventsStartDate, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetEventsStartDate() {
	o.EventsStartDate.Unset()
}

// GetHypothesis returns the Hypothesis field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetHypothesis() string {
	if o == nil || o.Hypothesis == nil {
		var ret string
		return ret
	}
	return *o.Hypothesis
}

// GetHypothesisOk returns a tuple with the Hypothesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetHypothesisOk() (*string, bool) {
	if o == nil || o.Hypothesis == nil {
		return nil, false
	}
	return o.Hypothesis, true
}

// HasHypothesis returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasHypothesis() bool {
	return o != nil && o.Hypothesis != nil
}

// SetHypothesis gets a reference to the given string and assigns it to the Hypothesis field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetHypothesis(v string) {
	o.Hypothesis = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRelatedLinks returns the RelatedLinks field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetRelatedLinks() []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems {
	if o == nil || o.RelatedLinks == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems
		return ret
	}
	return o.RelatedLinks
}

// GetRelatedLinksOk returns a tuple with the RelatedLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetRelatedLinksOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems, bool) {
	if o == nil || o.RelatedLinks == nil {
		return nil, false
	}
	return &o.RelatedLinks, true
}

// HasRelatedLinks returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasRelatedLinks() bool {
	return o != nil && o.RelatedLinks != nil
}

// SetRelatedLinks gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems and assigns it to the RelatedLinks field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetRelatedLinks(v []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) {
	o.RelatedLinks = v
}

// GetSplitByProperties returns the SplitByProperties field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSplitByProperties() []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems {
	if o == nil || o.SplitByProperties == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems
		return ret
	}
	return o.SplitByProperties
}

// GetSplitByPropertiesOk returns a tuple with the SplitByProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSplitByPropertiesOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems, bool) {
	if o == nil || o.SplitByProperties == nil {
		return nil, false
	}
	return &o.SplitByProperties, true
}

// HasSplitByProperties returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasSplitByProperties() bool {
	return o != nil && o.SplitByProperties != nil
}

// SetSplitByProperties gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems and assigns it to the SplitByProperties field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetSplitByProperties(v []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems) {
	o.SplitByProperties = v
}

// GetStructuredMetadata returns the StructuredMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetStructuredMetadata() []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems {
	if o == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems
		return ret
	}
	return o.StructuredMetadata
}

// GetStructuredMetadataOk returns a tuple with the StructuredMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetStructuredMetadataOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems, bool) {
	if o == nil || o.StructuredMetadata == nil {
		return nil, false
	}
	return &o.StructuredMetadata, true
}

// HasStructuredMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasStructuredMetadata() bool {
	return o != nil && o.StructuredMetadata != nil
}

// SetStructuredMetadata gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems and assigns it to the StructuredMetadata field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetStructuredMetadata(v []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) {
	o.StructuredMetadata = v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil || o.SubjectTypeId == nil {
		return nil, false
	}
	return o.SubjectTypeId, true
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId != nil
}

// SetSubjectTypeId gets a reference to the given string and assigns it to the SubjectTypeId field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasSummary() bool {
	return o != nil && o.Summary.IsSet()
}

// SetSummary gets a reference to the given datadog.NullableString and assigns it to the Summary field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetSummary() {
	o.Summary.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetTrafficExposure returns the TrafficExposure field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTrafficExposure() ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure {
	if o == nil || o.TrafficExposure == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure
		return ret
	}
	return *o.TrafficExposure
}

// GetTrafficExposureOk returns a tuple with the TrafficExposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetTrafficExposureOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure, bool) {
	if o == nil || o.TrafficExposure == nil {
		return nil, false
	}
	return o.TrafficExposure, true
}

// HasTrafficExposure returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasTrafficExposure() bool {
	return o != nil && o.TrafficExposure != nil
}

// SetTrafficExposure gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure and assigns it to the TrafficExposure field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetTrafficExposure(v ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) {
	o.TrafficExposure = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetVariants() []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems {
	if o == nil || o.Variants == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetVariantsOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems and assigns it to the Variants field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetVariants(v []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) {
	o.Variants = v
}

// GetWarehouseExposureConfiguration returns the WarehouseExposureConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetWarehouseExposureConfiguration() ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration {
	if o == nil || o.WarehouseExposureConfiguration.Get() == nil {
		var ret ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration
		return ret
	}
	return *o.WarehouseExposureConfiguration.Get()
}

// GetWarehouseExposureConfigurationOk returns a tuple with the WarehouseExposureConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) GetWarehouseExposureConfigurationOk() (*ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehouseExposureConfiguration.Get(), o.WarehouseExposureConfiguration.IsSet()
}

// HasWarehouseExposureConfiguration returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) HasWarehouseExposureConfiguration() bool {
	return o != nil && o.WarehouseExposureConfiguration.IsSet()
}

// SetWarehouseExposureConfiguration gets a reference to the given NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration and assigns it to the WarehouseExposureConfiguration field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetWarehouseExposureConfiguration(v ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) {
	o.WarehouseExposureConfiguration.Set(&v)
}

// SetWarehouseExposureConfigurationNil sets the value for WarehouseExposureConfiguration to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) SetWarehouseExposureConfigurationNil() {
	o.WarehouseExposureConfiguration.Set(nil)
}

// UnsetWarehouseExposureConfiguration ensures that no value is present for WarehouseExposureConfiguration, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnsetWarehouseExposureConfiguration() {
	o.WarehouseExposureConfiguration.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.DatadogFlagConfiguration.IsSet() {
		toSerialize["datadog_flag_configuration"] = o.DatadogFlagConfiguration.Get()
	}
	if o.DecisionMetrics != nil {
		toSerialize["decision_metrics"] = o.DecisionMetrics
	}
	if o.EventsEndDate.IsSet() {
		toSerialize["events_end_date"] = o.EventsEndDate.Get()
	}
	if o.EventsStartDate.IsSet() {
		toSerialize["events_start_date"] = o.EventsStartDate.Get()
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
	if o.RelatedLinks != nil {
		toSerialize["related_links"] = o.RelatedLinks
	}
	if o.SplitByProperties != nil {
		toSerialize["split_by_properties"] = o.SplitByProperties
	}
	if o.StructuredMetadata != nil {
		toSerialize["structured_metadata"] = o.StructuredMetadata
	}
	if o.SubjectTypeId != nil {
		toSerialize["subject_type_id"] = o.SubjectTypeId
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
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
func (o *ExperimentsPatchExperimentV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentsEndDate             datadog.NullableTime                                                                     `json:"assignments_end_date,omitempty"`
		AssignmentsStartDate           datadog.NullableTime                                                                     `json:"assignments_start_date,omitempty"`
		DatadogFlagConfiguration       NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration        `json:"datadog_flag_configuration,omitempty"`
		DecisionMetrics                []ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems                 `json:"decision_metrics,omitempty"`
		EventsEndDate                  datadog.NullableTime                                                                     `json:"events_end_date,omitempty"`
		EventsStartDate                datadog.NullableTime                                                                     `json:"events_start_date,omitempty"`
		Hypothesis                     *string                                                                                  `json:"hypothesis,omitempty"`
		MigrationMetadata              interface{}                                                                              `json:"migration_metadata,omitempty"`
		Name                           *string                                                                                  `json:"name,omitempty"`
		RelatedLinks                   []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems                    `json:"related_links,omitempty"`
		SplitByProperties              []ExperimentsCreateExperimentV2RequestDataAttributesSplitByPropertiesItems               `json:"split_by_properties,omitempty"`
		StructuredMetadata             []ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems              `json:"structured_metadata,omitempty"`
		SubjectTypeId                  *string                                                                                  `json:"subject_type_id,omitempty"`
		Summary                        datadog.NullableString                                                                   `json:"summary,omitempty"`
		Tags                           []string                                                                                 `json:"tags,omitempty"`
		Teams                          []string                                                                                 `json:"teams,omitempty"`
		TrafficExposure                *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure                       `json:"traffic_exposure,omitempty"`
		Variants                       []ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems                        `json:"variants,omitempty"`
		WarehouseExposureConfiguration NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration `json:"warehouse_exposure_configuration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignments_end_date", "assignments_start_date", "datadog_flag_configuration", "decision_metrics", "events_end_date", "events_start_date", "hypothesis", "migration_metadata", "name", "related_links", "split_by_properties", "structured_metadata", "subject_type_id", "summary", "tags", "teams", "traffic_exposure", "variants", "warehouse_exposure_configuration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignmentsEndDate = all.AssignmentsEndDate
	o.AssignmentsStartDate = all.AssignmentsStartDate
	o.DatadogFlagConfiguration = all.DatadogFlagConfiguration
	o.DecisionMetrics = all.DecisionMetrics
	o.EventsEndDate = all.EventsEndDate
	o.EventsStartDate = all.EventsStartDate
	o.Hypothesis = all.Hypothesis
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.RelatedLinks = all.RelatedLinks
	o.SplitByProperties = all.SplitByProperties
	o.StructuredMetadata = all.StructuredMetadata
	o.SubjectTypeId = all.SubjectTypeId
	o.Summary = all.Summary
	o.Tags = all.Tags
	o.Teams = all.Teams
	if all.TrafficExposure != nil && all.TrafficExposure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrafficExposure = all.TrafficExposure
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
