// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchByVersionRequestAttributes Attributes for patching a DORA deployment event identified by service, environment, and version.
type DORADeploymentPatchByVersionRequestAttributes struct {
	// Indicates whether the deployment resulted in a change failure.
	ChangeFailure bool `json:"change_failure"`
	// The environment the deployment was performed in.
	Env string `json:"env"`
	// Remediation details for the deployment. Optional, but required to calculate failed deployment recovery time. Specify either `id` or `version` to identify the remediation deployment, but not both.
	Remediation *DORADeploymentPatchByVersionRemediation `json:"remediation,omitempty"`
	// The name of the service that was deployed.
	Service string `json:"service"`
	// The version deployed. This can be seen in the Service Catalog or in the APM Deployment Tracking.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentPatchByVersionRequestAttributes instantiates a new DORADeploymentPatchByVersionRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentPatchByVersionRequestAttributes(changeFailure bool, env string, service string, version string) *DORADeploymentPatchByVersionRequestAttributes {
	this := DORADeploymentPatchByVersionRequestAttributes{}
	this.ChangeFailure = changeFailure
	this.Env = env
	this.Service = service
	this.Version = version
	return &this
}

// NewDORADeploymentPatchByVersionRequestAttributesWithDefaults instantiates a new DORADeploymentPatchByVersionRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentPatchByVersionRequestAttributesWithDefaults() *DORADeploymentPatchByVersionRequestAttributes {
	this := DORADeploymentPatchByVersionRequestAttributes{}
	return &this
}

// GetChangeFailure returns the ChangeFailure field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetChangeFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ChangeFailure
}

// GetChangeFailureOk returns a tuple with the ChangeFailure field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetChangeFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeFailure, true
}

// SetChangeFailure sets field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) SetChangeFailure(v bool) {
	o.ChangeFailure = v
}

// GetEnv returns the Env field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) SetEnv(v string) {
	o.Env = v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetRemediation() DORADeploymentPatchByVersionRemediation {
	if o == nil || o.Remediation == nil {
		var ret DORADeploymentPatchByVersionRemediation
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetRemediationOk() (*DORADeploymentPatchByVersionRemediation, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given DORADeploymentPatchByVersionRemediation and assigns it to the Remediation field.
func (o *DORADeploymentPatchByVersionRequestAttributes) SetRemediation(v DORADeploymentPatchByVersionRemediation) {
	o.Remediation = &v
}

// GetService returns the Service field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) SetService(v string) {
	o.Service = v
}

// GetVersion returns the Version field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRequestAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *DORADeploymentPatchByVersionRequestAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentPatchByVersionRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["change_failure"] = o.ChangeFailure
	toSerialize["env"] = o.Env
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}
	toSerialize["service"] = o.Service
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentPatchByVersionRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeFailure *bool                                    `json:"change_failure"`
		Env           *string                                  `json:"env"`
		Remediation   *DORADeploymentPatchByVersionRemediation `json:"remediation,omitempty"`
		Service       *string                                  `json:"service"`
		Version       *string                                  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChangeFailure == nil {
		return fmt.Errorf("required field change_failure missing")
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_failure", "env", "remediation", "service", "version"})
	} else {
		return err
	}
	o.ChangeFailure = *all.ChangeFailure
	o.Env = *all.Env
	o.Remediation = all.Remediation
	o.Service = *all.Service
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
