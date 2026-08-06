// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlMitigationDefinition The definition of a mitigation available for a control.
type GovernanceControlMitigationDefinition struct {
	// A human-readable description of the mitigation.
	Description string `json:"description"`
	// The execution modes the mitigation supports, such as `manual` or `automatic`.
	ExecutionModes []string `json:"execution_modes"`
	// The unique identifier of the mitigation.
	Id string `json:"id"`
	// The permissions required to apply the mitigation.
	Permissions []string `json:"permissions"`
	// An array of parameter definitions.
	SupportedParameters []GovernanceControlParameterDefinition `json:"supported_parameters"`
	// A short, human-readable name for the mitigation.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlMitigationDefinition instantiates a new GovernanceControlMitigationDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlMitigationDefinition(description string, executionModes []string, id string, permissions []string, supportedParameters []GovernanceControlParameterDefinition, title string) *GovernanceControlMitigationDefinition {
	this := GovernanceControlMitigationDefinition{}
	this.Description = description
	this.ExecutionModes = executionModes
	this.Id = id
	this.Permissions = permissions
	this.SupportedParameters = supportedParameters
	this.Title = title
	return &this
}

// NewGovernanceControlMitigationDefinitionWithDefaults instantiates a new GovernanceControlMitigationDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlMitigationDefinitionWithDefaults() *GovernanceControlMitigationDefinition {
	this := GovernanceControlMitigationDefinition{}
	return &this
}

// GetDescription returns the Description field value.
func (o *GovernanceControlMitigationDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceControlMitigationDefinition) SetDescription(v string) {
	o.Description = v
}

// GetExecutionModes returns the ExecutionModes field value.
func (o *GovernanceControlMitigationDefinition) GetExecutionModes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExecutionModes
}

// GetExecutionModesOk returns a tuple with the ExecutionModes field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetExecutionModesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionModes, true
}

// SetExecutionModes sets field value.
func (o *GovernanceControlMitigationDefinition) SetExecutionModes(v []string) {
	o.ExecutionModes = v
}

// GetId returns the Id field value.
func (o *GovernanceControlMitigationDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GovernanceControlMitigationDefinition) SetId(v string) {
	o.Id = v
}

// GetPermissions returns the Permissions field value.
func (o *GovernanceControlMitigationDefinition) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetPermissionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *GovernanceControlMitigationDefinition) SetPermissions(v []string) {
	o.Permissions = v
}

// GetSupportedParameters returns the SupportedParameters field value.
func (o *GovernanceControlMitigationDefinition) GetSupportedParameters() []GovernanceControlParameterDefinition {
	if o == nil {
		var ret []GovernanceControlParameterDefinition
		return ret
	}
	return o.SupportedParameters
}

// GetSupportedParametersOk returns a tuple with the SupportedParameters field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetSupportedParametersOk() (*[]GovernanceControlParameterDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedParameters, true
}

// SetSupportedParameters sets field value.
func (o *GovernanceControlMitigationDefinition) SetSupportedParameters(v []GovernanceControlParameterDefinition) {
	o.SupportedParameters = v
}

// GetTitle returns the Title field value.
func (o *GovernanceControlMitigationDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *GovernanceControlMitigationDefinition) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlMitigationDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["execution_modes"] = o.ExecutionModes
	toSerialize["id"] = o.Id
	toSerialize["permissions"] = o.Permissions
	toSerialize["supported_parameters"] = o.SupportedParameters
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlMitigationDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description         *string                                 `json:"description"`
		ExecutionModes      *[]string                               `json:"execution_modes"`
		Id                  *string                                 `json:"id"`
		Permissions         *[]string                               `json:"permissions"`
		SupportedParameters *[]GovernanceControlParameterDefinition `json:"supported_parameters"`
		Title               *string                                 `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ExecutionModes == nil {
		return fmt.Errorf("required field execution_modes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	if all.SupportedParameters == nil {
		return fmt.Errorf("required field supported_parameters missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "execution_modes", "id", "permissions", "supported_parameters", "title"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.ExecutionModes = *all.ExecutionModes
	o.Id = *all.Id
	o.Permissions = *all.Permissions
	o.SupportedParameters = *all.SupportedParameters
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
