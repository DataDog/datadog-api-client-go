// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsCoreUserJourney A sequence of events that users follow to use a feature.
type ProductAnalyticsCoreUserJourney struct {
	// Order of the steps, as step IDs joined by `->`.
	Expression *string `json:"expression,omitempty"`
	// Relevance score of the journey. Higher is more relevant.
	Score *int64 `json:"score,omitempty"`
	// Map of step ID to the event query of the step.
	Steps map[string]string `json:"steps,omitempty"`
	// Number of sessions that followed the journey.
	Support *int64 `json:"support,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsCoreUserJourney instantiates a new ProductAnalyticsCoreUserJourney object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsCoreUserJourney() *ProductAnalyticsCoreUserJourney {
	this := ProductAnalyticsCoreUserJourney{}
	return &this
}

// NewProductAnalyticsCoreUserJourneyWithDefaults instantiates a new ProductAnalyticsCoreUserJourney object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsCoreUserJourneyWithDefaults() *ProductAnalyticsCoreUserJourney {
	this := ProductAnalyticsCoreUserJourney{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ProductAnalyticsCoreUserJourney) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCoreUserJourney) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ProductAnalyticsCoreUserJourney) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ProductAnalyticsCoreUserJourney) SetExpression(v string) {
	o.Expression = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ProductAnalyticsCoreUserJourney) GetScore() int64 {
	if o == nil || o.Score == nil {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCoreUserJourney) GetScoreOk() (*int64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProductAnalyticsCoreUserJourney) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *ProductAnalyticsCoreUserJourney) SetScore(v int64) {
	o.Score = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductAnalyticsCoreUserJourney) GetSteps() map[string]string {
	if o == nil || o.Steps == nil {
		var ret map[string]string
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCoreUserJourney) GetStepsOk() (*map[string]string, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductAnalyticsCoreUserJourney) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given map[string]string and assigns it to the Steps field.
func (o *ProductAnalyticsCoreUserJourney) SetSteps(v map[string]string) {
	o.Steps = v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *ProductAnalyticsCoreUserJourney) GetSupport() int64 {
	if o == nil || o.Support == nil {
		var ret int64
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsCoreUserJourney) GetSupportOk() (*int64, bool) {
	if o == nil || o.Support == nil {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *ProductAnalyticsCoreUserJourney) HasSupport() bool {
	return o != nil && o.Support != nil
}

// SetSupport gets a reference to the given int64 and assigns it to the Support field.
func (o *ProductAnalyticsCoreUserJourney) SetSupport(v int64) {
	o.Support = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsCoreUserJourney) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Support != nil {
		toSerialize["support"] = o.Support
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsCoreUserJourney) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string           `json:"expression,omitempty"`
		Score      *int64            `json:"score,omitempty"`
		Steps      map[string]string `json:"steps,omitempty"`
		Support    *int64            `json:"support,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expression", "score", "steps", "support"})
	} else {
		return err
	}
	o.Expression = all.Expression
	o.Score = all.Score
	o.Steps = all.Steps
	o.Support = all.Support

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
