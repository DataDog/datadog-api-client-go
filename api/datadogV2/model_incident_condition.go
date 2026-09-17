// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCondition Conditions that determine which incidents trigger the workflow.
type IncidentCondition struct {
	// Incident tags and values used to filter matching incidents.
	TagValues []IncidentTagValue `json:"tagValues,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCondition instantiates a new IncidentCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCondition() *IncidentCondition {
	this := IncidentCondition{}
	return &this
}

// NewIncidentConditionWithDefaults instantiates a new IncidentCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentConditionWithDefaults() *IncidentCondition {
	this := IncidentCondition{}
	return &this
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *IncidentCondition) GetTagValues() []IncidentTagValue {
	if o == nil || o.TagValues == nil {
		var ret []IncidentTagValue
		return ret
	}
	return o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCondition) GetTagValuesOk() (*[]IncidentTagValue, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return &o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *IncidentCondition) HasTagValues() bool {
	return o != nil && o.TagValues != nil
}

// SetTagValues gets a reference to the given []IncidentTagValue and assigns it to the TagValues field.
func (o *IncidentCondition) SetTagValues(v []IncidentTagValue) {
	o.TagValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TagValues != nil {
		toSerialize["tagValues"] = o.TagValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagValues []IncidentTagValue `json:"tagValues,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tagValues"})
	} else {
		return err
	}
	o.TagValues = all.TagValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
