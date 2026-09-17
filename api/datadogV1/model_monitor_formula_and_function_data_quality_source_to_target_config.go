// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualitySourceToTargetConfig Configuration for a source to target monitor, which compares the same measure
// across two data entities and alerts on the difference between them.
type MonitorFormulaAndFunctionDataQualitySourceToTargetConfig struct {
	// How the difference between the source and target measures is computed.
	// `absolute` subtracts the two values, `diff_percent` expresses the difference
	// as a percentage of the source value.
	DiffType MonitorFormulaAndFunctionDataQualityDiffType `json:"diff_type"`
	// Type of the data entities being compared.
	EntityType string `json:"entity_type"`
	// Measure configuration for one side of a source to target comparison.
	Source MonitorFormulaAndFunctionDataQualityEntityMetricConfig `json:"source"`
	// Measure configuration for one side of a source to target comparison.
	Target MonitorFormulaAndFunctionDataQualityEntityMetricConfig `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualitySourceToTargetConfig instantiates a new MonitorFormulaAndFunctionDataQualitySourceToTargetConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualitySourceToTargetConfig(diffType MonitorFormulaAndFunctionDataQualityDiffType, entityType string, source MonitorFormulaAndFunctionDataQualityEntityMetricConfig, target MonitorFormulaAndFunctionDataQualityEntityMetricConfig) *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig {
	this := MonitorFormulaAndFunctionDataQualitySourceToTargetConfig{}
	this.DiffType = diffType
	this.EntityType = entityType
	this.Source = source
	this.Target = target
	return &this
}

// NewMonitorFormulaAndFunctionDataQualitySourceToTargetConfigWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualitySourceToTargetConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualitySourceToTargetConfigWithDefaults() *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig {
	this := MonitorFormulaAndFunctionDataQualitySourceToTargetConfig{}
	return &this
}

// GetDiffType returns the DiffType field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetDiffType() MonitorFormulaAndFunctionDataQualityDiffType {
	if o == nil {
		var ret MonitorFormulaAndFunctionDataQualityDiffType
		return ret
	}
	return o.DiffType
}

// GetDiffTypeOk returns a tuple with the DiffType field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetDiffTypeOk() (*MonitorFormulaAndFunctionDataQualityDiffType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiffType, true
}

// SetDiffType sets field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) SetDiffType(v MonitorFormulaAndFunctionDataQualityDiffType) {
	o.DiffType = v
}

// GetEntityType returns the EntityType field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) SetEntityType(v string) {
	o.EntityType = v
}

// GetSource returns the Source field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetSource() MonitorFormulaAndFunctionDataQualityEntityMetricConfig {
	if o == nil {
		var ret MonitorFormulaAndFunctionDataQualityEntityMetricConfig
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetSourceOk() (*MonitorFormulaAndFunctionDataQualityEntityMetricConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) SetSource(v MonitorFormulaAndFunctionDataQualityEntityMetricConfig) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetTarget() MonitorFormulaAndFunctionDataQualityEntityMetricConfig {
	if o == nil {
		var ret MonitorFormulaAndFunctionDataQualityEntityMetricConfig
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) GetTargetOk() (*MonitorFormulaAndFunctionDataQualityEntityMetricConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) SetTarget(v MonitorFormulaAndFunctionDataQualityEntityMetricConfig) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["diff_type"] = o.DiffType
	toSerialize["entity_type"] = o.EntityType
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualitySourceToTargetConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DiffType   *MonitorFormulaAndFunctionDataQualityDiffType           `json:"diff_type"`
		EntityType *string                                                 `json:"entity_type"`
		Source     *MonitorFormulaAndFunctionDataQualityEntityMetricConfig `json:"source"`
		Target     *MonitorFormulaAndFunctionDataQualityEntityMetricConfig `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DiffType == nil {
		return fmt.Errorf("required field diff_type missing")
	}
	if all.EntityType == nil {
		return fmt.Errorf("required field entity_type missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"diff_type", "entity_type", "source", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DiffType.IsValid() {
		hasInvalidField = true
	} else {
		o.DiffType = *all.DiffType
	}
	o.EntityType = *all.EntityType
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
