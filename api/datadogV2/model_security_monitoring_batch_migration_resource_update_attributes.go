// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceUpdateAttributes Attributes of the batch migration resource to update. At least one attribute is required.
type SecurityMonitoringBatchMigrationResourceUpdateAttributes struct {
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The source query or natural language description to translate. Must not be blank.
	Input *string `json:"input,omitempty"`
	// The type of the source rule: `spl` for a Splunk SPL search or saved search definition,
	// `natural_language` for a plain-English description of the behavior to detect, or `kql`
	// for a Microsoft KQL query. SPL inputs must be self-contained, with search macros expanded.
	InputType *SecurityMonitoringBatchMigrationResourceInputType `json:"inputType,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringBatchMigrationResourceUpdateAttributes instantiates a new SecurityMonitoringBatchMigrationResourceUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringBatchMigrationResourceUpdateAttributes() *SecurityMonitoringBatchMigrationResourceUpdateAttributes {
	this := SecurityMonitoringBatchMigrationResourceUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringBatchMigrationResourceUpdateAttributesWithDefaults instantiates a new SecurityMonitoringBatchMigrationResourceUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringBatchMigrationResourceUpdateAttributesWithDefaults() *SecurityMonitoringBatchMigrationResourceUpdateAttributes {
	this := SecurityMonitoringBatchMigrationResourceUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) HasInput() bool {
	return o != nil && o.Input != nil
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) SetInput(v string) {
	o.Input = &v
}

// GetInputType returns the InputType field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetInputType() SecurityMonitoringBatchMigrationResourceInputType {
	if o == nil || o.InputType == nil {
		var ret SecurityMonitoringBatchMigrationResourceInputType
		return ret
	}
	return *o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetInputTypeOk() (*SecurityMonitoringBatchMigrationResourceInputType, bool) {
	if o == nil || o.InputType == nil {
		return nil, false
	}
	return o.InputType, true
}

// HasInputType returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) HasInputType() bool {
	return o != nil && o.InputType != nil
}

// SetInputType gets a reference to the given SecurityMonitoringBatchMigrationResourceInputType and assigns it to the InputType field.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) SetInputType(v SecurityMonitoringBatchMigrationResourceInputType) {
	o.InputType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringBatchMigrationResourceUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.InputType != nil {
		toSerialize["inputType"] = o.InputType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringBatchMigrationResourceUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                            `json:"description,omitempty"`
		Input       *string                                            `json:"input,omitempty"`
		InputType   *SecurityMonitoringBatchMigrationResourceInputType `json:"inputType,omitempty"`
		Name        *string                                            `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Input = all.Input
	if all.InputType != nil && !all.InputType.IsValid() {
		hasInvalidField = true
	} else {
		o.InputType = all.InputType
	}
	o.Name = all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
