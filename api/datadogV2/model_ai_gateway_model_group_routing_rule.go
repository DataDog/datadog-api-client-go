// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupRoutingRule Routing rule in an AI Gateway model group.
type AiGatewayModelGroupRoutingRule struct {
	// Organization-unique routing rule slug.
	Slug string `json:"slug"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupRoutingRule instantiates a new AiGatewayModelGroupRoutingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupRoutingRule(slug string) *AiGatewayModelGroupRoutingRule {
	this := AiGatewayModelGroupRoutingRule{}
	this.Slug = slug
	return &this
}

// NewAiGatewayModelGroupRoutingRuleWithDefaults instantiates a new AiGatewayModelGroupRoutingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupRoutingRuleWithDefaults() *AiGatewayModelGroupRoutingRule {
	this := AiGatewayModelGroupRoutingRule{}
	return &this
}

// GetSlug returns the Slug field value.
func (o *AiGatewayModelGroupRoutingRule) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupRoutingRule) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *AiGatewayModelGroupRoutingRule) SetSlug(v string) {
	o.Slug = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupRoutingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelGroupRoutingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Slug *string `json:"slug"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Slug == nil {
		return fmt.Errorf("required field slug missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"slug"})
	} else {
		return err
	}
	o.Slug = *all.Slug

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
