// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTODataAttributes Details of the metric.
type ExperimentsMetricV2DTODataAttributes struct {
	// Time when this resource was certified.
	CertifiedAt datadog.NullableTime `json:"certified_at,omitempty"`
	// Time when this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Source of the data used to calculate the metric.
	DataSourceType *ExperimentsMetricV2DTODataAttributesDataSourceType `json:"data_source_type,omitempty"`
	// Source measure and aggregation settings for a metric value.
	DenominatorAggregation NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation `json:"denominator_aggregation,omitempty"`
	// Text that explains the metric.
	Description *string `json:"description,omitempty"`
	// Direction of metric change considered desirable.
	DesiredChange *ExperimentsMetricV2DTODataAttributesDesiredChange `json:"desired_change,omitempty"`
	// Number of experiments that reference this resource.
	ExperimentCount datadog.NullableInt64 `json:"experiment_count,omitempty"`
	// Whether to display the metric value as a percentage.
	FormatAsPercent *bool `json:"format_as_percent,omitempty"`
	// Threshold used when evaluating this metric as a guardrail.
	GuardrailCutoffThreshold datadog.NullableFloat64 `json:"guardrail_cutoff_threshold,omitempty"`
	// Type of metric calculation.
	MetricType *ExperimentsMetricV2DTODataAttributesMetricType `json:"metric_type,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the metric.
	Name *string `json:"name,omitempty"`
	// Source measure and aggregation settings for a metric value.
	NumeratorAggregation NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation `json:"numerator_aggregation,omitempty"`
	// Source measure and settings for a percentile metric.
	PercentileAggregation NullableExperimentsMetricV2DTODataAttributesPercentileAggregation `json:"percentile_aggregation,omitempty"`
	// URL with supporting information about the metric.
	ReferenceUrl *string `json:"reference_url,omitempty"`
	// Time when this resource was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTODataAttributes instantiates a new ExperimentsMetricV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTODataAttributes() *ExperimentsMetricV2DTODataAttributes {
	this := ExperimentsMetricV2DTODataAttributes{}
	return &this
}

// NewExperimentsMetricV2DTODataAttributesWithDefaults instantiates a new ExperimentsMetricV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTODataAttributesWithDefaults() *ExperimentsMetricV2DTODataAttributes {
	this := ExperimentsMetricV2DTODataAttributes{}
	return &this
}

// GetCertifiedAt returns the CertifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetCertifiedAt() time.Time {
	if o == nil || o.CertifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CertifiedAt.Get()
}

// GetCertifiedAtOk returns a tuple with the CertifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetCertifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertifiedAt.Get(), o.CertifiedAt.IsSet()
}

// HasCertifiedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasCertifiedAt() bool {
	return o != nil && o.CertifiedAt.IsSet()
}

// SetCertifiedAt gets a reference to the given datadog.NullableTime and assigns it to the CertifiedAt field.
func (o *ExperimentsMetricV2DTODataAttributes) SetCertifiedAt(v time.Time) {
	o.CertifiedAt.Set(&v)
}

// SetCertifiedAtNil sets the value for CertifiedAt to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetCertifiedAtNil() {
	o.CertifiedAt.Set(nil)
}

// UnsetCertifiedAt ensures that no value is present for CertifiedAt, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetCertifiedAt() {
	o.CertifiedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsMetricV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDataSourceType returns the DataSourceType field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetDataSourceType() ExperimentsMetricV2DTODataAttributesDataSourceType {
	if o == nil || o.DataSourceType == nil {
		var ret ExperimentsMetricV2DTODataAttributesDataSourceType
		return ret
	}
	return *o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetDataSourceTypeOk() (*ExperimentsMetricV2DTODataAttributesDataSourceType, bool) {
	if o == nil || o.DataSourceType == nil {
		return nil, false
	}
	return o.DataSourceType, true
}

// HasDataSourceType returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasDataSourceType() bool {
	return o != nil && o.DataSourceType != nil
}

// SetDataSourceType gets a reference to the given ExperimentsMetricV2DTODataAttributesDataSourceType and assigns it to the DataSourceType field.
func (o *ExperimentsMetricV2DTODataAttributes) SetDataSourceType(v ExperimentsMetricV2DTODataAttributesDataSourceType) {
	o.DataSourceType = &v
}

// GetDenominatorAggregation returns the DenominatorAggregation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetDenominatorAggregation() ExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	if o == nil || o.DenominatorAggregation.Get() == nil {
		var ret ExperimentsMetricV2DTODataAttributesNumeratorAggregation
		return ret
	}
	return *o.DenominatorAggregation.Get()
}

// GetDenominatorAggregationOk returns a tuple with the DenominatorAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetDenominatorAggregationOk() (*ExperimentsMetricV2DTODataAttributesNumeratorAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenominatorAggregation.Get(), o.DenominatorAggregation.IsSet()
}

// HasDenominatorAggregation returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasDenominatorAggregation() bool {
	return o != nil && o.DenominatorAggregation.IsSet()
}

// SetDenominatorAggregation gets a reference to the given NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation and assigns it to the DenominatorAggregation field.
func (o *ExperimentsMetricV2DTODataAttributes) SetDenominatorAggregation(v ExperimentsMetricV2DTODataAttributesNumeratorAggregation) {
	o.DenominatorAggregation.Set(&v)
}

// SetDenominatorAggregationNil sets the value for DenominatorAggregation to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetDenominatorAggregationNil() {
	o.DenominatorAggregation.Set(nil)
}

// UnsetDenominatorAggregation ensures that no value is present for DenominatorAggregation, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetDenominatorAggregation() {
	o.DenominatorAggregation.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsMetricV2DTODataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredChange returns the DesiredChange field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetDesiredChange() ExperimentsMetricV2DTODataAttributesDesiredChange {
	if o == nil || o.DesiredChange == nil {
		var ret ExperimentsMetricV2DTODataAttributesDesiredChange
		return ret
	}
	return *o.DesiredChange
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetDesiredChangeOk() (*ExperimentsMetricV2DTODataAttributesDesiredChange, bool) {
	if o == nil || o.DesiredChange == nil {
		return nil, false
	}
	return o.DesiredChange, true
}

// HasDesiredChange returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasDesiredChange() bool {
	return o != nil && o.DesiredChange != nil
}

// SetDesiredChange gets a reference to the given ExperimentsMetricV2DTODataAttributesDesiredChange and assigns it to the DesiredChange field.
func (o *ExperimentsMetricV2DTODataAttributes) SetDesiredChange(v ExperimentsMetricV2DTODataAttributesDesiredChange) {
	o.DesiredChange = &v
}

// GetExperimentCount returns the ExperimentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetExperimentCount() int64 {
	if o == nil || o.ExperimentCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentCount.Get()
}

// GetExperimentCountOk returns a tuple with the ExperimentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetExperimentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentCount.Get(), o.ExperimentCount.IsSet()
}

// HasExperimentCount returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasExperimentCount() bool {
	return o != nil && o.ExperimentCount.IsSet()
}

// SetExperimentCount gets a reference to the given datadog.NullableInt64 and assigns it to the ExperimentCount field.
func (o *ExperimentsMetricV2DTODataAttributes) SetExperimentCount(v int64) {
	o.ExperimentCount.Set(&v)
}

// SetExperimentCountNil sets the value for ExperimentCount to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetExperimentCountNil() {
	o.ExperimentCount.Set(nil)
}

// UnsetExperimentCount ensures that no value is present for ExperimentCount, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetExperimentCount() {
	o.ExperimentCount.Unset()
}

// GetFormatAsPercent returns the FormatAsPercent field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetFormatAsPercent() bool {
	if o == nil || o.FormatAsPercent == nil {
		var ret bool
		return ret
	}
	return *o.FormatAsPercent
}

// GetFormatAsPercentOk returns a tuple with the FormatAsPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetFormatAsPercentOk() (*bool, bool) {
	if o == nil || o.FormatAsPercent == nil {
		return nil, false
	}
	return o.FormatAsPercent, true
}

// HasFormatAsPercent returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasFormatAsPercent() bool {
	return o != nil && o.FormatAsPercent != nil
}

// SetFormatAsPercent gets a reference to the given bool and assigns it to the FormatAsPercent field.
func (o *ExperimentsMetricV2DTODataAttributes) SetFormatAsPercent(v bool) {
	o.FormatAsPercent = &v
}

// GetGuardrailCutoffThreshold returns the GuardrailCutoffThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetGuardrailCutoffThreshold() float64 {
	if o == nil || o.GuardrailCutoffThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GuardrailCutoffThreshold.Get()
}

// GetGuardrailCutoffThresholdOk returns a tuple with the GuardrailCutoffThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetGuardrailCutoffThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailCutoffThreshold.Get(), o.GuardrailCutoffThreshold.IsSet()
}

// HasGuardrailCutoffThreshold returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasGuardrailCutoffThreshold() bool {
	return o != nil && o.GuardrailCutoffThreshold.IsSet()
}

// SetGuardrailCutoffThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the GuardrailCutoffThreshold field.
func (o *ExperimentsMetricV2DTODataAttributes) SetGuardrailCutoffThreshold(v float64) {
	o.GuardrailCutoffThreshold.Set(&v)
}

// SetGuardrailCutoffThresholdNil sets the value for GuardrailCutoffThreshold to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetGuardrailCutoffThresholdNil() {
	o.GuardrailCutoffThreshold.Set(nil)
}

// UnsetGuardrailCutoffThreshold ensures that no value is present for GuardrailCutoffThreshold, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetGuardrailCutoffThreshold() {
	o.GuardrailCutoffThreshold.Unset()
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetMetricType() ExperimentsMetricV2DTODataAttributesMetricType {
	if o == nil || o.MetricType == nil {
		var ret ExperimentsMetricV2DTODataAttributesMetricType
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetMetricTypeOk() (*ExperimentsMetricV2DTODataAttributesMetricType, bool) {
	if o == nil || o.MetricType == nil {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasMetricType() bool {
	return o != nil && o.MetricType != nil
}

// SetMetricType gets a reference to the given ExperimentsMetricV2DTODataAttributesMetricType and assigns it to the MetricType field.
func (o *ExperimentsMetricV2DTODataAttributes) SetMetricType(v ExperimentsMetricV2DTODataAttributesMetricType) {
	o.MetricType = &v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsMetricV2DTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsMetricV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetNumeratorAggregation returns the NumeratorAggregation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetNumeratorAggregation() ExperimentsMetricV2DTODataAttributesNumeratorAggregation {
	if o == nil || o.NumeratorAggregation.Get() == nil {
		var ret ExperimentsMetricV2DTODataAttributesNumeratorAggregation
		return ret
	}
	return *o.NumeratorAggregation.Get()
}

// GetNumeratorAggregationOk returns a tuple with the NumeratorAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetNumeratorAggregationOk() (*ExperimentsMetricV2DTODataAttributesNumeratorAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeratorAggregation.Get(), o.NumeratorAggregation.IsSet()
}

// HasNumeratorAggregation returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasNumeratorAggregation() bool {
	return o != nil && o.NumeratorAggregation.IsSet()
}

// SetNumeratorAggregation gets a reference to the given NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation and assigns it to the NumeratorAggregation field.
func (o *ExperimentsMetricV2DTODataAttributes) SetNumeratorAggregation(v ExperimentsMetricV2DTODataAttributesNumeratorAggregation) {
	o.NumeratorAggregation.Set(&v)
}

// SetNumeratorAggregationNil sets the value for NumeratorAggregation to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetNumeratorAggregationNil() {
	o.NumeratorAggregation.Set(nil)
}

// UnsetNumeratorAggregation ensures that no value is present for NumeratorAggregation, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetNumeratorAggregation() {
	o.NumeratorAggregation.Unset()
}

// GetPercentileAggregation returns the PercentileAggregation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricV2DTODataAttributes) GetPercentileAggregation() ExperimentsMetricV2DTODataAttributesPercentileAggregation {
	if o == nil || o.PercentileAggregation.Get() == nil {
		var ret ExperimentsMetricV2DTODataAttributesPercentileAggregation
		return ret
	}
	return *o.PercentileAggregation.Get()
}

// GetPercentileAggregationOk returns a tuple with the PercentileAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricV2DTODataAttributes) GetPercentileAggregationOk() (*ExperimentsMetricV2DTODataAttributesPercentileAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentileAggregation.Get(), o.PercentileAggregation.IsSet()
}

// HasPercentileAggregation returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasPercentileAggregation() bool {
	return o != nil && o.PercentileAggregation.IsSet()
}

// SetPercentileAggregation gets a reference to the given NullableExperimentsMetricV2DTODataAttributesPercentileAggregation and assigns it to the PercentileAggregation field.
func (o *ExperimentsMetricV2DTODataAttributes) SetPercentileAggregation(v ExperimentsMetricV2DTODataAttributesPercentileAggregation) {
	o.PercentileAggregation.Set(&v)
}

// SetPercentileAggregationNil sets the value for PercentileAggregation to be an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) SetPercentileAggregationNil() {
	o.PercentileAggregation.Set(nil)
}

// UnsetPercentileAggregation ensures that no value is present for PercentileAggregation, not even an explicit nil.
func (o *ExperimentsMetricV2DTODataAttributes) UnsetPercentileAggregation() {
	o.PercentileAggregation.Unset()
}

// GetReferenceUrl returns the ReferenceUrl field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetReferenceUrl() string {
	if o == nil || o.ReferenceUrl == nil {
		var ret string
		return ret
	}
	return *o.ReferenceUrl
}

// GetReferenceUrlOk returns a tuple with the ReferenceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetReferenceUrlOk() (*string, bool) {
	if o == nil || o.ReferenceUrl == nil {
		return nil, false
	}
	return o.ReferenceUrl, true
}

// HasReferenceUrl returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasReferenceUrl() bool {
	return o != nil && o.ReferenceUrl != nil
}

// SetReferenceUrl gets a reference to the given string and assigns it to the ReferenceUrl field.
func (o *ExperimentsMetricV2DTODataAttributes) SetReferenceUrl(v string) {
	o.ReferenceUrl = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsMetricV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CertifiedAt.IsSet() {
		toSerialize["certified_at"] = o.CertifiedAt.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DataSourceType != nil {
		toSerialize["data_source_type"] = o.DataSourceType
	}
	if o.DenominatorAggregation.IsSet() {
		toSerialize["denominator_aggregation"] = o.DenominatorAggregation.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DesiredChange != nil {
		toSerialize["desired_change"] = o.DesiredChange
	}
	if o.ExperimentCount.IsSet() {
		toSerialize["experiment_count"] = o.ExperimentCount.Get()
	}
	if o.FormatAsPercent != nil {
		toSerialize["format_as_percent"] = o.FormatAsPercent
	}
	if o.GuardrailCutoffThreshold.IsSet() {
		toSerialize["guardrail_cutoff_threshold"] = o.GuardrailCutoffThreshold.Get()
	}
	if o.MetricType != nil {
		toSerialize["metric_type"] = o.MetricType
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumeratorAggregation.IsSet() {
		toSerialize["numerator_aggregation"] = o.NumeratorAggregation.Get()
	}
	if o.PercentileAggregation.IsSet() {
		toSerialize["percentile_aggregation"] = o.PercentileAggregation.Get()
	}
	if o.ReferenceUrl != nil {
		toSerialize["reference_url"] = o.ReferenceUrl
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
func (o *ExperimentsMetricV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CertifiedAt              datadog.NullableTime                                              `json:"certified_at,omitempty"`
		CreatedAt                *time.Time                                                        `json:"created_at,omitempty"`
		DataSourceType           *ExperimentsMetricV2DTODataAttributesDataSourceType               `json:"data_source_type,omitempty"`
		DenominatorAggregation   NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation  `json:"denominator_aggregation,omitempty"`
		Description              *string                                                           `json:"description,omitempty"`
		DesiredChange            *ExperimentsMetricV2DTODataAttributesDesiredChange                `json:"desired_change,omitempty"`
		ExperimentCount          datadog.NullableInt64                                             `json:"experiment_count,omitempty"`
		FormatAsPercent          *bool                                                             `json:"format_as_percent,omitempty"`
		GuardrailCutoffThreshold datadog.NullableFloat64                                           `json:"guardrail_cutoff_threshold,omitempty"`
		MetricType               *ExperimentsMetricV2DTODataAttributesMetricType                   `json:"metric_type,omitempty"`
		MigrationMetadata        interface{}                                                       `json:"migration_metadata,omitempty"`
		Name                     *string                                                           `json:"name,omitempty"`
		NumeratorAggregation     NullableExperimentsMetricV2DTODataAttributesNumeratorAggregation  `json:"numerator_aggregation,omitempty"`
		PercentileAggregation    NullableExperimentsMetricV2DTODataAttributesPercentileAggregation `json:"percentile_aggregation,omitempty"`
		ReferenceUrl             *string                                                           `json:"reference_url,omitempty"`
		UpdatedAt                *time.Time                                                        `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"certified_at", "created_at", "data_source_type", "denominator_aggregation", "description", "desired_change", "experiment_count", "format_as_percent", "guardrail_cutoff_threshold", "metric_type", "migration_metadata", "name", "numerator_aggregation", "percentile_aggregation", "reference_url", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CertifiedAt = all.CertifiedAt
	o.CreatedAt = all.CreatedAt
	if all.DataSourceType != nil && !all.DataSourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSourceType = all.DataSourceType
	}
	o.DenominatorAggregation = all.DenominatorAggregation
	o.Description = all.Description
	if all.DesiredChange != nil && !all.DesiredChange.IsValid() {
		hasInvalidField = true
	} else {
		o.DesiredChange = all.DesiredChange
	}
	o.ExperimentCount = all.ExperimentCount
	o.FormatAsPercent = all.FormatAsPercent
	o.GuardrailCutoffThreshold = all.GuardrailCutoffThreshold
	if all.MetricType != nil && !all.MetricType.IsValid() {
		hasInvalidField = true
	} else {
		o.MetricType = all.MetricType
	}
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.NumeratorAggregation = all.NumeratorAggregation
	o.PercentileAggregation = all.PercentileAggregation
	o.ReferenceUrl = all.ReferenceUrl
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
