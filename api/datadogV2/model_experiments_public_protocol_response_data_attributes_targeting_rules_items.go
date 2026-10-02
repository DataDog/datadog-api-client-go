// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems A targeting rule supplied by the protocol.
type ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems struct {
	// Conditions that define this targeting rule.
	Conditions []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems `json:"conditions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems instantiates a new ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems() *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems {
	this := ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems {
	this := ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) GetConditions() []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems {
	if o == nil || o.Conditions == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) GetConditionsOk() (*[]ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems and assigns it to the Conditions field.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) SetConditions(v []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems) {
	o.Conditions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions []ExperimentsPublicProtocolResponseDataAttributesTargetingRulesItemsConditionsItems `json:"conditions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions"})
	} else {
		return err
	}
	o.Conditions = all.Conditions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
