// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleCreateAttributes Attributes used to create an AI Gateway routing rule.
type AiGatewayRoutingRuleCreateAttributes struct {
	// Organization-unique resource slug.
	Slug string `json:"slug"`
	//
	Tiers []AiGatewayRoutingRuleTierRequest `json:"tiers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleCreateAttributes instantiates a new AiGatewayRoutingRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleCreateAttributes(slug string, tiers []AiGatewayRoutingRuleTierRequest) *AiGatewayRoutingRuleCreateAttributes {
	this := AiGatewayRoutingRuleCreateAttributes{}
	this.Slug = slug
	this.Tiers = tiers
	return &this
}

// NewAiGatewayRoutingRuleCreateAttributesWithDefaults instantiates a new AiGatewayRoutingRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleCreateAttributesWithDefaults() *AiGatewayRoutingRuleCreateAttributes {
	this := AiGatewayRoutingRuleCreateAttributes{}
	return &this
}

// GetSlug returns the Slug field value.
func (o *AiGatewayRoutingRuleCreateAttributes) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *AiGatewayRoutingRuleCreateAttributes) SetSlug(v string) {
	o.Slug = v
}

// GetTiers returns the Tiers field value.
func (o *AiGatewayRoutingRuleCreateAttributes) GetTiers() []AiGatewayRoutingRuleTierRequest {
	if o == nil {
		var ret []AiGatewayRoutingRuleTierRequest
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) GetTiersOk() (*[]AiGatewayRoutingRuleTierRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value.
func (o *AiGatewayRoutingRuleCreateAttributes) SetTiers(v []AiGatewayRoutingRuleTierRequest) {
	o.Tiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["slug"] = o.Slug
	toSerialize["tiers"] = o.Tiers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Slug  *string                            `json:"slug"`
		Tiers *[]AiGatewayRoutingRuleTierRequest `json:"tiers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Slug == nil {
		return fmt.Errorf("required field slug missing")
	}
	if all.Tiers == nil {
		return fmt.Errorf("required field tiers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"slug", "tiers"})
	} else {
		return err
	}
	o.Slug = *all.Slug
	o.Tiers = *all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
