// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleUpdateData Changes to a recommendation notification rule.
type CostRecommendationNotificationRuleUpdateData struct {
	// Attributes to change. Attributes left out keep their value; `mention_targets` and `message_customization`
	// replace the stored value when sent.
	Attributes CostRecommendationNotificationRuleUpdateAttributes `json:"attributes"`
	// The ID of the notification rule. It must match the ID in the path.
	Id string `json:"id"`
	// Recommendation notification rule resource type.
	Type CostRecommendationNotificationRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleUpdateData instantiates a new CostRecommendationNotificationRuleUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleUpdateData(attributes CostRecommendationNotificationRuleUpdateAttributes, id string, typeVar CostRecommendationNotificationRuleType) *CostRecommendationNotificationRuleUpdateData {
	this := CostRecommendationNotificationRuleUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCostRecommendationNotificationRuleUpdateDataWithDefaults instantiates a new CostRecommendationNotificationRuleUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleUpdateDataWithDefaults() *CostRecommendationNotificationRuleUpdateData {
	this := CostRecommendationNotificationRuleUpdateData{}
	var typeVar CostRecommendationNotificationRuleType = COSTRECOMMENDATIONNOTIFICATIONRULETYPE_NOTIFICATION_RULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CostRecommendationNotificationRuleUpdateData) GetAttributes() CostRecommendationNotificationRuleUpdateAttributes {
	if o == nil {
		var ret CostRecommendationNotificationRuleUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleUpdateData) GetAttributesOk() (*CostRecommendationNotificationRuleUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CostRecommendationNotificationRuleUpdateData) SetAttributes(v CostRecommendationNotificationRuleUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CostRecommendationNotificationRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CostRecommendationNotificationRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CostRecommendationNotificationRuleUpdateData) GetType() CostRecommendationNotificationRuleType {
	if o == nil {
		var ret CostRecommendationNotificationRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleUpdateData) GetTypeOk() (*CostRecommendationNotificationRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CostRecommendationNotificationRuleUpdateData) SetType(v CostRecommendationNotificationRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CostRecommendationNotificationRuleUpdateAttributes `json:"attributes"`
		Id         *string                                             `json:"id"`
		Type       *CostRecommendationNotificationRuleType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
