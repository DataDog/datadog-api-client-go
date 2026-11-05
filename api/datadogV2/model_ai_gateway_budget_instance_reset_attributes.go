// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstanceResetAttributes Attributes used to reset one budget instance's active-period spend.
type AiGatewayBudgetInstanceResetAttributes struct {
	// Counter key of the instance to reset, as returned by the list budget instances endpoint.
	InstanceKey string `json:"instance_key"`
	// Optional reason recorded with the spend reset.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetInstanceResetAttributes instantiates a new AiGatewayBudgetInstanceResetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetInstanceResetAttributes(instanceKey string) *AiGatewayBudgetInstanceResetAttributes {
	this := AiGatewayBudgetInstanceResetAttributes{}
	this.InstanceKey = instanceKey
	return &this
}

// NewAiGatewayBudgetInstanceResetAttributesWithDefaults instantiates a new AiGatewayBudgetInstanceResetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetInstanceResetAttributesWithDefaults() *AiGatewayBudgetInstanceResetAttributes {
	this := AiGatewayBudgetInstanceResetAttributes{}
	return &this
}

// GetInstanceKey returns the InstanceKey field value.
func (o *AiGatewayBudgetInstanceResetAttributes) GetInstanceKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceKey
}

// GetInstanceKeyOk returns a tuple with the InstanceKey field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceResetAttributes) GetInstanceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceKey, true
}

// SetInstanceKey sets field value.
func (o *AiGatewayBudgetInstanceResetAttributes) SetInstanceKey(v string) {
	o.InstanceKey = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiGatewayBudgetInstanceResetAttributes) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceResetAttributes) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiGatewayBudgetInstanceResetAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiGatewayBudgetInstanceResetAttributes) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetInstanceResetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_key"] = o.InstanceKey
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetInstanceResetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceKey *string `json:"instance_key"`
		Reason      *string `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceKey == nil {
		return fmt.Errorf("required field instance_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_key", "reason"})
	} else {
		return err
	}
	o.InstanceKey = *all.InstanceKey
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
