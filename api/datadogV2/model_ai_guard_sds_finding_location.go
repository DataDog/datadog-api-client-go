// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardSdsFindingLocation The location of a sensitive data match within the evaluated request.
type AIGuardSdsFindingLocation struct {
	// The end character index (exclusive) of the sensitive data match.
	EndIndexExclusive int64 `json:"end_index_exclusive"`
	// The JSON path to the field containing the sensitive data.
	Path string `json:"path"`
	// The start character index of the sensitive data match.
	StartIndex int64 `json:"start_index"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardSdsFindingLocation instantiates a new AIGuardSdsFindingLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardSdsFindingLocation(endIndexExclusive int64, path string, startIndex int64) *AIGuardSdsFindingLocation {
	this := AIGuardSdsFindingLocation{}
	this.EndIndexExclusive = endIndexExclusive
	this.Path = path
	this.StartIndex = startIndex
	return &this
}

// NewAIGuardSdsFindingLocationWithDefaults instantiates a new AIGuardSdsFindingLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardSdsFindingLocationWithDefaults() *AIGuardSdsFindingLocation {
	this := AIGuardSdsFindingLocation{}
	return &this
}

// GetEndIndexExclusive returns the EndIndexExclusive field value.
func (o *AIGuardSdsFindingLocation) GetEndIndexExclusive() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EndIndexExclusive
}

// GetEndIndexExclusiveOk returns a tuple with the EndIndexExclusive field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFindingLocation) GetEndIndexExclusiveOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndexExclusive, true
}

// SetEndIndexExclusive sets field value.
func (o *AIGuardSdsFindingLocation) SetEndIndexExclusive(v int64) {
	o.EndIndexExclusive = v
}

// GetPath returns the Path field value.
func (o *AIGuardSdsFindingLocation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFindingLocation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *AIGuardSdsFindingLocation) SetPath(v string) {
	o.Path = v
}

// GetStartIndex returns the StartIndex field value.
func (o *AIGuardSdsFindingLocation) GetStartIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *AIGuardSdsFindingLocation) GetStartIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value.
func (o *AIGuardSdsFindingLocation) SetStartIndex(v int64) {
	o.StartIndex = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardSdsFindingLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end_index_exclusive"] = o.EndIndexExclusive
	toSerialize["path"] = o.Path
	toSerialize["start_index"] = o.StartIndex

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardSdsFindingLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndIndexExclusive *int64  `json:"end_index_exclusive"`
		Path              *string `json:"path"`
		StartIndex        *int64  `json:"start_index"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EndIndexExclusive == nil {
		return fmt.Errorf("required field end_index_exclusive missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.StartIndex == nil {
		return fmt.Errorf("required field start_index missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_index_exclusive", "path", "start_index"})
	} else {
		return err
	}
	o.EndIndexExclusive = *all.EndIndexExclusive
	o.Path = *all.Path
	o.StartIndex = *all.StartIndex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
