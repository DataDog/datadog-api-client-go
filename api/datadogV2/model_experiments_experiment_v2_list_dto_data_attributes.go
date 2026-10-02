// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2ListDTODataAttributes Summary fields for an experiment returned in a list.
type ExperimentsExperimentV2ListDTODataAttributes struct {
	// End of the window used to read experiment assignments.
	AssignmentsEndDate datadog.NullableTime `json:"assignments_end_date,omitempty"`
	// Start of the window used to read experiment assignments.
	AssignmentsStartDate datadog.NullableTime `json:"assignments_start_date,omitempty"`
	// Time when the experiment was concluded.
	ConcludedAt datadog.NullableTime `json:"concluded_at,omitempty"`
	// Outcome and supporting text recorded when the experiment is concluded.
	Conclusion *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion `json:"conclusion,omitempty"`
	// Time when the experiment was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// End of the window used to read metric events.
	EventsEndDate datadog.NullableTime `json:"events_end_date,omitempty"`
	// Start of the window used to read metric events.
	EventsStartDate datadog.NullableTime `json:"events_start_date,omitempty"`
	// Kind of experiment. STANDARD is an ordinary experiment. Other values, such as CANARY and HOLDOUT, identify experiments owned by another workflow. New kinds may be added; treat unknown values as non-standard.
	ExperimentType *string `json:"experiment_type,omitempty"`
	// Expected effect that the experiment is designed to test.
	Hypothesis *string `json:"hypothesis,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment.
	Name *string `json:"name,omitempty"`
	// Suffix used for the experiment tables in the analysis pipeline.
	PipelineTableSuffix *string `json:"pipeline_table_suffix,omitempty"`
	// Identifier of the protocol associated with the experiment.
	ProtocolId *string `json:"protocol_id,omitempty"`
	// External links associated with the experiment.
	RelatedLinks []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems `json:"related_links,omitempty"`
	// Time when the experiment results were last updated.
	ResultsLastUpdated datadog.NullableTime `json:"results_last_updated,omitempty"`
	// Current stage in the experiment lifecycle.
	Status *ExperimentsExperimentV2DTODataAttributesStatus `json:"status,omitempty"`
	// Custom metadata fields and their values for the experiment.
	StructuredMetadata []ExperimentsStructuredMetadataResponse `json:"structured_metadata,omitempty"`
	// Identifier of the subject type used for experiment assignments.
	SubjectTypeId datadog.NullableString `json:"subject_type_id,omitempty"`
	// Free-form summary of the experiment.
	Summary *string `json:"summary,omitempty"`
	// Tag names associated with the experiment.
	Tags []string `json:"tags,omitempty"`
	// Team handles associated with the experiment.
	Teams []string `json:"teams,omitempty"`
	// Time when the experiment was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentV2ListDTODataAttributes instantiates a new ExperimentsExperimentV2ListDTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentV2ListDTODataAttributes() *ExperimentsExperimentV2ListDTODataAttributes {
	this := ExperimentsExperimentV2ListDTODataAttributes{}
	return &this
}

// NewExperimentsExperimentV2ListDTODataAttributesWithDefaults instantiates a new ExperimentsExperimentV2ListDTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentV2ListDTODataAttributesWithDefaults() *ExperimentsExperimentV2ListDTODataAttributes {
	this := ExperimentsExperimentV2ListDTODataAttributes{}
	return &this
}

// GetAssignmentsEndDate returns the AssignmentsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetAssignmentsEndDate() time.Time {
	if o == nil || o.AssignmentsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsEndDate.Get()
}

// GetAssignmentsEndDateOk returns a tuple with the AssignmentsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetAssignmentsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsEndDate.Get(), o.AssignmentsEndDate.IsSet()
}

// HasAssignmentsEndDate returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasAssignmentsEndDate() bool {
	return o != nil && o.AssignmentsEndDate.IsSet()
}

// SetAssignmentsEndDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsEndDate field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetAssignmentsEndDate(v time.Time) {
	o.AssignmentsEndDate.Set(&v)
}

// SetAssignmentsEndDateNil sets the value for AssignmentsEndDate to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetAssignmentsEndDateNil() {
	o.AssignmentsEndDate.Set(nil)
}

// UnsetAssignmentsEndDate ensures that no value is present for AssignmentsEndDate, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetAssignmentsEndDate() {
	o.AssignmentsEndDate.Unset()
}

// GetAssignmentsStartDate returns the AssignmentsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetAssignmentsStartDate() time.Time {
	if o == nil || o.AssignmentsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignmentsStartDate.Get()
}

// GetAssignmentsStartDateOk returns a tuple with the AssignmentsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetAssignmentsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignmentsStartDate.Get(), o.AssignmentsStartDate.IsSet()
}

// HasAssignmentsStartDate returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasAssignmentsStartDate() bool {
	return o != nil && o.AssignmentsStartDate.IsSet()
}

// SetAssignmentsStartDate gets a reference to the given datadog.NullableTime and assigns it to the AssignmentsStartDate field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetAssignmentsStartDate(v time.Time) {
	o.AssignmentsStartDate.Set(&v)
}

// SetAssignmentsStartDateNil sets the value for AssignmentsStartDate to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetAssignmentsStartDateNil() {
	o.AssignmentsStartDate.Set(nil)
}

// UnsetAssignmentsStartDate ensures that no value is present for AssignmentsStartDate, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetAssignmentsStartDate() {
	o.AssignmentsStartDate.Unset()
}

// GetConcludedAt returns the ConcludedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetConcludedAt() time.Time {
	if o == nil || o.ConcludedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ConcludedAt.Get()
}

// GetConcludedAtOk returns a tuple with the ConcludedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetConcludedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConcludedAt.Get(), o.ConcludedAt.IsSet()
}

// HasConcludedAt returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasConcludedAt() bool {
	return o != nil && o.ConcludedAt.IsSet()
}

// SetConcludedAt gets a reference to the given datadog.NullableTime and assigns it to the ConcludedAt field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetConcludedAt(v time.Time) {
	o.ConcludedAt.Set(&v)
}

// SetConcludedAtNil sets the value for ConcludedAt to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetConcludedAtNil() {
	o.ConcludedAt.Set(nil)
}

// UnsetConcludedAt ensures that no value is present for ConcludedAt, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetConcludedAt() {
	o.ConcludedAt.Unset()
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetConclusion() ExperimentsPatchExperimentV2ResponseDataAttributesConclusion {
	if o == nil || o.Conclusion == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesConclusion
		return ret
	}
	return *o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetConclusionOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesConclusion, bool) {
	if o == nil || o.Conclusion == nil {
		return nil, false
	}
	return o.Conclusion, true
}

// HasConclusion returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasConclusion() bool {
	return o != nil && o.Conclusion != nil
}

// SetConclusion gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesConclusion and assigns it to the Conclusion field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetConclusion(v ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) {
	o.Conclusion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEventsEndDate returns the EventsEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetEventsEndDate() time.Time {
	if o == nil || o.EventsEndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsEndDate.Get()
}

// GetEventsEndDateOk returns a tuple with the EventsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetEventsEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsEndDate.Get(), o.EventsEndDate.IsSet()
}

// HasEventsEndDate returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasEventsEndDate() bool {
	return o != nil && o.EventsEndDate.IsSet()
}

// SetEventsEndDate gets a reference to the given datadog.NullableTime and assigns it to the EventsEndDate field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetEventsEndDate(v time.Time) {
	o.EventsEndDate.Set(&v)
}

// SetEventsEndDateNil sets the value for EventsEndDate to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetEventsEndDateNil() {
	o.EventsEndDate.Set(nil)
}

// UnsetEventsEndDate ensures that no value is present for EventsEndDate, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetEventsEndDate() {
	o.EventsEndDate.Unset()
}

// GetEventsStartDate returns the EventsStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetEventsStartDate() time.Time {
	if o == nil || o.EventsStartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EventsStartDate.Get()
}

// GetEventsStartDateOk returns a tuple with the EventsStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetEventsStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsStartDate.Get(), o.EventsStartDate.IsSet()
}

// HasEventsStartDate returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasEventsStartDate() bool {
	return o != nil && o.EventsStartDate.IsSet()
}

// SetEventsStartDate gets a reference to the given datadog.NullableTime and assigns it to the EventsStartDate field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetEventsStartDate(v time.Time) {
	o.EventsStartDate.Set(&v)
}

// SetEventsStartDateNil sets the value for EventsStartDate to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetEventsStartDateNil() {
	o.EventsStartDate.Set(nil)
}

// UnsetEventsStartDate ensures that no value is present for EventsStartDate, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetEventsStartDate() {
	o.EventsStartDate.Unset()
}

// GetExperimentType returns the ExperimentType field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetExperimentType() string {
	if o == nil || o.ExperimentType == nil {
		var ret string
		return ret
	}
	return *o.ExperimentType
}

// GetExperimentTypeOk returns a tuple with the ExperimentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetExperimentTypeOk() (*string, bool) {
	if o == nil || o.ExperimentType == nil {
		return nil, false
	}
	return o.ExperimentType, true
}

// HasExperimentType returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasExperimentType() bool {
	return o != nil && o.ExperimentType != nil
}

// SetExperimentType gets a reference to the given string and assigns it to the ExperimentType field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetExperimentType(v string) {
	o.ExperimentType = &v
}

// GetHypothesis returns the Hypothesis field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetHypothesis() string {
	if o == nil || o.Hypothesis == nil {
		var ret string
		return ret
	}
	return *o.Hypothesis
}

// GetHypothesisOk returns a tuple with the Hypothesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetHypothesisOk() (*string, bool) {
	if o == nil || o.Hypothesis == nil {
		return nil, false
	}
	return o.Hypothesis, true
}

// HasHypothesis returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasHypothesis() bool {
	return o != nil && o.Hypothesis != nil
}

// SetHypothesis gets a reference to the given string and assigns it to the Hypothesis field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetHypothesis(v string) {
	o.Hypothesis = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPipelineTableSuffix returns the PipelineTableSuffix field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetPipelineTableSuffix() string {
	if o == nil || o.PipelineTableSuffix == nil {
		var ret string
		return ret
	}
	return *o.PipelineTableSuffix
}

// GetPipelineTableSuffixOk returns a tuple with the PipelineTableSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetPipelineTableSuffixOk() (*string, bool) {
	if o == nil || o.PipelineTableSuffix == nil {
		return nil, false
	}
	return o.PipelineTableSuffix, true
}

// HasPipelineTableSuffix returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasPipelineTableSuffix() bool {
	return o != nil && o.PipelineTableSuffix != nil
}

// SetPipelineTableSuffix gets a reference to the given string and assigns it to the PipelineTableSuffix field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetPipelineTableSuffix(v string) {
	o.PipelineTableSuffix = &v
}

// GetProtocolId returns the ProtocolId field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetProtocolId() string {
	if o == nil || o.ProtocolId == nil {
		var ret string
		return ret
	}
	return *o.ProtocolId
}

// GetProtocolIdOk returns a tuple with the ProtocolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetProtocolIdOk() (*string, bool) {
	if o == nil || o.ProtocolId == nil {
		return nil, false
	}
	return o.ProtocolId, true
}

// HasProtocolId returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasProtocolId() bool {
	return o != nil && o.ProtocolId != nil
}

// SetProtocolId gets a reference to the given string and assigns it to the ProtocolId field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetProtocolId(v string) {
	o.ProtocolId = &v
}

// GetRelatedLinks returns the RelatedLinks field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetRelatedLinks() []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems {
	if o == nil || o.RelatedLinks == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems
		return ret
	}
	return o.RelatedLinks
}

// GetRelatedLinksOk returns a tuple with the RelatedLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetRelatedLinksOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems, bool) {
	if o == nil || o.RelatedLinks == nil {
		return nil, false
	}
	return &o.RelatedLinks, true
}

// HasRelatedLinks returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasRelatedLinks() bool {
	return o != nil && o.RelatedLinks != nil
}

// SetRelatedLinks gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems and assigns it to the RelatedLinks field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetRelatedLinks(v []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) {
	o.RelatedLinks = v
}

// GetResultsLastUpdated returns the ResultsLastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetResultsLastUpdated() time.Time {
	if o == nil || o.ResultsLastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResultsLastUpdated.Get()
}

// GetResultsLastUpdatedOk returns a tuple with the ResultsLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetResultsLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsLastUpdated.Get(), o.ResultsLastUpdated.IsSet()
}

// HasResultsLastUpdated returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasResultsLastUpdated() bool {
	return o != nil && o.ResultsLastUpdated.IsSet()
}

// SetResultsLastUpdated gets a reference to the given datadog.NullableTime and assigns it to the ResultsLastUpdated field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetResultsLastUpdated(v time.Time) {
	o.ResultsLastUpdated.Set(&v)
}

// SetResultsLastUpdatedNil sets the value for ResultsLastUpdated to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetResultsLastUpdatedNil() {
	o.ResultsLastUpdated.Set(nil)
}

// UnsetResultsLastUpdated ensures that no value is present for ResultsLastUpdated, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetResultsLastUpdated() {
	o.ResultsLastUpdated.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetStatus() ExperimentsExperimentV2DTODataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret ExperimentsExperimentV2DTODataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetStatusOk() (*ExperimentsExperimentV2DTODataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given ExperimentsExperimentV2DTODataAttributesStatus and assigns it to the Status field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetStatus(v ExperimentsExperimentV2DTODataAttributesStatus) {
	o.Status = &v
}

// GetStructuredMetadata returns the StructuredMetadata field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetStructuredMetadata() []ExperimentsStructuredMetadataResponse {
	if o == nil || o.StructuredMetadata == nil {
		var ret []ExperimentsStructuredMetadataResponse
		return ret
	}
	return o.StructuredMetadata
}

// GetStructuredMetadataOk returns a tuple with the StructuredMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetStructuredMetadataOk() (*[]ExperimentsStructuredMetadataResponse, bool) {
	if o == nil || o.StructuredMetadata == nil {
		return nil, false
	}
	return &o.StructuredMetadata, true
}

// HasStructuredMetadata returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasStructuredMetadata() bool {
	return o != nil && o.StructuredMetadata != nil
}

// SetStructuredMetadata gets a reference to the given []ExperimentsStructuredMetadataResponse and assigns it to the StructuredMetadata field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetStructuredMetadata(v []ExperimentsStructuredMetadataResponse) {
	o.StructuredMetadata = v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId.Get()
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypeId.Get(), o.SubjectTypeId.IsSet()
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId.IsSet()
}

// SetSubjectTypeId gets a reference to the given datadog.NullableString and assigns it to the SubjectTypeId field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId.Set(&v)
}

// SetSubjectTypeIdNil sets the value for SubjectTypeId to be an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetSubjectTypeIdNil() {
	o.SubjectTypeId.Set(nil)
}

// UnsetSubjectTypeId ensures that no value is present for SubjectTypeId, not even an explicit nil.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnsetSubjectTypeId() {
	o.SubjectTypeId.Unset()
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetSummary(v string) {
	o.Summary = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2ListDTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsExperimentV2ListDTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentV2ListDTODataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentV2ListDTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentsEndDate   datadog.NullableTime                                                  `json:"assignments_end_date,omitempty"`
		AssignmentsStartDate datadog.NullableTime                                                  `json:"assignments_start_date,omitempty"`
		ConcludedAt          datadog.NullableTime                                                  `json:"concluded_at,omitempty"`
		Conclusion           *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion         `json:"conclusion,omitempty"`
		CreatedAt            *time.Time                                                            `json:"created_at,omitempty"`
		EventsEndDate        datadog.NullableTime                                                  `json:"events_end_date,omitempty"`
		EventsStartDate      datadog.NullableTime                                                  `json:"events_start_date,omitempty"`
		ExperimentType       *string                                                               `json:"experiment_type,omitempty"`
		Hypothesis           *string                                                               `json:"hypothesis,omitempty"`
		MigrationMetadata    interface{}                                                           `json:"migration_metadata,omitempty"`
		Name                 *string                                                               `json:"name,omitempty"`
		PipelineTableSuffix  *string                                                               `json:"pipeline_table_suffix,omitempty"`
		ProtocolId           *string                                                               `json:"protocol_id,omitempty"`
		RelatedLinks         []ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems `json:"related_links,omitempty"`
		ResultsLastUpdated   datadog.NullableTime                                                  `json:"results_last_updated,omitempty"`
		Status               *ExperimentsExperimentV2DTODataAttributesStatus                       `json:"status,omitempty"`
		StructuredMetadata   []ExperimentsStructuredMetadataResponse                               `json:"structured_metadata,omitempty"`
		SubjectTypeId        datadog.NullableString                                                `json:"subject_type_id,omitempty"`
		Summary              *string                                                               `json:"summary,omitempty"`
		Tags                 []string                                                              `json:"tags,omitempty"`
		Teams                []string                                                              `json:"teams,omitempty"`
		UpdatedAt            *time.Time                                                            `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignments_end_date", "assignments_start_date", "concluded_at", "conclusion", "created_at", "events_end_date", "events_start_date", "experiment_type", "hypothesis", "migration_metadata", "name", "pipeline_table_suffix", "protocol_id", "related_links", "results_last_updated", "status", "structured_metadata", "subject_type_id", "summary", "tags", "teams", "updated_at"})
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
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
