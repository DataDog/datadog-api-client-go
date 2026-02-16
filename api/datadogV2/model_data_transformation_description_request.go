// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformationDescriptionRequest
type DataTransformationDescriptionRequest struct {
	// The fully qualified name (FQN) of the action.
	ActionId string `json:"actionId"`
	// The transformation script to describe.
	Script string `json:"script"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataTransformationDescriptionRequest instantiates a new DataTransformationDescriptionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataTransformationDescriptionRequest(actionId string, script string) *DataTransformationDescriptionRequest {
	this := DataTransformationDescriptionRequest{}
	this.ActionId = actionId
	this.Script = script
	return &this
}

// NewDataTransformationDescriptionRequestWithDefaults instantiates a new DataTransformationDescriptionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataTransformationDescriptionRequestWithDefaults() *DataTransformationDescriptionRequest {
	this := DataTransformationDescriptionRequest{}
	return &this
}

// GetActionId returns the ActionId field value.
func (o *DataTransformationDescriptionRequest) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *DataTransformationDescriptionRequest) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value.
func (o *DataTransformationDescriptionRequest) SetActionId(v string) {
	o.ActionId = v
}

// GetScript returns the Script field value.
func (o *DataTransformationDescriptionRequest) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *DataTransformationDescriptionRequest) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value.
func (o *DataTransformationDescriptionRequest) SetScript(v string) {
	o.Script = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataTransformationDescriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actionId"] = o.ActionId
	toSerialize["script"] = o.Script

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataTransformationDescriptionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionId *string `json:"actionId"`
		Script   *string `json:"script"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionId == nil {
		return fmt.Errorf("required field actionId missing")
	}
	if all.Script == nil {
		return fmt.Errorf("required field script missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actionId", "script"})
	} else {
		return err
	}
	o.ActionId = *all.ActionId
	o.Script = *all.Script

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
