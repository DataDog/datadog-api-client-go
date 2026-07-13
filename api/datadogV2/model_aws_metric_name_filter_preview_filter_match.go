// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilterPreviewFilterMatch A metric name filter pattern and how many metrics it matched.
type AWSMetricNameFilterPreviewFilterMatch struct {
	// The number of Datadog metric names matched by this pattern.
	MatchCount int64 `json:"match_count"`
	// The metric name filter pattern.
	Pattern string `json:"pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricNameFilterPreviewFilterMatch instantiates a new AWSMetricNameFilterPreviewFilterMatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricNameFilterPreviewFilterMatch(matchCount int64, pattern string) *AWSMetricNameFilterPreviewFilterMatch {
	this := AWSMetricNameFilterPreviewFilterMatch{}
	this.MatchCount = matchCount
	this.Pattern = pattern
	return &this
}

// NewAWSMetricNameFilterPreviewFilterMatchWithDefaults instantiates a new AWSMetricNameFilterPreviewFilterMatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricNameFilterPreviewFilterMatchWithDefaults() *AWSMetricNameFilterPreviewFilterMatch {
	this := AWSMetricNameFilterPreviewFilterMatch{}
	return &this
}

// GetMatchCount returns the MatchCount field value.
func (o *AWSMetricNameFilterPreviewFilterMatch) GetMatchCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MatchCount
}

// GetMatchCountOk returns a tuple with the MatchCount field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewFilterMatch) GetMatchCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchCount, true
}

// SetMatchCount sets field value.
func (o *AWSMetricNameFilterPreviewFilterMatch) SetMatchCount(v int64) {
	o.MatchCount = v
}

// GetPattern returns the Pattern field value.
func (o *AWSMetricNameFilterPreviewFilterMatch) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewFilterMatch) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *AWSMetricNameFilterPreviewFilterMatch) SetPattern(v string) {
	o.Pattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricNameFilterPreviewFilterMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match_count"] = o.MatchCount
	toSerialize["pattern"] = o.Pattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricNameFilterPreviewFilterMatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchCount *int64  `json:"match_count"`
		Pattern    *string `json:"pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchCount == nil {
		return fmt.Errorf("required field match_count missing")
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match_count", "pattern"})
	} else {
		return err
	}
	o.MatchCount = *all.MatchCount
	o.Pattern = *all.Pattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
