// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideRequestDataAttributes Attributes of the severity override request.
type SeverityOverrideRequestDataAttributes struct {
	// Severity override to apply to the findings.
	// Set `action` to `set` to apply a manual severity override with the given `value`.
	// Set `action` to `clear` to remove a manual severity override.
	Severity SeverityOverrideAttributes `json:"severity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideRequestDataAttributes instantiates a new SeverityOverrideRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideRequestDataAttributes(severity SeverityOverrideAttributes) *SeverityOverrideRequestDataAttributes {
	this := SeverityOverrideRequestDataAttributes{}
	this.Severity = severity
	return &this
}

// NewSeverityOverrideRequestDataAttributesWithDefaults instantiates a new SeverityOverrideRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideRequestDataAttributesWithDefaults() *SeverityOverrideRequestDataAttributes {
	this := SeverityOverrideRequestDataAttributes{}
	return &this
}

// GetSeverity returns the Severity field value.
func (o *SeverityOverrideRequestDataAttributes) GetSeverity() SeverityOverrideAttributes {
	if o == nil {
		var ret SeverityOverrideAttributes
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideRequestDataAttributes) GetSeverityOk() (*SeverityOverrideAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *SeverityOverrideRequestDataAttributes) SetSeverity(v SeverityOverrideAttributes) {
	o.Severity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["severity"] = o.Severity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Severity *SeverityOverrideAttributes `json:"severity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"severity"})
	} else {
		return err
	}
	o.Severity = *all.Severity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
