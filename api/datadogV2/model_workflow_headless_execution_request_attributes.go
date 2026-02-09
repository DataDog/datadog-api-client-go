// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowHeadlessExecutionRequestAttributes
type WorkflowHeadlessExecutionRequestAttributes struct {
	//
	Config WorkflowHeadlessExecutionConfig `json:"config"`
	// The ID of the workflow template to execute
	TemplateId string `json:"template_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowHeadlessExecutionRequestAttributes instantiates a new WorkflowHeadlessExecutionRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowHeadlessExecutionRequestAttributes(config WorkflowHeadlessExecutionConfig, templateId string) *WorkflowHeadlessExecutionRequestAttributes {
	this := WorkflowHeadlessExecutionRequestAttributes{}
	this.Config = config
	this.TemplateId = templateId
	return &this
}

// NewWorkflowHeadlessExecutionRequestAttributesWithDefaults instantiates a new WorkflowHeadlessExecutionRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowHeadlessExecutionRequestAttributesWithDefaults() *WorkflowHeadlessExecutionRequestAttributes {
	this := WorkflowHeadlessExecutionRequestAttributes{}
	return &this
}

// GetConfig returns the Config field value.
func (o *WorkflowHeadlessExecutionRequestAttributes) GetConfig() WorkflowHeadlessExecutionConfig {
	if o == nil {
		var ret WorkflowHeadlessExecutionConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionRequestAttributes) GetConfigOk() (*WorkflowHeadlessExecutionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *WorkflowHeadlessExecutionRequestAttributes) SetConfig(v WorkflowHeadlessExecutionConfig) {
	o.Config = v
}

// GetTemplateId returns the TemplateId field value.
func (o *WorkflowHeadlessExecutionRequestAttributes) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionRequestAttributes) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value.
func (o *WorkflowHeadlessExecutionRequestAttributes) SetTemplateId(v string) {
	o.TemplateId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowHeadlessExecutionRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	toSerialize["template_id"] = o.TemplateId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowHeadlessExecutionRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config     *WorkflowHeadlessExecutionConfig `json:"config"`
		TemplateId *string                          `json:"template_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.TemplateId == nil {
		return fmt.Errorf("required field template_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "template_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config
	o.TemplateId = *all.TemplateId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
