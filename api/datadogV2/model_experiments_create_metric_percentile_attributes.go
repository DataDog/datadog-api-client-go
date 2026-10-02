// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricPercentileAttributes Configuration for a percentile metric. Omit numerator_aggregation and denominator_aggregation.
type ExperimentsCreateMetricPercentileAttributes struct {
	// Source of the data backing this metric.
	DataSourceType ExperimentsCreateMetricV2RequestDataAttributesDataSourceType `json:"data_source_type"`
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
	// Measure and percentile to calculate for the metric. Supply exactly one non-null measure.
	PercentileAggregation ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation `json:"percentile_aggregation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewExperimentsCreateMetricPercentileAttributes instantiates a new ExperimentsCreateMetricPercentileAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateMetricPercentileAttributes(dataSourceType ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, desiredChange ExperimentsCreateMetricV2RequestDataAttributesDesiredChange, name string, percentileAggregation ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) *ExperimentsCreateMetricPercentileAttributes {
	this := ExperimentsCreateMetricPercentileAttributes{}
	this.DataSourceType = dataSourceType
	this.DesiredChange = desiredChange
	this.Name = name
	this.PercentileAggregation = percentileAggregation
	return &this
}

// NewExperimentsCreateMetricPercentileAttributesWithDefaults instantiates a new ExperimentsCreateMetricPercentileAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateMetricPercentileAttributesWithDefaults() *ExperimentsCreateMetricPercentileAttributes {
	this := ExperimentsCreateMetricPercentileAttributes{}
	return &this
}

// GetDataSourceType returns the DataSourceType field value.
func (o *ExperimentsCreateMetricPercentileAttributes) GetDataSourceType() ExperimentsCreateMetricV2RequestDataAttributesDataSourceType {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDataSourceType
		return ret
	}
	return o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetDataSourceTypeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDataSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceType, true
}

// SetDataSourceType sets field value.
func (o *ExperimentsCreateMetricPercentileAttributes) SetDataSourceType(v ExperimentsCreateMetricV2RequestDataAttributesDataSourceType) {
	o.DataSourceType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricPercentileAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricPercentileAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsCreateMetricPercentileAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsCreateMetricPercentileAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsCreateMetricPercentileAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDesiredChange returns the DesiredChange field value.
func (o *ExperimentsCreateMetricPercentileAttributes) GetDesiredChange() ExperimentsCreateMetricV2RequestDataAttributesDesiredChange {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesDesiredChange
		return ret
	}
	return o.DesiredChange
}

// GetDesiredChangeOk returns a tuple with the DesiredChange field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetDesiredChangeOk() (*ExperimentsCreateMetricV2RequestDataAttributesDesiredChange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredChange, true
}

// SetDesiredChange sets field value.
func (o *ExperimentsCreateMetricPercentileAttributes) SetDesiredChange(v ExperimentsCreateMetricV2RequestDataAttributesDesiredChange) {
	o.DesiredChange = v
}

// GetFormatAsPercent returns the FormatAsPercent field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricPercentileAttributes) GetFormatAsPercent() bool {
	if o == nil || o.FormatAsPercent == nil {
		var ret bool
		return ret
	}
	return *o.FormatAsPercent
}

// GetFormatAsPercentOk returns a tuple with the FormatAsPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetFormatAsPercentOk() (*bool, bool) {
	if o == nil || o.FormatAsPercent == nil {
		return nil, false
	}
	return o.FormatAsPercent, true
}

// HasFormatAsPercent returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) HasFormatAsPercent() bool {
	return o != nil && o.FormatAsPercent != nil
}

// SetFormatAsPercent gets a reference to the given bool and assigns it to the FormatAsPercent field.
func (o *ExperimentsCreateMetricPercentileAttributes) SetFormatAsPercent(v bool) {
	o.FormatAsPercent = &v
}

// GetGuardrailCutoffThreshold returns the GuardrailCutoffThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateMetricPercentileAttributes) GetGuardrailCutoffThreshold() float64 {
	if o == nil || o.GuardrailCutoffThreshold.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GuardrailCutoffThreshold.Get()
}

// GetGuardrailCutoffThresholdOk returns a tuple with the GuardrailCutoffThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateMetricPercentileAttributes) GetGuardrailCutoffThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailCutoffThreshold.Get(), o.GuardrailCutoffThreshold.IsSet()
}

// HasGuardrailCutoffThreshold returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) HasGuardrailCutoffThreshold() bool {
	return o != nil && o.GuardrailCutoffThreshold.IsSet()
}

// SetGuardrailCutoffThreshold gets a reference to the given datadog.NullableFloat64 and assigns it to the GuardrailCutoffThreshold field.
func (o *ExperimentsCreateMetricPercentileAttributes) SetGuardrailCutoffThreshold(v float64) {
	o.GuardrailCutoffThreshold.Set(&v)
}

// SetGuardrailCutoffThresholdNil sets the value for GuardrailCutoffThreshold to be an explicit nil.
func (o *ExperimentsCreateMetricPercentileAttributes) SetGuardrailCutoffThresholdNil() {
	o.GuardrailCutoffThreshold.Set(nil)
}

// UnsetGuardrailCutoffThreshold ensures that no value is present for GuardrailCutoffThreshold, not even an explicit nil.
func (o *ExperimentsCreateMetricPercentileAttributes) UnsetGuardrailCutoffThreshold() {
	o.GuardrailCutoffThreshold.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsCreateMetricPercentileAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateMetricPercentileAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateMetricPercentileAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateMetricPercentileAttributes) SetName(v string) {
	o.Name = v
}

// GetPercentileAggregation returns the PercentileAggregation field value.
func (o *ExperimentsCreateMetricPercentileAttributes) GetPercentileAggregation() ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation
		return ret
	}
	return o.PercentileAggregation
}

// GetPercentileAggregationOk returns a tuple with the PercentileAggregation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricPercentileAttributes) GetPercentileAggregationOk() (*ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentileAggregation, true
}

// SetPercentileAggregation sets field value.
func (o *ExperimentsCreateMetricPercentileAttributes) SetPercentileAggregation(v ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation) {
	o.PercentileAggregation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateMetricPercentileAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source_type"] = o.DataSourceType
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
	toSerialize["percentile_aggregation"] = o.PercentileAggregation
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateMetricPercentileAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSourceType           *ExperimentsCreateMetricV2RequestDataAttributesDataSourceType        `json:"data_source_type"`
		Description              datadog.NullableString                                               `json:"description,omitempty"`
		DesiredChange            *ExperimentsCreateMetricV2RequestDataAttributesDesiredChange         `json:"desired_change"`
		FormatAsPercent          *bool                                                                `json:"format_as_percent,omitempty"`
		GuardrailCutoffThreshold datadog.NullableFloat64                                              `json:"guardrail_cutoff_threshold,omitempty"`
		MigrationMetadata        interface{}                                                          `json:"migration_metadata,omitempty"`
		Name                     *string                                                              `json:"name"`
		PercentileAggregation    *ExperimentsCreateMetricV2RequestDataAttributesPercentileAggregation `json:"percentile_aggregation"`
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
	if all.PercentileAggregation == nil {
		return fmt.Errorf("required field percentile_aggregation missing")
	}

	hasInvalidField := false
	if !all.DataSourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSourceType = *all.DataSourceType
	}
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
	o.PercentileAggregation = *all.PercentileAggregation

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
