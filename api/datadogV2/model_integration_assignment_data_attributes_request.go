// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAssignmentDataAttributesRequest
type IntegrationAssignmentDataAttributesRequest struct {
	// Action to perform on the assignment. Can be "assign" or "un_assign".
	Action IntegrationAssignmentDataAttributesRequestAction `json:"action"`
	//
	Assignment IntegrationAssignmentDataAttributesRequestAssignment `json:"assignment"`
	// Type of the security issues. Can be "findings" or "vulnerabilities".
	Type IntegrationAssignmentDataAttributesRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAssignmentDataAttributesRequest instantiates a new IntegrationAssignmentDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAssignmentDataAttributesRequest(action IntegrationAssignmentDataAttributesRequestAction, assignment IntegrationAssignmentDataAttributesRequestAssignment, typeVar IntegrationAssignmentDataAttributesRequestType) *IntegrationAssignmentDataAttributesRequest {
	this := IntegrationAssignmentDataAttributesRequest{}
	this.Action = action
	this.Assignment = assignment
	this.Type = typeVar
	return &this
}

// NewIntegrationAssignmentDataAttributesRequestWithDefaults instantiates a new IntegrationAssignmentDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAssignmentDataAttributesRequestWithDefaults() *IntegrationAssignmentDataAttributesRequest {
	this := IntegrationAssignmentDataAttributesRequest{}
	return &this
}

// GetAction returns the Action field value.
func (o *IntegrationAssignmentDataAttributesRequest) GetAction() IntegrationAssignmentDataAttributesRequestAction {
	if o == nil {
		var ret IntegrationAssignmentDataAttributesRequestAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IntegrationAssignmentDataAttributesRequest) GetActionOk() (*IntegrationAssignmentDataAttributesRequestAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *IntegrationAssignmentDataAttributesRequest) SetAction(v IntegrationAssignmentDataAttributesRequestAction) {
	o.Action = v
}

// GetAssignment returns the Assignment field value.
func (o *IntegrationAssignmentDataAttributesRequest) GetAssignment() IntegrationAssignmentDataAttributesRequestAssignment {
	if o == nil {
		var ret IntegrationAssignmentDataAttributesRequestAssignment
		return ret
	}
	return o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value
// and a boolean to check if the value has been set.
func (o *IntegrationAssignmentDataAttributesRequest) GetAssignmentOk() (*IntegrationAssignmentDataAttributesRequestAssignment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignment, true
}

// SetAssignment sets field value.
func (o *IntegrationAssignmentDataAttributesRequest) SetAssignment(v IntegrationAssignmentDataAttributesRequestAssignment) {
	o.Assignment = v
}

// GetType returns the Type field value.
func (o *IntegrationAssignmentDataAttributesRequest) GetType() IntegrationAssignmentDataAttributesRequestType {
	if o == nil {
		var ret IntegrationAssignmentDataAttributesRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationAssignmentDataAttributesRequest) GetTypeOk() (*IntegrationAssignmentDataAttributesRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IntegrationAssignmentDataAttributesRequest) SetType(v IntegrationAssignmentDataAttributesRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAssignmentDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["assignment"] = o.Assignment
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAssignmentDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *IntegrationAssignmentDataAttributesRequestAction     `json:"action"`
		Assignment *IntegrationAssignmentDataAttributesRequestAssignment `json:"assignment"`
		Type       *IntegrationAssignmentDataAttributesRequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Assignment == nil {
		return fmt.Errorf("required field assignment missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "assignment", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	if all.Assignment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Assignment = *all.Assignment
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
