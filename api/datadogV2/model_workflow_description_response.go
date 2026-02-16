// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowDescriptionResponse
type WorkflowDescriptionResponse struct {
	// The generated workflow description.
	Description string `json:"description"`
	// The generation ID.
	Id string `json:"id"`
	// A brief summary of the workflow.
	Summary string `json:"summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowDescriptionResponse instantiates a new WorkflowDescriptionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowDescriptionResponse(description string, id string, summary string) *WorkflowDescriptionResponse {
	this := WorkflowDescriptionResponse{}
	this.Description = description
	this.Id = id
	this.Summary = summary
	return &this
}

// NewWorkflowDescriptionResponseWithDefaults instantiates a new WorkflowDescriptionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowDescriptionResponseWithDefaults() *WorkflowDescriptionResponse {
	this := WorkflowDescriptionResponse{}
	return &this
}

// GetDescription returns the Description field value.
func (o *WorkflowDescriptionResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WorkflowDescriptionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *WorkflowDescriptionResponse) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value.
func (o *WorkflowDescriptionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowDescriptionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *WorkflowDescriptionResponse) SetId(v string) {
	o.Id = v
}

// GetSummary returns the Summary field value.
func (o *WorkflowDescriptionResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *WorkflowDescriptionResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *WorkflowDescriptionResponse) SetSummary(v string) {
	o.Summary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowDescriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["summary"] = o.Summary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowDescriptionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description"`
		Id          *string `json:"id"`
		Summary     *string `json:"summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "id", "summary"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Id = *all.Id
	o.Summary = *all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
