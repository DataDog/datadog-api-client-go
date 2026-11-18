// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageAction Where a message action appears.
type CostRecommendationNotificationRuleMessageAction struct {
	// Whether the action appears once for the whole digest.
	Bulk *bool `json:"bulk,omitempty"`
	// Whether the action appears on each recommendation.
	PerRec *bool `json:"per_rec,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleMessageAction instantiates a new CostRecommendationNotificationRuleMessageAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleMessageAction() *CostRecommendationNotificationRuleMessageAction {
	this := CostRecommendationNotificationRuleMessageAction{}
	return &this
}

// NewCostRecommendationNotificationRuleMessageActionWithDefaults instantiates a new CostRecommendationNotificationRuleMessageAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleMessageActionWithDefaults() *CostRecommendationNotificationRuleMessageAction {
	this := CostRecommendationNotificationRuleMessageAction{}
	return &this
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageAction) GetBulk() bool {
	if o == nil || o.Bulk == nil {
		var ret bool
		return ret
	}
	return *o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageAction) GetBulkOk() (*bool, bool) {
	if o == nil || o.Bulk == nil {
		return nil, false
	}
	return o.Bulk, true
}

// HasBulk returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageAction) HasBulk() bool {
	return o != nil && o.Bulk != nil
}

// SetBulk gets a reference to the given bool and assigns it to the Bulk field.
func (o *CostRecommendationNotificationRuleMessageAction) SetBulk(v bool) {
	o.Bulk = &v
}

// GetPerRec returns the PerRec field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageAction) GetPerRec() bool {
	if o == nil || o.PerRec == nil {
		var ret bool
		return ret
	}
	return *o.PerRec
}

// GetPerRecOk returns a tuple with the PerRec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageAction) GetPerRecOk() (*bool, bool) {
	if o == nil || o.PerRec == nil {
		return nil, false
	}
	return o.PerRec, true
}

// HasPerRec returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageAction) HasPerRec() bool {
	return o != nil && o.PerRec != nil
}

// SetPerRec gets a reference to the given bool and assigns it to the PerRec field.
func (o *CostRecommendationNotificationRuleMessageAction) SetPerRec(v bool) {
	o.PerRec = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleMessageAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Bulk != nil {
		toSerialize["bulk"] = o.Bulk
	}
	if o.PerRec != nil {
		toSerialize["per_rec"] = o.PerRec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleMessageAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bulk   *bool `json:"bulk,omitempty"`
		PerRec *bool `json:"per_rec,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bulk", "per_rec"})
	} else {
		return err
	}
	o.Bulk = all.Bulk
	o.PerRec = all.PerRec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
