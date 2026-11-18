// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageActions Links the digest shows. `view_in_ccm` is the link to review the recommendations in Cloud Cost Management.
type CostRecommendationNotificationRuleMessageActions struct {
	// Where a message action appears.
	ViewInCcm *CostRecommendationNotificationRuleMessageAction `json:"view_in_ccm,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleMessageActions instantiates a new CostRecommendationNotificationRuleMessageActions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleMessageActions() *CostRecommendationNotificationRuleMessageActions {
	this := CostRecommendationNotificationRuleMessageActions{}
	return &this
}

// NewCostRecommendationNotificationRuleMessageActionsWithDefaults instantiates a new CostRecommendationNotificationRuleMessageActions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleMessageActionsWithDefaults() *CostRecommendationNotificationRuleMessageActions {
	this := CostRecommendationNotificationRuleMessageActions{}
	return &this
}

// GetViewInCcm returns the ViewInCcm field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageActions) GetViewInCcm() CostRecommendationNotificationRuleMessageAction {
	if o == nil || o.ViewInCcm == nil {
		var ret CostRecommendationNotificationRuleMessageAction
		return ret
	}
	return *o.ViewInCcm
}

// GetViewInCcmOk returns a tuple with the ViewInCcm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageActions) GetViewInCcmOk() (*CostRecommendationNotificationRuleMessageAction, bool) {
	if o == nil || o.ViewInCcm == nil {
		return nil, false
	}
	return o.ViewInCcm, true
}

// HasViewInCcm returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageActions) HasViewInCcm() bool {
	return o != nil && o.ViewInCcm != nil
}

// SetViewInCcm gets a reference to the given CostRecommendationNotificationRuleMessageAction and assigns it to the ViewInCcm field.
func (o *CostRecommendationNotificationRuleMessageActions) SetViewInCcm(v CostRecommendationNotificationRuleMessageAction) {
	o.ViewInCcm = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleMessageActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ViewInCcm != nil {
		toSerialize["view_in_ccm"] = o.ViewInCcm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleMessageActions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ViewInCcm *CostRecommendationNotificationRuleMessageAction `json:"view_in_ccm,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"view_in_ccm"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ViewInCcm != nil && all.ViewInCcm.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewInCcm = all.ViewInCcm

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
