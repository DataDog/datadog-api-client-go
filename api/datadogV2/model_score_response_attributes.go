// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScoreResponseAttributes Score attributes.
type ScoreResponseAttributes struct {
	// The aggregation type.
	Aggregation string `json:"aggregation"`
	// Score denominator.
	Denominator int64 `json:"denominator"`
	// Score numerator.
	Numerator int64 `json:"numerator"`
	// Calculated score value.
	Score float64 `json:"score"`
	// Total number of failing rules.
	TotalFail int64 `json:"total_fail"`
	// Total number of rules with no data.
	TotalNoData int64 `json:"total_no_data"`
	// Total number of passing rules.
	TotalPass int64 `json:"total_pass"`
	// Total number of skipped rules.
	TotalSkip int64 `json:"total_skip"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScoreResponseAttributes instantiates a new ScoreResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScoreResponseAttributes(aggregation string, denominator int64, numerator int64, score float64, totalFail int64, totalNoData int64, totalPass int64, totalSkip int64) *ScoreResponseAttributes {
	this := ScoreResponseAttributes{}
	this.Aggregation = aggregation
	this.Denominator = denominator
	this.Numerator = numerator
	this.Score = score
	this.TotalFail = totalFail
	this.TotalNoData = totalNoData
	this.TotalPass = totalPass
	this.TotalSkip = totalSkip
	return &this
}

// NewScoreResponseAttributesWithDefaults instantiates a new ScoreResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScoreResponseAttributesWithDefaults() *ScoreResponseAttributes {
	this := ScoreResponseAttributes{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ScoreResponseAttributes) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ScoreResponseAttributes) SetAggregation(v string) {
	o.Aggregation = v
}

// GetDenominator returns the Denominator field value.
func (o *ScoreResponseAttributes) GetDenominator() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetDenominatorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Denominator, true
}

// SetDenominator sets field value.
func (o *ScoreResponseAttributes) SetDenominator(v int64) {
	o.Denominator = v
}

// GetNumerator returns the Numerator field value.
func (o *ScoreResponseAttributes) GetNumerator() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetNumeratorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numerator, true
}

// SetNumerator sets field value.
func (o *ScoreResponseAttributes) SetNumerator(v int64) {
	o.Numerator = v
}

// GetScore returns the Score field value.
func (o *ScoreResponseAttributes) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *ScoreResponseAttributes) SetScore(v float64) {
	o.Score = v
}

// GetTotalFail returns the TotalFail field value.
func (o *ScoreResponseAttributes) GetTotalFail() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFail
}

// GetTotalFailOk returns a tuple with the TotalFail field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetTotalFailOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFail, true
}

// SetTotalFail sets field value.
func (o *ScoreResponseAttributes) SetTotalFail(v int64) {
	o.TotalFail = v
}

// GetTotalNoData returns the TotalNoData field value.
func (o *ScoreResponseAttributes) GetTotalNoData() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalNoData
}

// GetTotalNoDataOk returns a tuple with the TotalNoData field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetTotalNoDataOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNoData, true
}

// SetTotalNoData sets field value.
func (o *ScoreResponseAttributes) SetTotalNoData(v int64) {
	o.TotalNoData = v
}

// GetTotalPass returns the TotalPass field value.
func (o *ScoreResponseAttributes) GetTotalPass() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalPass
}

// GetTotalPassOk returns a tuple with the TotalPass field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetTotalPassOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPass, true
}

// SetTotalPass sets field value.
func (o *ScoreResponseAttributes) SetTotalPass(v int64) {
	o.TotalPass = v
}

// GetTotalSkip returns the TotalSkip field value.
func (o *ScoreResponseAttributes) GetTotalSkip() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalSkip
}

// GetTotalSkipOk returns a tuple with the TotalSkip field value
// and a boolean to check if the value has been set.
func (o *ScoreResponseAttributes) GetTotalSkipOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSkip, true
}

// SetTotalSkip sets field value.
func (o *ScoreResponseAttributes) SetTotalSkip(v int64) {
	o.TotalSkip = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScoreResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["denominator"] = o.Denominator
	toSerialize["numerator"] = o.Numerator
	toSerialize["score"] = o.Score
	toSerialize["total_fail"] = o.TotalFail
	toSerialize["total_no_data"] = o.TotalNoData
	toSerialize["total_pass"] = o.TotalPass
	toSerialize["total_skip"] = o.TotalSkip

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScoreResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string  `json:"aggregation"`
		Denominator *int64   `json:"denominator"`
		Numerator   *int64   `json:"numerator"`
		Score       *float64 `json:"score"`
		TotalFail   *int64   `json:"total_fail"`
		TotalNoData *int64   `json:"total_no_data"`
		TotalPass   *int64   `json:"total_pass"`
		TotalSkip   *int64   `json:"total_skip"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Denominator == nil {
		return fmt.Errorf("required field denominator missing")
	}
	if all.Numerator == nil {
		return fmt.Errorf("required field numerator missing")
	}
	if all.Score == nil {
		return fmt.Errorf("required field score missing")
	}
	if all.TotalFail == nil {
		return fmt.Errorf("required field total_fail missing")
	}
	if all.TotalNoData == nil {
		return fmt.Errorf("required field total_no_data missing")
	}
	if all.TotalPass == nil {
		return fmt.Errorf("required field total_pass missing")
	}
	if all.TotalSkip == nil {
		return fmt.Errorf("required field total_skip missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "denominator", "numerator", "score", "total_fail", "total_no_data", "total_pass", "total_skip"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Denominator = *all.Denominator
	o.Numerator = *all.Numerator
	o.Score = *all.Score
	o.TotalFail = *all.TotalFail
	o.TotalNoData = *all.TotalNoData
	o.TotalPass = *all.TotalPass
	o.TotalSkip = *all.TotalSkip

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
