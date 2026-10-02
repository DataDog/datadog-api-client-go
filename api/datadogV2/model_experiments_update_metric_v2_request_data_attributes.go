// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateMetricV2RequestDataAttributes Fields supplied to update the metric. Every attribute is optional; omit an attribute to leave it unchanged.
type ExperimentsUpdateMetricV2RequestDataAttributes struct {
	// Source of the data backing this metric.
	DataSourceType *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType `json:"data_source_type,omitempty"`
	// Measure and calculation settings for a numerator or denominator aggregation. Supply exactly one non-null measure.
	DenominatorAggregation *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"denominator_aggregation,omitempty"`
	// Send null to clear the description. Omit to leave it unchanged.
	Description datadog.NullableString `json:"description,omitempty"`
	// Direction of change that represents an improvement for this metric.
	DesiredChange *ExperimentsCreateMetricV2RequestDataAttributesDesiredChange `json:"desired_change,omitempty"`
	// Whether results render as a percentage. Omit to leave it unchanged.
	FormatAsPercent *bool `json:"format_as_percent,omitempty"`
	// Send null to clear a stored threshold. Omit to leave it unchanged.
	GuardrailCutoffThreshold datadog.NullableFloat64 `json:"guardrail_cutoff_threshold,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Name of the metric. Omit to leave it unchanged.
	Name *string `json:"name,omitempty"`
	// Measure and calculation settings for a numerator or denominator aggregation. Supply exactly one non-null measure.
	NumeratorAggregation *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation `json:"numerator_aggregation,omitempty"`
	// Measure and percentile to calculate for the metric. Supply exactly one non-null measure.
	PercentileAggregation *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation `json:"percentile_aggregation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateMetricV2RequestDataAttributes instantiates a new ExperimentsUpdateMetricV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateMetricV2RequestDataAttributes() *ExperimentsUpdateMetricV2RequestDataAttributes {
	this := ExperimentsUpdateMetricV2RequestDataAttributes{}
	return &this
}

// NewExperimentsUpdateMetricV2RequestDataAttributesWithDefaults instantiates a new ExperimentsUpdateMetricV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateMetricV2RequestDataAttributesWithDefaults() *ExperimentsUpdateMetricV2RequestDataAttributes {
	this := ExperimentsUpdateMetricV2RequestDataAttributes{}
	return &this
}

// GetDataSourceType returns the DataSourceType field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDataSourceType() ExperimentsCreateMetricV2RequestDataAttributesDataSourceType {
	if o == nil || o.DataSourceType == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDataSourceType
		return ret
	}
	return *o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDataSourceTypeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, bool) {
	if o == nil || o.DataSourceType == nil {
		return nil, false
	}
	return o.DataSourceType, true
}

// HasDataSourceType returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasDataSourceType() bool {
	return o != nil && o.DataSourceType != nil
}

// SetDataSourceType gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesDataSourceType and assigns it to the DataSourceType field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetDataSourceType(v ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) {
	o.DataSourceType = &v
}

// GetDenominatorAggregation returns the DenominatorAggregation field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDenominatorAggregation() ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	if o == nil || o.DenominatorAggregation == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation
		return ret
	}
	return *o.DenominatorAggregation
}

// GetDenominatorAggregationOk returns a tuple with the DenominatorAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDenominatorAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation, bool) {
	if o == nil || o.DenominatorAggregation == nil {
		return nil, false
	}
	return o.DenominatorAggregation, true
}

// HasDenominatorAggregation returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasDenominatorAggregation() bool {
	return o != nil && o.DenominatorAggregation != nil
}

// SetDenominatorAggregation gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation and assigns it to the DenominatorAggregation field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetDenominatorAggregation(v ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) {
	o.DenominatorAggregation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDesiredChange returns the DesiredChange field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDesiredChange() ExperimentsCreateMetricV2RequestDataAttributesDesiredChange {
	if o == nil || o.DesiredChange == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDesiredChange
		return ret
	}
	return *o.DesiredChange
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetDesiredChangeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDesiredChange, bool) {
	if o == nil || o.DesiredChange == nil {
		return nil, false
	}
	return o.DesiredChange, true
}

// HasDesiredChange returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasDesiredChange() bool {
	return o != nil && o.DesiredChange != nil
}

// SetDesiredChange gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesDesiredChange and assigns it to the DesiredChange field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetDesiredChange(v ExperimentsCreateMetricV2RequestDataAttributesDesiredChange) {
	o.DesiredChange = &v
}

// GetFormatAsPercent returns the FormatAsPercent field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetFormatAsPercent() bool {
	if o == nil || o.FormatAsPercent == nil {
		var ret bool
		return ret
	}
	return *o.FormatAsPercent
}

// GetFormatAsPercentOk returns a tuple with the FormatAsPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetFormatAsPercentOk() (*bool, bool) {
	if o == nil || o.FormatAsPercent == nil {
		return nil, false
	}
	return o.FormatAsPercent, true
}

// HasFormatAsPercent returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasFormatAsPercent() bool {
	return o != nil && o.FormatAsPercent != nil
}

// SetFormatAsPercent gets a reference to the given bool and assigns it to the FormatAsPercent field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetFormatAsPercent(v bool) {
	o.FormatAsPercent = &v
}

// GetGuardrailCutoffThreshold returns the GuardrailCutoffThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetGuardrailCutoffThreshold() float64 {
	if o == nil || o.GuardrailCutoffThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GuardrailCutoffThreshold.Get()
}

// GetGuardrailCutoffThresholdOk returns a tuple with the GuardrailCutoffThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetGuardrailCutoffThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailCutoffThreshold.Get(), o.GuardrailCutoffThreshold.IsSet()
}

// HasGuardrailCutoffThreshold returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasGuardrailCutoffThreshold() bool {
	return o != nil && o.GuardrailCutoffThreshold.IsSet()
}

// SetGuardrailCutoffThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the GuardrailCutoffThreshold field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetGuardrailCutoffThreshold(v float64) {
	o.GuardrailCutoffThreshold.Set(&v)
}

// SetGuardrailCutoffThresholdNil sets the value for GuardrailCutoffThreshold to be an explicit nil.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetGuardrailCutoffThresholdNil() {
	o.GuardrailCutoffThreshold.Set(nil)
}

// UnsetGuardrailCutoffThreshold ensures that no value is present for GuardrailCutoffThreshold, not even an explicit nil.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) UnsetGuardrailCutoffThreshold() {
	o.GuardrailCutoffThreshold.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetNumeratorAggregation returns the NumeratorAggregation field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetNumeratorAggregation() ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation {
	if o == nil || o.NumeratorAggregation == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation
		return ret
	}
	return *o.NumeratorAggregation
}

// GetNumeratorAggregationOk returns a tuple with the NumeratorAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetNumeratorAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation, bool) {
	if o == nil || o.NumeratorAggregation == nil {
		return nil, false
	}
	return o.NumeratorAggregation, true
}

// HasNumeratorAggregation returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasNumeratorAggregation() bool {
	return o != nil && o.NumeratorAggregation != nil
}

// SetNumeratorAggregation gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation and assigns it to the NumeratorAggregation field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetNumeratorAggregation(v ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation) {
	o.NumeratorAggregation = &v
}

// GetPercentileAggregation returns the PercentileAggregation field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetPercentileAggregation() ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation {
	if o == nil || o.PercentileAggregation == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation
		return ret
	}
	return *o.PercentileAggregation
}

// GetPercentileAggregationOk returns a tuple with the PercentileAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) GetPercentileAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation, bool) {
	if o == nil || o.PercentileAggregation == nil {
		return nil, false
	}
	return o.PercentileAggregation, true
}

// HasPercentileAggregation returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) HasPercentileAggregation() bool {
	return o != nil && o.PercentileAggregation != nil
}

// SetPercentileAggregation gets a reference to the given ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation and assigns it to the PercentileAggregation field.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) SetPercentileAggregation(v ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) {
	o.PercentileAggregation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateMetricV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataSourceType != nil {
		toSerialize["data_source_type"] = o.DataSourceType
	}
	if o.DenominatorAggregation != nil {
		toSerialize["denominator_aggregation"] = o.DenominatorAggregation
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DesiredChange != nil {
		toSerialize["desired_change"] = o.DesiredChange
	}
	if o.FormatAsPercent != nil {
		toSerialize["format_as_percent"] = o.FormatAsPercent
	}
	if o.GuardrailCutoffThreshold.IsSet() {
		toSerialize["guardrail_cutoff_threshold"] = o.GuardrailCutoffThreshold.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumeratorAggregation != nil {
		toSerialize["numerator_aggregation"] = o.NumeratorAggregation
	}
	if o.PercentileAggregation != nil {
		toSerialize["percentile_aggregation"] = o.PercentileAggregation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateMetricV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSourceType           *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType        `json:"data_source_type,omitempty"`
		DenominatorAggregation   *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation  `json:"denominator_aggregation,omitempty"`
		Description              datadog.NullableString                                               `json:"description,omitempty"`
		DesiredChange            *ExperimentsCreateMetricV2RequestDataAttributesDesiredChange         `json:"desired_change,omitempty"`
		FormatAsPercent          *bool                                                                `json:"format_as_percent,omitempty"`
		GuardrailCutoffThreshold datadog.NullableFloat64                                              `json:"guardrail_cutoff_threshold,omitempty"`
		MigrationMetadata        interface{}                                                          `json:"migration_metadata,omitempty"`
		Name                     *string                                                              `json:"name,omitempty"`
		NumeratorAggregation     *ExperimentsCreateMetricV2RequestDataAttributesNumeratorAggregation  `json:"numerator_aggregation,omitempty"`
		PercentileAggregation    *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation `json:"percentile_aggregation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source_type", "denominator_aggregation", "description", "desired_change", "format_as_percent", "guardrail_cutoff_threshold", "migration_metadata", "name", "numerator_aggregation", "percentile_aggregation"})
	} else {
		return err
	}

	hasInvalidField := false
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
	o.FormatAsPercent = all.FormatAsPercent
	o.GuardrailCutoffThreshold = all.GuardrailCutoffThreshold
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name
	o.NumeratorAggregation = all.NumeratorAggregation
	o.PercentileAggregation = all.PercentileAggregation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
