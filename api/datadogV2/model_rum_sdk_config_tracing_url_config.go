// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigTracingUrlConfig Configuration for a URL that should have distributed tracing enabled.
type RUMSdkConfigTracingUrlConfig struct {
	// A match option used for URL or origin pattern matching.
	Match RUMSdkConfigMatchOption `json:"match"`
	// The list of trace propagator types to use for this URL.
	PropagatorTypes []RUMSdkConfigTracingUrlPropagatorType `json:"propagator_types"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMSdkConfigTracingUrlConfig instantiates a new RUMSdkConfigTracingUrlConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMSdkConfigTracingUrlConfig(match RUMSdkConfigMatchOption, propagatorTypes []RUMSdkConfigTracingUrlPropagatorType) *RUMSdkConfigTracingUrlConfig {
	this := RUMSdkConfigTracingUrlConfig{}
	this.Match = match
	this.PropagatorTypes = propagatorTypes
	return &this
}

// NewRUMSdkConfigTracingUrlConfigWithDefaults instantiates a new RUMSdkConfigTracingUrlConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMSdkConfigTracingUrlConfigWithDefaults() *RUMSdkConfigTracingUrlConfig {
	this := RUMSdkConfigTracingUrlConfig{}
	return &this
}

// GetMatch returns the Match field value.
func (o *RUMSdkConfigTracingUrlConfig) GetMatch() RUMSdkConfigMatchOption {
	if o == nil {
		var ret RUMSdkConfigMatchOption
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigTracingUrlConfig) GetMatchOk() (*RUMSdkConfigMatchOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *RUMSdkConfigTracingUrlConfig) SetMatch(v RUMSdkConfigMatchOption) {
	o.Match = v
}

// GetPropagatorTypes returns the PropagatorTypes field value.
func (o *RUMSdkConfigTracingUrlConfig) GetPropagatorTypes() []RUMSdkConfigTracingUrlPropagatorType {
	if o == nil {
		var ret []RUMSdkConfigTracingUrlPropagatorType
		return ret
	}
	return o.PropagatorTypes
}

// GetPropagatorTypesOk returns a tuple with the PropagatorTypes field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigTracingUrlConfig) GetPropagatorTypesOk() (*[]RUMSdkConfigTracingUrlPropagatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropagatorTypes, true
}

// SetPropagatorTypes sets field value.
func (o *RUMSdkConfigTracingUrlConfig) SetPropagatorTypes(v []RUMSdkConfigTracingUrlPropagatorType) {
	o.PropagatorTypes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMSdkConfigTracingUrlConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match"] = o.Match
	toSerialize["propagator_types"] = o.PropagatorTypes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMSdkConfigTracingUrlConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match           *RUMSdkConfigMatchOption                `json:"match"`
		PropagatorTypes *[]RUMSdkConfigTracingUrlPropagatorType `json:"propagator_types"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}
	if all.PropagatorTypes == nil {
		return fmt.Errorf("required field propagator_types missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match", "propagator_types"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Match.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Match = *all.Match
	o.PropagatorTypes = *all.PropagatorTypes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
