// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsStopBanditV2RequestDataAttributes The variant to serve once the bandit stops.
type ExperimentsStopBanditV2RequestDataAttributes struct {
	// Key of the flag variant to serve to everyone once the bandit stops. It can be any of the flag's variants, and overrides the configured fallback.
	ServedVariantKey string `json:"served_variant_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsStopBanditV2RequestDataAttributes instantiates a new ExperimentsStopBanditV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsStopBanditV2RequestDataAttributes(servedVariantKey string) *ExperimentsStopBanditV2RequestDataAttributes {
	this := ExperimentsStopBanditV2RequestDataAttributes{}
	this.ServedVariantKey = servedVariantKey
	return &this
}

// NewExperimentsStopBanditV2RequestDataAttributesWithDefaults instantiates a new ExperimentsStopBanditV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsStopBanditV2RequestDataAttributesWithDefaults() *ExperimentsStopBanditV2RequestDataAttributes {
	this := ExperimentsStopBanditV2RequestDataAttributes{}
	return &this
}

// GetServedVariantKey returns the ServedVariantKey field value.
func (o *ExperimentsStopBanditV2RequestDataAttributes) GetServedVariantKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServedVariantKey
}

// GetServedVariantKeyOk returns a tuple with the ServedVariantKey field value
// and a boolean to check if the value has been set.
func (o *ExperimentsStopBanditV2RequestDataAttributes) GetServedVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServedVariantKey, true
}

// SetServedVariantKey sets field value.
func (o *ExperimentsStopBanditV2RequestDataAttributes) SetServedVariantKey(v string) {
	o.ServedVariantKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsStopBanditV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["served_variant_key"] = o.ServedVariantKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsStopBanditV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServedVariantKey *string `json:"served_variant_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServedVariantKey == nil {
		return fmt.Errorf("required field served_variant_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"served_variant_key"})
	} else {
		return err
	}
	o.ServedVariantKey = *all.ServedVariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
