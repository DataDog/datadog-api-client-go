// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardSdsFinding A sensitive data finding detected by the SDS scanner.
type AIGuardSdsFinding struct {
	// The category of sensitive data detected.
	Category string `json:"category"`
	// The location of a sensitive data match within the evaluated request.
	Location AIGuardSdsFindingLocation `json:"location"`
	// The human-readable name of the SDS rule that triggered.
	RuleDisplayName string `json:"rule_display_name"`
	// The tag identifier of the SDS rule that triggered.
	RuleTag string `json:"rule_tag"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardSdsFinding instantiates a new AIGuardSdsFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardSdsFinding(category string, location AIGuardSdsFindingLocation, ruleDisplayName string, ruleTag string) *AIGuardSdsFinding {
	this := AIGuardSdsFinding{}
	this.Category = category
	this.Location = location
	this.RuleDisplayName = ruleDisplayName
	this.RuleTag = ruleTag
	return &this
}

// NewAIGuardSdsFindingWithDefaults instantiates a new AIGuardSdsFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardSdsFindingWithDefaults() *AIGuardSdsFinding {
	this := AIGuardSdsFinding{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AIGuardSdsFinding) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFinding) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AIGuardSdsFinding) SetCategory(v string) {
	o.Category = v
}

// GetLocation returns the Location field value.
func (o *AIGuardSdsFinding) GetLocation() AIGuardSdsFindingLocation {
	if o == nil {
		var ret AIGuardSdsFindingLocation
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFinding) GetLocationOk() (*AIGuardSdsFindingLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value.
func (o *AIGuardSdsFinding) SetLocation(v AIGuardSdsFindingLocation) {
	o.Location = v
}

// GetRuleDisplayName returns the RuleDisplayName field value.
func (o *AIGuardSdsFinding) GetRuleDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleDisplayName
}

// GetRuleDisplayNameOk returns a tuple with the RuleDisplayName field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFinding) GetRuleDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleDisplayName, true
}

// SetRuleDisplayName sets field value.
func (o *AIGuardSdsFinding) SetRuleDisplayName(v string) {
	o.RuleDisplayName = v
}

// GetRuleTag returns the RuleTag field value.
func (o *AIGuardSdsFinding) GetRuleTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleTag
}

// GetRuleTagOk returns a tuple with the RuleTag field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFinding) GetRuleTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleTag, true
}

// SetRuleTag sets field value.
func (o *AIGuardSdsFinding) SetRuleTag(v string) {
	o.RuleTag = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardSdsFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["location"] = o.Location
	toSerialize["rule_display_name"] = o.RuleDisplayName
	toSerialize["rule_tag"] = o.RuleTag

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardSdsFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category        *string                    `json:"category"`
		Location        *AIGuardSdsFindingLocation `json:"location"`
		RuleDisplayName *string                    `json:"rule_display_name"`
		RuleTag         *string                    `json:"rule_tag"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Location == nil {
		return fmt.Errorf("required field location missing")
	}
	if all.RuleDisplayName == nil {
		return fmt.Errorf("required field rule_display_name missing")
	}
	if all.RuleTag == nil {
		return fmt.Errorf("required field rule_tag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "location", "rule_display_name", "rule_tag"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = *all.Category
	if all.Location.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Location = *all.Location
	o.RuleDisplayName = *all.RuleDisplayName
	o.RuleTag = *all.RuleTag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
