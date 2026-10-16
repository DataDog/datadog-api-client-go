// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceCreateItem A source rule to translate into a Cloud SIEM detection rule.
type SecurityMonitoringBatchMigrationResourceCreateItem struct {
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The source query or natural language description to translate. Must not be blank.
	Input string `json:"input"`
	// The type of the source rule: `spl` for a Splunk SPL search or saved search definition,
	// `natural_language` for a plain-English description of the behavior to detect, or `kql`
	// for a Microsoft KQL query. SPL inputs must be self-contained, with search macros expanded.
	InputType SecurityMonitoringBatchMigrationResourceInputType `json:"inputType"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourceCreateItem instantiates a new SecurityMonitoringBatchMigrationResourceCreateItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourceCreateItem(input string, inputType SecurityMonitoringBatchMigrationResourceInputType) *SecurityMonitoringBatchMigrationResourceCreateItem {
	this := SecurityMonitoringBatchMigrationResourceCreateItem{}
	this.Input = input
	this.InputType = inputType
	return &this
}

// NewSecurityMonitoringBatchMigrationResourceCreateItemWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourceCreateItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourceCreateItemWithDefaults() *SecurityMonitoringBatchMigrationResourceCreateItem {
	this := SecurityMonitoringBatchMigrationResourceCreateItem{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) SetDescription(v string) {
	o.Description = &v
}

// GetInput returns the Input field value.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) SetInput(v string) {
	o.Input = v
}

// GetInputType returns the InputType field value.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetInputType() SecurityMonitoringBatchMigrationResourceInputType {
	if o == nil {
		var ret SecurityMonitoringBatchMigrationResourceInputType
		return ret
	}
	return o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetInputTypeOk() (*SecurityMonitoringBatchMigrationResourceInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputType, true
}

// SetInputType sets field value.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) SetInputType(v SecurityMonitoringBatchMigrationResourceInputType) {
	o.InputType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourceCreateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["input"] = o.Input
	toSerialize["inputType"] = o.InputType
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourceCreateItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                            `json:"description,omitempty"`
		Input       *string                                            `json:"input"`
		InputType   *SecurityMonitoringBatchMigrationResourceInputType `json:"inputType"`
		Name        *string                                            `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Input == nil {
		return fmt.Errorf("required field input missing")
	}
	if all.InputType == nil {
		return fmt.Errorf("required field inputType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "input", "inputType", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Input = *all.Input
	if !all.InputType.IsValid() {
		hasInvalidField = true
	} else {
		o.InputType = *all.InputType
	}
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
