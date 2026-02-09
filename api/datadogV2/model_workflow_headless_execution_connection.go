// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowHeadlessExecutionConnection
type WorkflowHeadlessExecutionConnection struct {
	// The ID of the connection
	ConnectionId uuid.UUID `json:"connection_id"`
	// The label for the connection
	Label string `json:"label"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowHeadlessExecutionConnection instantiates a new WorkflowHeadlessExecutionConnection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowHeadlessExecutionConnection(connectionId uuid.UUID, label string) *WorkflowHeadlessExecutionConnection {
	this := WorkflowHeadlessExecutionConnection{}
	this.ConnectionId = connectionId
	this.Label = label
	return &this
}

// NewWorkflowHeadlessExecutionConnectionWithDefaults instantiates a new WorkflowHeadlessExecutionConnection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowHeadlessExecutionConnectionWithDefaults() *WorkflowHeadlessExecutionConnection {
	this := WorkflowHeadlessExecutionConnection{}
	return &this
}

// GetConnectionId returns the ConnectionId field value.
func (o *WorkflowHeadlessExecutionConnection) GetConnectionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionConnection) GetConnectionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *WorkflowHeadlessExecutionConnection) SetConnectionId(v uuid.UUID) {
	o.ConnectionId = v
}

// GetLabel returns the Label field value.
func (o *WorkflowHeadlessExecutionConnection) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionConnection) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *WorkflowHeadlessExecutionConnection) SetLabel(v string) {
	o.Label = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowHeadlessExecutionConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["label"] = o.Label

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowHeadlessExecutionConnection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionId *uuid.UUID `json:"connection_id"`
		Label        *string    `json:"label"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionId == nil {
		return fmt.Errorf("required field connection_id missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connection_id", "label"})
	} else {
		return err
	}
	o.ConnectionId = *all.ConnectionId
	o.Label = *all.Label

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
