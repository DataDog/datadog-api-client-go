// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackStateAttributes Attributes of a content pack state.
type SecurityMonitoringContentPackStateAttributes struct {
	// Type-specific details for a content pack state. The set of fields present depends
	// on the content pack's `type`. When Cloud SIEM is inactive for the requesting organization, `onboarding` is returned instead of the content pack's usual type, such as `logs` or `vulnerability`.`
	Details SecurityMonitoringContentPackStateDetails `json:"details"`
	// The current operational status of a content pack.
	Status SecurityMonitoringContentPackStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackStateAttributes instantiates a new SecurityMonitoringContentPackStateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackStateAttributes(details SecurityMonitoringContentPackStateDetails, status SecurityMonitoringContentPackStatus) *SecurityMonitoringContentPackStateAttributes {
	this := SecurityMonitoringContentPackStateAttributes{}
	this.Details = details
	this.Status = status
	return &this
}

// NewSecurityMonitoringContentPackStateAttributesWithDefaults instantiates a new SecurityMonitoringContentPackStateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackStateAttributesWithDefaults() *SecurityMonitoringContentPackStateAttributes {
	this := SecurityMonitoringContentPackStateAttributes{}
	return &this
}

// GetDetails returns the Details field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetDetails() SecurityMonitoringContentPackStateDetails {
	if o == nil {
		var ret SecurityMonitoringContentPackStateDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetDetailsOk() (*SecurityMonitoringContentPackStateDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetDetails(v SecurityMonitoringContentPackStateDetails) {
	o.Details = v
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetStatus() SecurityMonitoringContentPackStatus {
	if o == nil {
		var ret SecurityMonitoringContentPackStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetStatusOk() (*SecurityMonitoringContentPackStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetStatus(v SecurityMonitoringContentPackStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackStateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["details"] = o.Details
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackStateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Details *SecurityMonitoringContentPackStateDetails `json:"details"`
		Status  *SecurityMonitoringContentPackStatus       `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Details == nil {
		return fmt.Errorf("required field details missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"details", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Details = *all.Details
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
