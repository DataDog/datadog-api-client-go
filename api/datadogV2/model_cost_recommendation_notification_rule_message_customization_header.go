// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationNotificationRuleMessageCustomizationHeader What the digest header shows.
type CostRecommendationNotificationRuleMessageCustomizationHeader struct {
	// Text shown in the header, up to 500 characters. It can use the `{{ total_savings }}`, `{{ rec_count }}`,
	// and `{{ digest_date }}` variables.
	CustomText *string `json:"custom_text,omitempty"`
	// Whether the header shows the rule name.
	NotificationName *bool `json:"notification_name,omitempty"`
	// Whether the header shows the total potential savings.
	TotalSavings *bool `json:"total_savings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationNotificationRuleMessageCustomizationHeader instantiates a new CostRecommendationNotificationRuleMessageCustomizationHeader object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationNotificationRuleMessageCustomizationHeader() *CostRecommendationNotificationRuleMessageCustomizationHeader {
	this := CostRecommendationNotificationRuleMessageCustomizationHeader{}
	return &this
}

// NewCostRecommendationNotificationRuleMessageCustomizationHeaderWithDefaults instantiates a new CostRecommendationNotificationRuleMessageCustomizationHeader object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationNotificationRuleMessageCustomizationHeaderWithDefaults() *CostRecommendationNotificationRuleMessageCustomizationHeader {
	this := CostRecommendationNotificationRuleMessageCustomizationHeader{}
	return &this
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetCustomText() string {
	if o == nil || o.CustomText == nil {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetCustomTextOk() (*string, bool) {
	if o == nil || o.CustomText == nil {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) HasCustomText() bool {
	return o != nil && o.CustomText != nil
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) SetCustomText(v string) {
	o.CustomText = &v
}

// GetNotificationName returns the NotificationName field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetNotificationName() bool {
	if o == nil || o.NotificationName == nil {
		var ret bool
		return ret
	}
	return *o.NotificationName
}

// GetNotificationNameOk returns a tuple with the NotificationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetNotificationNameOk() (*bool, bool) {
	if o == nil || o.NotificationName == nil {
		return nil, false
	}
	return o.NotificationName, true
}

// HasNotificationName returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) HasNotificationName() bool {
	return o != nil && o.NotificationName != nil
}

// SetNotificationName gets a reference to the given bool and assigns it to the NotificationName field.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) SetNotificationName(v bool) {
	o.NotificationName = &v
}

// GetTotalSavings returns the TotalSavings field value if set, zero value otherwise.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetTotalSavings() bool {
	if o == nil || o.TotalSavings == nil {
		var ret bool
		return ret
	}
	return *o.TotalSavings
}

// GetTotalSavingsOk returns a tuple with the TotalSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) GetTotalSavingsOk() (*bool, bool) {
	if o == nil || o.TotalSavings == nil {
		return nil, false
	}
	return o.TotalSavings, true
}

// HasTotalSavings returns a boolean if a field has been set.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) HasTotalSavings() bool {
	return o != nil && o.TotalSavings != nil
}

// SetTotalSavings gets a reference to the given bool and assigns it to the TotalSavings field.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) SetTotalSavings(v bool) {
	o.TotalSavings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationNotificationRuleMessageCustomizationHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomText != nil {
		toSerialize["custom_text"] = o.CustomText
	}
	if o.NotificationName != nil {
		toSerialize["notification_name"] = o.NotificationName
	}
	if o.TotalSavings != nil {
		toSerialize["total_savings"] = o.TotalSavings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationNotificationRuleMessageCustomizationHeader) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomText       *string `json:"custom_text,omitempty"`
		NotificationName *bool   `json:"notification_name,omitempty"`
		TotalSavings     *bool   `json:"total_savings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_text", "notification_name", "total_savings"})
	} else {
		return err
	}
	o.CustomText = all.CustomText
	o.NotificationName = all.NotificationName
	o.TotalSavings = all.TotalSavings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
