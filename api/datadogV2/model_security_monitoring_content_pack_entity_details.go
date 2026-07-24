// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackEntityDetails Details for an entity or identity content pack.
type SecurityMonitoringContentPackEntityDetails struct {
	// The activation status of a content pack.
	CpActivation SecurityMonitoringContentPackActivation `json:"cp_activation"`
	// Type for entity content pack details.
	Type SecurityMonitoringContentPackEntityDetailsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackEntityDetails instantiates a new SecurityMonitoringContentPackEntityDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackEntityDetails(cpActivation SecurityMonitoringContentPackActivation, typeVar SecurityMonitoringContentPackEntityDetailsType) *SecurityMonitoringContentPackEntityDetails {
	this := SecurityMonitoringContentPackEntityDetails{}
	this.CpActivation = cpActivation
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringContentPackEntityDetailsWithDefaults instantiates a new SecurityMonitoringContentPackEntityDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackEntityDetailsWithDefaults() *SecurityMonitoringContentPackEntityDetails {
	this := SecurityMonitoringContentPackEntityDetails{}
	return &this
}

// GetCpActivation returns the CpActivation field value.
func (o *SecurityMonitoringContentPackEntityDetails) GetCpActivation() SecurityMonitoringContentPackActivation {
	if o == nil {
		var ret SecurityMonitoringContentPackActivation
		return ret
	}
	return o.CpActivation
}

// GetCpActivationOk returns a tuple with the CpActivation field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackEntityDetails) GetCpActivationOk() (*SecurityMonitoringContentPackActivation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpActivation, true
}

// SetCpActivation sets field value.
func (o *SecurityMonitoringContentPackEntityDetails) SetCpActivation(v SecurityMonitoringContentPackActivation) {
	o.CpActivation = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringContentPackEntityDetails) GetType() SecurityMonitoringContentPackEntityDetailsType {
	if o == nil {
		var ret SecurityMonitoringContentPackEntityDetailsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackEntityDetails) GetTypeOk() (*SecurityMonitoringContentPackEntityDetailsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringContentPackEntityDetails) SetType(v SecurityMonitoringContentPackEntityDetailsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackEntityDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cp_activation"] = o.CpActivation
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackEntityDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpActivation *SecurityMonitoringContentPackActivation        `json:"cp_activation"`
		Type         *SecurityMonitoringContentPackEntityDetailsType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CpActivation == nil {
		return fmt.Errorf("required field cp_activation missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cp_activation", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CpActivation.IsValid() {
		hasInvalidField = true
	} else {
		o.CpActivation = *all.CpActivation
	}
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
