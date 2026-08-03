// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentV2DetailAgent Per-host status entry for a deployment.
type FleetDeploymentV2DetailAgent struct {
	// Error message if the deployment failed on this host.
	Error *string `json:"error,omitempty"`
	// Hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// Name of the step currently executing on this host.
	RunningStep *string `json:"running_step,omitempty"`
	// Deployment status for this host (for example, "pending", "running", "succeeded", "failed").
	Status *string `json:"status,omitempty"`
	// Additional details about the current deployment status on this host.
	StatusDetails *string `json:"status_details,omitempty"`
	// Package version details for this host.
	Versions []FleetDeploymentHostPackage `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentV2DetailAgent instantiates a new FleetDeploymentV2DetailAgent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentV2DetailAgent() *FleetDeploymentV2DetailAgent {
	this := FleetDeploymentV2DetailAgent{}
	return &this
}

// NewFleetDeploymentV2DetailAgentWithDefaults instantiates a new FleetDeploymentV2DetailAgent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentV2DetailAgentWithDefaults() *FleetDeploymentV2DetailAgent {
	this := FleetDeploymentV2DetailAgent{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FleetDeploymentV2DetailAgent) SetError(v string) {
	o.Error = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetDeploymentV2DetailAgent) SetHostname(v string) {
	o.Hostname = &v
}

// GetRunningStep returns the RunningStep field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetRunningStep() string {
	if o == nil || o.RunningStep == nil {
		var ret string
		return ret
	}
	return *o.RunningStep
}

// GetRunningStepOk returns a tuple with the RunningStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetRunningStepOk() (*string, bool) {
	if o == nil || o.RunningStep == nil {
		return nil, false
	}
	return o.RunningStep, true
}

// HasRunningStep returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasRunningStep() bool {
	return o != nil && o.RunningStep != nil
}

// SetRunningStep gets a reference to the given string and assigns it to the RunningStep field.
func (o *FleetDeploymentV2DetailAgent) SetRunningStep(v string) {
	o.RunningStep = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FleetDeploymentV2DetailAgent) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasStatusDetails() bool {
	return o != nil && o.StatusDetails != nil
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *FleetDeploymentV2DetailAgent) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *FleetDeploymentV2DetailAgent) GetVersions() []FleetDeploymentHostPackage {
	if o == nil || o.Versions == nil {
		var ret []FleetDeploymentHostPackage
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentV2DetailAgent) GetVersionsOk() (*[]FleetDeploymentHostPackage, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *FleetDeploymentV2DetailAgent) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []FleetDeploymentHostPackage and assigns it to the Versions field.
func (o *FleetDeploymentV2DetailAgent) SetVersions(v []FleetDeploymentHostPackage) {
	o.Versions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentV2DetailAgent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.RunningStep != nil {
		toSerialize["running_step"] = o.RunningStep
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["status_details"] = o.StatusDetails
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentV2DetailAgent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error         *string                      `json:"error,omitempty"`
		Hostname      *string                      `json:"hostname,omitempty"`
		RunningStep   *string                      `json:"running_step,omitempty"`
		Status        *string                      `json:"status,omitempty"`
		StatusDetails *string                      `json:"status_details,omitempty"`
		Versions      []FleetDeploymentHostPackage `json:"versions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "hostname", "running_step", "status", "status_details", "versions"})
	} else {
		return err
	}
	o.Error = all.Error
	o.Hostname = all.Hostname
	o.RunningStep = all.RunningStep
	o.Status = all.Status
	o.StatusDetails = all.StatusDetails
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
