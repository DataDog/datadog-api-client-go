// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterTargetingRuleRequest A replacement saved filter targeting rule.
type SavedFilterTargetingRuleRequest struct {
	// Inline conditions combined with AND. Must contain at least one condition.
	Conditions []SavedFilterConditionRequest `json:"conditions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterTargetingRuleRequest instantiates a new SavedFilterTargetingRuleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterTargetingRuleRequest(conditions []SavedFilterConditionRequest) *SavedFilterTargetingRuleRequest {
	this := SavedFilterTargetingRuleRequest{}
	this.Conditions = conditions
	return &this
}

// NewSavedFilterTargetingRuleRequestWithDefaults instantiates a new SavedFilterTargetingRuleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterTargetingRuleRequestWithDefaults() *SavedFilterTargetingRuleRequest {
	this := SavedFilterTargetingRuleRequest{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *SavedFilterTargetingRuleRequest) GetConditions() []SavedFilterConditionRequest {
	if o == nil {
		var ret []SavedFilterConditionRequest
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *SavedFilterTargetingRuleRequest) GetConditionsOk() (*[]SavedFilterConditionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *SavedFilterTargetingRuleRequest) SetConditions(v []SavedFilterConditionRequest) {
	o.Conditions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterTargetingRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterTargetingRuleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]SavedFilterConditionRequest `json:"conditions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions"})
	} else {
		return err
	}
	o.Conditions = *all.Conditions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
