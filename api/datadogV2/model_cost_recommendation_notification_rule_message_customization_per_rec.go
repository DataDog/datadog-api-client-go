// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageCustomizationPerRec What each recommendation in the digest shows.
type CostRecommendationNotificationRuleMessageCustomizationPerRec struct {
	// Whether each recommendation shows its effort.
	Effort *bool `json:"effort,omitempty"`
	// Whether each recommendation shows its risk.
	Risk *bool `json:"risk,omitempty"`
	// Whether each recommendation shows why it was made.
	Why *bool `json:"why,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleMessageCustomizationPerRec instantiates a new CostRecommendationNotificationRuleMessageCustomizationPerRec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleMessageCustomizationPerRec() *CostRecommendationNotificationRuleMessageCustomizationPerRec {
	this := CostRecommendationNotificationRuleMessageCustomizationPerRec{}
	return &this
}

// NewCostRecommendationNotificationRuleMessageCustomizationPerRecWithDefaults instantiates a new CostRecommendationNotificationRuleMessageCustomizationPerRec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleMessageCustomizationPerRecWithDefaults() *CostRecommendationNotificationRuleMessageCustomizationPerRec {
	this := CostRecommendationNotificationRuleMessageCustomizationPerRec{}
	return &this
}

// GetEffort returns the Effort field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetEffort() bool {
	if o == nil || o.Effort == nil {
		var ret bool
		return ret
	}
	return *o.Effort
}

// GetEffortOk returns a tuple with the Effort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetEffortOk() (*bool, bool) {
	if o == nil || o.Effort == nil {
		return nil, false
	}
	return o.Effort, true
}

// HasEffort returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) HasEffort() bool {
	return o != nil && o.Effort != nil
}

// SetEffort gets a reference to the given bool and assigns it to the Effort field.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) SetEffort(v bool) {
	o.Effort = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetRisk() bool {
	if o == nil || o.Risk == nil {
		var ret bool
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetRiskOk() (*bool, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given bool and assigns it to the Risk field.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) SetRisk(v bool) {
	o.Risk = &v
}

// GetWhy returns the Why field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetWhy() bool {
	if o == nil || o.Why == nil {
		var ret bool
		return ret
	}
	return *o.Why
}

// GetWhyOk returns a tuple with the Why field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) GetWhyOk() (*bool, bool) {
	if o == nil || o.Why == nil {
		return nil, false
	}
	return o.Why, true
}

// HasWhy returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) HasWhy() bool {
	return o != nil && o.Why != nil
}

// SetWhy gets a reference to the given bool and assigns it to the Why field.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) SetWhy(v bool) {
	o.Why = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleMessageCustomizationPerRec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Effort != nil {
		toSerialize["effort"] = o.Effort
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.Why != nil {
		toSerialize["why"] = o.Why
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleMessageCustomizationPerRec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Effort *bool `json:"effort,omitempty"`
		Risk   *bool `json:"risk,omitempty"`
		Why    *bool `json:"why,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effort", "risk", "why"})
	} else {
		return err
	}
	o.Effort = all.Effort
	o.Risk = all.Risk
	o.Why = all.Why

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
