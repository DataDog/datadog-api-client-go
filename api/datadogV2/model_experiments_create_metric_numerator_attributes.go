// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricNumeratorAttributes Configuration for a metric calculated from a numerator and an optional denominator. Omit percentile_aggregation. Omit denominator_aggregation when unused.
type ExperimentsCreateMetricNumeratorAttributes struct {
	// Source of the data backing this metric.
	DataSourceType ExperimentsCreateMetricV2RequestDataAttributesDataSourceType `json:"data_source_type"`
	// Measure and calculation settings for a numerator or denominator aggregation. Supply exactly one non-null measure.
	DenominatorAggregation *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"denominator_aggregation,omitempty"`
	// Description of the metric. Send null to leave it unset.
	Description datadog.NullableString `json:"description,omitempty"`
	// Direction of change that represents an improvement for this metric.
	DesiredChange ExperimentsCreateMetricV2RequestDataAttributesDesiredChange `json:"desired_change"`
	// Whether results render as a percentage. Defaults to false when omitted.
	FormatAsPercent *bool `json:"format_as_percent,omitempty"`
	// Guardrail cutoff threshold. Send null to leave it unset.
	GuardrailCutoffThreshold datadog.NullableFloat64 `json:"guardrail_cutoff_threshold,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Name of the metric.
	Name string `json:"name"`
	// Measure and calculation settings for a numerator or denominator aggregation. Supply exactly one non-null measure.
	NumeratorAggregation ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"numerator_aggregation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewExperimentsCreateMetricNumeratorAttributes instantiates a new ExperimentsCreateMetricNumeratorAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateMetricNumeratorAttributes(dataSourceType ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, desiredChange ExperimentsCreateMetricV2RequestDataAttributesDesiredChange, name string, numeratorAggregation ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) *ExperimentsCreateMetricNumeratorAttributes {
	this := ExperimentsCreateMetricNumeratorAttributes{}
	this.DataSourceType = dataSourceType
	this.DesiredChange = desiredChange
	this.Name = name
	this.NumeratorAggregation = numeratorAggregation
	return &this
}

// NewExperimentsCreateMetricNumeratorAttributesWithDefaults instantiates a new ExperimentsCreateMetricNumeratorAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateMetricNumeratorAttributesWithDefaults() *ExperimentsCreateMetricNumeratorAttributes {
	this := ExperimentsCreateMetricNumeratorAttributes{}
	return &this
}

// GetDataSourceType returns the DataSourceType field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDataSourceType() ExperimentsCreateMetricV2RequestDataAttributesDataSourceType {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDataSourceType
		return ret
	}
	return o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDataSourceTypeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceType, true
}

// SetDataSourceType sets field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetDataSourceType(v ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) {
	o.DataSourceType = v
}

// GetDenominatorAggregation returns the DenominatorAggregation field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDenominatorAggregation() ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	if o == nil || o.DenominatorAggregation == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation
		return ret
	}
	return *o.DenominatorAggregation
}

// GetDenominatorAggregationOk returns a tuple with the DenominatorAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDenominatorAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation, bool) {
	if o == nil || o.DenominatorAggregation == nil {
		return nil, false
	}
	return o.DenominatorAggregation, true
}

// HasDenominatorAggregation returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) HasDenominatorAggregation() bool {
	return o != nil && o.DenominatorAggregation != nil
}

// SetDenominatorAggregation gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation and assigns it to the DenominatorAggregation field.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetDenominatorAggregation(v ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) {
	o.DenominatorAggregation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsCreateMetricNumeratorAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDesiredChange returns the DesiredChange field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDesiredChange() ExperimentsCreateMetricV2RequestDataAttributesDesiredChange {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDesiredChange
		return ret
	}
	return o.DesiredChange
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetDesiredChangeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDesiredChange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredChange, true
}

// SetDesiredChange sets field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetDesiredChange(v ExperimentsCreateMetricV2RequestDataAttributesDesiredChange) {
	o.DesiredChange = v
}

// GetFormatAsPercent returns the FormatAsPercent field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetFormatAsPercent() bool {
	if o == nil || o.FormatAsPercent == nil {
		var ret bool
		return ret
	}
	return *o.FormatAsPercent
}

// GetFormatAsPercentOk returns a tuple with the FormatAsPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetFormatAsPercentOk() (*bool, bool) {
	if o == nil || o.FormatAsPercent == nil {
		return nil, false
	}
	return o.FormatAsPercent, true
}

