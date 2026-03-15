// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceEnrollmentAttributes Attributes for creating an enrollment.
type OnPremManagementServiceEnrollmentAttributes struct {
	// List of allowed actions for the runner.
	ActionsAllowlist []string `json:"actions_allowlist"`
	// The hostname for the runner. Required when push mode is enabled.
	RunnerHost *string `json:"runner_host,omitempty"`
	// The modes the runner should operate in.
	RunnerModes []OnPremManagementServiceEnrollmentAttributesRunnerModesItems `json:"runner_modes"`
	// The name of the on-prem runner.
	RunnerName string `json:"runner_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnPremManagementServiceEnrollmentAttributes instantiates a new OnPremManagementServiceEnrollmentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnPremManagementServiceEnrollmentAttributes(actionsAllowlist []string, runnerModes []OnPremManagementServiceEnrollmentAttributesRunnerModesItems, runnerName string) *OnPremManagementServiceEnrollmentAttributes {
	this := OnPremManagementServiceEnrollmentAttributes{}
	this.ActionsAllowlist = actionsAllowlist
	this.RunnerModes = runnerModes
	this.RunnerName = runnerName
	return &this
}

// NewOnPremManagementServiceEnrollmentAttributesWithDefaults instantiates a new OnPremManagementServiceEnrollmentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnPremManagementServiceEnrollmentAttributesWithDefaults() *OnPremManagementServiceEnrollmentAttributes {
	this := OnPremManagementServiceEnrollmentAttributes{}
	return &this
}

// GetActionsAllowlist returns the ActionsAllowlist field value.
func (o *OnPremManagementServiceEnrollmentAttributes) GetActionsAllowlist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActionsAllowlist
}

// GetActionsAllowlistOk returns a tuple with the ActionsAllowlist field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceEnrollmentAttributes) GetActionsAllowlistOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionsAllowlist, true
}

// SetActionsAllowlist sets field value.
func (o *OnPremManagementServiceEnrollmentAttributes) SetActionsAllowlist(v []string) {
	o.ActionsAllowlist = v
}

// GetRunnerHost returns the RunnerHost field value if set, zero value otherwise.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerHost() string {
	if o == nil || o.RunnerHost == nil {
		var ret string
		return ret
	}
	return *o.RunnerHost
}

// GetRunnerHostOk returns a tuple with the RunnerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerHostOk() (*string, bool) {
	if o == nil || o.RunnerHost == nil {
		return nil, false
	}
	return o.RunnerHost, true
}

// HasRunnerHost returns a boolean if a field has been set.
func (o *OnPremManagementServiceEnrollmentAttributes) HasRunnerHost() bool {
	return o != nil && o.RunnerHost != nil
}

// SetRunnerHost gets a reference to the given string and assigns it to the RunnerHost field.
func (o *OnPremManagementServiceEnrollmentAttributes) SetRunnerHost(v string) {
	o.RunnerHost = &v
}

// GetRunnerModes returns the RunnerModes field value.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerModes() []OnPremManagementServiceEnrollmentAttributesRunnerModesItems {
	if o == nil {
		var ret []OnPremManagementServiceEnrollmentAttributesRunnerModesItems
		return ret
	}
	return o.RunnerModes
}

// GetRunnerModesOk returns a tuple with the RunnerModes field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerModesOk() (*[]OnPremManagementServiceEnrollmentAttributesRunnerModesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerModes, true
}

// SetRunnerModes sets field value.
func (o *OnPremManagementServiceEnrollmentAttributes) SetRunnerModes(v []OnPremManagementServiceEnrollmentAttributesRunnerModesItems) {
	o.RunnerModes = v
}

// GetRunnerName returns the RunnerName field value.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunnerName
}

// GetRunnerNameOk returns a tuple with the RunnerName field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceEnrollmentAttributes) GetRunnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerName, true
}

// SetRunnerName sets field value.
func (o *OnPremManagementServiceEnrollmentAttributes) SetRunnerName(v string) {
	o.RunnerName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnPremManagementServiceEnrollmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actions_allowlist"] = o.ActionsAllowlist
	if o.RunnerHost != nil {
		toSerialize["runner_host"] = o.RunnerHost
	}
	toSerialize["runner_modes"] = o.RunnerModes
	toSerialize["runner_name"] = o.RunnerName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnPremManagementServiceEnrollmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionsAllowlist *[]string                                                      `json:"actions_allowlist"`
		RunnerHost       *string                                                        `json:"runner_host,omitempty"`
		RunnerModes      *[]OnPremManagementServiceEnrollmentAttributesRunnerModesItems `json:"runner_modes"`
		RunnerName       *string                                                        `json:"runner_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionsAllowlist == nil {
		return fmt.Errorf("required field actions_allowlist missing")
	}
	if all.RunnerModes == nil {
		return fmt.Errorf("required field runner_modes missing")
	}
	if all.RunnerName == nil {
		return fmt.Errorf("required field runner_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions_allowlist", "runner_host", "runner_modes", "runner_name"})
	} else {
		return err
	}
	o.ActionsAllowlist = *all.ActionsAllowlist
	o.RunnerHost = all.RunnerHost
	o.RunnerModes = *all.RunnerModes
	o.RunnerName = *all.RunnerName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
