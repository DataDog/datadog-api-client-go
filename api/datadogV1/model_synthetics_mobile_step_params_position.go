// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsPosition The definition of `SyntheticsMobileStepParamsPosition` object.
type SyntheticsMobileStepParamsPosition struct {
	// The `position` `positions`.
	Positions []SyntheticsMobileStepParamsPositionPositionsItems `json:"positions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsPosition instantiates a new SyntheticsMobileStepParamsPosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsPosition() *SyntheticsMobileStepParamsPosition {
	this := SyntheticsMobileStepParamsPosition{}
	return &this
}

// NewSyntheticsMobileStepParamsPositionWithDefaults instantiates a new SyntheticsMobileStepParamsPosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsPositionWithDefaults() *SyntheticsMobileStepParamsPosition {
	this := SyntheticsMobileStepParamsPosition{}
	return &this
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsPosition) GetPositions() []SyntheticsMobileStepParamsPositionPositionsItems {
	if o == nil || o.Positions == nil {
		var ret []SyntheticsMobileStepParamsPositionPositionsItems
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsPosition) GetPositionsOk() (*[]SyntheticsMobileStepParamsPositionPositionsItems, bool) {
	if o == nil || o.Positions == nil {
		return nil, false
	}
	return &o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsPosition) HasPositions() bool {
	return o != nil && o.Positions != nil
}

// SetPositions gets a reference to the given []SyntheticsMobileStepParamsPositionPositionsItems and assigns it to the Positions field.
func (o *SyntheticsMobileStepParamsPosition) SetPositions(v []SyntheticsMobileStepParamsPositionPositionsItems) {
	o.Positions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Positions != nil {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParamsPosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Positions []SyntheticsMobileStepParamsPositionPositionsItems `json:"positions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"positions"})
	} else {
		return err
	}
	o.Positions = all.Positions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
