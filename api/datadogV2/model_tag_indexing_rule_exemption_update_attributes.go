// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionUpdateAttributes Attributes for updating a tag indexing rule exemption.
type TagIndexingRuleExemptionUpdateAttributes struct {
	// The updated reason the metric is exempt from tag indexing rules.
	Reason string `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleExemptionUpdateAttributes instantiates a new TagIndexingRuleExemptionUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleExemptionUpdateAttributes(reason string) *TagIndexingRuleExemptionUpdateAttributes {
	this := TagIndexingRuleExemptionUpdateAttributes{}
	this.Reason = reason
	return &this
}

// NewTagIndexingRuleExemptionUpdateAttributesWithDefaults instantiates a new TagIndexingRuleExemptionUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleExemptionUpdateAttributesWithDefaults() *TagIndexingRuleExemptionUpdateAttributes {
	this := TagIndexingRuleExemptionUpdateAttributes{}
	return &this
}

// GetReason returns the Reason field value.
func (o *TagIndexingRuleExemptionUpdateAttributes) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionUpdateAttributes) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *TagIndexingRuleExemptionUpdateAttributes) SetReason(v string) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleExemptionUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleExemptionUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason *string `json:"reason"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reason"})
	} else {
		return err
	}
	o.Reason = *all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
