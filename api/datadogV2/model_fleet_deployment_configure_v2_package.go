// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentConfigureV2Package A package and its target version to additionally deploy alongside a configuration change.
type FleetDeploymentConfigureV2Package struct {
	// APM auto-instrumentation mode to enable for this package, if applicable.
	ApmInstrumentation *string `json:"apm_instrumentation,omitempty"`
	// The name of the package to deploy.
	Name string `json:"name"`
	// The target version of the package to deploy.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentConfigureV2Package instantiates a new FleetDeploymentConfigureV2Package object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentConfigureV2Package(name string, version string) *FleetDeploymentConfigureV2Package {
	this := FleetDeploymentConfigureV2Package{}
	this.Name = name
	this.Version = version
	return &this
}

// NewFleetDeploymentConfigureV2PackageWithDefaults instantiates a new FleetDeploymentConfigureV2Package object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentConfigureV2PackageWithDefaults() *FleetDeploymentConfigureV2Package {
	this := FleetDeploymentConfigureV2Package{}
	return &this
}

// GetApmInstrumentation returns the ApmInstrumentation field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2Package) GetApmInstrumentation() string {
	if o == nil || o.ApmInstrumentation == nil {
		var ret string
		return ret
	}
	return *o.ApmInstrumentation
}

// GetApmInstrumentationOk returns a tuple with the ApmInstrumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Package) GetApmInstrumentationOk() (*string, bool) {
	if o == nil || o.ApmInstrumentation == nil {
		return nil, false
	}
	return o.ApmInstrumentation, true
}

// HasApmInstrumentation returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2Package) HasApmInstrumentation() bool {
	return o != nil && o.ApmInstrumentation != nil
}

// SetApmInstrumentation gets a reference to the given string and assigns it to the ApmInstrumentation field.
func (o *FleetDeploymentConfigureV2Package) SetApmInstrumentation(v string) {
	o.ApmInstrumentation = &v
}

// GetName returns the Name field value.
func (o *FleetDeploymentConfigureV2Package) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Package) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FleetDeploymentConfigureV2Package) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value.
func (o *FleetDeploymentConfigureV2Package) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Package) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FleetDeploymentConfigureV2Package) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentConfigureV2Package) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApmInstrumentation != nil {
		toSerialize["apm_instrumentation"] = o.ApmInstrumentation
	}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentConfigureV2Package) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmInstrumentation *string `json:"apm_instrumentation,omitempty"`
		Name               *string `json:"name"`
		Version            *string `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_instrumentation", "name", "version"})
	} else {
		return err
	}
	o.ApmInstrumentation = all.ApmInstrumentation
	o.Name = *all.Name
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
