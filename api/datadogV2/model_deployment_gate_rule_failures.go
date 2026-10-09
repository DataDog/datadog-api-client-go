// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateRuleFailures Rule failure details.
type DeploymentGateRuleFailures struct {
	// Names of faulty APM resources.
	FaultyApmResources []string `json:"faulty_apm_resources"`
	//
	Monitors []DeploymentGateRuleFailureMonitor `json:"monitors"`
	//
	Narratives []DeploymentGateRuleFailureNarrative `json:"narratives"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateRuleFailures instantiates a new DeploymentGateRuleFailures object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateRuleFailures(faultyApmResources []string, monitors []DeploymentGateRuleFailureMonitor, narratives []DeploymentGateRuleFailureNarrative) *DeploymentGateRuleFailures {
	this := DeploymentGateRuleFailures{}
	this.FaultyApmResources = faultyApmResources
	this.Monitors = monitors
	this.Narratives = narratives
	return &this
}

// NewDeploymentGateRuleFailuresWithDefaults instantiates a new DeploymentGateRuleFailures object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateRuleFailuresWithDefaults() *DeploymentGateRuleFailures {
	this := DeploymentGateRuleFailures{}
	return &this
}

// GetFaultyApmResources returns the FaultyApmResources field value.
func (o *DeploymentGateRuleFailures) GetFaultyApmResources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FaultyApmResources
}

// GetFaultyApmResourcesOk returns a tuple with the FaultyApmResources field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailures) GetFaultyApmResourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FaultyApmResources, true
}

// SetFaultyApmResources sets field value.
func (o *DeploymentGateRuleFailures) SetFaultyApmResources(v []string) {
	o.FaultyApmResources = v
}

// GetMonitors returns the Monitors field value.
func (o *DeploymentGateRuleFailures) GetMonitors() []DeploymentGateRuleFailureMonitor {
	if o == nil {
		var ret []DeploymentGateRuleFailureMonitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailures) GetMonitorsOk() (*[]DeploymentGateRuleFailureMonitor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Monitors, true
}

// SetMonitors sets field value.
func (o *DeploymentGateRuleFailures) SetMonitors(v []DeploymentGateRuleFailureMonitor) {
	o.Monitors = v
}

// GetNarratives returns the Narratives field value.
func (o *DeploymentGateRuleFailures) GetNarratives() []DeploymentGateRuleFailureNarrative {
	if o == nil {
		var ret []DeploymentGateRuleFailureNarrative
		return ret
	}
	return o.Narratives
}

// GetNarrativesOk returns a tuple with the Narratives field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleFailures) GetNarrativesOk() (*[]DeploymentGateRuleFailureNarrative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Narratives, true
}

// SetNarratives sets field value.
func (o *DeploymentGateRuleFailures) SetNarratives(v []DeploymentGateRuleFailureNarrative) {
	o.Narratives = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateRuleFailures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["faulty_apm_resources"] = o.FaultyApmResources
	toSerialize["monitors"] = o.Monitors
	toSerialize["narratives"] = o.Narratives

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateRuleFailures) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FaultyApmResources *[]string                             `json:"faulty_apm_resources"`
		Monitors           *[]DeploymentGateRuleFailureMonitor   `json:"monitors"`
		Narratives         *[]DeploymentGateRuleFailureNarrative `json:"narratives"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FaultyApmResources == nil {
		return fmt.Errorf("required field faulty_apm_resources missing")
	}
	if all.Monitors == nil {
		return fmt.Errorf("required field monitors missing")
	}
	if all.Narratives == nil {
		return fmt.Errorf("required field narratives missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"faulty_apm_resources", "monitors", "narratives"})
	} else {
		return err
	}
	o.FaultyApmResources = *all.FaultyApmResources
	o.Monitors = *all.Monitors
	o.Narratives = *all.Narratives

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
