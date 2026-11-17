// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelAutomaticPricingMapping The reviewed endpoint mapping used to identify a custom provider. Omitted when the provider and model identity matched exactly.
type AiGatewayModelAutomaticPricingMapping struct {
	// Mapping identifier.
	Id string `json:"id"`
	// Mapping revision.
	Revision string `json:"revision"`
	// Mapping scope, such as `provider`.
	Scope string `json:"scope"`
	// Mapping version.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelAutomaticPricingMapping instantiates a new AiGatewayModelAutomaticPricingMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelAutomaticPricingMapping(id string, revision string, scope string, version int64) *AiGatewayModelAutomaticPricingMapping {
	this := AiGatewayModelAutomaticPricingMapping{}
	this.Id = id
	this.Revision = revision
	this.Scope = scope
	this.Version = version
	return &this
}

// NewAiGatewayModelAutomaticPricingMappingWithDefaults instantiates a new AiGatewayModelAutomaticPricingMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelAutomaticPricingMappingWithDefaults() *AiGatewayModelAutomaticPricingMapping {
	this := AiGatewayModelAutomaticPricingMapping{}
	return &this
}

// GetId returns the Id field value.
func (o *AiGatewayModelAutomaticPricingMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayModelAutomaticPricingMapping) SetId(v string) {
	o.Id = v
}

// GetRevision returns the Revision field value.
func (o *AiGatewayModelAutomaticPricingMapping) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingMapping) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayModelAutomaticPricingMapping) SetRevision(v string) {
	o.Revision = v
}

// GetScope returns the Scope field value.
func (o *AiGatewayModelAutomaticPricingMapping) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingMapping) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiGatewayModelAutomaticPricingMapping) SetScope(v string) {
	o.Scope = v
}

// GetVersion returns the Version field value.
func (o *AiGatewayModelAutomaticPricingMapping) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingMapping) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *AiGatewayModelAutomaticPricingMapping) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelAutomaticPricingMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["scope"] = o.Scope
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelAutomaticPricingMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *string `json:"id"`
		Revision *string `json:"revision"`
		Scope    *string `json:"scope"`
		Version  *int64  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Revision == nil {
		return fmt.Errorf("required field revision missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "revision", "scope", "version"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Revision = *all.Revision
	o.Scope = *all.Scope
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
