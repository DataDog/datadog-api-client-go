// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeDuration20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type ChangeDuration20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomTimeframe FixedSpan2 `json:"custom_timeframe"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ChangeDurationAnyOfAnyOfItemType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeDuration20270101OneOf instantiates a new ChangeDuration20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeDuration20270101OneOf(customTimeframe FixedSpan2, typeVar ChangeDurationAnyOfAnyOfItemType) *ChangeDuration20270101OneOf {
	this := ChangeDuration20270101OneOf{}
	this.CustomTimeframe = customTimeframe
	this.Type = typeVar
	return &this
}

// NewChangeDuration20270101OneOfWithDefaults instantiates a new ChangeDuration20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeDuration20270101OneOfWithDefaults() *ChangeDuration20270101OneOf {
	this := ChangeDuration20270101OneOf{}
	return &this
}

// GetCustomTimeframe returns the CustomTimeframe field value.
func (o *ChangeDuration20270101OneOf) GetCustomTimeframe() FixedSpan2 {
	if o == nil {
		var ret FixedSpan2
		return ret
	}
	return o.CustomTimeframe
}

// GetCustomTimeframeOk returns a tuple with the CustomTimeframe field value
// and a boolean to check if the value has been set.
func (o *ChangeDuration20270101OneOf) GetCustomTimeframeOk() (*FixedSpan2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomTimeframe, true
}

// SetCustomTimeframe sets field value.
func (o *ChangeDuration20270101OneOf) SetCustomTimeframe(v FixedSpan2) {
	o.CustomTimeframe = v
}

// GetType returns the Type field value.
func (o *ChangeDuration20270101OneOf) GetType() ChangeDurationAnyOfAnyOfItemType {
	if o == nil {
		var ret ChangeDurationAnyOfAnyOfItemType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeDuration20270101OneOf) GetTypeOk() (*ChangeDurationAnyOfAnyOfItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeDuration20270101OneOf) SetType(v ChangeDurationAnyOfAnyOfItemType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeDuration20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["custom_timeframe"] = o.CustomTimeframe
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeDuration20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTimeframe *FixedSpan2                       `json:"custom_timeframe"`
		Type            *ChangeDurationAnyOfAnyOfItemType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CustomTimeframe == nil {
		return fmt.Errorf("required field custom_timeframe missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.CustomTimeframe.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CustomTimeframe = *all.CustomTimeframe
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
