// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration Warehouse model and experiment key used to read assignment data.
type ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration struct {
	// Optional Warehouse measure that scopes analyzed subjects.
	EntryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint `json:"entry_point"`
	// Warehouse experiment key. Must not be blank.
	ExperimentKey string `json:"experiment_key"`
	// ID of the exposure SQL model that provides assignment data.
	ExposureSqlModelId string `json:"exposure_sql_model_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration(entryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint, experimentKey string, exposureSqlModelId string) *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration {
	this := ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration{}
	this.EntryPoint = entryPoint
	this.ExperimentKey = experimentKey
	this.ExposureSqlModelId = exposureSqlModelId
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfigurationWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfigurationWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration {
	this := ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration{}
	return &this
}

// GetEntryPoint returns the EntryPoint field value.
// If the value is explicit nil, the zero value for ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetEntryPoint() ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint {
	if o == nil || o.EntryPoint.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint
		return ret
	}
	return *o.EntryPoint.Get()
}

// GetEntryPointOk returns a tuple with the EntryPoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetEntryPointOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryPoint.Get(), o.EntryPoint.IsSet()
}

// SetEntryPoint sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) SetEntryPoint(v ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint) {
	o.EntryPoint.Set(&v)
}

// GetExperimentKey returns the ExperimentKey field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetExperimentKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExperimentKey
}

// GetExperimentKeyOk returns a tuple with the ExperimentKey field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetExperimentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentKey, true
}

// SetExperimentKey sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) SetExperimentKey(v string) {
	o.ExperimentKey = v
}

// GetExposureSqlModelId returns the ExposureSqlModelId field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetExposureSqlModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExposureSqlModelId
}

// GetExposureSqlModelIdOk returns a tuple with the ExposureSqlModelId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) GetExposureSqlModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExposureSqlModelId, true
}

// SetExposureSqlModelId sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) SetExposureSqlModelId(v string) {
	o.ExposureSqlModelId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entry_point"] = o.EntryPoint.Get()
	toSerialize["experiment_key"] = o.ExperimentKey
	toSerialize["exposure_sql_model_id"] = o.ExposureSqlModelId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntryPoint         NullableExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPoint `json:"entry_point"`
		ExperimentKey      *string                                                                                            `json:"experiment_key"`
		ExposureSqlModelId *string                                                                                            `json:"exposure_sql_model_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.EntryPoint.IsSet() {
		return fmt.Errorf("required field entry_point missing")
	}
	if all.ExperimentKey == nil {
		return fmt.Errorf("required field experiment_key missing")
	}
	if all.ExposureSqlModelId == nil {
		return fmt.Errorf("required field exposure_sql_model_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entry_point", "experiment_key", "exposure_sql_model_id"})
	} else {
		return err
	}
	o.EntryPoint = all.EntryPoint
	o.ExperimentKey = *all.ExperimentKey
	o.ExposureSqlModelId = *all.ExposureSqlModelId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration handles when a null is used for ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration.
type NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration struct {
	value *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) Get() *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) Set(val *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration initializes the struct as if Set has been called.
func NewNullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration(val *ExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) *NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration {
	return &NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesWarehouseExposureConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
