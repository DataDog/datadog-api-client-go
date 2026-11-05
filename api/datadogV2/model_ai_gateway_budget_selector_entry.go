// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetSelectorEntry One dimension rule in a template budget selector.
type AiGatewayBudgetSelectorEntry struct {
	// Dimension to match: user_id, team_id, model, provider, or tag:<key>. Reserved metadata tag keys are rejected.
	Dim string `json:"dim"`
	// Separate counters by the matched value when true. Otherwise, pool spend within the other separated dimensions.
	PerValue bool `json:"per_value"`
	// Allowed dimension values. Omit to match any value, including an absent dimension. An explicit list must contain at least one value.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetSelectorEntry instantiates a new AiGatewayBudgetSelectorEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetSelectorEntry(dim string, perValue bool) *AiGatewayBudgetSelectorEntry {
	this := AiGatewayBudgetSelectorEntry{}
	this.Dim = dim
	this.PerValue = perValue
	return &this
}

// NewAiGatewayBudgetSelectorEntryWithDefaults instantiates a new AiGatewayBudgetSelectorEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetSelectorEntryWithDefaults() *AiGatewayBudgetSelectorEntry {
	this := AiGatewayBudgetSelectorEntry{}
	var perValue bool = false
	this.PerValue = perValue
	return &this
}

// GetDim returns the Dim field value.
func (o *AiGatewayBudgetSelectorEntry) GetDim() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dim
}

// GetDimOk returns a tuple with the Dim field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetSelectorEntry) GetDimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dim, true
}

// SetDim sets field value.
func (o *AiGatewayBudgetSelectorEntry) SetDim(v string) {
	o.Dim = v
}

// GetPerValue returns the PerValue field value.
func (o *AiGatewayBudgetSelectorEntry) GetPerValue() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PerValue
}

// GetPerValueOk returns a tuple with the PerValue field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetSelectorEntry) GetPerValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerValue, true
}

// SetPerValue sets field value.
func (o *AiGatewayBudgetSelectorEntry) SetPerValue(v bool) {
	o.PerValue = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AiGatewayBudgetSelectorEntry) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetSelectorEntry) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AiGatewayBudgetSelectorEntry) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *AiGatewayBudgetSelectorEntry) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetSelectorEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dim"] = o.Dim
	toSerialize["per_value"] = o.PerValue
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetSelectorEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dim      *string  `json:"dim"`
		PerValue *bool    `json:"per_value"`
		Values   []string `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dim == nil {
		return fmt.Errorf("required field dim missing")
	}
	if all.PerValue == nil {
		return fmt.Errorf("required field per_value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dim", "per_value", "values"})
	} else {
		return err
	}
	o.Dim = *all.Dim
	o.PerValue = *all.PerValue
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
