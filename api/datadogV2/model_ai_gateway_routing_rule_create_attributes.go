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
	// Optional description of what the routing rule is for. Up to 500 characters. Line breaks are allowed.
	// Leading and trailing spaces are trimmed, and a blank value is the same as none.
	Description datadog.NullableString `json:"description,omitempty"`
	// Optional human-readable name shown in place of the slug. Up to 100 characters on one line. Need not be unique.
	// Leading and trailing spaces are trimmed, and a blank value is the same as none.
	DisplayName datadog.NullableString `json:"display_name,omitempty"`
	// Reserved for initial View and Edit permissions. Supplying this attribute returns HTTP 400 and doesn't
	// create the route. Create the route, then set its permissions through the restriction policy API.
	InitialBindings []AiGatewayProviderInitialBinding `json:"initial_bindings,omitempty"`
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

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleCreateAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *AiGatewayRoutingRuleCreateAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *AiGatewayRoutingRuleCreateAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *AiGatewayRoutingRuleCreateAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayRoutingRuleCreateAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayRoutingRuleCreateAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName.IsSet()
}

// SetDisplayName gets a reference to the given datadog.NullableString and assigns it to the DisplayName field.
func (o *AiGatewayRoutingRuleCreateAttributes) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil.
func (o *AiGatewayRoutingRuleCreateAttributes) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil.
func (o *AiGatewayRoutingRuleCreateAttributes) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetInitialBindings returns the InitialBindings field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleCreateAttributes) GetInitialBindings() []AiGatewayProviderInitialBinding {
	if o == nil || o.InitialBindings == nil {
		var ret []AiGatewayProviderInitialBinding
		return ret
	}
	return o.InitialBindings
}

// GetInitialBindingsOk returns a tuple with the InitialBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) GetInitialBindingsOk() (*[]AiGatewayProviderInitialBinding, bool) {
	if o == nil || o.InitialBindings == nil {
		return nil, false
	}
	return &o.InitialBindings, true
}

// HasInitialBindings returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleCreateAttributes) HasInitialBindings() bool {
	return o != nil && o.InitialBindings != nil
}

// SetInitialBindings gets a reference to the given []AiGatewayProviderInitialBinding and assigns it to the InitialBindings field.
func (o *AiGatewayRoutingRuleCreateAttributes) SetInitialBindings(v []AiGatewayProviderInitialBinding) {
	o.InitialBindings = v
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.InitialBindings != nil {
		toSerialize["initial_bindings"] = o.InitialBindings
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
		Description     datadog.NullableString             `json:"description,omitempty"`
		DisplayName     datadog.NullableString             `json:"display_name,omitempty"`
		InitialBindings []AiGatewayProviderInitialBinding  `json:"initial_bindings,omitempty"`
		Slug            *string                            `json:"slug"`
		Tiers           *[]AiGatewayRoutingRuleTierRequest `json:"tiers"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "initial_bindings", "slug", "tiers"})
	} else {
		return err
	}
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.InitialBindings = all.InitialBindings
	o.Slug = *all.Slug
	o.Tiers = *all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
