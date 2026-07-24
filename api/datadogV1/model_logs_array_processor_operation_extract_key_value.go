// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationExtractKeyValue Operation that extracts key-value pairs from a `source` array and stores the result in the `target` attribute.
type LogsArrayProcessorOperationExtractKeyValue struct {
	// Key of the attribute in each array element that holds the name to use for the extracted attribute.
	KeyToExtract string `json:"key_to_extract"`
	// Whether to override the target element if it's already set.
	OverrideOnConflict *bool `json:"override_on_conflict,omitempty"`
	// Attribute path of the array to extract key-value pairs from.
	Source string `json:"source"`
	// Attribute that receives the extracted key-value pairs. If not specified, the extracted attributes are added at the root level of the log.
	Target *string `json:"target,omitempty"`
	// Operation type.
	Type LogsArrayProcessorOperationExtractKeyValueType `json:"type"`
	// Key of the attribute in each array element that holds the value to use for the extracted attribute.
	ValueToExtract string `json:"value_to_extract"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayProcessorOperationExtractKeyValue instantiates a new LogsArrayProcessorOperationExtractKeyValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayProcessorOperationExtractKeyValue(keyToExtract string, source string, typeVar LogsArrayProcessorOperationExtractKeyValueType, valueToExtract string) *LogsArrayProcessorOperationExtractKeyValue {
	this := LogsArrayProcessorOperationExtractKeyValue{}
	this.KeyToExtract = keyToExtract
	var overrideOnConflict bool = false
	this.OverrideOnConflict = &overrideOnConflict
	this.Source = source
	this.Type = typeVar
	this.ValueToExtract = valueToExtract
	return &this
}

// NewLogsArrayProcessorOperationExtractKeyValueWithDefaults instantiates a new LogsArrayProcessorOperationExtractKeyValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayProcessorOperationExtractKeyValueWithDefaults() *LogsArrayProcessorOperationExtractKeyValue {
	this := LogsArrayProcessorOperationExtractKeyValue{}
	var overrideOnConflict bool = false
	this.OverrideOnConflict = &overrideOnConflict
	return &this
}

// GetKeyToExtract returns the KeyToExtract field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetKeyToExtract() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.KeyToExtract
}

// GetKeyToExtractOk returns a tuple with the KeyToExtract field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetKeyToExtractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyToExtract, true
}

// SetKeyToExtract sets field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetKeyToExtract(v string) {
	o.KeyToExtract = v
}

// GetOverrideOnConflict returns the OverrideOnConflict field value if set, zero value otherwise.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetOverrideOnConflict() bool {
	if o == nil || o.OverrideOnConflict == nil {
		var ret bool
		return ret
	}
	return *o.OverrideOnConflict
}

// GetOverrideOnConflictOk returns a tuple with the OverrideOnConflict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetOverrideOnConflictOk() (*bool, bool) {
	if o == nil || o.OverrideOnConflict == nil {
		return nil, false
	}
	return o.OverrideOnConflict, true
}

// HasOverrideOnConflict returns a boolean if a field has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) HasOverrideOnConflict() bool {
	return o != nil && o.OverrideOnConflict != nil
}

// SetOverrideOnConflict gets a reference to the given bool and assigns it to the OverrideOnConflict field.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetOverrideOnConflict(v bool) {
	o.OverrideOnConflict = &v
}

// GetSource returns the Source field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetTarget(v string) {
	o.Target = &v
}

// GetType returns the Type field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetType() LogsArrayProcessorOperationExtractKeyValueType {
	if o == nil {
		var ret LogsArrayProcessorOperationExtractKeyValueType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetTypeOk() (*LogsArrayProcessorOperationExtractKeyValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetType(v LogsArrayProcessorOperationExtractKeyValueType) {
	o.Type = v
}

// GetValueToExtract returns the ValueToExtract field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetValueToExtract() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValueToExtract
}

// GetValueToExtractOk returns a tuple with the ValueToExtract field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationExtractKeyValue) GetValueToExtractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueToExtract, true
}

// SetValueToExtract sets field value.
func (o *LogsArrayProcessorOperationExtractKeyValue) SetValueToExtract(v string) {
	o.ValueToExtract = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayProcessorOperationExtractKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key_to_extract"] = o.KeyToExtract
	if o.OverrideOnConflict != nil {
		toSerialize["override_on_conflict"] = o.OverrideOnConflict
	}
	toSerialize["source"] = o.Source
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	toSerialize["type"] = o.Type
	toSerialize["value_to_extract"] = o.ValueToExtract

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayProcessorOperationExtractKeyValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KeyToExtract       *string                                         `json:"key_to_extract"`
		OverrideOnConflict *bool                                           `json:"override_on_conflict,omitempty"`
		Source             *string                                         `json:"source"`
		Target             *string                                         `json:"target,omitempty"`
		Type               *LogsArrayProcessorOperationExtractKeyValueType `json:"type"`
		ValueToExtract     *string                                         `json:"value_to_extract"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.KeyToExtract == nil {
		return fmt.Errorf("required field key_to_extract missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ValueToExtract == nil {
		return fmt.Errorf("required field value_to_extract missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key_to_extract", "override_on_conflict", "source", "target", "type", "value_to_extract"})
	} else {
		return err
	}

	hasInvalidField := false
	o.KeyToExtract = *all.KeyToExtract
	o.OverrideOnConflict = all.OverrideOnConflict
	o.Source = *all.Source
	o.Target = all.Target
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.ValueToExtract = *all.ValueToExtract

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
