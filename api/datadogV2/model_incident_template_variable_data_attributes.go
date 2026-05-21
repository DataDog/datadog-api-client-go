// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTemplateVariableDataAttributes Attributes of a template variable.
type IncidentTemplateVariableDataAttributes struct {
	// A description of the template variable.
	Description string `json:"description"`
	// The display name of the template variable.
	DisplayName string `json:"display_name"`
	// The domain of the template variable.
	Domain string `json:"domain"`
	// The variable name used in templates.
	Variable string `json:"variable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTemplateVariableDataAttributes instantiates a new IncidentTemplateVariableDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTemplateVariableDataAttributes(description string, displayName string, domain string, variable string) *IncidentTemplateVariableDataAttributes {
	this := IncidentTemplateVariableDataAttributes{}
	this.Description = description
	this.DisplayName = displayName
	this.Domain = domain
	this.Variable = variable
	return &this
}

// NewIncidentTemplateVariableDataAttributesWithDefaults instantiates a new IncidentTemplateVariableDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTemplateVariableDataAttributesWithDefaults() *IncidentTemplateVariableDataAttributes {
	this := IncidentTemplateVariableDataAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *IncidentTemplateVariableDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplateVariableDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *IncidentTemplateVariableDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *IncidentTemplateVariableDataAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplateVariableDataAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *IncidentTemplateVariableDataAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDomain returns the Domain field value.
func (o *IncidentTemplateVariableDataAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplateVariableDataAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *IncidentTemplateVariableDataAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetVariable returns the Variable field value.
func (o *IncidentTemplateVariableDataAttributes) GetVariable() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Variable
}

// GetVariableOk returns a tuple with the Variable field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplateVariableDataAttributes) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variable, true
}

// SetVariable sets field value.
func (o *IncidentTemplateVariableDataAttributes) SetVariable(v string) {
	o.Variable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTemplateVariableDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	toSerialize["domain"] = o.Domain
	toSerialize["variable"] = o.Variable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTemplateVariableDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description"`
		DisplayName *string `json:"display_name"`
		Domain      *string `json:"domain"`
		Variable    *string `json:"variable"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	if all.Variable == nil {
		return fmt.Errorf("required field variable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "domain", "variable"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	o.Domain = *all.Domain
	o.Variable = *all.Variable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
