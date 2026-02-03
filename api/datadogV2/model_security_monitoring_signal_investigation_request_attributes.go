// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationRequestAttributes Attributes for creating a signal investigation.
type SecurityMonitoringSignalInvestigationRequestAttributes struct {
	// Optional deployment override for the investigation.
	Deployment *string `json:"deployment,omitempty"`
	// The unique ID of the security signal.
	SignalId string `json:"signal_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationRequestAttributes instantiates a new SecurityMonitoringSignalInvestigationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationRequestAttributes(signalId string) *SecurityMonitoringSignalInvestigationRequestAttributes {
	this := SecurityMonitoringSignalInvestigationRequestAttributes{}
	this.SignalId = signalId
	return &this
}

// NewSecurityMonitoringSignalInvestigationRequestAttributesWithDefaults instantiates a new SecurityMonitoringSignalInvestigationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationRequestAttributesWithDefaults() *SecurityMonitoringSignalInvestigationRequestAttributes {
	this := SecurityMonitoringSignalInvestigationRequestAttributes{}
	return &this
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) GetDeployment() string {
	if o == nil || o.Deployment == nil {
		var ret string
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) GetDeploymentOk() (*string, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) HasDeployment() bool {
	return o != nil && o.Deployment != nil
}

// SetDeployment gets a reference to the given string and assigns it to the Deployment field.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) SetDeployment(v string) {
	o.Deployment = &v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}
	toSerialize["signal_id"] = o.SignalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deployment *string `json:"deployment,omitempty"`
		SignalId   *string `json:"signal_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deployment", "signal_id"})
	} else {
		return err
	}
	o.Deployment = all.Deployment
	o.SignalId = *all.SignalId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
