// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelResponseAttributes Attributes of a journey funnel response.
type ProductAnalyticsJourneyFunnelResponseAttributes struct {
	// Conversion rate from the first step to the last step.
	EndToEndConversionRate float64 `json:"end_to_end_conversion_rate"`
	// Elapsed time statistics (min/max/avg in milliseconds).
	EndToEndElapsedTime ProductAnalyticsElapsedTime `json:"end_to_end_elapsed_time"`
	// The funnel steps, in the order given by the search expression.
	FunnelSteps []ProductAnalyticsJourneyFunnelStep `json:"funnel_steps"`
	// Number of entities that entered the funnel.
	InitialCount int64 `json:"initial_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelResponseAttributes instantiates a new ProductAnalyticsJourneyFunnelResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelResponseAttributes(endToEndConversionRate float64, endToEndElapsedTime ProductAnalyticsElapsedTime, funnelSteps []ProductAnalyticsJourneyFunnelStep, initialCount int64) *ProductAnalyticsJourneyFunnelResponseAttributes {
	this := ProductAnalyticsJourneyFunnelResponseAttributes{}
	this.EndToEndConversionRate = endToEndConversionRate
	this.EndToEndElapsedTime = endToEndElapsedTime
	this.FunnelSteps = funnelSteps
	this.InitialCount = initialCount
	return &this
}

// NewProductAnalyticsJourneyFunnelResponseAttributesWithDefaults instantiates a new ProductAnalyticsJourneyFunnelResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelResponseAttributesWithDefaults() *ProductAnalyticsJourneyFunnelResponseAttributes {
	this := ProductAnalyticsJourneyFunnelResponseAttributes{}
	return &this
}

// GetEndToEndConversionRate returns the EndToEndConversionRate field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetEndToEndConversionRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.EndToEndConversionRate
}

// GetEndToEndConversionRateOk returns a tuple with the EndToEndConversionRate field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetEndToEndConversionRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndToEndConversionRate, true
}

// SetEndToEndConversionRate sets field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) SetEndToEndConversionRate(v float64) {
	o.EndToEndConversionRate = v
}

// GetEndToEndElapsedTime returns the EndToEndElapsedTime field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetEndToEndElapsedTime() ProductAnalyticsElapsedTime {
	if o == nil {
		var ret ProductAnalyticsElapsedTime
		return ret
	}
	return o.EndToEndElapsedTime
}

// GetEndToEndElapsedTimeOk returns a tuple with the EndToEndElapsedTime field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetEndToEndElapsedTimeOk() (*ProductAnalyticsElapsedTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndToEndElapsedTime, true
}

// SetEndToEndElapsedTime sets field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) SetEndToEndElapsedTime(v ProductAnalyticsElapsedTime) {
	o.EndToEndElapsedTime = v
}

// GetFunnelSteps returns the FunnelSteps field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetFunnelSteps() []ProductAnalyticsJourneyFunnelStep {
	if o == nil {
		var ret []ProductAnalyticsJourneyFunnelStep
		return ret
	}
	return o.FunnelSteps
}

// GetFunnelStepsOk returns a tuple with the FunnelSteps field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetFunnelStepsOk() (*[]ProductAnalyticsJourneyFunnelStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunnelSteps, true
}

// SetFunnelSteps sets field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) SetFunnelSteps(v []ProductAnalyticsJourneyFunnelStep) {
	o.FunnelSteps = v
}

// GetInitialCount returns the InitialCount field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetInitialCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.InitialCount
}

// GetInitialCountOk returns a tuple with the InitialCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) GetInitialCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialCount, true
}

// SetInitialCount sets field value.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) SetInitialCount(v int64) {
	o.InitialCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["end_to_end_conversion_rate"] = o.EndToEndConversionRate
	toSerialize["end_to_end_elapsed_time"] = o.EndToEndElapsedTime
	toSerialize["funnel_steps"] = o.FunnelSteps
	toSerialize["initial_count"] = o.InitialCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndToEndConversionRate *float64                             `json:"end_to_end_conversion_rate"`
		EndToEndElapsedTime    *ProductAnalyticsElapsedTime         `json:"end_to_end_elapsed_time"`
		FunnelSteps            *[]ProductAnalyticsJourneyFunnelStep `json:"funnel_steps"`
		InitialCount           *int64                               `json:"initial_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EndToEndConversionRate == nil {
		return fmt.Errorf("required field end_to_end_conversion_rate missing")
	}
	if all.EndToEndElapsedTime == nil {
		return fmt.Errorf("required field end_to_end_elapsed_time missing")
	}
	if all.FunnelSteps == nil {
		return fmt.Errorf("required field funnel_steps missing")
	}
	if all.InitialCount == nil {
		return fmt.Errorf("required field initial_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_to_end_conversion_rate", "end_to_end_elapsed_time", "funnel_steps", "initial_count"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EndToEndConversionRate = *all.EndToEndConversionRate
	if all.EndToEndElapsedTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EndToEndElapsedTime = *all.EndToEndElapsedTime
	o.FunnelSteps = *all.FunnelSteps
	o.InitialCount = *all.InitialCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
