// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowScaffoldAgenticStreamResponse
type WorkflowScaffoldAgenticStreamResponse struct {
	//
	Event StreamEventV1 `json:"event"`
	// The generation ID.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowScaffoldAgenticStreamResponse instantiates a new WorkflowScaffoldAgenticStreamResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowScaffoldAgenticStreamResponse(event StreamEventV1, id string) *WorkflowScaffoldAgenticStreamResponse {
	this := WorkflowScaffoldAgenticStreamResponse{}
	this.Event = event
	this.Id = id
	return &this
}

// NewWorkflowScaffoldAgenticStreamResponseWithDefaults instantiates a new WorkflowScaffoldAgenticStreamResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowScaffoldAgenticStreamResponseWithDefaults() *WorkflowScaffoldAgenticStreamResponse {
	this := WorkflowScaffoldAgenticStreamResponse{}
	return &this
}

// GetEvent returns the Event field value.
func (o *WorkflowScaffoldAgenticStreamResponse) GetEvent() StreamEventV1 {
	if o == nil {
		var ret StreamEventV1
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamResponse) GetEventOk() (*StreamEventV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value.
func (o *WorkflowScaffoldAgenticStreamResponse) SetEvent(v StreamEventV1) {
	o.Event = v
}

// GetId returns the Id field value.
func (o *WorkflowScaffoldAgenticStreamResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowScaffoldAgenticStreamResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *WorkflowScaffoldAgenticStreamResponse) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowScaffoldAgenticStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event"] = o.Event
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowScaffoldAgenticStreamResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Event *StreamEventV1 `json:"event"`
		Id    *string        `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Event == nil {
		return fmt.Errorf("required field event missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event", "id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Event.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Event = *all.Event
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