// HasFormatAsPercent returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) HasFormatAsPercent() bool {
	return o != nil && o.FormatAsPercent != nil
}

// SetFormatAsPercent gets a reference to the given bool and assigns it to the FormatAsPercent field.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetFormatAsPercent(v bool) {
	o.FormatAsPercent = &v
}

// GetGuardrailCutoffThreshold returns the GuardrailCutoffThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricNumeratorAttributes) GetGuardrailCutoffThreshold() float64 {
	if o == nil || o.GuardrailCutoffThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GuardrailCutoffThreshold.Get()
}

// GetGuardrailCutoffThresholdOk returns a tuple with the GuardrailCutoffThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetGuardrailCutoffThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailCutoffThreshold.Get(), o.GuardrailCutoffThreshold.IsSet()
}

// HasGuardrailCutoffThreshold returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) HasGuardrailCutoffThreshold() bool {
	return o != nil && o.GuardrailCutoffThreshold.IsSet()
}

// SetGuardrailCutoffThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the GuardrailCutoffThreshold field.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetGuardrailCutoffThreshold(v float64) {
	o.GuardrailCutoffThreshold.Set(&v)
}

// SetGuardrailCutoffThresholdNil sets the value for GuardrailCutoffThreshold to be an explicit nil.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetGuardrailCutoffThresholdNil() {
	o.GuardrailCutoffThreshold.Set(nil)
}

// UnsetGuardrailCutoffThreshold ensures that no value is present for GuardrailCutoffThreshold, not even an explicit nil.
func (o *ExperimentsCreateMetricNumeratorAttributes) UnsetGuardrailCutoffThreshold() {
	o.GuardrailCutoffThreshold.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetName(v string) {
	o.Name = v
}

// GetNumeratorAggregation returns the NumeratorAggregation field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetNumeratorAggregation() ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation
		return ret
	}
	return o.NumeratorAggregation
}

// GetNumeratorAggregationOk returns a tuple with the NumeratorAggregation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricNumeratorAttributes) GetNumeratorAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeratorAggregation, true
}

// SetNumeratorAggregation sets field value.
func (o *ExperimentsCreateMetricNumeratorAttributes) SetNumeratorAggregation(v ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) {
	o.NumeratorAggregation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateMetricNumeratorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source_type"] = o.DataSourceType
	if o.DenominatorAggregation != nil {
		toSerialize["denominator_aggregation"] = o.DenominatorAggregation
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["desired_change"] = o.DesiredChange
	if o.FormatAsPercent != nil {
		toSerialize["format_as_percent"] = o.FormatAsPercent
	}
	if o.GuardrailCutoffThreshold.IsSet() {
		toSerialize["guardrail_cutoff_threshold"] = o.GuardrailCutoffThreshold.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	toSerialize["name"] = o.Name
	toSerialize["numerator_aggregation"] = o.NumeratorAggregation
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateMetricNumeratorAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSourceType           *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType       `json:"data_source_type"`
		DenominatorAggregation   *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"denominator_aggregation,omitempty"`
		Description              datadog.NullableString                                              `json:"description,omitempty"`
		DesiredChange            *ExperimentsCreateMetricV2RequestDataAttributesDesiredChange        `json:"desired_change"`
		FormatAsPercent          *bool                                                               `json:"format_as_percent,omitempty"`
		GuardrailCutoffThreshold datadog.NullableFloat64                                             `json:"guardrail_cutoff_threshold,omitempty"`
		MigrationMetadata        interface{}                                                         `json:"migration_metadata,omitempty"`
		Name                     *string                                                             `json:"name"`
		NumeratorAggregation     *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"numerator_aggregation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSourceType == nil {
		return fmt.Errorf("required field data_source_type missing")
	}
	if all.DesiredChange == nil {
		return fmt.Errorf("required field desired_change missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NumeratorAggregation == nil {
		return fmt.Errorf("required field numerator_aggregation missing")
	}

	hasInvalidField := false
	if !all.DataSourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSourceType = *all.DataSourceType
	}
	o.DenominatorAggregation = all.DenominatorAggregation
	o.Description = all.Description
	if !all.DesiredChange.IsValid() {
		hasInvalidField = true
	} else {
		o.DesiredChange = *all.DesiredChange
	}
	o.FormatAsPercent = all.FormatAsPercent
	o.GuardrailCutoffThreshold = all.GuardrailCutoffThreshold
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name
	o.NumeratorAggregation = *all.NumeratorAggregation

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
