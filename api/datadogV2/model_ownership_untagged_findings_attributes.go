// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipUntaggedFindingsAttributes The counts of findings without a team tag by ownership confidence.
type OwnershipUntaggedFindingsAttributes struct {
	// The number of high confidence findings without a team tag.
	HighConfidence int64 `json:"high_confidence"`
	// The number of low confidence findings without a team tag.
	LowConfidence int64 `json:"low_confidence"`
	// The number of medium confidence findings without a team tag.
	MediumConfidence int64 `json:"medium_confidence"`
	// The total number of findings without a team tag.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipUntaggedFindingsAttributes instantiates a new OwnershipUntaggedFindingsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipUntaggedFindingsAttributes(highConfidence int64, lowConfidence int64, mediumConfidence int64, total int64) *OwnershipUntaggedFindingsAttributes {
	this := OwnershipUntaggedFindingsAttributes{}
	this.HighConfidence = highConfidence
	this.LowConfidence = lowConfidence
	this.MediumConfidence = mediumConfidence
	this.Total = total
	return &this
}

// NewOwnershipUntaggedFindingsAttributesWithDefaults instantiates a new OwnershipUntaggedFindingsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipUntaggedFindingsAttributesWithDefaults() *OwnershipUntaggedFindingsAttributes {
	this := OwnershipUntaggedFindingsAttributes{}
	return &this
}

// GetHighConfidence returns the HighConfidence field value.
func (o *OwnershipUntaggedFindingsAttributes) GetHighConfidence() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.HighConfidence
}

// GetHighConfidenceOk returns a tuple with the HighConfidence field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsAttributes) GetHighConfidenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighConfidence, true
}

// SetHighConfidence sets field value.
func (o *OwnershipUntaggedFindingsAttributes) SetHighConfidence(v int64) {
	o.HighConfidence = v
}

// GetLowConfidence returns the LowConfidence field value.
func (o *OwnershipUntaggedFindingsAttributes) GetLowConfidence() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LowConfidence
}

// GetLowConfidenceOk returns a tuple with the LowConfidence field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsAttributes) GetLowConfidenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowConfidence, true
}

// SetLowConfidence sets field value.
func (o *OwnershipUntaggedFindingsAttributes) SetLowConfidence(v int64) {
	o.LowConfidence = v
}

// GetMediumConfidence returns the MediumConfidence field value.
func (o *OwnershipUntaggedFindingsAttributes) GetMediumConfidence() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MediumConfidence
}

// GetMediumConfidenceOk returns a tuple with the MediumConfidence field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsAttributes) GetMediumConfidenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediumConfidence, true
}

// SetMediumConfidence sets field value.
func (o *OwnershipUntaggedFindingsAttributes) SetMediumConfidence(v int64) {
	o.MediumConfidence = v
}

// GetTotal returns the Total field value.
func (o *OwnershipUntaggedFindingsAttributes) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsAttributes) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *OwnershipUntaggedFindingsAttributes) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipUntaggedFindingsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["high_confidence"] = o.HighConfidence
	toSerialize["low_confidence"] = o.LowConfidence
	toSerialize["medium_confidence"] = o.MediumConfidence
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipUntaggedFindingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HighConfidence   *int64 `json:"high_confidence"`
		LowConfidence    *int64 `json:"low_confidence"`
		MediumConfidence *int64 `json:"medium_confidence"`
		Total            *int64 `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HighConfidence == nil {
		return fmt.Errorf("required field high_confidence missing")
	}
	if all.LowConfidence == nil {
		return fmt.Errorf("required field low_confidence missing")
	}
	if all.MediumConfidence == nil {
		return fmt.Errorf("required field medium_confidence missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"high_confidence", "low_confidence", "medium_confidence", "total"})
	} else {
		return err
	}
	o.HighConfidence = *all.HighConfidence
	o.LowConfidence = *all.LowConfidence
	o.MediumConfidence = *all.MediumConfidence
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
