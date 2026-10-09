// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideResponseMeta Security findings skipped while processing the severity override request.
type SeverityOverrideResponseMeta struct {
	// Findings skipped because an automation rule set their severity.
	Warnings []SeverityOverrideResult `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideResponseMeta instantiates a new SeverityOverrideResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideResponseMeta() *SeverityOverrideResponseMeta {
	this := SeverityOverrideResponseMeta{}
	return &this
}

// NewSeverityOverrideResponseMetaWithDefaults instantiates a new SeverityOverrideResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideResponseMetaWithDefaults() *SeverityOverrideResponseMeta {
	this := SeverityOverrideResponseMeta{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SeverityOverrideResponseMeta) GetWarnings() []SeverityOverrideResult {
	if o == nil || o.Warnings == nil {
		var ret []SeverityOverrideResult
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResponseMeta) GetWarningsOk() (*[]SeverityOverrideResult, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SeverityOverrideResponseMeta) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []SeverityOverrideResult and assigns it to the Warnings field.
func (o *SeverityOverrideResponseMeta) SetWarnings(v []SeverityOverrideResult) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Warnings []SeverityOverrideResult `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"warnings"})
	} else {
		return err
	}
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
