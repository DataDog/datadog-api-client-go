// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupTarget A target the scanning group applies to.
type SensitiveDataScannerGroupTarget struct {
	// Type of target whose scanning groups are served by the target-scoped group routes. `llm_observability` is Agent Observability.
	Type SensitiveDataScannerGroupTargetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerGroupTarget instantiates a new SensitiveDataScannerGroupTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerGroupTarget(typeVar SensitiveDataScannerGroupTargetType) *SensitiveDataScannerGroupTarget {
	this := SensitiveDataScannerGroupTarget{}
	this.Type = typeVar
	return &this
}

// NewSensitiveDataScannerGroupTargetWithDefaults instantiates a new SensitiveDataScannerGroupTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerGroupTargetWithDefaults() *SensitiveDataScannerGroupTarget {
	this := SensitiveDataScannerGroupTarget{}
	return &this
}

// GetType returns the Type field value.
func (o *SensitiveDataScannerGroupTarget) GetType() SensitiveDataScannerGroupTargetType {
	if o == nil {
		var ret SensitiveDataScannerGroupTargetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupTarget) GetTypeOk() (*SensitiveDataScannerGroupTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SensitiveDataScannerGroupTarget) SetType(v SensitiveDataScannerGroupTargetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerGroupTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerGroupTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *SensitiveDataScannerGroupTargetType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
