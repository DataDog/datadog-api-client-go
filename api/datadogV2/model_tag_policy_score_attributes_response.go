// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyScoreAttributesResponse Attributes of a tag policy score.
type TagPolicyScoreAttributesResponse struct {
	// The compliance score for the tag policy.
	Score float64 `json:"score"`
	// End timestamp for the score calculation period.
	TsEnd int64 `json:"ts_end"`
	// Start timestamp for the score calculation period.
	TsStart int64 `json:"ts_start"`
	// The version of the score.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyScoreAttributesResponse instantiates a new TagPolicyScoreAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyScoreAttributesResponse(score float64, tsEnd int64, tsStart int64) *TagPolicyScoreAttributesResponse {
	this := TagPolicyScoreAttributesResponse{}
	this.Score = score
	this.TsEnd = tsEnd
	this.TsStart = tsStart
	return &this
}

// NewTagPolicyScoreAttributesResponseWithDefaults instantiates a new TagPolicyScoreAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyScoreAttributesResponseWithDefaults() *TagPolicyScoreAttributesResponse {
	this := TagPolicyScoreAttributesResponse{}
	return &this
}

// GetScore returns the Score field value.
func (o *TagPolicyScoreAttributesResponse) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributesResponse) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *TagPolicyScoreAttributesResponse) SetScore(v float64) {
	o.Score = v
}

// GetTsEnd returns the TsEnd field value.
func (o *TagPolicyScoreAttributesResponse) GetTsEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TsEnd
}

// GetTsEndOk returns a tuple with the TsEnd field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributesResponse) GetTsEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsEnd, true
}

// SetTsEnd sets field value.
func (o *TagPolicyScoreAttributesResponse) SetTsEnd(v int64) {
	o.TsEnd = v
}

// GetTsStart returns the TsStart field value.
func (o *TagPolicyScoreAttributesResponse) GetTsStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TsStart
}

// GetTsStartOk returns a tuple with the TsStart field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributesResponse) GetTsStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsStart, true
}

// SetTsStart sets field value.
func (o *TagPolicyScoreAttributesResponse) SetTsStart(v int64) {
	o.TsStart = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TagPolicyScoreAttributesResponse) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributesResponse) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TagPolicyScoreAttributesResponse) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *TagPolicyScoreAttributesResponse) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyScoreAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["score"] = o.Score
	toSerialize["ts_end"] = o.TsEnd
	toSerialize["ts_start"] = o.TsStart
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyScoreAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Score   *float64 `json:"score"`
		TsEnd   *int64   `json:"ts_end"`
		TsStart *int64   `json:"ts_start"`
		Version *int64   `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Score == nil {
		return fmt.Errorf("required field score missing")
	}
	if all.TsEnd == nil {
		return fmt.Errorf("required field ts_end missing")
	}
	if all.TsStart == nil {
		return fmt.Errorf("required field ts_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"score", "ts_end", "ts_start", "version"})
	} else {
		return err
	}
	o.Score = *all.Score
	o.TsEnd = *all.TsEnd
	o.TsStart = *all.TsStart
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
